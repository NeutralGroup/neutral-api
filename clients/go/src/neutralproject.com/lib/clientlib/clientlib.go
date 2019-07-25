/*
   Neutral Account Management Client
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license.
*/


package clientlib

import (
	"log"
	"net/url"
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"
	"google.golang.org/grpc/credentials"

)

type NeutralClient interface {
	New(connection *grpc.ClientConn)
	Close() error
}

func New(nc NeutralClient,
	connection_path string,
	tls bool,
	ca_file string,
	tls_certificate_file string,
	tls_key_file string) error {
	var opts []grpc.DialOption
	if tls {
		if ca_file == "" {
			ca_file = testdata.Path("ca.pem")
		}

		path,err := url.Parse(connection_path)
		if err != nil {
			log.Fatal(err)
		}
		creds, err := credentials.NewClientTLSFromFile(ca_file, path.Hostname())
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(connection_path, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	nc.New(conn)
	return nil	
}
