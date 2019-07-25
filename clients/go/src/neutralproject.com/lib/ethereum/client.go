/* 
   Neutral Ethereum Client
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license.
*/


package neutral_ethereum_client

import (
	"fmt"
	"log"
	"context"
	"math/big"
	"neutralproject.com/lib/ethereum/packer"
	"neutralproject.com/lib/ethereum/token"
	"neutralproject.com/lib/ethereum/mock"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethtypes "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	pb "neutralproject.com/lib/protocol"
	"github.com/btcsuite/btcd/btcec"
	"crypto/ecdsa"
)

type EthereumClient struct {
	privateKey *btcec.PrivateKey
	publicEcdsa *ecdsa.PublicKey
	fromAddress common.Address
	neutralValidatorAddress common.Address
	nonce uint64
	auth *bind.TransactOpts
	networkID *big.Int
	gasPrice *big.Int
	client *ethclient.Client
	packer *eth_packer.Packer
	neutralToken *neutral_token.NeutralToken
}


func (e *EthereumClient) updateNonce() error {
	// Nonce for Authorization
	nonce, err := e.client.PendingNonceAt(
		context.Background(),
		e.fromAddress)
	
	if err != nil {
		log.Printf("Pending Nonce Error")
		return err
	}
	e.nonce = nonce
	return nil
}

func (e *EthereumClient) updateAuth() error {
	e.updateNonce()
	e.auth.Nonce = big.NewInt(int64(e.nonce))
	e.auth.Value = big.NewInt(0)     // in wei
	e.auth.GasLimit = uint64(3000000) // in units
	e.auth.GasPrice = e.gasPrice
	return nil
}

func (e *EthereumClient) Connect(
	network_path string,
	validator_address string,
	key *ecdsa.PrivateKey) error {

	var err error
	
	e.packer, err = eth_packer.New()
	if err != nil {
		log.Printf("Unable to instantiate an ethereum packer")
		return err
	}
	e.neutralValidatorAddress = ethtypes.HexToAddress(validator_address)
	
	// Key Setup
	publicKey := key.Public()
	e.publicEcdsa = publicKey.(*ecdsa.PublicKey)
	e.fromAddress = crypto.PubkeyToAddress(*e.publicEcdsa)

	// Client Setup
	client,err := ethclient.Dial(network_path)
	if err != nil {
		log.Printf("Unable to connect to the specified Ethereum network");
		return err
	}
	e.client = client

	// Network Identifier
	networkID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Printf("Cannot query network id: %v", err)
		return err
	}
	log.Printf("Network: %s", networkID.String())
	e.networkID = networkID

	// GAS
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Unable to get suggested gas price")
		return err
	}
	e.gasPrice = gasPrice

	lastHeader,err := client.HeaderByNumber(context.Background(), nil)
	
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Last Block: %s", lastHeader.Number.String())
	
	// Authorization
	e.auth = bind.NewKeyedTransactor(key)
	e.updateAuth()
	
	// Contract setup

	e.neutralToken, err = neutral_token.NewNeutralToken(
		e.neutralValidatorAddress,
		e.client)

	if err != nil {
		log.Printf("Unable to instantiate NeutralToken contract")
		return err
	}
	
	return nil
}


func (e *EthereumClient) Execute(response *pb.SignedFinalCommitResponse) error {
	e.updateAuth()
	// We need to repack the response
	bq := response.Wrapped
	fNonce := e.packer.DecodeUint256("Nonce", bq.XNonce)
	fFeeDestination := ethtypes.HexToAddress(bq.XFeeDestination)
	log.Printf("FeeDestination: %s", fFeeDestination.String())
	fExpirationDate := e.packer.DecodeUint256("ExpirationDate", bq.XExpirationDate)
	fExpirationBlock := e.packer.DecodeUint256("ExpirationBlock", bq.XExpirationBlock)
	fSource := ethtypes.HexToAddress(bq.XSource)
	log.Printf("Source: %s", fSource.String())
	fDestination := ethtypes.HexToAddress(bq.XDestination)
	log.Printf("Destination: %s", fDestination.String())
	fInstrument := ethtypes.HexToAddress(bq.XInstrument)
	log.Printf("Instrument: %s", fInstrument.String())
	fInstrumentQuantity := e.packer.DecodeUint256("InstrumentQuantity", bq.XInstrumentQuantity)
	fInstrumentOperation := uint8(bq.XInstrumentOperation)
	log.Printf("InstrumentOperation: %d", fInstrumentOperation)
	fNeutralQuantity  := e.packer.DecodeUint256("NeutralQuantity", bq.XNeutralQuantity)
	fFee := e.packer.DecodeUint256("Fee", bq.XFee)
	fNeutralBoundary := e.packer.DecodeUint256("NeutralBoundary", bq.XNeutralBoundary)
	
	// Decode the Signature into the appropriate parameters
	/*
	v1 := bq.Signature.Signature[0:1]
	v2 := v1[0]
	r1 := bq.Signature.Signature[1:33]
	var r2 [32]byte
	for i,j := range r1 {
		r2[i] = j
	}
	s1 := bq.Signature.Signature[33:]
	var s2 [32]byte
	for i,j := range s1 {
		s2[i] = j
	}
    */
	
	tx, err := e.neutralToken.Settle(e.auth,
		fNonce,
		fFeeDestination,
		fExpirationDate,
		fExpirationBlock,
		fSource,
		fDestination,
		fInstrument,
		fInstrumentQuantity,
		fInstrumentOperation,
		fNeutralQuantity,
		fFee,
		fNeutralBoundary,
		bq.Signature.Signature)
	
	if err != nil {
		log.Fatalf("%v", err)
	}
	
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return nil
}

func (e *EthereumClient) MockSetup() error {
	e.updateAuth()
	quantity := big.NewInt(100000000)
	base := new(big.Int)
	base = base.Exp(big.NewInt(10), big.NewInt(18), nil)
	quantity = quantity.Mul(quantity, base)
	mockToken, err := detailed_mock.NewDetailedMock(
		common.HexToAddress("0xfaAa50616a81f9EFa41E44e41a05504e37c5eB3E"),
		e.client)
	tx, err := mockToken.SetBalance(e.auth,
		e.fromAddress,
		quantity)

	if err != nil {
		log.Fatalf("%v", err)
	}
	
	fmt.Printf("Token configured: %s", tx.Hash().Hex())
	return nil

	return nil
}

func (e *EthereumClient) SendEther(privkey *ecdsa.PrivateKey) error {
	e.updateAuth()

	publicKey := privkey.Public()
	fromAddress := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))
	
	var rawTx *types.Transaction
	rawTx = types.NewTransaction(
		e.nonce,
		fromAddress,
		big.NewInt(1000),
		1000000,
		e.gasPrice,
		nil)

	SignedTx, err := e.auth.Signer(types.HomesteadSigner{}, e.auth.From, rawTx)
	if err != nil {
		return err
	}

	err = e.client.SendTransaction(context.Background(), SignedTx)
	if err != nil {
		return err
	}
	return nil
}
