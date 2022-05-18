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

	isVariableLength := false
	file := wire.NewFile()
	fwm := wire.FEDWireMessage{}

	// Mandatory Fields
	ss := wire.NewSenderSupplied(isVariableLength)
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = wire.MessageDuplicationOriginal
	fwm.SenderSupplied = ss

	tst := wire.NewTypeSubType(isVariableLength)
	tst.TypeCode = wire.SettlementTransfer
	tst.SubTypeCode = wire.BasicFundsTransfer
	fwm.TypeSubType = tst

	imad := wire.NewInputMessageAccountabilityData(isVariableLength)
	imad.InputCycleDate = time.Now().Format("20060102")
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	fwm.InputMessageAccountabilityData = imad

	amt := wire.NewAmount(isVariableLength)
	amt.Amount = "000001234567"
	fwm.Amount = amt

	sdi := wire.NewSenderDepositoryInstitution(isVariableLength)
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	fwm.SenderDepositoryInstitution = sdi

	rdi := wire.NewReceiverDepositoryInstitution(isVariableLength)
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	fwm.ReceiverDepositoryInstitution = rdi

	bfc := wire.NewBusinessFunctionCode(isVariableLength)
	bfc.BusinessFunctionCode = wire.CheckSameDaySettlement
	bfc.TransactionTypeCode = "   "
	fwm.BusinessFunctionCode = bfc

	// Other Transfer Information
	sr := wire.NewSenderReference(isVariableLength)
	sr.SenderReference = "Sender Reference"
	fwm.SenderReference = sr

	pmi := wire.NewPreviousMessageIdentifier(isVariableLength)
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	fwm.PreviousMessageIdentifier = pmi

	// Beneficiary
	bifi := wire.NewBeneficiaryIntermediaryFI(isVariableLength)
	bifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryIntermediaryFI = bifi

	bfi := wire.NewBeneficiaryFI(isVariableLength)
	bfi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.BeneficiaryFI = bfi

	ben := wire.NewBeneficiary(isVariableLength)
	ben.Personal.IdentificationCode = wire.DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	fwm.Beneficiary = ben

	br := wire.NewBeneficiaryReference(isVariableLength)
	br.BeneficiaryReference = "Reference"
	fwm.BeneficiaryReference = br

	// Originator
	o := wire.NewOriginator(isVariableLength)
	o.Personal.IdentificationCode = wire.PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address One"
	o.Personal.Address.AddressLineTwo = "Address Two"
	o.Personal.Address.AddressLineThree = "Address Three"
	fwm.Originator = o

	ofi := wire.NewOriginatorFI(isVariableLength)
	ofi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.OriginatorFI = ofi

	ifi := wire.NewInstructingFI(isVariableLength)
	ifi.FinancialInstitution.IdentificationCode = wire.DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	fwm.InstructingFI = ifi

	ob := wire.NewOriginatorToBeneficiary(isVariableLength)
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	fwm.OriginatorToBeneficiary = ob

	// FI to FI
	firfi := wire.NewFIReceiverFI(isVariableLength)
	firfi.FIToFI.LineOne = "Line One"
	firfi.FIToFI.LineOne = "Line Two"
	firfi.FIToFI.LineOne = "Line Three"
	firfi.FIToFI.LineOne = "Line Four"
	firfi.FIToFI.LineOne = "Line Five"
	firfi.FIToFI.LineOne = "Line Six"
	fwm.FIReceiverFI = firfi

	fiifi := wire.NewFIIntermediaryFI(isVariableLength)
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	fwm.FIIntermediaryFI = fiifi

	fiifia := wire.NewFIIntermediaryFIAdvice(isVariableLength)
	fiifia.Advice.AdviceCode = wire.AdviceCodeLetter
	fiifia.Advice.LineOne = "Line One"
	fiifia.Advice.LineTwo = "Line Two"
	fiifia.Advice.LineThree = "Line Three"
	fiifia.Advice.LineFour = "Line Four"
	fiifia.Advice.LineFive = "Line Five"
	fiifia.Advice.LineSix = "Line Six"
	fwm.FIIntermediaryFIAdvice = fiifia

	fibfi := wire.NewFIBeneficiaryFI(isVariableLength)
	fibfi.FIToFI.LineOne = "Line One"
	fibfi.FIToFI.LineTwo = "Line Two"
	fibfi.FIToFI.LineThree = "Line Three"
	fibfi.FIToFI.LineFour = "Line Four"
	fibfi.FIToFI.LineFive = "Line Five"
	fibfi.FIToFI.LineSix = "Line Six"
	fwm.FIBeneficiaryFI = fibfi

	fibfia := wire.NewFIBeneficiaryFIAdvice(isVariableLength)
	fibfia.Advice.AdviceCode = wire.AdviceCodeTelex
	fibfia.Advice.LineOne = "Line One"
	fibfia.Advice.LineTwo = "Line Two"
	fibfia.Advice.LineThree = "Line Three"
	fibfia.Advice.LineFour = "Line Four"
	fibfia.Advice.LineFive = "Line Five"
	fibfia.Advice.LineSix = "Line Six"
	fwm.FIBeneficiaryFIAdvice = fibfia

	fib := wire.NewFIBeneficiary(isVariableLength)
	fib.FIToFI.LineOne = "Line One"
	fib.FIToFI.LineTwo = "Line Two"
	fib.FIToFI.LineThree = "Line Three"
	fib.FIToFI.LineFour = "Line Four"
	fib.FIToFI.LineFive = "Line Five"
	fib.FIToFI.LineSix = "Line Six"
	fwm.FIBeneficiary = fib

	fiba := wire.NewFIBeneficiaryAdvice(isVariableLength)
	fiba.Advice.AdviceCode = wire.AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	fwm.FIBeneficiaryAdvice = fiba

	pm := wire.NewFIPaymentMethodToBeneficiary(isVariableLength)
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	fwm.FIPaymentMethodToBeneficiary = pm

	fifi := wire.NewFIAdditionalFIToFI(isVariableLength)
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	fwm.FIAdditionalFIToFI = fifi

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
