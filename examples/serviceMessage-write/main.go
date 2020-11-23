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
	// SenderSupplied
	ss := wire.NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = wire.MessageDuplicationOriginal
	fwm.SenderSupplied = ss

	tst := wire.NewTypeSubType()
	tst.TypeCode = wire.FundsTransfer
	tst.SubTypeCode = wire.RequestReversal
	fwm.TypeSubType = tst

	// InputMessageAccountabilityData
	imad := wire.NewInputMessageAccountabilityData()
	imad.InputCycleDate = time.Now().Format("20060102")
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	fwm.InputMessageAccountabilityData = imad

	// Amount
	amt := wire.NewAmount()
	amt.Amount = "000001234567"
	fwm.Amount = amt

	// SenderDepositoryInstitution
	sdi := wire.NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	fwm.SenderDepositoryInstitution = sdi

	rdi := wire.NewReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	fwm.ReceiverDepositoryInstitution = rdi

	bfc := wire.NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = wire.BFCServiceMessage
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	// Sender Reference
	sr := wire.NewSenderReference()
	sr.SenderReference = "Sender Reference"
	fwm.SenderReference = sr

	// Previous Message Identifier
	pmi := wire.NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	// Beneficiary Intermediary FI
	bifi := wire.NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryIntermediaryFI = bifi

	// Beneficiary FI
	bfi := wire.NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryFI = bfi

	// Beneficiary
	ben := wire.NewBeneficiary()
	ben.Personal.IdentificationCode = wire.DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	fwm.Beneficiary = ben

	// Beneficiary Reference
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

	// Originator FI
	ofi := wire.NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.OriginatorFI = ofi

	// Instructing FI
	ifi := wire.NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.InstructingFI = ifi

	// Originator To Beneficiary
	ob := wire.NewOriginatorToBeneficiary()
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	// FIReceiverFI
	firfi := wire.NewFIReceiverFI()
	firfi.FIToFI.LineOne = "FIToFI Line One"
	firfi.FIToFI.LineOne = "FIToFI Line Two"
	firfi.FIToFI.LineOne = "FIToFI Line Three"
	firfi.FIToFI.LineOne = "FIToFI Line Four"
	firfi.FIToFI.LineOne = "FIToFI Line Five"
	firfi.FIToFI.LineOne = "FIToFI Line Six"
	fwm.FIReceiverFI = firfi

	// FIIntermediaryFI
	fiifi := wire.NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line One"
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line Two"
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line Three"
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line Four"
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line Five"
	fiifi.FIToFI.LineOne = "FIIntermediaryFI Line Six"
	fwm.FIIntermediaryFI = fiifi

	// FIIntermediaryFIAdvice
	fiifia := wire.NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = wire.AdviceCodeLetter
	fiifia.Advice.LineOne = "FIInterFIAdvice Line One"
	fiifia.Advice.LineTwo = "FIInterFIAdvice Line Two"
	fiifia.Advice.LineThree = "FIInterFIAdvice Line Three"
	fiifia.Advice.LineFour = "FIInterFIAdvice Line Four"
	fiifia.Advice.LineFive = "FIInterFIAdvice Line Five"
	fiifia.Advice.LineSix = "FIInterFIAdvice Line Six"
	fwm.FIIntermediaryFIAdvice = fiifia

	// FIBeneficiaryFI
	fibfi := wire.NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "FIBenFI Line One"
	fibfi.FIToFI.LineTwo = "FIBenFI Line Two"
	fibfi.FIToFI.LineThree = "FIBenFI Line Three"
	fibfi.FIToFI.LineFour = "FIBenFI Line Four"
	fibfi.FIToFI.LineFive = "FIBenFI Line Five"
	fibfi.FIToFI.LineSix = "FIBenFI Line Six"
	fwm.FIBeneficiaryFI = fibfi

	// FIBeneficiaryFIAdvice
	fibfia := wire.NewFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = wire.AdviceCodeTelex
	fibfia.Advice.LineOne = "FIBenFIAdvice Line One"
	fibfia.Advice.LineTwo = "FIBenFIAdvice Line Two"
	fibfia.Advice.LineThree = "FIBenFIAdvice Line Three"
	fibfia.Advice.LineFour = "FIBenFIAdvice Line Four"
	fibfia.Advice.LineFive = "FIBenFIAdvice Line Five"
	fibfia.Advice.LineSix = "FIBenFIAdvice Line Six"
	fwm.FIBeneficiaryFIAdvice = fibfia

	// FIBeneficiary
	fib := wire.NewFIBeneficiary()
	fib.FIToFI.LineOne = "FIBen Line One"
	fib.FIToFI.LineTwo = "FIBen Line Two"
	fib.FIToFI.LineThree = "FIBen Line Three"
	fib.FIToFI.LineFour = "FIBen Line Four"
	fib.FIToFI.LineFive = "FIBen Line Five"
	fib.FIToFI.LineSix = "FIBen Line Six"
	fwm.FIBeneficiary = fib

	// FIBeneficiaryAdvice
	fiba := wire.NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = wire.AdviceCodeLetter
	fiba.Advice.LineOne = "FIBenAdvice Line One"
	fiba.Advice.LineTwo = "FIBenAdvice Line Two"
	fiba.Advice.LineThree = "FIBenAdvice Line Three"
	fiba.Advice.LineFour = "FIBenAdvice Line Four"
	fiba.Advice.LineFive = "FIBenAdvice Line Five"
	fiba.Advice.LineSix = "FIBenAdvice Line Six"
	fwm.FIBeneficiaryAdvice = fiba

	// FIPaymentMethodToBeneficiary
	pm := wire.NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	fwm.FIPaymentMethodToBeneficiary = pm

	// FIAdditionalFIToFI
	fifi := wire.NewFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "FIAddFI Line One"
	fifi.AdditionalFIToFI.LineTwo = "FIAddFI Line Two"
	fifi.AdditionalFIToFI.LineThree = "FIAddFI Line Three"
	fifi.AdditionalFIToFI.LineFour = "FIAddFI Line Four"
	fifi.AdditionalFIToFI.LineFive = "FIAddFI Line Five"
	fifi.AdditionalFIToFI.LineSix = "FIAddFI Line Six"
	fwm.FIAdditionalFIToFI = fifi

	sm := wire.NewServiceMessage()
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
