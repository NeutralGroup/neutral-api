# -*- coding: utf-8 -*-
#
#   NeutralApiKey
#
#   Only ECDSA 256k1 is supported
#
# -----------------------------------------------------------------------------

import binascii
import ecdsa
from ecdsa.curves import SECP256k1
from enum import Enum
import json

from neutral.libs.client.hash import (Hasher, HashType)

# -----------------------------------------------------------------------------

class KeyType(Enum):
    PUBLIC = 'PUBLIC'
    PRIVATE = 'PRIVATE'

    @staticmethod
    def parse(s):
        return KeyType[s.upper()]

# -----------------------------------------------------------------------------

class KeyFormat(Enum):
    DER = 'DER'
    RAW = 'RAW'
    PEM = 'PEM'

    @staticmethod
    def parse(s):
        return KeyFormat[s.upper()]

# -----------------------------------------------------------------------------

class NeutralApiKey(object):
    def __init__(self, key_type, key_format, key_value):
        assert isinstance(key_type, KeyType)
        assert isinstance(key_format, KeyFormat)
        assert isinstance(key_value, str)
        self.key_type = key_type
        self.key_format = key_format
        self.key_value = key_value

    def get_signing_key(self, hash_type):
        hashfunc = Hasher.create(hash_type)
        if self.key_type == KeyType.PUBLIC:
            raise ValueError('this is a public key!')
        if self.key_format == KeyFormat.DER:
            return ecdsa.SigningKey.from_der(
                binascii.unhexlify(self.key_value),
                hashfunc=hashfunc
            )
        elif self.key_format == KeyFormat.RAW:
            return ecdsa.SigningKey.from_string(
                binascii.unhexlify(self.key_value),
                curve=SECP256k1,
                hashfunc=hashfunc
            )
        elif self.key_format == KeyFormat.PEM:
            return ecdsa.SigningKey.from_pem(
                self.key_value,
                hashfunc=hashfunc
            )
        else:
            raise ValueError('unsupported key format {}'.format(self.key_format))

    def get_verifying_key(self, hash_type):
        hashfunc = Hasher.create(hash_type)
        if self.key_type == KeyType.PRIVATE:
            signing_key = self.get_signing_key(hash_type)
            return signing_key.get_verifying_key()

        if self.key_format == KeyFormat.RAW:
            return ecdsa.VerifyingKey.from_string(
                binascii.unhexlify(self.key_value),
                curve=SECP256k1,
                hashfunc=hashfunc
            )
        # from_der and from_pem don't allow specifying a hash_func
        if self.key_format == KeyFormat.DER:
            key = ecdsa.VerifyingKey.from_der(
                binascii.unhexlify(self.key_value)
            )
        elif self.key_format == KeyFormat.PEM:
            key = ecdsa.VerifyingKey.from_pem(
                self.key_value
            )
        else:
            raise ValueError('unsupported key format {}'.format(self.key_format))
        # raw = binascii.hexlify(key.to_string()).decode('utf-8')
        return ecdsa.VerifyingKey.from_string(
            key.to_string(),
            curve=SECP256k1,
            hashfunc=hashfunc
        )

    @classmethod
    def from_dict(cls, d):
        return cls(
            KeyType.parse(d['type']),
            KeyFormat.parse(d['format']),
            d['value']
        )

# =============================================================================
