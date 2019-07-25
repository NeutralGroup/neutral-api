# -*- coding: utf-8 -*-
#
#   NeutralApiChannel
#
# -----------------------------------------------------------------------------

import grpc

# -----------------------------------------------------------------------------

class NeutralApiChannel(object):
    def __init__(self, hostport, credentials=None):
        self.hostport = hostport
        self.credentials = credentials

    def connect(self):
        if self.credentials:
            return grpc.secure_channel(self.hostport, self.credentials)
        else:
            return grpc.insecure_channel(self.hostport)

# =============================================================================
