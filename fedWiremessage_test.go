package wire

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"
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

func TestFEDWireMessage_invalidAmount(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error (can only be zeros is TypeSubType code is 90)
	fwm.Amount.Amount = "000000000000"
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
	err := file.Validate()

	require.Equal(t, NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount, "SubTypeCode",
		fwm.TypeSubType.SubTypeCode), err)
}

func TestFEDWireMessage_previousMessageIdentifierInvalid(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.TypeSubType.SubTypeCode = ReversalTransfer
	fwm.PreviousMessageIdentifier = nil // required when SubTypeCode is ReversalTransfer

	err := fwm.checkPreviousMessageIdentifier()

	require.Equal(t, fieldError("PreviousMessageIdentifier", ErrFieldRequired), err)
}

func TestFEDWireMessage_invalidLocalInstrument(t *testing.T) {
	fwm := mockCustomerTransferData()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	fwm.BusinessFunctionCode.BusinessFunctionCode = BankTransfer // local instrument only permitted for CTP

	err := fwm.validateLocalInstrumentCode()

	require.Equal(t, fieldError("LocalInstrument", ErrLocalInstrumentNotPermitted), err)
}

func TestFEDWireMessage_invalidCharges(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	c := mockCharges()
	fwm.SetCharges(c)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateCharges()

	require.Equal(t, NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"Charges", fwm.Charges.String()), err)
}

func TestFEDWireMessage_invalidInstructedAmount(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateInstructedAmount()

	require.Equal(t, NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"Instructed Amount", fwm.InstructedAmount.String()), err)
}

func TestFEDWireMessage_validateExchangeRate_missingInstructedAmount(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateExchangeRate()

	require.Equal(t, fieldError("InstructedAmount", ErrFieldRequired), err)
}

func TestFEDWireMessage_isExchangeRateValid_missingLocalInstrumentCode(t *testing.T) {
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

	err := fwm.validateExchangeRate()

	require.Equal(t, NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"ExchangeRate", fwm.ExchangeRate.ExchangeRate), err)
}

func TestFEDWireMessage_validateBeneficiaryIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateBeneficiaryIntermediaryFI()
	require.Equal(t, fieldError("BeneficiaryFI", ErrFieldRequired), err)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateBeneficiaryIntermediaryFI()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateBeneficiaryFI()

	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateOriginatorFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateOriginatorFI()
	require.Equal(t, fieldError("Originator", ErrFieldRequired), err)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateOriginatorFI()
	require.Equal(t, fieldError("OriginatorOptionF", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateInstructingFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateInstructingFI()
	require.Equal(t, fieldError("Originator", ErrFieldRequired), err)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateInstructingFI()
	require.Equal(t, fieldError("OriginatorOptionF", ErrFieldRequired), err)
}

func TestNewFEDWireMessage_validateOriginatorToBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateOriginatorToBeneficiary()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)

	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	// Originator required Field check
	err = fwm.validateOriginatorToBeneficiary()
	require.Equal(t, fieldError("Originator", ErrFieldRequired), err)

	o := mockOriginator()
	fwm.SetOriginator(o)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()
	require.Equal(t, fieldError("OriginatorOptionF", ErrFieldRequired), err)

	// check beneficiary still required
	fwm.SetBeneficiary(nil)
	err = fwm.validateOriginatorToBeneficiary()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateFIIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer
	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFI()
	require.Equal(t, fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired), err)

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFI()
	require.Equal(t, fieldError("BeneficiaryFI", ErrFieldRequired), err)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFI()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateFIIntermediaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFIAdvice()
	require.Equal(t, fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired), err)

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFIAdvice()
	require.Equal(t, fieldError("BeneficiaryFI", ErrFieldRequired), err)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFIAdvice()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateFIBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFI()
	require.Equal(t, fieldError("BeneficiaryFI", ErrFieldRequired), err)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFI()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateFIBeneficiaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFIAdvice()
	require.Equal(t, fieldError("BeneficiaryFI", ErrFieldRequired), err)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFIAdvice()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateFIBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	if err := fwm.validateFIBeneficiary(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_validateFIBeneficiaryAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiaryAdvice()
	require.Equal(t, fieldError("Beneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validateUnstructuredAddenda(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)

	// UnstructuredAddenda Invalid Property
	err := fwm.validateUnstructuredAddenda()
	require.Equal(t, NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.String(),
		"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode), err)
}

func TestFEDWireMessage_validateRelatedRemittance(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)

	// RelatedRemittance Invalid Property
	err := fwm.validateRelatedRemittance()
	require.Equal(t, fieldError("RelatedRemittance", ErrNotPermitted), err)
}

func TestFEDWireMessage_validateRemittanceOriginator(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)

	// RemittanceOriginator Invalid Property
	err := fwm.validateRemittanceOriginator()
	require.Equal(t, fieldError("RemittanceOriginator", ErrNotPermitted), err)
}

func TestFEDWireMessage_validateRemittanceBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)

	// RemittanceBeneficiary Invalid Property
	err := fwm.validateRemittanceBeneficiary()
	require.Equal(t, fieldError("RemittanceBeneficiary", ErrNotPermitted), err)

	fwm.RemittanceBeneficiary = nil
	fwm.LocalInstrument.LocalInstrumentCode = RemittanceInformationStructured
	// RemittanceBeneficiary Invalid Property
	err = fwm.validateRemittanceBeneficiary()
	require.Equal(t, fieldError("RemittanceBeneficiary", ErrFieldRequired), err)
}

func TestFEDWireMessage_validatePrimaryRemittanceDocument(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)

	// PrimaryRemittanceDocument Invalid Property
	err := fwm.validatePrimaryRemittanceDocument()
	if err != NewErrInvalidPropertyForProperty("PrimaryRemittanceDocument", fwm.PrimaryRemittanceDocument.String(),
		"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
		t.Errorf("%T: %s", err, err)
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
	if err := fwm.validateActualAmountPaid(); err != nil {
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
	if err := fwm.validateGrossAmountRemittanceDocument(); err != nil {
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
	if err := fwm.validateAdjustment(); err != nil {
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
	if err := fwm.validateDateRemittanceDocument(); err != nil {
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
	if err := fwm.validateSecondaryRemittanceDocument(); err != nil {
		if err != NewErrInvalidPropertyForProperty("SecondaryRemittanceDocument", fwm.SecondaryRemittanceDocument.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFEDWireMessage_isRemittanceFreeTextValid(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)
	file.AddFEDWireMessage(fwm)
	// RemittanceFreeTextValid Invalid Property
	if err := fwm.validateRemittanceFreeText(); err != nil {
		if err != NewErrInvalidPropertyForProperty("RemittanceFreeText", fwm.RemittanceFreeText.String(),
			"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBankTransferInValid test an invalid BankTransfer
func TestBankTransferInValid(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	tst := mockTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = RequestCredit
	fwm.SetTypeSubType(tst)
	if err := fwm.validateBankTransfer(); err != nil {
		if err != NewErrBusinessFunctionCodeProperty("TypeSubType", tst.TypeCode+tst.SubTypeCode,
			fwm.BusinessFunctionCode.BusinessFunctionCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestTransactionTypeCodeForBankTransfer test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)
	if err := fwm.checkProhibitedBankTransferTags(); err != nil {
		if !base.Match(err, ErrTransactionTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidLocalInstrumentForBankTransfer test an invalid LocalInstrument
func TestInvalidLocalInstrumentForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	if err := fwm.checkProhibitedBankTransferTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidPaymentNotificationForBankTransfer test an invalid PaymentNotification
func TestInvalidPaymentNotificationForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidChargesForBankTransfer test an invalid Charges
func TestInvalidChargesForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	c := mockCharges()
	fwm.SetCharges(c)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidInstructedAmountForBankTransfer test an invalid InstructedAmount
func TestInvalidInstructedAmountForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidExchangeRateForBankTransfer test an invalid ExchangeRate
func TestInvalidExchangeRateForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidBeneficiaryIdentificationCodeForBankTransfer test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetBeneficiary(ben)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountDebitedDrawdownForBankTransfer test an invalid AccountDebitedDrawdown
func TestInvalidAccountDebitedDrawdownForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorIdentificationCodeForBankTransfer test an invalid OriginatorIdentificationCode
func TestInvalidOriginatorIdentificationCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetOriginator(o)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorOptionFForBankTransfer test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	off := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(off)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountCreditedDrawdownForBankTransfer test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	creditDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(creditDD)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidFIDrawdownDebitAccountAdviceForBankTransfer test an invalid FIDrawdownDebitAccountAdvice
func TestInvalidFIDrawdownDebitAccountAdviceForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	fwm.SetFIDrawdownDebitAccountAdvice(debitDDAdvice)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidServiceMessageForBankTransfer test an invalid ServiceMessage
func TestInvalidServiceMessageForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidUnstructuredAddendaForBankTransfer test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidCurrencyInstructedAmountForBankTransfer test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidRelatedRemittanceForBankTransfer test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	err := fwm.checkProhibitedBankTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestTransactionTypeCodeForCustomerTransfer test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)
	if err := fwm.checkProhibitedCustomerTransferTags(); err != nil {
		if !base.Match(err, ErrTransactionTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidLocalInstrumentForCustomerTransfer test an invalid LocalInstrument
func TestInvalidLocalInstrumentForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	if err := fwm.checkProhibitedCustomerTransferTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidPaymentNotificationForCustomerTransfer test an invalid PaymentNotification
func TestInvalidPaymentNotificationForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidChargesForCustomerTransfer test an invalid Charges
func TestInvalidChargesForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	c := mockCharges()
	fwm.SetCharges(c)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidInstructedAmountForCustomerTransfer test an invalid InstructedAmount
func TestInvalidInstructedAmountForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidExchangeRateForCustomerTransfer test an invalid ExchangeRate
func TestInvalidExchangeRateForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidBeneficiaryIdentificationCodeForCustomerTransfer test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetBeneficiary(ben)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountDebitedDrawdownForCustomerTransfer test an invalid AccountDebitedDrawdown
func TestInvalidAccountDebitedDrawdownForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorIdentificationCodeForCustomerTransfer test an invalid OriginatorIdentificationCode
func TestInvalidOriginatorIdentificationCodeForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetOriginator(o)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorOptionFForCustomerTransfer test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	off := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(off)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountCreditedDrawdownForCustomerTransfer test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	creditDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(creditDD)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidFIDrawdownDebitAccountAdviceForCustomerTransfer test an invalid FIDrawdownDebitAccountAdvice
func TestInvalidFIDrawdownDebitAccountAdviceForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	fwm.SetFIDrawdownDebitAccountAdvice(debitDDAdvice)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidServiceMessageForCustomerTransfer test an invalid ServiceMessage
func TestInvalidServiceMessageForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidUnstructuredAddendaForCustomerTransfer test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidCurrencyInstructedAmountForCustomerTransfer test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidRelatedRemittanceForCustomerTransfer test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	err := fwm.checkProhibitedCustomerTransferTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentUnstructuredAddendaForCustomerTransferPlus tests UnstructuredAddenda is required for
// LocalInstrumentCode ANSIX12format
func TestLocalInstrumentUnstructuredAddendaForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentRelatedRemittanceForCustomerTransferPlus tests RelatedRemittance is required for
// LocalInstrumentCode RelatedRemittanceInformation
func TestLocalInstrumentRelatedRemittanceForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentBeneficiaryReferenceForCustomerTransferPlus tests BeneficiaryReference is required for
// LocalInstrumentCode SequenceBCoverPaymentStructured
func TestLocalInstrumentBeneficiaryReferenceForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentOrderingCustomerForCustomerTransferPlus tests OrderingCustomer is required for
// LocalInstrumentCode SequenceBCoverPaymentStructured
func TestLocalInstrumentOrderingCustomerForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentBeneficiaryCustomerForCustomerTransferPlus tests BeneficiaryCustomer is required for
// LocalInstrumentCode SequenceBCoverPaymentStructured
func TestLocalInstrumentBeneficiaryCustomerForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	oc := mockOrderingCustomer()
	fwm.SetOrderingCustomer(oc)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentProprietaryCodeForCustomerTransferPlus tests ProprietaryCode is required for
// LocalInstrumentCode SequenceBCoverPaymentStructured
func TestLocalInstrumentProprietaryCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	fwm.SetLocalInstrument(li)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentRemittanceOriginatorForCustomerTransferPlus tests RemittanceOriginator is required for
// LocalInstrumentCode RemittanceInformationStructured
func TestLocalInstrumentRemittanceOriginatorForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentRemittanceBeneficiaryForCustomerTransferPlus tests RemittanceBeneficiary is required for
// LocalInstrumentCode RemittanceInformationStructured
func TestLocalInstrumentRemittanceBeneficiaryForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentPrimaryRemittanceDocumentForCustomerTransferPlus tests PrimaryRemittanceDocument is required for
// LocalInstrumentCode RemittanceInformationStructured
func TestLocalInstrumentPrimaryRemittanceDocumentForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLocalInstrumentActualAmountPaidForCustomerTransferPlus tests ActualAmountPaid is required for
// LocalInstrumentCode RemittanceInformationStructured
func TestLocalInstrumentActualAmountPaidForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	o := mockOriginator()
	fwm.SetOriginator(o)
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.SetLocalInstrument(li)
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryForCustomerTransferPlus tests a Beneficiary is required
func TestBeneficiaryIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	err := fwm.checkMandatoryCustomerTransferPlusTags()
	if err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorForCustomerTransferPlus tests an Originator is required
func TestOriginatorIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	err := fwm.checkMandatoryCustomerTransferPlusTags()
	if err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountDebitedDrawdownForCustomerTransferPlus test an invalid AccountDebitedDrawdown
func TestInvalidAccountDebitedDrawdownForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)
	err := fwm.checkProhibitedCustomerTransferPlusTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidAccountCreditedDrawdownForCustomerTransferPlus test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)
	creditDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(creditDD)
	err := fwm.checkProhibitedCustomerTransferPlusTags()
	if err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

/*// Beneficiary
ben := mockBeneficiary()
fwm.SetBeneficiary(ben)

// Originator
o := mockOriginator()
fwm.SetOriginator(o)*/
