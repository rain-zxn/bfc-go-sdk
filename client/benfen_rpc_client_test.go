package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"testing"
)

func client_bfc_getChainIdentifier(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetChainIdentifier(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfcx_getTotalSupply(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetTotalSupply(
		context.Background(),
		types.BFC_COIN_TYPE,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfcx_getValidatorsApy(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetValidatorsApy(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}
func client_bfcx_getReferenceGasPrice(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetReferenceGasPrice(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func TestAllCrossChainEvent(t *testing.T) {
	TestClient_QueryEvents(t)

}

func TestAll(t *testing.T) {
	client_bfc_getChainIdentifier(t)
	client_bfcx_getTotalSupply(t)
	client_bfcx_getValidatorsApy(t)
	client_bfcx_getReferenceGasPrice(t)

	TestClient_GetObject(t)
	TestClient_MultiGetObjects(t)
	TestClient_GetLatestCheckpointSequenceNumber(t)
	TestClient_GetLatestSuiSystemState(t)
	TestGetDelegatedStakes(t)
	TestGetStakesByIds(t)
	TestRequestAddDelegation(t)
	//need account has available stake object
	//TestRequestWithdrawDelegation(t)

	TestClient_GetAllBalances(t)
	TestClient_GetAllCoins(t)
	TestClient_GetCoins(t)
	TestClient_GetBalance(t)
	TestClient_GetOwnedObjects(t)
	TestClient_GetCoinMetadata(t)
	TestClient_QueryTransactionBlocks(t)
	TestClient_GetEvents(t)
	TestClient_GetTotalTransactionBlocks(t)
	TestClient_GetTransaction(t)
	TestClient_MultiGetObjects(t)
	TestClient_TryGetPastObject(t)
	client_bfc_getProtocolConfig(t)
	client_bfc_getCheckpoint(t)
	client_bfc_getCheckpoints(t)
	client_bfc_getCommiteeInfo(t)
	TestBCS_TransferBFC(t)
	TestClient_PayBFC(t)
	TestBCS_PayAllBFC(t)
	TestBCS_Pay(t)
	TestClient_SplitCoin(t)
	TestClient_SplitCoinEqual(t)
	TestClient_MergeCoins(t)
	TestClient_TransferObject(t)
	TestClient_QueryEvents(t)
	client_bfc_multiGetTransactionBlocks(t)
	client_bfc_tryMultiGetPastObjects(t)
	client_bfc_getMoveFunctionArgTypes(t)
	client_bfc_getNormalizedMoveFunction(t)
	client_bfc_getNormalizedMoveModule(t)
	client_bfc_getNormalizedMoveStruct(t)
	client_bfc_getNormalizedMoveModulesByPackage(t)

	TestClient_DryRunTransaction(t)
	client_bfc_resolveNameServiceAddress(t)
	client_bfc_resolveNameServiceNames(t)
	client_bfc_publish(t)
	client_bfc_executeTransactionBlock(t)
	client_bfc_moveCall(t)
	client_bfc_getLoadedChildObjects(t)
	client_bfc_requestAddStake(t)
	client_bfc_requestWithdrawStake(t)
	client_bfc_batTransactions(t)

}
func main() {

}

func TestSingle(t *testing.T) {
	//client_bfc_subscribeEvent(t)
	//client_bfc_devInspectTransactionBlock(t)
	//client_bfc_requestWithdrawStake(t)
	//client_bfc_getDynamicFieldObject(t)
	//client_bfc_getCheckpoint(t)
	client_bfc_subscribeEvent(t)
}

func client_GetInnerDaoInf(t *testing.T) {
	chain := LocalnetClient(t)

	resp, err := chain.GetInnerDaoInfo(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_GetStablePools(t *testing.T) {
	chain := LocalnetClient(t)

	address, _ := bfc_types.NewAddressFromHex(
		"0xa162390d3e52eab4bb7b75653a30577048de3328bcb16e4c955ee1c40d2b8ece",
	)
	resp, err := chain.GetStablePools(
		context.Background(),
		*address,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_GetProposalInfo(t *testing.T) {
	chain := LocalnetClient(t)

	address, _ := bfc_types.NewAddressFromHex(
		"0xd1701b76bc920eedc21c4b97c5729209d9443748979de783ec020166dad31fc8",
	)
	resp, err := chain.GetProposalInfo(
		context.Background(),
		*address,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getDynamicFieldObject(t *testing.T) {
	chain := LocalnetClient(t)
	address, err := bfc_types.NewAddressFromHex(
		"0x6337de842b9daae517401654860d7072d6fb1d78f27dd10e562acfc7e1f09b57",
	)

	name := bfc_types.DynamicFieldName{
		Type:  "0x1::ascii::String",
		Value: "00000000000000000000000000000000000000000000000000000000000000c8::busd::BUSD",
	}

	resp, err := chain.GetDynamicFieldObject(
		context.Background(),
		*address,
		name,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_batTransactions(t *testing.T) {
	chain := ChainClient(t)
	signer, err := bfc_types.NewAddressFromHex("0x7113a31aa484dfca371f854ae74918c7463c7b3f1bf4c1fe8ef28835e88fd590")

	coin1, err := bfc_types.NewAddressFromHex("0xe05bac2f7ba94d4e9fd4fae3bb80a17be122559dea76dd887431da363882ca34")

	data := types.TransferObjectParams{
		ObjectId:  *coin1,
		Recipient: *signer,
	}

	transferParamMap := map[string]interface{}{
		"transferObjectRequestParams": data,
	}

	transferMaps := []map[string]interface{}{transferParamMap}

	resp, err := chain.BatchTransaction(
		context.Background(),
		*signer,
		transferMaps,
		nil,
		"100000000",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_requestWithdrawStake(t *testing.T) {
	chain := ChainClient(t)
	signer, err := bfc_types.NewAddressFromHex("0x7113a31aa484dfca371f854ae74918c7463c7b3f1bf4c1fe8ef28835e88fd590")
	coin1, err := bfc_types.NewAddressFromHex("0x3f7e138fc33c47279abb65d92c1859b203958ea5dd3ce77c8c09ebf4796d7cc7")

	resp, err := chain.RequestWithdrawStake(
		context.Background(),
		*signer,
		*coin1,
		nil,
		decimal.NewFromInt(100000000),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_requestAddStake(t *testing.T) {
	chain := ChainClient(t)
	signer, err := bfc_types.NewAddressFromHex("0x7113a31aa484dfca371f854ae74918c7463c7b3f1bf4c1fe8ef28835e88fd590")

	validator, err := bfc_types.NewAddressFromHex(
		"0xc1e11c130f55ee714387c801963bbf0ed8ff8da629e45ccc491fdee3751de867",
	)

	coin1, err := bfc_types.NewAddressFromHex("0x902b2dff406e2d194d52f4ad59efa1e27401e42288d23393ec90af91f29a8e50")
	coins := []bfc_types.ObjectID{*coin1}
	resp, err := chain.RequestAddStake(
		context.Background(),
		*signer,
		coins,
		decimal.NewFromInt(10000000000),
		*validator,
		nil,
		decimal.NewFromInt(100000000),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getLoadedChildObjects(t *testing.T) {
	chain := ChainClient(t)

	resp, err := chain.getLoadedChildObjects(
		context.Background(),
		"E1BGygDpR7QhXpiZz5DEUyh11hCKKE4eUBJeSvp9Yn26",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_moveCall(t *testing.T) {
	chain := ChainClient(t)
	signer, err := bfc_types.NewAddressFromHex("0x7113a31aa484dfca371f854ae74918c7463c7b3f1bf4c1fe8ef28835e88fd590")
	packageId, _ := bfc_types.NewAddressFromHex(
		"0x7916c5becc284217bc9654fe4b724eeedd8ec45bf1b80050804e34a218f1f034",
	)

	args := []any{}
	resp, err := chain.MoveCall(
		context.Background(),
		*signer,
		*packageId,
		"counter",
		"getCounter",
		[]string{},
		args,
		nil,
		types.NewSafeBfcBigInt(uint64(100000000)),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_devInspectTransactionBlock(t *testing.T) {
	chain := ChainClient(t)

	txbyte := "AAACAQCYsnriMoGiNcCeg1APEjY2zUM/anYqD0p/n0D7cyG/gw0AAAAAAAAAID5rjkVyVoh/AkLNHb9Dsp6eQEaPFgm3Oj4vVxezQe0SAAkBAJQ1dwAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA3BheQlzcGxpdF92ZWMBBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA2JmYwNCRkMAAgEAAAEBAPKGGQaxc4jiT1NGltL7Czb8SK+tALU0yk2iGyB+plSlASKpgXP23aQ0DTmaynSy5Gg1R40BmY9mXZvs7iyLocZSDQAAAAAAAAAgw3hk7piJ/5miZ/vE7ZCTUSsTI6ZYN/b8iLuAiTBhLzvyhhkGsXOI4k9TRpbS+ws2/EivrQC1NMpNohsgfqZUpWQAAAAAAAAAgJaYAAAAAAAA"
	objId, err := bfc_types.NewAddressFromHex("0xf2861906b17388e24f534696d2fb0b36fc48afad00b534ca4da21b207ea654a5")

	resp, err := chain.devInspectTransactionBlock(
		context.Background(),
		objId,
		txbyte,
		nil,
		"10000000",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_executeTransactionBlock(t *testing.T) {
	chain := ChainClient(t)
	txbytes := "AAACAQCYsnriMoGiNcCeg1APEjY2zUM/anYqD0p/n0D7cyG/gwsAAAAAAAAAICXEzWGQpDVzmCeJacFIDnFOiIYcZizYj2Nb5tFWJOdVAAkBAJQ1dwAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA3BheQlzcGxpdF92ZWMBBwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACA2JmYwNCRkMAAgEAAAEBAPKGGQaxc4jiT1NGltL7Czb8SK+tALU0yk2iGyB+plSlAfQfrz0REIerGmKd1QQ1FeA4cWHdWghNBx7rb2OjblFeDAAAAAAAAAAg2AnQiqJtMCcLOI/5Rw3WYNWAmmsMgw6Azzu9k70FIrTyhhkGsXOI4k9TRpbS+ws2/EivrQC1NMpNohsgfqZUpWQAAAAAAAAAgJaYAAAAAAAA"
	signatures := []string{"AOIM8aEcJEkrjKHbCOfcQaObAOW3+ZFRV+mEz258A1R64GxEv8nChi26xQcjBMrwGRRzAIYLx3EF1Rk7d62V/wgiCgf7H2+YIRmOSlhqEyhPdpBTyuRVL/rUCay7O5BLfw=="}
	resp, err := chain.executeTransactionBlock(
		context.Background(),
		txbytes,
		signatures,
		&types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_subscribeEvent(t *testing.T) {
	chain := ChainClient(t)
	objId, err := bfc_types.NewAddressFromHex("0x7113a31aa484dfca371f854ae74918c7463c7b3f1bf4c1fe8ef28835e88fd590")
	if err != nil {
		fmt.Println("NewAddressFromHex err = ", err)
	}
	resp, err := chain.SubscribeEvent(
		context.Background(),
		objId,
	)
	if err != nil {
		fmt.Println("SubscribeEvent err = ", err)

	}

	PrintJson(resp)
}

// sender
func client_bfc_publish(t *testing.T) {
	chain := ChainClient(t)

	sender, err := bfc_types.NewAddressFromHex("0xf2861906b17388e24f534696d2fb0b36fc48afad00b534ca4da21b207ea654a5")
	modules := []string{
		"oRzrCwYAAAADAQACBwIZCBsgAAAYYmZjX2Rhb192b3RpbmdfcG9vbF90ZXN0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		"oRzrCwYAAAAEAQACBwINCA8gBi8SAAAMYmZjX2Rhb190ZXN0AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAQICAQMCAQQCAQUCAQYCAQcA",
		"oRzrCwYAAAALAQAMAgwUAyAqBEoEBU4lB3OXAQiKAmAG6gIPCvkCDgyHA0sN0gMCAAUBEAIIAg4CEQISAAEIAAAABwABAgcAAwQEAAUDAgAABgABAAA" +
			"JAAEAAAsCAQABEwMEAAIHBgEBAwMNAAkABBELAQEIBQ8IBwAEBQYKAQcIBAABBwgAAQoCAQgCAQgBAQkAAQUBBggEAQgDAQgAAgkABQpDb3VudEV2ZW50B" +
			"0NvdW50ZXIGU3RyaW5nCVR4Q29udGV4dANVSUQHY291bnRlcgtjcmVhdGVFdmVudARlbWl0BWV2ZW50CmdldENvdW50ZXICaWQEaW5jcgRuYW1lA25ldwZv" +
			"YmplY3QGc2VuZGVyBnN0cmluZwh0cmFuc2Zlcgp0eF9jb250ZXh0BHV0ZjgFdmFsdWUAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAIKAgwLaGVsbG8gd29ybGQAAgIKCAMUAwECAQwIAgABBAABBQc" +
			"AEQMSATgAAgEBBAAHCwoALhEHDAELABEFBgAAAAAAAAAAEgALATgBAgIBBAABCQoAEAAUBgEAAAAAAAAAFgsADwAVAgABAA==",
	}

	dep1, err := bfc_types.NewAddressFromHex("0x0000000000000000000000000000000000000000000000000000000000000001")
	dep2, err := bfc_types.NewAddressFromHex("0x0000000000000000000000000000000000000000000000000000000000000002")
	dependencies := []bfc_types.ObjectID{*dep1, *dep2}
	resp, err := chain.publish(
		context.Background(),
		sender,
		modules,
		dependencies,
		nil,
		"1000000000",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_resolveNameServiceNames(t *testing.T) {
	chain := ChainClient(t)

	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.resolveNameServiceNames(
		context.Background(),
		objId,
		nil,
		10,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_resolveNameServiceAddress(t *testing.T) {
	chain := ChainClient(t)

	resp, err := chain.resolveNameServiceAddress(
		context.Background(),
		"example.bfc",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getNormalizedMoveModulesByPackage(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.getNormalizedMoveModulesByPackage(
		context.Background(),
		objId,
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getNormalizedMoveStruct(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.getNormalizedMoveStruct(
		context.Background(),
		objId,
		"bfc_system",
		"BfcSystemState",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getNormalizedMoveModule(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.getNormalizedMoveModule(
		context.Background(),
		objId,
		"bfc_system",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getNormalizedMoveFunction(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.getNormalizedMoveFunction(
		context.Background(),
		objId,
		"bfc_system",
		"create_stake_manager_key",
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getMoveFunctionArgTypes(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	objId, err := bfc_types.NewAddressFromHex("0x00000000000000000000000000000000000000000000000000000000000000c8")

	resp, err := chain.getMoveFunctionArgTypes(
		context.Background(),
		objId,
		"bfc_system",
		"create_stake_manager_key",
	)
	require.NoError(t, err)

	PrintJson(resp)
}
func client_bfc_tryMultiGetPastObjects(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	objId, err := bfc_types.NewAddressFromHex("0x0000000000000000000000000000000000000000000000000000000000000006")
	r1 := types.GetPastObjectRequest{
		ObjectId: *objId,
		Version:  "33",
	}

	resp, err := chain.tryMultiGetPastObjects(
		context.Background(),
		[]types.GetPastObjectRequest{r1},
		nil,
	)
	require.NoError(t, err)

	PrintJson(resp)
}
func client_bfc_multiGetTransactionBlocks(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)
	digest := "2uDcP2NdAFteFnZQhKFqx79QsidzsSfHx5r2r4FDhGZL"
	d1, err := bfc_types.NewDigest(digest)

	digest2 := "CxMqqMCxnD47DG2KaExu9remr7nHcqxhiwZJaH3mQj7X"
	d2, err := bfc_types.NewDigest(digest2)

	resp, err := chain.multiGetTransactionBlocks(
		context.Background(),
		[]bfcDigest{*d1, *d2},
		&types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func client_bfc_getCommiteeInfo(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetCommiteeInfo(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)

}

func client_bfc_getCheckpoints(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetCheckPoints(
		context.Background(),
		nil,
		10,
		false,
	)
	require.NoError(t, err)

	PrintJson(resp)

}

func client_bfc_getCheckpoint(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetCheckPoint(
		context.Background(),
		"1",
	)
	require.NoError(t, err)

	PrintJson(resp)

	var obj types.CheckPointObject
	err = json.Unmarshal([]byte(resp), &obj)
	if err != nil {
		fmt.Println("decode JSON fail:", err)
		return
	}

	fmt.Println("decode result:")
	fmt.Println("Epoch:", obj.Epoch)
	fmt.Println("SequenceNumber:", obj.SequenceNumber)
	fmt.Println("Digest:", obj.Digest)

}
func client_bfc_getProtocolConfig(t *testing.T) {
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetProtocolConfig(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}
