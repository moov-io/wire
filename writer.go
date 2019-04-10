// Copyright 2019 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"io"
)

// A Writer writes an x9.file to an encoded file.
//
// As returned by NewWriter, a Writer writes x9file structs into
// x9 formatted files.

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

// Writer writes a single x9.file record to w
func (w *Writer) Write(file *File) error {
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeFedWireMessage(file); err != nil {
		return err
	}
	w.lineNum++

	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

func (w *Writer) writeFedWireMessage(file *File) error {
		fwm := file.FedWireMessage
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
		if _, err := w.w.WriteString(fwm.GetUnstructuredAddenda().String() + "\n"); err != nil {
			return err
		}
		if err := w.writeRemittance(fwm); err != nil {
			return err
		}
		if _, err := w.w.WriteString(fwm.GetServiceMessage().String() + "\n"); err != nil {
			return err
		}
	return nil
}


func (w *Writer) writeMandatory(fwm FedWireMessage) error {
	if fwm.SenderSupplied != nil {
		if _, err := w.w.WriteString(fwm.GetSenderSupplied().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	if fwm.TypeSubType != nil {
		if _, err := w.w.WriteString(fwm.GetTypeSubType().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	if fwm.InputMessageAccountabilityData != nil {
		if _, err := w.w.WriteString(fwm.GetInputMessageAccountabilityData().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	if fwm.Amount != nil {
		if _, err := w.w.WriteString(fwm.GetAmount().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	if fwm.SenderDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.GetSenderDepositoryInstitution().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}

	if fwm.ReceiverDepositoryInstitution != nil {
		if _, err := w.w.WriteString(fwm.GetReceiverDepositoryInstitution().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	if fwm.BusinessFunctionCode != nil {
		if _, err := w.w.WriteString(fwm.GetBusinessFunctionCode().String() + "\n"); err != nil {
			return err
		}
		// ToDo; Return error
	}
	return nil
}

func (w *Writer) writeOtherTransferInfo(fwm FedWireMessage) error {
	if _, err := w.w.WriteString(fwm.GetSenderReference().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetPreviousMessageIdentifier().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetLocalInstrument().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetPaymentNotification().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetCharges().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetInstructedAmount().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetExchangeRate().String() + "\n"); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeBeneficiary(fwm FedWireMessage) error {
	if _, err := w.w.WriteString(fwm.GetBeneficiaryIntermediaryFI().String() + "\n"); err != nil {
		return err
	}

	if _, err := w.w.WriteString(fwm.GetBeneficiaryFI().String() + "\n"); err != nil {
		return err
	}

	if _, err := w.w.WriteString(fwm.GetBeneficiary().String() + "\n"); err != nil {
		return err
	}

	if _, err := w.w.WriteString(fwm.GetBeneficiaryReference().String() + "\n"); err != nil {
		return err
	}

	if _, err := w.w.WriteString(fwm.GetAccountDebitedDrawdown().String() + "\n"); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeOriginator(fwm FedWireMessage) error {
	if _, err := w.w.WriteString(fwm.GetAccountDebitedDrawdown().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOriginator().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOriginatorOptionF().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOriginatorFI().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetInstructingFI().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetAccountCreditedDrawdown().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOriginatorToBeneficiary().String() + "\n"); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeFinancialInstitution(fwm FedWireMessage) error {
	if _, err := w.w.WriteString(fwm.GetFIReceiverFI().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIDrawdownDebitAccountAdvice().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIIntermediaryFI().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIIntermediaryFIAdvice().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIBeneficiaryFI().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIBeneficiaryFIAdvice().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIBeneficiary().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIBeneficiaryAdvice().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIPaymentMethodToBeneficiary().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetFIAdditionalFIToFI().String() + "\n"); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeCoverPayment(fwm FedWireMessage) error {
	if _, err := w.w.WriteString(fwm.GetCurrencyInstructedAmount().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOrderingCustomer().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetOrderingInstitution().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetIntermediaryInstitution().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetInstitutionAccount().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetBeneficiaryCustomer().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetRemittance().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetSenderToReceiver().String() + "\n"); err != nil {
		return err
	}
	return nil
}

func (w *Writer) writeRemittance(fwm FedWireMessage) error {

	// Related Remittance

	if _, err := w.w.WriteString(fwm.GetRelatedRemittance().String() + "\n"); err != nil {
		return err
	}

	// Structured Remittance

	if _, err := w.w.WriteString(fwm.GetRemittanceOriginator().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetRemittanceBeneficiary().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetPrimaryRemittanceDocument().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetActualAmountPaid().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetGrossAmountRemittanceDocument().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetAmountNegotiatedDiscount().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetAdjustment().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetDateRemittanceDocument().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetSecondaryRemittanceDocument().String() + "\n"); err != nil {
		return err
	}
	if _, err := w.w.WriteString(fwm.GetRemittanceFreeText().String() + "\n"); err != nil {
		return err
	}
	return nil
}