/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import java.util.Optional;
import java.util.concurrent.atomic.AtomicReference;

import com.neutralproject.lib.services.rpc.SessionToken;
import com.neutralproject.lib.services.rpc.UtcMicroTime;


public class SessionTokenManager {
  private final AtomicReference<Optional<SessionToken>> _token =
    new AtomicReference<Optional<SessionToken>>(Optional.empty());
  private final Long slackMicros = 30 * 1000L;  // how much time before token expires

  public SessionTokenManager() {
  }

  private Long nowMicros() {
    return System.currentTimeMillis() * 1000L;
  }

  public Boolean isTokenValid() {
    Optional<SessionToken> token = this._token.get();
    return token.map(t -> {
      Optional<UtcMicroTime> duration = Optional.ofNullable(t.getDuration());
      return duration
        .map(d -> d.getUtcMicroTime() > this.nowMicros() + this.slackMicros)
        .orElse(false);
    }).orElse(false);
  }

  public SessionToken updateToken(SessionToken token) {
    this._token.set(Optional.of(token));
    return token;
  }

  public SessionToken token() {
    assert(this.isTokenValid());
    return this._token.get().get();
  }
}
