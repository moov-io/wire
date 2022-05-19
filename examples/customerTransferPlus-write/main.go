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
	fwm := wire.FEDWireMessage{}

	// Mandatory Fields
	ss := wire.NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = wire.MessageDuplicationOriginal
	fwm.SenderSupplied = ss

	tst := wire.NewTypeSubType()
	tst.TypeCode = wire.FundsTransfer
	tst.SubTypeCode = wire.BasicFundsTransfer
	fwm.TypeSubType = tst

	imad := wire.NewInputMessageAccountabilityData()
	imad.InputCycleDate = time.Now().Format("20060102")
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	fwm.InputMessageAccountabilityData = imad

	amt := wire.NewAmount()
	amt.Amount = "000001234567"
	fwm.Amount = amt

	sdi := wire.NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	fwm.SenderDepositoryInstitution = sdi

	rdi := wire.NewReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	fwm.ReceiverDepositoryInstitution = rdi

	bfc := wire.NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = wire.CustomerTransferPlus
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := wire.NewSenderReference()
	sr.SenderReference = "Sender Reference"
	fwm.SenderReference = sr

	pmi := wire.NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	fwm.PreviousMessageIdentifier = pmi

	li := wire.NewLocalInstrument()
	li.LocalInstrumentCode = wire.ProprietaryLocalInstrumentCode
	li.ProprietaryCode = "PROP CODE"
	fwm.LocalInstrument = li

	pn := wire.NewPaymentNotification()
	pn.PaymentNotificationIndicator = "1"
	pn.ContactNotificationElectronicAddress = "http://moov.io"
	pn.ContactName = "Contact Name"
	pn.ContactPhoneNumber = "5555551212"
	pn.ContactMobileNumber = "5551231212"
	pn.ContactFaxNumber = "5554561212"
	fwm.PaymentNotification = pn

	c := wire.NewCharges()
	c.ChargeDetails = "B"
	c.SendersChargesOne = "USD0,99"
	c.SendersChargesTwo = "USD2,99"
	c.SendersChargesThree = "USD3,99"
	c.SendersChargesFour = "USD1,00"
	fwm.Charges = c

	ia := wire.NewInstructedAmount()
	ia.CurrencyCode = "USD"
	ia.Amount = "4567,89"
	fwm.InstructedAmount = ia

	eRate := wire.NewExchangeRate()
	eRate.ExchangeRate = "1,2345"
	fwm.ExchangeRate = eRate

	// Beneficiary
	bifi := wire.NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryIntermediaryFI = bifi

	bfi := wire.NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryFI = bfi

	ben := wire.NewBeneficiary()
	ben.Personal.IdentificationCode = wire.DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	fwm.Beneficiary = ben

	br := wire.NewBeneficiaryReference()
	br.BeneficiaryReference = "Reference"
	fwm.BeneficiaryReference = br

	// Originator
	o := wire.NewOriginator()
	o.Personal.IdentificationCode = wire.PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address One"
	o.Personal.Address.AddressLineTwo = "Address Two"
	o.Personal.Address.AddressLineThree = "Address Three"
	fwm.Originator = o

	oof := wire.NewOriginatorOptionF()
	oof.PartyIdentifier = "TXID/123-45-6789"
	oof.Name = "1/Name"
	oof.LineOne = "1/1234"
	oof.LineTwo = "2/1000 Colonial Farm Rd"
	oof.LineThree = "5/Pottstown"
	fwm.OriginatorOptionF = oof

	ofi := wire.NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.OriginatorFI = ofi

	ifi := wire.NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "FI Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "FI Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "FI Address Three"
	fwm.InstructingFI = ifi

	ob := wire.NewOriginatorToBeneficiary()
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	fiifi := wire.NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "FI Intermediary Line One"
	fiifi.FIToFI.LineOne = "FI Intermediary Line Two"
	fiifi.FIToFI.LineOne = "FI Intermediary Line Three"
	fiifi.FIToFI.LineOne = ""
	fiifi.FIToFI.LineOne = ""
	fiifi.FIToFI.LineOne = ""
	fwm.FIIntermediaryFI = fiifi

	fiifia := wire.NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = wire.AdviceCodeLetter
	fiifia.Advice.LineOne = "Intermediary Advice Line One"
	fiifia.Advice.LineTwo = "Intermediary Advice Line Two"
	fiifia.Advice.LineThree = "Intermediary Advice Line Three"
	fiifia.Advice.LineFour = "Intermediary Advice Line Four"
	fiifia.Advice.LineFive = "Line Five"
	fiifia.Advice.LineSix = "Line Six"
	fwm.FIIntermediaryFIAdvice = fiifia

	fibfi := wire.NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "Beneficiary FI Line One"
	fibfi.FIToFI.LineTwo = "Beneficiary FI Line Two"
	fibfi.FIToFI.LineThree = "Beneficiary FI Line Three"
	fibfi.FIToFI.LineFour = "Beneficiary FI Line Four"
	fibfi.FIToFI.LineFive = "Beneficiary FI Line Five"
	fibfi.FIToFI.LineSix = "Beneficiary FI Line Six"
	fwm.FIBeneficiaryFI = fibfi

	fibfia := wire.NewFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = wire.AdviceCodeTelex
	fibfia.Advice.LineOne = "Line One"
	fibfia.Advice.LineTwo = "Line Two"
	fibfia.Advice.LineThree = "Line Three"
	fibfia.Advice.LineFour = "Line Four"
	fibfia.Advice.LineFive = "Line Five"
	fibfia.Advice.LineSix = "Line Six"
	fwm.FIBeneficiaryFIAdvice = fibfia

	fib := wire.NewFIBeneficiary()
	fib.FIToFI.LineOne = "Line One"
	fib.FIToFI.LineTwo = "Line Two"
	fib.FIToFI.LineThree = "Line Three"
	fib.FIToFI.LineFour = "Line Four"
	fib.FIToFI.LineFive = "Line Five"
	fib.FIToFI.LineSix = "Line Six"
	fwm.FIBeneficiary = fib

	fiba := wire.NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = wire.AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	fwm.FIBeneficiaryAdvice = fiba

	pm := wire.NewFIPaymentMethodToBeneficiary()
	pm.AdditionalInformation = "Additional Information"
	fwm.FIPaymentMethodToBeneficiary = pm

	fifi := wire.NewFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "Line One"
	fifi.AdditionalFIToFI.LineTwo = "Line Two"
	fifi.AdditionalFIToFI.LineThree = "Line Three"
	fifi.AdditionalFIToFI.LineFour = "Line Four"
	fifi.AdditionalFIToFI.LineFive = "Line Five"
	fifi.AdditionalFIToFI.LineSix = "Line Six"
	fwm.FIAdditionalFIToFI = fifi

	// ServiceMessage
	sm := wire.NewServiceMessage()
	sm.LineOne = "Line One"
	sm.LineTwo = "Line Two"
	sm.LineThree = "Line Three"
	sm.LineFour = "Line Four"
	sm.LineFive = "Line Five"
	sm.LineSix = "Line Six"
	sm.LineSeven = "Line Seven"
	sm.LineEight = "Line Eight"
	sm.LineNine = "Line Nine"
	sm.LineTen = "Line Ten"
	sm.LineEleven = "Line Eleven"
	sm.LineTwelve = "line Twelve"
	fwm.ServiceMessage = sm

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
