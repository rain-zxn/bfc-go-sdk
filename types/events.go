package types

import "github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"

type EventId struct {
	TxDigest bfc_types.TransactionDigest `json:"txDigest"`
	EventSeq SafeBfcBigInt[uint64]       `json:"eventSeq"`
}

type MoveEventField struct {
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type BfcEvent struct {
	Id EventId `json:"id"`
	// Move package where this event was emitted.
	PackageId bfc_types.ObjectID `json:"packageId"`
	// Move module where this event was emitted.
	TransactionModule string `json:"transactionModule"`
	// Sender's Bfc bfc_types.address.
	Sender bfc_types.BfcAddress `json:"sender"`
	// Move event type.
	Type string `json:"type"`
	// Parsed json value of the event
	ParsedJson interface{} `json:"parsedJson,omitempty"`
	// Base 58 encoded bcs bytes of the move event
	Bcs         string                 `json:"bcs"`
	TimestampMs *SafeBfcBigInt[uint64] `json:"timestampMs,omitempty"`
}

type EventFilter struct {
	/// Query by sender bfc_types.address.
	Sender *bfc_types.BfcAddress `json:"Sender,omitempty"`
	/// Return events emitted by the given transaction.
	Transaction *bfc_types.TransactionDigest `json:"Transaction,omitempty"`
	///digest of the transaction, as base-64 encoded string

	/// Return events emitted in a specified Package.
	Package *bfc_types.ObjectID `json:"Package,omitempty"`
	/// Return events emitted in a specified Move module.
	MoveModule *MoveModule `json:"MoveModule,omitempty"`
	/// Return events with the given move event struct name
	MoveEventType  *string         `json:"MoveEventType,omitempty"`
	MoveEventField *MoveEventField `json:"MoveEventField,omitempty"`
	/// Return events emitted in [start_time, end_time] interval
	TimeRange *TimeRange `json:"TimeRange,omitempty"`

	All *[]EventFilter `json:"All,omitempty"`
	Any *[]EventFilter `json:"Any,omitempty"`
	//And *struct {
	//	*EventFilter
	//	*EventFilter
	//} `json:"And,omitempty"`
	//Or *struct {
	//	EventFilter
	//	EventFilter
	//} `json:"Or,omitempty"`
}

type EventPage = Page[BfcEvent, EventId]
