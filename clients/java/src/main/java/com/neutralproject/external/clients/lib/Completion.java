/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;
import java.util.Optional;

public class Completion {
  private final Lock lock = new ReentrantLock();
  private final Condition isDone = lock.newCondition();
  private volatile Optional<Throwable> error = Optional.empty();

  public void onError(Throwable stone) {
    this.lock.lock();
    try {
      this.error = Optional.of(stone);
      this.isDone.signal();
    } finally {
      this.lock.unlock();
    }
  }

  public void onCompleted() {
    this.lock.lock();
    try {
      this.isDone.signal();
    } finally {
      this.lock.unlock();
    }
  }
  public Optional<Throwable> getError() {
    this.lock.lock();
    try {
      return this.error;
    } finally {
      this.lock.unlock();
    }
  }

  public Optional<Throwable> waitUntilDone() throws InterruptedException {
    this.lock.lock();
    try {
      this.isDone.await();
      return this.error;
    } finally {
      this.lock.unlock();
    }
  }
}
