package types

import (
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

type AuthSignInfo interface{}

type CertifiedTransaction struct {
	TransactionDigest string        `json:"transactionDigest"`
	TxSignature       string        `json:"txSignature"`
	AuthSignInfo      *AuthSignInfo `json:"authSignInfo"`

	Data *SenderSignedData `json:"data"`
}

type ParsedTransactionResponse interface{}

type ExecuteTransactionEffects struct {
	TransactionEffectsDigest string `json:"transactionEffectsDigest"`

	Effects      lib.TagJson[BfcTransactionBlockEffects] `json:"effects"`
	AuthSignInfo *AuthSignInfo                           `json:"authSignInfo"`
}

type ExecuteTransactionResponse struct {
	Certificate CertifiedTransaction      `json:"certificate"`
	Effects     ExecuteTransactionEffects `json:"effects"`

	ConfirmedLocalExecution bool `json:"confirmed_local_execution"`
}

func (r *ExecuteTransactionResponse) TransactionDigest() string {
	return r.Certificate.TransactionDigest
}

type BfcCoinMetadata struct {
	Decimals    uint8              `json:"decimals"`
	Description string             `json:"description"`
	IconUrl     string             `json:"iconUrl,omitempty"`
	Id          bfc_types.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Symbol      string             `json:"symbol"`
}

type DevInspectResult struct {
	Err string `json:"Err,omitempty"`
	Ok  any    `json:"Ok,omitempty"` //Result_of_Array_of_Tuple_of_uint_and_BfcExecutionResult_or_String
}

type GetPastObjectRequest struct {
	ObjectId bfc_types.ObjectID `json:"objectId"`
	Version  string             `json:"version"`
}

type TransferObjectParams struct {
	ObjectId  bfc_types.ObjectID   `json:"objectId"`
	Recipient bfc_types.BfcAddress `json:"recipient"`
}

type CheckPointObject struct {
	Epoch                      string   `json:"epoch"`
	SequenceNumber             string   `json:"sequenceNumber"`
	Digest                     string   `json:"digest"`
	NetworkTotalTransactions   string   `json:"networkTotalTransactions"`
	PreviousDigest             string   `json:"previousDigest"`
	TimestampMs                string   `json:"timestampMs"`
	Transactions               []string `json:"transactions"`
	CheckpointCommitments      []string `json:"checkpointCommitments"`
	ValidatorSignature         string   `json:"validatorSignature"`
	EpochRollingGasCostSummary struct {
		ComputationCost         string `json:"computationCost"`
		StorageCost             string `json:"storageCost"`
		StorageRebate           string `json:"storageRebate"`
		NonRefundableStorageFee string `json:"nonRefundableStorageFee"`
	} `json:"epochRollingGasCostSummary"`
}
