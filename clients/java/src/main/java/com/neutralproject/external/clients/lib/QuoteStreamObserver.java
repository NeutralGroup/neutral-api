/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import com.neutralproject.lib.services.rpc.SignedUserQuote;
import io.grpc.stub.StreamObserver;

public class QuoteStreamObserver implements StreamObserver<SignedUserQuote> {
  private final UserQuoteProcessor processor;
  public final Completion completion = new Completion();

  public QuoteStreamObserver(UserQuoteProcessor processor) {
    this.processor = processor;
  }

  @Override
  public void onCompleted() {
    this.completion.onCompleted();
  }

  @Override
  public void onError(Throwable stone) {
    this.completion.onError(stone);
  }

  @Override
  public void onNext(SignedUserQuote signedUserQuote) {
    this.processor.process(signedUserQuote);
  }
}
