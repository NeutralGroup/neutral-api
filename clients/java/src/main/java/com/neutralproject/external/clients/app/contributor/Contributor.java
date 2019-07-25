/* Neutral Java Client
   Copyright 2018, Neutral Empire Group, Inc 228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.app.contributor;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.Map;

import com.lambdista.util.Try;

import com.neutralproject.external.clients.configs.NeutralClientConfigs;
import com.neutralproject.external.clients.lib.Completion;
import com.neutralproject.external.clients.lib.EcdsaSigner;
import com.neutralproject.external.clients.lib.NeutralClient;
import com.neutralproject.external.clients.lib.NeutralClient.CommitType;
import com.neutralproject.external.clients.lib.NeutralClientConfig;
import com.neutralproject.external.clients.lib.NeutralSettler;
import com.neutralproject.external.clients.lib.UserQuoteProcessor;
import com.neutralproject.lib.services.rpc.*;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.core.methods.response.TransactionReceipt;

public class Contributor {
  private static NeutralClientConfig getConfig() {
    Map<String, String> map = System.getenv();
    String env = map.get("NTL_ENV");
    return env == null ? NeutralClientConfigs.Local : NeutralClientConfigs.forEnv(env);
  }

  private static InstrumentDefinition getInstrumentDefinition(String symbol, InstrumentResponse instrumentResponse) {
    return instrumentResponse
      .getInstrumentDefinitionsList()
      .stream()
      .filter(def -> def.hasInstrumentID())
      .filter(def -> def.getInstrumentID().getSymbol().equalsIgnoreCase(symbol))   // N.B. == does not work!
      .findAny()
      .get();
  }

  private static BigDecimal translateQuantity(
      byte[] quantityInBytes,
      InstrumentDefinition def
  ) {
    BigDecimal nativeQuantity = new BigDecimal(new BigInteger(quantityInBytes));
    return nativeQuantity.scaleByPowerOfTen(-def.getDecimals());
  }

  public static void main(String [] args) throws Throwable {
    // A real contributor app would be set up as a gRPC or REST service so that
    // it can keep track of the real-time quotes without the startup cost.
    if (args.length != 4) {
      System.err.println("usage: contribute symbol quantity private-key-for-withdraw-address deposite-address");
      System.exit(1);
    }
    String symbol = args[0];
    Double quantity = Double.parseDouble(args[1]);
    Credentials credentials = Credentials.create(args[2]);
    Address withdrawAddress = Address
      .newBuilder()
      .setValue(credentials.getAddress())
      .setAddressType("ethereum")
      .build();
    Address depositAddress = Address
      .newBuilder()
      .setValue(args[3])
      .setAddressType("ethereum")
      .build();
    NeutralClientConfig config = getConfig();

    // N.B. this is for demo. A real app needs a secure way to import the private key.
    String userPrivateKey = "efd08437a03c8df74e8ee8b10774d67e203f732d6d664613470d85d1f4a32562";
    EcdsaSigner signer = new EcdsaSigner(userPrivateKey);
    NeutralClient client = NeutralClient.build(config, signer);
    System.out.println("AccountInfo: " + client.accountInfo());
    InstrumentResponse instrumentResponse = client.getInstrumentDefinitions();
    InstrumentDefinition NUSD = getInstrumentDefinition("NUSD", instrumentResponse);
    InstrumentDefinition coin = getInstrumentDefinition(symbol, instrumentResponse);

    String neutralContractAddress = NUSD.getInstrumentID().getAddress().getValue();

    UserQuoteProcessor userQuoteProcessor = new UserQuoteProcessor() {
      private NeutralSettler settler = new NeutralSettler(neutralContractAddress, config, credentials);

      @Override
      public void process(SignedUserQuote signedUserQuote) {
        UserQuote userQuote = signedUserQuote.getWrapped();
        // We have to make sure it's a quote in the contribution window.
        BaseQuote baseQuote = userQuote.getSignedBaseQuote().getWrapped();

        Boolean inContributionWindow = baseQuote
          .getFlagsList()
          .stream()
          .anyMatch(flag -> flag.equalsIgnoreCase("InContributionWindow"));

        if (!inContributionWindow) {
          System.err.println("Error: Not in contribution window.");
          System.exit(1);
        }
        try {
          System.out.println("Obtaining an executable record for contribution against quote " + signedUserQuote.getWrapped());
          FinalCommitResponse executableRecord = client.getExecutableRecord(
              CommitType.Sell,
              coin.getInstrumentID(),
              quantity,
              withdrawAddress,
              depositAddress,
              signedUserQuote
          );
          Try<TransactionReceipt> settlementResult = this.settler.settle(executableRecord);
          if (settlementResult.isSuccess()) {
            BigDecimal nusdQuantity = translateQuantity(
              executableRecord.getNeutralQuantity().toByteArray(),
              NUSD
            );
            BigDecimal coinQuantity = translateQuantity(
              executableRecord.getInstrumentQuantity().toByteArray(),
              coin
            );
            settlementResult.forEach(receipt -> {
              System.out.println("Succeeded in selling " + coinQuantity + " " + symbol + " for " + nusdQuantity + " NUSD. TransactionReceipt = " + receipt);
            });
            // We are done
            System.exit(0);
          } else {
            System.err.println("Error: failed to execute the commitment " + settlementResult);
          }
        } catch (Throwable stone) {
          stone.printStackTrace();
          System.err.println("Failed to execute: " + stone);
          // try again
        }
      }
    };

    Completion completion = client.subscribeToUserQuotes(NUSD.getInstrumentID(), userQuoteProcessor);
    System.out.println("waiting for contribution to be done");
    completion.waitUntilDone();
  }
}
