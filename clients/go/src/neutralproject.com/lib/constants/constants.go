/* Cryptographic Utilities
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package neutral_constants

const (
	// Network Literals
	NETWORK_ETHEREUM = "ethereum"
	
	// Token Literals
	TOKEN_ETHEREUM = "ethereum"
	
	// Exchange Literals
	EXCHANGE_BINANCE = "binance"
	
	PRINT_BASEQUOTE = "base_quote"

	// Static Symbols
	SYMBOL_NEUTRAL = "NUSD"

	QUOTE_BLOCK_EXPIRY = 20
	
	ADDRESS_FEES = "0x2650cEC9F5d63A08bC0a5fCE03C2D00E58341D1e"

	ECDSA_METHOD = "ecdsa-secp256k1-sha256"
)

const (
	DeconstructNativeNative = iota 
	DeconstructVaultVault = iota
	DeconstructNativeVault = iota
	DeconstructVaultNative  = iota
	DividerStart = iota
	DividerEnd = iota     
	ConstructNativeNative = iota  
	ConstructVaultVault = iota	
	ConstructNativeVault = iota 	
	ConstructVaultNative = iota
)

var (
	TIERS = map[string]bool{
		"public": true,
		"private": true,
		"tier-1": true,
		"tier-2": true,
		"tier-3": true,
		"tier-4": true,
		"tier-5": true,
		"tier-6": true}
)
