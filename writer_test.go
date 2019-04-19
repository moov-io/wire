package wire

import (
	"bytes"
	"github.com/moov-io/base"
	"log"
	"os"
	"strings"
	"testing"
)

/*// TestFedWireMessageWriteCustomerTransfer writes a FedWireMessage to a file with BusinessFunctionCode = CTR
func TestFedWireMessageWriteCustomerTransfer(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	c := mockCharges()
	fwm.SetCharges(c)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)

	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	oof := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(oof)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	creditDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(creditDD)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.SetFIReceiverFI(firfi)
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	fwm.SetFIDrawdownDebitAccountAdvice(debitDDAdvice)
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)

	// Cover Payment Information
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	oc := mockOrderingCustomer()
	fwm.SetOrderingCustomer(oc)
	oi := mockOrderingInstitution()
	fwm.SetOrderingInstitution(oi)
	ii := mockIntermediaryInstitution()
	fwm.SetIntermediaryInstitution(ii)
	iAccount := mockInstitutionAccount()
	fwm.SetInstitutionAccount(iAccount)
	bc := mockBeneficiaryCustomer()
	fwm.SetBeneficiaryCustomer(bc)
	ri := mockRemittance()
	fwm.SetRemittance(ri)
	str := mockSenderToReceiver()
	fwm.SetSenderToReceiver(str)

	// Unstructured Addenda
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)

	// Related Remittance Information
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)

	// Structured Remittance Information
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)

	// Additional Remittance Data
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)
	gard := mockGrossAmountRemittanceDocument()
	fwm.SetGrossAmountRemittanceDocument(gard)
	nd := mockAmountNegotiatedDiscount()
	fwm.SetAmountNegotiatedDiscount(nd)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)
	drd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(drd)
	srd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(srd)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)

	// ServiceMessage
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// We want to write the file to an io.Writer
	w := NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))

	fwmFile, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}*/

func TestSenderSupplied_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

	// Mandatory Fields
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	amt := mockAmount()
	fwm.SetAmount(amt)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestTypeSubType_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SetSenderSupplied(ss)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	amt := mockAmount()
	fwm.SetAmount(amt)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestInputMessageAccountabilityData_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SetSenderSupplied(ss)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	amt := mockAmount()
	fwm.SetAmount(amt)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAmount_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SetSenderSupplied(ss)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	amt := mockAmount()
	fwm.SetAmount(amt)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestSenderDepositoryInstitution_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

	// Mandatory Fields
	ss := mockSenderSupplied()
	fwm.SetSenderSupplied(ss)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	amt := mockAmount()
	fwm.SetAmount(amt)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestReceiverDepositoryInstitution_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBusinessFunctionCode_Mandatory(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFedWireMessageWriteBankTransfer writes a FedWireMessage to a file with BusinessFunctionCode = BTR
func TestFedWireMessageWriteBankTransfer(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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
	bfc.BusinessFunctionCode = "BTR"
	bfc.TransactionTypeCode = "   "
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)


	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.SetFIReceiverFI(firfi)
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// We want to write the file to an io.Writer
	w := NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))

	fwmFile, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFedWireMessageWriteCustomerTransfer writes a FedWireMessage to a file with BusinessFunctionCode = CTR
func TestFedWireMessageWriteCustomerTransfer(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)
	//li := mockLocalInstrument()
	//fwm.SetLocalInstrument(li)
	c := mockCharges()
	fwm.SetCharges(c)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)

	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	firfi := mockFIReceiverFI()
	fwm.SetFIReceiverFI(firfi)
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// We want to write the file to an io.Writer
	w := NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))

	fwmFile, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFedWireMessageWriteCustomerTransferPlus writes a FedWireMessage to a file with BusinessFunctionCode = CTP
func TestFedWireMessageWriteCustomerTransferPlus(t *testing.T) {
	file := NewFile()
	fwm := NewFedWireMessage()

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
	bfc.BusinessFunctionCode = CustomerTransferPlus
	bfc.TransactionTypeCode = "   "
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	c := mockCharges()
	fwm.SetCharges(c)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)

	// Beneficiary
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	bfi := mockBeneficiaryFI()
	fwm.SetBeneficiaryFI(bfi)
	ben := mockBeneficiary()
	fwm.SetBeneficiary(ben)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	//debitDD := mockAccountDebitedDrawdown()
	//fwm.SetAccountDebitedDrawdown(debitDD)

	// Originator
	o := mockOriginator()
	fwm.SetOriginator(o)
	oof := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(oof)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	ifi := mockInstructingFI()
	fwm.SetInstructingFI(ifi)
	//creditDD := mockAccountCreditedDrawdown()
	//fwm.SetAccountCreditedDrawdown(creditDD)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	//firfi := mockFIReceiverFI()
	//fwm.SetFIReceiverFI(firfi)
	//debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	//fwm.SetFIDrawdownDebitAccountAdvice(debitDDAdvice)
	fiifi := mockFIIntermediaryFI()
	fwm.SetFIIntermediaryFI(fiifi)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	fibfia := mockFIBeneficiaryFIAdvice()
	fwm.SetFIBeneficiaryFIAdvice(fibfia)
	fib := mockFIBeneficiary()
	fwm.SetFIBeneficiary(fib)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	pm := mockFIPaymentMethodToBeneficiary()
	fwm.SetFIPaymentMethodToBeneficiary(pm)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)

	// Cover Payment Information
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	oc := mockOrderingCustomer()
	fwm.SetOrderingCustomer(oc)
	oi := mockOrderingInstitution()
	fwm.SetOrderingInstitution(oi)
	ii := mockIntermediaryInstitution()
	fwm.SetIntermediaryInstitution(ii)
	iAccount := mockInstitutionAccount()
	fwm.SetInstitutionAccount(iAccount)
	bc := mockBeneficiaryCustomer()
	fwm.SetBeneficiaryCustomer(bc)
	ri := mockRemittance()
	fwm.SetRemittance(ri)
	str := mockSenderToReceiver()
	fwm.SetSenderToReceiver(str)

	// Unstructured Addenda
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)

	// Related Remittance Information
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)

	// Structured Remittance Information
	ro := mockRemittanceOriginator()
	fwm.SetRemittanceOriginator(ro)
	rb := mockRemittanceBeneficiary()
	fwm.SetRemittanceBeneficiary(rb)

	// Additional Remittance Data
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)
	gard := mockGrossAmountRemittanceDocument()
	fwm.SetGrossAmountRemittanceDocument(gard)
	nd := mockAmountNegotiatedDiscount()
	fwm.SetAmountNegotiatedDiscount(nd)
	adj := mockAdjustment()
	fwm.SetAdjustment(adj)
	drd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(drd)
	srd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(srd)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)

	// ServiceMessage
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)

	file.AddFedWireMessage(fwm)

	// Create file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// We want to write the file to an io.Writer
	w := NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))

	fwmFile, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}