package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/fardream/go-bcs/bcs"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"github.com/hellokittyboy-code/benfen-go-sdk/move_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
)

// NOTE: This copys the query limit from our Rust JSON RPC backend, this needs to be kept in sync!
const QUERY_MAX_RESULT_LIMIT = 1000

type BfcAddress = bfc_types.BfcAddress
type bfcObjectID = bfc_types.ObjectID
type bfcDigest = bfc_types.TransactionDigest
type bfcBase64Data = lib.Base64Data

type DevInspectArgs struct {
	Type    string
	Value   any
	Version uint64
}

// MARK - Getter Function

// GetBalance to use default bfc coin(0x2::bfc::BFC) when coinType is empty
func (c *Client) GetBalance(ctx context.Context, owner BfcAddress, coinType string) (*types.Balance, error) {
	resp := types.Balance{}
	if coinType == "" {
		return &resp, c.CallContext(ctx, &resp, getBalance, owner)
	} else {
		return &resp, c.CallContext(ctx, &resp, getBalance, owner, coinType)
	}
}

func (c *Client) GetAllBalances(ctx context.Context, owner BfcAddress) ([]types.Balance, error) {
	var resp []types.Balance
	return resp, c.CallContext(ctx, &resp, getAllBalances, owner)
}

// GetBfcCoinsOwnedByAddress This function will retrieve a maximum of 200 coins.
func (c *Client) GetBfcCoinsOwnedByAddress(ctx context.Context, address BfcAddress) (types.Coins, error) {
	coinType := types.BFCoinType
	page, err := c.GetCoins(ctx, address, &coinType, nil, 200)
	if err != nil {
		return nil, err
	}
	return page.Data, nil
}

// GetCoins to use default bfc coin(0x2::bfc::BFC) when coinType is nil
// start with the first object when cursor is nil
func (c *Client) GetCoins(
	ctx context.Context,
	owner BfcAddress,

	coinType *string,
	cursor *bfcObjectID,
	limit uint,
) (*types.CoinPage, error) {
	var resp types.CoinPage
	return &resp, c.CallContext(ctx, &resp, getCoins, owner, coinType, cursor, limit)
}

// GetAllCoins
// start with the first object when cursor is nil
func (c *Client) GetAllCoins(
	ctx context.Context,
	owner BfcAddress,
	cursor *bfcObjectID,
	limit uint,
) (*types.CoinPage, error) {
	var resp types.CoinPage
	return &resp, c.CallContext(ctx, &resp, getAllCoins, owner, cursor, limit)
}

func (c *Client) GetCoinMetadata(ctx context.Context, coinType string) (*types.BfcCoinMetadata, error) {
	var resp types.BfcCoinMetadata
	return &resp, c.CallContext(ctx, &resp, getCoinMetadata, coinType)
}

func (c *Client) GetObject(
	ctx context.Context,
	objID bfcObjectID,
	options *types.BfcObjectDataOptions,
) (*types.BfcObjectResponse, error) {
	var resp types.BfcObjectResponse
	return &resp, c.CallContext(ctx, &resp, getObject, objID, options)
}

func (c *Client) MultiGetObjects(
	ctx context.Context,
	objIDs []bfcObjectID,
	options *types.BfcObjectDataOptions,
) ([]types.BfcObjectResponse, error) {
	var resp []types.BfcObjectResponse
	return resp, c.CallContext(ctx, &resp, multiGetObjects, objIDs, options)
}

// address : <BfcAddress> - the owner's Bfc address
// query : <ObjectResponseQuery> - the objects query criteria.
// cursor : <CheckpointedObjectID> - An optional paging cursor. If provided, the query will start from the next item after the specified cursor. Default to start from the first item if not specified.
// limit : <uint> - Max number of items returned per page, default to [QUERY_MAX_RESULT_LIMIT_OBJECTS] if is 0
func (c *Client) GetOwnedObjects(
	ctx context.Context,
	address BfcAddress,
	query *types.BfcObjectResponseQuery,
	cursor *types.CheckpointedObjectId,
	limit *uint,
) (*types.ObjectsPage, error) {
	var resp types.ObjectsPage
	return &resp, c.CallContext(ctx, &resp, getOwnedObjects, address, query, cursor, limit)
}

func (c *Client) GetTotalSupply(ctx context.Context, coinType string) (*types.Supply, error) {
	var resp types.Supply
	return &resp, c.CallContext(ctx, &resp, getTotalSupply, coinType)
}

func (c *Client) GetTotalTransactionBlocks(ctx context.Context) (string, error) {
	var resp string
	return resp, c.CallContext(ctx, &resp, getTotalTransactionBlocks)
}

func (c *Client) GetLatestCheckpointSequenceNumber(ctx context.Context) (string, error) {
	var resp string
	return resp, c.CallContext(ctx, &resp, getLatestCheckpointSequenceNumber)
}

// BatchGetObjectsOwnedByAddress @param filterType You can specify filtering out the specified resources, this will fetch all resources if it is not empty ""
func (c *Client) BatchGetObjectsOwnedByAddress(
	ctx context.Context,
	address BfcAddress,
	options types.BfcObjectDataOptions,
	filterType string,
) ([]types.BfcObjectResponse, error) {
	filterType = strings.TrimSpace(filterType)
	return c.BatchGetFilteredObjectsOwnedByAddress(
		ctx, address, options, func(sod *types.BfcObjectData) bool {
			return filterType == "" || filterType == *sod.Type
		},
	)
}

func (c *Client) BatchGetFilteredObjectsOwnedByAddress(
	ctx context.Context,
	address BfcAddress,
	options types.BfcObjectDataOptions,
	filter func(*types.BfcObjectData) bool,
) ([]types.BfcObjectResponse, error) {
	query := types.BfcObjectResponseQuery{
		Options: &types.BfcObjectDataOptions{
			ShowType: true,
		},
	}
	filteringObjs, err := c.GetOwnedObjects(ctx, address, &query, nil, nil)
	if err != nil {
		return nil, err
	}
	objIds := make([]bfcObjectID, 0)
	for _, obj := range filteringObjs.Data {
		if obj.Data == nil {
			continue // error obj
		}
		if filter != nil && filter(obj.Data) == false {
			continue // ignore objects if non-specified type
		}
		objIds = append(objIds, obj.Data.ObjectId)
	}

	return c.MultiGetObjects(ctx, objIds, &options)
}

func (c *Client) GetTransactionBlock(
	ctx context.Context,
	digest bfcDigest,
	options types.BfcTransactionBlockResponseOptions,
) (*types.BfcTransactionBlockResponse, error) {
	resp := types.BfcTransactionBlockResponse{}
	return &resp, c.CallContext(ctx, &resp, getTransactionBlock, digest, options)
}

func (c *Client) GetTransactionBlockString(
	ctx context.Context,
	digest bfcDigest,
	options types.BfcTransactionBlockResponseOptions,
) (string, error) {
	resp := types.BfcTransactionBlockResponse{}
	return c.CallContextString(ctx, &resp, getTransactionBlock, digest, options)
}

func (c *Client) GetReferenceGasPrice(ctx context.Context) (*types.SafeBfcBigInt[uint64], error) {
	var resp types.SafeBfcBigInt[uint64]
	return &resp, c.CallContext(ctx, &resp, getReferenceGasPrice)
}

func (c *Client) GetEvents(ctx context.Context, digest bfcDigest) ([]types.BfcEvent, error) {
	var resp []types.BfcEvent
	return resp, c.CallContext(ctx, &resp, getEvents, digest)
}

func (c *Client) TryGetPastObject(
	ctx context.Context,
	objectId bfcObjectID,
	version uint64,
	options *types.BfcObjectDataOptions,
) (*types.BfcPastObjectResponse, error) {
	var resp types.BfcPastObjectResponse
	return &resp, c.CallContext(ctx, &resp, tryGetPastObject, objectId, version, options)
}

func (c *Client) DevInspectTransactionBlock(
	ctx context.Context,
	senderAddress BfcAddress,
	txByte bfcBase64Data,
	gasPrice *types.SafeBfcBigInt[uint64],
	epoch *uint64,
) (*types.DevInspectResults, error) {
	var resp types.DevInspectResults
	return &resp, c.CallContext(ctx, &resp, devInspectTransactionBlock, senderAddress, txByte, gasPrice, epoch)
}

func (c *Client) DryRunTransaction(
	ctx context.Context,
	txBytes bfcBase64Data,
) (*types.DryRunTransactionBlockResponse, error) {
	var resp types.DryRunTransactionBlockResponse
	return &resp, c.CallContext(ctx, &resp, dryRunTransactionBlock, txBytes)
}

func (c *Client) ExecuteTransactionBlock(
	ctx context.Context, txBytes bfcBase64Data, signatures []any,
	options *types.BfcTransactionBlockResponseOptions, requestType types.ExecuteTransactionRequestType,
) (*types.BfcTransactionBlockResponse, error) {
	resp := types.BfcTransactionBlockResponse{}
	return &resp, c.CallContext(ctx, &resp, executeTransactionBlock, txBytes, signatures, options, requestType)
}

func (c *Client) ExecuteTransactionBlockStr(
	ctx context.Context, txBytes string, signatures []string,
	options *types.BfcTransactionBlockResponseOptions, requestType types.ExecuteTransactionRequestType,
) (*types.BfcTransactionBlockResponse, error) {
	resp := types.BfcTransactionBlockResponse{}
	return &resp, c.CallContext(ctx, &resp, executeTransactionBlock, txBytes, signatures, options, requestType)
}

// TransferObject Create an unsigned transaction to transfer an object from one address to another. The object's type must allow public transfers
func (c *Client) TransferObject(
	ctx context.Context,
	signer, recipient BfcAddress,
	objID bfcObjectID,
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, transferObject, signer, objID, gas, gasBudget, recipient)
}

// TransferBfc Create an unsigned transaction to send BFC coin object to a Bfc address.
// The BFC object is also used as the gas object.
func (c *Client) TransferBFC(
	ctx context.Context, signer, recipient BfcAddress, ObjID bfcObjectID, amount,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, transferBfc, signer, ObjID, gasBudget, recipient, amount)
}

// PayAllBfc Create an unsigned transaction to send all BFC coins to one recipient.
func (c *Client) PayAllBFC(
	ctx context.Context,
	signer, recipient BfcAddress,
	inputCoins []bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, payAllBfc, signer, inputCoins, recipient, gasBudget)
}

func (c *Client) Pay(
	ctx context.Context,
	signer BfcAddress,
	inputCoins []bfcObjectID,
	recipients []BfcAddress,
	amount []types.SafeBfcBigInt[uint64],
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, pay, signer, inputCoins, recipients, amount, gas, gasBudget)
}

func (c *Client) PayBFC(
	ctx context.Context,
	signer BfcAddress,
	inputCoins []bfcObjectID,
	recipients []BfcAddress,
	amount []types.SafeBfcBigInt[uint64],
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, payBfc, signer, inputCoins, recipients, amount, gasBudget)
}

// SplitCoin Create an unsigned transaction to split a coin object into multiple coins.
func (c *Client) SplitCoin(
	ctx context.Context,
	signer BfcAddress,
	Coin bfcObjectID,
	splitAmounts []types.SafeBfcBigInt[uint64],
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, splitCoin, signer, Coin, splitAmounts, gas, gasBudget)
}

// SplitCoinEqual Create an unsigned transaction to split a coin object into multiple equal-size coins.
func (c *Client) SplitCoinEqual(
	ctx context.Context,
	signer BfcAddress,
	Coin bfcObjectID,
	splitCount types.SafeBfcBigInt[uint64],
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, splitCoinEqual, signer, Coin, splitCount, gas, gasBudget)
}

// MergeCoins Create an unsigned transaction to merge multiple coins into one coin.
func (c *Client) MergeCoins(
	ctx context.Context,
	signer BfcAddress,
	primaryCoin, coinToMerge bfcObjectID,
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, mergeCoins, signer, primaryCoin, coinToMerge, gas, gasBudget)
}

func (c *Client) Publish(
	ctx context.Context,
	sender BfcAddress,
	compiledModules []*bfcBase64Data,
	dependencies []bfcObjectID,
	gas bfcObjectID,
	gasBudget uint,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, publish, sender, compiledModules, dependencies, gas, gasBudget)
}

// MoveCall Create an unsigned transaction to execute a Move call on the network, by calling the specified function in the module of a given package.
// TODO: not support param `typeArguments` yet.
// So now only methods with `typeArguments` are supported
// TODO: execution_mode : <BfcTransactionBlockBuilderMode>
func (c *Client) MoveCall(
	ctx context.Context,
	signer BfcAddress,
	packageId bfcObjectID,
	module, function string,
	typeArgs []string,
	arguments []any,
	gas *bfcObjectID,
	gasBudget types.SafeBfcBigInt[uint64],
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(
		ctx,
		&resp,
		moveCall,
		signer,
		packageId,
		module,
		function,
		typeArgs,
		arguments,
		gas,
		gasBudget,
	)
}

//getLoadedChildObjects

// TODO: execution_mode : <BfcTransactionBlockBuilderMode>
func (c *Client) BatchTransaction(
	ctx context.Context,
	signer BfcAddress,
	txnParams []map[string]interface{},
	gas *bfcObjectID,
	gasBudget string,
) (*types.TransactionBytes, error) {
	resp := types.TransactionBytes{}
	return &resp, c.CallContext(ctx, &resp, batchTransaction, signer, txnParams, gas, gasBudget)
}

func (c *Client) QueryTransactionBlocks(
	ctx context.Context, query types.BfcTransactionBlockResponseQuery,
	cursor *bfcDigest, limit *uint, descendingOrder bool,
) (*types.TransactionBlocksPage, error) {
	resp := types.TransactionBlocksPage{}
	return &resp, c.CallContext(ctx, &resp, queryTransactionBlocks, query, cursor, limit, descendingOrder)
}

func (c *Client) QueryEvents(
	ctx context.Context, query types.EventFilter, cursor *types.EventId, limit *uint,
	descendingOrder bool,
) (*types.EventPage, error) {
	var resp types.EventPage
	return &resp, c.CallContext(ctx, &resp, queryEvents, query, cursor, limit, descendingOrder)
}

func (c *Client) GetDynamicFields(
	ctx context.Context, parentObjectId bfcObjectID, cursor *bfcObjectID,
	limit *uint,
) (*types.DynamicFieldPage, error) {
	var resp types.DynamicFieldPage
	return &resp, c.CallContext(ctx, &resp, getDynamicFields, parentObjectId, cursor, limit)
}

func (c *Client) GetInnerDaoInfo(
	ctx context.Context,
) (string, error) {
	var resp types.BfcObjectResponse
	return c.CallContextString(ctx, &resp, getInnerDaoInfo)
}

func (c *Client) GetStablePools(
	ctx context.Context,
	bfcaddress BfcAddress,
) (string, error) {
	var resp types.BfcObjectResponse
	return c.CallContextString(ctx, &resp, getStablePools, bfcaddress)
}

func (c *Client) GetProposalInfo(
	ctx context.Context,
	bfcaddress BfcAddress,
) (string, error) {
	var resp types.BfcObjectResponse
	return c.CallContextString(ctx, &resp, getProposal, bfcaddress)
}

func (c *Client) GetDynamicFieldObject(
	ctx context.Context,
	parentObjectId bfcObjectID,
	name bfc_types.DynamicFieldName,
) (*types.BfcObjectResponse, error) {
	var resp types.BfcObjectResponse
	return &resp, c.CallContext(ctx, &resp, getDynamicFieldObject, parentObjectId, name)
}

func (c *Client) GetChainIdentifier(ctx context.Context) (*string, error) {
	var resp string
	return &resp, c.CallContext(ctx, &resp, getChainIdentifier)
}

/*
input target: 0xc8::bfc_system::vault_info
return: 0xc8, bfc_system, vault_info
*/
func (c *Client) GetFunctions(target string) (*BfcAddress, move_types.Identifier, move_types.Identifier, error) {
	parts := strings.Split(target, "::")
	if len(parts) != 3 {
		return nil, "", "", fmt.Errorf("target %v is error", target)
	}
	packageId, err := bfc_types.NewObjectIdFromHex(parts[0])
	if err != nil {
		return nil, "", "", err
	}
	return packageId, move_types.Identifier(parts[1]), move_types.Identifier(parts[2]), nil
}

/*
DevInspectTransactionBlockV2(

		"0xc8::bfc_system::vault_info",
		[]string{
			"0xc8::busd::BUSD",
		},
		[]*DevInspectArgs{
			{
				Type:    "object",
				Value:   "0xc9",
				Version: 1,
			},
		},
	)
*/
func (c *Client) DevInspectTransactionBlockV2(
	ctx context.Context,
	senderAddress BfcAddress,
	target string,
	typeArgs []string,
	args []*DevInspectArgs,
) (*types.DevInspectResults, error) {
	packageId, module, function, err := c.GetFunctions(target)
	if err != nil {
		return nil, err
	}

	ptb := bfc_types.NewProgrammableTransactionBuilder()

	// make TypeTags
	typeTags := []move_types.TypeTag{}
	for _, arg := range typeArgs {
		typeAddr, typeModule, typeFunction, err := c.GetFunctions(arg)
		if err != nil {
			return nil, fmt.Errorf("type tag %v is error", arg)
		}
		typeTags = append(
			typeTags, move_types.TypeTag{
				Struct: &move_types.StructTag{
					Address: *typeAddr,
					Module:  typeModule,
					Name:    typeFunction,
				},
			},
		)
	}

	// make CallArgs
	arguments := []bfc_types.CallArg{}
	for _, arg := range args {
		var a bfc_types.CallArg
		if arg.Type == "object" {
			a, _ = ptb.SharedObjCallArg(arg.Value.(string), arg.Version)
		} else if arg.Type == "pure" {
			a, _ = ptb.PureCallArg(arg.Value)
		} else {
			continue
		}
		arguments = append(arguments, a)
	}

	ptb.MoveCall(*packageId, module, function, typeTags, arguments)
	pt := ptb.Finish()

	// Get tx bytes
	tx := bfc_types.NewProgrammable(senderAddress, nil, pt, 0, 0)
	jsonBytes, err := bcs.Marshal(tx)
	if err != nil {
		return nil, err
	}
	txBytes := jsonBytes[1 : len(jsonBytes)-82]

	// Request
	return c.DevInspectTransactionBlock(ctx, senderAddress, txBytes, nil, nil)
}
