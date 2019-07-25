/*
   Neutral Account Management Client
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license.
*/

package userGatewayClient

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	nc "neutralproject.com/lib/crypto"
	ne "neutralproject.com/lib/ethereum"
	pb "neutralproject.com/lib/protocol"
	session "neutralproject.com/lib/session"
	"time"
)

// implements clientlib/NeutralClient interface

type UserGatewayClient struct {
	connection     *grpc.ClientConn
	client         pb.UserGatewayClient
	keyStore       nc.Keystore
	baseClient     *ne.EthereumClient
	ethereumClient *ne.EthereumClient
	sessionToken   *session.SessionToken
	privateKey     *btcec.PrivateKey
	baseKey        *btcec.PrivateKey
	withSettlement bool
}

func (cn *UserGatewayClient) Close() error {
	return cn.connection.Close()
}

func (cn *UserGatewayClient) New(connection *grpc.ClientConn) {
	cn.connection = connection
	cn.client = pb.NewUserGatewayClient(connection)
}

func (cn *UserGatewayClient) Setup(
	privkey *btcec.PrivateKey,
	basekey *btcec.PrivateKey,
	tokenSetup bool,
	withSettlement bool) error {

	cn.privateKey = privkey
	cn.baseKey = basekey

	// If tokenSetup is true we will do additioanl initialization to create
	// default balances in a test or sandbox environment.

    cn.withSettlement = withSettlement

    if withSettlement {
        cn.ethereumClient = new(ne.EthereumClient)
        cn.ethereumClient.Connect(
            viper.GetString("ethereum_network"),
            viper.GetString("contract_address"),
            privkey.ToECDSA())
    } else {
        log.Printf("running without settlement - eth client not initialised\n")
    }

	if tokenSetup {
		cn.baseClient = new(ne.EthereumClient)
		cn.baseClient.Connect(
			viper.GetString("ethereum_network"),
			viper.GetString("contract_address"),
			basekey.ToECDSA())

		cn.baseClient.MockSetup()
		err := cn.baseClient.SendEther(privkey.ToECDSA())
		if err != nil {
			log.Fatalf("Problem seeding account: %v", err)
		}
	}
	return nil
}

func (client *UserGatewayClient) ProtoString(in proto.Message) string {
	buf := new(bytes.Buffer)
	m := jsonpb.Marshaler{}
	m.Marshal(buf, in)
	return buf.String()
}

func (client *UserGatewayClient) RenewToken() {
	if client.sessionToken.IsPrivate() {
		client.GetPrivateSessionToken()
	} else {
		client.GetPublicSessionToken()
	}
}

func (client *UserGatewayClient) GetPublicSessionToken() *session.SessionToken {
    clientDeadline := time.Now().Add(time.Duration(20 * 1000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := client.client.GetPublicSessionToken(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("could not get public session token: %v", err)
	}
	client.sessionToken = session.New(r, client.privateKey, false)
	return client.sessionToken
}

func (client *UserGatewayClient) GetPrivateSessionToken() *session.SessionToken {
	clientDeadline := time.Now().Add(time.Duration(20 * 1000) * time.Millisecond)
    ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()
	r, err := client.client.GetPrivateSessionToken(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("could not get private session token: %v", err)
	}
	client.sessionToken = session.New(r, client.privateKey, true)
	return client.sessionToken
}

func (client *UserGatewayClient) GetInstrumentDefinitions() *pb.SignedInstrumentResponse {
	ctx := client.sessionToken.MetaData()
	r, err := client.client.GetInstrumentDefinitions(ctx, &pb.InstrumentRequest{})

	if err != nil {
		log.Fatalf("failed to get instrument definitions: %v", err)
	}

    return r
}

func (client *UserGatewayClient) Subscribe(instrument pb.InstrumentID) pb.UserGateway_SubscribeToUserQuotesClient {
	ctx := client.sessionToken.MetaData()
	stream, err := client.client.SubscribeToUserQuotes(ctx,
		&pb.SubscribeRequest{Instrument: &instrument})

	if err != nil {
		log.Fatalf("Subscribe failed: %v", err)
	}
	return stream
}

func (client *UserGatewayClient) Execute(
	commitRequest pb.ImmediateCommitRequest) {

	if client.sessionToken.IsExpired() {
		log.Printf("Token is expired, renewing...")
		client.RenewToken()
	}

	ctx := client.sessionToken.MetaData()

    reqId := commitRequest.Commitment.Quote.Wrapped.QuoteId

	r, err := client.client.Execute(ctx, &commitRequest)
	if err != nil {
		errStatus, _ := status.FromError(err)
        fmt.Printf("executing failed for quote id = %v\n", reqId)
		log.Fatal(fmt.Println(errStatus.Message()))
	} else {
		// We have the FinalSignedCommitResponse from the execution service
		// We need to verify, pack it for Ethereum and submit it to the network

		if client.withSettlement {
		    err = client.ethereumClient.Execute(r)
		    if err != nil {
		    	log.Fatalf("Unable to execute on ethereum network: %v", err)
		    }
		}
	}
}

func (client *UserGatewayClient) GenerateCommitRequest(
    sourceAccount pb.Address,
    destAccount pb.Address,
    instrument pb.InstrumentID,
    orderType string,
    quantity float64,
    quote *pb.SignedUserQuote) pb.ImmediateCommitRequest {

    return pb.ImmediateCommitRequest{
                               SourceAccount: &sourceAccount,
                               DestinationAccount: &destAccount,
                               Commitment: &pb.CommitmentRecord{
                                   Instrument: &instrument,
                                   OrderType: orderType,
                                   Quantity:  quantity,
                                   Quote:     quote},
                           }

}

