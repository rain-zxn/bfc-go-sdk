package types

import (
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

type ExecuteTransactionRequestType string

const (
	TxnRequestTypeWaitForEffectsCert    ExecuteTransactionRequestType = "WaitForEffectsCert"
	TxnRequestTypeWaitForLocalExecution ExecuteTransactionRequestType = "WaitForLocalExecution"
)

type EpochId = uint64

type GasCostSummary struct {
	ComputationCost         SafeBfcBigInt[uint64] `json:"computationCost"`
	StorageCost             SafeBfcBigInt[uint64] `json:"storageCost"`
	StorageRebate           SafeBfcBigInt[uint64] `json:"storageRebate"`
	NonRefundableStorageFee SafeBfcBigInt[uint64] `json:"nonRefundableStorageFee"`
}

const (
	ExecutionStatusSuccess = "success"
	ExecutionStatusFailure = "failure"
)

type ExecutionStatus struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type OwnedObjectRef struct {
	Owner     lib.TagJson[bfc_types.Owner] `json:"owner"`
	Reference BfcObjectRef                 `json:"reference"`
}

type BfcTransactionBlockEffectsModifiedAtVersions struct {
	ObjectId       bfc_types.ObjectID                      `json:"objectId"`
	SequenceNumber SafeBfcBigInt[bfc_types.SequenceNumber] `json:"sequenceNumber"`
}

type BfcTransactionBlockEffectsV1 struct {
	/** The status of the execution */
	Status ExecutionStatus `json:"status"`
	/** The epoch when this transaction was executed */
	ExecutedEpoch SafeBfcBigInt[EpochId] `json:"executedEpoch"`
	/** The version that every modified (mutated or deleted) object had before it was modified by this transaction. **/
	ModifiedAtVersions []BfcTransactionBlockEffectsModifiedAtVersions `json:"modifiedAtVersions,omitempty"`
	GasUsed            GasCostSummary                                 `json:"gasUsed"`
	/** The object references of the shared objects used in this transaction. Empty if no shared objects were used. */
	SharedObjects []BfcObjectRef `json:"sharedObjects,omitempty"`
	/** The transaction digest */
	TransactionDigest bfc_types.TransactionDigest `json:"transactionDigest"`
	/** ObjectRef and owner of new objects created */
	Created []OwnedObjectRef `json:"created,omitempty"`
	/** ObjectRef and owner of mutated objects, including gas object */
	Mutated []OwnedObjectRef `json:"mutated,omitempty"`
	/**
	 * ObjectRef and owner of objects that are unwrapped in this transaction.
	 * Unwrapped objects are objects that were wrapped into other objects in the past,
	 * and just got extracted out.
	 */
	Unwrapped []OwnedObjectRef `json:"unwrapped,omitempty"`
	/** Object Refs of objects now deleted (the old refs) */
	Deleted []BfcObjectRef `json:"deleted,omitempty"`
	/** Object Refs of objects now deleted (the old refs) */
	UnwrappedThenDeleted []BfcObjectRef `json:"unwrapped_then_deleted,omitempty"`
	/** Object refs of objects now wrapped in other objects */
	Wrapped []BfcObjectRef `json:"wrapped,omitempty"`
	/**
	 * The updated gas object reference. Have a dedicated field for convenient access.
	 * It's also included in mutated.
	 */
	GasObject OwnedObjectRef `json:"gasObject"`
	/** The events emitted during execution. Note that only successful transactions emit events */
	EventsDigest *bfc_types.TransactionEventsDigest `json:"eventsDigest,omitempty"`
	/** The set of transaction digests this transaction depends on */
	Dependencies []bfc_types.TransactionDigest `json:"dependencies,omitempty"`
}

type BfcTransactionBlockEffects struct {
	V1 *BfcTransactionBlockEffectsV1 `json:"v1"`
}

func (t BfcTransactionBlockEffects) Tag() string {
	return "messageVersion"
}

func (t BfcTransactionBlockEffects) Content() string {
	return ""
}

func (t BfcTransactionBlockEffects) GasFee() int64 {
	if t.V1 == nil {
		return 0
	}
	fee := t.V1.GasUsed.StorageCost.Int64() -
		t.V1.GasUsed.StorageRebate.Int64() +
		t.V1.GasUsed.ComputationCost.Int64()
	return fee
}

func (t BfcTransactionBlockEffects) IsSuccess() bool {
	return t.V1.Status.Status == ExecutionStatusSuccess
}

const (
	BfcTransactionBlockKindBfcChangeEpoch             = "ChangeEpoch"
	BfcTransactionBlockKindBfcConsensusCommitPrologue = "ConsensusCommitPrologue"
	BfcTransactionBlockKindGenesis                    = "Genesis"
	BfcTransactionBlockKindProgrammableTransaction    = "ProgrammableTransaction"
)

type BfcTransactionBlockKind = lib.TagJson[TransactionBlockKind]

type TransactionBlockKind struct {
	/// A system transaction that will update epoch information on-chain.
	ChangeEpoch *BfcChangeEpoch `json:"ChangeEpoch,omitempty"`
	/// A system transaction used for initializing the initial state of the chain.
	Genesis *BfcGenesisTransaction `json:"Genesis,omitempty"`
	/// A system transaction marking the start of a series of transactions scheduled as part of a
	/// checkpoint
	ConsensusCommitPrologue *BfcConsensusCommitPrologue `json:"ConsensusCommitPrologue,omitempty"`
	/// A series of transactions where the results of one transaction can be used in future
	/// transactions
	ProgrammableTransaction *BfcProgrammableTransactionBlock `json:"ProgrammableTransaction,omitempty"`
	// .. more transaction types go here
}

func (t TransactionBlockKind) Tag() string {
	return "kind"
}

func (t TransactionBlockKind) Content() string {
	return ""
}

type BfcChangeEpoch struct {
	Epoch                 SafeBfcBigInt[EpochId] `json:"epoch"`
	StorageCharge         uint64                 `json:"storage_charge"`
	ComputationCharge     uint64                 `json:"computation_charge"`
	StorageRebate         uint64                 `json:"storage_rebate"`
	EpochStartTimestampMs uint64                 `json:"epoch_start_timestamp_ms"`
}

type BfcGenesisTransaction struct {
	Objects []bfc_types.ObjectID `json:"objects"`
}

type BfcConsensusCommitPrologue struct {
	Epoch             uint64 `json:"epoch"`
	Round             uint64 `json:"round"`
	CommitTimestampMs uint64 `json:"commit_timestamp_ms"`
}

type BfcProgrammableTransactionBlock struct {
	Inputs []interface{} `json:"inputs"`
	/// The transactions to be executed sequentially. A failure in any transaction will
	/// result in the failure of the entire programmable transaction block.
	Commands []interface{} `json:"transactions"`
}

type BfcTransactionBlockDataV1 struct {
	Transaction BfcTransactionBlockKind `json:"transaction"`
	Sender      bfc_types.BfcAddress    `json:"sender"`
	GasData     BfcGasData              `json:"gasData"`
}

type BfcTransactionBlockData struct {
	V1 *BfcTransactionBlockDataV1 `json:"v1,omitempty"`
}

func (t BfcTransactionBlockData) Tag() string {
	return "messageVersion"
}

func (t BfcTransactionBlockData) Content() string {
	return ""
}

type BfcTransactionBlock struct {
	Data         lib.TagJson[BfcTransactionBlockData] `json:"data"`
	TxSignatures []string                             `json:"txSignatures"`
}

type ObjectChange struct {
	Published *struct {
		PackageId bfc_types.ObjectID                      `json:"packageId"`
		Version   SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
		Digest    bfc_types.ObjectDigest                  `json:"digest"`
		Nodules   []string                                `json:"nodules"`
	} `json:"published,omitempty"`
	/// Transfer objects to new address / wrap in another object
	Transferred *struct {
		Sender     bfc_types.BfcAddress                    `json:"sender"`
		Recipient  ObjectOwner                             `json:"recipient"`
		ObjectType string                                  `json:"objectType"`
		ObjectId   bfc_types.ObjectID                      `json:"objectId"`
		Version    SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
		Digest     bfc_types.ObjectDigest                  `json:"digest"`
	} `json:"transferred,omitempty"`
	/// Object mutated.
	Mutated *struct {
		Sender          bfc_types.BfcAddress                    `json:"sender"`
		Owner           ObjectOwner                             `json:"owner"`
		ObjectType      string                                  `json:"objectType"`
		ObjectId        bfc_types.ObjectID                      `json:"objectId"`
		Version         SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
		PreviousVersion SafeBfcBigInt[bfc_types.SequenceNumber] `json:"previousVersion"`
		Digest          bfc_types.ObjectDigest                  `json:"digest"`
	} `json:"mutated,omitempty"`
	/// Delete object j
	Deleted *struct {
		Sender     bfc_types.BfcAddress                    `json:"sender"`
		ObjectType string                                  `json:"objectType"`
		ObjectId   bfc_types.ObjectID                      `json:"objectId"`
		Version    SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
	} `json:"deleted,omitempty"`
	/// Wrapped object
	Wrapped *struct {
		Sender     bfc_types.BfcAddress                    `json:"sender"`
		ObjectType string                                  `json:"objectType"`
		ObjectId   bfc_types.ObjectID                      `json:"objectId"`
		Version    SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
	} `json:"wrapped,omitempty"`
	/// New object creation
	Created *struct {
		Sender     bfc_types.BfcAddress                    `json:"sender"`
		Owner      ObjectOwner                             `json:"owner"`
		ObjectType string                                  `json:"objectType"`
		ObjectId   bfc_types.ObjectID                      `json:"objectId"`
		Version    SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
		Digest     bfc_types.ObjectDigest                  `json:"digest"`
	} `json:"created,omitempty"`
}

func (o ObjectChange) Tag() string {
	return "type"
}

func (o ObjectChange) Content() string {
	return ""
}

type BalanceChange struct {
	Owner    ObjectOwner `json:"owner"`
	CoinType string      `json:"coinType"`
	/* Coin balance change(positive means receive, negative means send) */
	Amount string `json:"amount"`
}

type BfcTransactionBlockResponse struct {
	Digest                  bfc_types.TransactionDigest              `json:"digest"`
	Transaction             *BfcTransactionBlock                     `json:"transaction,omitempty"`
	RawTransaction          []byte                                   `json:"rawTransaction,omitempty"`
	Effects                 *lib.TagJson[BfcTransactionBlockEffects] `json:"effects,omitempty"`
	Events                  []BfcEvent                               `json:"events,omitempty"`
	TimestampMs             *SafeBfcBigInt[uint64]                   `json:"timestampMs,omitempty"`
	Checkpoint              *SafeBfcBigInt[CheckpointSequenceNumber] `json:"checkpoint,omitempty"`
	ConfirmedLocalExecution *bool                                    `json:"confirmedLocalExecution,omitempty"`
	ObjectChanges           []lib.TagJson[ObjectChange]              `json:"objectChanges,omitempty"`
	BalanceChanges          []BalanceChange                          `json:"balanceChanges,omitempty"`
	/* Errors that occurred in fetching/serializing the transaction. */
	Errors []string `json:"errors,omitempty"`
}

type ReturnValueType interface{}
type MutableReferenceOutputType interface{}
type ExecutionResultType struct {
	MutableReferenceOutputs []MutableReferenceOutputType `json:"mutableReferenceOutputs,omitempty"`
	ReturnValues            []ReturnValueType            `json:"returnValues,omitempty"`
}

type DevInspectResults struct {
	Effects lib.TagJson[BfcTransactionBlockEffects] `json:"effects"`
	Events  []BfcEvent                              `json:"events"`
	Results []ExecutionResultType                   `json:"results,omitempty"`
	Error   *string                                 `json:"error,omitempty"`
}

type TransactionFilter struct {
	Checkpoint   *bfc_types.SequenceNumber `json:"Checkpoint,omitempty"`
	MoveFunction *struct {
		Package  bfc_types.ObjectID `json:"package"`
		Module   string             `json:"module,omitempty"`
		Function string             `json:"function,omitempty"`
	} `json:"MoveFunction,omitempty"`
	InputObject      *bfc_types.ObjectID   `json:"InputObject,omitempty"`
	ChangedObject    *bfc_types.ObjectID   `json:"ChangedObject,omitempty"`
	FromAddress      *bfc_types.BfcAddress `json:"FromAddress,omitempty"`
	ToAddress        *bfc_types.BfcAddress `json:"ToAddress,omitempty"`
	FromAndToAddress *struct {
		From *bfc_types.BfcAddress `json:"from"`
		To   *bfc_types.BfcAddress `json:"to"`
	} `json:"FromAndToAddress,omitempty"`
	TransactionKind *string `json:"TransactionKind,omitempty"`
}

type BfcTransactionBlockResponseOptions struct {
	/* Whether to show transaction input data. Default to be false. */
	ShowInput bool `json:"showInput,omitempty"`
	/* Whether to show transaction effects. Default to be false. */
	ShowEffects bool `json:"showEffects,omitempty"`
	/* Whether to show transaction events. Default to be false. */
	ShowEvents bool `json:"showEvents,omitempty"`
	/* Whether to show object changes. Default to be false. */
	ShowObjectChanges bool `json:"showObjectChanges,omitempty"`
	/* Whether to show coin balance changes. Default to be false. */
	ShowBalanceChanges bool `json:"showBalanceChanges,omitempty"`
	ShowRawInput       bool `json:"showRawInput,omitempty"`
}

type BfcTransactionBlockResponseQuery struct {
	Filter  *TransactionFilter                  `json:"filter,omitempty"`
	Options *BfcTransactionBlockResponseOptions `json:"options,omitempty"`
}

type TransactionBlocksPage = Page[BfcTransactionBlockResponse, bfc_types.TransactionDigest]

type DryRunTransactionBlockResponse struct {
	Effects        lib.TagJson[BfcTransactionBlockEffects] `json:"effects"`
	Events         []BfcEvent                              `json:"events"`
	ObjectChanges  []lib.TagJson[ObjectChange]             `json:"objectChanges"`
	BalanceChanges []BalanceChange                         `json:"balanceChanges"`
	Input          lib.TagJson[BfcTransactionBlockData]    `json:"input"`
}

type TxSign struct {
	Hash []string `json:"hash"`
}
