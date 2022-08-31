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

	if err := w.writeMandatory(fwm); err != nil {
		return err
	}

	if err := w.writeOtherTransferInfo(fwm); err != nil {
		return err
	}

	if err := w.writeBeneficiary(fwm); err != nil {
		return err
	}

	if err := w.writeOriginator(fwm); err != nil {
		return err
	}

	if err := w.writeFinancialInstitution(fwm); err != nil {
		return err
	}

	if err := w.writeCoverPayment(fwm); err != nil {
		return err
	}

	if fwm.UnstructuredAddenda != nil {
		if _, err := w.w.WriteString(fwm.UnstructuredAddenda.String() + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if err := w.writeRemittance(fwm); err != nil {
		return err
	}

	if fwm.ServiceMessage != nil {
		if _, err := w.w.WriteString(fwm.ServiceMessage.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if err := w.writeFedAppended(fwm); err != nil {
		return err
	}

	return nil
}

func (w *Writer) writeFedAppended(fwm FEDWireMessage) error {

	if fwm.MessageDisposition != nil {
		if _, err := w.w.WriteString(fwm.MessageDisposition.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.ReceiptTimeStamp != nil {
		if _, err := w.w.WriteString(fwm.ReceiptTimeStamp.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OutputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.OutputMessageAccountabilityData.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.ErrorWire != nil {
		if _, err := w.w.WriteString(fwm.ErrorWire.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeMandatory(fwm FEDWireMessage) error {

	if fwm.SenderSupplied != nil {
		if _, err := w.w.WriteString(fwm.SenderSupplied.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}

	if fwm.TypeSubType != nil {
		if _, err := w.w.WriteString(fwm.TypeSubType.String() + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("TypeSubType", ErrFieldRequired)
	}

	if fwm.InputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.InputMessageAccountabilityData.String() + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}

	if fwm.Amount != nil {
		if _, err := w.w.WriteString(fwm.Amount.String() + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("Amount", ErrFieldRequired)
	}

	if fwm.SenderDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.SenderDepositoryInstitution.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.ReceiverDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.ReceiverDepositoryInstitution.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	if fwm.BusinessFunctionCode != nil {
		if _, err := w.w.WriteString(fwm.BusinessFunctionCode.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}

	return nil
}

func (w *Writer) writeOtherTransferInfo(fwm FEDWireMessage) error {

	if fwm.SenderReference != nil {
		if _, err := w.w.WriteString(fwm.SenderReference.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.PreviousMessageIdentifier != nil {
		if _, err := w.w.WriteString(fwm.PreviousMessageIdentifier.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.LocalInstrument != nil {
		if _, err := w.w.WriteString(fwm.LocalInstrument.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.PaymentNotification != nil {
		if _, err := w.w.WriteString(fwm.PaymentNotification.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.Charges != nil {
		if _, err := w.w.WriteString(fwm.Charges.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.InstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.InstructedAmount.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.ExchangeRate != nil {
		if _, err := w.w.WriteString(fwm.ExchangeRate.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeBeneficiary(fwm FEDWireMessage) error {

	if fwm.BeneficiaryIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryIntermediaryFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryFI != nil {
		if fwm.BeneficiaryFI != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
				return err
			}
		}
	}

	if fwm.Beneficiary != nil {
		if fwm.Beneficiary != nil {
			if _, err := w.w.WriteString(fwm.Beneficiary.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
				return err
			}
		}
	}

	if fwm.BeneficiaryReference != nil {
		if fwm.BeneficiaryReference != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryReference.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
				return err
			}
		}
	}

	if fwm.AccountDebitedDrawdown != nil {
		if fwm.AccountDebitedDrawdown != nil {
			if _, err := w.w.WriteString(fwm.AccountDebitedDrawdown.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *Writer) writeOriginator(fwm FEDWireMessage) error {

	if fwm.Originator != nil {
		if _, err := w.w.WriteString(fwm.Originator.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OriginatorOptionF != nil {
		if _, err := w.w.WriteString(fwm.OriginatorOptionF.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OriginatorFI != nil {
		if _, err := w.w.WriteString(fwm.OriginatorFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.InstructingFI != nil {
		if _, err := w.w.WriteString(fwm.InstructingFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.AccountCreditedDrawdown != nil {
		if _, err := w.w.WriteString(fwm.AccountCreditedDrawdown.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OriginatorToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.OriginatorToBeneficiary.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeFinancialInstitution(fwm FEDWireMessage) error {

	if fwm.FIReceiverFI != nil {
		if _, err := w.w.WriteString(fwm.FIReceiverFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIDrawdownDebitAccountAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIDrawdownDebitAccountAdvice.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIIntermediaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFIAdvice.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFIAdvice.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiary.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIBeneficiaryAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryAdvice.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIPaymentMethodToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIPaymentMethodToBeneficiary.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.FIAdditionalFIToFI != nil {
		if _, err := w.w.WriteString(fwm.FIAdditionalFIToFI.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeCoverPayment(fwm FEDWireMessage) error {

	if fwm.CurrencyInstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.CurrencyInstructedAmount.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OrderingCustomer != nil {
		if _, err := w.w.WriteString(fwm.OrderingCustomer.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.OrderingInstitution != nil {
		if _, err := w.w.WriteString(fwm.OrderingInstitution.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.IntermediaryInstitution != nil {
		if _, err := w.w.WriteString(fwm.IntermediaryInstitution.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.InstitutionAccount != nil {
		if _, err := w.w.WriteString(fwm.InstitutionAccount.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.BeneficiaryCustomer != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryCustomer.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.Remittance != nil {
		if _, err := w.w.WriteString(fwm.Remittance.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.SenderToReceiver != nil {
		if _, err := w.w.WriteString(fwm.SenderToReceiver.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeRemittance(fwm FEDWireMessage) error {

	// Related Remittance
	if fwm.RelatedRemittance != nil {
		if _, err := w.w.WriteString(fwm.RelatedRemittance.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	// Structured Remittance
	if fwm.RemittanceOriginator != nil {
		if _, err := w.w.WriteString(fwm.RemittanceOriginator.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.RemittanceBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.RemittanceBeneficiary.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.PrimaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.PrimaryRemittanceDocument.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.ActualAmountPaid != nil {
		if _, err := w.w.WriteString(fwm.ActualAmountPaid.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.GrossAmountRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.GrossAmountRemittanceDocument.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.AmountNegotiatedDiscount != nil {
		if _, err := w.w.WriteString(fwm.AmountNegotiatedDiscount.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.Adjustment != nil {
		if _, err := w.w.WriteString(fwm.Adjustment.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.DateRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.DateRemittanceDocument.String() + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.SecondaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.SecondaryRemittanceDocument.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	if fwm.RemittanceFreeText != nil {
		if _, err := w.w.WriteString(fwm.RemittanceFreeText.Format(w.FormatOptions) + w.NewlineCharacter); err != nil {
			return err
		}
	}

	return nil
}
