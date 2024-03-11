package client

const (
	BfcXPrefix   = "bfcx_"
	BfcPrefix    = "bfc_"
	UnsafePrefix = "unsafe_"
)

type Method interface {
	String() string
}

type BfcMethod string

func (s BfcMethod) String() string {
	return BfcPrefix + string(s)
}

type BfcXMethod string

func (s BfcXMethod) String() string {
	return BfcXPrefix + string(s)
}

type UnsafeMethod string

func (u UnsafeMethod) String() string {
	return UnsafePrefix + string(u)
}

const (
	getLoadedChildObjects             BfcMethod = "getLoadedChildObjects"
	devInspectTransactionBlock        BfcMethod = "devInspectTransactionBlock"
	dryRunTransactionBlock            BfcMethod = "dryRunTransactionBlock"
	executeTransactionBlock           BfcMethod = "executeTransactionBlock"
	getChainIdentifier                BfcMethod = "getChainIdentifier"
	getCheckpoint                     BfcMethod = "getCheckpoint"
	getCheckpoints                    BfcMethod = "getCheckpoints"
	getEvents                         BfcMethod = "getEvents"
	getLatestCheckpointSequenceNumber BfcMethod = "getLatestCheckpointSequenceNumber"
	getMoveFunctionArgTypes           BfcMethod = "getMoveFunctionArgTypes"
	getNormalizedMoveFunction         BfcMethod = "getNormalizedMoveFunction"
	getNormalizedMoveModule           BfcMethod = "getNormalizedMoveModule"
	getNormalizedMoveModulesByPackage BfcMethod = "getNormalizedMoveModulesByPackage"
	getNormalizedMoveStruct           BfcMethod = "getNormalizedMoveStruct"
	getObject                         BfcMethod = "getObject"
	getTotalTransactionBlocks         BfcMethod = "getTotalTransactionBlocks"
	getTransactionBlock               BfcMethod = "getTransactionBlock"
	multiGetObjects                   BfcMethod = "multiGetObjects"
	multiGetTransactionBlocks         BfcMethod = "multiGetTransactionBlocks"
	tryGetPastObject                  BfcMethod = "tryGetPastObject"
	tryMultiGetPastObjects            BfcMethod = "tryMultiGetPastObjects"
	getProtocolConfig                 BfcMethod = "getProtocolConfig"
	getInnerDaoInfo                   BfcMethod = "getInnerDaoInfo"

	getAllBalances            BfcXMethod = "getAllBalances"
	getAllCoins               BfcXMethod = "getAllCoins"
	getBalance                BfcXMethod = "getBalance"
	getCoinMetadata           BfcXMethod = "getCoinMetadata"
	getCoins                  BfcXMethod = "getCoins"
	getCommitteeInfo          BfcXMethod = "getCommitteeInfo"
	getCurrentEpoch           BfcXMethod = "getCurrentEpoch"
	getDynamicFieldObject     BfcXMethod = "getDynamicFieldObject"
	getProposal               BfcXMethod = "getProposal"
	getStablePools            BfcXMethod = "getStablePools"
	getDynamicFields          BfcXMethod = "getDynamicFields"
	getEpochs                 BfcXMethod = "getEpochs"
	getLatestSuiSystemState   BfcXMethod = "getLatestSuiSystemState"
	getMoveCallMetrics        BfcXMethod = "getMoveCallMetrics"
	getNetworkMetrics         BfcXMethod = "getNetworkMetrics"
	getOwnedObjects           BfcXMethod = "getOwnedObjects"
	getReferenceGasPrice      BfcXMethod = "getReferenceGasPrice"
	getStakes                 BfcXMethod = "getStakes"
	getStakesByIds            BfcXMethod = "getStakesByIds"
	getTotalSupply            BfcXMethod = "getTotalSupply"
	getValidatorsApy          BfcXMethod = "getValidatorsApy"
	queryEvents               BfcXMethod = "queryEvents"
	queryObjects              BfcXMethod = "queryObjects"
	queryTransactionBlocks    BfcXMethod = "queryTransactionBlocks"
	subscribeEvent            BfcXMethod = "subscribeEvent"
	resolveNameServiceAddress BfcXMethod = "resolveNameServiceAddress"
	resolveNameServiceNames   BfcXMethod = "resolveNameServiceNames"

	batchTransaction     UnsafeMethod = "batchTransaction"
	mergeCoins           UnsafeMethod = "mergeCoins"
	moveCall             UnsafeMethod = "moveCall"
	pay                  UnsafeMethod = "pay"
	payAllBfc            UnsafeMethod = "payAllBfc"
	payBfc               UnsafeMethod = "payBfc"
	publish              UnsafeMethod = "publish"
	requestAddStake      UnsafeMethod = "requestAddStake"
	requestWithdrawStake UnsafeMethod = "requestWithdrawStake"
	splitCoin            UnsafeMethod = "splitCoin"
	splitCoinEqual       UnsafeMethod = "splitCoinEqual"
	transferObject       UnsafeMethod = "transferObject"
	transferBfc          UnsafeMethod = "transferBfc"
)
