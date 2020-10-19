package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFedWireMessage_verifyIssue92(t *testing.T) {
	fwm := issue92FedWireMessage()
	require.NoError(t, fwm.verify())
}

// this is the payload reported in issue 92 (bug in fwm validation)
func issue92FedWireMessage() *FEDWireMessage {
	fwm := &FEDWireMessage{}
	fwm.MessageDisposition = &MessageDisposition{
		FormatVersion:      FormatVersion,
		TestProductionCode: EnvironmentTest,
	}
	fwm.SenderSupplied = &SenderSupplied{
		FormatVersion:          FormatVersion,
		UserRequestCorrelation: "TESTDATA",
		TestProductionCode:     EnvironmentTest,
		MessageDuplicationCode: " ",
	}
	fwm.TypeSubType = &TypeSubType{
		TypeCode:    FundsTransfer,
		SubTypeCode: BasicFundsTransfer,
	}
	fwm.InputMessageAccountabilityData = &InputMessageAccountabilityData{
		InputCycleDate:      "20180922",
		InputSource:         "XYZ ABC",
		InputSequenceNumber: "000001",
	}
	fwm.Amount = &Amount{
		Amount: "000000000250",
	}
	fwm.SenderDepositoryInstitution = &SenderDepositoryInstitution{
		SenderABANumber: "000714895",
		SenderShortName: "Fake Institution",
	}
	fwm.ReceiverDepositoryInstitution = &ReceiverDepositoryInstitution{
		ReceiverABANumber: "000738119",
		ReceiverShortName: "Fake Institution",
	}
	fwm.BusinessFunctionCode = &BusinessFunctionCode{
		BusinessFunctionCode: CustomerTransfer,
		TransactionTypeCode:  "   ",
	}
	fwm.Charges = &Charges{
		ChargeDetails:       CDBeneficiary,
		SendersChargesOne:   "USD0",
		SendersChargesTwo:   "USD0",
		SendersChargesThree: "USD0",
		SendersChargesFour:  "USD0",
	}
	fwm.InstructedAmount = &InstructedAmount{
		CurrencyCode: "USD",
		Amount:       "000000000250",
	}
	fwm.ExchangeRate = &ExchangeRate{
		ExchangeRate: "1.3624055125",
	}
	fwm.Beneficiary = &Beneficiary{
		Personal: Personal{
			IdentificationCode: SWIFTBICORBEIANDAccountNumber,
			Identifier:         "755756",
			Name:               "string",
			Address: Address{
				AddressLineOne:   " ",
				AddressLineTwo:   " ",
				AddressLineThree: " ",
			},
		},
	}
	fwm.Originator = &Originator{
		Personal: Personal{
			IdentificationCode: SWIFTBICORBEIANDAccountNumber,
			Identifier:         "798260",
			Name:               "string",
			Address: Address{
				AddressLineOne:   " ",
				AddressLineTwo:   " ",
				AddressLineThree: " ",
			},
		},
	}

	return fwm
}
