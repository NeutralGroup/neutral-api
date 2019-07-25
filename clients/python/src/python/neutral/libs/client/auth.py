# -*- coding: utf-8 -*-
#
#   auth
#
# -----------------------------------------------------------------------------

import binascii
import ecdsa
import re
from web3.auto import w3

from neutral.libs.client.hash import HashType, Hasher

# -----------------------------------------------------------------------------

class EcdsaVerifier(object):
    '''Generic ECDSA verifier using python-ecdsa.'''

    def __init__(self, verifying_key):
        assert isinstance(verifying_key, ecdsa.VerifyingKey)
        self.verifying_key = verifying_key

    def verify_bytes(self, sig, data):
        assert isinstance(sig, str), sig
        sig = binascii.unhexlify(str(sig))
        if len(sig) == 65:
            sig = sig[1:]
        elif len(sig) != 64:
            raise ValueError('invalid signature length {}'.format(sig))
        self.verifying_key.verify(sig, data)

    def verify_string(self, sig, text, encoding='utf-8'):
        self.verify_bytes(sig, bytes(text, encoding=encoding))

# -----------------------------------------------------------------------------

class Web3EcdsaSigner(object):
    '''
    Web3 adds a recovery byte to the signature. We have to use this to ensure
    correct extraction of the public key in the backend (since we are not
    transmitting the public key over.
    '''

    def __init__(self, signing_key, hash_type):
        assert isinstance(signing_key, ecdsa.SigningKey)
        assert isinstance(hash_type, HashType)

        self.signing_key = signing_key
        self.private_key = signing_key.to_string()
        self.hash_type = hash_type
        self.hasher = Hasher.create(hash_type)

    def get_hash_type():
        return self.hash_type

    def sign_bytes(self, data):
        return self.sign_hash(self.hasher(data).digest())

    def sign_string(self, text, encoding='utf-8'):
        return self.sign_bytes(bytes(text, encoding=encoding))

    '''
    The backend uses (v + r + s)
    w3.eth.account.signHash returns (r + s + v)
    '''
    def sign_hash(self, data_hash):
        # print('hash ({}}: {}'.format(type(data_hash), data_hash))
        signed = w3.eth.account.signHash(data_hash, private_key=self.private_key)
        signature = re.sub(r'^0x', '', signed.signature.hex())
        signature = signature[-2:] + signature[:-2]
        return signature

# -----------------------------------------------------------------------------

class UserAuthentication(object):
    def __init__(self, signer):
        self.signer = signer
        verifying_key = signer.signing_key.get_verifying_key()
        self.verifier = EcdsaVerifier(verifying_key)

    def gen_authentication_header(self, token_str):
        signature = self.signer.sign_string(token_str)
        print('token: {}'.format(token_str))
        print('signature: {}'.format(signature))
        self.verifier.verify_string(
            signature[2:],  # strip the recovery byte
            token_str
        )
        return [
            ('session-token', token_str),
            ('user-signature', signature),
        ]

# =============================================================================
