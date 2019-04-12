package wire

import (
	"bytes"
	"log"
	"os"
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
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}