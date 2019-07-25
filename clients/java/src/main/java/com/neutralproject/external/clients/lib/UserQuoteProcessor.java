/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import com.neutralproject.lib.services.rpc.SignedUserQuote;

public interface UserQuoteProcessor {
  public void process(SignedUserQuote signedUserQuote);
}
