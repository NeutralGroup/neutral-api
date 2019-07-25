from unittest import TestCase
from neutral.libs.client.keys import *
from neutral.libs.client.hash import *

class EcdsaKeyTest(TestCase):

    der_private_key_config = {
        'type': 'private',
        'format': 'der',
        'value': '3074020101042097ddae0f3a25b92268175400149d65d6887b9cefaf28ea2c078e05cdc15a3c0aa00706052b8104000aa144034200047b83ad6afb1209f3c82ebeb08c0c5fa9bf6724548506f2fb4f991e2287a77090177316ca82b0bdf70cd9dee145c3002c0da1d92626449875972a27807b73b42e'
    }

    raw_private_key_config = {
        'type': 'private',
        'format': 'raw',
        'value': '97ddae0f3a25b92268175400149d65d6887b9cefaf28ea2c078e05cdc15a3c0a'
    }

    pem_private_key_config = {
        'type': 'private',
        'format': 'pem',
        'value': '''-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIJfdrg86JbkiaBdUABSdZdaIe5zvryjqLAeOBc3BWjwKoAcGBSuBBAAK
oUQDQgAEe4OtavsSCfPILr6wjAxfqb9nJFSFBvL7T5keIoencJAXcxbKgrC99wzZ
3uFFwwAsDaHZJiZEmHWXKieAe3O0Lg==
-----END EC PRIVATE KEY-----'''
    }

    private_key_configs = [der_private_key_config, raw_private_key_config, pem_private_key_config]

    der_public_key_config = {
        'type': 'public',
        'format': 'der',
        'value': '3056301006072a8648ce3d020106052b8104000a034200047b83ad6afb1209f3c82ebeb08c0c5fa9bf6724548506f2fb4f991e2287a77090177316ca82b0bdf70cd9dee145c3002c0da1d92626449875972a27807b73b42e'
    }

    raw_public_key_config = {
        'type': 'public',
        'format': 'raw',
        'value': '7b83ad6afb1209f3c82ebeb08c0c5fa9bf6724548506f2fb4f991e2287a77090177316ca82b0bdf70cd9dee145c3002c0da1d92626449875972a27807b73b42e'
    }

    pem_public_key_config = {
        'type': 'public',
        'format': 'pem',
        'value': '''-----BEGIN PUBLIC KEY-----
MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAEe4OtavsSCfPILr6wjAxfqb9nJFSFBvL7
T5keIoencJAXcxbKgrC99wzZ3uFFwwAsDaHZJiZEmHWXKieAe3O0Lg==
-----END PUBLIC KEY-----'''
    }

    public_key_configs = [der_public_key_config, raw_public_key_config, pem_public_key_config]

    def test_parsing(self):
        private_keys = [NeutralApiKey.from_dict(config) for config in self.private_key_configs]
        public_keys = [NeutralApiKey.from_dict(config) for config in self.public_key_configs]

    def test_signing_key_generator(self):
        msg = b'hello world'
        private_keys = [NeutralApiKey.from_dict(config) for config in self.private_key_configs]
        for hash_type in [HashType.SHA1, HashType.SHA256, HashType.Ethereum]:
            for key1 in private_keys:
                for key2 in private_keys:
                    signing_key = key1.get_signing_key(hash_type)
                    verifying_key = key2.get_verifying_key(hash_type)
                    signature = signing_key.sign(msg)
                    verifying_key.verify(signature, msg)
                    # make sure this is not a no-op
                    with self.assertRaises(Exception):
                        verifying_key.verify(signature, b'HELLO WORLD')

    def test_verifying_key_generator(self):
        msg = b'hello world'
        private_keys = [NeutralApiKey.from_dict(config) for config in self.private_key_configs]
        public_keys = [NeutralApiKey.from_dict(config) for config in self.public_key_configs]
        for hash_type in [HashType.SHA1, HashType.SHA256, HashType.Ethereum]:
            for key1 in private_keys:
                for key2 in public_keys:
                    signing_key = key1.get_signing_key(hash_type)
                    with self.assertRaises(ValueError):
                        # cannot get a signing key from a public key
                        key2.get_signing_key(hash_type)
                    verifying_key = key2.get_verifying_key(hash_type)
                    signature = signing_key.sign(msg)
                    verifying_key.verify(signature, msg)
                    # make sure this is not a no-op
                    with self.assertRaises(Exception):
                        verifying_key.verify(signature, b'HELLO WORLD')
