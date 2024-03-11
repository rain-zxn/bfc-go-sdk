package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

type StakeStatus = lib.TagJson[Status]

type Status struct {
	Pending *struct{} `json:"Pending,omitempty"`
	Active  *struct {
		EstimatedReward SafeBfcBigInt[uint64] `json:"estimatedReward"`
	} `json:"Active,omitempty"`
	Unstaked *struct{} `json:"Unstaked,omitempty"`
}

func (s Status) Tag() string {
	return "status"
}

func (s Status) Content() string {
	return ""
}

const (
	StakeStatusActive   = "Active"
	StakeStatusPending  = "Pending"
	StakeStatusUnstaked = "Unstaked"
)

type EpochPage struct {
	Data []*EpochInfo `json:"data"`
}

type EpochInfo struct {
	Epoch             string         `json:"epoch"`
	FirstCheckpointId string         `json:"firstCheckpointId"`
	EndOfEpochInfo    EndOfEpochInfo `json:"endOfEpochInfo"`
}

type EndOfEpochInfo struct {
	LastCheckpointId string `json:"lastCheckpointId"`
	ProtocolVersion  string `json:"protocolVersion"`
}

type Stake struct {
	StakedBfcId       bfc_types.ObjectID     `json:"stakedSuiId"`
	StakeRequestEpoch SafeBfcBigInt[EpochId] `json:"stakeRequestEpoch"`
	StakeActiveEpoch  SafeBfcBigInt[EpochId] `json:"stakeActiveEpoch"`
	Principal         SafeBfcBigInt[uint64]  `json:"principal"`
	StakeStatus       *StakeStatus           `json:"-,flatten"`
}

func (s *Stake) IsActive() bool {
	return s.StakeStatus.Data.Active != nil
}

type JsonFlatten[T Stake] struct {
	Data T
}

func (s *JsonFlatten[T]) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &s.Data)
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(s).Elem().Field(0)
	for i := 0; i < rv.Type().NumField(); i++ {
		tag := rv.Type().Field(i).Tag.Get("json")
		if strings.Contains(tag, "flatten") {
			if rv.Field(i).Kind() != reflect.Pointer {
				return fmt.Errorf("field %s not pointer", rv.Field(i).Type().Name())
			}
			if rv.Field(i).IsNil() {
				rv.Field(i).Set(reflect.New(rv.Field(i).Type().Elem()))
			}
			err = json.Unmarshal(data, rv.Field(i).Interface())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DelegatedStake struct {
	ValidatorAddress bfc_types.BfcAddress `json:"validatorAddress"`
	StakingPool      bfc_types.ObjectID   `json:"stakingPool"`
	Stakes           []JsonFlatten[Stake] `json:"stakes"`
}

type BfcValidatorSummary struct {
	BfcAddress             bfc_types.BfcAddress `json:"BfcAddress"`
	ProtocolPubkeyBytes    lib.Base64Data       `json:"protocolPubkeyBytes"`
	NetworkPubkeyBytes     lib.Base64Data       `json:"networkPubkeyBytes"`
	WorkerPubkeyBytes      lib.Base64Data       `json:"workerPubkeyBytes"`
	ProofOfPossessionBytes lib.Base64Data       `json:"proofOfPossessionBytes"`
	OperationCapId         bfc_types.ObjectID   `json:"operationCapId"`
	Name                   string               `json:"name"`
	Description            string               `json:"description"`
	ImageUrl               string               `json:"imageUrl"`
	ProjectUrl             string               `json:"projectUrl"`
	P2pAddress             string               `json:"p2pAddress"`
	NetAddress             string               `json:"netAddress"`
	PrimaryAddress         string               `json:"primaryAddress"`
	WorkerAddress          string               `json:"workerAddress"`

	NextEpochProtocolPubkeyBytes lib.Base64Data `json:"nextEpochProtocolPubkeyBytes"`
	NextEpochProofOfPossession   lib.Base64Data `json:"nextEpochProofOfPossession"`
	NextEpochNetworkPubkeyBytes  lib.Base64Data `json:"nextEpochNetworkPubkeyBytes"`
	NextEpochWorkerPubkeyBytes   lib.Base64Data `json:"nextEpochWorkerPubkeyBytes"`
	NextEpochNetAddress          string         `json:"nextEpochNetAddress"`
	NextEpochP2pAddress          string         `json:"nextEpochP2pAddress"`
	NextEpochPrimaryAddress      string         `json:"nextEpochPrimaryAddress"`
	NextEpochWorkerAddress       string         `json:"nextEpochWorkerAddress"`

	VotingPower             SafeBfcBigInt[uint64] `json:"votingPower"`
	GasPrice                SafeBfcBigInt[uint64] `json:"gasPrice"`
	CommissionRate          SafeBfcBigInt[uint64] `json:"commissionRate"`
	NextEpochStake          SafeBfcBigInt[uint64] `json:"nextEpochStake"`
	NextEpochGasPrice       SafeBfcBigInt[uint64] `json:"nextEpochGasPrice"`
	NextEpochCommissionRate SafeBfcBigInt[uint64] `json:"nextEpochCommissionRate"`
	StakingPoolId           bfc_types.ObjectID    `json:"stakingPoolId"`

	StakingPoolActivationEpoch   SafeBfcBigInt[uint64] `json:"stakingPoolActivationEpoch"`
	StakingPoolDeactivationEpoch SafeBfcBigInt[uint64] `json:"stakingPoolDeactivationEpoch"`

	StakingPoolBfcBalance    SafeBfcBigInt[uint64] `json:"stakingPoolSuiBalance"`
	RewardsPool              SafeBfcBigInt[uint64] `json:"rewardsPool"`
	PoolTokenBalance         SafeBfcBigInt[uint64] `json:"poolTokenBalance"`
	PendingStake             SafeBfcBigInt[uint64] `json:"pendingStake"`
	PendingPoolTokenWithdraw SafeBfcBigInt[uint64] `json:"pendingPoolTokenWithdraw"`
	PendingTotalBfcWithdraw  SafeBfcBigInt[uint64] `json:"pendingTotalSuiWithdraw"`
	ExchangeRatesId          bfc_types.ObjectID    `json:"exchangeRatesId"`
	ExchangeRatesSize        SafeBfcBigInt[uint64] `json:"exchangeRatesSize"`
}

type TypeName []bfc_types.BfcAddress
type BfcSystemStateSummary struct {
	Epoch                                 SafeBfcBigInt[uint64]   `json:"epoch"`
	ProtocolVersion                       SafeBfcBigInt[uint64]   `json:"protocolVersion"`
	SystemStateVersion                    SafeBfcBigInt[uint64]   `json:"systemStateVersion"`
	StorageFundTotalObjectStorageRebates  SafeBfcBigInt[uint64]   `json:"storageFundTotalObjectStorageRebates"`
	StorageFundNonRefundableBalance       SafeBfcBigInt[uint64]   `json:"storageFundNonRefundableBalance"`
	ReferenceGasPrice                     SafeBfcBigInt[uint64]   `json:"referenceGasPrice"`
	SafeMode                              bool                    `json:"safeMode"`
	SafeModeStorageRewards                SafeBfcBigInt[uint64]   `json:"safeModeStorageRewards"`
	SafeModeComputationRewards            SafeBfcBigInt[uint64]   `json:"safeModeComputationRewards"`
	SafeModeStorageRebates                SafeBfcBigInt[uint64]   `json:"safeModeStorageRebates"`
	SafeModeNonRefundableStorageFee       SafeBfcBigInt[uint64]   `json:"safeModeNonRefundableStorageFee"`
	EpochStartTimestampMs                 SafeBfcBigInt[uint64]   `json:"epochStartTimestampMs"`
	EpochDurationMs                       SafeBfcBigInt[uint64]   `json:"epochDurationMs"`
	StakeSubsidyStartEpoch                SafeBfcBigInt[uint64]   `json:"stakeSubsidyStartEpoch"`
	MaxValidatorCount                     SafeBfcBigInt[uint64]   `json:"maxValidatorCount"`
	MinValidatorJoiningStake              SafeBfcBigInt[uint64]   `json:"minValidatorJoiningStake"`
	ValidatorLowStakeThreshold            SafeBfcBigInt[uint64]   `json:"validatorLowStakeThreshold"`
	ValidatorVeryLowStakeThreshold        SafeBfcBigInt[uint64]   `json:"validatorVeryLowStakeThreshold"`
	ValidatorLowStakeGracePeriod          SafeBfcBigInt[uint64]   `json:"validatorLowStakeGracePeriod"`
	StakeSubsidyBalance                   SafeBfcBigInt[uint64]   `json:"stakeSubsidyBalance"`
	StakeSubsidyDistributionCounter       SafeBfcBigInt[uint64]   `json:"stakeSubsidyDistributionCounter"`
	StakeSubsidyCurrentDistributionAmount SafeBfcBigInt[uint64]   `json:"stakeSubsidyCurrentDistributionAmount"`
	StakeSubsidyPeriodLength              SafeBfcBigInt[uint64]   `json:"stakeSubsidyPeriodLength"`
	StakeSubsidyDecreaseRate              uint16                  `json:"stakeSubsidyDecreaseRate"`
	TotalStake                            SafeBfcBigInt[uint64]   `json:"totalStake"`
	ActiveValidators                      []BfcValidatorSummary   `json:"activeValidators"`
	PendingActiveValidatorsId             bfc_types.ObjectID      `json:"pendingActiveValidatorsId"`
	PendingActiveValidatorsSize           SafeBfcBigInt[uint64]   `json:"pendingActiveValidatorsSize"`
	PendingRemovals                       []SafeBfcBigInt[uint64] `json:"pendingRemovals"`
	StakingPoolMappingsId                 bfc_types.ObjectID      `json:"stakingPoolMappingsId"`
	StakingPoolMappingsSize               SafeBfcBigInt[uint64]   `json:"stakingPoolMappingsSize"`
	InactivePoolsId                       bfc_types.ObjectID      `json:"inactivePoolsId"`
	InactivePoolsSize                     SafeBfcBigInt[uint64]   `json:"inactivePoolsSize"`
	ValidatorCandidatesId                 bfc_types.ObjectID      `json:"validatorCandidatesId"`
	ValidatorCandidatesSize               SafeBfcBigInt[uint64]   `json:"validatorCandidatesSize"`
	AtRiskValidators                      interface{}             `json:"atRiskValidators"`
	ValidatorReportRecords                interface{}             `json:"validatorReportRecords"`
}

type ValidatorsApy struct {
	Epoch SafeBfcBigInt[EpochId] `json:"epoch"`
	Apys  []struct {
		Address string  `json:"address"`
		Apy     float64 `json:"apy"`
	} `json:"apys"`
}

type ProtocolConfig struct {
}

type CommonStruct struct {
}

type CheckPointsPage struct {
	points []Checkpoints `json:"data"`
}
type Checkpoints struct {
	Epoch SafeBfcBigInt[EpochId] `json:"epoch"`
}

func (apys *ValidatorsApy) ApyMap() map[string]float64 {
	res := make(map[string]float64)
	for _, apy := range apys.Apys {
		res[apy.Address] = apy.Apy
	}
	return res
}
