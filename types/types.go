package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

const (
	BFCoinType     = "0x2::bfc::BFC"
	LocalNetRpcUrl = "http://0.0.0.0:9000"

	DevNetRpcUrl  = "https://obcrpc.openblock.vip"
	TestnetRpcUrl = "https://testrpc.benfen.org/"
	MainnetRpcUrl = "https://obcrpc.openblock.vip/"
	DevIndexRpc   = "https://obcindex.openblock.vip"
)

// ShortString Returns the address with leading zeros trimmed, e.g. 0x2

type InputObjectKind map[string]interface{}

type TransactionBytes struct {
	// the gas object to be used
	Gas []bfc_types.ObjectRef `json:"gas"`

	// objects to be used in this transaction
	InputObjects []InputObjectKind `json:"inputObjects"`

	// transaction data bytes
	TxBytes lib.Base64Data `json:"txBytes"`
}

type TransferObject struct {
	Recipient bfc_types.BfcAddress `json:"recipient"`
	ObjectRef bfc_types.ObjectRef  `json:"object_ref"`
}
type ModulePublish struct {
	Modules [][]byte `json:"modules"`
}
type MoveCall struct {
	Package  bfc_types.ObjectID `json:"package"`
	Module   string             `json:"module"`
	Function string             `json:"function"`
	TypeArgs []interface{}      `json:"typeArguments"`
	Args     []interface{}      `json:"arguments"`
}
type TransferBfc struct {
	Recipient bfc_types.BfcAddress `json:"recipient"`
	Amount    uint64               `json:"amount"`
}
type Pay struct {
	Coins      []bfc_types.ObjectRef  `json:"coins"`
	Recipients []bfc_types.BfcAddress `json:"recipients"`
	Amounts    []uint64               `json:"amounts"`
}
type PayBfc struct {
	Coins      []bfc_types.ObjectRef  `json:"coins"`
	Recipients []bfc_types.BfcAddress `json:"recipients"`
	Amounts    []uint64               `json:"amounts"`
}
type PayAllBfc struct {
	Coins     []bfc_types.ObjectRef `json:"coins"`
	Recipient bfc_types.BfcAddress  `json:"recipient"`
}
type ChangeEpoch struct {
	Epoch             interface{} `json:"epoch"`
	StorageCharge     uint64      `json:"storage_charge"`
	ComputationCharge uint64      `json:"computation_charge"`
}

type SingleTransactionKind struct {
	TransferObject *TransferObject `json:"TransferObject,omitempty"`
	Publish        *ModulePublish  `json:"Publish,omitempty"`
	Call           *MoveCall       `json:"Call,omitempty"`
	TransferBfc    *TransferBfc    `json:"TransferBfc,omitempty"`
	ChangeEpoch    *ChangeEpoch    `json:"ChangeEpoch,omitempty"`
	PayBfc         *PayBfc         `json:"PaySui,omitempty"`
	Pay            *Pay            `json:"Pay,omitempty"`
	PayAllBfc      *PayAllBfc      `json:"PayAllSui,omitempty"`
}

type SenderSignedData struct {
	Transactions []SingleTransactionKind `json:"transactions,omitempty"`

	Sender     *bfc_types.BfcAddress `json:"sender"`
	GasPayment *bfc_types.ObjectRef  `json:"gasPayment"`
	GasBudget  uint64                `json:"gasBudget"`
	// GasPrice     uint64      `json:"gasPrice"`
}

type TimeRange struct {
	StartTime uint64 `json:"startTime"` // left endpoint of time interval, milliseconds since epoch, inclusive
	EndTime   uint64 `json:"endTime"`   // right endpoint of time interval, milliseconds since epoch, exclusive
}

type MoveModule struct {
	Package bfc_types.ObjectID `json:"package"`
	Module  string             `json:"module"`
}

func (o ObjectOwner) MarshalJSON() ([]byte, error) {
	if o.string != nil {
		data, err := json.Marshal(o.string)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if o.ObjectOwnerInternal != nil {
		data, err := json.Marshal(o.ObjectOwnerInternal)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("nil value")
}

func (o *ObjectOwner) UnmarshalJSON(data []byte) error {
	if bytes.HasPrefix(data, []byte("\"")) {
		stringData := string(data[1 : len(data)-1])
		o.string = &stringData
		return nil
	}
	if bytes.HasPrefix(data, []byte("{")) {
		oOI := ObjectOwnerInternal{}
		err := json.Unmarshal(data, &oOI)
		if err != nil {
			return err
		}
		o.ObjectOwnerInternal = &oOI
		return nil
	}
	return errors.New("value not json")
}

func IsSameStringAddress(addr1, addr2 string) bool {
	addr1 = strings.TrimPrefix(addr1, "0x")
	addr2 = strings.TrimPrefix(addr2, "0x")
	addr1 = strings.TrimLeft(addr1, "0")
	return strings.TrimLeft(addr1, "0") == strings.TrimLeft(addr2, "0")
}
