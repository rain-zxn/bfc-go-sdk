package client

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"

	"github.com/hellokittyboy-code/benfen-go-sdk/account"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestClient_TransferObject(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	recipient := signer
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(coins.Data), 2)
	coin := coins.Data[0]

	txn, err := cli.TransferObject(
		context.Background(), *signer, *recipient,
		coin.CoinObjectId, nil, types.NewSafeBfcBigInt(BFC(0.01).Uint64()),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_TransferBFC(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	recipient := signer
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.0001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 1, 0)
	require.Nil(t, err)

	txn, err := cli.TransferBFC(
		context.Background(), *signer, *recipient,
		pickedCoins.Coins[0].CoinObjectId,
		types.NewSafeBfcBigInt(amount),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_PayAllBFC(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	recipient := signer
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)
	require.Nil(t, err)

	txn, err := cli.PayAllBFC(
		context.Background(), *signer, *recipient,
		pickedCoins.CoinIds(),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_Pay(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	recipient := Address
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	limit := len(coins.Data) - 1 // need reserve a coin for gas

	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, limit, 0)
	require.NoError(t, err)

	txn, err := cli.Pay(
		context.Background(), *signer,
		pickedCoins.CoinIds(),
		[]BfcAddress{*recipient},
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(amount),
		},
		nil,
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_PayBFC(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	recipient := Address
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)
	require.NoError(t, err)

	txn, err := cli.PayBFC(
		context.Background(), *signer,
		pickedCoins.CoinIds(),
		[]BfcAddress{*recipient},
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(amount),
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func CallClient_onChain_PayBFC() {
	cli, _ := Dial(types.TestnetRpcUrl)

	signer, _ := account.NewAccountWithMnemonic(M1Mnemonic)

	recipient := Address

	coins, err := cli.GetCoins(context.Background(), *Address, nil, nil, 10)
	if err != nil {
		println("err = ", err)
	}

	amount := BFC(0.001).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)

	txn, err := cli.PayBFC(
		context.Background(),
		*Address,
		pickedCoins.CoinIds(),
		[]BfcAddress{*recipient},
		[]types.SafeBfcBigInt[uint64]{
			types.NewSafeBfcBigInt(amount),
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	// sign and send
	signature, err := signer.SignSecureWithoutEncode(txn.TxBytes, bfc_types.DefaultIntent())
	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(), txn.TxBytes, []any{signature}, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)

	println("resp = ", resp)
}

func TestClient_SplitCoin(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), 0, 1, 0)
	require.NoError(t, err)
	splitCoins := []types.SafeBfcBigInt[uint64]{types.NewSafeBfcBigInt(amount / 2)}

	txn, err := cli.SplitCoin(
		context.Background(), *signer,
		pickedCoins.Coins[0].CoinObjectId,
		splitCoins,
		nil, types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, false)
}

func TestClient_SplitCoinEqual(t *testing.T) {
	cli := ChainClient(t)
	signer := M1Address(t)
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), 0, 1, 0)
	require.NoError(t, err)

	txn, err := cli.SplitCoinEqual(
		context.Background(), *signer,
		pickedCoins.Coins[0].CoinObjectId,
		types.NewSafeBfcBigInt(uint64(2)),
		nil, types.NewSafeBfcBigInt(gasBudget),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_MergeCoins(t *testing.T) {
	cli := ChainClient(t)
	signer := Address
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)
	require.True(t, len(coins.Data) >= 3)

	coin1 := coins.Data[0]
	coin2 := coins.Data[1]
	coin3 := coins.Data[2] // gas coin

	gasBudge := BFC(0.01).Uint64()

	txn, err := cli.MergeCoins(
		context.Background(), *signer,
		coin1.CoinObjectId,
		coin2.CoinObjectId,
		&coin3.CoinObjectId,
		types.NewSafeBfcBigInt(gasBudge),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txn.TxBytes, true)
}

func TestClient_Publish(t *testing.T) {
	t.Log("TestClient_Publish TODO")
	// cli := ChainClient(t)

	// txnBytes, err := cli.Publish(context.Background(), *signer, *coin1, *coin2, nil, 10000)
	// require.Nil(t, err)
	// simulateCheck(t, cli, txnBytes, M1Account(t))
}

func TestClient_MoveCall(t *testing.T) {
	t.Log("TestClient_MoveCall TODO")
	// cli := ChainClient(t)

	// txnBytes, err := cli.MoveCall(context.Background(), *signer, *coin1, *coin2, nil, 10000)
	// require.Nil(t, err)
	// simulateCheck(t, cli, txnBytes, M1Account(t))
}

func TestClient_BatchTransaction(t *testing.T) {
	t.Log("TestClient_BatchTransaction TODO")
	// cli := ChainClient(t)

	// txnBytes, err := cli.BatchTransaction(context.Background(), *signer, *coin1, *coin2, nil, 10000)
	// require.Nil(t, err)
	// simulateCheck(t, cli, txnBytes, M1Account(t))
}

// @return types.DryRunTransactionBlockResponse
func simulateCheck(
	t *testing.T,
	cli *Client,
	txBytes bfcBase64Data,
	showJson bool,
) *types.DryRunTransactionBlockResponse {
	simulate, err := cli.DryRunTransaction(context.Background(), txBytes)
	require.Nil(t, err)
	require.Equal(t, simulate.Effects.Data.V1.Status.Error, "")
	require.True(t, simulate.Effects.Data.IsSuccess())
	if showJson {
		data, err := json.Marshal(simulate)
		require.Nil(t, err)
		t.Log(string(data))
		t.Log("gasFee = ", simulate.Effects.Data.GasFee())
	}
	return simulate
}

func executeTxn(
	t *testing.T,
	cli *Client,
	txBytes bfcBase64Data,
	acc *account.Account,
) *types.BfcTransactionBlockResponse {
	// First of all, make sure that there are no problems with simulated trading.
	simulate, err := cli.DryRunTransaction(context.Background(), txBytes)
	require.Nil(t, err)
	require.True(t, simulate.Effects.Data.IsSuccess())

	// sign and send
	signature, err := acc.SignSecureWithoutEncode(txBytes, bfc_types.DefaultIntent())
	require.NoError(t, err)
	options := types.BfcTransactionBlockResponseOptions{
		ShowEffects: true,
	}
	resp, err := cli.ExecuteTransactionBlock(
		context.TODO(), txBytes, []any{signature}, &options,
		types.TxnRequestTypeWaitForLocalExecution,
	)
	require.NoError(t, err)
	t.Log(resp)
	return resp
}
