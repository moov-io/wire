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
//  first option : hasVariableLength
//  second option : hasNotNewLine
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
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeFEDWireMessage(file *File, options ...bool) error {
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

	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.UnstructuredAddenda != nil {
		if _, err := w.w.WriteString(fwm.UnstructuredAddenda.String() + newLine); err != nil {
			return err
		}
	}
	if err := w.writeRemittance(fwm); err != nil {
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

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeFedAppended(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.MessageDisposition != nil {
		if _, err := w.w.WriteString(fwm.MessageDisposition.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}
	if fwm.ReceiptTimeStamp != nil {
		if _, err := w.w.WriteString(fwm.ReceiptTimeStamp.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}
	if fwm.OutputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.OutputMessageAccountabilityData.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}
	if fwm.ErrorWire != nil {
		if _, err := w.w.WriteString(fwm.ErrorWire.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeMandatory(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.SenderSupplied != nil {
		if _, err := w.w.WriteString(fwm.SenderSupplied.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}

	if fwm.TypeSubType != nil {
		if _, err := w.w.WriteString(fwm.TypeSubType.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("TypeSubType", ErrFieldRequired)
	}

	if fwm.InputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.InputMessageAccountabilityData.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}

	if fwm.Amount != nil {
		if _, err := w.w.WriteString(fwm.Amount.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("Amount", ErrFieldRequired)
	}

	if fwm.SenderDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.SenderDepositoryInstitution.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.ReceiverDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.ReceiverDepositoryInstitution.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.BusinessFunctionCode != nil {
		if _, err := w.w.WriteString(fwm.BusinessFunctionCode.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeOtherTransferInfo(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.SenderReference != nil {
		if _, err := w.w.WriteString(fwm.SenderReference.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.PreviousMessageIdentifier != nil {
		if _, err := w.w.WriteString(fwm.PreviousMessageIdentifier.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.LocalInstrument != nil {
		if _, err := w.w.WriteString(fwm.LocalInstrument.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.PaymentNotification != nil {
		if _, err := w.w.WriteString(fwm.PaymentNotification.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.Charges != nil {
		if _, err := w.w.WriteString(fwm.Charges.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.InstructedAmount.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.ExchangeRate != nil {
		if _, err := w.w.WriteString(fwm.ExchangeRate.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeBeneficiary(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.BeneficiaryIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryIntermediaryFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryFI != nil {
		if fwm.BeneficiaryFI != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryFI.String(hasVariableLength) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.Beneficiary != nil {
		if fwm.Beneficiary != nil {
			if _, err := w.w.WriteString(fwm.Beneficiary.String(hasVariableLength) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.BeneficiaryReference != nil {
		if fwm.BeneficiaryReference != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryReference.String(hasVariableLength) + newLine); err != nil {
				return err
			}
		}
	}

	if fwm.AccountDebitedDrawdown != nil {
		if fwm.AccountDebitedDrawdown != nil {
			if _, err := w.w.WriteString(fwm.AccountDebitedDrawdown.String(hasVariableLength) + newLine); err != nil {
				return err
			}
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeOriginator(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.Originator != nil {
		if _, err := w.w.WriteString(fwm.Originator.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorOptionF != nil {
		if _, err := w.w.WriteString(fwm.OriginatorOptionF.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorFI != nil {
		if _, err := w.w.WriteString(fwm.OriginatorFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstructingFI != nil {
		if _, err := w.w.WriteString(fwm.InstructingFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.AccountCreditedDrawdown != nil {
		if _, err := w.w.WriteString(fwm.AccountCreditedDrawdown.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.OriginatorToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.OriginatorToBeneficiary.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeFinancialInstitution(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.FIReceiverFI != nil {
		if _, err := w.w.WriteString(fwm.FIReceiverFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIDrawdownDebitAccountAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIDrawdownDebitAccountAdvice.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFIAdvice.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFIAdvice.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiary.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryAdvice.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIPaymentMethodToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIPaymentMethodToBeneficiary.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.FIAdditionalFIToFI != nil {
		if _, err := w.w.WriteString(fwm.FIAdditionalFIToFI.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeCoverPayment(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	if fwm.CurrencyInstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.CurrencyInstructedAmount.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.OrderingCustomer != nil {
		if _, err := w.w.WriteString(fwm.OrderingCustomer.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.OrderingInstitution != nil {
		if _, err := w.w.WriteString(fwm.OrderingInstitution.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.IntermediaryInstitution != nil {
		if _, err := w.w.WriteString(fwm.IntermediaryInstitution.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.InstitutionAccount != nil {
		if _, err := w.w.WriteString(fwm.InstitutionAccount.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryCustomer != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryCustomer.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.Remittance != nil {
		if _, err := w.w.WriteString(fwm.Remittance.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.SenderToReceiver != nil {
		if _, err := w.w.WriteString(fwm.SenderToReceiver.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}

// options
//  first option : hasVariableLength
//  second option : hasNotNewLine
func (w *Writer) writeRemittance(fwm FEDWireMessage, options ...bool) error {

	hasVariableLength := false
	if len(options) > 0 {
		hasVariableLength = options[0]
	}
	newLine := "\n"
	if len(options) > 1 {
		if options[1] == true {
			newLine = ""
		}
	}

	// Related Remittance
	if fwm.RelatedRemittance != nil {
		if _, err := w.w.WriteString(fwm.RelatedRemittance.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	// Structured Remittance
	if fwm.RemittanceOriginator != nil {
		if _, err := w.w.WriteString(fwm.RemittanceOriginator.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.RemittanceBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.RemittanceBeneficiary.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.PrimaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.PrimaryRemittanceDocument.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.ActualAmountPaid != nil {
		if _, err := w.w.WriteString(fwm.ActualAmountPaid.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.GrossAmountRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.GrossAmountRemittanceDocument.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.AmountNegotiatedDiscount != nil {
		if _, err := w.w.WriteString(fwm.AmountNegotiatedDiscount.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.Adjustment != nil {
		if _, err := w.w.WriteString(fwm.Adjustment.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.DateRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.DateRemittanceDocument.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.SecondaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.SecondaryRemittanceDocument.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	if fwm.RemittanceFreeText != nil {
		if _, err := w.w.WriteString(fwm.RemittanceFreeText.String(hasVariableLength) + newLine); err != nil {
			return err
		}
	}

	return nil
}
