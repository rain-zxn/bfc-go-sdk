package bfc_types

import (
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_protocol"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"github.com/hellokittyboy-code/benfen-go-sdk/move_types"
)

type TransactionData struct {
	V1 *TransactionDataV1
}

func (t TransactionData) IsBcsEnum() {

}

type TransactionDataV1 struct {
	Kind       TransactionKind
	Sender     BfcAddress
	GasData    GasData
	Expiration TransactionExpiration
}

type TransactionExpiration struct {
	None  *lib.EmptyEnum
	Epoch *EpochId
}

func (t TransactionExpiration) IsBcsEnum() {
}

type GasData struct {
	Payment []*ObjectRef
	Owner   BfcAddress
	Price   uint64
	Budget  uint64
}

type TransactionKind struct {
	ProgrammableTransaction *ProgrammableTransaction
	ChangeEpoch             *ChangeEpoch
	Genesis                 *GenesisTransaction
	ConsensusCommitPrologue *ConsensusCommitPrologue
}

func (t TransactionKind) IsBcsEnum() {
}

type ConsensusCommitPrologue struct {
	Epoch             uint64
	Round             uint64
	CommitTimestampMs CheckpointTimestamp
}

type ProgrammableTransaction struct {
	Inputs   []CallArg
	Commands []Command
}

type Command struct {
	MoveCall        *ProgrammableMoveCall
	TransferObjects *struct {
		Arguments []Argument
		Argument  Argument
	}
	SplitCoins *struct {
		Argument  Argument
		Arguments []Argument
	}
	MergeCoins *struct {
		Argument  Argument
		Arguments []Argument
	}
	Publish *struct {
		Bytes   [][]uint8
		Objects []ObjectID
	}
	MakeMoveVec *struct {
		TypeTag   *move_types.TypeTag `bcs:"optional"`
		Arguments []Argument
	}
	Upgrade *struct {
		Bytes    [][]uint8
		Objects  []ObjectID
		ObjectID ObjectID
		Argument Argument
	}
}

func (c Command) IsBcsEnum() {

}

type Argument struct {
	GasCoin      *lib.EmptyEnum
	Input        *uint16
	Result       *uint16
	NestedResult *struct {
		Result1 uint16
		Result2 uint16
	}
}

func (a Argument) IsBcsEnum() {

}

type ProgrammableMoveCall struct {
	Package       ObjectID
	Module        move_types.Identifier
	Function      move_types.Identifier
	TypeArguments []move_types.TypeTag
	Arguments     []Argument
}

type SingleTransactionKind struct {
	TransferObject *TransferObject
	Publish        *MoveModulePublish
	Call           *MoveCall
	TransferBfc    *TransferBfc
	Pay            *Pay
	PayBfc         *PayBfc
	PayAllBfc      *PayAllBfc
	ChangeEpoch    *ChangeEpoch
	Genesis        *GenesisTransaction
}

func (s SingleTransactionKind) IsBcsEnum() {
}

type TransferObject struct {
	Recipient BfcAddress
	ObjectRef ObjectRef
}

type MoveModulePublish struct {
	Modules [][]byte
}

type MoveCall struct {
	Package       ObjectID
	Module        string
	Function      string
	TypeArguments []*move_types.TypeTag
	Arguments     []*CallArg
}

type TransferBfc struct {
	Recipient BfcAddress
	Amount    *uint64 `bcs:"optional"`
}

type Pay struct {
	Coins      []*ObjectRef
	Recipients []*BfcAddress
	Amounts    []*uint64
}

type PayBfc = Pay

type PayAllBfc struct {
	Coins     []*ObjectRef
	Recipient BfcAddress
}

type ChangeEpoch struct {
	Epoch                   EpochId
	ProtocolVersion         bfc_protocol.ProtocolVersion
	StorageCharge           uint64
	ComputationCharge       uint64
	StorageRebate           uint64
	NonRefundableStorageFee uint64
	EpochStartTimestampMs   uint64
	SystemPackages          []*struct {
		SequenceNumber SequenceNumber
		Bytes          [][]uint8
		Objects        []*ObjectID
	}
}

type GenesisTransaction struct {
	Objects []GenesisObject
}

type GenesisObject struct {
	RawObject *struct {
		Data  Data
		Owner Owner
	}
}

type CallArg struct {
	Pure   *[]byte
	Object *ObjectArg
}

func (c CallArg) IsBcsEnum() {
}

type SharedObject struct {
	Id                   ObjectID
	InitialSharedVersion SequenceNumber
	Mutable              bool
}

type ObjectArg struct {
	ImmOrOwnedObject *ObjectRef
	SharedObject     *SharedObject
}

func (o ObjectArg) IsBcsEnum() {
}

func (o ObjectArg) id() ObjectID {
	switch {
	case o.ImmOrOwnedObject != nil:
		return o.ImmOrOwnedObject.ObjectId
	case o.SharedObject != nil:
		return o.SharedObject.Id
	default:
		return ObjectID{}
	}
}
