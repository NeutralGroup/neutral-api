---
title: Neutral gRPC API

language_tabs: # must be one of https://git.io/vQNgJ
  - proto 
  - go
  - java
  - python
  - javascript

toc_footers:

includes:
  - errors

search: true
---

# Introduction
The Neutral platform provides programatical access to real-time
pricing streams and execution service via the gRPC protocol.
The client connects to a single gRPC endpoint for all Neutral
platform services.

# API Endpoints
## Production Environment
### Production gRPC API
`production-grpc-api.neutralproject.com:8081`

### Production gRPC Web Proxy

`https://production-web-api.neutralproject.com:10000`

Neutral provides a [reverse proxy](https://github.com/improbable-eng/grpc-web/tree/master/go/grpcwebproxy) for web support.

## Ropsten Sandbox Environment

### Ropsten sandbox gRPC API
`ropsten-grpc-api.neutralproject.com:8081`

### Ropsten sandbox gRPC Web Proxy
`https://ropsten-web-api.neutralproject.com:10000`


# gRPC API Basics
The API is specified in the following files.

[`common.proto`] (./protobuf/common.proto)  contains some common data structures.

[`instruments.proto`] (./protobuf/instruments.proto) contains the instrument definition structure.

[`quotes.proto`] (./protobuf/quotes.proto) contains the data structures for pricing and execution.

[`neutralservices.proto`] (./protobuf/neutralservices.proto) contains the definitions for the API calls.

## Calls (defined in neutralservices.proto)
```
service UserGateway {
  rpc getPrivateSessionToken (Empty) returns (SessionToken) {}
  rpc getPublicSessionToken (Empty) returns (SessionToken) {}
  rpc getInstrumentDefinitions (InstrumentRequest) returns (SignedInstrumentResponse) {}
  rpc accountInfo(Empty) returns (UserAccountInfo) {}
  rpc subscribeToUserQuotes (SubscribeRequest) returns (stream SignedUserQuote) {}
  rpc execute (ImmediateCommitRequest) returns (SignedFinalCommitResponse) {}
}
```

`getPrivateSessionToken` and `getPublicSessionToken` are unauthenticated calls to obtain
a session token.

`getInstrumentDefinitions` is an unauthenticated calls to obtain the instrument definitions
for all products on the platform.

`accountInfo`, `subscribeToUserQuotes` and `execute` are all authenticated calls that
required a signed session token in the gRPC metadata, as explained below.

## Signed Responses
```
message Signature {
    string method = 1; // ecdsa-ethereum
    bytes signature = 2;
    bytes pubKeyFingerprint = 3; // optional
}

message SignedResponse {
  Signature signature = 1;
  SomeType  wrapped = 2;
}
```
Some API calls return signed responses. In addition to allowing the client to verify
the authenticity of the message, the signature is also used by the backend to detect
tampering on an execution request which is accompanied by a **signed quote** the client
has most recently received.

```python
for signed_user_quote in stub.subscribeToUserQuotes(.....):
  user_quote = signed_user_quote.wrapped
```
To obtain the actual data in a signed message, the client simply access the `wrapped` field.


## Unauthenticated Calls
```go
TODO
```
```java
TODO
```
```python
import grpc
import instruments_pb2
import neutralservices_pb2_grpc

channel = grpc.insecure_channel('ropsten-grpc-api.neutralproject.com:8081')
stub = neutralservices_pb2_grpc.UserGatewayStub(channel)
signed = self.stub.getInstrumentDefinitions(instruments_pb2.InstrumentRequest())
for definition in signed.wrapped.instrumentDefinitions:
    print(definition)
```
```javascript
TODO
```
These are `getPrivateSessionToken`, `getPublicSessionToken` and `getInstrumentDefinitions`.
A client can simply call them to obtain the results.

## Authenticated Calls: Sessions, Authentication and Accounts
```proto
message SessionToken {
    string token = 1; // opaque
    UtcMicroTime duration = 2; // Time until expiration
}
```
Most service calls exposed by the gRPC endpoint requires an authenticated session.
To authenticate itself, a client must first obtain a private or public **session token**.
The private tokens are for those clients that have accounts with Neutral; the public
tokens are for everyone else. The corresponding gRPC calls are `getPrivateSessionToken`
and `getPublicSessionToken`, respectively.

```go
TODO
```
```java
TODO
```
```python
import common_pb2
import grpc
import neutralservices_pb2_grpc

channel = grpc.insecure_channel('ropsten-grpc-api.neutralproject.com:8081')
stub = neutralservices_pb2_grpc.UserGatewayStub(channel)
token = stub.getPublicSessionToken(common_pb2.Empty())
signature = ecdsa_signer.sign_string(token.token)
stub.accountInfo(
    common_pb2.Empty(),
    metadata=[('session-token', token.token), ('user-signature', signature)]
)
```
```javascript
const { Empty } = require('./common_pb');
const { UserGatewayPromiseClient } = require('./neutralservices_grpc_web_pb');

const client = new UserGatewayPromiseClient(NEUTRAL_PROXY_URL);
const tokenResponse = await client.getPublicSessionToken(new Empty(), {});
const token = tokenResponse.getToken();

// Sign token, and reformat signature.
const s = await web3Signer.signMessage(token);
const signature = s.substr(-2) + s.substr(0, s.length - 2);
const grpcMeta = { 'session-token': token, 'user-signature': signature };
```
To authenticate itself, a client computes an Ethereum signature and attaches the token itself
(i.e., `SessionToken.token`) and the signature to the `metadata` field of subsequent gRPC
calls. The key used in computing the signature is the private key for an Ethereum
wallet.

### Accounts and Account Keys
Upon receiving an authenticated request, the Neutral backend extracts the public
key used in computing the session signature (thanks to ECDSA).
This public key uniquely identifies a private account on the Neutral Platform, when a
**private** session token is used. The private accounts entail preferred bid-ask spreads
and are set up by Neutral business development team working with our
partners. If an account cannot be located by the public key, an error (in the form of
a gRPC exception) will be returned to the client.

Note that we do not require and store a client's private key. Only the public key is
required for authentication. In next release, we may store the public key checksum or 
the Ethereum address for a wallet instead.

When a public session token is used, regardless of the signing key, subsequent requests
will be treated as from a public user who may see higher bid-ask spreads than a private
client.

# NUSD Pricing Quotes

## Streaming Quotes
```go
TODO
```
```java
TODO
```
```python
import common_pb2
import grpc
import instruments_pb2
import neutralservices_pb2_grpc
import quotes_pb

channel = grpc.insecure_channel('ropsten-grpc-api.neutralproject.com:8081')
stub = neutralservices_pb2_grpc.UserGatewayStub(channel)
signed = self.stub.getInstrumentDefinitions(instruments_pb2.InstrumentRequest())
nusd_definition = [
    definition for definition in signed.wrapped.instrumentDefinitions
    if definition.symbol == 'NUSD'
]
token = stub.getPublicSessionToken(common_pb2.Empty())
signature = ecdsa_signer.sign_string(token.token)
stub.subscribeToUserQuotes(
    quotes_pb2.SubscribeRequest(instrument=nusd_definition)
    common_pb2.Empty(),
    metadata=[('session-token', token.token), ('user-signature', signature)]
)
```
```javascript
const Common = require('./common_pb');
const Quotes = require('./quotes_pb');

// Create a subscription request object
const request = new Quotes.SubscribeRequest();

// Instrument definition
const instrument = new Common.InstrumentID();
instrument.setSymbol('NUSD');
const instrumentAddress = new Common.Address();
instrumentAddress.setValue(NUSD_ADDRESS);
instrumentAddress.setAddresstype('ethereum');
instrument.setAddress(INSTRUMENT_ADDRESS);
subRequest.setInstrument(instrument);

// Connect!
const quoteStream = await client.subscribeToUserQuotes(request, grpcMeta);
```
Neutral prices are streamed to the clients via the authenticated gRPC call `subscribeToUserQuotes`.


## Parsing Quotes
```proto
/* protobuf definition begins */
message UserQuote {
    string tier = 1; // public, tier-1, etc
    SignedBaseQuote signedBaseQuote = 2; // Signed Base Quote
    repeated SubQuote subQuotes = 3; // Quote Sub Record
    string quoteId = 4; // unique Identifier of the quote
    UtcMicroTime createdAt = 5; // User quote generated time
}

message SubQuote {
    Pair pair = 1; //
    Price bid = 2; // Bid
    Price ask = 3; //
    double minBidQuantity = 4; //
    double maxBidQuantity = 5; //
    double minAskQuantity = 6; //
    double maxAskQuantity = 7; //
}
/* protobuf definition ends */
```
Each `UserQuote` contains a list of `SubQuotes`, one for each instrument in the `NUSD`
basket. There is also a `SignedBaseQuote` from which the user quote is generated.
This field should be treated as an opaque data structure and may change in future. The
`quoteId` field uniquely identifies the `UserQuote`.


```python
for signed_user_quote in stub.subscribeToUserQuotes(...):
    user_quote = signed_user_quote.wrapped
    for sub_quote in user_quote.subQuotes:
        print('{}/{}: bid={}, ask={}'.format(
          sub_quote.pair.baseCurrency.symbol,
          sub_quote.pair.quoteCurrency.symbol,
          sub_quote.bid.value,
          sub_quote.ask.value
        ))
```
```javascript
quoteStream.on('data', (data)  => {
  // Quote Sub Record
  const sub = data.getWrapped().getSubquotesList();

  // Signed Base Quote
  const baseQuote = data.getWrapped().getSignedbasequote();

  // Quote Records
  const quotes = baseQuote.getWrapped().getSubquotesList();

  // Nav
  const nav = quotes[index].getNav();

  // Quote Details
  const bid = quotes[index].getQuote().getBid();
  const ask = quotes[index].getQuote().getAsk();
});

quoteStream.on('error', (error)  => { /* ... */ });
```

# NUSD Execution

Executing against an `NUSD` consists of two steps: Securing a price commitment and Settling against
the `NUSD` contract.

## Secure a Price Commitment
```proto
message CommitmentRecord {
    InstrumentID Instrument = 1; // ID of instrument
    string OrderType = 2; // BID or ASK
    double Quantity = 3; // Quantity
    SignedUserQuote Quote = 4; // Relayed price packet
}

message ImmediateCommitRequest {
    Address SourceAccount = 1; // Vault address of funds, msg.sender validator
    Address DestinationAccount = 2; // Vault address where funds will be deposited
    CommitmentRecord Commitment = 3; // A single commitment sub record
}

message FinalCommitResponse {
    Signature Signature = 1; // Signature Authority Signed canonical
    bytes _nonce = 2; // On chain nonce
    string _feeDestination = 3; // Fee Vault Address
    bytes _expirationDate = 4; // Expiration Date
    bytes _expirationBlock = 5; // Expiration Block
    string _source = 6; // Source of funds
    string _destination = 7; // Destination of funds
    string _instrument = 8; // Address of token
    bytes _instrumentQuantity = 9; // Amount to add or subtract
    uint32 _instrumentOperation = 10; // Addition and Subtraction
    bytes _neutralQuantity = 11; // Amount to add or subtract
    bytes _fee = 12; // Assessed spread fee in Neutral
}

message SignedFinalCommitResponse {
    Signature signature = 1;
    FinalCommitResponse wrapped = 2;
}
```
To secure a commitment, be it for buying some amount of `NUSD` using one of its constituent instruments or
selling some `NUSD` for a constituent instrument, one calls the `execute` function with an
`ImmediateCommitRequest`.

The result of the `execute` call is a `SignedFinalCommitResponse` from which the client
can extract the `FinalCommitResponse`. It is the `FinalCommitResponse` that the client
uses to create a transaction for the `NUSD` contract.

## Settle on the NUSD contract
```go
TODO
```

```java
TODO
```

```python
commitment = stub.execute(...).wrapped
signature = commitment.Signature.signature
assert len(signature) == 65
args = (
    bytes_to_uint256(commitment._nonce),
    Web3.toChecksumAddress(commitment._feeDestination.lower()),
    #re.sub('^0x', '', commitment._feeDestination),
    bytes_to_uint256(commitment._expirationDate),
    bytes_to_uint256(commitment._expirationBlock),
    Web3.toChecksumAddress(commitment._source.lower()),
    Web3.toChecksumAddress(commitment._destination.lower()),
    Web3.toChecksumAddress(commitment._instrument.lower()),
    bytes_to_uint256(commitment._instrumentQuantity),
    commitment._instrumentOperation,
    bytes_to_uint256(commitment._neutralQuantity),
    bytes_to_uint256(commitment._fee),
    signature[0],
    signature[1:33],
    signature[33:65]
)
contract = web3.eth.contract(
    address=Web3.toChecksumAddress(contract_address),
    abi=abi
)
nonce = web3.eth.getTransactionCount(
    Web3.toChecksumAddress(account.address.lower())
)
txn = contract.functions.settle(*args).buildTransaction({
    'from': account.address,
    'gas': gas,
    'gasPrice': web3.toWei('1', 'gwei'),
    'nonce': nonce
})
signed_txn = web3.eth.account.signTransaction(
    txn,
    private_key=private_key
)
web3.eth.sendRawTransaction(signed_txn.rawTransaction)
```

```javascript
TODO
```
The client constructs an Ethereum transaction against the `NUSD` contract using the
parameters provided in the `FinalCommitResponse`. The parameters including the
signature must be passed exactly as in the commit response, with appropriate
conversion to the `Ethereum` types of course.
