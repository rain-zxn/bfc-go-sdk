package types

import (
	"fmt"
	"math/big"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/shopspring/decimal"
)

type BfcBigInt = decimal.Decimal

type SafeBigInt interface {
	~int64 | ~uint64
}

func NewSafeBfcBigInt[T SafeBigInt](num T) SafeBfcBigInt[T] {
	return SafeBfcBigInt[T]{
		data: num,
	}
}

type SafeBfcBigInt[T SafeBigInt] struct {
	data T
}

func (s *SafeBfcBigInt[T]) UnmarshalText(data []byte) error {
	return s.UnmarshalJSON(data)
}

func (s *SafeBfcBigInt[T]) UnmarshalJSON(data []byte) error {
	num := decimal.NewFromInt(0)
	err := num.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	if num.BigInt().IsInt64() {
		s.data = T(num.BigInt().Int64())
		return nil
	}

	if num.BigInt().IsUint64() {
		s.data = T(num.BigInt().Uint64())
		return nil
	}
	return fmt.Errorf("json data [%s] is not T", string(data))
}

func (s SafeBfcBigInt[T]) MarshalJSON() ([]byte, error) {
	return decimal.NewFromInt(int64(s.data)).MarshalJSON()
}

func (s SafeBfcBigInt[T]) Int64() int64 {
	return int64(s.data)
}

func (s SafeBfcBigInt[T]) Uint64() uint64 {
	return uint64(s.data)
}

func (s *SafeBfcBigInt[T]) Decimal() decimal.Decimal {
	return decimal.NewFromBigInt(big.NewInt(0).SetUint64(s.Uint64()), 0)
}

// export const ObjectID = string();
// export type ObjectID = Infer<typeof ObjectID>;

// export const BfcAddress = string();
// export type BfcAddress = Infer<typeof BfcAddress>;

type ObjectOwnerInternal struct {
	AddressOwner *bfc_types.BfcAddress `json:"AddressOwner,omitempty"`
	ObjectOwner  *bfc_types.BfcAddress `json:"ObjectOwner,omitempty"`
	SingleOwner  *bfc_types.BfcAddress `json:"SingleOwner,omitempty"`
	Shared       *struct {
		InitialSharedVersion *bfc_types.SequenceNumber `json:"initial_shared_version"`
	} `json:"Shared,omitempty"`
}

type ObjectOwner struct {
	*ObjectOwnerInternal
	*string
}

type Page[T BfcTransactionBlockResponse | BfcEvent | Coin | BfcObjectResponse | DynamicFieldInfo,
	C bfc_types.TransactionDigest | EventId | bfc_types.ObjectID] struct {
	Data        []T  `json:"data"`
	NextCursor  *C   `json:"nextCursor,omitempty"`
	HasNextPage bool `json:"hasNextPage"`
}
