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

class EthContractCore(object):
    def __init__(
            self,
            contract_address,
            abi,
            eth_node,
            private_key,
            gas_price=None,
            gas_limit=400000):
        self.contract_address = contract_address.lower()
        self.eth_node = eth_node
        self.private_key = private_key
        self.account = Account.privateKeyToAccount(private_key)
        provider = Web3.HTTPProvider(eth_node)
        self.web3 = Web3(provider)
        self.gas_price = self.web3.toWei('1', 'gwei') if gas_price is None else gas_price
        self.gas_limit = gas_limit
        self.contract = self.web3.eth.contract(
            address=Web3.toChecksumAddress(contract_address),
            abi=abi
        )

    def transact(self, call):
        nonce = self.web3.eth.getTransactionCount(
            Web3.toChecksumAddress(self.account.address.lower())
        )
        txn = call.buildTransaction({
            'from': self.account.address,
            'gas': self.gas_limit,
            'gasPrice': self.gas_price,
            'nonce': nonce
        })
        print(txn)
        signed_txn = self.web3.eth.account.signTransaction(
            txn,
            private_key=self.private_key
        )
        return self.web3.eth.sendRawTransaction(signed_txn.rawTransaction)

# -----------------------------------------------------------------------------

class NeutralSpendingApprover(EthContractCore):
    def __init__(
            self,
            contract_address,
            eth_node,
            private_key,
            gas_price=None,
            gas_limit=400000):
        super().__init__(
            contract_address,
            NeutralUtils.get_resource_json('compiled-contracts/ERC20Detailed.json')['abi'],
            eth_node,
            private_key,
            gas_price=None,
            gas_limit=400000)

    def approve(self, spender, amount):
        args = (
            Web3.toChecksumAddress(spender),
            bytes_to_uint256(amount)
        )
        return self.transact(self.contract.functions.approve(*args))

# -----------------------------------------------------------------------------

class NeutralSettler(EthContractCore):

    def __init__(
            self,
            contract_address,
            eth_node,
            private_key,
            gas_price=None,
            gas_limit=400000):
        super().__init__(
            contract_address,
            NeutralUtils.get_resource_json('compiled-contracts/NeutralValidator.json')['abi'],
            eth_node,
            private_key,
            gas_price=None,
            gas_limit=400000)
        self.approvers = dict()
        '''
        self.contract_address = contract_address.lower()
        self.private_key = private_key
        self.account = Account.privateKeyToAccount(private_key)
        provider = Web3.HTTPProvider(eth_node)
        self.web3 = Web3(provider)
        self.gas_price = self.web3.toWei('1', 'gwei') if gas_price is None else gas_price
        self.gas_limit = gas_limit
        compiled = NeutralUtils.get_resource_json(
            'compiled-contracts/NeutralValidator.json'
        )
        self.contract = self.web3.eth.contract(
            address=Web3.toChecksumAddress(contract_address),
            abi=compiled['abi']
        )
        '''

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
            uint256|number _neutralBoundary
            bytes          _signature
        '''
        return (
            bytes_to_uint256(commitment._nonce),
            Web3.toChecksumAddress(commitment._feeDestination.lower()),
            bytes_to_uint256(commitment._expirationDate),
            bytes_to_uint256(commitment._expirationBlock),
            Web3.toChecksumAddress(commitment._source.lower()),
            Web3.toChecksumAddress(commitment._destination.lower()),
            Web3.toChecksumAddress(commitment._instrument.lower()),
            bytes_to_uint256(commitment._instrumentQuantity),
            int(commitment._instrumentOperation),
            bytes_to_uint256(commitment._neutralQuantity),
            bytes_to_uint256(commitment._fee),
            bytes_to_uint256(commitment._neutralBoundary),
            commitment.Signature.signature
        )

    def get_or_create_approver(self, instrument):
        instrument = instrument.lower()
        if instrument not in self.approvers:
            self.approvers[instrument] = NeutralSpendingApprover(
                instrument,
                self.eth_node,
                self.private_key,
                gas_price=self.gas_price,
                gas_limit=self.gas_limit)
        return self.approvers[instrument]

    def settle(self, commitment):
        self.get_or_create_approver(commitment._instrument).approve(commitment._source, commitment._instrumentQuantity)
        args = self._commit_to_args(commitment)
        return self.transact(self.contract.functions.settle(*args))

# =============================================================================
