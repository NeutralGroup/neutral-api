/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import java.util.concurrent.Callable;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.TimeUnit;
import java.util.Optional;

import com.neutralproject.lib.services.rpc.*;
import com.neutralproject.services.rpc.UserGatewayGrpc;
import com.neutralproject.services.rpc.UserGatewayGrpc.UserGatewayFutureStub;
import com.neutralproject.services.rpc.UserGatewayGrpc.UserGatewayStub;

import com.google.common.util.concurrent.ListenableFuture;
import com.google.common.util.concurrent.ListeningExecutorService;
import com.google.common.util.concurrent.MoreExecutors;
import com.google.common.util.concurrent.ThreadFactoryBuilder;

import io.grpc.ManagedChannel;
import io.grpc.Metadata;
import io.grpc.stub.MetadataUtils;
import io.grpc.stub.StreamObserver;

public class NeutralClient {
  private final UserGatewayStub stub; // for streaming calls
  private final UserGatewayFutureStub futureStub; // for unary calls
  private final EcdsaSigner signer;
  private final SessionTokenManager sessionTokenManager;
  private final Empty empty = Empty.newBuilder().build();

  public enum CommitType {
    Sell, // sell the collateral for NUSD
    Buy   // buy the collateral using NUSD
  }

  public NeutralClient(ManagedChannel channel, EcdsaSigner signer) {
    this.stub = UserGatewayGrpc.newStub(channel);
    this.futureStub = UserGatewayGrpc.newFutureStub(channel);
    this.signer = signer;
    this.sessionTokenManager = new SessionTokenManager();
  } 

  public SessionToken getValidSessionToken() throws Throwable {
    if (this.sessionTokenManager.isTokenValid()) {
      return this.sessionTokenManager.token();
    }
    // retrieve one
    ListenableFuture<SessionToken> tokenFuture = this.futureStub.getPrivateSessionToken(this.empty);
    // tokenFuture.addListener(() -> (), MoreExecutors.directExecutor());
    // Futures.addCallback(tokenFuture, new FutureCallback<CreateResponse>() { }
    SessionToken token = tokenFuture.get(30, TimeUnit.SECONDS);
    this.sessionTokenManager.updateToken(token);
    return token;
  }

  private Metadata tokenToHeader(SessionToken token) {
    String tokenString = token.getToken();
    String signature = this.signer.sign(tokenString);
    Metadata header=new Metadata();
    header.put(
      Metadata.Key.of("session-token", Metadata.ASCII_STRING_MARSHALLER),
      tokenString
    );
    header.put(
      Metadata.Key.of("user-signature", Metadata.ASCII_STRING_MARSHALLER),
      signature
    );
    return header;
  }

  private UserGatewayStub getStubForAuthenticatedCall() throws Throwable {
    SessionToken token = this.getValidSessionToken();
    return MetadataUtils.attachHeaders(this.stub, this.tokenToHeader(token));
  }

  private UserGatewayFutureStub getFutureStubForAuthenticatedCall() throws Throwable {
    SessionToken token = this.getValidSessionToken();
    return MetadataUtils.attachHeaders(this.futureStub, this.tokenToHeader(token));
  }

  public UserAccountInfo accountInfo() throws Throwable {
    ListenableFuture<UserAccountInfo> accountInfoFuture =
      this.getFutureStubForAuthenticatedCall().accountInfo(this.empty);
    return accountInfoFuture.get(30, TimeUnit.SECONDS);
  }

  public InstrumentResponse getInstrumentDefinitions() throws Throwable {
    ListenableFuture<SignedInstrumentResponse> signedResponseFuture =
      this.getFutureStubForAuthenticatedCall().getInstrumentDefinitions(
          InstrumentRequest.newBuilder().build()
      );
    SignedInstrumentResponse signedResponse = signedResponseFuture.get(30, TimeUnit.SECONDS);
    return signedResponse.getWrapped();
  }

  public FinalCommitResponse getExecutableRecord(
      CommitType commitType,
      InstrumentID coin,
      Double quantity,
      Address withdrawAddress,
      Address depositAddress,
      SignedUserQuote signedUserQuote) throws Throwable {

    CommitmentRecord commitment = CommitmentRecord
      .newBuilder()
      .setOrderType(commitType == CommitType.Sell ? "BID" : "ASK")
      .setInstrument(coin)
      .setQuantity(quantity)
      .setQuote(signedUserQuote)
      .build();

    ImmediateCommitRequest request = ImmediateCommitRequest
      .newBuilder()
      .setSourceAccount(withdrawAddress)
      .setDestinationAccount(depositAddress)
      .setCommitment(commitment)
      .build();

    ListenableFuture<SignedFinalCommitResponse> signedCommitResponseFuture =
      this.getFutureStubForAuthenticatedCall().execute(request);
    SignedFinalCommitResponse signedCommitResponse = signedCommitResponseFuture.get(30, TimeUnit.SECONDS);
    return signedCommitResponse.getWrapped();
  }

  public Completion subscribeToUserQuotes(InstrumentID instrument, UserQuoteProcessor processor) throws Throwable {
    UserGatewayStub stub = this.getStubForAuthenticatedCall();
    SubscribeRequest request = SubscribeRequest
      .newBuilder()
      .setInstrument(instrument)
      .build();
    QuoteStreamObserver observer = new QuoteStreamObserver(processor);
    stub.subscribeToUserQuotes(request, observer);
    return observer.completion;
  }

  public static NeutralClient build(NeutralClientConfig config, EcdsaSigner signer) throws Exception {
    return new NeutralClient((new ChannelBuilder(config)).build(), signer);
  }
}
