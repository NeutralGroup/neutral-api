/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

import java.math.BigInteger;
import java.nio.charset.StandardCharsets;

import org.bouncycastle.util.encoders.Hex;
// import org.web3j.crypto.ECDSASignature;
import org.web3j.crypto.ECKeyPair;
import org.web3j.crypto.Sign;

public class EcdsaSigner {
  private final BigInteger privKey;
  private final BigInteger pubKey;
  private final ECKeyPair keyPair;

  public EcdsaSigner(String privateKey) {
    this.privKey = new BigInteger(privateKey, 16);
    this.pubKey = Sign.publicKeyFromPrivate(this.privKey);
    this.keyPair = new ECKeyPair(this.privKey, this.pubKey);
  }

  public String sign(byte[] bytes) {
    Sign.SignatureData  signature = Sign.signPrefixedMessage(bytes, this.keyPair);
    return Hex.toHexString(new byte[]{signature.getV()}) + Hex.toHexString(signature.getR()) + Hex.toHexString(signature.getS());
  }

  public String sign(String msg) {
    return this.sign(msg.getBytes(StandardCharsets.UTF_8));
  }
}


