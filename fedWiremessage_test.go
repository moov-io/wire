package wire

import (
	"github.com/moov-io/base"
	"testing"
)

func mockCustomerTransferData() FEDWireMessage {
	fwm := NewFEDWireMessage()

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SetSenderSupplied(ss)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	amt := mockAmount()
	fwm.SetAmount(amt)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "   "
	fwm.SetBusinessFunctionCode(bfc)
	// Beneficiary
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	return fwm
}

func TestFEDWireMessage_isAmountValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.Amount.Amount = "000000000000"
	//fwm.SetAmount(fwm.Amount)
	file.AddFEDWireMessage(fwm)
	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// Validate File
	if err := file.Validate(); err != nil {
		if err != NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount, "SubTypeCode",
			fwm.TypeSubType.SubTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isPreviousMessageIdentifierValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.TypeSubType.SubTypeCode = "02"
	file.AddFEDWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// Validate File
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isLocalInstrumentCodeValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	file.AddFEDWireMessage(fwm)

	if err := fwm.isLocalInstrumentCodeValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isChargesValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	c := mockCharges()
	fwm.SetCharges(c)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	file.AddFEDWireMessage(fwm)

	if err := fwm.isChargesValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
			"Charges", fwm.Charges.String()) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isInstructedAmountValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	file.AddFEDWireMessage(fwm)

	if err := fwm.isInstructedAmountValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("LocalInstrumentCode",
			fwm.LocalInstrument.LocalInstrumentCode, "Instructed Amount", fwm.InstructedAmount.String()) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isExchangeRateRequired(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	file.AddFEDWireMessage(fwm)

	if err := fwm.isExchangeRateValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isExchangeRateValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	file.AddFEDWireMessage(fwm)

	if err := fwm.isExchangeRateValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("LocalInstrumentCode",
			fwm.LocalInstrument.LocalInstrumentCode, "ExchangeRate", fwm.ExchangeRate.ExchangeRate) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isBeneficiaryIntermediaryFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	//li := mockLocalInstrument()
	//li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	//fwm.SetLocalInstrument(li)
	//eRate := mockExchangeRate()
	//fwm.SetExchangeRate(eRate)
	//ia := mockInstructedAmount()
	//fwm.SetInstructedAmount(ia)
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	file.AddFEDWireMessage(fwm)

	// BeneficiaryFI required field check
	if err := fwm.isBeneficiaryIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	if err := fwm.isBeneficiaryIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isBeneficiaryFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	file.AddFEDWireMessage(fwm)
	// Beneficiary required field check
	if err := fwm.isBeneficiaryIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isOriginatorFIValid(t *testing.T) {
}

func TestFEDWireMessage_isFIIntermediaryFIValid(t *testing.T) {
}

func TestFEDWireMessage_isFIIntermediaryFIAdviceValid(t *testing.T) {
}

func TestFEDWireMessage_FIBeneficiaryFI(t *testing.T) {
}

func TestFEDWireMessage_isFIBeneficiaryFIAdvice(t *testing.T) {
}

func TestFEDWireMessage_isFIBeneficiary(t *testing.T) {
}

func TestFEDWireMessage_isFIBeneficiaryAdvice(t *testing.T) {
}

func TestFEDWireMessage_isUnstructuredAddendaValid(t *testing.T) {
}

func TestFEDWireMessage_isRemittanceOriginatorValid(t *testing.T) {
}

func TestFEDWireMessage_isRemittanceBeneficiaryValid(t *testing.T) {
}

func TestFEDWireMessage_isPrimaryRemittanceDocumentValid(t *testing.T) {
}

func TestFEDWireMessage_isActualAmountPaidValid(t *testing.T) {
}

func TestFEDWireMessage_isGrossAmountRemittanceDocument(t *testing.T) {
}

func TestFEDWireMessage_is(t *testing.T) {
}

func TestFEDWireMessage_isAdjustmentValid(t *testing.T) {
}

func TestFEDWireMessage_isDateRemittanceDocumentValid(t *testing.T) {
}

func TestFEDWireMessage_isSecondaryRemittanceDocumen(t *testing.T) {
}

func TestFEDWireMessage_isRemittanceFreeText(t *testing.T) {
}

func TestFEDWireMessage_isRelatedRemittanceValid(t *testing.T) {
}

func TestFEDWireMessage_isGrossAmountRemittanceDocumentValid(t *testing.T) {
}

func TestFEDWireMessage_isRemittanceFreeTextValid(t *testing.T) {
}


