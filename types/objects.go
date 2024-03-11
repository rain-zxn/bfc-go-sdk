package types

import (
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

type BfcObjectRef struct {
	/** Base64 string representing the object digest */
	Digest bfc_types.TransactionDigest `json:"digest"`
	/** Hex code as string representing the object id */
	ObjectId string `json:"objectId"`
	/** Object version */
	Version bfc_types.SequenceNumber `json:"version"`
}

type BfcGasData struct {
	Payment []BfcObjectRef `json:"payment"`
	/** Gas Object's owner */
	Owner  string                `json:"owner"`
	Price  SafeBfcBigInt[uint64] `json:"price"`
	Budget SafeBfcBigInt[uint64] `json:"budget"`
}

type BfcParsedData struct {
	MoveObject *BfcParsedMoveObject `json:"moveObject,omitempty"`
	Package    *BfcMovePackage      `json:"package,omitempty"`
}

func (p BfcParsedData) Tag() string {
	return "dataType"
}

func (p BfcParsedData) Content() string {
	return ""
}

type BfcMovePackage struct {
	Disassembled map[string]interface{} `json:"disassembled"`
}

type BfcParsedMoveObject struct {
	Type              string `json:"type"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Fields            any    `json:"fields"`
}

type BfcRawData struct {
	MoveObject *BfcRawMoveObject  `json:"moveObject,omitempty"`
	Package    *BfcRawMovePackage `json:"package,omitempty"`
}

func (r BfcRawData) Tag() string {
	return "dataType"
}

func (r BfcRawData) Content() string {
	return ""
}

type BfcRawMoveObject struct {
	Type              string                   `json:"type"`
	HasPublicTransfer bool                     `json:"hasPublicTransfer"`
	Version           bfc_types.SequenceNumber `json:"version"`
	BcsBytes          lib.Base64Data           `json:"bcsBytes"`
}

type BfcRawMovePackage struct {
	Id              bfc_types.ObjectID        `json:"id"`
	Version         bfc_types.SequenceNumber  `json:"version"`
	ModuleMap       map[string]lib.Base64Data `json:"moduleMap"`
	TypeOriginTable []TypeOrigin              `json:"typeOriginTable"`
	LinkageTable    map[string]UpgradeInfo
}

type UpgradeInfo struct {
	UpgradedId      bfc_types.ObjectID
	UpgradedVersion bfc_types.SequenceNumber
}

type TypeOrigin struct {
	ModuleName string             `json:"moduleName"`
	StructName string             `json:"structName"`
	Package    bfc_types.ObjectID `json:"package"`
}

type BfcObjectData struct {
	ObjectId bfc_types.ObjectID                      `json:"objectId"`
	Version  SafeBfcBigInt[bfc_types.SequenceNumber] `json:"version"`
	Digest   bfc_types.ObjectDigest                  `json:"digest"`
	/**
	 * Type of the object, default to be undefined unless BfcObjectDataOptions.showType is set to true
	 */
	Type *string `json:"type,omitempty"`
	/**
	 * Move object content or package content, default to be undefined unless BfcObjectDataOptions.showContent is set to true
	 */
	Content *lib.TagJson[BfcParsedData] `json:"content,omitempty"`
	/**
	 * Move object content or package content in BCS bytes, default to be undefined unless BfcObjectDataOptions.showBcs is set to true
	 */
	Bcs *lib.TagJson[BfcRawData] `json:"bcs,omitempty"`
	/**
	 * The owner of this object. Default to be undefined unless BfcObjectDataOptions.showOwner is set to true
	 */
	Owner *ObjectOwner `json:"owner,omitempty"`
	/**
	 * The digest of the transaction that created or last mutated this object.
	 * Default to be undefined unless BfcObjectDataOptions.showPreviousTransaction is set to true
	 */
	PreviousTransaction *bfc_types.TransactionDigest `json:"previousTransaction,omitempty"`
	/**
	 * The amount of BFC we would rebate if this object gets deleted.
	 * This number is re-calculated each time the object is mutated based on
	 * the present storage gas price.
	 * Default to be undefined unless BfcObjectDataOptions.showStorageRebate is set to true
	 */
	StorageRebate *SafeBfcBigInt[uint64] `json:"storageRebate,omitempty"`
	/**
	 * Display metadata for this object, default to be undefined unless BfcObjectDataOptions.showDisplay is set to true
	 * This can also be None if the struct type does not have Display defined
	 */
	Display interface{} `json:"display,omitempty"`
}

func (data *BfcObjectData) Reference() bfc_types.ObjectRef {
	return bfc_types.ObjectRef{
		ObjectId: data.ObjectId,
		Version:  data.Version.data,
		Digest:   data.Digest.Data(),
	}
}

type BfcObjectDataOptions struct {
	/* Whether to fetch the object type, default to be false */
	ShowType bool `json:"showType,omitempty"`
	/* Whether to fetch the object content, default to be false */
	ShowContent bool `json:"showContent,omitempty"`
	/* Whether to fetch the object content in BCS bytes, default to be false */
	ShowBcs bool `json:"showBcs,omitempty"`
	/* Whether to fetch the object owner, default to be false */
	ShowOwner bool `json:"showOwner,omitempty"`
	/* Whether to fetch the previous transaction digest, default to be false */
	ShowPreviousTransaction bool `json:"showPreviousTransaction,omitempty"`
	/* Whether to fetch the storage rebate, default to be false */
	ShowStorageRebate bool `json:"showStorageRebate,omitempty"`
	/* Whether to fetch the display metadata, default to be false */
	ShowDisplay bool `json:"showDisplay,omitempty"`
}

type BfcObjectResponseError struct {
	NotExists *struct {
		ObjectId bfc_types.ObjectID `json:"object_id"`
	} `json:"notExists,omitempty"`
	Deleted *struct {
		ObjectId bfc_types.ObjectID       `json:"object_id"`
		Version  bfc_types.SequenceNumber `json:"version"`
		Digest   bfc_types.ObjectDigest   `json:"digest"`
	} `json:"deleted,omitempty"`
	UnKnown      *struct{} `json:"unKnown"`
	DisplayError *struct {
		Error string `json:"error"`
	} `json:"displayError"`
}

func (e BfcObjectResponseError) Tag() string {
	return "code"
}

func (e BfcObjectResponseError) Content() string {
	return ""
}

type BfcObjectResponse struct {
	Data  *BfcObjectData                       `json:"data,omitempty"`
	Error *lib.TagJson[BfcObjectResponseError] `json:"error,omitempty"`
}

type CheckpointSequenceNumber = uint64
type CheckpointedObjectId struct {
	ObjectId     bfc_types.ObjectID                       `json:"objectId"`
	AtCheckpoint *SafeBfcBigInt[CheckpointSequenceNumber] `json:"atCheckpoint"`
}

type ObjectsPage = Page[BfcObjectResponse, bfc_types.ObjectID]

// TODO need use Enum
type BfcObjectDataFilter struct {
	Package    *bfc_types.ObjectID `json:"Package,omitempty"`
	MoveModule *MoveModule         `json:"MoveModule,omitempty"`
	StructType string              `json:"StructType,omitempty"`
}

type BfcObjectResponseQuery struct {
	Filter  *BfcObjectDataFilter  `json:"filter,omitempty"`
	Options *BfcObjectDataOptions `json:"options,omitempty"`
}

type BfcPastObjectResponse = lib.TagJson[BfcPastObject]

// TODO need test VersionNotFound
type BfcPastObject struct {
	/// The object exists and is found with this version
	VersionFound *BfcObjectData `json:"VersionFound,omitempty"`
	/// The object does not exist
	ObjectNotExists *bfc_types.ObjectID `json:"ObjectNotExists,omitempty"`
	/// The object is found to be deleted with this version
	ObjectDeleted *BfcObjectRef `json:"ObjectDeleted,omitempty"`
	/// The object exists but not found with this version
	VersionNotFound *struct{ ObjectId bfc_types.SequenceNumber } `json:"VersionNotFound,omitempty"`
	/// The asked object version is higher than the latest
	VersionTooHigh *struct {
		ObjectId      bfc_types.ObjectID       `json:"object_id"`
		AskedVersion  bfc_types.SequenceNumber `json:"asked_version"`
		LatestVersion bfc_types.SequenceNumber `json:"latest_version"`
	} `json:"VersionTooHigh,omitempty"`
}

func (s BfcPastObject) Tag() string {
	return "status"
}

func (s BfcPastObject) Content() string {
	return "details"
}
