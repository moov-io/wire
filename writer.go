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
		if _, err := w.w.WriteString(fwm.UnstructuredAddenda.String() + "\n"); err != nil {
			return err
		}
	}
	if err := w.writeRemittance(fwm); err != nil {
		return err
	}
	if fwm.ServiceMessage != nil {
		if _, err := w.w.WriteString(fwm.ServiceMessage.String() + "\n"); err != nil {
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
		if _, err := w.w.WriteString(fwm.MessageDisposition.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.ReceiptTimeStamp != nil {
		if _, err := w.w.WriteString(fwm.ReceiptTimeStamp.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OutputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.OutputMessageAccountabilityData.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.ErrorWire != nil {
		if _, err := w.w.WriteString(fwm.ErrorWire.String() + "\n"); err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) writeMandatory(fwm FEDWireMessage) error {
	if fwm.SenderSupplied != nil {
		if _, err := w.w.WriteString(fwm.SenderSupplied.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}

	if fwm.TypeSubType != nil {
		if _, err := w.w.WriteString(fwm.TypeSubType.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("TypeSubType", ErrFieldRequired)
	}
	if fwm.InputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.InputMessageAccountabilityData.String() + "\n"); err != nil {
			return err
		}
	} else if fwm.validateOpts.AllowMissingIMAD == false { // IMAD is optional for FedLine Advantage customers
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}
	if fwm.Amount != nil {
		if _, err := w.w.WriteString(fwm.Amount.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("Amount", ErrFieldRequired)
	}
	if fwm.SenderDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.SenderDepositoryInstitution.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}
	if fwm.ReceiverDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.ReceiverDepositoryInstitution.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}
	if fwm.BusinessFunctionCode != nil {
		if _, err := w.w.WriteString(fwm.BusinessFunctionCode.String() + "\n"); err != nil {
			return err
		}
	} else {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}
	return nil
}

func (w *Writer) writeOtherTransferInfo(fwm FEDWireMessage) error {
	if fwm.SenderReference != nil {
		if _, err := w.w.WriteString(fwm.SenderReference.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.PreviousMessageIdentifier != nil {
		if _, err := w.w.WriteString(fwm.PreviousMessageIdentifier.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.LocalInstrument != nil {
		if _, err := w.w.WriteString(fwm.LocalInstrument.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.PaymentNotification != nil {
		if _, err := w.w.WriteString(fwm.PaymentNotification.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.Charges != nil {
		if _, err := w.w.WriteString(fwm.Charges.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.InstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.InstructedAmount.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.ExchangeRate != nil {
		if _, err := w.w.WriteString(fwm.ExchangeRate.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeBeneficiary(fwm FEDWireMessage) error {
	if fwm.BeneficiaryIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryIntermediaryFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.BeneficiaryFI != nil {
		if fwm.BeneficiaryFI != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryFI.String() + "\n"); err != nil {
				return err
			}
		}
	}
	if fwm.Beneficiary != nil {
		if fwm.Beneficiary != nil {
			if _, err := w.w.WriteString(fwm.Beneficiary.String() + "\n"); err != nil {
				return err
			}
		}
	}
	if fwm.BeneficiaryReference != nil {
		if fwm.BeneficiaryReference != nil {
			if _, err := w.w.WriteString(fwm.BeneficiaryReference.String() + "\n"); err != nil {
				return err
			}
		}
	}
	if fwm.AccountDebitedDrawdown != nil {
		if fwm.AccountDebitedDrawdown != nil {
			if _, err := w.w.WriteString(fwm.AccountDebitedDrawdown.String() + "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *Writer) writeOriginator(fwm FEDWireMessage) error {
	if fwm.Originator != nil {
		if _, err := w.w.WriteString(fwm.Originator.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OriginatorOptionF != nil {
		if _, err := w.w.WriteString(fwm.OriginatorOptionF.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OriginatorFI != nil {
		if _, err := w.w.WriteString(fwm.OriginatorFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.InstructingFI != nil {
		if _, err := w.w.WriteString(fwm.InstructingFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.AccountCreditedDrawdown != nil {
		if _, err := w.w.WriteString(fwm.AccountCreditedDrawdown.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OriginatorToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.OriginatorToBeneficiary.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeFinancialInstitution(fwm FEDWireMessage) error {
	if fwm.FIReceiverFI != nil {
		if _, err := w.w.WriteString(fwm.FIReceiverFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIDrawdownDebitAccountAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIDrawdownDebitAccountAdvice.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIIntermediaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIIntermediaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIIntermediaryFIAdvice.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryFI != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFI.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryFIAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryFIAdvice.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiary.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIBeneficiaryAdvice != nil {
		if _, err := w.w.WriteString(fwm.FIBeneficiaryAdvice.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIPaymentMethodToBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.FIPaymentMethodToBeneficiary.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.FIAdditionalFIToFI != nil {
		if _, err := w.w.WriteString(fwm.FIAdditionalFIToFI.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeCoverPayment(fwm FEDWireMessage) error {
	if fwm.CurrencyInstructedAmount != nil {
		if _, err := w.w.WriteString(fwm.CurrencyInstructedAmount.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OrderingCustomer != nil {
		if _, err := w.w.WriteString(fwm.OrderingCustomer.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.OrderingInstitution != nil {
		if _, err := w.w.WriteString(fwm.OrderingInstitution.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.IntermediaryInstitution != nil {
		if _, err := w.w.WriteString(fwm.IntermediaryInstitution.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.InstitutionAccount != nil {
		if _, err := w.w.WriteString(fwm.InstitutionAccount.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.BeneficiaryCustomer != nil {
		if _, err := w.w.WriteString(fwm.BeneficiaryCustomer.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.Remittance != nil {
		if _, err := w.w.WriteString(fwm.Remittance.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.SenderToReceiver != nil {
		if _, err := w.w.WriteString(fwm.SenderToReceiver.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (w *Writer) writeRemittance(fwm FEDWireMessage) error {

	// Related Remittance
	if fwm.RelatedRemittance != nil {
		if _, err := w.w.WriteString(fwm.RelatedRemittance.String() + "\n"); err != nil {
			return err
		}
	}
	// Structured Remittance
	if fwm.RemittanceOriginator != nil {
		if _, err := w.w.WriteString(fwm.RemittanceOriginator.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.RemittanceBeneficiary != nil {
		if _, err := w.w.WriteString(fwm.RemittanceBeneficiary.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.PrimaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.PrimaryRemittanceDocument.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.ActualAmountPaid != nil {
		if _, err := w.w.WriteString(fwm.ActualAmountPaid.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.GrossAmountRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.GrossAmountRemittanceDocument.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.AmountNegotiatedDiscount != nil {
		if _, err := w.w.WriteString(fwm.AmountNegotiatedDiscount.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.Adjustment != nil {
		if _, err := w.w.WriteString(fwm.Adjustment.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.DateRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.DateRemittanceDocument.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.SecondaryRemittanceDocument != nil {
		if _, err := w.w.WriteString(fwm.SecondaryRemittanceDocument.String() + "\n"); err != nil {
			return err
		}
	}
	if fwm.RemittanceFreeText != nil {
		if _, err := w.w.WriteString(fwm.RemittanceFreeText.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}
