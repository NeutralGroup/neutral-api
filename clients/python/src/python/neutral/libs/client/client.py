# -*- coding: utf-8 -*-
#
#   NeutralApiClient
#
# -----------------------------------------------------------------------------

import binascii
from google.protobuf.json_format import MessageToJson
import json
import logging

from neutral.libs.client.auth import EcdsaVerifier
from neutral.libs.client.logging import get_logger

import common_pb2
import instruments_pb2
import neutralservices_pb2
import neutralservices_pb2_grpc
import quotes_pb2

# -----------------------------------------------------------------------------

logger = get_logger(__name__)

# -----------------------------------------------------------------------------

class NeutralApiClient(object):
    '''
    Abstracts away the authentication header. Most API calls, except the two
    calls to get the private and public tokens, require authentication header
    in the meta data.

    The authentication protocol goes as follows.

    The user obtains a private or public session token, depending on whether
    he has an account with Neutral.

    The user then signs the token and attaches the token and the signature to
    the metadata in subsequent API calls made within the validity window of the
    token.
    '''
    def __init__(
            self,
            channel,
            user_auth,
            verifier,
            is_public_account=False,
            validate=True):
        if validate:
            assert verifier is not None
            assert isinstance(verifier, EcdsaVerifier)
        self.user_auth = user_auth
        self.verifier = verifier    # gateway verifier
        self.validate = validate
        self.is_public_account = is_public_account
        self.empty = common_pb2.Empty()
        self.stub = neutralservices_pb2_grpc.UserGatewayStub(channel)

    def validate_token(self, token_str):
        # validates the token sent by the gateway
        logger.info('validating session token {}'.format(token_str))
        token = json.loads(token_str)
        unsigned = token.get('unsigned', None)
        assert unsigned is not None, \
            'missing unsigned token string in {}'.format(token)
        signature = token.get('signature', None)
        assert signature is not None, \
            'missing signature string in {}'.format(token)
        self.verifier.verify_string(signature, unsigned)
        # make sure self.verifier.verify_string is not a null test
        is_a_null_test = False
        try:
            self.verifier.verify_string(signature, unsigned + 'hello world')
            is_a_null_test = True
        except:
            pass
        if is_a_null_test:
            raise Exception(
                'self.verifier.verify_string is a NULL TEST!'
            )
        logger.info('Token passed signature validation: {}'.format(token))

    def prepare_header(self):
        if self.is_public_account:
            token = self.stub.getPublicSessionToken(self.empty)
        else:
            token = self.stub.getPrivateSessionToken(self.empty)
        self.validate_token(token.token)
        return self.user_auth.gen_authentication_header(token.token)

    def _authenticated_call(self, func, *params):
        header = self.prepare_header()
        logger.debug('header = {}'.format(header))
        return func(*params, metadata=header)

    def authenticate_server_response(self, wrapped, signature):
        if not self.validate:
            return
        # The following code will not work because Python grpc lib
        # produces an incompatible JSON string. We will supply a
        # stand-alone canonicalizer binary executable (written in Go).
        sig = binascii.hexlify(signature.signature).decode('utf-8')
        preserving_proto_field_name = False
        sort_keys = False
        json_str = MessageToJson(
            wrapped,
            indent=None,
            preserving_proto_field_name=preserving_proto_field_name,
            sort_keys=sort_keys
        )
        json_obj = json.loads(json_str)
        # remove the spaces
        json_str = json.dumps(
            json_obj,
            indent=None,
            sort_keys=sort_keys,
            separators=(',', ':')
        )
        logger.debug(json_str)
        self.verifier.verify_string(sig, json_str)

    def account_info(self):
        return self._authenticated_call(self.stub.accountInfo, self.empty)

    def stream_quotes(self, instrumentID):
        assert instrumentID.symbol == 'NUSD', 'only NUSD is supported at this time'
        subscriptionRequest = quotes_pb2.SubscribeRequest(instrument=instrumentID)
        return self._authenticated_call(self.stub.subscribeToUserQuotes, subscriptionRequest)

    def execute(self, instrument_id, order_type, quantity, signed_user_quote, src_address, dst_address):
        commitment = quotes_pb2.CommitmentRecord(
            Instrument=instrument_id,
            OrderType=order_type,
            Quantity=quantity,
            Quote=signed_user_quote
        )
        request = quotes_pb2.ImmediateCommitRequest(
            SourceAccount=src_address,
            DestinationAccount=dst_address,
            Commitment=commitment
        )
        result = self._authenticated_call(self.stub.execute, request)
        self.authenticate_server_response(
            result.wrapped, result.signature
        )
        return result


    def instruments(self):
        signed = self.stub.getInstrumentDefinitions(instruments_pb2.InstrumentRequest())
        self.authenticate_server_response(
            signed.wrapped, signed.signature
        )
        return signed.wrapped.instrumentDefinitions

    def eth_address(self, hex_value):
        return common_pb2.Address(
            value = hex_value,
            addressType = "ethereum"
        )

# =============================================================================
