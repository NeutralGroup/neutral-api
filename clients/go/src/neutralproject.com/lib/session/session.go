/*
   Neutral Session
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license.
*/

// Session management

package neutral_session

import (
	"fmt"
	"time"
	"encoding/hex"
	"google.golang.org/grpc/metadata"
	"golang.org/x/net/context"
	pb "neutralproject.com/lib/protocol"
	nc "neutralproject.com/lib/crypto"
	"github.com/btcsuite/btcd/btcec"

)

type SessionToken struct {
	private bool
	token *pb.SessionToken
	signature string
	expiration time.Time
}

func New(t *pb.SessionToken, privateKey *btcec.PrivateKey, private bool) *SessionToken {
	temp := &SessionToken{
		token: t,
		private: private}
	temp.SignSessionToken(privateKey)
	if t.Duration == nil {
		panic("Malformed session token from host.")
	}
	endDuration := time.Duration(t.Duration.UtcMicroTime) * time.Microsecond
	endDuration = endDuration - (10 * time.Second)
	temp.expiration = time.Now().Add(endDuration)
	return temp
}

func (token *SessionToken) SignSessionToken(privateKey *btcec.PrivateKey) error {
	token.signature = token.SignEth(privateKey)
	return nil
}

func (token *SessionToken) SignEth(key *btcec.PrivateKey) string {
	msg := fmt.Sprintf(
		"\x19Ethereum Signed Message:\n%d%s",
		len(token.token.Token),
		token.token.Token)
    return hex.EncodeToString(nc.SignBytesKeccak(key, []byte(msg)))
}

func (token *SessionToken) MetaData() context.Context {
    return metadata.AppendToOutgoingContext(
		context.Background(),
		"session-token", token.token.Token,
		"user-signature", token.signature)
}

func (token *SessionToken) IsExpired() bool {
	return token.expiration.Before(time.Now())
}

func (token *SessionToken) IsPrivate() bool {
	return token.private
}

