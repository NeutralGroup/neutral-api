# -*- coding: utf-8 -*-
#
#   hasher
#
# -----------------------------------------------------------------------------

import binascii
from enum import Enum
from eth_account.messages import defunct_hash_message
import hashlib

# -----------------------------------------------------------------------------

class HashType(Enum):
    SHA1 = 'sha1'
    SHA256 = 'sha256'
    Ethereum = 'ethereum'

def _eth_hashser(data):
    return _EthHasher(defunct_hash_message(data))

class _EthHasher(object):
    def __init__(self, result):
        self.result = result

    def digest(self):
        return self.result

class Hasher(object):
    @staticmethod
    def create(hash_type):
        assert isinstance(hash_type, HashType), hash_type
        if hash_type == HashType.SHA1:
            return hashlib.sha1
        elif hash_type == HashType.SHA256:
            return hashlib.sha256
        elif hash_type == HashType.Ethereum:
            return _eth_hashser

# =============================================================================
