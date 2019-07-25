package com.neutralproject.contracts;

import io.reactivex.Flowable;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import org.web3j.abi.EventEncoder;
import org.web3j.abi.TypeReference;
import org.web3j.abi.datatypes.Address;
import org.web3j.abi.datatypes.Bool;
import org.web3j.abi.datatypes.Event;
import org.web3j.abi.datatypes.Function;
import org.web3j.abi.datatypes.Type;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.abi.datatypes.generated.Uint256;
import org.web3j.crypto.Credentials;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.core.DefaultBlockParameter;
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.request.EthFilter;
import org.web3j.protocol.core.methods.response.Log;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.tx.Contract;
import org.web3j.tx.TransactionManager;
import org.web3j.tx.gas.ContractGasProvider;

/**
 * <p>Auto generated code.
 * <p><strong>Do not modify!</strong>
 * <p>Please use the <a href="https://docs.web3j.io/command_line.html">web3j command line tools</a>,
 * or the org.web3j.codegen.SolidityFunctionWrapperGenerator in the 
 * <a href="https://github.com/web3j/web3j/tree/master/codegen">codegen module</a> to update.
 *
 * <p>Generated with web3j version 4.2.0.
 */
public class NeutralVault extends Contract {
    private static final String BINARY = "Bin file was not provided";

    public static final String FUNC_PERMISSIONS = "permissions";

    public static final String FUNC_DEPOSITS = "deposits";

    public static final String FUNC_NAME = "NAME";

    public static final String FUNC_VERSION = "VERSION";

    public static final String FUNC_READAVAILABLEBALANCE = "readAvailableBalance";

    public static final String FUNC_READPERMISSION = "readPermission";

    public static final String FUNC_DEPOSIT = "deposit";

    public static final String FUNC_WITHDRAW = "withdraw";

    public static final String FUNC_TRANSFER = "transfer";

    public static final String FUNC_GRANT = "grant";

    public static final String FUNC_REVOKE = "revoke";

    public static final String FUNC_OPTIONALMANAGEMENT = "optionalManagement";

    public static final Event EVENTDEPOSIT_EVENT = new Event("EventDeposit", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event EVENTWITHDRAW_EVENT = new Event("EventWithdraw", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event EVENTTRANSFER_EVENT = new Event("EventTransfer", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}, new TypeReference<Address>() {}, new TypeReference<Address>(true) {}, new TypeReference<Uint256>() {}));
    ;

    public static final Event EVENTGRANT_EVENT = new Event("EventGrant", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}));
    ;

    public static final Event EVENTREVOKE_EVENT = new Event("EventRevoke", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Address>(true) {}));
    ;

    public static final Event EVENTOPTIONALMANAGEMENT_EVENT = new Event("EventOptionalManagement", 
            Arrays.<TypeReference<?>>asList(new TypeReference<Address>(true) {}, new TypeReference<Bool>() {}));
    ;

    @Deprecated
    protected NeutralVault(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    protected NeutralVault(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, credentials, contractGasProvider);
    }

    @Deprecated
    protected NeutralVault(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        super(BINARY, contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    protected NeutralVault(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        super(BINARY, contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public RemoteCall<Boolean> permissions(String param0, String param1) {
        final Function function = new Function(FUNC_PERMISSIONS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(param0), 
                new org.web3j.abi.datatypes.Address(param1)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteCall<BigInteger> deposits(String param0, String param1) {
        final Function function = new Function(FUNC_DEPOSITS, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(param0), 
                new org.web3j.abi.datatypes.Address(param1)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<String> NAME() {
        final Function function = new Function(FUNC_NAME, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public RemoteCall<String> VERSION() {
        final Function function = new Function(FUNC_VERSION, 
                Arrays.<Type>asList(), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Utf8String>() {}));
        return executeRemoteCallSingleValueReturn(function, String.class);
    }

    public List<EventDepositEventResponse> getEventDepositEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTDEPOSIT_EVENT, transactionReceipt);
        ArrayList<EventDepositEventResponse> responses = new ArrayList<EventDepositEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventDepositEventResponse typedResponse = new EventDepositEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._forBenefitOf = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
            typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventDepositEventResponse> eventDepositEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventDepositEventResponse>() {
            @Override
            public EventDepositEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTDEPOSIT_EVENT, log);
                EventDepositEventResponse typedResponse = new EventDepositEventResponse();
                typedResponse.log = log;
                typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._forBenefitOf = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
                typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventDepositEventResponse> eventDepositEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTDEPOSIT_EVENT));
        return eventDepositEventFlowable(filter);
    }

    public List<EventWithdrawEventResponse> getEventWithdrawEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTWITHDRAW_EVENT, transactionReceipt);
        ArrayList<EventWithdrawEventResponse> responses = new ArrayList<EventWithdrawEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventWithdrawEventResponse typedResponse = new EventWithdrawEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._to = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
            typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventWithdrawEventResponse> eventWithdrawEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventWithdrawEventResponse>() {
            @Override
            public EventWithdrawEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTWITHDRAW_EVENT, log);
                EventWithdrawEventResponse typedResponse = new EventWithdrawEventResponse();
                typedResponse.log = log;
                typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._to = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
                typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventWithdrawEventResponse> eventWithdrawEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTWITHDRAW_EVENT));
        return eventWithdrawEventFlowable(filter);
    }

    public List<EventTransferEventResponse> getEventTransferEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTTRANSFER_EVENT, transactionReceipt);
        ArrayList<EventTransferEventResponse> responses = new ArrayList<EventTransferEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventTransferEventResponse typedResponse = new EventTransferEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._to = (String) eventValues.getIndexedValues().get(1).getValue();
            typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
            typedResponse._caller = (String) eventValues.getNonIndexedValues().get(0).getValue();
            typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventTransferEventResponse> eventTransferEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventTransferEventResponse>() {
            @Override
            public EventTransferEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTTRANSFER_EVENT, log);
                EventTransferEventResponse typedResponse = new EventTransferEventResponse();
                typedResponse.log = log;
                typedResponse._from = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._to = (String) eventValues.getIndexedValues().get(1).getValue();
                typedResponse._token = (String) eventValues.getIndexedValues().get(2).getValue();
                typedResponse._caller = (String) eventValues.getNonIndexedValues().get(0).getValue();
                typedResponse._quantity = (BigInteger) eventValues.getNonIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventTransferEventResponse> eventTransferEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTTRANSFER_EVENT));
        return eventTransferEventFlowable(filter);
    }

    public List<EventGrantEventResponse> getEventGrantEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTGRANT_EVENT, transactionReceipt);
        ArrayList<EventGrantEventResponse> responses = new ArrayList<EventGrantEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventGrantEventResponse typedResponse = new EventGrantEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._who = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._grant = (String) eventValues.getIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventGrantEventResponse> eventGrantEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventGrantEventResponse>() {
            @Override
            public EventGrantEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTGRANT_EVENT, log);
                EventGrantEventResponse typedResponse = new EventGrantEventResponse();
                typedResponse.log = log;
                typedResponse._who = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._grant = (String) eventValues.getIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventGrantEventResponse> eventGrantEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTGRANT_EVENT));
        return eventGrantEventFlowable(filter);
    }

    public List<EventRevokeEventResponse> getEventRevokeEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTREVOKE_EVENT, transactionReceipt);
        ArrayList<EventRevokeEventResponse> responses = new ArrayList<EventRevokeEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventRevokeEventResponse typedResponse = new EventRevokeEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._who = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._revokedCaller = (String) eventValues.getIndexedValues().get(1).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventRevokeEventResponse> eventRevokeEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventRevokeEventResponse>() {
            @Override
            public EventRevokeEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTREVOKE_EVENT, log);
                EventRevokeEventResponse typedResponse = new EventRevokeEventResponse();
                typedResponse.log = log;
                typedResponse._who = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._revokedCaller = (String) eventValues.getIndexedValues().get(1).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventRevokeEventResponse> eventRevokeEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTREVOKE_EVENT));
        return eventRevokeEventFlowable(filter);
    }

    public List<EventOptionalManagementEventResponse> getEventOptionalManagementEvents(TransactionReceipt transactionReceipt) {
        List<Contract.EventValuesWithLog> valueList = extractEventParametersWithLog(EVENTOPTIONALMANAGEMENT_EVENT, transactionReceipt);
        ArrayList<EventOptionalManagementEventResponse> responses = new ArrayList<EventOptionalManagementEventResponse>(valueList.size());
        for (Contract.EventValuesWithLog eventValues : valueList) {
            EventOptionalManagementEventResponse typedResponse = new EventOptionalManagementEventResponse();
            typedResponse.log = eventValues.getLog();
            typedResponse._token = (String) eventValues.getIndexedValues().get(0).getValue();
            typedResponse._status = (Boolean) eventValues.getNonIndexedValues().get(0).getValue();
            responses.add(typedResponse);
        }
        return responses;
    }

    public Flowable<EventOptionalManagementEventResponse> eventOptionalManagementEventFlowable(EthFilter filter) {
        return web3j.ethLogFlowable(filter).map(new io.reactivex.functions.Function<Log, EventOptionalManagementEventResponse>() {
            @Override
            public EventOptionalManagementEventResponse apply(Log log) {
                Contract.EventValuesWithLog eventValues = extractEventParametersWithLog(EVENTOPTIONALMANAGEMENT_EVENT, log);
                EventOptionalManagementEventResponse typedResponse = new EventOptionalManagementEventResponse();
                typedResponse.log = log;
                typedResponse._token = (String) eventValues.getIndexedValues().get(0).getValue();
                typedResponse._status = (Boolean) eventValues.getNonIndexedValues().get(0).getValue();
                return typedResponse;
            }
        });
    }

    public Flowable<EventOptionalManagementEventResponse> eventOptionalManagementEventFlowable(DefaultBlockParameter startBlock, DefaultBlockParameter endBlock) {
        EthFilter filter = new EthFilter(startBlock, endBlock, getContractAddress());
        filter.addSingleTopic(EventEncoder.encode(EVENTOPTIONALMANAGEMENT_EVENT));
        return eventOptionalManagementEventFlowable(filter);
    }

    public RemoteCall<BigInteger> readAvailableBalance(String _who, String _token) {
        final Function function = new Function(FUNC_READAVAILABLEBALANCE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_who), 
                new org.web3j.abi.datatypes.Address(_token)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Uint256>() {}));
        return executeRemoteCallSingleValueReturn(function, BigInteger.class);
    }

    public RemoteCall<Boolean> readPermission(String _vaultid, String _caller) {
        final Function function = new Function(FUNC_READPERMISSION, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_vaultid), 
                new org.web3j.abi.datatypes.Address(_caller)), 
                Arrays.<TypeReference<?>>asList(new TypeReference<Bool>() {}));
        return executeRemoteCallSingleValueReturn(function, Boolean.class);
    }

    public RemoteCall<TransactionReceipt> deposit(String _forBenefitOf, String _token, BigInteger _quantity) {
        final Function function = new Function(
                FUNC_DEPOSIT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_forBenefitOf), 
                new org.web3j.abi.datatypes.Address(_token), 
                new org.web3j.abi.datatypes.generated.Uint256(_quantity)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> withdraw(String _where, String _token, BigInteger _quantity) {
        final Function function = new Function(
                FUNC_WITHDRAW, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_where), 
                new org.web3j.abi.datatypes.Address(_token), 
                new org.web3j.abi.datatypes.generated.Uint256(_quantity)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> transfer(String _from, String _to, String _token, BigInteger _quantity) {
        final Function function = new Function(
                FUNC_TRANSFER, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_from), 
                new org.web3j.abi.datatypes.Address(_to), 
                new org.web3j.abi.datatypes.Address(_token), 
                new org.web3j.abi.datatypes.generated.Uint256(_quantity)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> grant(String _grant) {
        final Function function = new Function(
                FUNC_GRANT, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_grant)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> revoke(String _revokedCaller) {
        final Function function = new Function(
                FUNC_REVOKE, 
                Arrays.<Type>asList(new org.web3j.abi.datatypes.Address(_revokedCaller)), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    public RemoteCall<TransactionReceipt> optionalManagement() {
        final Function function = new Function(
                FUNC_OPTIONALMANAGEMENT, 
                Arrays.<Type>asList(), 
                Collections.<TypeReference<?>>emptyList());
        return executeRemoteCallTransaction(function);
    }

    @Deprecated
    public static NeutralVault load(String contractAddress, Web3j web3j, Credentials credentials, BigInteger gasPrice, BigInteger gasLimit) {
        return new NeutralVault(contractAddress, web3j, credentials, gasPrice, gasLimit);
    }

    @Deprecated
    public static NeutralVault load(String contractAddress, Web3j web3j, TransactionManager transactionManager, BigInteger gasPrice, BigInteger gasLimit) {
        return new NeutralVault(contractAddress, web3j, transactionManager, gasPrice, gasLimit);
    }

    public static NeutralVault load(String contractAddress, Web3j web3j, Credentials credentials, ContractGasProvider contractGasProvider) {
        return new NeutralVault(contractAddress, web3j, credentials, contractGasProvider);
    }

    public static NeutralVault load(String contractAddress, Web3j web3j, TransactionManager transactionManager, ContractGasProvider contractGasProvider) {
        return new NeutralVault(contractAddress, web3j, transactionManager, contractGasProvider);
    }

    public static class EventDepositEventResponse {
        public Log log;

        public String _from;

        public String _forBenefitOf;

        public String _token;

        public BigInteger _quantity;
    }

    public static class EventWithdrawEventResponse {
        public Log log;

        public String _from;

        public String _to;

        public String _token;

        public BigInteger _quantity;
    }

    public static class EventTransferEventResponse {
        public Log log;

        public String _from;

        public String _to;

        public String _token;

        public String _caller;

        public BigInteger _quantity;
    }

    public static class EventGrantEventResponse {
        public Log log;

        public String _who;

        public String _grant;
    }

    public static class EventRevokeEventResponse {
        public Log log;

        public String _who;

        public String _revokedCaller;
    }

    public static class EventOptionalManagementEventResponse {
        public Log log;

        public String _token;

        public Boolean _status;
    }
}

