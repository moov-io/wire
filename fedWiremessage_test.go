package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func mockCustomerTransferData() FEDWireMessage {
	fwm := FEDWireMessage{}

	// Mandatory Fields
	fwm.SenderSupplied = mockSenderSupplied()
	fwm.TypeSubType = mockTypeSubType()
	fwm.InputMessageAccountabilityData = mockInputMessageAccountabilityData()
	fwm.Amount = mockAmount()
	fwm.SenderDepositoryInstitution = mockSenderDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = mockReceiverDepositoryInstitution()
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc
	return fwm
}

func TestFEDWireMessage_invalidAmount(t *testing.T) {
	file := NewFile()
	fwm := mockCustomerTransferData()
	// Override to trigger error (can only be zeros is TypeSubType code is 90)
	fwm.Amount.Amount = "000000000000"
	// Beneficiary
	fwm.Beneficiary = mockBeneficiary()
	// Originator
	fwm.Originator = mockOriginator()
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
	fwm.LocalInstrument = li
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
	fwm.LocalInstrument = li
	fwm.Charges = mockCharges()
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
	fwm.LocalInstrument = li
	fwm.InstructedAmount = mockInstructedAmount()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateInstructedAmount()

	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode, "Instructed Amount", fwm.InstructedAmount.String()).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateExchangeRate_missingInstructedAmount(t *testing.T) {
	fwm := mockCustomerTransferData()
	// Override to trigger error
	fwm.ExchangeRate = mockExchangeRate()
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
	fwm.LocalInstrument = li
	fwm.ExchangeRate = mockExchangeRate()
	fwm.InstructedAmount = mockInstructedAmount()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	err := fwm.validateExchangeRate()

	expected := NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode, "ExchangeRate", fwm.ExchangeRate.ExchangeRate).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateBeneficiaryIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BeneficiaryIntermediaryFI = mockBeneficiaryIntermediaryFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateBeneficiaryIntermediaryFI()

	require.EqualError(t, err, fieldError("BeneficiaryFI", ErrFieldRequired).Error())

	fwm.BeneficiaryFI = mockBeneficiaryFI()

	// Beneficiary required field check
	err = fwm.validateBeneficiaryIntermediaryFI()

	require.EqualError(t, err, fieldError("Beneficiary", ErrFieldRequired).Error())
}

func TestFEDWireMessage_validateBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.BeneficiaryFI = mockBeneficiaryFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateBeneficiaryFI()

	require.EqualError(t, err, fieldError("Beneficiary", ErrFieldRequired).Error())
}

func TestFEDWireMessage_validateOriginatorFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.OriginatorFI = mockOriginatorFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateOriginatorFI()

	require.EqualError(t, err, fieldError("Originator", ErrFieldRequired).Error())

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	err = fwm.validateOriginatorFI()

	require.EqualError(t, err, fieldError("Originator or OriginatorOptionF", ErrFieldRequired).Error())

	fwm.Originator = mockOriginator()
	err = fwm.validateOriginatorFI()

	require.NoError(t, err)

	fwm.Originator = nil
	fwm.OriginatorOptionF = mockOriginatorOptionF()
	err = fwm.validateOriginatorFI()

	require.NoError(t, err)
}

func TestFEDWireMessage_validateInstructingFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.InstructingFI = mockInstructingFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Originator required field check
	err := fwm.validateInstructingFI()

	expected := fieldError("Originator", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus
	fwm.Originator = mockOriginator()

	// OriginatorOptionF required field check
	err = fwm.validateInstructingFI()

	expected = fieldError("OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestNewFEDWireMessage_validateOriginatorToBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.OriginatorToBeneficiary = mockOriginatorToBeneficiary()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateOriginatorToBeneficiary()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.Beneficiary = mockBeneficiary()

	// Originator required Field check
	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("Originator", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransferPlus

	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("Originator or OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.Originator = mockOriginator()

	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()

	require.NoError(t, err)

	fwm.Originator = nil
	fwm.OriginatorOptionF = mockOriginatorOptionF()

	// OriginatorOptionF required Field check
	err = fwm.validateOriginatorToBeneficiary()

	require.NoError(t, err)

	// check beneficiary still required
	fwm.Beneficiary = nil

	err = fwm.validateOriginatorToBeneficiary()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIIntermediaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIIntermediaryFI = mockFIIntermediaryFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFI()

	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryIntermediaryFI = mockBeneficiaryIntermediaryFI()

	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFI()

	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryFI = mockBeneficiaryFI()

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFI()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIIntermediaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIIntermediaryFIAdvice = mockFIIntermediaryFIAdvice()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryIntermediaryFI required field check
	err := fwm.validateFIIntermediaryFIAdvice()

	expected := fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryIntermediaryFI = mockBeneficiaryIntermediaryFI()
	// BeneficiaryFI required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	expected = fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryFI = mockBeneficiaryFI()

	// Beneficiary required field check
	err = fwm.validateFIIntermediaryFIAdvice()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryFI(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIBeneficiaryFI = mockFIBeneficiaryFI()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFI()

	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryFI = mockBeneficiaryFI()

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFI()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryFIAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIBeneficiaryFIAdvice = mockFIBeneficiaryFIAdvice()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// BeneficiaryFI required field check
	err := fwm.validateFIBeneficiaryFIAdvice()

	expected := fieldError("BeneficiaryFI", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)

	fwm.BeneficiaryFI = mockBeneficiaryFI()

	// Beneficiary required field check
	err = fwm.validateFIBeneficiaryFIAdvice()

	expected = fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiary(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIBeneficiary = mockFIBeneficiary()
	fwm.BusinessFunctionCode.BusinessFunctionCode = CustomerTransfer

	// Beneficiary required field check
	err := fwm.validateFIBeneficiary()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

func TestFEDWireMessage_validateFIBeneficiaryAdvice(t *testing.T) {
	fwm := mockCustomerTransferData()
	fwm.FIBeneficiaryAdvice = mockFIBeneficiaryAdvice()
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
	fwm.LocalInstrument = li
	fwm.UnstructuredAddenda = mockUnstructuredAddenda()

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
	fwm.LocalInstrument = li
	fwm.RelatedRemittance = mockRelatedRemittance()

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
	fwm.LocalInstrument = li
	fwm.RemittanceOriginator = mockRemittanceOriginator()

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
	fwm.LocalInstrument = li
	fwm.RemittanceBeneficiary = mockRemittanceBeneficiary()

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
	fwm.LocalInstrument = li
	fwm.PrimaryRemittanceDocument = mockPrimaryRemittanceDocument()

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
	fwm.LocalInstrument = li
	fwm.ActualAmountPaid = mockActualAmountPaid()

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
	fwm.LocalInstrument = li
	fwm.GrossAmountRemittanceDocument = mockGrossAmountRemittanceDocument()

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
	fwm.LocalInstrument = li
	fwm.Adjustment = mockAdjustment()

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
	fwm.LocalInstrument = li
	fwm.DateRemittanceDocument = mockDateRemittanceDocument()

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
	fwm.LocalInstrument = li
	fwm.SecondaryRemittanceDocument = mockSecondaryRemittanceDocument()

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
	fwm.LocalInstrument = li
	fwm.RemittanceFreeText = mockRemittanceFreeText()

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
	fwm.BusinessFunctionCode = bfc
	tst := mockTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = RequestCredit
	fwm.TypeSubType = tst

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
	fwm.BusinessFunctionCode = bfc

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
	fwm.BusinessFunctionCode = bfc
	fwm.LocalInstrument = mockLocalInstrument()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidPaymentNotificationForBankTransfer test an invalid PaymentNotification
func TestInvalidPaymentNotificationForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.PaymentNotification = mockPaymentNotification()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidChargesForBankTransfer test an invalid Charges
func TestInvalidChargesForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.Charges = mockCharges()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("Charges", ErrInvalidProperty, fwm.Charges).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidInstructedAmountForBankTransfer test an invalid InstructedAmount
func TestInvalidInstructedAmountForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.InstructedAmount = mockInstructedAmount()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidExchangeRateForBankTransfer test an invalid ExchangeRate
func TestInvalidExchangeRateForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.ExchangeRate = mockExchangeRate()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidBeneficiaryIdentificationCodeForBankTransfer test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.Beneficiary = ben

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
	fwm.BusinessFunctionCode = bfc
	fwm.AccountDebitedDrawdown = mockAccountDebitedDrawdown()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidOriginatorIdentificationCodeForBankTransfer test an invalid Originator Personal.IdentificationCode
func TestInvalidOriginatorIdentificationCodeForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.Originator = o

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
	fwm.BusinessFunctionCode = bfc
	fwm.OriginatorOptionF = mockOriginatorOptionF()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidAccountCreditedDrawdownForBankTransfer test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.AccountCreditedDrawdown = mockAccountCreditedDrawdown()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidFIDrawdownDebitAccountAdviceForBankTransfer test an invalid FIDrawdownDebitAccountAdvice
func TestInvalidFIDrawdownDebitAccountAdviceForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.FIDrawdownDebitAccountAdvice = mockFIDrawdownDebitAccountAdvice()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidServiceMessageForBankTransfer test an invalid ServiceMessage
func TestInvalidServiceMessageForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.ServiceMessage = mockServiceMessage()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidUnstructuredAddendaForBankTransfer test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.UnstructuredAddenda = mockUnstructuredAddenda()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("UnstructuredAddenda", ErrInvalidProperty, fwm.UnstructuredAddenda).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidCurrencyInstructedAmountForBankTransfer test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.CurrencyInstructedAmount = mockCurrencyInstructedAmount()

	err := fwm.checkProhibitedBankTransferTags()

	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidRelatedRemittanceForBankTransfer test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForBankTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.RelatedRemittance = mockRelatedRemittance()

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
	fwm.BusinessFunctionCode = bfc

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
	fwm.BusinessFunctionCode = bfc
	fwm.LocalInstrument = mockLocalInstrument()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidPaymentNotificationForCustomerTransfer test an invalid PaymentNotification
func TestInvalidPaymentNotificationForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.PaymentNotification = mockPaymentNotification()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidAccountDebitedDrawdownForCustomerTransfer test an invalid AccountDebitedDrawdown
func TestInvalidAccountDebitedDrawdownForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.AccountDebitedDrawdown = mockAccountDebitedDrawdown()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidOriginatorOptionFForCustomerTransfer test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.OriginatorOptionF = mockOriginatorOptionF()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidAccountCreditedDrawdownForCustomerTransfer test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.AccountCreditedDrawdown = mockAccountCreditedDrawdown()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidFIDrawdownDebitAccountAdviceForCustomerTransfer test an invalid FIDrawdownDebitAccountAdvice
func TestInvalidFIDrawdownDebitAccountAdviceForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.FIDrawdownDebitAccountAdvice = mockFIDrawdownDebitAccountAdvice()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidServiceMessageForCustomerTransfer test an invalid ServiceMessage
func TestInvalidServiceMessageForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.ServiceMessage = mockServiceMessage()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("ServiceMessage", ErrInvalidProperty, fwm.ServiceMessage).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidUnstructuredAddendaForCustomerTransfer test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.UnstructuredAddenda = mockUnstructuredAddenda()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("UnstructuredAddenda", ErrInvalidProperty, fwm.UnstructuredAddenda).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidCurrencyInstructedAmountForCustomerTransfer test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.CurrencyInstructedAmount = mockCurrencyInstructedAmount()

	err := fwm.checkProhibitedCustomerTransferTags()

	expected := fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidRelatedRemittanceForCustomerTransfer test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForCustomerTransfer(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	fwm.BusinessFunctionCode = bfc
	fwm.RelatedRemittance = mockRelatedRemittance()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	fwm.LocalInstrument = mockLocalInstrument()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.LocalInstrument = li

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.LocalInstrument = li

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.LocalInstrument = li
	fwm.BeneficiaryReference = mockBeneficiaryReference()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.LocalInstrument = li
	fwm.BeneficiaryReference = mockBeneficiaryReference()
	fwm.OrderingCustomer = mockOrderingCustomer()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	fwm.LocalInstrument = li

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.LocalInstrument = li

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.LocalInstrument = li
	fwm.RemittanceOriginator = mockRemittanceOriginator()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.LocalInstrument = li
	fwm.RemittanceOriginator = mockRemittanceOriginator()
	fwm.RemittanceBeneficiary = mockRemittanceBeneficiary()

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
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	li := mockLocalInstrument()
	li.LocalInstrumentCode = RemittanceInformationStructured
	fwm.LocalInstrument = li
	fwm.RemittanceOriginator = mockRemittanceOriginator()
	fwm.RemittanceBeneficiary = mockRemittanceBeneficiary()
	fwm.PrimaryRemittanceDocument = mockPrimaryRemittanceDocument()

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	expected := fieldError("ActualAmountPaid", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryForCustomerTransferPlus tests a Beneficiary is required
func TestBeneficiaryIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.BusinessFunctionCode = bfc

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	expected := fieldError("Beneficiary", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

// TestOriginatorForCustomerTransferPlus tests an Originator is required
func TestOriginatorIdentificationCodeForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.BusinessFunctionCode = bfc
	fwm.Beneficiary = mockBeneficiary()

	err := fwm.checkMandatoryCustomerTransferPlusTags()

	expected := fieldError("Originator OR OriginatorOptionF", ErrFieldRequired).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidAccountDebitedDrawdownForCustomerTransferPlus test an invalid AccountDebitedDrawdown
func TestInvalidAccountDebitedDrawdownForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.BusinessFunctionCode = bfc
	fwm.AccountDebitedDrawdown = mockAccountDebitedDrawdown()

	err := fwm.checkProhibitedCustomerTransferPlusTags()

	expected := fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown).Error()
	require.EqualError(t, err, expected)
}

// TestInvalidAccountCreditedDrawdownForCustomerTransferPlus test an invalid AccountCreditedDrawdown
func TestInvalidAccountCreditedDrawdownForCustomerTransferPlus(t *testing.T) {
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	fwm.BusinessFunctionCode = bfc
	fwm.AccountCreditedDrawdown = mockAccountCreditedDrawdown()

	err := fwm.checkProhibitedCustomerTransferPlusTags()

	expected := fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown).Error()
	require.EqualError(t, err, expected)
}
