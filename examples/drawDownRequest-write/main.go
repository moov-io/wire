// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/wire"
)

func main() {
	file := wire.NewFile()
	fwm := wire.NewFEDWireMessage()
	// Mandatory Fields
	ss := wire.NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = wire.MessageDuplicationOriginal
	fwm.SetSenderSupplied(ss)

	tst := wire.NewTypeSubType()
	tst.TypeCode = wire.SettlementTransfer
	tst.SubTypeCode = wire.RequestCredit

	fwm.SetTypeSubType(tst)

	imad := wire.NewInputMessageAccountabilityData()
	imad.InputCycleDate = time.Now().Format("20060102")
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	fwm.SetInputMessageAccountabilityData(imad)

	amt := wire.NewAmount()
	amt.Amount = "000001234567"
	fwm.SetAmount(amt)

	sdi := wire.NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	fwm.SetSenderDepositoryInstitution(sdi)

	rdi := wire.NewReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	fwm.SetReceiverDepositoryInstitution(rdi)

	bfc := wire.NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = wire.BankDrawDownRequest
	bfc.TransactionTypeCode = "   "
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := wire.NewSenderReference()
	sr.SenderReference = "Sender Reference"
	fwm.SetSenderReference(sr)

	pmi := wire.NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	fwm.SetPreviousMessageIdentifier(pmi)

	// Beneficiary
	bifi := wire.NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "DrawDown FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "DrawDown Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "DrawDown Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "DrawDown Address Three"
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	bfi := wire.NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "DrawDown FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "DrawDown Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "DrawDown Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "DrawDown Address Three"
	fwm.SetBeneficiaryFI(bfi)

	ben := wire.NewBeneficiary()
	ben.Personal.IdentificationCode = wire.DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "DrawDown Name"
	ben.Personal.Address.AddressLineOne = "DrawDown Address One"
	ben.Personal.Address.AddressLineTwo = "DrawDown Address Two"
	ben.Personal.Address.AddressLineThree = "DrawDown Address Three"
	fwm.SetBeneficiary(ben)

	br := wire.NewBeneficiaryReference()
	br.BeneficiaryReference = "Drawdown Reference"
	fwm.SetBeneficiaryReference(br)

	debitDD := wire.NewAccountDebitedDrawdown()
	debitDD.IdentificationCode = wire.DemandDepositAccountNumber
	debitDD.Identifier = "123456789"
	debitDD.Name = "DrawDown Name"
	debitDD.Address.AddressLineOne = "DrawDown Address One"
	debitDD.Address.AddressLineTwo = "DrawDown Address Two"
	debitDD.Address.AddressLineThree = "DrawDown Address Three"
	fwm.SetAccountDebitedDrawdown(debitDD)

	// Originator
	o := wire.NewOriginator()
	o.Personal.IdentificationCode = wire.PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address"
	o.Personal.Address.AddressLineTwo = "Address"
	o.Personal.Address.AddressLineThree = "Address"
	fwm.SetOriginator(o)

	ofi := wire.NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.SetOriginatorFI(ofi)

	ifi := wire.NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.SetInstructingFI(ifi)

	creditDD := wire.NewAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "123456789"
	fwm.SetAccountCreditedDrawdown(creditDD)

	ob := wire.NewOriginatorToBeneficiary()
	ob.LineOne = "Line 1"
	ob.LineTwo = "Line 2 "
	ob.LineThree = "Line 3"
	ob.LineFour = "Line 4"
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	firfi := wire.NewFIReceiverFI()
	firfi.FIToFI.LineOne = "Line 1"
	firfi.FIToFI.LineOne = "Line 2"
	firfi.FIToFI.LineOne = "Line 3"
	firfi.FIToFI.LineOne = "Line 4"
	firfi.FIToFI.LineOne = "Line 5"
	firfi.FIToFI.LineOne = "Line 6"
	fwm.SetFIReceiverFI(firfi)

	fiifi := wire.NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "Line 1"
	fiifi.FIToFI.LineOne = "Line 2"
	fiifi.FIToFI.LineOne = "Line 3"
	fiifi.FIToFI.LineOne = "Line 4"
	fiifi.FIToFI.LineOne = "Line 5"
	fiifi.FIToFI.LineOne = "Line 6"
	fwm.SetFIIntermediaryFI(fiifi)

	fiifia := wire.NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = wire.AdviceCodeLetter
	fiifia.Advice.LineOne = "Line 1"
	fiifia.Advice.LineTwo = "Line 2"
	fiifia.Advice.LineThree = "Line 3"
	fiifia.Advice.LineFour = "Line 4"
	fiifia.Advice.LineFive = "Line 5"
	fiifia.Advice.LineSix = "Line 6"
	fwm.SetFIIntermediaryFIAdvice(fiifia)

	fibfi := wire.NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "Line One"
	fibfi.FIToFI.LineTwo = "Line Two"
	fibfi.FIToFI.LineThree = "Line Three"
	fibfi.FIToFI.LineFour = "Line Four"
	fibfi.FIToFI.LineFive = "Line Five"
	fibfi.FIToFI.LineSix = "Line Six"
	fwm.SetFIBeneficiaryFI(fibfi)

	fibfia := wire.NewFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = wire.AdviceCodeTelex
	fibfia.Advice.LineOne = "Line One"
	fibfia.Advice.LineTwo = "Line Two"
	fibfia.Advice.LineThree = "Line Three"
	fibfia.Advice.LineFour = "Line Four"
	fibfia.Advice.LineFive = "Line Five"
	fibfia.Advice.LineSix = "Line Six"
	fwm.SetFIBeneficiaryFIAdvice(fibfia)

	fib := wire.NewFIBeneficiary()
	fib.FIToFI.LineOne = "Line One"
	fib.FIToFI.LineTwo = "Line Two"
	fib.FIToFI.LineThree = "Line Three"
	fib.FIToFI.LineFour = "Line Four"
	fib.FIToFI.LineFive = "Line Five"
	fib.FIToFI.LineSix = "Line Six"
	fwm.SetFIBeneficiary(fib)

	fiba := wire.NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = wire.AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	fwm.SetFIBeneficiaryAdvice(fiba)

	pm := wire.NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	fwm.SetFIPaymentMethodToBeneficiary(pm)

	fifi := wire.NewFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "Line One"
	fifi.AdditionalFIToFI.LineTwo = "Line Two"
	fifi.AdditionalFIToFI.LineThree = "Line Three"
	fifi.AdditionalFIToFI.LineFour = "Line Four"
	fifi.AdditionalFIToFI.LineFive = "Line Five"
	fifi.AdditionalFIToFI.LineSix = "Line Six"
	fwm.SetFIAdditionalFIToFI(fifi)

	file.AddFEDWireMessage(fwm)

	if err := file.Create(); err != nil {
		log.Fatalf("Could not create FEDWireMessage: %s\n", err)
	}
	if err := file.Validate(); err != nil {
		log.Fatalf("Could not validate FEDWireMessage: %s\n", err)
	}

	w := wire.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()

}
