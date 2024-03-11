package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFaucetFundAccount_Devnet(t *testing.T) {
	// addr := M1Account(t).Address
	addr := Address.String()

	res, err := FaucetFundAccount(addr, DevNetFaucetUrl)
	require.Nil(t, err)
	t.Log("hash = ", res)
}

// func TestFaucetFundAccount_Testnet(t *testing.T) {
// 	addr := "0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea"
// 	res, err := FaucetFundAccount(addr, TestNetFaucetUrl)
// 	require.Nil(t, err)
// }
