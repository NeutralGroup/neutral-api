# -*- coding: utf-8 -*-
#
#   Neutral Service gRPC Demo Client
#
# -----------------------------------------------------------------------------

import argparse
import binascii
import grpc
from grpc._channel import _Rendezvous
import logging
import random
import re
import sys

from eth_account import Account
from neutral.libs.client.auth import (
    Web3EcdsaSigner,
    EcdsaVerifier,
    UserAuthentication
)
from neutral.libs.client.logging import (set_logging_level, get_logger)
from neutral.libs.client.channel import NeutralApiChannel
from neutral.libs.client.client import NeutralApiClient
from neutral.libs.client.config import NeutralApiClientConfig
from neutral.libs.client.hash import HashType
from neutral.libs.client.keys import NeutralApiKey
from neutral.libs.client.settler import NeutralSettler
from neutral.libs.client.utils import NeutralUtils

import common_pb2
import instruments_pb2
import neutralservices_pb2
import neutralservices_pb2_grpc
import quotes_pb2

# -----------------------------------------------------------------------------

logger = logging.getLogger(__name__)

# -----------------------------------------------------------------------------
def read_pem_file(path): 
    with open(path, 'rt') as pemfile:
        return pemfile.read()

# -----------------------------------------------------------------------------

def parse_args():
    parser = argparse.ArgumentParser(description='Neutral Pricing Service Client.')
    parser.add_argument(
        '--config-file-override',
        dest='config_file_override',
        type=str,
        default=None,
        help='override the pre-compiled config in the resources'
    )
    parser.add_argument(
        '--env',
        dest='env',
        type=str,
        default='local',
        help='service environment to run against'
    )
    parser.add_argument(
        '-P', '--is-private-account',
        dest='is_private_account',
        action='store_true'
    )
    parser.add_argument(
        '-e', '--execute',
        dest='execute_every_n_quotes',
        type=int,
        default=0,
        help='obtain an execution authorization on every N quotes'
    )
    parser.add_argument(
        '--log-level',
        dest='loglevel',
        type=str,
        default='INFO',
        help='set the logger level'
    )

    '''
    0x119c26b62c2a28546382db56f223f42f59d10759db72a4cae135be54b4565d5d
    is the private key for Ganache test address
    0xf1f0862711fd40259522cabfa4f99e9c80d61f18
    '''
    parser.add_argument(
        '-k', '--src-private-key',
        dest='src_private_key',
        type=str,
        default='119c26b62c2a28546382db56f223f42f59d10759db72a4cae135be54b4565d5d',
        help='Private key for src address'
    )
    parser.add_argument(
        '--dst-address',
        dest='dst_address',
        type=str,
        default=None, # use the src address
        help='Ethereum address for the destination of the transaction; default is src address'
    )
    parser.add_argument(
        '-S', '--settle-the-commitment',
        dest='settle',
        action='store_true',
        help='settle the commitment to the Neutral smart contract'
    )
    return parser.parse_args()

# -----------------------------------------------------------------------------

class DemoTrader(object):

    def __init__(self, client, contract_address, config, args):
        self.client = client
        self.args = args
        self.src_account = Account.privateKeyToAccount(args.src_private_key)
        self.settler = NeutralSettler(
            contract_address,
            config.eth_node,
            args.src_private_key,
            gas_price=config.gas_price,
            gas_limit=config.gas_limit
        )
        self.quotes_seen = 0
        # we do one settlement at a time
        self.previous_transaction_hash = None
        self.logger = get_logger('DemoTrader')

    def previous_transaction_mined(self):
        if self.previous_transaction_hash is None:
            return True
        return self.settler.is_settlement_mined(self.previous_transaction_hash)

    def on_quote(self, signed_user_quote):
        self.client.authenticate_server_response(
            signed_user_quote.wrapped, signed_user_quote.signature
        )
        quote = signed_user_quote.wrapped
        self.logger.info('received a quote')
        self.logger.debug(quote)
        if self.args.execute_every_n_quotes <= 0:
            return
        self.quotes_seen += 1
        if self.quotes_seen < self.args.execute_every_n_quotes:
            return
        if not self.previous_transaction_mined():
            self.logger.warn('previous transaction {} is not mined yet'.format(
                binascii.hexlify(self.previous_transaction_hash)))
            return
        try:
            self.execute(signed_user_quote)
        except ValueError as e:
            if 'replacement transaction underpriced' in str(e):
                # the program might have been killed and restarted immediately
                self.logger.warn('previous transaction (unknown) is not mined yet')
                return
            raise
        except _Rendezvous as e:
            if 'Commitment is Expired' in e.details():
                self.logger.warn('Commitment is Expired. Will retry on next quote')
                return
            raise
        self.quotes_seen = 0

    def execute(self, signed_user_quote):
        quote = signed_user_quote.wrapped
        # We only collateralize because the basket is most
        # likely empty during testing. A real app will check
        # balances in the contract and in the account and
        # decide whether to buy or sell.
        subquote = random.choice([
            subquote for subquote in quote.subQuotes
            if subquote.minAskQuantity > 0
        ])
        instrument = subquote.pair.quoteCurrency
        # BID hits the ASK
        order_type = 'BID'  # BUY NUSD using instrument
        quantity = subquote.minAskQuantity
        src_address = self.client.eth_address(self.src_account.address)
        dst_address = self.client.eth_address(
            self.args.dst_address if self.args.dst_address
            else self.src_account.address)
        self.logger.info(
            '{} {} {}'.format(order_type, quantity, instrument)
        )
        commitment = self.client.execute(
            instrument,
            order_type,
            quantity,
            signed_user_quote,
            src_address,
            dst_address
        )
        self.logger.info('received the commitment')
        self.logger.debug(commitment)
        if self.args.settle:
            self.logger.info('settling the commitment ......')
            txn_hash = self.settler.settle(commitment.wrapped)
            self.logger.info('settlement transaction hash = {}'.format(
                binascii.hexlify(txn_hash)))
            if txn_hash:
                self.previous_transaction_hash = txn_hash

# -----------------------------------------------------------------------------

def main():
    args = parse_args()
    loglevel = getattr(logging, args.loglevel.upper(), logging.INFO)
    set_logging_level(loglevel)
    logger = get_logger(__name__)

    # Validation requires a canonicalizer service; not available at this moment.
    # Python grpc code produces an incompatible JSON string.
    args.validate_server_response = False

    if args.config_file_override:
        config = NeutralApiClientConfig.from_file(args.config_file_override)
    else:
        config = NeutralApiClientConfig.for_env(args.env)

    print('connecting to {} ......'.format(config.hostport()))
    hostport = config.hostport()
    if config.use_tls:
        credentials = grpc.ssl_channel_credentials(
            root_certificates=NeutralUtils.get_resource_data('certs/production.ca')
        )
    else:
        credentials = None
    gateway_key = NeutralApiKey.from_dict(
        config.gateway_public_key_config
    ).get_verifying_key(HashType.Ethereum)
    verifier = EcdsaVerifier(gateway_key)

    print('src private key = {}'.format(args.src_private_key))
    user_key = NeutralApiKey.from_dict({
        'type': 'private',
        'format': 'raw',
        'value': args.src_private_key
    }).get_signing_key(HashType.Ethereum)
    signer = Web3EcdsaSigner(user_key, HashType.Ethereum)
    user_auth = UserAuthentication(signer)

    with NeutralApiChannel(hostport, credentials).connect() as channel:
        client = NeutralApiClient(
            channel, user_auth, verifier,
            is_private_account=args.is_private_account,
            validate=args.validate_server_response
        )

        logger.info('getting account info ......')
        account_info = client.account_info()
        logger.info('received account info: {}'.format(account_info))

        instruments = client.instruments()
        filtered = [
            inst for inst in instruments
            if inst.instrumentID.symbol == 'NUSD'
        ]
        assert len(filtered) == 1, \
            'missing NUSD in instruments {}'.format(instruments)
        NUSD = filtered[0]
        contract_address = NUSD.instrumentID.address.value
        trader = DemoTrader(client, contract_address, config, args)
        for signed_user_quote in client.stream_quotes(NUSD.instrumentID):
            trader.on_quote(signed_user_quote)

# =============================================================================
