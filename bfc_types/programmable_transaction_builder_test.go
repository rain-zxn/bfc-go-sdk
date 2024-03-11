package bfc_types

import (
	"testing"

	"github.com/fardream/go-bcs/bcs"
	"github.com/hellokittyboy-code/benfen-go-sdk/move_types"
	"github.com/stretchr/testify/require"
)

func TestTransferBFC(t *testing.T) {
	ptb := NewProgrammableTransactionBuilder()
	recipient, err := NewAddressFromHex("0x7e875ea78ee09f08d72e2676cf84e0f1c8ac61d94fa339cc8e37cace85bebc6e")
	require.NoError(t, err)
	amount := uint64(100000)
	err = ptb.TransferBFC(*recipient, &amount)
	require.NoError(t, err)
	pt := ptb.Finish()
	digest, err := NewDigest("HvbE2UZny6cP4KukaXetmj4jjpKTDTjVo23XEcu7VgSn")
	require.NoError(t, err)
	objectId, err := move_types.NewAccountAddressHex("0x13c1c3d0e15b4039cec4291c75b77c972c10c8e8e70ab4ca174cf336917cb4db")
	require.NoError(t, err)
	tx := NewProgrammable(
		*recipient, []*ObjectRef{
			{
				ObjectId: *objectId,
				Version:  14924029,
				Digest:   *digest,
			},
		}, pt, 100000000, 1000,
	)
	txByte, err := bcs.Marshal(tx)
	require.NoError(t, err)
	t.Logf("%x", txByte)
}