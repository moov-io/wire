// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"

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
	// ToDo:  Do we need a current FEDWireMessage, just use FEDWireMessage
	// currentFEDWireMessage is the current FEDWireMessage being parsed
	currentFEDWireMessage FEDWireMessage
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

// NewReader returns a new ACH Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	reader := &Reader{
		scanner: bufio.NewScanner(r),
	}

	reader.scanner.Split(scanLinesWithSegmentFormat)

	return reader
}

// addCurrentFEDWireMessage creates the current FEDWireMessage for the file being read. A successful
// current FEDWireMessage will be added to r.File once parsed.
/*func (r *Reader) addCurrentFEDWireMessage(fwm FEDWireMessage) {
	r.currentFEDWireMessage = FEDWireMessage{}
}*/

// Read reads each line of the FED Wire file and defines which parser to use based
// on the first character of each line. It also enforces FED Wire formatting rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {

	stripNewLines := func(input string) string {
		// windows
		input = strings.ReplaceAll(input, "\r\n", "")

		// mac, linux
		input = strings.ReplaceAll(input, "\n", "")

		return input
	}

	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		// ToDo: File length Check?
		r.line = stripNewLines(line)
		if err := r.parseLine(); err != nil {
			r.errors.Add(err)
		}
	}

	r.File.AddFEDWireMessage(r.currentFEDWireMessage)
	r.currentFEDWireMessage = FEDWireMessage{}

	if r.errors.Empty() {
		err := r.File.Validate()
		if err == nil {
			return r.File, nil
		}
		r.errors.Add(fmt.Errorf("file validation failed: %v", err))
	}
	return r.File, r.errors
}

func (r *Reader) parseLine() error { //nolint:gocyclo
	if n := utf8.RuneCountInString(r.line); n < 6 {
		return fmt.Errorf("line %q is too short for tag", r.line)
	}
	switch r.line[:6] {
	case TagMessageDisposition:
		if err := r.parseMessageDisposition(); err != nil {
			return err
		}
	case TagReceiptTimeStamp:
		if err := r.parseReceiptTimeStamp(); err != nil {
			return err
		}
	case TagOutputMessageAccountabilityData:
		if err := r.parseOutputMessageAccountabilityData(); err != nil {
			return err
		}
	case TagErrorWire:
		if err := r.parseErrorWire(); err != nil {
			return err
		}
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
		if err := r.parseFIBeneficiaryFIAdvice(); err != nil {
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
	case TagFIAdditionalFIToFI:
		if err := r.parseFIAdditionalFIToFI(); err != nil {
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
	case TagRemittance:
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
	ss := new(SenderSupplied)
	if err := ss.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ss.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.SenderSupplied = ss
	return nil
}

func (r *Reader) parseTypeSubType() error {
	r.tagName = "TypeSubType"
	tst := new(TypeSubType)
	if err := tst.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := tst.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.TypeSubType = tst
	return nil
}

func (r *Reader) parseInputMessageAccountabilityData() error {
	r.tagName = "InputMessageAccountabilityData"
	imad := new(InputMessageAccountabilityData)
	if err := imad.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := imad.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.InputMessageAccountabilityData = imad
	return nil
}

func (r *Reader) parseAmount() error {
	r.tagName = "Amount"
	amt := new(Amount)
	if err := amt.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := amt.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Amount = amt
	return nil
}

func (r *Reader) parseSenderDepositoryInstitution() error {
	r.tagName = "SenderDepositoryInstitution"
	sdi := new(SenderDepositoryInstitution)
	if err := sdi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := sdi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.SenderDepositoryInstitution = sdi
	return nil
}

func (r *Reader) parseReceiverDepositoryInstitution() error {
	r.tagName = "ReceiverDepositoryInstitution"
	rdi := new(ReceiverDepositoryInstitution)
	if err := rdi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := rdi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ReceiverDepositoryInstitution = rdi
	return nil
}

func (r *Reader) parseBusinessFunctionCode() error {
	r.tagName = "BusinessFunctionCode"
	bfc := new(BusinessFunctionCode)
	if err := bfc.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := bfc.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.BusinessFunctionCode = bfc
	return nil
}

func (r *Reader) parseSenderReference() error {
	r.tagName = "SenderReference"
	sr := new(SenderReference)
	if err := sr.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := sr.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.SenderReference = sr
	return nil
}

func (r *Reader) parsePreviousMessageIdentifier() error {
	r.tagName = "PreviousMessageIdentifier"
	pmi := new(PreviousMessageIdentifier)
	if err := pmi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := pmi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.PreviousMessageIdentifier = pmi
	return nil
}

func (r *Reader) parseLocalInstrument() error {
	r.tagName = "LocalInstrument"
	li := new(LocalInstrument)
	if err := li.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := li.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.LocalInstrument = li
	return nil
}

func (r *Reader) parsePaymentNotification() error {
	r.tagName = "PaymentNotification"
	pn := new(PaymentNotification)
	if err := pn.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := pn.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.PaymentNotification = pn
	return nil
}

func (r *Reader) parseCharges() error {
	r.tagName = "Charges"
	c := new(Charges)
	if err := c.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := c.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Charges = c
	return nil
}

func (r *Reader) parseInstructedAmount() error {
	r.tagName = "InstructedAmount"
	ia := new(InstructedAmount)
	if err := ia.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ia.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.InstructedAmount = ia
	return nil
}

func (r *Reader) parseExchangeRate() error {
	r.tagName = "ExchangeRate"
	eRate := new(ExchangeRate)
	if err := eRate.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := eRate.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ExchangeRate = eRate
	return nil
}

func (r *Reader) parseBeneficiaryIntermediaryFI() error {
	r.tagName = "BeneficiaryIntermediaryFI"
	bifi := new(BeneficiaryIntermediaryFI)
	if err := bifi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := bifi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.BeneficiaryIntermediaryFI = bifi
	return nil
}

func (r *Reader) parseBeneficiaryFI() error {
	r.tagName = "BeneficiaryFI"
	bfi := new(BeneficiaryFI)
	if err := bfi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := bfi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.BeneficiaryFI = bfi
	return nil
}

func (r *Reader) parseBeneficiary() error {
	r.tagName = "Beneficiary"
	ben := new(Beneficiary)
	if err := ben.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ben.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Beneficiary = ben
	return nil
}

func (r *Reader) parseBeneficiaryReference() error {
	r.tagName = "BeneficiaryReference"
	br := new(BeneficiaryReference)
	if err := br.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := br.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.BeneficiaryReference = br
	return nil
}

func (r *Reader) parseAccountDebitedDrawdown() error {
	r.tagName = "AccountDebitedDrawdown"
	debitDD := new(AccountDebitedDrawdown)
	if err := debitDD.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := debitDD.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.AccountDebitedDrawdown = debitDD
	return nil
}

func (r *Reader) parseOriginator() error {
	r.tagName = "Originator"
	o := new(Originator)
	if err := o.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := o.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Originator = o
	return nil
}

func (r *Reader) parseOriginatorOptionF() error {
	r.tagName = "OriginatorOptionF"
	oof := new(OriginatorOptionF)
	if err := oof.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := oof.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OriginatorOptionF = oof
	return nil
}

func (r *Reader) parseOriginatorFI() error {
	r.tagName = "OriginatorFI"
	ofi := new(OriginatorFI)
	if err := ofi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ofi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OriginatorFI = ofi
	return nil
}

func (r *Reader) parseInstructingFI() error {
	r.tagName = "InstructingFI"
	ifi := new(InstructingFI)
	if err := ifi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ifi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.InstructingFI = ifi
	return nil
}

func (r *Reader) parseAccountCreditedDrawdown() error {
	r.tagName = "AccountCreditedDrawdown"
	creditDD := new(AccountCreditedDrawdown)
	if err := creditDD.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := creditDD.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.AccountCreditedDrawdown = creditDD
	return nil
}

func (r *Reader) parseOriginatorToBeneficiary() error {
	r.tagName = "OriginatorToBeneficiary"
	ob := new(OriginatorToBeneficiary)
	if err := ob.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ob.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OriginatorToBeneficiary = ob
	return nil
}

func (r *Reader) parseFIReceiverFI() error {
	r.tagName = "FIReceiverFI"
	firfi := new(FIReceiverFI)
	if err := firfi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := firfi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIReceiverFI = firfi
	return nil
}

func (r *Reader) parseFIDrawdownDebitAccountAdvice() error {
	r.tagName = "FIDrawdownDebitAccountAdvice"
	debitDDAdvice := new(FIDrawdownDebitAccountAdvice)
	if err := debitDDAdvice.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := debitDDAdvice.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIDrawdownDebitAccountAdvice = debitDDAdvice
	return nil
}

func (r *Reader) parseFIIntermediaryFI() error {
	r.tagName = "FIIntermediaryFI"
	fiifi := new(FIIntermediaryFI)
	if err := fiifi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fiifi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIIntermediaryFI = fiifi
	return nil
}

func (r *Reader) parseFIIntermediaryFIAdvice() error {
	r.tagName = "FIIntermediaryFIAdvice"
	fiifia := new(FIIntermediaryFIAdvice)
	if err := fiifia.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fiifia.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIIntermediaryFIAdvice = fiifia
	return nil
}

func (r *Reader) parseFIBeneficiaryFI() error {
	r.tagName = "FIBeneficiaryFI"
	fibfi := new(FIBeneficiaryFI)
	if err := fibfi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fibfi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIBeneficiaryFI = fibfi
	return nil
}

func (r *Reader) parseFIBeneficiaryFIAdvice() error {
	r.tagName = "FIBeneficiaryFIAdvice"
	fibfia := new(FIBeneficiaryFIAdvice)
	if err := fibfia.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fibfia.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIBeneficiaryFIAdvice = fibfia
	return nil
}

func (r *Reader) parseFIBeneficiary() error {
	r.tagName = "FIBeneficiary"
	fib := new(FIBeneficiary)
	if err := fib.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fib.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIBeneficiary = fib
	return nil
}

func (r *Reader) parseFIBeneficiaryAdvice() error {
	r.tagName = "FIBeneficiaryAdvice"
	fiba := new(FIBeneficiaryAdvice)
	if err := fiba.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fiba.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIBeneficiaryAdvice = fiba
	return nil
}

func (r *Reader) parseFIPaymentMethodToBeneficiary() error {
	r.tagName = "FIPaymentMethodToBeneficiary"
	pm := new(FIPaymentMethodToBeneficiary)
	if err := pm.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := pm.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIPaymentMethodToBeneficiary = pm
	return nil
}

func (r *Reader) parseFIAdditionalFIToFI() error {
	r.tagName = "FIAdditionalFiToFi"
	fifi := new(FIAdditionalFIToFI)
	if err := fifi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := fifi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.FIAdditionalFIToFI = fifi
	return nil
}

func (r *Reader) parseCurrencyInstructedAmount() error {
	r.tagName = "CurrencyInstructedAmount"
	cia := new(CurrencyInstructedAmount)
	if err := cia.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := cia.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.CurrencyInstructedAmount = cia
	return nil
}

func (r *Reader) parseOrderingCustomer() error {
	r.tagName = "OrderingCustomer"
	oc := new(OrderingCustomer)
	if err := oc.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := oc.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OrderingCustomer = oc
	return nil
}

func (r *Reader) parseOrderingInstitution() error {
	r.tagName = "OrderingInstitution"
	oi := new(OrderingInstitution)
	if err := oi.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := oi.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OrderingInstitution = oi
	return nil
}

func (r *Reader) parseIntermediaryInstitution() error {
	r.tagName = "IntermediaryInstitution"
	ii := new(IntermediaryInstitution)
	if err := ii.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ii.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.IntermediaryInstitution = ii
	return nil
}

func (r *Reader) parseInstitutionAccount() error {
	r.tagName = "InstitutionAccount"
	iAccount := new(InstitutionAccount)
	if err := iAccount.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := iAccount.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.InstitutionAccount = iAccount
	return nil
}

func (r *Reader) parseBeneficiaryCustomer() error {
	r.tagName = "BeneficiaryCustomer"
	bc := new(BeneficiaryCustomer)
	if err := bc.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := bc.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.BeneficiaryCustomer = bc
	return nil
}

func (r *Reader) parseRemittance() error {
	r.tagName = "Remittance"
	ri := new(Remittance)
	if err := ri.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ri.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Remittance = ri
	return nil
}

func (r *Reader) parseSenderToReceiver() error {
	r.tagName = "SenderToReceiver"
	sr := new(SenderToReceiver)
	if err := sr.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := sr.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.SenderToReceiver = sr
	return nil
}

func (r *Reader) parseUnstructuredAddenda() error {
	r.tagName = "UnstructuredAddenda"
	ua := new(UnstructuredAddenda)
	if err := ua.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ua.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.UnstructuredAddenda = ua
	return nil
}

func (r *Reader) parseRelatedRemittance() error {
	r.tagName = "RelatedRemittance"
	rr := new(RelatedRemittance)
	if err := rr.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := rr.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.RelatedRemittance = rr
	return nil
}

func (r *Reader) parseRemittanceOriginator() error {
	r.tagName = "RemittanceOriginator"
	ro := new(RemittanceOriginator)
	if err := ro.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ro.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.RemittanceOriginator = ro
	return nil
}

func (r *Reader) parseRemittanceBeneficiary() error {
	r.tagName = "RemittanceBeneficiary"
	rb := new(RemittanceBeneficiary)
	if err := rb.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := rb.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.RemittanceBeneficiary = rb
	return nil
}

func (r *Reader) parsePrimaryRemittanceDocument() error {
	r.tagName = "PrimaryRemittanceDocument"
	prd := new(PrimaryRemittanceDocument)
	if err := prd.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := prd.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.PrimaryRemittanceDocument = prd
	return nil
}

func (r *Reader) parseActualAmountPaid() error {
	r.tagName = "ActualAmountPaid"
	aap := new(ActualAmountPaid)
	if err := aap.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := aap.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ActualAmountPaid = aap
	return nil
}

func (r *Reader) parseGrossAmountRemittanceDocument() error {
	r.tagName = "GrossAmountRemittanceDocument"
	gard := new(GrossAmountRemittanceDocument)
	if err := gard.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := gard.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.GrossAmountRemittanceDocument = gard
	return nil
}

func (r *Reader) parseAmountNegotiatedDiscount() error {
	r.tagName = "AmountNegotiatedDiscount"
	nd := new(AmountNegotiatedDiscount)
	if err := nd.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := nd.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.AmountNegotiatedDiscount = nd
	return nil
}

func (r *Reader) parseAdjustment() error {
	r.tagName = "Adjustment"
	adj := new(Adjustment)
	if err := adj.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := adj.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.Adjustment = adj
	return nil
}

func (r *Reader) parseDateRemittanceDocument() error {
	r.tagName = "DateRemittanceDocument"
	drd := new(DateRemittanceDocument)
	if err := drd.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := drd.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.DateRemittanceDocument = drd
	return nil
}

func (r *Reader) parseSecondaryRemittanceDocument() error {
	r.tagName = "SecondaryRemittanceDocument"
	srd := new(SecondaryRemittanceDocument)
	if err := srd.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := srd.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.SecondaryRemittanceDocument = srd
	return nil
}

func (r *Reader) parseRemittanceFreeText() error {
	r.tagName = "RemittanceFreeText"
	rft := new(RemittanceFreeText)
	if err := rft.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := rft.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.RemittanceFreeText = rft
	return nil
}

func (r *Reader) parseServiceMessage() error {
	r.tagName = "ServiceMessage"
	sm := new(ServiceMessage)
	if err := sm.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := sm.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ServiceMessage = sm
	return nil
}

func (r *Reader) parseMessageDisposition() error {
	r.tagName = "MessageDisposition"
	md := new(MessageDisposition)
	if err := md.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := md.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.MessageDisposition = md
	return nil
}

func (r *Reader) parseReceiptTimeStamp() error {
	r.tagName = "ReceiptTimeStamp"
	rts := new(ReceiptTimeStamp)
	if err := rts.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := rts.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ReceiptTimeStamp = rts
	return nil
}

func (r *Reader) parseOutputMessageAccountabilityData() error {
	r.tagName = "OutputMessageAccountabilityData"
	omad := new(OutputMessageAccountabilityData)
	if err := omad.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := omad.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.OutputMessageAccountabilityData = omad
	return nil
}

func (r *Reader) parseErrorWire() error {
	r.tagName = "ErrorWire"
	ew := new(ErrorWire)
	if err := ew.Parse(r.line); err != nil {
		return r.parseError(err)
	}
	if err := ew.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentFEDWireMessage.ErrorWire = ew
	return nil
}

var (
	tagRegex = regexp.MustCompile(`{([0-9]{4})}`)
)

//scanLinesWithSegmentFormat allows Reader to read each segment
func scanLinesWithSegmentFormat(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	indexes := tagRegex.FindAllIndex(data, -1)
	if len(indexes) == 0 {
		return len(data), data, nil
	}

	if len(indexes) < 2 && !atEOF {
		// need more data
		return 0, nil, nil
	}

	firstIndex := indexes[0]
	if firstIndex[0] > 0 {
		return firstIndex[0], data[:firstIndex[0]], nil
	}

	if len(indexes) == 1 {
		return len(data), data, nil
	}

	secondIndex := indexes[1]
	length := secondIndex[0]

	return length, data[:length], nil
}
