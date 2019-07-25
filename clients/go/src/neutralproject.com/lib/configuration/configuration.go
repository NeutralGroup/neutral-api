/* Neutral API Gateway
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

// Common protocol definitions in definitions/


package config

import (
	"log"
	"flag"
	"github.com/spf13/viper"
)

var (
	environment = flag.String("environment", "test", "The Neutral Deployment Environment")
)

func Environment() (string) {
	return *environment
}

func LoadConfiguration(params []string) {
	flag.Parse()
	viper.SetConfigName(*environment)
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}
	
	// Verify the existence of every parameter we care about
	for _, x := range params {
		if viper.IsSet(x) != true {
			log.Fatalf("Configuration missing parameter %s", x)
		}
	}
}
