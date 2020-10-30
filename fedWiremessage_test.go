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

	expected := NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount, "SubTypeCode", fwm.TypeSubType.SubTypeCode).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_previousMessageIdentifierInvalid(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.TypeSubType.SubTypeCode = ReversalTransfer
	fwm.PreviousMessageIdentifier = nil // required when SubTypeCode is ReversalTransfer

	err := fwm.checkPreviousMessageIdentifier()

	expected := fieldError("PreviousMessageIdentifier", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_invalidLocalInstrument(t *testing.T) {
	fwm := mockCustomerTransferData()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.SetLocalInstrument(li)
	fwm.BusinessFunctionCode.BusinessFunctionCode = BankTransfer // local instrument only permitted for CTP

	err := fwm.validateLocalInstrumentCode()

	expected := fieldError("LocalInstrument", ErrLocalInstrumentNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode, "Charges", fwm.Charges.String()).Error()
	require.EqualError(t, err, expected)
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

	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode, "Instructed Amount", fwm.InstructedAmount.String()).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateExchangeRate_missingInstructedAmount(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateExchangeRate()

	expected := fieldError("InstructedAmount", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode, "ExchangeRate", fwm.ExchangeRate.ExchangeRate).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateBeneficiaryIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateBeneficiaryIntermediaryFI()

	require.EqualError(t, err, fieldError("BeneficiaryFI", ErrFieldRequired).Error())

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateBeneficiaryIntermediaryFI()

	require.EqualError(t, err, fieldError("Beneficiary", ErrFieldRequired).Error())
}

func TestFEDWireMessage_validateBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateBeneficiaryFI()

	require.EqualError(t, err, fieldError("Beneficiary", ErrFieldRequired).Error())
}

func TestFEDWireMessage_validateOriginatorFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateOriginatorFI()

	require.EqualError(t, err, fieldError("Originator", ErrFieldRequired).Error())

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateOriginatorFI()

	require.EqualError(t, err, fieldError("OriginatorOptionF", ErrFieldRequired).Error())
}

func TestFEDWireMessage_validateInstructingFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateInstructingFI()

	expected := fieldError("Originator", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	o := mockOriginator()
	fwm.SetOriginator(o)

	// OriginatorOptionF required field check
	err = fwm.validateInstructingFI()

	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestNewFEDWireMessage_validateOriginatorToBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateOriginatorToBeneficiary()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)

	// Originator required Field check
	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("Originator", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	o := mockOriginator()
	fwm.SetOriginator(o)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	// check beneficiary still required
	fwm.SetBeneficiary(nil)

	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFI()

	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFI()

	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFI()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIIntermediaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFIAdvice()

	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFI()

	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFI()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFIAdvice()

	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFIAdvice()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiary()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiaryAdvice()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.String(),
		"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RelatedRemittance", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RemittanceOriginator", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RemittanceBeneficiary", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)

	fwm.RemittanceBeneficiary = nil
	fwm.LocalInstrument.LocalInstrumentCode = RemittanceInformationStructured

	// RemittanceBeneficiary Invalid Property
	err = fwm.validateRemittanceBeneficiary()

	expected = fieldError("RemittanceBeneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("PrimaryRemittanceDocument", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ActualAmountPaid", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)

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

	expected := fieldError("GrossAmountRemittanceDocument", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("Adjustment", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)

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

	expected := fieldError("DateRemittanceDocument", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("SecondaryRemittanceDocument", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RemittanceFreeText", ErrNotPermitted).Error()
	require.EqualError(t, err, expected)
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

	expected := NewErrBusinessFunctionCodeProperty("TypeSubType", tst.TypeCode+tst.SubTypeCode,
		fwm.BusinessFunctionCode.BusinessFunctionCode).Error()
	require.EqualError(t, err, expected)
}

// TestFEDWireMessage_invalidTransTypeCodeBankTransfer test an invalid TransactionTypeCode
func TestFEDWireMessage_invalidTransTypeCodeBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode,
		fwm.BusinessFunctionCode.TransactionTypeCode).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("Charges", ErrInvalidProperty, fwm.Charges).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty,
		fwm.Beneficiary.Personal.IdentificationCode).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty,
		fwm.Originator.Personal.IdentificationCode).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("UnstructuredAddenda", ErrInvalidProperty, fwm.UnstructuredAddenda).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error()
	require.EqualError(t, err, expected)
}

// TestTransactionTypeCodeForCustomerTransfer test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode,
		fwm.BusinessFunctionCode.TransactionTypeCode).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("UnstructuredAddenda", ErrInvalidProperty, fwm.UnstructuredAddenda).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("UnstructuredAddenda", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RelatedRemittance", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("BeneficiaryReference", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("OrderingCustomer", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("BeneficiaryCustomer", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ProprietaryCode", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RemittanceOriginator", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("RemittanceBeneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("PrimaryRemittanceDocument", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("ActualAmountPaid", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryForCustomerTransferPlus tests a Beneficiary is required
func TestBeneficiaryIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("Originator OR OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
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

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
}
