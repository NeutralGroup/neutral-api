/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.configs;

import com.google.common.collect.ImmutableMap;
import com.neutralproject.external.clients.lib.NeutralClientConfig;

import java.util.Map;

public class NeutralClientConfigs {

  public static final NeutralClientConfig Local = new NeutralClientConfig(
          "localhost",
          8081,
          false,
          "http://localhost:8646",
          200000L, // gas price
          400000L // gas limit
  );

  public static final NeutralClientConfig Test = new NeutralClientConfig(
          "test-grpc-api.neutralproject.com",
          8081,
          false,
          "http://test-ethereum-api.neutralproject.com:8646",
          200000L, // gas price
          400000L // gas limit
  );

  public static final NeutralClientConfig Ropsten = new NeutralClientConfig(
          "ropsten-grpc-api.neutralproject.com",
          8081,
          false,
          "https://ropsten.infura.io/v3/c59131dc3793450ea035c087faae3772",
          200000L, // gas price
          400000L // gas limit
  );

  public static final NeutralClientConfig Production = new NeutralClientConfig(
          "production-grpc-api.neutralproject.com",
          8081,
          true,
          "https://mainnet.infura.io/v3/cdda06dcefac458d9454e05a8a14cea7",
          200000L, // gas price
          400000L, // gas limit
          "certs/production.ca"
  );

  private static Map<String, NeutralClientConfig> map = ImmutableMap.of(
          "local", Local,
          "test", Test,
          "ropsten", Ropsten,
          "production", Production
  );

  public static NeutralClientConfig forEnv(String env) {
    env = env.toLowerCase().trim();
    NeutralClientConfig config = map.get(env);
    if (null == config) {
      throw new IllegalArgumentException(String.format("Missing config for %s", env));
    }
    return config;
  }
}

