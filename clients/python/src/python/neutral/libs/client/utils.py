# -*- coding: utf-8 -*-
#
#   NeutralUtils
#
# -----------------------------------------------------------------------------

import os
import pkg_resources
import json

# -----------------------------------------------------------------------------

class NeutralUtils:
    '''Includes useful APIs'''

    def format_datetime(self, dt):
        return dt.strftime('%Y-%m-%d %H:%M:%S %Z%z')

    @staticmethod
    def get_resource_data(path, read_mode='r'):
        path_from_root = os.path.join('resources', path)
        try:
            # assume we are in the PEX environment
            return __loader__.get_data(path_from_root)
        except FileNotFoundError:
            # We are in the test or pants run environment, not PEX
            p = pkg_resources.resource_filename('neutral', '1234567') #'configs/{}/{}'.format(dirname, filename))
            p = os.path.dirname(os.path.dirname(p))
            p = os.path.join(p, path_from_root)
            with open(p, read_mode) as f:
                return f.read()

    @staticmethod
    def get_resource_json(path):
        data = NeutralUtils.get_resource_data(path)
        assert data
        return json.loads(data) if data else None

# =============================================================================
