package client

import (
	"context"
	"math/big"
	"testing"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	ValidatorAddress = "0xd961ce7fdc406116d6dca6f1c17436d1196a97ae877e5599f12b9f882123fe4f"
)

func TestClient_GetLatestSuiSystemState(t *testing.T) {
	cli := ChainClient(t)
	state, err := cli.GetLatestSuiSystemState(context.Background())
	require.Nil(t, err)
	t.Logf("system state = %v", state)
}

func TestClient_GetValidatorsApy(t *testing.T) {
	cli := ChainClient(t)
	apys, err := cli.GetValidatorsApy(context.Background())
	require.Nil(t, err)
	t.Logf("current epoch %v", apys.Epoch)
	apyMap := apys.ApyMap()
	defaultValidatorCount := 4
	for idx := 0; idx < defaultValidatorCount; idx++ {
		key := apys.Apys[idx].Address
		t.Logf("%v apy = %v", key, apyMap[key])
	}
}

func TestGetDelegatedStakes(t *testing.T) {
	cli := ChainClient(t)

	address, err := bfc_types.NewAddressFromHex(Address.String())
	require.Nil(t, err)
	stakes, err := cli.GetStakes(context.Background(), *address)
	require.Nil(t, err)

	for _, validator := range stakes {
		for _, stake := range validator.Stakes {
			if stake.Data.StakeStatus.Data.Active != nil {
				t.Logf(
					"earned amount %10v at %v",
					stake.Data.StakeStatus.Data.Active.EstimatedReward.Uint64(),
					validator.ValidatorAddress,
				)
			}
		}
	}
}

func TestGetStakesByIds(t *testing.T) {
	cli := ChainClient(t)
	owner, err := bfc_types.NewAddressFromHex(Address.String())
	stakes, err := cli.GetStakes(context.Background(), *owner)
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(stakes), 1)

	stake1 := stakes[0].Stakes[0].Data
	stakeId := stake1.StakedBfcId
	stakesFromId, err := cli.GetStakesByIds(context.Background(), []bfcObjectID{stakeId})
	require.Nil(t, err)
	require.GreaterOrEqual(t, len(stakesFromId), 1)

	queriedStake := stakesFromId[0].Stakes[0].Data
	require.Equal(t, stake1, queriedStake)
	t.Log(stakesFromId)
}

// important , need check if the validator address is right...
func TestRequestAddDelegation(t *testing.T) {
	cli := ChainClient(t)
	signer := Address

	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.Nil(t, err)

	amount := BFC(10).Uint64()
	gasBudget := BFC(0.1).Uint64()
	gasPrice := 100
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0).SetUint64(amount), 0, 0, 0)
	require.Nil(t, err)

	validatorAddress := ValidatorAddress
	validator, err := bfc_types.NewAddressFromHex(validatorAddress)
	require.Nil(t, err)

	txBytes, err := BCS_RequestAddStake(
		*signer,
		pickedCoins.CoinRefs(),
		types.NewSafeBfcBigInt(amount),
		*validator,
		gasBudget,
		uint64(gasPrice),
	)
	require.Nil(t, err)

	simulateCheck(t, cli, txBytes, true)
}

// single run can pass...
func TestRequestWithdrawDelegation(t *testing.T) {
	cli := ChainClient(t)
	gasBudget := BFC(0.1).Uint64()
	gasPrice := 100

	signer, err := bfc_types.NewAddressFromHex(Address.String())
	require.Nil(t, err)
	stakes, err := cli.GetStakes(context.Background(), *signer)
	require.Nil(t, err)
	require.True(t, len(stakes) > 0)
	require.True(t, len(stakes[0].Stakes) > 0)

	coins, err := cli.GetCoins(context.Background(), *signer, nil, nil, 10)
	require.Nil(t, err)
	pickedCoins, err := types.PickupCoins(coins, *big.NewInt(0), gasBudget, 0, 0)
	require.Nil(t, err)

	stakeId := stakes[0].Stakes[0].Data.StakedBfcId
	detail, err := cli.GetObject(context.Background(), stakeId, nil)
	require.Nil(t, err)
	txBytes, err := BCS_RequestWithdrawStake(
		*signer, detail.Data.Reference(), pickedCoins.CoinRefs(), gasBudget,
		uint64(gasPrice),
	)
	require.Nil(t, err)
	println("start simulate check .....")

	simulateCheck(t, cli, txBytes, true)
}
