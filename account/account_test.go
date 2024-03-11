package account

import (
	"encoding/json"
	"testing"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/stretchr/testify/require"
)

var Mnemonic = "monkey tragic drive owner fade mimic taxi despair endorse peasant amused woman"

// expected : monkey tragic drive owner fade mimic taxi despair endorse peasant amused woman
func TestMyAccouunt(t *testing.T) {
	account, err := NewAccountWithMnemonic(Mnemonic)
	require.Nil(t, err)

	t.Logf("addr = %v", account.Address)
}

func Test_Signature_Marshal_Unmarshal(t *testing.T) {
	account, err := NewAccountWithMnemonic(Mnemonic)
	require.Nil(t, err)

	msg := "Test_Signature_Marshal_Unmarshal"
	msgBytes := []byte(msg)

	signature1, err := account.SignSecureWithoutEncode(msgBytes, bfc_types.DefaultIntent())
	require.Nil(t, err)

	marshaedData, err := json.Marshal(signature1)
	require.Nil(t, err)

	var signature2 bfc_types.Signature
	err = json.Unmarshal(marshaedData, &signature2)
	require.Nil(t, err)

	require.Equal(t, signature1, signature2)
}
