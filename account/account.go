package account

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/crypto"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

const (
	ADDRESS_LENGTH = 64
)

type Account struct {
	KeyPair bfc_types.BfcKeyPair
	Address string
}

func NewAccount(scheme bfc_types.SignatureScheme, seed []byte) *Account {
	bfcKeyPair := bfc_types.NewBfcKeyPair(scheme, seed)
	tmp := []byte{scheme.Flag()}
	tmp = append(tmp, bfcKeyPair.PublicKey()...)
	addrBytes := blake2b.Sum256(tmp)
	address := "0x" + hex.EncodeToString(addrBytes[:])[:ADDRESS_LENGTH]

	return &Account{
		KeyPair: bfcKeyPair,
		Address: address,
	}
}

func NewAccountWithKeystore(keystore string) (*Account, error) {
	ksByte, err := base64.StdEncoding.DecodeString(keystore)
	if err != nil {
		return nil, err
	}
	scheme, err := bfc_types.NewSignatureScheme(ksByte[0])
	if err != nil {
		return nil, err
	}
	return NewAccount(scheme, ksByte[1:]), nil
}

func NewAccountWithMnemonic(mnemonic string) (*Account, error) {
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}
	key, err := crypto.DeriveForPath("m/44'/784'/0'/0'/0'", seed)
	if err != nil {
		return nil, err
	}
	scheme, err := bfc_types.NewSignatureScheme(0)
	if err != nil {
		return nil, err
	}
	return NewAccount(scheme, key.Key), nil
}

func (a *Account) Sign(data []byte) []byte {
	switch a.KeyPair.Flag() {
	case 0:
		return a.KeyPair.Ed25519.Sign(data)
	default:
		return []byte{}
	}
}

func (a *Account) SignSecureWithoutEncode(msg lib.Base64Data, intent bfc_types.Intent) (bfc_types.Signature, error) {
	signature, err := bfc_types.NewSignatureSecure(
		bfc_types.NewIntentMessage(intent, msg), &a.KeyPair,
	)
	if err != nil {
		return bfc_types.Signature{}, err
	}
	return signature, nil
}
