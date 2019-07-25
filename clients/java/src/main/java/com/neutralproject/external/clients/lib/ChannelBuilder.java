/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import io.grpc.ManagedChannel;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.grpc.netty.NegotiationType;
import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;

import java.io.InputStream;
import java.lang.reflect.Field;
import java.lang.reflect.Modifier;
import java.util.Arrays;
import java.util.List;

public class ChannelBuilder {
  private final NeutralClientConfig config;

  public ChannelBuilder(NeutralClientConfig config) {
    this.config = config;
  }

  public ManagedChannel build() throws Exception {
    NettyChannelBuilder builder = NettyChannelBuilder.forAddress(this.config.host, this.config.port);

    if (this.config.secure) {
      if (this.config.useSelfSignedCertificate()) {
        // https://stackoverflow.com/a/47933523
        throw new RuntimeException("self-signed certificate is not supported");
      } else {
        enableProtocol();
        InputStream is = this.getClass().getClassLoader().getResourceAsStream(this.config.certificateName);
        SslContext context = GrpcSslContexts.forClient().trustManager(is).build();
        builder = builder.negotiationType(NegotiationType.TLS)
                .sslContext(context);
      }
    } else {
      builder = builder.usePlaintext();
    }
    return builder.build();
  }

  private void enableProtocol() throws Exception {
    Field field = GrpcSslContexts.class.getDeclaredField("NEXT_PROTOCOL_VERSIONS");
    List<String> newValue = Arrays.asList(null, "h2");
    field.setAccessible(true);
    Field modifiersField = Field.class.getDeclaredField("modifiers");
    modifiersField.setAccessible(true);
    modifiersField.setInt(field, field.getModifiers() & ~Modifier.FINAL);
    field.set(null, newValue);
  }
}