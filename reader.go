// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"github.com/moov-io/base"
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// line is the current line being parsed from the input r
	line string
	// currentFedWireMessage is the current FedWireMessage being parsed
	currentFedWireMessage FedWireMessage
	// lineNum is the line number of the file being parsed
	lineNum int
	// tagName holds the current tag name being parsed.
	tagName string
	// errors holds each error encountered when attempting to parse the file
	errors base.ErrorList
}

// error returns a new ParseError based on err
func (r *Reader) parseError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*base.ParseError); ok {
		return err
	}
	return &base.ParseError{
		Line:   r.lineNum,
		Record: r.tagName,
		Err:    err,
	}
}

// addCurrentFedWireMessage creates the current FedWireMessage for the file being read. A successful
// current FedWireMessage will be added to r.File once parsed.
func (r *Reader) addCurrentFedWireMessage(fwm FedWireMessage) {
	r.currentFedWireMessage = FedWireMessage{}
}

// Read reads each line of the FED Wire file and defines which parser to use based
// on the first character of each line. It also enforces FED Wire formatting rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		// ToDo: File length Check?
		r.line = line
		if err := r.parseLine(); err != nil {
			return r.File, err
		}
	}
	return r.File, nil
}

func (r *Reader) parseLine() error {
	switch r.line[:6] {
	case TagSenderSupplied:
		if err := r.parseSenderSupplied(); err != nil {
			return err
		}
	case TagTypeSubType:
		if err := r.parseTypeSubType(); err != nil {
			return err
		}
	case TagInputMessageAccountabilityData:
		if err := r.parseInputMessageAccountabilityData(); err != nil {
			return err
		}
	case TagAmount:
		if err := r.parseAmount(); err != nil {
			return err
		}
	case TagSenderDepositoryInstitution:
		if err := r.parseSenderDepositoryInstitution(); err != nil {
			return err
		}
	case TagReceiverDepositoryInstitution:
		if err := r.parseReceiverDepositoryInstitution(); err != nil {
			return err
		}
	case TagBusinessFunctionCode:
		if err := r.parseBusinessFunctionCode(); err != nil {
			return err
		}
	case TagSenderReference:
		if err := r.parseSenderReference(); err != nil {
			return err
		}
	case TagPreviousMessageIdentifier:
		if err := r.parsePreviousMessageIdentifier(); err != nil {
			return err
		}
	case TagLocalInstrument:
		if err := r.parseLocalInstrument(); err != nil {
			return err
		}
	case TagPaymentNotification:
		if err := r.parsePaymentNotification(); err != nil {
			return err
		}
	case TagCharges:
		if err := r.parseCharges(); err != nil {
			return err
		}
	case TagInstructedAmount:
		if err := r.parseInstructedAmount(); err != nil {
			return err
		}
	case TagExchangeRate:
		if err := r.parseExchangeRate(); err != nil {
			return err
		}

	case TagBeneficiaryIntermediaryFI:
		if err := r.parseBeneficiaryIntermediaryFI(); err != nil {
			return err
		}
	case TagBeneficiaryFI:
		if err := r.parseBeneficiaryFI(); err != nil {
			return err
		}
	case TagBeneficiary:
		if err := r.parseBeneficiary(); err != nil {
			return err
		}
	case TagBeneficiaryReference:
		if err := r.parseBeneficiaryReference(); err != nil {
			return err
		}
	case TagAccountDebitedDrawdown:
		if err := r.parseAccountDebitedDrawdown(); err != nil {
			return err
		}
	case TagOriginator:
		if err := r.parseOriginator(); err != nil {
			return err
		}
	case TagOriginatorOptionF:
		if err := r.parseOriginatorOptionF(); err != nil {
			return err
		}
	case TagOriginatorFI:
		if err := r.parseOriginatorFI(); err != nil {
			return err
		}
	case TagInstructingFI:
		if err := r.parseInstructingFI(); err != nil {
			return err
		}
	case TagAccountCreditedDrawdown:
		if err := r.parseAccountCreditedDrawdown(); err != nil {
			return err
		}
	case TagOriginatorToBeneficiary:
		if err := r.parseOriginatorToBeneficiary(); err != nil {
			return err
		}
	case TagFIReceiverFI:
		if err := r.parseFIReceiverFI(); err != nil {
			return err
		}
	case TagFIDrawdownDebitAccountAdvice:
		if err := r.parseFIDrawdownDebitAccountAdvice(); err != nil {
			return err
		}
	case TagFIIntermediaryFI:
		if err := r.parseFIIntermediaryFI(); err != nil {
			return err
		}
	case TagFIIntermediaryFIAdvice:
		if err := r.parseFIIntermediaryFIAdvice(); err != nil {
			return err
		}
	case TagFIBeneficiaryFI:
		if err := r.parseFIBeneficiaryFI(); err != nil {
			return err
		}
	case TagFIBeneficiaryFIAdvice:
		if err := r.parseFIBeneficiaryFI(); err != nil {
			return err
		}
	case TagFIBeneficiary:
		if err := r.parseFIBeneficiary(); err != nil {
			return err
		}
	case TagFIBeneficiaryAdvice:
		if err := r.parseFIBeneficiaryAdvice(); err != nil {
			return err
		}
	case TagFIPaymentMethodToBeneficiary:
		if err := r.parseFIPaymentMethodToBeneficiary(); err != nil {
			return err
		}
	case TagFIAdditionalFiToFi:
		if err := r.parseFIAdditionalFiToFi(); err != nil {
			return err
		}
	case TagCurrencyInstructedAmount:
		if err := r.parseCurrencyInstructedAmount(); err != nil {
			return err
		}
	case TagOrderingCustomer:
		if err := r.parseOrderingCustomer(); err != nil {
			return err
		}
	case TagOrderingInstitution:
		if err := r.parseOrderingInstitution(); err != nil {
			return err
		}
	case TagIntermediaryInstitution:
		if err := r.parseIntermediaryInstitution(); err != nil {
			return err
		}
	case TagInstitutionAccount:
		if err := r.parseInstitutionAccount(); err != nil {
			return err
		}
	case TagBeneficiaryCustomer:
		if err := r.parseBeneficiaryCustomer(); err != nil {
			return err
		}
		// ToDo: Remove Information
	case TagRemittanceInformation:
		if err := r.parseRemittance(); err != nil {
			return err
		}
	case TagSenderToReceiver:
		if err := r.parseSenderToReceiver(); err != nil {
			return err
		}
	case TagUnstructuredAddenda:
		if err := r.parseUnstructuredAddenda(); err != nil {
			return err
		}
	case TagRelatedRemittance:
		if err := r.parseRelatedRemittance(); err != nil {
			return err
		}
	case TagRemittanceOriginator:
		if err := r.parseRemittanceOriginator(); err != nil {
			return err
		}
	case TagRemittanceBeneficiary:
		if err := r.parseRemittanceBeneficiary(); err != nil {
			return err
		}
	case TagPrimaryRemittanceDocument:
		if err := r.parsePrimaryRemittanceDocument(); err != nil {
			return err
		}
	case TagActualAmountPaid:
		if err := r.parseActualAmountPaid(); err != nil {
			return err
		}
	case TagGrossAmountRemittanceDocument:
		if err := r.parseGrossAmountRemittanceDocument(); err != nil {
			return err
		}
	case TagAmountNegotiatedDiscount:
		if err := r.parseAmountNegotiatedDiscount(); err != nil {
			return err
		}
	case TagAdjustment:
		if err := r.parseAdjustment(); err != nil {
			return err
		}
	case TagDateRemittanceDocument:
		if err := r.parseDateRemittanceDocument(); err != nil {
			return err
		}
	case TagSecondaryRemittanceDocument:
		if err := r.parseSecondaryRemittanceDocument(); err != nil {
			return err
		}
	case TagRemittanceFreeText:
		if err := r.parseRemittanceFreeText(); err != nil {
			return err
		}
	case TagServiceMessage:
		if err := r.parseServiceMessage(); err != nil {
			return err
		}
	default:
		return NewErrInvalidTag(r.line[:6])
	}
	return nil
}

func (r *Reader) parseSenderSupplied() error {
	r.tagName = "SenderSupplied"
	if len(r.line) != 18 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(18, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.SenderSupplied.Parse(r.line)
	if err := r.File.FedWireMessage.SenderSupplied.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseTypeSubType() error {
	r.tagName = "TypeSubType"
	if len(r.line) != 10 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(10, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.TypeSubType.Parse(r.line)
	if err := r.File.FedWireMessage.TypeSubType.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseInputMessageAccountabilityData() error {
	r.tagName = "InputMessageAccountabilityData"
	if len(r.line) != 22 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(22, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.InputMessageAccountabilityData.Parse(r.line)
	if err := r.File.FedWireMessage.InputMessageAccountabilityData.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAmount() error {
	r.tagName = "Amount"
	if len(r.line) != 18 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(18, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.Amount.Parse(r.line)
	if err := r.File.FedWireMessage.Amount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseSenderDepositoryInstitution() error {
	r.tagName = "SenderDepositoryInstitution"
	if len(r.line) != 39 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(39, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.SenderDepositoryInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.SenderDepositoryInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseReceiverDepositoryInstitution() error {
	r.tagName = "ReceiverDepositoryInstitution"
	if len(r.line) != 39 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(39, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.ReceiverDepositoryInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.ReceiverDepositoryInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBusinessFunctionCode() error {
	r.tagName = "BusinessFunctionCode"
	if len(r.line) != 12 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(12, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.BusinessFunctionCode.Parse(r.line)
	if err := r.File.FedWireMessage.BusinessFunctionCode.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseSenderReference() error {
	r.tagName = "SenderReference"
	r.File.FedWireMessage.SenderReference.Parse(r.line)
	if err := r.File.FedWireMessage.SenderReference.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parsePreviousMessageIdentifier() error {
	r.tagName = "PreviousMessageIdentifier"
	r.File.FedWireMessage.PreviousMessageIdentifier.Parse(r.line)
	if err := r.File.FedWireMessage.PreviousMessageIdentifier.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseLocalInstrument() error {
	r.tagName = "LocalInstrument"
	r.File.FedWireMessage.LocalInstrument.Parse(r.line)
	if err := r.File.FedWireMessage.LocalInstrument.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parsePaymentNotification() error {
	r.tagName = "PaymentNotification"
	r.File.FedWireMessage.PaymentNotification.Parse(r.line)
	if err := r.File.FedWireMessage.PaymentNotification.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseCharges() error {
	r.tagName = "Charges"
	r.File.FedWireMessage.Charges.Parse(r.line)
	if err := r.File.FedWireMessage.Charges.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseInstructedAmount() error {
	r.tagName = "InstructedAmount"
	r.File.FedWireMessage.InstructedAmount.Parse(r.line)
	if err := r.File.FedWireMessage.InstructedAmount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseExchangeRate() error {
	r.tagName = "ExchangeRate"
	r.File.FedWireMessage.ExchangeRate.Parse(r.line)
	if err := r.File.FedWireMessage.ExchangeRate.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBeneficiaryIntermediaryFI() error {
	r.tagName = "BeneficiaryIntermediaryFI"
	r.File.FedWireMessage.BeneficiaryIntermediaryFI.Parse(r.line)
	if err := r.File.FedWireMessage.BeneficiaryIntermediaryFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBeneficiaryFI() error {
	r.tagName = "BeneficiaryFI"
	r.File.FedWireMessage.BeneficiaryFI.Parse(r.line)
	if err := r.File.FedWireMessage.BeneficiaryFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBeneficiary() error {
	r.tagName = "Beneficiary"
	r.File.FedWireMessage.Beneficiary.Parse(r.line)
	if err := r.File.FedWireMessage.Beneficiary.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBeneficiaryReference() error {
	r.tagName = "BeneficiaryReference"
	r.File.FedWireMessage.BeneficiaryReference.Parse(r.line)
	if err := r.File.FedWireMessage.BeneficiaryReference.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAccountDebitedDrawdown() error {
	r.tagName = "BeneficiaryReference"
	r.File.FedWireMessage.AccountDebitedDrawdown.Parse(r.line)
	if err := r.File.FedWireMessage.AccountDebitedDrawdown.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOriginator() error {
	r.tagName = "Originator"
	r.File.FedWireMessage.Originator.Parse(r.line)
	if err := r.File.FedWireMessage.Originator.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOriginatorOptionF() error {
	r.tagName = "OriginatorOptionF"
	r.File.FedWireMessage.OriginatorOptionF.Parse(r.line)
	if err := r.File.FedWireMessage.OriginatorOptionF.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOriginatorFI() error {
	r.tagName = "OriginatorFI"
	r.File.FedWireMessage.OriginatorFI.Parse(r.line)
	if err := r.File.FedWireMessage.OriginatorFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseInstructingFI() error {
	r.tagName = "InstructingFI"
	r.File.FedWireMessage.InstructingFI.Parse(r.line)
	if err := r.File.FedWireMessage.InstructingFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAccountCreditedDrawdown() error {
	r.tagName = "AccountCreditedDrawdown"
	r.File.FedWireMessage.AccountCreditedDrawdown.Parse(r.line)
	if err := r.File.FedWireMessage.AccountCreditedDrawdown.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOriginatorToBeneficiary() error {
	r.tagName = "OriginatorToBeneficiary"
	r.File.FedWireMessage.OriginatorToBeneficiary.Parse(r.line)
	if err := r.File.FedWireMessage.OriginatorToBeneficiary.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIReceiverFI() error {
	r.tagName = "FIReceiverFI"
	r.File.FedWireMessage.FIReceiverFI.Parse(r.line)
	if err := r.File.FedWireMessage.FIReceiverFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIDrawdownDebitAccountAdvice() error {
	r.tagName = "FIDrawdownDebitAccountAdvice"
	r.File.FedWireMessage.FIDrawdownDebitAccountAdvice.Parse(r.line)
	if err := r.File.FedWireMessage.FIDrawdownDebitAccountAdvice.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIIntermediaryFI() error {
	r.tagName = "FIIntermediaryFI"
	r.File.FedWireMessage.FIIntermediaryFI.Parse(r.line)
	if err := r.File.FedWireMessage.FIIntermediaryFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIIntermediaryFIAdvice() error {
	r.tagName = "FIIntermediaryFIAdvice"
	r.File.FedWireMessage.FIIntermediaryFIAdvice.Parse(r.line)
	if err := r.File.FedWireMessage.FIIntermediaryFIAdvice.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIBeneficiaryFI() error {
	r.tagName = "FIBeneficiaryFI"
	r.File.FedWireMessage.FIBeneficiaryFI.Parse(r.line)
	if err := r.File.FedWireMessage.FIBeneficiaryFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIBeneficiaryFIAdvice() error {
	r.tagName = "FIBeneficiaryFIAdvice"
	r.File.FedWireMessage.FIBeneficiaryFIAdvice.Parse(r.line)
	if err := r.File.FedWireMessage.FIBeneficiaryFIAdvice.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIBeneficiary() error {
	r.tagName = "FIBeneficiary"
	r.File.FedWireMessage.FIBeneficiary.Parse(r.line)
	if err := r.File.FedWireMessage.FIBeneficiary.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIBeneficiaryAdvice() error {
	r.tagName = "FIBeneficiaryAdvice"
	r.File.FedWireMessage.FIBeneficiaryAdvice.Parse(r.line)
	if err := r.File.FedWireMessage.FIBeneficiaryAdvice.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIPaymentMethodToBeneficiary() error {
	r.tagName = "FIPaymentMethodToBeneficiary"
	r.File.FedWireMessage.FIPaymentMethodToBeneficiary.Parse(r.line)
	if err := r.File.FedWireMessage.FIPaymentMethodToBeneficiary.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseFIAdditionalFiToFi() error {
	r.tagName = "FIAdditionalFiToFi"
	r.File.FedWireMessage.FIAdditionalFIToFI.Parse(r.line)
	if err := r.File.FedWireMessage.FIAdditionalFIToFI.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseCurrencyInstructedAmount() error {
	r.tagName = "CurrencyInstructedAmount"
	r.File.FedWireMessage.CurrencyInstructedAmount.Parse(r.line)
	if err := r.File.FedWireMessage.CurrencyInstructedAmount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOrderingCustomer() error {
	r.tagName = "OrderingCustomer"
	r.File.FedWireMessage.OrderingCustomer.Parse(r.line)
	if err := r.File.FedWireMessage.OrderingCustomer.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseOrderingInstitution() error {
	r.tagName = "OrderingInstitution"
	r.File.FedWireMessage.OrderingInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.OrderingInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseIntermediaryInstitution() error {
	r.tagName = "IntermediaryInstitution"
	r.File.FedWireMessage.IntermediaryInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.IntermediaryInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseInstitutionAccount() error {
	r.tagName = "InstitutionAccount"
	r.File.FedWireMessage.InstitutionAccount.Parse(r.line)
	if err := r.File.FedWireMessage.InstitutionAccount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBeneficiaryCustomer() error {
	r.tagName = "BeneficiaryCustomer"
	r.File.FedWireMessage.BeneficiaryCustomer.Parse(r.line)
	if err := r.File.FedWireMessage.BeneficiaryCustomer.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseRemittance() error {
	r.tagName = "Remittance"
	r.File.FedWireMessage.Remittance.Parse(r.line)
	if err := r.File.FedWireMessage.Remittance.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseSenderToReceiver() error {
	r.tagName = "SenderToReceiver"
	r.File.FedWireMessage.SenderToReceiver.Parse(r.line)
	if err := r.File.FedWireMessage.SenderToReceiver.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseUnstructuredAddenda() error {
	r.tagName = "UnstructuredAddenda"
	r.File.FedWireMessage.UnstructuredAddenda.Parse(r.line)
	if err := r.File.FedWireMessage.UnstructuredAddenda.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseRelatedRemittance() error {
	r.tagName = "RelatedRemittance"
	r.File.FedWireMessage.RelatedRemittance.Parse(r.line)
	if err := r.File.FedWireMessage.RelatedRemittance.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseRemittanceOriginator() error {
	r.tagName = "RemittanceOriginator"
	r.File.FedWireMessage.RemittanceOriginator.Parse(r.line)
	if err := r.File.FedWireMessage.RemittanceOriginator.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseRemittanceBeneficiary() error {
	r.tagName = "RemittanceBeneficiary"
	r.File.FedWireMessage.RemittanceBeneficiary.Parse(r.line)
	if err := r.File.FedWireMessage.RemittanceBeneficiary.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parsePrimaryRemittanceDocument() error {
	r.tagName = "PrimaryRemittanceDocument"
	r.File.FedWireMessage.PrimaryRemittanceDocument.Parse(r.line)
	if err := r.File.FedWireMessage.PrimaryRemittanceDocument.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseActualAmountPaid() error {
	r.tagName = "ActualAmountPaid"
	r.File.FedWireMessage.ActualAmountPaid.Parse(r.line)
	if err := r.File.FedWireMessage.ActualAmountPaid.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseGrossAmountRemittanceDocument() error {
	r.tagName = "GrossAmountRemittanceDocument"
	r.File.FedWireMessage.GrossAmountRemittanceDocument.Parse(r.line)
	if err := r.File.FedWireMessage.GrossAmountRemittanceDocument.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAmountNegotiatedDiscount() error {
	r.tagName = "AmountNegotiatedDiscount"
	r.File.FedWireMessage.AmountNegotiatedDiscount.Parse(r.line)
	if err := r.File.FedWireMessage.AmountNegotiatedDiscount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAdjustment() error {
	r.tagName = "Adjustment"
	r.File.FedWireMessage.Adjustment.Parse(r.line)
	if err := r.File.FedWireMessage.Adjustment.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseDateRemittanceDocument() error {
	r.tagName = "DateRemittanceDocument"
	r.File.FedWireMessage.DateRemittanceDocument.Parse(r.line)
	if err := r.File.FedWireMessage.DateRemittanceDocument.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseSecondaryRemittanceDocument() error {
	r.tagName = "SecondaryRemittanceDocument"
	r.File.FedWireMessage.SecondaryRemittanceDocument.Parse(r.line)
	if err := r.File.FedWireMessage.SecondaryRemittanceDocument.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseRemittanceFreeText() error {
	r.tagName = "RemittanceFreeText"
	r.File.FedWireMessage.RemittanceFreeText.Parse(r.line)
	if err := r.File.FedWireMessage.RemittanceFreeText.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseServiceMessage() error {
	r.tagName = "ServiceMessage"
	r.File.FedWireMessage.ServiceMessage.Parse(r.line)
	if err := r.File.FedWireMessage.ServiceMessage.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}
