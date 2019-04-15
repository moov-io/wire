package wire

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestFedWireMessage writes an FedWireMessage toa file
func TestFedWireMessageWrite(t *testing.T) {
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

	fd, err := os.Create(filepath.Join("", "test/testdata", "fedWireMessage.txt"))
	if err != nil {
		log.Fatalf("Unexpected error creating output file: %s\n", err)
	}
	defer func() {
		fd.Sync()
		fd.Close()
	}()
	w := NewWriter(fd)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}

	// We want to write the file to an io.Writer
	w = NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

	r := NewReader(strings.NewReader(b.String()))
	_, err = r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

/*// write the file to std out. Anything io.Writer
fd, err := os.Create(filepath.Join("..", "ach-rck-read", "rck-debit.ach"))
if err != nil {
log.Fatalf("Unexpected error creating output file: %s\n", err)
}
defer func() {
	fd.Sync()
	fd.Close()
}()
w := ach.NewWriter(fd)
if err := w.Write(file); err != nil {
log.Fatalf("Unexpected error: %s\n", err)
}
w.Flush()*/
