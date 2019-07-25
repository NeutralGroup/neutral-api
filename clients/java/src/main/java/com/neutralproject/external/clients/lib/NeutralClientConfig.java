/* Neutral Java Client Library
   Copyright 2018, Neutral Empire Group, Inc
   228 Park Ave S #45204, New York, NY 10021
   This is commercial software and is not available under an open source license. */

package com.neutralproject.external.clients.lib;

public class NeutralClientConfig {
  public String host;
  public Integer port;
  public Boolean secure;
  public String ethNode;
  public String neutralContractAddress;
  public Long gasPrice;
  public Long gasLimit;
  public String certificateName;

  public NeutralClientConfig(String host, Integer port, Boolean secure, String ethNode, String neutralContractAddress, Long gasPrice, Long gasLimit) {
    this(host, port, secure, ethNode, neutralContractAddress, gasPrice, gasLimit, "");
  }

  public NeutralClientConfig(String host, Integer port, Boolean secure, String ethNode, String neutralContractAddress, Long gasPrice, Long gasLimit, String certificateName) {
    this.host = host;
    this.port = port;
    this.secure = secure;
    this.ethNode = ethNode;
    this.neutralContractAddress = neutralContractAddress;
    this.gasPrice = gasPrice;
    this.gasLimit = gasLimit;
    this.certificateName = certificateName;
  }

  public Boolean useSelfSignedCertificate() {
    return false;
  }
}
