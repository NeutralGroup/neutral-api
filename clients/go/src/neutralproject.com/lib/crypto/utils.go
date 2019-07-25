/* Cryptographic Utilities
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package neutral_crypto

import (
	"github.com/golang/protobuf/proto"
	"github.com/btcsuite/btcd/btcec"
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"bytes"
	"regexp"
	"errors"
	"io/ioutil"
	"os"
	"encoding/hex"
	"encoding/base64"
	pb "neutralproject.com/lib/protocol"
	con "neutralproject.com/lib/constants"
)

//jsonpb is supposed to generate canonical json for proto3 encodings
// Yay
// no generics
// no external interface definitions on concrete types
// no external receivers
// no open classes
// no macros

func IsSigned(in interface{}, fingerprint []byte) bool {
	switch in.(type) {
	case *pb.SignedBaseQuote:
		return VerifyProtocolBuffer(fingerprint,
			in.(*pb.SignedBaseQuote).Signature,
			in.(*pb.SignedBaseQuote).Wrapped)
	case *pb.SignedAccountRequest:
		return VerifyProtocolBuffer(fingerprint,
			in.(*pb.SignedAccountRequest).Signature,
			in.(*pb.SignedAccountRequest).Wrapped)
	case *pb.SignedInstrumentResponse:
		return VerifyProtocolBuffer(fingerprint,
			in.(*pb.SignedInstrumentResponse).Signature,
			in.(*pb.SignedInstrumentResponse).Wrapped)
	case *pb.SignedUserQuote:
		return VerifyProtocolBuffer(fingerprint,
			in.(*pb.SignedUserQuote).Signature,
			in.(*pb.SignedUserQuote).Wrapped)
	case *pb.SignedFinalCommitResponse:
		return VerifyProtocolBuffer(fingerprint,
			in.(*pb.SignedFinalCommitResponse).Signature,
			in.(*pb.SignedFinalCommitResponse).Wrapped)
	default:
		log.Printf("Unhandled IsSigned class")
		return false
	}
}

func Sign(in interface{}, key *btcec.PrivateKey) interface{} {
	switch in.(type) {
	case *pb.BaseQuote:
		return &pb.SignedBaseQuote{
		Signature:
			NewSignature(SignProtocolBuffer(key, in.(*pb.BaseQuote)),
				Fingerprint(key)),
			Wrapped: in.(*pb.BaseQuote)}
	default:
		return nil
	}
}

func Fingerprint(privkey *btcec.PrivateKey) ([]byte) {
	b := sha256.Sum256(PublicUncompressed(privkey))
	return b[:]
}

func GenerateKey() (*btcec.PrivateKey) {
	key,err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatalf("Unable to generate keypair")
	}
	return key
}

func PublicUncompressed(in *btcec.PrivateKey)([]byte) {
	pubkey := in.PubKey()
	return pubkey.SerializeUncompressed()
}

// XXX Marshal returns err
func SignProtocolBuffer(key *btcec.PrivateKey, in proto.Message) ([]byte) {
	m := jsonpb.Marshaler{}
	buffer := bytes.NewBufferString("")
	m.Marshal(buffer, in)
	return SignBytes(key, buffer.Bytes())
}

func SignBytes(key *btcec.PrivateKey, in []byte) ([]byte) {
	hashedBytes := sha256.Sum256(in)
	compact,err := btcec.SignCompact(btcec.S256(), key, ([]byte)(hashedBytes[:]), true)
	if err != nil {
		log.Fatalf("Cannot sign protocol buffer: %v", err)
	}
	return compact
}


func SignBytesKeccak(key *btcec.PrivateKey, in []byte) ([]byte) {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(in)
	hash := hasher.Sum(nil)
	compact,err := btcec.SignCompact(btcec.S256(), key, hash, false)
	if err != nil {
		log.Fatalf("Cannot sign byte buffer: %v", err)
	}
	return compact
}

func VerifyProtocolBuffer(fingerprint []byte, sig *pb.Signature, in proto.Message) bool {
	if ! VerifySignatureStructure(sig) {
		return false
	}
	m := jsonpb.Marshaler{}
	buffer := bytes.NewBufferString("")
	m.Marshal(buffer, in)
        log.Printf("verifying [%s]", buffer)
	hashedBytes := sha256.Sum256(buffer.Bytes())
	return VerifySignatureBytes(fingerprint, sig.Signature, ([]byte)(hashedBytes[:]))
}

func VerifySignatureBytes(fingerprint []byte, sig []byte, hash []byte) (bool) {
	RecoveredKey, _, err := btcec.RecoverCompact(btcec.S256(), sig, hash)
	RecoveredFingerprint := sha256.Sum256(RecoveredKey.SerializeUncompressed())
	if err != nil {
		return false
	}
	return string(RecoveredFingerprint[:]) == string(fingerprint)
}

func VerifySignatureStructure(sig *pb.Signature) bool {
	if sig.Method != con.ECDSA_METHOD {
		return false
	}
	return true
}

func NewSignature(sig []byte, fingerprint []byte) *pb.Signature {
	return &pb.Signature{Signature: sig,
		Method: con.ECDSA_METHOD,
	    PubKeyFingerprint: fingerprint}
}

/* ECDSA PEM Functions
   Improvised ECDSA PEM implementation
   This probably should not be used for encoding in favor of Openssl directly
*/

func EncodePEM(privateKey *btcec.PrivateKey) (string, error) {
	priv_serial := privateKey.Serialize()
	inner := bytes.NewBufferString("302e0201010420")
	inner.WriteString(hex.EncodeToString(priv_serial))
	inner.WriteString("a00706052b8104000a")
	bits,err := hex.DecodeString(inner.String())
	if err != nil {
		return "", err
	}
	outer := bytes.NewBufferString("-----BEGIN EC PRIVATE KEY-----\n")
	outer.WriteString(base64.StdEncoding.EncodeToString(bits))
	outer.WriteString("\n-----END EC PRIVATE KEY-----\n")
	return outer.String(), nil
}

func DecodePEM(pemEncoded string) (*btcec.PrivateKey, *btcec.PublicKey, error) {
	pem_wrapping := regexp.MustCompile("-+BEGIN EC PRIVATE KEY-+\n(.*?)\n-+END EC PRIVATE KEY-+")
	if pem_wrapping.MatchString(pemEncoded) == false {
		return nil, nil, errors.New("PEM does not include correct key markers")
	}
	
	b64encoded_key := pem_wrapping.FindStringSubmatch(pemEncoded)[1]
	binstr,err := base64.StdEncoding.DecodeString(b64encoded_key)
	if (err != nil) {
		return nil, nil, err
	}
	hexstr := hex.EncodeToString(binstr)
	header_def := regexp.MustCompile("302e0201010420(.*?)a00706052b8104000a")
	if header_def.MatchString(hexstr) == false {
		return nil, nil, errors.New("PEM does not have correct EC single Private Key definition")
	}
	bytes,err := hex.DecodeString(header_def.FindStringSubmatch(hexstr)[1])
	if err != nil {
		return nil, nil, err
	}
	privkey,pubkey := btcec.PrivKeyFromBytes(btcec.S256(), bytes)
	return privkey, pubkey, nil
}

func ReadPEMFile(file_name string) (*btcec.PrivateKey, *btcec.PublicKey, error) {
	str,err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil,nil,err
	}
	return DecodePEM(string(str))
	
}

func WritePEMFile(file_name string, privkey *btcec.PrivateKey) error {
	enc_priv,err := EncodePEM(privkey)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file_name, []byte(enc_priv), os.FileMode(0400))
	if err != nil {
		return err
	}
	return nil
}

func WriteFingerprint(file_name string, privkey *btcec.PrivateKey) error {
	fingerprint := Fingerprint(privkey)
	err := ioutil.WriteFile(file_name,
		[]byte(base64.StdEncoding.EncodeToString(fingerprint[:])),
		os.FileMode(0600))
	if err != nil {
		return err
	}
	return nil
}

func ReadFingerprint(file_name string) ([]byte, error) {
	str,err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil, err
	}
	fingerprint,err := base64.StdEncoding.DecodeString(string(str))
	if err != nil {
		return nil, err
	}
	return fingerprint, nil
}
