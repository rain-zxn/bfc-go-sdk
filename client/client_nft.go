package client

import (
	"context"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/types"
)

// MintNFT
// Create an unsigned transaction to mint a nft at devnet
func (c *Client) MintNFT(
	ctx context.Context,
	signer BfcAddress,
	nftName, nftDescription, nftUri string,
	gas *bfcObjectID,
	gasBudget uint64,
) (*types.TransactionBytes, error) {
	packageId, _ := bfc_types.NewAddressFromHex("0x2")
	args := []any{
		nftName, nftDescription, nftUri,
	}
	return c.MoveCall(
		ctx,
		signer,
		*packageId,
		"devnet_nft",
		"mint",
		[]string{},
		args,
		gas,
		types.NewSafeBfcBigInt(gasBudget),
	)
}

func (c *Client) GetNFTsOwnedByAddress(ctx context.Context, address BfcAddress) ([]types.BfcObjectResponse, error) {
	return c.BatchGetObjectsOwnedByAddress(
		ctx, address, types.BfcObjectDataOptions{
			ShowType:    true,
			ShowContent: true,
			ShowOwner:   true,
		}, "0x2::devnet_nft::DevNetNFT",
	)
}
