// Copyright 2020 The WIRE Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"io"
)

// A Writer writes an fedWireMessage to an encoded file.
//
// As returned by NewWriter, a Writer writes FEDWireMessage file structs into
// FEDWireMessage formatted files.

// Writer struct
type Writer struct {
	w       *bufio.Writer
	lineNum int //current line being written
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Writer writes a single FEDWireMessage record to w
// options
//  first bool : has variable length
//  second bool : has not new line
func (w *Writer) Write(file *File, options ...bool) error {
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeFEDWireMessage(file, options...); err != nil {
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

// options
//  first bool : has variable length
//  second bool : has not new line
func (w *Writer) writeFEDWireMessage(file *File, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	fwm := file.FEDWireMessage
	if err := w.writeMandatory(fwm, options...); err != nil {
		return err
	}

	if err := w.writeOtherTransferInfo(fwm, options...); err != nil {
		return err
	}

	if err := w.writeBeneficiary(fwm, options...); err != nil {
		return err
	}

	if err := w.writeOriginator(fwm, options...); err != nil {
		return err
	}

	if err := w.writeFinancialInstitution(fwm, options...); err != nil {
		return err
	}

	if err := w.writeCoverPayment(fwm, options...); err != nil {
		return err
	}

	if fwm.UnstructuredAddenda != nil {
		if _, err := w.w.WriteString(fwm.UnstructuredAddenda.String() + newLine); err != nil {
			return err
		}
	}

	if err := w.writeRemittance(fwm, options...); err != nil {
		return err
	}

	if fwm.ServiceMessage != nil {
		if _, err := w.w.WriteString(fwm.ServiceMessage.String(options...) + newLine); err != nil {
			return err
		}
	}

	if err := w.writeFedAppended(fwm, options...); err != nil {
		return err
	}

	return nil
}

func (w *Writer) writeFedAppended(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.MessageDisposition != nil {
		if _, err := w.w.WriteString(fwm.MessageDisposition.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.ReceiptTimeStamp != nil {
		if _, err := w.w.WriteString(fwm.ReceiptTimeStamp.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OutputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.OutputMessageAccountabilityData.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.ErrorWire != nil {
		if _, err := w.w.WriteString(fwm.ErrorWire.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeMandatory(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.SenderSupplied != nil {
		if _, err := w.w.WriteString(fwm.SenderSupplied.String(options...) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}

	if fwm.TypeSubType != nil {
		if _, err := w.w.WriteString(fwm.TypeSubType.String() + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("TypeSubType", ErrFieldRequired)
	}

	if fwm.InputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.InputMessageAccountabilityData.String() + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}

	if fwm.Amount != nil {
		if _, err := w.w.WriteString(fwm.Amount.String() + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("Amount", ErrFieldRequired)
	}

	if fwm.SenderDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.SenderDepositoryInstitution.String(options...) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.ReceiverDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.ReceiverDepositoryInstitution.String(options...) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.BusinessFunctionCode != nil {
		if _, err := w.w.WriteString(fwm.BusinessFunctionCode.String(options...) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	return nil
}

func (w *Writer) writeOtherTransferInfo(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.SenderReference != nil {
		if _, err := w.w.WriteString(fwm.SenderReference.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.PreviousMessageIdentifier != nil {
		if _, err := w.w.WriteString(fwm.PreviousMessageIdentifier.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.LocalInstrument != nil {
		if _, err := w.w.WriteString(fwm.LocalInstrument.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.PaymentNotification != nil {
		if _, err := w.w.WriteString(fwm.PaymentNotification.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.Charges != nil {
		if _, err := w.w.WriteString(fwm.Charges.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.InstructedAmount.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.ExchangeRate != nil {
		if _, err := w.w.WriteString(fwm.ExchangeRate.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeBeneficiary(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.BeneficiaryIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryIntermediaryFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryFI != nil {
		if fwm.BeneficiaryFI != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryFI.String(options...) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.Beneficiary != nil {
		if fwm.Beneficiary != nil {
			if _, err := w.w.WriteString(fwm.Beneficiary.String(options...) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.BeneficiaryReference != nil {
		if fwm.BeneficiaryReference != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryReference.String(options...) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.AccountDebitedDrawdown != nil {
		if fwm.AccountDebitedDrawdown != nil {
			if _, err := w.w.WriteString(fwm.AccountDebitedDrawdown.String(options...) + newLine); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *Writer) writeOriginator(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.Originator != nil {
		if _, err := w.w.WriteString(fwm.Originator.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorOptionF != nil {
		if _, err := w.w.WriteString(fwm.OriginatorOptionF.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorFI != nil {
		if _, err := w.w.WriteString(fwm.OriginatorFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstructingFI != nil {
		if _, err := w.w.WriteString(fwm.InstructingFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.AccountCreditedDrawdown != nil {
		if _, err := w.w.WriteString(fwm.AccountCreditedDrawdown.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.OriginatorToBeneficiary.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeFinancialInstitution(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.FIReceiverFI != nil {
		if _, err := w.w.WriteString(fwm.FIReceiverFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIDrawdownDebitAccountAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIDrawdownDebitAccountAdvice.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFIAdvice.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFIAdvice.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiary.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryAdvice.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIPaymentMethodToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIPaymentMethodToBeneficiary.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIAdditionalFIToFI != nil {
		if _, err := w.w.WriteString(fwm.FIAdditionalFIToFI.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeCoverPayment(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	if fwm.CurrencyInstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.CurrencyInstructedAmount.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OrderingCustomer != nil {
		if _, err := w.w.WriteString(fwm.OrderingCustomer.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.OrderingInstitution != nil {
		if _, err := w.w.WriteString(fwm.OrderingInstitution.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.IntermediaryInstitution != nil {
		if _, err := w.w.WriteString(fwm.IntermediaryInstitution.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstitutionAccount != nil {
		if _, err := w.w.WriteString(fwm.InstitutionAccount.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryCustomer != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryCustomer.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.Remittance != nil {
		if _, err := w.w.WriteString(fwm.Remittance.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.SenderToReceiver != nil {
		if _, err := w.w.WriteString(fwm.SenderToReceiver.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeRemittance(fwm FEDWireMessage, options ...bool) error {

	newLine := "\n"
	if hasNotNewLine(options) {
		newLine = ""
	}

	// Related Remittance
	if fwm.RelatedRemittance != nil {
		if _, err := w.w.WriteString(fwm.RelatedRemittance.String(options...) + newLine); err != nil {
			return err
		}
	}

	// Structured Remittance
	if fwm.RemittanceOriginator != nil {
		if _, err := w.w.WriteString(fwm.RemittanceOriginator.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.RemittanceBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.RemittanceBeneficiary.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.PrimaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.PrimaryRemittanceDocument.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.ActualAmountPaid != nil {
		if _, err := w.w.WriteString(fwm.ActualAmountPaid.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.GrossAmountRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.GrossAmountRemittanceDocument.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.AmountNegotiatedDiscount != nil {
		if _, err := w.w.WriteString(fwm.AmountNegotiatedDiscount.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.Adjustment != nil {
		if _, err := w.w.WriteString(fwm.Adjustment.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.DateRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.DateRemittanceDocument.String() + newLine); err != nil {
			return err
		}
	}

	if fwm.SecondaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.SecondaryRemittanceDocument.String(options...) + newLine); err != nil {
			return err
		}
	}

	if fwm.RemittanceFreeText != nil {
		if _, err := w.w.WriteString(fwm.RemittanceFreeText.String(options...) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// get second option from options, has not new line
func hasNotNewLine(options []bool) bool {

	firstOption := false

	if len(options) > 1 {
		firstOption = options[1]
	}

	return firstOption
}
