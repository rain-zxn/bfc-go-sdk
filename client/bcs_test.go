package client

import (
	"context"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types/sui_system_state"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"testing"

	"github.com/fardream/go-bcs/bcs"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/stretchr/testify/require"
)

// 0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea
// [monkey tragic drive owner fade mimic taxi despair endorse peasant amused woman]
func TestBCS_TransferObject(t *testing.T) {
	sender, err := bfc_types.NewAddressFromHex(Address.String())
	require.NoError(t, err)
	recipient := sender
	gasBudget := BFC(0.1).Uint64()

	cli := ChainClient(t)
	coins := GetCoins(t, cli, *sender, 2)
	coin, gas := coins[0], coins[1]

	gasPrice := uint64(100)
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err = ptb.TransferObject(*recipient, []*bfc_types.ObjectRef{coin.Reference()})
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			gas.Reference(),
		},
		pt,
		gasBudget,
		gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)

	// build with remote rpc
	txn, err := cli.TransferObject(
		context.Background(), *sender, *recipient,
		coin.CoinObjectId,
		&gas.CoinObjectId,
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	txBytesRemote := txn.TxBytes.Data()

	require.Equal(t, txBytesBCS, txBytesRemote)
}

func TestBCS_TransferBFC(t *testing.T) {
	sender, err := bfc_types.NewAddressFromHex(Address.String())
	require.NoError(t, err)
	recipient := sender
	amount := BFC(0.1).Uint64()
	gasBudget := BFC(0.1).Uint64()

	cli := ChainClient(t)
	coin := GetCoins(t, cli, *sender, 1)[0]

	gasPrice := uint64(100)
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err = ptb.TransferBFC(*recipient, &amount)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)

	// build with remote rpc
	txn, err := cli.TransferBFC(
		context.Background(), *sender, *recipient, coin.CoinObjectId,
		types.NewSafeBfcBigInt(amount),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	txBytesRemote := txn.TxBytes.Data()

	require.Equal(t, txBytesBCS, txBytesRemote)
}

func TestBCS_PayBFC(t *testing.T) {
	sender := Address
	//require.NoError(t, err)
	// recipient := sender
	recipient2, _ := bfc_types.NewAddressFromHex("0x123456")
	amount := BFC(0.1).Uint64()
	gasBudget := BFC(0.1).Uint64()

	cli := ChainClient(t)
	coin := GetCoins(t, cli, *sender, 1)[0]

	gasPrice := uint64(100)
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err := ptb.PayBFC([]BfcAddress{*recipient2, *recipient2}, []uint64{amount, amount})
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)

	resp := simulateCheck(t, cli, txBytesBCS, true)
	gasFee := resp.Effects.Data.GasFee()
	t.Log(gasFee)

	// build with remote rpc
	// txn, err := cli.PayBFC(context.Background(), *sender, []bfcObjectID{coin.CoinObjectId},
	// 	[]BfcAddress{*recipient2, *recipient2},
	// 	[]types.SafeBfcBigInt[uint64]{types.NewSafeBfcBigInt(amount), types.NewSafeBfcBigInt(amount)},
	// 	types.NewSafeBfcBigInt(gasBudget))
	// require.NoError(t, err)
	// txBytesRemote := txn.TxBytes.Data()

	// XXX: Fail when there are multiple recipients
	// require.Equal(t, txBytesBCS, txBytesRemote)
}

func TestBCS_PayAllBFC(t *testing.T) {
	sender, err := bfc_types.NewAddressFromHex(Address.String())
	require.NoError(t, err)
	recipient := sender
	gasBudget := BFC(0.1).Uint64()

	cli := ChainClient(t)
	coins := GetCoins(t, cli, *sender, 2)
	coin, coin2 := coins[0], coins[1]

	gasPrice := uint64(100)
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err = ptb.PayAllBFC(*recipient)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
			coin2.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)

	// build with remote rpc
	txn, err := cli.PayAllBFC(
		context.Background(), *sender, *recipient,
		[]bfcObjectID{
			coin.CoinObjectId, coin2.CoinObjectId,
		},
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)
	txBytesRemote := txn.TxBytes.Data()

	require.Equal(t, txBytesBCS, txBytesRemote)
}

func TestBCS_Pay(t *testing.T) {
	sender, err := bfc_types.NewAddressFromHex(Address.String())
	require.NoError(t, err)
	// recipient := sender
	recipient2, _ := bfc_types.NewAddressFromHex("0x123456")
	amount := BFC(0.1).Uint64()
	gasBudget := BFC(0.1).Uint64()

	cli := ChainClient(t)
	coins := GetCoins(t, cli, *sender, 2)
	coin, gas := coins[0], coins[1]

	gasPrice := uint64(100)
	// gasPrice, err := cli.GetReferenceGasPrice(context.Background())

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	err = ptb.Pay(
		[]*bfc_types.ObjectRef{coin.Reference()},
		[]BfcAddress{*recipient2, *recipient2},
		[]uint64{amount, amount},
	)
	require.NoError(t, err)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			gas.Reference(),
		},
		pt, gasBudget, gasPrice,
	)
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)

	resp := simulateCheck(t, cli, txBytesBCS, true)
	gasfee := resp.Effects.Data.GasFee()
	t.Log(gasfee)

	// build with remote rpc
	// txn, err := cli.Pay(context.Background(), *sender,
	// 	[]bfcObjectID{coin.CoinObjectId},
	// 	[]BfcAddress{*recipient, *recipient2},
	// 	[]types.SafeBfcBigInt[uint64]{types.NewSafeBfcBigInt(amount), types.NewSafeBfcBigInt(amount)},
	// 	&gas.CoinObjectId,
	// 	types.NewSafeBfcBigInt(gasBudget))
	// require.NoError(t, err)
	// txBytesRemote := txn.TxBytes.Data()

	// XXX: Fail when there are multiple recipients
	// require.Equal(t, txBytesBCS, txBytesRemote)
}

// important , need check if the validator address is right...
func TestBCS_MoveCall(t *testing.T) {
	sender, err := bfc_types.NewAddressFromHex(Address.String())
	require.NoError(t, err)
	gasBudget := BFC(0.2).Uint64()
	gasPrice := uint64(100)

	cli := ChainClient(t)
	coins := GetCoins(t, cli, *sender, 2)
	coin, coin2 := coins[0], coins[1]

	validatorAddress, err := bfc_types.NewAddressFromHex(ValidatorAddress)
	require.NoError(t, err)

	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()

	// case 1: split target amount
	amtArg, err := ptb.Pure(BFC(1).Uint64())
	require.NoError(t, err)
	arg1 := ptb.Command(
		bfc_types.Command{
			SplitCoins: &struct {
				Argument  bfc_types.Argument
				Arguments []bfc_types.Argument
			}{
				Argument:  bfc_types.Argument{GasCoin: &lib.EmptyEnum{}},
				Arguments: []bfc_types.Argument{amtArg},
			},
		},
	) // the coin is split result argument
	arg2, err := ptb.Pure(validatorAddress)
	require.NoError(t, err)
	arg0, err := ptb.Obj(bfc_types.SuiSystemMutObj)
	require.NoError(t, err)
	ptb.Command(
		bfc_types.Command{
			MoveCall: &bfc_types.ProgrammableMoveCall{
				Package:  *bfc_types.SuiSystemAddress,
				Module:   sui_system_state.SuiSystemModuleName,
				Function: bfc_types.AddStakeFunName,
				Arguments: []bfc_types.Argument{
					arg0, arg1, arg2,
				},
			},
		},
	)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		*sender, []*bfc_types.ObjectRef{
			coin.Reference(),
			coin2.Reference(),
		},
		pt, gasBudget, gasPrice,
	)

	// case 2: direct stake the specified coin
	// coinArg := bfc_types.CallArg{
	// 	Object: &bfc_types.ObjectArg{
	// 		ImmOrOwnedObject: coin.Reference(),
	// 	},
	// }
	// addrBytes := validatorAddress.Data()
	// addrArg := bfc_types.CallArg{
	// 	Pure: &addrBytes,
	// }
	// err = ptb.MoveCall(
	// 	*bfc_types.SuiSystemAddress,
	// 	sui_system_state.SuiSystemModuleName,
	// 	bfc_types.AddStakeFunName,
	// 	[]move_types.TypeTag{},
	// 	[]bfc_types.CallArg{
	// 		bfc_types.BfcSystemMut,
	// 		coinArg,
	// 		addrArg,
	// 	},
	// )
	// require.NoError(t, err)
	// pt := ptb.Finish()
	// tx := bfc_types.NewProgrammable(
	// 	*sender, []*bfc_types.ObjectRef{
	// 		coin2.Reference(),
	// 	},
	// 	pt, gasBudget, gasPrice,
	// )

	// build & simulate
	txBytesBCS, err := bcs.Marshal(tx)
	require.NoError(t, err)
	resp := simulateCheck(t, cli, txBytesBCS, true)
	t.Log(resp.Effects.Data.GasFee())
}

func GetCoins(t *testing.T, cli *Client, sender BfcAddress, needCount int) []types.Coin {
	coins, err := cli.GetCoins(context.Background(), sender, nil, nil, uint(needCount))
	require.NoError(t, err)
	require.True(t, len(coins.Data) >= needCount)
	return coins.Data
}
