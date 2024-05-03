package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFedWireMessage_verifyIssue92(t *testing.T) {
	fwm := issue92FedWireMessage()
	require.NoError(t, fwm.verifyWithOpts(ValidateOpts{}))
}

// this is the payload reported in issue 92 (bug in fwm validation)
func issue92FedWireMessage() *FEDWireMessage {
	fwm := &FEDWireMessage{}
	fwm.MessageDisposition = &MessageDisposition{
		tag:                TagMessageDisposition,
		FormatVersion:      FormatVersion,
		TestProductionCode: EnvironmentTest,
	}
	fwm.SenderSupplied = &SenderSupplied{
		tag:                    TagSenderSupplied,
		FormatVersion:          FormatVersion,
		UserRequestCorrelation: "TESTDATA",
		TestProductionCode:     EnvironmentTest,
		MessageDuplicationCode: MessageDuplicationOriginal,
	}
	fwm.TypeSubType = &TypeSubType{
		tag:         TagTypeSubType,
		TypeCode:    FundsTransfer,
		SubTypeCode: BasicFundsTransfer,
	}
	fwm.InputMessageAccountabilityData = &InputMessageAccountabilityData{
		tag:                 TagInputMessageAccountabilityData,
		InputCycleDate:      "20180922",
		InputSource:         "XYZ ABC",
		InputSequenceNumber: "000001",
	}
	fwm.Amount = &Amount{
		tag:    TagAmount,
		Amount: "000000000250",
	}
	fwm.SenderDepositoryInstitution = &SenderDepositoryInstitution{
		tag:             TagSenderDepositoryInstitution,
		SenderABANumber: "000714895",
		SenderShortName: "Fake Institution",
	}
	fwm.ReceiverDepositoryInstitution = &ReceiverDepositoryInstitution{
		tag:               TagReceiverDepositoryInstitution,
		ReceiverABANumber: "000738119",
		ReceiverShortName: "Fake Institution",
	}
	fwm.BusinessFunctionCode = &BusinessFunctionCode{
		tag:                  TagBusinessFunctionCode,
		BusinessFunctionCode: CustomerTransfer,
		TransactionTypeCode:  "   ",
	}
	fwm.Charges = &Charges{
		tag:                 TagCharges,
		ChargeDetails:       CDBeneficiary,
		SendersChargesOne:   "USD0",
		SendersChargesTwo:   "USD0",
		SendersChargesThree: "USD0",
		SendersChargesFour:  "USD0",
	}
	fwm.InstructedAmount = &InstructedAmount{
		tag:          TagInstructedAmount,
		CurrencyCode: "USD",
		Amount:       "000000000250",
	}
	fwm.ExchangeRate = &ExchangeRate{
		tag:          TagExchangeRate,
		ExchangeRate: "1.3624055125",
	}
	fwm.Beneficiary = &Beneficiary{
		tag: TagBeneficiary,
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
		tag: TagOriginator,
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
