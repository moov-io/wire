// Copyright 2020 The WIRE Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"io"
	"strings"

	"golang.org/x/exp/slices"
)

// A Writer writes an fedWireMessage to an encoded file.
//
// As returned by NewWriter, a Writer writes FEDWireMessage file structs into
// FEDWireMessage formatted files.

// Writer struct
type Writer struct {
	w       *bufio.Writer
	lineNum int // current line being written
	FormatOptions
}

type OptionFunc func(*Writer)

// VariableLengthFields specify to support variable length
func VariableLengthFields(variableLength bool) OptionFunc {
	return func(w *Writer) {
		w.VariableLengthFields = variableLength
	}
}

// NewlineCharacter specify new line character
func NewlineCharacter(newline string) OptionFunc {
	return func(w *Writer) {
		w.NewlineCharacter = newline
	}
}

// NewWriter returns a new Writer that writes to w.
// If no opts are provided, the writer will default to fixed-length fields and use "\n" for newlines.
func NewWriter(w io.Writer, opts ...OptionFunc) *Writer {
	writer := &Writer{
		w: bufio.NewWriter(w),
		FormatOptions: FormatOptions{
			NewlineCharacter: "\n",
		},
	}

	for _, opt := range opts {
		opt(writer)
	}

	return writer
}

// Writer writes a single FEDWireMessage record to w
// options
//
//	first bool : has variable length
//	second bool : has not new line
func (w *Writer) Write(file *File) error {
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeFEDWireMessage(file); err != nil {
		return err
	}
	w.lineNum++

	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
// Flush writes any buffered data to the underlying io.Writer.
func (w *Writer) Flush() error {
	return w.w.Flush()
}

func (w *Writer) writeFEDWireMessage(file *File) error {
	fwm := file.FEDWireMessage

	var outputLines []string

	mandatoryLines, err := w.writeMandatory(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, mandatoryLines...)

	otherTransferLines, err := w.writeOtherTransferInfo(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, otherTransferLines...)

	beneficiaryLines, err := w.writeBeneficiary(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, beneficiaryLines...)

	originatorLines, err := w.writeOriginator(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, originatorLines...)

	financialInstitutionLines, err := w.writeFinancialInstitution(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, financialInstitutionLines...)

	coverPaymentLines, err := w.writeCoverPayment(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, coverPaymentLines...)

	if fwm.UnstructuredAddenda != nil {
		outputLines = append(outputLines, fwm.UnstructuredAddenda.String())
	}

	remittanceLines, err := w.writeRemittance(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, remittanceLines...)

	if fwm.ServiceMessage != nil {
		outputLines = append(outputLines, fwm.ServiceMessage.Format(w.FormatOptions))
	}

	fedAppendedLines, err := w.writeFedAppended(fwm)
	if err != nil {
		return err
	}
	outputLines = append(outputLines, fedAppendedLines...)

	slices.Sort(outputLines)
	w.w.WriteString(strings.Join(outputLines, w.NewlineCharacter))
	w.w.WriteString(w.NewlineCharacter)

	return nil
}

func (w *Writer) writeFedAppended(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.MessageDisposition != nil {
		lines = append(lines, fwm.MessageDisposition.Format(w.FormatOptions))
	}

	if fwm.ReceiptTimeStamp != nil {
		lines = append(lines, fwm.ReceiptTimeStamp.Format(w.FormatOptions))
	}

	if fwm.OutputMessageAccountabilityData != nil {
		lines = append(lines, fwm.OutputMessageAccountabilityData.Format(w.FormatOptions))
	}

	if fwm.ErrorWire != nil {
		lines = append(lines, fwm.ErrorWire.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeMandatory(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.SenderSupplied != nil {
		lines = append(lines, fwm.SenderSupplied.Format(w.FormatOptions))
	} else {
		if fwm.requireSenderSupplied() {
			return nil, fieldError("SenderSupplied", ErrFieldRequired)
		}
	}

	if fwm.TypeSubType != nil {
		lines = append(lines, fwm.TypeSubType.String())
	} else {
		return nil, fieldError("TypeSubType", ErrFieldRequired)
	}

	if fwm.InputMessageAccountabilityData != nil {
		lines = append(lines, fwm.InputMessageAccountabilityData.String())
	} else {
		return nil, fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}

	if fwm.Amount != nil {
		lines = append(lines, fwm.Amount.String())
	} else {
		return nil, fieldError("Amount", ErrFieldRequired)
	}

	if fwm.SenderDepositoryInstitution != nil {
		lines = append(lines, fwm.SenderDepositoryInstitution.Format(w.FormatOptions))
	} else {
		return nil, fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.ReceiverDepositoryInstitution != nil {
		lines = append(lines, fwm.ReceiverDepositoryInstitution.Format(w.FormatOptions))
	} else {
		return nil, fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.BusinessFunctionCode != nil {
		lines = append(lines, fwm.BusinessFunctionCode.Format(w.FormatOptions))
	} else {
		return nil, fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	return lines, nil
}

func (w *Writer) writeOtherTransferInfo(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.SenderReference != nil {
		lines = append(lines, fwm.SenderReference.Format(w.FormatOptions))
	}

	if fwm.PreviousMessageIdentifier != nil {
		lines = append(lines, fwm.PreviousMessageIdentifier.Format(w.FormatOptions))
	}

	if fwm.LocalInstrument != nil {
		lines = append(lines, fwm.LocalInstrument.Format(w.FormatOptions))
	}

	if fwm.PaymentNotification != nil {
		lines = append(lines, fwm.PaymentNotification.Format(w.FormatOptions))
	}

	if fwm.Charges != nil {
		lines = append(lines, fwm.Charges.Format(w.FormatOptions))
	}

	if fwm.InstructedAmount != nil {
		lines = append(lines, fwm.InstructedAmount.Format(w.FormatOptions))
	}

	if fwm.ExchangeRate != nil {
		lines = append(lines, fwm.ExchangeRate.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeBeneficiary(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.BeneficiaryIntermediaryFI != nil {
		lines = append(lines, fwm.BeneficiaryIntermediaryFI.Format(w.FormatOptions))
	}

	if fwm.BeneficiaryFI != nil {
		lines = append(lines, fwm.BeneficiaryFI.Format(w.FormatOptions))
	}

	if fwm.Beneficiary != nil {
		lines = append(lines, fwm.Beneficiary.Format(w.FormatOptions))
	}

	if fwm.BeneficiaryReference != nil {
		lines = append(lines, fwm.BeneficiaryReference.Format(w.FormatOptions))
	}

	if fwm.AccountDebitedDrawdown != nil {
		lines = append(lines, fwm.AccountDebitedDrawdown.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeOriginator(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.Originator != nil {
		lines = append(lines, fwm.Originator.Format(w.FormatOptions))
	}

	if fwm.OriginatorOptionF != nil {
		lines = append(lines, fwm.OriginatorOptionF.Format(w.FormatOptions))
	}

	if fwm.OriginatorFI != nil {
		lines = append(lines, fwm.OriginatorFI.Format(w.FormatOptions))
	}

	if fwm.InstructingFI != nil {
		lines = append(lines, fwm.InstructingFI.Format(w.FormatOptions))
	}

	if fwm.AccountCreditedDrawdown != nil {
		lines = append(lines, fwm.AccountCreditedDrawdown.Format(w.FormatOptions))
	}

	if fwm.OriginatorToBeneficiary != nil {
		lines = append(lines, fwm.OriginatorToBeneficiary.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeFinancialInstitution(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.FIReceiverFI != nil {
		lines = append(lines, fwm.FIReceiverFI.Format(w.FormatOptions))
	}

	if fwm.FIDrawdownDebitAccountAdvice != nil {
		lines = append(lines, fwm.FIDrawdownDebitAccountAdvice.Format(w.FormatOptions))
	}

	if fwm.FIIntermediaryFI != nil {
		lines = append(lines, fwm.FIIntermediaryFI.Format(w.FormatOptions))
	}

	if fwm.FIIntermediaryFIAdvice != nil {
		lines = append(lines, fwm.FIIntermediaryFIAdvice.Format(w.FormatOptions))
	}

	if fwm.FIBeneficiaryFI != nil {
		lines = append(lines, fwm.FIBeneficiaryFI.Format(w.FormatOptions))
	}

	if fwm.FIBeneficiaryFIAdvice != nil {
		lines = append(lines, fwm.FIBeneficiaryFIAdvice.Format(w.FormatOptions))
	}

	if fwm.FIBeneficiary != nil {
		lines = append(lines, fwm.FIBeneficiary.Format(w.FormatOptions))
	}

	if fwm.FIBeneficiaryAdvice != nil {
		lines = append(lines, fwm.FIBeneficiaryAdvice.Format(w.FormatOptions))
	}

	if fwm.FIPaymentMethodToBeneficiary != nil {
		lines = append(lines, fwm.FIPaymentMethodToBeneficiary.Format(w.FormatOptions))
	}

	if fwm.FIAdditionalFIToFI != nil {
		lines = append(lines, fwm.FIAdditionalFIToFI.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeCoverPayment(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	if fwm.CurrencyInstructedAmount != nil {
		lines = append(lines, fwm.CurrencyInstructedAmount.Format(w.FormatOptions))
	}

	if fwm.OrderingCustomer != nil {
		lines = append(lines, fwm.OrderingCustomer.Format(w.FormatOptions))
	}

	if fwm.OrderingInstitution != nil {
		lines = append(lines, fwm.OrderingInstitution.Format(w.FormatOptions))
	}

	if fwm.IntermediaryInstitution != nil {
		lines = append(lines, fwm.IntermediaryInstitution.Format(w.FormatOptions))
	}

	if fwm.InstitutionAccount != nil {
		lines = append(lines, fwm.InstitutionAccount.Format(w.FormatOptions))
	}

	if fwm.BeneficiaryCustomer != nil {
		lines = append(lines, fwm.BeneficiaryCustomer.Format(w.FormatOptions))
	}

	if fwm.Remittance != nil {
		lines = append(lines, fwm.Remittance.Format(w.FormatOptions))
	}

	if fwm.SenderToReceiver != nil {
		lines = append(lines, fwm.SenderToReceiver.Format(w.FormatOptions))
	}

	return lines, nil
}

func (w *Writer) writeRemittance(fwm FEDWireMessage) ([]string, error) {
	var lines []string

	// Related Remittance
	if fwm.RelatedRemittance != nil {
		lines = append(lines, fwm.RelatedRemittance.Format(w.FormatOptions))
	}

	// Structured Remittance
	if fwm.RemittanceOriginator != nil {
		lines = append(lines, fwm.RemittanceOriginator.Format(w.FormatOptions))
	}

	if fwm.RemittanceBeneficiary != nil {
		lines = append(lines, fwm.RemittanceBeneficiary.Format(w.FormatOptions))
	}

	if fwm.PrimaryRemittanceDocument != nil {
		lines = append(lines, fwm.PrimaryRemittanceDocument.Format(w.FormatOptions))
	}

	if fwm.ActualAmountPaid != nil {
		lines = append(lines, fwm.ActualAmountPaid.Format(w.FormatOptions))
	}

	if fwm.GrossAmountRemittanceDocument != nil {
		lines = append(lines, fwm.GrossAmountRemittanceDocument.Format(w.FormatOptions))
	}

	if fwm.AmountNegotiatedDiscount != nil {
		lines = append(lines, fwm.AmountNegotiatedDiscount.Format(w.FormatOptions))
	}

	if fwm.Adjustment != nil {
		lines = append(lines, fwm.Adjustment.Format(w.FormatOptions))
	}

	if fwm.DateRemittanceDocument != nil {
		lines = append(lines, fwm.DateRemittanceDocument.String())
	}

	if fwm.SecondaryRemittanceDocument != nil {
		lines = append(lines, fwm.SecondaryRemittanceDocument.Format(w.FormatOptions))
	}

	if fwm.RemittanceFreeText != nil {
		lines = append(lines, fwm.RemittanceFreeText.Format(w.FormatOptions))
	}

	return lines, nil
}
