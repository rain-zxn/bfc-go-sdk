package bfc_types

var (
	SuiSystemAddress, _                  = NewAddressFromHex("0x3")
	SuiSystemPackageId                   = SuiSystemAddress
	SuiSystemStateObjectId, _            = NewObjectIdFromHex("0x5")
	SuiSystemStateObjectSharedVersion    = ObjectStartVersion
	BenfenSystemStateObjectId, _         = NewObjectIdFromHex("0xc9")
	BenfenSystemStateObjectSharedVersion = ObjectStartVersion
)
