package wire

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSenderSupplied_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("SenderSupplied", ErrFieldRequired).Error())
}

func TestTypeSubType_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("TypeSubType", ErrFieldRequired).Error())
}

func TestInputMessageAccountabilityData_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	amt := mockAmount()
	fwm.Amount = amt
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("InputMessageAccountabilityData", ErrFieldRequired).Error())
}

func TestAmount_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

func TestSenderDepositoryInstitution_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("SenderDepositoryInstitution", ErrFieldRequired).Error())
}

func TestReceiverDepositoryInstitution_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("ReceiverDepositoryInstitution", ErrFieldRequired).Error())
}

func TestBusinessFunctionCode_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi

	file.AddFEDWireMessage(fwm)

	// Create file
	require.NoError(t, file.Create())

	err := file.Validate()

	require.EqualError(t, err, fieldError("BusinessFunctionCode", ErrFieldRequired).Error())
}

// TestFEDWireMessageWriteBankTransfer writes a FEDWireMessage to a file with BusinessFunctionCode = BTR
func TestFEDWireMessageWriteBankTransfer(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = "BTR"
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerTransfer writes a FEDWireMessage to a file with BusinessFunctionCode = CTR
func TestFEDWireMessageWriteCustomerTransfer(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi
	//li := mockLocalInstrument()
	//fwm.LocalInstrument = li
	c := mockCharges()
	fwm.Charges = c
	ia := mockInstructedAmount()
	fwm.InstructedAmount = ia
	eRate := mockExchangeRate()
	fwm.ExchangeRate = eRate

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerTransferPlus writes a FEDWireMessage to a file with BusinessFunctionCode = CTP
func TestFEDWireMessageWriteCustomerTransferPlus(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	li.ProprietaryCode = "PROP CODE"
	fwm.LocalInstrument = li
	pn := mockPaymentNotification()
	fwm.PaymentNotification = pn
	c := mockCharges()
	fwm.Charges = c
	ia := mockInstructedAmount()
	fwm.InstructedAmount = ia
	eRate := mockExchangeRate()
	fwm.ExchangeRate = eRate

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	oof := mockOriginatorOptionF()
	fwm.OriginatorOptionF = oof
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	// ServiceMessage
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCheckSameDaySettlement writes a FEDWireMessage to a file with BusinessFunctionCode = CKS
func TestFEDWireMessageWriteCheckSameDaySettlement(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = SettlementTransfer
	tst.SubTypeCode = BasicFundsTransfer
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CheckSameDaySettlement
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteDepositSendersAccount writes a FEDWireMessage to a file with BusinessFunctionCode = DEP
func TestFEDWireMessageWriteDepositSendersAccount(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = SettlementTransfer
	tst.SubTypeCode = BasicFundsTransfer
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = DepositSendersAccount
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteFEDFundsReturned writes a FEDWireMessage to a file with BusinessFunctionCode = FFR
func TestFEDWireMessageWriteFEDFundsReturned(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = "16"
	tst.SubTypeCode = "00"
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = FEDFundsReturned
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteFEDFundsSold writes a FEDWireMessage to a file with BusinessFunctionCode = FFS
func TestFEDWireMessageWriteFEDFundsSold(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = SettlementTransfer
	tst.SubTypeCode = BasicFundsTransfer
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = FEDFundsSold
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteDrawdownRequest writes a FEDWireMessage to a file with BusinessFunctionCode = DRW
func TestFEDWireMessageWriteDrawdownRequest(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = FundsTransferRequestCredit
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = DrawdownResponse
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteBankDrawdownRequest writes a FEDWireMessage to a file with BusinessFunctionCode = DRB
func TestFEDWireMessageWriteBankDrawdownRequest(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = SettlementTransfer
	tst.SubTypeCode = RequestCredit
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankDrawDownRequest
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br
	debitDD := mockAccountDebitedDrawdown()
	fwm.AccountDebitedDrawdown = debitDD

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	creditDD := mockAccountCreditedDrawdown()
	fwm.AccountCreditedDrawdown = creditDD
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerCorporateDrawdownRequest writes a FEDWireMessage to a file with BusinessFunctionCode = DRC
func TestFEDWireMessageWriteCustomerCorporateDrawdownRequest(t *testing.T) {
	file := NewFile()
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = RequestCredit
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerCorporateDrawdownRequest
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br
	debitDD := mockAccountDebitedDrawdown()
	fwm.AccountDebitedDrawdown = debitDD

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	creditDD := mockAccountCreditedDrawdown()
	fwm.AccountCreditedDrawdown = creditDD
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	fwm.FIDrawdownDebitAccountAdvice = debitDDAdvice
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteServiceMessage writes a FEDWireMessage to a file with BusinessFunctionCode = SVC
func TestFEDWireMessageWriteServiceMessage(t *testing.T) {
	file := NewFile()
	fwm := createMockServiceMessageData()
	fwm.TypeSubType.TypeCode = FundsTransfer
	fwm.TypeSubType.SubTypeCode = RequestReversal

	fwm.BusinessFunctionCode.BusinessFunctionCode = BFCServiceMessage
	fwm.BusinessFunctionCode.TransactionTypeCode = "   "

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// writeFile writes a FEDWireMessage File and ensures the File can be read
func writeFile(file *File) error {
	if err := file.Create(); err != nil {
		return err
	}
	if err := file.Validate(); err != nil {
		return err
	}
	b := &bytes.Buffer{}
	f := NewWriter(b)
	if err := f.Write(file); err != nil {
		return err
	}

	r := NewReader(strings.NewReader(b.String()))
	fwmFile, err := r.Read()
	if err != nil {
		return err
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		return err
	}
	return nil
}

func createMockServiceMessageData() FEDWireMessage {
	fwm := FEDWireMessage{}
	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = RequestReversal
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = BFCServiceMessage
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br
	debitDD := mockAccountDebitedDrawdown()
	fwm.AccountDebitedDrawdown = debitDD

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	creditDD := mockAccountCreditedDrawdown()
	fwm.AccountCreditedDrawdown = creditDD
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.FIReceiverFI = firfi
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	fwm.FIDrawdownDebitAccountAdvice = debitDDAdvice
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	// ServiceMessage
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	return fwm
}

func createCustomerTransferData() FEDWireMessage {
	fwm := FEDWireMessage{}

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SenderSupplied = ss
	tst := mockTypeSubType()
	fwm.TypeSubType = tst
	imad := mockInputMessageAccountabilityData()
	fwm.InputMessageAccountabilityData = imad
	amt := mockAmount()
	fwm.Amount = amt
	sdi := mockSenderDepositoryInstitution()
	fwm.SenderDepositoryInstitution = sdi
	rdi := mockReceiverDepositoryInstitution()
	fwm.ReceiverDepositoryInstitution = rdi
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransferPlus
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SenderReference = sr
	pmi := mockPreviousMessageIdentifier()
	fwm.PreviousMessageIdentifier = pmi
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	li.ProprietaryCode = "PROP CODE"
	fwm.LocalInstrument = li
	pn := mockPaymentNotification()
	fwm.PaymentNotification = pn
	/*	c := mockCharges()
		fwm.Charges = c
		ia := mockInstructedAmount()
		fwm.InstructedAmount = ia
		eRate := mockExchangeRate()
		fwm.ExchangeRate = eRate*/

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi
	bfi := mockBeneficiaryFI()
	fwm.BeneficiaryFI = bfi
	ben := mockBeneficiary()
	fwm.Beneficiary = ben
	br := mockBeneficiaryReference()
	fwm.BeneficiaryReference = br

	// Originator
	o := mockOriginator()
	fwm.Originator = o
	oof := mockOriginatorOptionF()
	fwm.OriginatorOptionF = oof
	ofi := mockOriginatorFI()
	fwm.OriginatorFI = ofi
	ifi := mockInstructingFI()
	fwm.InstructingFI = ifi
	ob := mockOriginatorToBeneficiary()
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	fiifi := mockFIIntermediaryFI()
	fwm.FIIntermediaryFI = fiifi
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.FIIntermediaryFIAdvice = fiifia
	fibfi := mockFIBeneficiaryFI()
	fwm.FIBeneficiaryFI = fibfi
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.FIBeneficiaryFIAdvice = fibfia
	fib := mockFIBeneficiary()
	fwm.FIBeneficiary = fib
	fiba := mockFIBeneficiaryAdvice()
	fwm.FIBeneficiaryAdvice = fiba
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.FIPaymentMethodToBeneficiary = pm
	fifi := mockFIAdditionalFIToFI()
	fwm.FIAdditionalFIToFI = fifi

	return fwm
}

// TestFEDWireMessageWriteCustomerTransferPlusCOVS writes a FEDWireMessage to a file with BusinessFunctionCode = CTP and
// LocalInstrumentCode = "COVS"
func TestFEDWireMessageWriteCustomerTransferPlusCOVS(t *testing.T) {
	file := NewFile()
	fwm := createCustomerTransferData()

	fwm.LocalInstrument.LocalInstrumentCode = SequenceBCoverPaymentStructured
	fwm.LocalInstrument.ProprietaryCode = ""

	// Cover Payment Information
	cia := mockCurrencyInstructedAmount()
	fwm.CurrencyInstructedAmount = cia
	oc := mockOrderingCustomer()
	fwm.OrderingCustomer = oc
	oi := mockOrderingInstitution()
	fwm.OrderingInstitution = oi
	ii := mockIntermediaryInstitution()
	fwm.IntermediaryInstitution = ii
	iAccount := mockInstitutionAccount()
	fwm.InstitutionAccount = iAccount
	bc := mockBeneficiaryCustomer()
	fwm.BeneficiaryCustomer = bc
	ri := mockRemittance()
	fwm.Remittance = ri
	str := mockSenderToReceiver()
	fwm.SenderToReceiver = str

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerTransferPlusRelatedRemittance writes a FEDWireMessage to a file with BusinessFunctionCode = CTP and
// LocalInstrumentCode = "RRMT"
func TestFEDWireMessageWriteCustomerTransferPlusRelatedRemittance(t *testing.T) {
	file := NewFile()
	fwm := createCustomerTransferData()

	fwm.LocalInstrument.LocalInstrumentCode = RelatedRemittanceInformation
	fwm.LocalInstrument.ProprietaryCode = ""

	// Related Remittance Information
	rr := mockRelatedRemittance()
	fwm.RelatedRemittance = rr

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerTransferPlusRemittanceInformationStructured writes a FEDWireMessage to a file with BusinessFunctionCode = CTP and
// LocalInstrumentCode = "RMTS"
func TestFEDWireMessageWriteCustomerTransferPlusRemittanceInformationStructured(t *testing.T) {
	file := NewFile()
	fwm := createCustomerTransferData()

	fwm.LocalInstrument.LocalInstrumentCode = RemittanceInformationStructured
	fwm.LocalInstrument.ProprietaryCode = ""

	// Structured Remittance Information
	ro := mockRemittanceOriginator()
	fwm.RemittanceOriginator = ro
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.DateBirthPlace = ""
	fwm.RemittanceBeneficiary = rb

	// Additional Remittance Data
	prd := mockPrimaryRemittanceDocument()
	fwm.PrimaryRemittanceDocument = prd
	aap := mockActualAmountPaid()
	fwm.ActualAmountPaid = aap
	gard := mockGrossAmountRemittanceDocument()
	fwm.GrossAmountRemittanceDocument = gard
	nd := mockAmountNegotiatedDiscount()
	fwm.AmountNegotiatedDiscount = nd
	adj := mockAdjustment()
	fwm.Adjustment = adj
	drd := mockDateRemittanceDocument()
	fwm.DateRemittanceDocument = drd
	srd := mockSecondaryRemittanceDocument()
	fwm.SecondaryRemittanceDocument = srd
	rft := mockRemittanceFreeText()
	fwm.RemittanceFreeText = rft

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}

// TestFEDWireMessageWriteCustomerTransferPlusUnstructuredAddenda writes a FEDWireMessage to a file with BusinessFunctionCode = CTP and
// LocalInstrumentCode = "ANSI"
func TestFEDWireMessageWriteCustomerTransferPlusUnstructuredAddenda(t *testing.T) {
	file := NewFile()
	fwm := createCustomerTransferData()

	fwm.LocalInstrument.LocalInstrumentCode = ANSIX12format
	fwm.LocalInstrument.ProprietaryCode = ""

	// Unstructured Addenda
	ua := mockUnstructuredAddenda()
	fwm.UnstructuredAddenda = ua

	file.AddFEDWireMessage(fwm)

	require.NoError(t, writeFile(file))
}
