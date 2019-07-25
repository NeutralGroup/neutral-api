/* Ethereum Neutral ABI Packer
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package eth_packer

import (
	"log"
	"errors"
	"strings"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethtypes "github.com/ethereum/go-ethereum/common"
)

const (
	simpleAbi = `
[
  {
    "constant":false,
    "inputs":[
      {
        "name":"v",
        "type":"uint256"
      }
    ],
    "name":"uint256",
    "outputs":[
    ],
    "type":"function"
  },
  {
    "constant":false,
    "inputs":[
      {
        "name":"v",
        "type":"bool"
      }
    ],
    "name":"bool",
    "outputs":[
    ],
    "type":"function"
  },
  {
    "constant":false,
    "inputs":[
      {
        "name":"v",
        "type":"uint8"
      }
    ],
    "name":"uint8",
    "outputs":[
    ],
    "type":"function"
  },
  {
    "constant":false,
    "inputs":[
      {
        "name":"v",
        "type":"address"
      }
    ],
    "name":"address",
    "outputs":[
    ],
    "type":"function"
  },
  {
      "constant": false,
      "inputs": [
        {
          "name": "_nonce",
          "type": "uint256"
        },
        {
          "name": "_feeDestination",
          "type": "address"
        },
        {
          "name": "_expirationDate",
          "type": "uint256"
        },
        {
          "name": "_expirationBlock",
          "type": "uint256"
        },
        {
          "name": "_source",
          "type": "address"
        },
        {
          "name": "_destination",
          "type": "address"
        },
        {
          "name": "_instrument",
          "type": "address"
        },
        {
          "name": "_instrumentQuantity",
          "type": "uint256"
        },
        {
          "name": "_instrumentOperation",
          "type": "uint8"
        },
        {
          "name": "_neutralQuantity",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_neutralBoundary",
          "type": "uint256"
        }
      ],
      "name": "settle",
      "type": "function"
    }
]
`)

type Packer struct {
	packerAbi abi.ABI
}

func New() (*Packer, error) {
	packerAbi,err := abi.JSON(strings.NewReader(simpleAbi))
	if err != nil {
		return nil, err
	}
	return &Packer{packerAbi: packerAbi}, nil
}

func (p *Packer) Encode(typeName string, args ...interface{}) ([]byte, error) {

	ethType, exist := p.packerAbi.Methods[typeName]
	if !exist {
		return nil, errors.New("Could not find ethereum type")
	}
	packedArgs, err := ethType.Inputs.Pack(args...)
	if err != nil {
		return nil, err
	}
	return packedArgs, nil
}

func (p *Packer) Decode(typeName string, data []byte) ([]interface{}, error) {

	ethType, exist := p.packerAbi.Methods[typeName]
	if !exist {
		return nil, errors.New("Could not find ethereum type")
	}
	packedArgs, err := ethType.Inputs.UnpackValues(data)
	if err != nil {
		return nil, err
	}
	return packedArgs, nil
}

func (p *Packer) ToAddress(address string) ethtypes.Address {
	return ethtypes.HexToAddress(address)
}

func (p *Packer) DecodeUint256(info string, b []byte) (*big.Int) {
	v ,err := p.Decode("uint256", b)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Decode %s: %d", info, v[0].(*big.Int))
	return v[0].(*big.Int)
}
