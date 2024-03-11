package client

import (
	"context"

	"github.com/fardream/go-bcs/bcs"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types/sui_system_state"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
)

func (c *Client) GetLatestSuiSystemState(ctx context.Context) (*types.BfcSystemStateSummary, error) {
	var resp types.BfcSystemStateSummary
	return &resp, c.CallContext(ctx, &resp, getLatestSuiSystemState)
}

func (c *Client) GetValidatorsApy(ctx context.Context) (*types.ValidatorsApy, error) {
	var resp types.ValidatorsApy
	return &resp, c.CallContext(ctx, &resp, getValidatorsApy)
}
func (c *Client) GetProtocolConfig(ctx context.Context) (string, error) {
	var resp types.ProtocolConfig
	return c.CallContextString(ctx, &resp, getProtocolConfig)
}

// CommiteeInfo
func (c *Client) GetCommiteeInfo(ctx context.Context) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getCommitteeInfo)
}

// getMoveFunctionArgTypes
func (c *Client) getMoveFunctionArgTypes(
	ctx context.Context,
	pkg *bfcObjectID, module string,
	function string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getMoveFunctionArgTypes, pkg, module, function)
}

// getNormalizedMoveFunction
func (c *Client) getNormalizedMoveFunction(
	ctx context.Context,
	pkg *bfcObjectID,
	module string,
	function string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getNormalizedMoveFunction, pkg, module, function)
}

// publish
func (c *Client) publish(
	ctx context.Context,
	sender *bfcObjectID,
	modules []string,
	dependencies []bfcObjectID,
	gas *bfcObjectID,
	gasBudget string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, publish, sender, modules, dependencies, gas, gasBudget)
}

// resolveNameServiceNames
func (c *Client) resolveNameServiceNames(
	ctx context.Context,
	address *bfcObjectID,
	cursor *bfcObjectID,
	limit uint,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, resolveNameServiceNames, address, cursor, limit)
}

// resolveNameServiceAddress
func (c *Client) resolveNameServiceAddress(
	ctx context.Context,
	name string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, resolveNameServiceAddress, name)
}

// devInspectTransactionBlock
func (c *Client) devInspectTransactionBlock(
	ctx context.Context,
	sender *bfcObjectID,
	txbytes string,
	gas *bfcObjectID,
	gasBudget string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, devInspectTransactionBlock, sender, txbytes, gas, gasBudget)
}

// getLoadedChildObjects
func (c *Client) getLoadedChildObjects(
	ctx context.Context,
	txHash string,

) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getLoadedChildObjects, txHash)
}

// executeTransactionBlock
func (c *Client) executeTransactionBlock(
	ctx context.Context,
	txbytes string,
	signature []string,
	options *types.BfcTransactionBlockResponseOptions,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, executeTransactionBlock, txbytes, signature, options)
}

// subscribeEvent
func (c *Client) SubscribeEvent(
	ctx context.Context,
	sender *bfcObjectID,
	// packageId *bfcObjectID,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, subscribeEvent, sender)
}

// getNormalizedMoveModulesByPackage
func (c *Client) getNormalizedMoveModulesByPackage(
	ctx context.Context,
	pkg *bfcObjectID,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getNormalizedMoveModulesByPackage, pkg)
}

// getNormalizedMoveStruct
func (c *Client) getNormalizedMoveStruct(
	ctx context.Context,
	pkg *bfcObjectID,
	module string,
	structName string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getNormalizedMoveStruct, pkg, module, structName)
}

// getNormalizedMoveModule
func (c *Client) getNormalizedMoveModule(
	ctx context.Context,
	pkg *bfcObjectID,
	module string,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, getNormalizedMoveModule, pkg, module)
}

// tryMultiGetPastObjects
func (c *Client) tryMultiGetPastObjects(
	ctx context.Context,
	pastObjects []types.GetPastObjectRequest,
	options *types.BfcObjectDataOptions,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, tryMultiGetPastObjects, pastObjects, options)
}

// multiGetTransactionBlocks
func (c *Client) multiGetTransactionBlocks(
	ctx context.Context,
	digests []bfcDigest,
	options *types.BfcTransactionBlockResponseOptions,
) (string, error) {
	var resp types.CommonStruct
	return c.CallContextString(ctx, &resp, multiGetTransactionBlocks, digests, options)
}

func (c *Client) GetCheckPoints(
	ctx context.Context,
	cursor *bfcObjectID,
	limit uint,
	descendingOrder bool,
) (string, error) {
	var resp types.CheckPointsPage
	return c.CallContextString(ctx, &resp, getCheckpoints, cursor, limit, descendingOrder)
}

func (c *Client) GetEpochs(ctx context.Context, epochNum string, limit int64) (types.EpochPage, error) {
	var resp types.EpochPage
	return resp, c.CallContext(ctx, &resp, getEpochs, epochNum, limit, true)
}

func (c *Client) GetCheckPoint(ctx context.Context, CheckpointId string) (string, error) {
	var resp types.CheckPointsPage
	return c.CallContextString(ctx, &resp, getCheckpoint, CheckpointId)
}

func (c *Client) GetStakes(ctx context.Context, owner BfcAddress) ([]types.DelegatedStake, error) {
	var resp []types.DelegatedStake
	return resp, c.CallContext(ctx, &resp, getStakes, owner)
}

func (c *Client) GetStakesByIds(ctx context.Context, stakedBfcIds []bfcObjectID) ([]types.DelegatedStake, error) {
	var resp []types.DelegatedStake
	return resp, c.CallContext(ctx, &resp, getStakesByIds, stakedBfcIds)
}

func (c *Client) RequestAddStake(
	ctx context.Context,
	signer BfcAddress,
	coins []bfcObjectID,
	amount types.BfcBigInt,
	validator BfcAddress,
	gas *bfcObjectID,
	gasBudget types.BfcBigInt,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, requestAddStake, signer, coins, amount, validator, gas, gasBudget)
}

func (c *Client) RequestWithdrawStake(
	ctx context.Context,
	signer BfcAddress,
	stakedBfcId bfcObjectID,
	gas *bfcObjectID,
	gasBudget types.BfcBigInt,
) (*types.TransactionBytes, error) {
	var resp types.TransactionBytes
	return &resp, c.CallContext(ctx, &resp, requestWithdrawStake, signer, stakedBfcId, gas, gasBudget)
}

func BCS_RequestAddStake(
	signer BfcAddress,
	coins []*bfc_types.ObjectRef,
	amount types.SafeBfcBigInt[uint64],
	validator BfcAddress,
	gasBudget, gasPrice uint64,
) ([]byte, error) {
	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	amtArg, err := ptb.Pure(amount.Uint64())
	if err != nil {
		return nil, err
	}
	arg0, err := ptb.Obj(bfc_types.SuiSystemMutObj)
	if err != nil {
		return nil, err
	}
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
	arg2, err := ptb.Pure(validator)
	if err != nil {
		return nil, err
	}

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
		signer, coins, pt, gasBudget, gasPrice,
	)
	return bcs.Marshal(tx)
}

func BCS_RequestWithdrawStake(
	signer BfcAddress,
	stakedBfcRef bfc_types.ObjectRef,
	gas []*bfc_types.ObjectRef,
	gasBudget, gasPrice uint64,
) ([]byte, error) {
	// build with BCS
	ptb := bfc_types.NewProgrammableTransactionBuilder()
	arg0, err := ptb.Obj(bfc_types.SuiSystemMutObj)
	if err != nil {
		return nil, err
	}
	arg1, err := ptb.Obj(
		bfc_types.ObjectArg{
			ImmOrOwnedObject: &stakedBfcRef,
		},
	)
	if err != nil {
		return nil, err
	}
	ptb.Command(
		bfc_types.Command{
			MoveCall: &bfc_types.ProgrammableMoveCall{
				Package:  *bfc_types.SuiSystemAddress,
				Module:   sui_system_state.SuiSystemModuleName,
				Function: bfc_types.WithdrawStakeFunName,
				Arguments: []bfc_types.Argument{
					arg0, arg1,
				},
			},
		},
	)
	pt := ptb.Finish()
	tx := bfc_types.NewProgrammable(
		signer, gas, pt, gasBudget, gasPrice,
	)
	return bcs.Marshal(tx)
}
