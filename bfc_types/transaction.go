package bfc_types

import "github.com/hellokittyboy-code/benfen-go-sdk/lib"

var (
	SuiSystemMut = CallArg{
		Object: &SuiSystemMutObj,
	}

	SuiSystemMutObj = ObjectArg{
		SharedObject: &SharedObject{
			Id:                   *SuiSystemStateObjectId,
			InitialSharedVersion: SuiSystemStateObjectSharedVersion,
			Mutable:              true,
		},
	}

	BenfenSystemMut = CallArg{
		Object: &BenfenSystemMutObj,
	}

	BenfenSystemMutObj = ObjectArg{
		SharedObject: &SharedObject{
			Id:                   *BenfenSystemStateObjectId,
			InitialSharedVersion: BenfenSystemStateObjectSharedVersion,
			Mutable:              true,
		},
	}
)

func NewProgrammableAllowSponsor(
	sender BfcAddress,
	gasPayment []*ObjectRef,
	pt ProgrammableTransaction,
	gasBudge,
	gasPrice uint64,
	sponsor BfcAddress,
) TransactionData {
	kind := TransactionKind{
		ProgrammableTransaction: &pt,
	}
	return newWithGasCoinsAllowSponsor(kind, sender, gasPayment, gasBudge, gasPrice, sponsor)
}

func NewProgrammable(
	sender BfcAddress,
	gasPayment []*ObjectRef,
	pt ProgrammableTransaction,
	gasBudget uint64,
	gasPrice uint64,
) TransactionData {
	return NewProgrammableAllowSponsor(sender, gasPayment, pt, gasBudget, gasPrice, sender)
}

func newWithGasCoinsAllowSponsor(
	kind TransactionKind,
	sender BfcAddress,
	gasPayment []*ObjectRef,
	gasBudget uint64,
	gasPrice uint64,
	gasSponsor BfcAddress,
) TransactionData {
	return TransactionData{
		V1: &TransactionDataV1{
			Kind:   kind,
			Sender: sender,
			GasData: GasData{
				Price:   gasPrice,
				Owner:   gasSponsor,
				Payment: gasPayment,
				Budget:  gasBudget,
			},
			Expiration: TransactionExpiration{
				None: &lib.EmptyEnum{},
			},
		},
	}
}
