package move_types

import (
	"github.com/fardream/go-bcs/bcs"
)

type U128 [2]uint64

func (a U128) ToUint128() bcs.Uint128 {
	return *bcs.NewUint128FromUint64(a[0], a[1])
}

func (a U128) String() string {
	return a.ToUint128().String()
}

func (a U128) MarshalJSON() ([]byte, error) {
	return a.ToUint128().MarshalJSON()
}

func (a *U128) UnmarshalJSON(data []byte) error {
	i := bcs.NewUint128FromUint64(0, 0)
	return i.UnmarshalJSON(data)
}

func (a U128) MarshalBCS() ([]byte, error) {
	return a.ToUint128().MarshalBCS()
}
