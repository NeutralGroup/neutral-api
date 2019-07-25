/*
   Neutral Account Client CLI
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license.
*/

package main

import (
    "io"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"crypto/ecdsa"
	"neutralproject.com/clients/demo/client"
        "neutralproject.com/lib/clientlib"
	"neutralproject.com/lib/configuration"
	nc "neutralproject.com/lib/crypto"
	pb "neutralproject.com/lib/protocol"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/btcsuite/btcd/btcec"
)

var (
	random = flag.Bool("random", false, "Create and setup a random account")
	withSettlement = flag.Bool("withSettlement", false, "settle transactions")
	withExecution = flag.Bool("withExecution", false, "execute")
)

func main() {
	config.LoadConfiguration(
		[]string{"user_gateway_server",
			"tls",
			"tls_certificate_file",
			"tls_key_file",
			"ca_file"})

	client := new(userGatewayClient.UserGatewayClient)
	err := clientlib.New(client,
		viper.GetString("user_gateway_server"),
		viper.GetBool("tls"),
		viper.GetString("ca_file"),
		viper.GetString("tls_certificate_file"),
		viper.GetString("tls_key_file"))

	defer client.Close()

	if err != nil {
		log.Fatalf("Connection issue: %v", err)
	}

	/* Setup the Private key used for challenge auth with Neutral
		   and in this example Ethereum network connectivity. In a production setup
		   these do not have to be the same.
	       baseAccount is used for the setup of random accounts i.e gifting Eth
	*/

	log.Printf("Environment: %s", config.Environment())

	baseAccount, _, _ := nc.ReadPEMFile(fmt.Sprintf("./keys/%s/%s.pem",
		config.Environment(),
		"client"))

	setupToken := false
	settle := false
	if *withSettlement {
	    settle = true
	}

	var privkey *btcec.PrivateKey

	if *random {
		setupToken = true
		privkey = nc.GenerateKey()
		client.Setup(privkey, baseAccount, setupToken, settle)
	} else {
		// Known preconfigured client
		privkey = baseAccount
		client.Setup(baseAccount, baseAccount, setupToken, settle)
	}

	// Private Sessions
	// get a session token
	client.GetPublicSessionToken()

    var nusdAddress *pb.Address
    var tusdAddress *pb.Address
    var daiAddress *pb.Address
    var usdcAddress *pb.Address
    var paxAddress *pb.Address

	// get instrument definitions
	definitions := client.GetInstrumentDefinitions()
	for _, v := range definitions.Wrapped.InstrumentDefinitions {
        if v.InstrumentID.Symbol == "NUSD" {
            nusdAddress = v.InstrumentID.Address
        }
        if v.InstrumentID.Symbol == "TUSD" {
            tusdAddress = v.InstrumentID.Address
        }
        if v.InstrumentID.Symbol == "DAI" {
            daiAddress = v.InstrumentID.Address
        }
        if v.InstrumentID.Symbol == "USDC" {
            usdcAddress = v.InstrumentID.Address
        }
        if v.InstrumentID.Symbol == "PAX" {
            paxAddress = v.InstrumentID.Address
        }
    }

    fmt.Printf("=========================================\n")
    fmt.Printf("NUSD Address=%v\n", nusdAddress.Value)
    fmt.Printf("TUSD Address=%v\n", tusdAddress.Value)
    fmt.Printf("PAX Address=%v\n", paxAddress.Value)
    fmt.Printf("DAI Address=%v\n", daiAddress.Value)
    fmt.Printf("USDC Address=%v\n", usdcAddress.Value)
    fmt.Printf("=========================================\n")

	stream := client.Subscribe(pb.InstrumentID{
                               			Symbol: "NUSD",
                               			Address: nusdAddress})


    publicKey := privkey.ToECDSA().Public()
    fromAddress := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))

    insts := []string{"TUSD", "DAI", "USDC", "PAX"}
    addresses := []*pb.Address{tusdAddress, daiAddress, usdcAddress, paxAddress}

    var currentIndex = 0

    waitc := make(chan struct{})
    go func() {
        for {
            res, err := stream.Recv()
            if err == io.EOF {
                close(waitc)
                return
            }
            if err != nil {
                log.Fatalf("Failed to receive subscription data : %v", err)
            }


            if *withExecution && len(res.Wrapped.SubQuotes) > 0 {
                fmt.Printf("executing for %v (%v) - CreatedAt=%v\n", insts[currentIndex],addresses[currentIndex], res.Wrapped.SignedBaseQuote.Wrapped.CreatedAt.UtcMicroTime)
                commitment := client.GenerateCommitRequest(
                    pb.Address{
                        Value:       fromAddress.String(),
                        AddressType: "ethereum"},
                    pb.Address{
                        Value:       "0x1a63c94aae69d28ddd810e4e9e3113eeecfcfe2d",
                        AddressType: "ethereum"},
                    pb.InstrumentID{
                        Symbol: insts[currentIndex],
                        Address: addresses[currentIndex]},
                    "BID",
                    1,
                    res)

                client.Execute(commitment)
                fmt.Printf("\n======================================\n")
                fmt.Printf("execution success\n")
                fmt.Printf("======================================\n")
                currentIndex = (currentIndex + 1) % 4
            } else {
                fmt.Printf("=========================================\n")
                fmt.Printf("Received a signed user quote: \n %v\n", client.ProtoString(res))
                fmt.Printf("=========================================\n")
            }
        }
    }()
    stream.CloseSend()
    <-waitc
}

