/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import java.math.BigInteger;
import java.util.Arrays;

import com.google.protobuf.ByteString;
import com.neutralproject.lib.services.rpc.FinalCommitResponse;
import com.lambdista.util.Try;

import org.bouncycastle.util.encoders.Hex;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.http.HttpService;
import org.web3j.protocol.Web3j;
import org.web3j.tx.gas.ContractGasProvider;


/**
 *  NeutralSettler
 *
 *  Takes the output of neutralservices.execute() and commit it to the smart
 *  contract by calling its Settle() function.
 */
public class NeutralSettler {
  private final NeutralTokenAbi abi;

  public NeutralSettler(NeutralClientConfig config, Credentials credentials) {
    System.out.println("Creating a NeutralSettler against " + config.neutralContractAddress + "@" + config.ethNode);
    this.abi = NeutralTokenAbi.load(
        config.neutralContractAddress,
        Web3j.build(new HttpService(config.ethNode)),
        credentials,
        new ContractGasProvider() {
          @Deprecated
          public BigInteger getGasPrice() {
            return new BigInteger("" + config.gasPrice);
          }

          @Deprecated
          public BigInteger getGasLimit() {
            return new BigInteger("" + config.gasLimit);
          }

          public BigInteger getGasPrice(String contractFunc) {
            return this.getGasPrice();
          }

          public BigInteger getGasLimit(String contractFunc) {
            return this.getGasLimit();
          }
        }
    );
  }

  private static BigInteger byteStringToBigInt(ByteString bs) {
    String hex = Hex.toHexString(bs.toByteArray());
    return new BigInteger("+" + hex, 16);
  }

  /**
   *  settle
   *
   *  Call NeutralValidator contract's Settle with the parameters from
   *  commitResponse.
   */
  public Try<TransactionReceipt> settle(FinalCommitResponse commitResponse) {
    System.out.println("---- trying to settle " + commitResponse);
    return Try.apply(() -> {
      try {
        byte[] signatureBytes = commitResponse.getSignature().getSignature().toByteArray();
        return this.abi.settle(
          byteStringToBigInt(commitResponse.getNonce()),
          commitResponse.getFeeDestination(),
          byteStringToBigInt(commitResponse.getExpirationDate()),
          byteStringToBigInt(commitResponse.getExpirationBlock()),
          commitResponse.getSource(),
          commitResponse.getDestination(),
          commitResponse.getInstrument(),
          byteStringToBigInt(commitResponse.getInstrumentQuantity()),
          new BigInteger("" + commitResponse.getInstrumentOperation()),
          byteStringToBigInt(commitResponse.getNeutralQuantity()),
          byteStringToBigInt(commitResponse.getFee()),
          new BigInteger("" + signatureBytes[0]),
          Arrays.copyOfRange(signatureBytes,  1, 33),
          Arrays.copyOfRange(signatureBytes, 33, 65)
        ).send();
      } catch (Throwable e) {
        e.printStackTrace();
        throw e;
      }
    });
  }
}
