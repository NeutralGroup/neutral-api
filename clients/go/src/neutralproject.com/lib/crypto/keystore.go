/* Cryptographic Utilities
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

// Basic keystore interface for managing readonly keys in memory

package neutral_crypto

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
        "log"
	"strings"
	"path/filepath"
	"github.com/btcsuite/btcd/btcec"
)

type Keystore struct {
	fingerprints map[string][]byte
	keys map[string]*btcec.PrivateKey
}

func (ks *Keystore) Load(directory string) error {
        log.Printf("loading keystore from %s ......", directory)

	if ks.fingerprints == nil {
		ks.fingerprints = make(map[string][]byte)
	}

	if ks.keys == nil {
		ks.keys = make(map[string]*btcec.PrivateKey)
	}
	
	files,err := filepath.Glob(directory + "/*.fingerprint")
	if err != nil {
		return err
	}
	for _, fingerprint_file := range files {
		fingerprint,err := ReadFingerprint(fingerprint_file)
		if err != nil {
			return err
		}
		fingerprint_file := filepath.Base(fingerprint_file)
		ks.fingerprints[strings.TrimSuffix(fingerprint_file,
			filepath.Ext(fingerprint_file))] = fingerprint
	}
	
	files,err = filepath.Glob(directory + "/*.pem")
	for _,privkey_file := range files {
		privkey, _, err := ReadPEMFile(privkey_file)
		if err != nil {
			return err
		}
		privkey_file := filepath.Base(privkey_file)
		ks.keys[strings.TrimSuffix(privkey_file, filepath.Ext(privkey_file))] = privkey
	}
	return nil
	}

func (ks *Keystore) GetFingerprint(name string) ([]byte, error) {
	if ks.fingerprints[name] == nil {
		return nil, status.Error(codes.NotFound, "Fingerprint missing")
	}
	return ks.fingerprints[name], nil
}

func (ks *Keystore) GetPrivkey(name string) (*btcec.PrivateKey, error) {
	if ks.keys[name] == nil {
		return nil, status.Error(codes.NotFound, "Privkey missing")
	}
	return ks.keys[name], nil
}

// Generate an ephemeral key
func (ks *Keystore) Generate(name string) error {
	p := GenerateKey()
	ks.keys[name] = p
	return nil
}
