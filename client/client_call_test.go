package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/stretchr/testify/require"
)

//func TestClient_BatchGetTransaction(t *testing.T) {
//	chain := ChainClient(t)
//	coins, err := chain.GetCoins(context.TODO(), *Address, nil, nil, 1)
//	require.NoError(t, err)
//	object, err := chain.GetObject(context.TODO(), coins.Data[0].CoinObjectId, nil)
//	require.NoError(t, err)
//	type args struct {
//		digests []string
//	}
//	tests := []struct {
//		name    string
//		chain   *Client
//		args    args
//		want    int
//		wantErr bool
//	}{
//		{
//			name:  "test for devnet transaction",
//			chain: chain,
//			args: args{
//				digests: []string{*object.Data.PreviousTransaction},
//			},
//			want:    1,
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.chain.BatchGetTransaction(tt.args.digests)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("BatchGetTransaction() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if len(got) != tt.want {
//				t.Errorf("BatchGetTransaction() got = %v, want %v", got, tt.want)
//			}
//			t.Logf("%+v", got)
//		})
//	}
//}

func Test_TagJson_Owner(t *testing.T) {
	test := func(str string) lib.TagJson[bfc_types.Owner] {
		var s lib.TagJson[bfc_types.Owner]
		data := []byte(str)
		err := json.Unmarshal(data, &s)
		require.Nil(t, err)
		return s
	}
	{
		v := test(`"Immutable"`).Data
		require.Nil(t, v.AddressOwner)
		require.Nil(t, v.ObjectOwner)
		require.Nil(t, v.Shared)
		require.NotNil(t, v.Immutable)
	}
	{
		v := test(`{"AddressOwner": "0x7e875ea78ee09f08d72e2676cf84e0f1c8ac61d94fa339cc8e37cace85bebc6e"}`).Data
		require.NotNil(t, v.AddressOwner)
		require.Nil(t, v.ObjectOwner)
		require.Nil(t, v.Shared)
		require.Nil(t, v.Immutable)
	}
}

func TestClient_DryRunTransaction(t *testing.T) {
	cli := ChainClient(t)
	signer := Address
	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.NoError(t, err)

	amount := BFC(0.01).Uint64()
	gasBudget := BFC(0.01).Uint64()
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), gasBudget, 0, 0)
	require.NoError(t, err)
	tx, err := cli.PayAllBFC(
		context.Background(), *signer, *signer,
		pickedCoins.CoinIds(),
		types.NewSafeBfcBigInt(gasBudget),
	)
	require.NoError(t, err)

	resp, err := cli.DryRunTransaction(context.Background(), tx.TxBytes)
	require.Nil(t, err)
	t.Log("dry run status:", resp.Effects.Data.IsSuccess())
	t.Log("dry run error:", resp.Effects.Data.V1.Status.Error)
}

// TestClient_ExecuteTransactionSerializedSig
// This test case will affect the real coin in the test case of account
// temporary disabled
//func TestClient_ExecuteTransactionSerializedSig(t *testing.T) {
//	chain := ChainClient(t)
//	coins, err := chain.GetBfcCoinsOwnedByAddress(context.TODO(), *Address)
//	require.NoError(t, err)
//	coin, err := coins.PickCoinNoLess(2000)
//	require.NoError(t, err)
//	tx, err := chain.TransferBFC(context.TODO(), *Address, *Address, coin.Reference.ObjectId, 1000, 1000)
//	require.NoError(t, err)
//	account := M1Account(t)
//	signedTx := tx.SignSerializedSigWith(account.PrivateKey)
//	txResult, err := chain.ExecuteTransactionSerializedSig(context.TODO(), *signedTx, types.TxnRequestTypeWaitForEffectsCert)
//	require.NoError(t, err)
//	t.Logf("%#v", txResult)
//}

//func TestClient_ExecuteTransaction(t *testing.T) {
//	chain := ChainClient(t)
//	coins, err := chain.GetBfcCoinsOwnedByAddress(context.TODO(), *Address)
//	require.NoError(t, err)
//	coin, err := coins.PickCoinNoLess(2000)
//	require.NoError(t, err)
//	tx, err := chain.TransferBFC(context.TODO(), *Address, *Address, coin.Reference.ObjectId, 1000, 1000)
//	require.NoError(t, err)
//	account := M1Account(t)
//	signedTx := tx.SignSerializedSigWith(account.PrivateKey)
//	txResult, err := chain.ExecuteTransaction(context.TODO(), *signedTx, types.TxnRequestTypeWaitForEffectsCert)
//	require.NoError(t, err)
//	t.Logf("%#v", txResult)
//}

func TestClient_BatchGetObjectsOwnedByAddress(t *testing.T) {
	cli := ChainClient(t)

	options := types.BfcObjectDataOptions{
		ShowType:    true,
		ShowContent: true,
	}
	coinType := fmt.Sprintf("0x2::coin::Coin<%v>", types.BFCoinType)
	filterObject, err := cli.BatchGetObjectsOwnedByAddress(context.TODO(), *Address, options, coinType)
	require.NoError(t, err)
	t.Log(filterObject)
}

func TestClient_GetCoinMetadata(t *testing.T) {
	chain := ChainClient(t)
	metadata, err := chain.GetCoinMetadata(context.TODO(), types.BFCoinType)
	require.Nil(t, err)
	t.Logf("%#v", metadata)
}

func TestClient_GetAllBalances(t *testing.T) {
	chain := ChainClient(t)
	balances, err := chain.GetAllBalances(context.TODO(), *Address)
	require.NoError(t, err)
	for _, balance := range balances {
		t.Logf(
			"Coin Name: %v, Count: %v, Total: %v, Locked: %v",
			balance.CoinType, balance.CoinObjectCount,
			balance.TotalBalance.String(), balance.LockedBalance,
		)
	}
}

func TestClient_GetBalance(t *testing.T) {
	chain := ChainClient(t)
	balance, err := chain.GetBalance(context.TODO(), *Address, "")
	require.NoError(t, err)
	t.Logf(
		"Coin Name: %v, Count: %v, Total: %v, Locked: %v",
		balance.CoinType, balance.CoinObjectCount,
		balance.TotalBalance.String(), balance.LockedBalance,
	)
}

func TestClient_GetCoins(t *testing.T) {
	chain := ChainClient(t)
	defaultCoinType := types.BFCoinType
	coins, err := chain.GetCoins(context.TODO(), *Address, &defaultCoinType, nil, 1)
	require.NoError(t, err)
	t.Logf("%#v", coins)
}

func TestClient_GetAllCoins(t *testing.T) {
	chain := ChainClient(t)
	type args struct {
		ctx     context.Context
		address BfcAddress
		cursor  *bfcObjectID
		limit   uint
	}
	tests := []struct {
		name    string
		chain   *Client
		args    args
		want    *types.CoinPage
		wantErr bool
	}{
		{
			name:  "test case 1",
			chain: chain,
			args: args{
				ctx:     context.TODO(),
				address: *Address,
				cursor:  nil,
				limit:   3,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := chain.GetAllCoins(tt.args.ctx, tt.args.address, tt.args.cursor, tt.args.limit)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetAllCoins() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func TestClient_GetTransaction(t *testing.T) {
	cli := ChainClient(t)
	digest := "DRdkHqn1J7UsN8wEGifNQ71qhSi6AG7WegG68sx6Yay5"
	d, err := bfc_types.NewDigest(digest)
	require.Nil(t, err)
	resp, err := cli.GetTransactionBlock(
		context.Background(), *d, types.BfcTransactionBlockResponseOptions{
			ShowInput:          true,
			ShowEffects:        true,
			ShowObjectChanges:  true,
			ShowBalanceChanges: true,
			ShowEvents:         true,
		},
	)
	//require.NoError(t, err)
	t.Logf("the response is %#v", resp)

	//require.Equal(t, int64(1997880), resp.Effects.Data.GasFee())
}

func TestBatchCall_GetObject(t *testing.T) {
	cli := ChainClient(t)

	if false {
		// get sepcified object
		idstr := "0x4ad2f0a918a241d6a19573212aeb56947bb9255a14e921a7ec78b262536826f0"
		objId, err := bfc_types.NewAddressFromHex(idstr)
		require.Nil(t, err)
		obj, err := cli.GetObject(
			context.Background(), *objId, &types.BfcObjectDataOptions{
				ShowType:    true,
				ShowContent: true,
			},
		)
		require.Nil(t, err)
		t.Log(obj.Data)
	}

	coins, err := cli.GetCoins(context.TODO(), *Address, nil, nil, 3)
	require.NoError(t, err)
	if len(coins.Data) == 0 {
		return
	}
	objId := coins.Data[0].CoinObjectId
	obj, err := cli.GetObject(context.Background(), objId, nil)
	require.Nil(t, err)
	t.Log(obj.Data)
}

func TestClient_GetObject(t *testing.T) {
	type args struct {
		ctx   context.Context
		objID bfcObjectID
	}
	chain := ChainClient(t)
	coins, err := chain.GetCoins(context.TODO(), *Address, nil, nil, 1)
	require.NoError(t, err)

	tests := []struct {
		name    string
		chain   *Client
		args    args
		want    int
		wantErr bool
	}{
		{
			name:  "test for devnet",
			chain: chain,
			args: args{
				ctx:   context.TODO(),
				objID: coins.Data[0].CoinObjectId,
			},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := tt.chain.GetObject(
					tt.args.ctx, tt.args.objID, &types.BfcObjectDataOptions{
						ShowType:                true,
						ShowOwner:               true,
						ShowContent:             true,
						ShowDisplay:             true,
						ShowBcs:                 true,
						ShowPreviousTransaction: true,
						ShowStorageRebate:       true,
					},
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetObject() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%+v", got)
			},
		)
	}
}

func TestClient_MultiGetObjects(t *testing.T) {
	chain := ChainClient(t)
	coins, err := chain.GetCoins(context.TODO(), *Address, nil, nil, 1)
	require.NoError(t, err)
	if len(coins.Data) == 0 {
		t.Log("Warning: No Object Id for test.")
		return
	}

	obj := coins.Data[0].CoinObjectId
	objs := []bfcObjectID{obj, obj}
	resp, err := chain.MultiGetObjects(
		context.Background(), objs, &types.BfcObjectDataOptions{
			ShowType:                true,
			ShowOwner:               true,
			ShowContent:             true,
			ShowDisplay:             true,
			ShowBcs:                 true,
			ShowPreviousTransaction: true,
			ShowStorageRebate:       true,
		},
	)
	require.Nil(t, err)
	require.Equal(t, len(objs), len(resp))
	require.Equal(t, resp[0], resp[1])
}

func TestClient_GetOwnedObjects(t *testing.T) {
	cli := ChainClient(t)

	obj, err := bfc_types.NewAddressFromHex("0x2")
	require.Nil(t, err)
	query := types.BfcObjectResponseQuery{
		Filter: &types.BfcObjectDataFilter{
			Package: obj,
			// StructType: "0x2::coin::Coin<0x2::bfc::BFC>",
		},
		Options: &types.BfcObjectDataOptions{
			ShowType: true,
		},
	}
	limit := uint(1)
	objs, err := cli.GetOwnedObjects(context.Background(), *Address, &query, nil, &limit)
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(objs.Data), int(limit))
}

func TestClient_GetTotalSupply(t *testing.T) {
	chain := ChainClient(t)
	type args struct {
		ctx      context.Context
		coinType string
	}
	tests := []struct {
		name    string
		chain   *Client
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:  "test 1",
			chain: chain,
			args: args{
				context.TODO(),
				types.BFCoinType,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := tt.chain.GetTotalSupply(tt.args.ctx, tt.args.coinType)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetTotalSupply() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%d", got)
			},
		)
	}
}
func TestClient_GetTotalTransactionBlocks(t *testing.T) {
	cli := ChainClient(t)
	res, err := cli.GetTotalTransactionBlocks(context.Background())
	require.Nil(t, err)
	t.Log(res)
}

func TestClient_GetLatestCheckpointSequenceNumber(t *testing.T) {
	cli := ChainClient(t)
	res, err := cli.GetLatestCheckpointSequenceNumber(context.Background())
	require.Nil(t, err)
	t.Log(res)
}

//func TestClient_Publish(t *testing.T) {
//	chain := ChainClient(t)
//	dmens, err := types.NewBase64Data(DmensDmensB64)
//	require.NoError(t, err)
//	profile, err := types.NewBase64Data(DmensProfileB64)
//	require.NoError(t, err)
//	coins, err := chain.GetBfcCoinsOwnedByAddress(context.TODO(), *Address)
//	require.NoError(t, err)
//	coin, err := coins.PickCoinNoLess(30000)
//	require.NoError(t, err)
//	type args struct {
//		ctx             context.Context
//		address         types.Address
//		compiledModules []*types.Base64Data
//		gas             types.ObjectId
//		gasBudget       uint
//	}
//	tests := []struct {
//		name    string
//		client  *Client
//		args    args
//		want    *types.TransactionBytes
//		wantErr bool
//	}{
//		{
//			name:   "test for dmens publish",
//			client: chain,
//			args: args{
//				ctx:             context.TODO(),
//				address:         *Address,
//				compiledModules: []*types.Base64Data{dmens, profile},
//				gas:             coin.CoinObjectId,
//				gasBudget:       30000,
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.client.Publish(tt.args.ctx, tt.args.address, tt.args.compiledModules, tt.args.gas, tt.args.gasBudget)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			t.Logf("%#v", got)
//
//			txResult, err := tt.client.DryRunTransaction(context.TODO(), got)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			t.Logf("%#v", txResult)
//		})
//	}
//}

func TestClient_TryGetPastObject(t *testing.T) {
	cli := ChainClient(t)
	objId, err := bfc_types.NewAddressFromHex("0x11462c88e74bb00079e3c043efb664482ee4551744ee691c7623b98503cb3f4d")
	require.Nil(t, err)
	data, err := cli.TryGetPastObject(context.Background(), *objId, 903, nil)
	require.Nil(t, err)
	t.Log(data)
}

// change to corrent event hash digest. 6vRkJ5TconcXL8SEsMAN3ZDU4dcFtJv6AVndXPNgn49Q
func TestClient_GetEvents(t *testing.T) {
	cli := ChainClient(t)
	digest := "9wfNFxA8Vd8kmmw4JYvNvqiNsr6V1uoTx8zgybrxf2MY"
	d, err := bfc_types.NewDigest(digest)
	require.Nil(t, err)
	res, err := cli.GetEvents(context.Background(), *d)
	require.NoError(t, err)
	t.Log(res)
}

func TestClient_GetReferenceGasPrice(t *testing.T) {
	cli := ChainClient(t)
	gasPrice, err := cli.GetReferenceGasPrice(context.Background())
	require.Nil(t, err)
	t.Logf("current gas price = %v", gasPrice)
}

// func TestClient_DevInspectTransactionBlock(t *testing.T) {
// 	chain := ChainClient(t)
// 	signer := Address
// 	price, err := chain.GetReferenceGasPrice(context.TODO())
// 	require.NoError(t, err)
// 	coins, err := chain.GetCoins(context.Background(), *signer, nil, nil, 10)
// 	require.NoError(t, err)

// 	amount := BFC(0.01).Int64()
// 	gasBudget := BFC(0.01).Uint64()
// 	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(amount * 2), 0, false)
// 	require.NoError(t, err)
// 	tx, err := chain.PayAllBFC(context.Background(),
// 		*signer, *signer,
// 		pickedCoins.CoinIds(),
// 		types.NewSafeBfcBigInt(gasBudget))
// 	require.NoError(t, err)

// 	resp, err := chain.DevInspectTransactionBlock(context.Background(), *signer, tx.TxBytes, price, nil)
// 	require.Nil(t, err)
// 	t.Log(resp)
// }

func TestClient_QueryTransactionBlocks(t *testing.T) {
	cli := ChainClient(t)
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           types.BfcTransactionBlockResponseQuery
		cursor          *bfcDigest
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *types.TransactionBlocksPage
		wantErr bool
	}{
		{
			name: "test for queryTransactionBlocks",
			args: args{
				ctx: context.TODO(),
				query: types.BfcTransactionBlockResponseQuery{
					Filter: &types.TransactionFilter{
						FromAddress: Address,
					},
					Options: &types.BfcTransactionBlockResponseOptions{
						ShowInput:   true,
						ShowEffects: true,
					},
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := cli.QueryTransactionBlocks(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryTransactionBlocks() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("%#v", got)
			},
		)
	}
}

func TestClient_QueryEvents(t *testing.T) {
	cli := ChainClient(t)
	limit := uint(10)
	type args struct {
		ctx             context.Context
		query           types.EventFilter
		cursor          *types.EventId
		limit           *uint
		descendingOrder bool
	}
	tests := []struct {
		name    string
		args    args
		want    *types.EventPage
		wantErr bool
	}{
		{
			name: "test for query events",
			args: args{
				ctx: context.TODO(),
				query: types.EventFilter{
					Sender: Address,
				},
				cursor:          nil,
				limit:           &limit,
				descendingOrder: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := cli.QueryEvents(
					tt.args.ctx,
					tt.args.query,
					tt.args.cursor,
					tt.args.limit,
					tt.args.descendingOrder,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("QueryEvents() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

func TestClient_GetDynamicFields(t *testing.T) {
	chain := ChainClient(t)
	parentObjectId, err := bfc_types.NewAddressFromHex(
		"0xbc0957602f610d0992d32d6ce601109931f7e546b6063783b66601f6552c590c",
	)
	require.NoError(t, err)
	limit := uint(5)
	type args struct {
		ctx            context.Context
		parentObjectId bfcObjectID
		cursor         *bfcObjectID
		limit          *uint
	}
	tests := []struct {
		name    string
		args    args
		want    *types.DynamicFieldPage
		wantErr bool
	}{
		{
			name: "case 1",
			args: args{
				ctx:            context.TODO(),
				parentObjectId: *parentObjectId,
				cursor:         nil,
				limit:          &limit,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := chain.GetDynamicFields(tt.args.ctx, tt.args.parentObjectId, tt.args.cursor, tt.args.limit)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetDynamicFields() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Log(got)
			},
		)
	}
}

//todo: need find dynamic test data.
//func TestClient_GetDynamicFieldObject(t *testing.T) {
//	chain := ChainClient(t)
//	parentObjectId, err := bfc_types.NewAddressFromHex("0xbc0957602f610d0992d32d6ce601109931f7e546b6063783b66601f6552c590c")
//	require.NoError(t, err)
//	type args struct {
//		ctx            context.Context
//		parentObjectId bfcObjectID
//		name           bfc_types.DynamicFieldName
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *types.BfcObjectResponse
//		wantErr bool
//	}{
//		{
//			name: "case 1",
//			args: args{
//				ctx:            context.TODO(),
//				parentObjectId: *parentObjectId,
//				name: bfc_types.DynamicFieldName{
//					Type:  "address",
//					Value: "0xf9ed7d8de1a6c44d703b64318a1cc687c324fdec35454281035a53ea3ba1a95a",
//				},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(
//			tt.name, func(t *testing.T) {
//				got, err := chain.GetDynamicFieldObject(tt.args.ctx, tt.args.parentObjectId, tt.args.name)
//				if (err != nil) != tt.wantErr {
//					t.Errorf("GetDynamicFieldObject() error = %v, wantErr %v", err, tt.wantErr)
//					return
//				}
//				t.Logf("%#v", got)
//			},
//		)
//	}
//}

//func TestClient_DevInspectTransactionBlockV2(t *testing.T) {
//	chain, err := Dial("http://yobcrpc.openblock.vip")
//	require.NoError(t, err)
//
//	fromAddr, err := bfc_types.NewAddressFromHex("0x0")
//	require.NoError(t, err)
//
//	resp, err := chain.DevInspectTransactionBlockV2(
//		context.Background(),
//		*fromAddr,
//		"0xc8::bfc_system::vault_info",
//		[]string{
//			"0xc8::busd::BUSD",
//		},
//		[]*DevInspectArgs{
//			{
//				Type:    "object",
//				Value:   "0xc9",
//				Version: 1,
//			},
//		},
//	)
//	require.NoError(t, err)
//
//	// parse result
//	resBytes := []byte{}
//	values := resp.Results[0].ReturnValues[0]
//	interfaceData := values.([]interface{})[0]
//	switch interfaceData.(type) {
//	case []interface{}:
//		for _, item := range interfaceData.([]interface{}) {
//			resBytes = append(resBytes, uint8(item.(float64)))
//		}
//	}
//
//	var vaultInfo struct {
//		VaultId          bfc_types.BfcAddress  `json:"vault_id"`
//		PositionNumber   uint32                `json:"position_number"`
//		State            uint8                 `json:"state"`
//		StateCounter     uint32                `json:"state_counter"`
//		MaxCounterTimes  uint32                `json:"max_counter_times"`
//		LastSqrtPrice    move_types.U128       `json:"last_sqrt_price"`
//		CoinABalance     uint64                `json:"coin_a_balance"`
//		CoinBBalance     uint64                `json:"coin_b_balance"`
//		CoinAType        move_types.Identifier `json:"coin_a_type"`
//		CoinBType        move_types.Identifier `json:"coin_b_type"`
//		TickSpacing      uint32                `json:"tick_spacing"`
//		SpacingTimes     uint32                `json:"spacing_times"`
//		Liquidity        move_types.U128       `json:"liquidity"`
//		CurrentSqrtPrice move_types.U128       `json:"current_sqrt_price"`
//		CurrentTickIndex uint32                `json:"current_tick_index"`
//		IsPause          bool                  `json:"is_pause"`
//		Index            uint64                `json:"index"`
//		BasePoint        uint64                `json:"base_point"`
//	}
//	bcs.Unmarshal(resBytes, &vaultInfo)
//	PrintJson(vaultInfo)
//}

func TestClient_GetChainIdentifier(t *testing.T) {
	//chain, err := Dial("https://obcrpc.openblock.vip")
	chain := ChainClient(t)
	//require.NoError(t, err)

	resp, err := chain.GetChainIdentifier(
		context.Background(),
	)
	require.NoError(t, err)

	PrintJson(resp)
}

func PrintJson(data any) {
	body, _ := json.Marshal(data)
	var str bytes.Buffer
	_ = json.Indent(&str, body, "", "    ")
	fmt.Println(str.String())
}
