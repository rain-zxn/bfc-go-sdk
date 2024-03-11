package types

import (
	"github.com/hellokittyboy-code/benfen-go-sdk/bfc_types"
	"github.com/hellokittyboy-code/benfen-go-sdk/lib"
)

type DynamicFieldInfo struct {
	Name bfc_types.DynamicFieldName `json:"name"`
	//Base58
	BcsName    lib.Base58                              `json:"bcsName"`
	Type       lib.TagJson[bfc_types.DynamicFieldType] `json:"type"`
	ObjectType string                                  `json:"objectType"`
	ObjectId   bfc_types.ObjectID                      `json:"objectId"`
	Version    bfc_types.SequenceNumber                `json:"version"`
	Digest     bfc_types.ObjectDigest                  `json:"digest"`
}

type DynamicFieldPage = Page[DynamicFieldInfo, bfc_types.ObjectID]
