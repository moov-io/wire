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
	return fwm
}

func TestFEDWireMessage_isAmountValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.Amount.Amount = "000000000000"
	//fwm.SetAmount(fwm.Amount)
	// Beneficiary
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
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
	if err := fwm.isPreviousMessageIdentifierValid(); err != nil {
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
	if err := fwm.isBeneficiaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}

}

func TestFEDWireMessage_isOriginatorFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// Originator required field check
	if err := fwm.isOriginatorFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	file.AddFEDWireMessage(fwm)
	// OriginatorOptionF required field check
	if err := fwm.isOriginatorFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isInstructingFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// Originator required field check
	if err := fwm.isInstructingFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	file.AddFEDWireMessage(fwm)
	// OriginatorOptionF required field check
	if err := fwm.isInstructingFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestNewFEDWireMessage_isOriginatorToBeneficiaryValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// Originator required field check
	if err := fwm.isOriginatorToBeneficiaryValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	file.AddFEDWireMessage(fwm)
	// Beneficiary required field check
	if err := fwm.isOriginatorToBeneficiaryValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	// Originator required Field check
	if err := fwm.isOriginatorToBeneficiaryValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	// OriginatorOptionF required Field check
	if err := fwm.isOriginatorToBeneficiaryValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isFIIntermediaryFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// BeneficiaryIntermediaryFI required field check
	if err := fwm.isFIIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	if err := fwm.isFIIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	if err := fwm.isFIIntermediaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isFIIntermediaryFIAdviceValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// BeneficiaryIntermediaryFI required field check
	if err := fwm.isFIIntermediaryFIAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	if err := fwm.isFIIntermediaryFIAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	if err := fwm.isFIIntermediaryFIAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_FIBeneficiaryFIValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// BeneficiaryFI required field check
	if err := fwm.isFIBeneficiaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	if err := fwm.isFIBeneficiaryFIValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isFIBeneficiaryFIAdvice(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// BeneficiaryFI required field check
	if err := fwm.isFIBeneficiaryFIAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	if err := fwm.isFIBeneficiaryFIAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isFIBeneficiary(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// Beneficiary required field check
	if err := fwm.isFIBeneficiaryValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isFIBeneficiaryAdvice(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	file.AddFEDWireMessage(fwm)
	// Beneficiary required field check
	if err := fwm.isFIBeneficiaryAdviceValid(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isUnstructuredAddendaValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	file.AddFEDWireMessage(fwm)
	// UnstructuredAddenda Invalid Property
	if err := fwm.isUnstructuredAddendaValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isRelatedRemittanceValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	file.AddFEDWireMessage(fwm)
	// RelatedRemittance Invalid Property
	if err := fwm.isRelatedRemittanceValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("RelatedRemittance", fwm.RelatedRemittance.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isRemittanceOriginatorValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	file.AddFEDWireMessage(fwm)
	// RemittanceOriginator Invalid Property
	if err := fwm.isRemittanceOriginatorValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("RemittanceOriginator", fwm.RemittanceOriginator.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isRemittanceBeneficiaryValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)
	file.AddFEDWireMessage(fwm)
	// RemittanceBeneficiary Invalid Property
	if err := fwm.isRemittanceBeneficiaryValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("RemittanceBeneficiary", fwm.RemittanceBeneficiary.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isPrimaryRemittanceDocumentValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	file.AddFEDWireMessage(fwm)
	// PrimaryRemittanceDocument Invalid Property
	if err := fwm.isPrimaryRemittanceDocumentValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("PrimaryRemittanceDocument", fwm.PrimaryRemittanceDocument.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isActualAmountPaidValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)
	file.AddFEDWireMessage(fwm)
	// ActualAmountPaid Invalid Property
	if err := fwm.isActualAmountPaidValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("ActualAmountPaid", fwm.ActualAmountPaid.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isGrossAmountRemittanceDocument(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	gard := mockGrossAmountRemittanceDocument()
	fwm.SetGrossAmountRemittanceDocument(gard)
	file.AddFEDWireMessage(fwm)
	// GrossAmountRemittanceDocument Invalid Property
	if err := fwm.isGrossAmountRemittanceDocumentValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("GrossAmountRemittanceDocument", fwm.GrossAmountRemittanceDocument.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isAdjustmentValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)
	file.AddFEDWireMessage(fwm)
	// Adjustment Invalid Property
	if err := fwm.isAdjustmentValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("Adjustment", fwm.Adjustment.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isDateRemittanceDocumentValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	drd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(drd)
	file.AddFEDWireMessage(fwm)
	// DateRemittanceDocument Invalid Property
	if err := fwm.isDateRemittanceDocumentValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("DateRemittanceDocument", fwm.DateRemittanceDocument.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isSecondaryRemittanceDocumentValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	srd := mockSecondaryRemittanceDocument()
	fwm.SetSecondaryRemittanceDocument(srd)
	file.AddFEDWireMessage(fwm)
	// SecondaryRemittanceDocument Invalid Property
	if err := fwm.isSecondaryRemittanceDocumentValid(); err != nil {
		if err != NewErrInvalidPropertyForProperty("SecondaryRemittanceDocument", fwm.SecondaryRemittanceDocument.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isRemittanceFreeText(t *testing.T) {
}

func TestFEDWireMessage_isGrossAmountRemittanceDocumentValid(t *testing.T) {
}

func TestFEDWireMessage_isRemittanceFreeTextValid(t *testing.T) {
}

/*// Beneficiary
ben := mockBeneficiary()
fwm.SetBeneficiary(ben)

// Originator
o := mockOriginator()
fwm.SetOriginator(o)*/
