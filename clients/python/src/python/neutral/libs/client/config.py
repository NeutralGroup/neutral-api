# -*- coding: utf-8 -*-
#
#   NeutralApiClientConfig
#
# -----------------------------------------------------------------------------

import json
import os

from neutral.libs.client.utils import NeutralUtils
# -----------------------------------------------------------------------------

class NeutralApiClientConfig(object):
    @classmethod
    def from_file(cls, path):
        with open(path, 'r') as f:
            return cls.from_json(json.loads(f.read()))

    @classmethod
    def for_env(cls, env):
        return cls.from_json(NeutralUtils.get_resource_json(
            'configs/{}/config.json'.format(env)
        ))

    @classmethod
    def from_json(cls, data):
        gateway_public_key_config = data.get('gateway-public-key', None)
        assert gateway_public_key_config, 'missing gateway-public-key'
        return cls(
            data['host'],
            data['port'],
            data.get('use-tls', False),
            gateway_public_key_config,
            data['gas-price'],
            data['gas-limit'],
            data['eth-node']
        )

    def __init__(
            self,
            host,
            port,
            use_tls,
            gateway_public_key_config,
            gas_price,
            gas_limit,
            eth_node):
        self.host = host
        self.port = port
        self.use_tls = use_tls
        self.gateway_public_key_config = gateway_public_key_config
        self.gas_price = gas_price
        self.gas_limit = gas_limit
        self.eth_node = eth_node

    def hostport(self):
        return '{}:{}'.format(self.host, self.port)

# =============================================================================
