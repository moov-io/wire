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
	tst.TypeCode = wire.FundsTransfer
	tst.SubTypeCode = wire.BasicFundsTransfer
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
	bfc.BusinessFunctionCode = wire.CustomerTransferPlus
	bfc.TransactionTypeCode = "   "
	fwm.SetBusinessFunctionCode(bfc)

	// Other Transfer Information
	sr := wire.NewSenderReference()
	sr.SenderReference = "Sender Reference"
	fwm.SetSenderReference(sr)

	pmi := wire.NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	fwm.SetPreviousMessageIdentifier(pmi)

	li := wire.NewLocalInstrument()
	li.LocalInstrumentCode = wire.SequenceBCoverPaymentStructured
	li.ProprietaryCode = ""
	fwm.SetLocalInstrument(li)

	pn := wire.NewPaymentNotification()
	pn.PaymentNotificationIndicator = "1"
	pn.ContactNotificationElectronicAddress = "http://moov.io"
	pn.ContactName = "Contact Name"
	pn.ContactPhoneNumber = "5555551212"
	pn.ContactMobileNumber = "5551231212"
	pn.ContactFaxNumber = "5554561212"
	fwm.SetPaymentNotification(pn)

	// Beneficiary
	bifi := wire.NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	bfi := wire.NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.SetBeneficiaryFI(bfi)

	ben := wire.NewBeneficiary()
	ben.Personal.IdentificationCode = wire.DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	fwm.SetBeneficiary(ben)

	br := wire.NewBeneficiaryReference()
	br.BeneficiaryReference = "Reference"
	fwm.SetBeneficiaryReference(br)

	// Originator
	o := wire.NewOriginator()
	o.Personal.IdentificationCode = wire.PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address One"
	o.Personal.Address.AddressLineTwo = "Address Two"
	o.Personal.Address.AddressLineThree = "Address Three"
	fwm.SetOriginator(o)

	oof := wire.NewOriginatorOptionF()
	oof.PartyIdentifier = "TXID/123-45-6789"
	oof.Name = "1/Name"
	oof.LineOne = "1/1234"
	oof.LineTwo = "2/1000 Colonial Farm Rd"
	oof.LineThree = "5/Pottstown"
	fwm.SetOriginatorOptionF(oof)

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

	ob := wire.NewOriginatorToBeneficiary()
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	fwm.SetOriginatorToBeneficiary(ob)

	// FI to FI
	fiifi := wire.NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	fwm.SetFIIntermediaryFI(fiifi)

	fiifia := wire.NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = wire.AdviceCodeLetter
	fiifia.Advice.LineOne = "Line One"
	fiifia.Advice.LineTwo = "Line Two"
	fiifia.Advice.LineThree = "Line Three"
	fiifia.Advice.LineFour = "Line Four"
	fiifia.Advice.LineFive = "Line Five"
	fiifia.Advice.LineSix = "Line Six"
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

	// Cover Payment Information
	cia := wire.NewCurrencyInstructedAmount()
	cia.SwiftFieldTag = "Swift Field Tag"
	cia.Amount = "1500,49"
	fwm.SetCurrencyInstructedAmount(cia)
	oc := wire.NewOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oc.CoverPayment.SwiftLineOne = "Swift Line One"
	oc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oc.CoverPayment.SwiftLineThree = "Swift Line Three"
	oc.CoverPayment.SwiftLineFour = "Swift Line Four"
	oc.CoverPayment.SwiftLineFive = "Swift Line Five"
	fwm.SetOrderingCustomer(oc)
	oi := wire.NewOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oi.CoverPayment.SwiftLineOne = "Swift Line One"
	oi.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oi.CoverPayment.SwiftLineThree = "Swift Line Three"
	oi.CoverPayment.SwiftLineFour = "Swift Line Four"
	oi.CoverPayment.SwiftLineFive = "Swift Line Five"
	fwm.SetOrderingInstitution(oi)
	ii := wire.NewIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ii.CoverPayment.SwiftLineOne = "Swift Line One"
	ii.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ii.CoverPayment.SwiftLineThree = "Swift Line Three"
	ii.CoverPayment.SwiftLineFour = "Swift Line Four"
	ii.CoverPayment.SwiftLineFive = "Swift Line Five"
	fwm.SetIntermediaryInstitution(ii)
	iAccount := wire.NewInstitutionAccount()
	iAccount.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	iAccount.CoverPayment.SwiftLineOne = "Swift Line One"
	iAccount.CoverPayment.SwiftLineTwo = "Swift Line Two"
	iAccount.CoverPayment.SwiftLineThree = "Swift Line Three"
	iAccount.CoverPayment.SwiftLineFour = "Swift Line Four"
	iAccount.CoverPayment.SwiftLineFive = "Swift Line Five"
	fwm.SetInstitutionAccount(iAccount)
	bc := wire.NewBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	bc.CoverPayment.SwiftLineOne = "Swift Line One"
	bc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	bc.CoverPayment.SwiftLineThree = "Swift Line Three"
	bc.CoverPayment.SwiftLineFour = "Swift Line Four"
	bc.CoverPayment.SwiftLineFive = "Swift Line Five"
	fwm.SetBeneficiaryCustomer(bc)
	ri := wire.NewRemittance()
	ri.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ri.CoverPayment.SwiftLineOne = "Swift Line One"
	ri.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ri.CoverPayment.SwiftLineThree = "Swift Line Three"
	ri.CoverPayment.SwiftLineFour = "Swift Line Four"
	fwm.SetRemittance(ri)
	str := wire.NewSenderToReceiver()
	str.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	str.CoverPayment.SwiftLineOne = "Swift Line One"
	str.CoverPayment.SwiftLineTwo = "Swift Line Two"
	str.CoverPayment.SwiftLineThree = "Swift Line Three"
	str.CoverPayment.SwiftLineFour = "Swift Line Four"
	str.CoverPayment.SwiftLineFive = "Swift Line Five"
	str.CoverPayment.SwiftLineSix = "Swift Line Six"
	fwm.SetSenderToReceiver(str)

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
