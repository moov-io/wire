package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
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
		t.Fatalf("%T: %s", err, err)
	}

	// Validate File
	err := file.Validate()

	require.NotNil(t, err)
	expected := NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount, "SubTypeCode",
		fwm.TypeSubType.SubTypeCode).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_previousMessageIdentifierInvalid(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.TypeSubType.SubTypeCode = ReversalTransfer
	fwm.PreviousMessageIdentifier = nil // required when SubTypeCode is ReversalTransfer

	err := fwm.checkPreviousMessageIdentifier()

	require.NotNil(t, err)
	expected := fieldError("PreviousMessageIdentifier", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_invalidLocalInstrument(t *testing.T) {
	fwm := mockCustomerTransferData()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	fwm.BusinessFunctionCode.BusinessFunctionCode = BankTransfer // local instrument only permitted for CTP

	err := fwm.validateLocalInstrumentCode()

	require.NotNil(t, err)
	expected := fieldError("LocalInstrument", ErrLocalInstrumentNotPermitted).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"Charges", fwm.Charges.String()).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"Instructed Amount", fwm.InstructedAmount.String()).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateExchangeRate_missingInstructedAmount(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateExchangeRate()

	require.NotNil(t, err)
	expected := fieldError("InstructedAmount", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
		"ExchangeRate", fwm.ExchangeRate.ExchangeRate).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateBeneficiaryIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateBeneficiaryIntermediaryFI()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateBeneficiaryIntermediaryFI()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateBeneficiaryFI()

	require.NotNil(t, err)
	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateOriginatorFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateOriginatorFI()

	require.NotNil(t, err)
	expected := fieldError("Originator", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateOriginatorFI()

	require.NotNil(t, err)
	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateInstructingFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateInstructingFI()

	require.NotNil(t, err)
	expected := fieldError("Originator", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateInstructingFI()

	require.NotNil(t, err)
	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestNewFEDWireMessage_validateOriginatorToBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateOriginatorToBeneficiary()

	require.NotNil(t, err)
	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)

	// Originator required Field check
	err = fwm.validateOriginatorToBeneficiary()

	require.NotNil(t, err)
	expected = fieldError("Originator", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	o := mockOriginator()
	fwm.SetOriginator(o)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()

	require.NotNil(t, err)
	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	// check beneficiary still required
	fwm.SetBeneficiary(nil)

	err = fwm.validateOriginatorToBeneficiary()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFI()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFI()

	require.NotNil(t, err)
	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFI()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIIntermediaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFIAdvice()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	require.NotNil(t, err)
	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFI()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFI()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIBeneficiaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFIAdvice()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFIAdvice()

	require.NotNil(t, err)
	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiary()

	require.NotNil(t, err)
	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateFIBeneficiaryAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiaryAdvice()

	require.NotNil(t, err)
	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.String(),
		"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("RelatedRemittance", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("RemittanceOriginator", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("RemittanceBeneficiary", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())

	fwm.RemittanceBeneficiary = nil
	fwm.LocalInstrument.LocalInstrumentCode = RemittanceInformationStructured

	// RemittanceBeneficiary Invalid Property
	err = fwm.validateRemittanceBeneficiary()

	require.NotNil(t, err)
	expected = fieldError("RemittanceBeneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("PrimaryRemittanceDocument", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateActualAmountPaid(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)

	// ActualAmountPaid only permitted for CTP and RMTS
	err := fwm.validateActualAmountPaid()

	require.NotNil(t, err)
	expected := fieldError("ActualAmountPaid", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())

}

func TestFEDWireMessage_validateGrossAmountRemittanceDocument(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	gard := mockGrossAmountRemittanceDocument()
	fwm.SetGrossAmountRemittanceDocument(gard)

	// GrossAmountRemittanceDocument only permitted for CTP and RMTS
	err := fwm.validateGrossAmountRemittanceDocument()

	require.NotNil(t, err)
	expected := fieldError("GrossAmountRemittanceDocument", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateAdjustment(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)

	// Adjustment Invalid Property
	err := fwm.validateAdjustment()

	require.NotNil(t, err)
	expected := fieldError("Adjustment", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())

}

func TestFEDWireMessage_validateDateRemittanceDocument(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	drd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(drd)

	// DateRemittanceDocument Invalid Property
	err := fwm.validateDateRemittanceDocument()

	require.NotNil(t, err)
	expected := fieldError("DateRemittanceDocument", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_validateSecondaryRemittanceDocument(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	srd := mockSecondaryRemittanceDocument()
	fwm.SetSecondaryRemittanceDocument(srd)

	// SecondaryRemittanceDocument Invalid Property
	err := fwm.validateSecondaryRemittanceDocument()

	require.NotNil(t, err)
	expected := fieldError("SecondaryRemittanceDocument", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
}

func TestFEDWireMessage_isRemittanceFreeTextValid(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	li := NewLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.SetLocalInstrument(li)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)

	// RemittanceFreeTextValid Invalid Property
	err := fwm.validateRemittanceFreeText()

	require.NotNil(t, err)
	expected := fieldError("RemittanceFreeText", ErrNotPermitted).Error()
	require.Equal(t, expected, err.Error())
}

// TestFEDWireMessage_validateBankTransfer test an invalid BankTransfer
func TestFEDWireMessage_validateBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	tst := mockTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = RequestCredit
	fwm.SetTypeSubType(tst)

	err := fwm.validateBankTransfer()

	require.NotNil(t, err)
	expected := NewErrBusinessFunctionCodeProperty("TypeSubType", tst.TypeCode+tst.SubTypeCode,
		fwm.BusinessFunctionCode.BusinessFunctionCode).Error()
	require.Equal(t, expected, err.Error())
}

// TestFEDWireMessage_invalidTransTypeCodeBankTransfer test an invalid TransactionTypeCode
func TestFEDWireMessage_invalidTransTypeCodeBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkProhibitedBankTransferTags()

	require.NotNil(t, err)
	expected := fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode,
		fwm.BusinessFunctionCode.TransactionTypeCode).Error()
	require.Equal(t, expected, err.Error())
}

// TestInvalidLocalInstrumentForBankTransfer test an invalid LocalInstrument
func TestInvalidLocalInstrumentForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)

	err := fwm.checkProhibitedBankTransferTags()

	require.NotNil(t, err)
	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("Charges", ErrInvalidProperty, fwm.Charges).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty,
		fwm.Beneficiary.Personal.IdentificationCode).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.Equal(t, expected, err.Error())
}

// TestInvalidOriginatorIdentificationCodeForBankTransfer test an invalid Originator Personal.IdentificationCode
func TestInvalidOriginatorIdentificationCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.SetBusinessFunctionCode(bfc)
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetOriginator(o)

	err := fwm.checkProhibitedBankTransferTags()

	require.NotNil(t, err)
	expected := fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty,
		fwm.Originator.Personal.IdentificationCode).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error()
	require.Equal(t, expected, err.Error())
}

// TestTransactionTypeCodeForCustomerTransfer test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkProhibitedCustomerTransferTags()

	require.NotNil(t, err)
	expected := fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode,
		fwm.BusinessFunctionCode.TransactionTypeCode).Error()
	require.Equal(t, expected, err.Error())
}

// TestInvalidLocalInstrumentForCustomerTransfer test an invalid LocalInstrument
func TestInvalidLocalInstrumentForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.SetBusinessFunctionCode(bfc)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)

	err := fwm.checkProhibitedCustomerTransferTags()

	require.NotNil(t, err)
	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("UnstructuredAddenda", ErrInvalidProperty, fwm.UnstructuredAddenda).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("UnstructuredAddenda", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("RelatedRemittance", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryReference", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("OrderingCustomer", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("BeneficiaryCustomer", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("ProprietaryCode", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("RemittanceOriginator", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("RemittanceBeneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("PrimaryRemittanceDocument", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("ActualAmountPaid", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

// TestBeneficiaryForCustomerTransferPlus tests a Beneficiary is required
func TestBeneficiaryIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	require.NotNil(t, err)
	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("Originator OR OriginatorOptionF", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.Equal(t, expected, err.Error())
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

	require.NotNil(t, err)
	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.Equal(t, expected, err.Error())
}
