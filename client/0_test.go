package client

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
	"github.com/shopspring/decimal"

	"github.com/hellokittyboy-code/benfen-go-sdk/account"
	"github.com/stretchr/testify/require"
)

var (
	M1Mnemonic = "monkey tragic drive owner fade mimic taxi despair endorse peasant amused woman"

	Address, _ = bfc_types.NewAddressFromHex("0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea")
	CrossChainAddress, _ = bfc_types.NewAddressFromHex("0x7419050e564485685f306e20060472fca1b3a4453b41bdace0010624801b11ea")
)

func MainnetClient(t *testing.T) *Client {
	c, err := Dial(types.MainnetRpcUrl)
	require.NoError(t, err)
	return c
}

func TestnetClient(t *testing.T) *Client {
	c, err := Dial(types.TestnetRpcUrl)
	require.NoError(t, err)
	return c
}

func LocalnetClient(t *testing.T) *Client {
	c, err := Dial(types.LocalNetRpcUrl)
	require.NoError(t, err)
	return c
}

func DevnetClient(t *testing.T) *Client {
	c, err := Dial(types.DevNetRpcUrl)
	require.NoError(t, err)

	balance, err := c.GetBalance(context.Background(), *Address, types.BFC_COIN_TYPE)
	require.NoError(t, err)
	if balance.TotalBalance.BigInt().Uint64() < BFC(0.3).Uint64() {
		_, err = FaucetFundAccount(Address.String(), DevNetFaucetUrl)
		require.NoError(t, err)
	}
	return c
}

func ChainClient(t *testing.T) *Client {
	bfcEnv := os.Getenv("BFC_NETWORK")
	bfcEnv = "testnet"
	switch bfcEnv {
	case "testnet":
		return TestnetClient(t)
	case "devnet":
		return DevnetClient(t)
	case "":
		fallthrough
	default:

		println("using default chain: localnet")
		return LocalnetClient(t)
	}
}

func M1Account(t *testing.T) *account.Account {
	a, err := account.NewAccountWithMnemonic(M1Mnemonic)
	require.NoError(t, err)
	return a
}

func M1Address(t *testing.T) *BfcAddress {
	return Address
}

func Signer(t *testing.T) *account.Account {
	return M1Account(t)
}

type BFC float64

func (s BFC) Int64() int64 {
	return int64(s * 1e9)
}
func (s BFC) Uint64() uint64 {
	return uint64(s * 1e9)
}
func (s BFC) Decimal() decimal.Decimal {
	return decimal.NewFromInt(s.Int64())
}
func (s BFC) String() string {
	return strconv.FormatInt(s.Int64(), 10)
}
