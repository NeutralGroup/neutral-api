# -*- coding: utf-8 -*-
#
#   NeutralSettler
#
# -----------------------------------------------------------------------------

import binascii
from eth_account import Account
import json
import logging
import re
from web3 import Web3

from neutral.libs.client.utils import NeutralUtils
from neutral.libs.client.logging import get_logger

# -----------------------------------------------------------------------------

logger = get_logger(__name__)

# -----------------------------------------------------------------------------

def bytes_to_uint256(b):
    return int(binascii.hexlify(b), 16)

# -----------------------------------------------------------------------------
class NeutralSettler(object):

    def __init__(
            self,
            contract_address,
            eth_node,
            private_key,
            gas=400000):
        self.contract_address = contract_address.lower()
        self.private_key = private_key
        self.account = Account.privateKeyToAccount(private_key)
        provider = Web3.HTTPProvider(eth_node)
        self.web3 = Web3(provider)
        self.gas = gas
        compiled = NeutralUtils.get_resource_json(
            'compiled-contracts/NeutralValidator.json'
        )
        self.contract = self.web3.eth.contract(
            address=Web3.toChecksumAddress(contract_address),
            abi=compiled['abi']
        )

    def is_settlement_mined(self, txn_hash):
        detail = self.web3.eth.getTransaction(txn_hash)
        logging.debug('transaction detail for {}: {}'.format(txn_hash, detail))
        receipt = self.web3.eth.getTransactionReceipt(txn_hash)
        logging.debug('transaction receipt for {}: {}'.format(txn_hash, receipt))
        return receipt is not None

    def _commit_to_args(self, commitment):
        print(commitment)
        signature = commitment.Signature.signature
        assert len(signature) == 65
        '''
        parameters for the settle function
            uint256|number _nonce
            address|string _feeDestination
            uint256|number _expirationDate
            uint256|number _expirationBlock
            address|string _source
            address|string _destination
            address|string _instrument
            uint256|number _instrumentQuantity
            uint8|number _instrumentOperation
            uint256|number _neutralQuantity
            uint256|number _fee
            uint8|string _v
            bytes32|string _r
            bytes32|string _s
        '''
        return (
            bytes_to_uint256(commitment._nonce),
            Web3.toChecksumAddress(commitment._feeDestination.lower()),
            #re.sub('^0x', '', commitment._feeDestination),
            bytes_to_uint256(commitment._expirationDate),
            bytes_to_uint256(commitment._expirationBlock),
            Web3.toChecksumAddress(commitment._source.lower()),
            Web3.toChecksumAddress(commitment._destination.lower()),
            Web3.toChecksumAddress(commitment._instrument.lower()),
            bytes_to_uint256(commitment._instrumentQuantity),
            commitment._instrumentOperation,
            bytes_to_uint256(commitment._neutralQuantity),
            bytes_to_uint256(commitment._fee),
            signature[0],
            signature[1:33],
            signature[33:65]
        )

    def settle(self, commitment):
        args = self._commit_to_args(commitment)
        print(args)
        nonce = self.web3.eth.getTransactionCount(
            Web3.toChecksumAddress(self.account.address.lower())
        )
        gas_price = self.web3.toWei('1', 'gwei')
        txn = self.contract.functions.settle(*args).buildTransaction({
            'from': self.account.address,
            'gas': self.gas,
            'gasPrice': gas_price,
            'nonce': nonce
        })
        print(txn)
        signed_txn = self.web3.eth.account.signTransaction(
            txn,
            private_key=self.private_key
        )
        return self.web3.eth.sendRawTransaction(signed_txn.rawTransaction)

# =============================================================================
