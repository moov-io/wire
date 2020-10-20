// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// FEDWireMessage is a FedWire Message
type FEDWireMessage struct {
	// ID
	ID string `json:"id"`
	// MessageDisposition
	MessageDisposition *MessageDisposition `json:"messageDisposition,omitempty"`
	// ReceiptTimeStamp
	ReceiptTimeStamp *ReceiptTimeStamp `json:"receiptTimeStamp,omitempty"`
	// OutputMessageAccountabilityData (OMAD)
	OutputMessageAccountabilityData *OutputMessageAccountabilityData `json:"outputMessageAccountabilityData,omitempty"`
	// ErrorWire
	ErrorWire *ErrorWire `json:"errorWire,omitempty"`
	// SenderSuppliedInformation
	SenderSupplied *SenderSupplied `json:"senderSupplied"`
	// TypeSubType
	TypeSubType *TypeSubType `json:"typeSubType"`
	// InputMessageAccountabilityData (IMAD)
	InputMessageAccountabilityData *InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	// Amount (up to a penny less than $10 billion)
	Amount *Amount `json:"amount"`
	// SenderDepositoryInstitution
	SenderDepositoryInstitution *SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	// ReceiverDepositoryInstitution
	ReceiverDepositoryInstitution *ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	// BusinessFunctionCode
	BusinessFunctionCode *BusinessFunctionCode `json:"businessFunctionCode"`
	// SenderReference
	SenderReference *SenderReference `json:"senderReference,omitempty"`
	// PreviousMessageIdentifier
	PreviousMessageIdentifier *PreviousMessageIdentifier `json:"previousMessageIdentifier,omitempty"`
	// LocalInstrument
	LocalInstrument *LocalInstrument `json:"localInstrument,omitempty"`
	// PaymentNotification
	PaymentNotification *PaymentNotification `json:"paymentNotification,omitempty"`
	// Charges
	Charges *Charges `json:"charges,omitempty"`
	// InstructedAmount
	InstructedAmount *InstructedAmount `json:"instructedAmount,omitempty"`
	// ExchangeRate
	ExchangeRate *ExchangeRate `json:"exchangeRate,omitempty"`
	// BeneficiaryIntermediaryFI
	BeneficiaryIntermediaryFI *BeneficiaryIntermediaryFI `json:"beneficiaryIntermediaryFI,omitempty"`
	// BeneficiaryFI
	BeneficiaryFI *BeneficiaryFI `json:"beneficiaryFI,omitempty"`
	// Beneficiary
	Beneficiary *Beneficiary `json:"beneficiary,omitempty"`
	// BeneficiaryReference
	BeneficiaryReference *BeneficiaryReference `json:"beneficiaryReference,omitempty"`
	// AccountDebitedDrawdown
	AccountDebitedDrawdown *AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	// Originator
	Originator *Originator `json:"originator,omitempty"`
	// OriginatorOptionF
	OriginatorOptionF *OriginatorOptionF `json:"originatorOptionF,omitempty"`
	// OriginatorFI
	OriginatorFI *OriginatorFI `json:"originatorFI,omitempty"`
	// InstructingFI
	InstructingFI *InstructingFI `json:"instructingFI,omitempty"`
	// AccountCreditedDrawdown
	AccountCreditedDrawdown *AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	// OriginatorToBeneficiary
	OriginatorToBeneficiary *OriginatorToBeneficiary `json:"originatorToBeneficiary,omitempty"`
	// FIReceiverFI
	FIReceiverFI *FIReceiverFI `json:"fiReceiverFI,omitempty"`
	// FIDrawdownDebitAccountAdvice
	FIDrawdownDebitAccountAdvice *FIDrawdownDebitAccountAdvice `json:"fiDrawdownDebitAccountAdvice,omitempty"`
	// FIIntermediaryFI
	FIIntermediaryFI *FIIntermediaryFI `json:"fiIntermediaryFI,omitempty"`
	// FIIntermediaryFIAdvice
	FIIntermediaryFIAdvice *FIIntermediaryFIAdvice `json:"fiIntermediaryFIAdvice,omitempty"`
	// FIBeneficiaryFI
	FIBeneficiaryFI *FIBeneficiaryFI `json:"fiBeneficiaryFI,omitempty"`
	// FIBeneficiaryFIAdvice
	FIBeneficiaryFIAdvice *FIBeneficiaryFIAdvice `json:"fiBeneficiaryFIAdvice,omitempty"`
	// FIBeneficiary

	FIBeneficiary *FIBeneficiary `json:"fiBeneficiary,omitempty"`
	// FIBeneficiaryAdvice
	FIBeneficiaryAdvice *FIBeneficiaryAdvice `json:"fiBeneficiaryAdvice,omitempty"`
	// FIPaymentMethodToBeneficiary
	FIPaymentMethodToBeneficiary *FIPaymentMethodToBeneficiary `json:"fiPaymentMethodToBeneficiary,omitempty"`
	// FIAdditionalFIToFI
	FIAdditionalFIToFI *FIAdditionalFIToFI `json:"fiAdditionalFiToFi,omitempty"`
	// CurrencyInstructedAmount
	CurrencyInstructedAmount *CurrencyInstructedAmount `json:"currencyInstructedAmount,omitempty"`
	// OrderingCustomer
	OrderingCustomer *OrderingCustomer `json:"orderingCustomer,omitempty"`
	// OrderingInstitution
	OrderingInstitution *OrderingInstitution `json:"orderingInstitution,omitempty"`
	// IntermediaryInstitution
	IntermediaryInstitution *IntermediaryInstitution `json:"intermediaryInstitution,omitempty"`
	// InstitutionAccount
	InstitutionAccount *InstitutionAccount `json:"institutionAccount,omitempty"`
	// BeneficiaryCustomer
	BeneficiaryCustomer *BeneficiaryCustomer `json:"beneficiaryCustomer,omitempty"`
	// Remittance
	Remittance *Remittance `json:"remittance,omitempty"`
	// SenderToReceiver
	SenderToReceiver *SenderToReceiver `json:"senderToReceiver,omitempty"`
	// UnstructuredAddenda
	UnstructuredAddenda *UnstructuredAddenda `json:"unstructuredAddenda,omitempty"`
	// RelatedRemittance
	RelatedRemittance *RelatedRemittance `json:"relatedRemittance,omitempty"`
	// RemittanceOriginator
	RemittanceOriginator *RemittanceOriginator `json:"remittanceOriginator,omitempty"`
	// RemittanceBeneficiary
	RemittanceBeneficiary *RemittanceBeneficiary `json:"remittanceBeneficiary,omitempty"`
	// PrimaryRemittanceDocument
	PrimaryRemittanceDocument *PrimaryRemittanceDocument `json:"primaryRemittanceDocument,omitempty"`
	// ActualAmountPaid
	ActualAmountPaid *ActualAmountPaid `json:"actualAmountPaid,omitempty"`
	// GrossAmountRemittanceDocument
	GrossAmountRemittanceDocument *GrossAmountRemittanceDocument `json:"grossAmountRemittanceDocument,omitempty"`
	// AmountNegotiatedDiscount
	AmountNegotiatedDiscount *AmountNegotiatedDiscount `json:"amountNegotiatedDiscount,omitempty"`
	// Adjustment
	Adjustment *Adjustment `json:"adjustment,omitempty"`
	// DateRemittanceDocument
	DateRemittanceDocument *DateRemittanceDocument `json:"dateRemittanceDocument,omitempty"`
	// SecondaryRemittanceDocument
	SecondaryRemittanceDocument *SecondaryRemittanceDocument `json:"secondaryRemittanceDocument,omitempty"`
	// RemittanceFreeText
	RemittanceFreeText *RemittanceFreeText `json:"remittanceFreeText,omitempty"`
	// ServiceMessage
	ServiceMessage *ServiceMessage `json:"serviceMessage,omitempty"`
}

// NewFEDWireMessage returns a new FEDWireMessage
func NewFEDWireMessage() FEDWireMessage {
	fwm := FEDWireMessage{}
	return fwm
}

// verify checks basic WIRE rules. Assumes properly parsed records. Each validation func should
// check for the expected relationships between fields within a FedWireMessage.
func (fwm *FEDWireMessage) verify() error {

	if err := fwm.mandatoryFields(); err != nil {
		return err
	}

	if err := fwm.otherTransferInformation(); err != nil {
		return err
	}
	if err := fwm.validateBeneficiaryIntermediaryFI(); err != nil {
		return err
	}
	if err := fwm.validateBeneficiaryFI(); err != nil {
		return err
	}
	if err := fwm.validateOriginatorFI(); err != nil {
		return err
	}
	if err := fwm.validateInstructingFI(); err != nil {
		return err
	}
	if err := fwm.validateOriginatorToBeneficiary(); err != nil {
		return err
	}
	if err := fwm.validateFIIntermediaryFI(); err != nil {
		return err
	}
	if err := fwm.validateFIIntermediaryFIAdvice(); err != nil {
		return err
	}
	if err := fwm.validateFIBeneficiaryFI(); err != nil {
		return err
	}
	if err := fwm.validateFIBeneficiaryFIAdvice(); err != nil {
		return err
	}
	if err := fwm.validateFIBeneficiary(); err != nil {
		return err
	}
	if err := fwm.validateFIBeneficiaryAdvice(); err != nil {
		return err
	}
	if err := fwm.validateFIPaymentMethodToBeneficiary(); err != nil {
		return err
	}
	if err := fwm.validateUnstructuredAddenda(); err != nil {
		return err
	}

	if err := fwm.validateRelatedRemittance(); err != nil {
		return err
	}

	if err := fwm.isRemittanceValid(); err != nil {
		return err
	}
	return nil
}

// mandatoryFields validates mandatory tags for a FEDWireMessage are defined
func (fwm *FEDWireMessage) mandatoryFields() error {
	if err := fwm.validateSenderSupplied(); err != nil {
		return err
	}
	if err := fwm.validateTypeSubType(); err != nil {
		return err
	}
	if err := fwm.validateIMAD(); err != nil {
		return err
	}
	if err := fwm.validateAmount(); err != nil {
		return err
	}
	if err := fwm.validateSenderDI(); err != nil {
		return err
	}
	if err := fwm.validateReceiverDI(); err != nil {
		return err
	}
	if err := fwm.validateBusinessFunctionCode(); err != nil {
		return err
	}
	return nil
}

// validateSenderSupplied validates TagSenderSupplied within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateSenderSupplied() error {
	if fwm.SenderSupplied == nil {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}
	return nil
}

// validateTypeSubType validates TagTypeSubType within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateTypeSubType() error {
	if fwm.TypeSubType == nil {
		return fieldError("TypeSubType", ErrFieldRequired)
	}
	return nil
}

// validateIMAD validates TagInputMessageAccountabilityData within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateIMAD() error {
	if fwm.InputMessageAccountabilityData == nil {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}
	return nil
}

// validateAmount validates TagAmount within a FEDWireMessage
// * Mandatory for all requests
// * Can be all zeros for TypeSubType code 90
func (fwm *FEDWireMessage) validateAmount() error {
	if fwm.Amount == nil {
		return fieldError("Amount", ErrFieldRequired)
	}
	if fwm.Amount.Amount == "000000000000" && fwm.TypeSubType.SubTypeCode != "90" {
		return NewErrInvalidPropertyForProperty("Amount", fwm.Amount.Amount,
			"SubTypeCode", fwm.TypeSubType.SubTypeCode)
	}
	return nil
}

// validateSenderDI validates TagSenderDepositoryInstitution within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateSenderDI() error {
	if fwm.SenderDepositoryInstitution == nil {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}
	return nil
}

// validateReceiverDI validates TagReceiverDepositoryInstitution within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateReceiverDI() error {
	if fwm.ReceiverDepositoryInstitution == nil {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}
	return nil
}

// validateBusinessFunctionCode validates TagBusinessFunctionCode within a FEDWireMessage
// Mandatory for all requests
func (fwm *FEDWireMessage) validateBusinessFunctionCode() error {
	if fwm.BusinessFunctionCode == nil {
		return fieldError("BusinessFunctionCode", ErrFieldRequired)
	}

	switch fwm.BusinessFunctionCode.BusinessFunctionCode {
	case BankTransfer:
		if err := fwm.validateBankTransfer(); err != nil {
			return err
		}
	case CustomerTransfer:
		if err := fwm.validateCustomerTransfer(); err != nil {
			return err
		}
	case CustomerTransferPlus:
		if err := fwm.validateCustomerTransferPlus(); err != nil {
			return err
		}
	case CheckSameDaySettlement:
		if err := fwm.validateCheckSameDaySettlement(); err != nil {
			return err
		}
	case DepositSendersAccount:
		if err := fwm.validateDepositSendersAccount(); err != nil {
			return err
		}
	case FEDFundsReturned:
		if err := fwm.validateFEDFundsReturned(); err != nil {
			return err
		}
	case FEDFundsSold:
		if err := fwm.validateFEDFundsSold(); err != nil {
			return err
		}
	case DrawdownResponse:
		if err := fwm.validateDrawdownResponse(); err != nil {
			return err
		}
	case BankDrawDownRequest:
		if err := fwm.validateBankDrawdownRequest(); err != nil {
			return err
		}
	case CustomerCorporateDrawdownRequest:
		if err := fwm.validateCustomerCorporateDrawdownRequest(); err != nil {
			return err
		}
	case BFCServiceMessage:
		if err := fwm.validateServiceMessage(); err != nil {
			return err
		}
	}
	return nil
}

// validateBankTransfer validates the BankTransfer code and associated tags
// Requires the standard "mandatory" tags checked in mandatoryFields
// If TypeSubType is ReversalTransfer or ReversalPriorDayTransfer, then PreviousMessageIdentifier is mandatory.
func (fwm *FEDWireMessage) validateBankTransfer() error {
	if err := fwm.checkProhibitedBankTransferTags(); err != nil {
		return err
	}
	if err := fwm.checkPreviousMessageIdentifier(); err != nil {
		return err
	}

	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !btrTypeSubTypes.Contains(typeSubType) {
		return NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode)
	}

	return nil
}

// checkProhibitedBankTransferTags ensures there are no tags present in the message that are incompatible with the BankTransfer code
// Tags NOT permitted:
//   BusinessFunctionCode Element 02, LocalInstrument, PaymentNotification, Charges, InstructedAmount, ExchangeRate,
//   Beneficiary Code SWIFTBICORBEIANDAccountNumber, AccountDebitedDrawdown, Originator Code SWIFTBICORBEIANDAccountNumber,
//   OriginatorOptionF, AccountCreditedDrawdown, FIDrawdownDebitAccountAdvice, Any CoverPayment Information tag ({7xxx}),
//   Any UnstructuredAddenda or remittance tags ({8xxx}), and ServiceMessage
func (fwm *FEDWireMessage) checkProhibitedBankTransferTags() error {
	if fwm.BusinessFunctionCode != nil {
		if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
			return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
		}
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.Charges != nil {
		return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
	}
	if fwm.InstructedAmount != nil {
		return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
	}
	if fwm.ExchangeRate != nil {
		return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
	}
	if fwm.Beneficiary != nil && fwm.Beneficiary.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.Originator != nil && fwm.Originator.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIDrawdownDebitAccountAdvice != nil {
		return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
	}
	if fwm.ServiceMessage != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// validateCustomerTransfer validates the CustomerTransfer business function code
func (fwm *FEDWireMessage) validateCustomerTransfer() error {
	if err := fwm.checkMandatoryCustomerTransferTags(); err != nil {
		return err
	}
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !ctrTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// checkMandatoryCustomerTransferTags checks for the tags required by CustomerTransfer in addition to the standard mandatoryFields.
// Additional mandatory tags: Beneficiary, Originator
// If TypeSubType = ReversalTransfer or ReversalPriorDayTransfer, then PreviousMessageIdentifier is mandatory.
func (fwm *FEDWireMessage) checkMandatoryCustomerTransferTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil {
		return fieldError("Originator", ErrFieldRequired)
	}
	if err := fwm.checkPreviousMessageIdentifier(); err != nil {
		return err
	}
	return nil
}

// checkProhibitedCustomerTransferTags ensures there are no tags present in the message that are incompatible with the CustomerTransfer code
// Tags NOT permitted:
//   BusinessFunctionCode Element 02 = COV, LocalInstrument, PaymentNotification, AccountDebitedDrawdown, OriginatorOptionF, AccountCreditedDrawdown,
//   FIDrawdownDebitAccountAdvice, any CoverPayment Information tag ({7xxx}), any UnstructuredAddenda or remittance tags ({8xxx}) and ServiceMessage
func (fwm *FEDWireMessage) checkProhibitedCustomerTransferTags() error {
	// This covers the edit requirement
	if fwm.BusinessFunctionCode.TransactionTypeCode == "COV" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIDrawdownDebitAccountAdvice != nil {
		return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
	}
	if fwm.ServiceMessage != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// validateCustomerTransferPlus validates the CustomerTransferPlus business function code
func (fwm *FEDWireMessage) validateCustomerTransferPlus() error {
	if err := fwm.checkMandatoryCustomerTransferPlusTags(); err != nil {
		return err
	}
	if err := fwm.checkProhibitedCustomerTransferPlusTags(); err != nil {
		return err
	}
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !ctpTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return nil
}

// checkMandatoryCustomerTransferPlusTags checks for the tags required by CustomerTransferPlus in addition to the standard mandatoryFields
// Additional mandatory fields:
//   Beneficiary and Originator OR OriginatorOptionF
// If TypeSubType = ReversalTransfer or ReversalPriorDayTransfer, then PreviousMessageIdentifier is mandatory.
// If LocalInstrument = SequenceBCoverPaymentStructured, then BeneficiaryReference, OrderingCustomer & BeneficiaryCustomer are mandatory.
// If LocalInstrument = ANSIX12format, GeneralXMLformat, ISO20022XMLformat, NarrativeText, STP820format, SWIFTfield70 or UNEDIFACTformat, then UnstructuredAddenda is mandatory.
// If LocalInstrument = RelatedRemittanceInformation, then RelatedRemittance is mandatory.
// If LocalInstrument = RemittanceInformationStructured, then RemittanceOriginator, RemittanceBeneficiary, PrimaryRemittanceDocument & ActualAmountPaid are mandatory.
// If LocalInstrument = ProprietaryLocalInstrumentCode, then LocalInstrument Element 02 is mandatory.
func (fwm *FEDWireMessage) checkMandatoryCustomerTransferPlusTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil && fwm.OriginatorOptionF == nil { // one or the other must be present
		return fieldError("Originator OR OriginatorOptionF", ErrFieldRequired)
	}
	if err := fwm.checkPreviousMessageIdentifier(); err != nil {
		return err
	}

	switch fwm.LocalInstrument.LocalInstrumentCode {
	case SequenceBCoverPaymentStructured:
		if fwm.BeneficiaryReference == nil {
			return fieldError("Beneficiary Reference", ErrFieldRequired)
		}
		if fwm.OrderingCustomer == nil {
			return fieldError("Ordering Customer", ErrFieldRequired)
		}
		if fwm.BeneficiaryCustomer == nil {
			return fieldError("BeneficiaryCustomer", ErrFieldRequired)
		}
	case ANSIX12format, GeneralXMLformat, ISO20022XMLformat,
		NarrativeText, STP820format, SWIFTfield70, UNEDIFACTformat:
		if fwm.UnstructuredAddenda == nil {
			return fieldError("UnstructuredAddenda", ErrFieldRequired)
		}
	case RelatedRemittanceInformation:
		if fwm.RelatedRemittance == nil {
			return fieldError("RelatedRemittance", ErrFieldRequired)
		}
	case RemittanceInformationStructured:
		if fwm.RemittanceOriginator == nil {
			return fieldError("RemittanceOriginator", ErrFieldRequired)
		}
		if fwm.RemittanceBeneficiary == nil {
			return fieldError("RemittanceBeneficiary", ErrFieldRequired)
		}
		if fwm.PrimaryRemittanceDocument == nil {
			return fieldError("PrimaryRemittanceDocument", ErrFieldRequired)
		}
		if fwm.ActualAmountPaid == nil {
			return fieldError("ActualAmountPaid", ErrFieldRequired)
		}
	case ProprietaryLocalInstrumentCode:
		if fwm.LocalInstrument.ProprietaryCode == "" {
			return fieldError("ProprietaryCode", ErrFieldRequired)
		}
	}
	return nil
}

// checkProhibitedCustomerTransferPlusTags ensures there are no tags present in the message that are incompatible with the CustomerTransferPlus code
// Tags NOT permitted:
//   BusinessFunctionCode.TransactionTypeCode, AccountDebitedDrawdown, AccountCreditedDrawdown, FIDrawdownDebitAccountAdvice, ServiceMessage
// If LocalInstrument = SequenceBCoverPaymentStructured, Charges, InstructedAmount & ExchangeRate are not permitted.
// Certain {7xxx} tags & {8xxx} tags may not be permitted depending upon value of LocalInstrument.
func (fwm *FEDWireMessage) checkProhibitedCustomerTransferPlusTags() error {
	if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.AccountDebitedDrawdown != nil {
		return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
	}
	if fwm.AccountCreditedDrawdown != nil {
		return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
	}
	if fwm.FIReceiverFI != nil {
		return fieldError("FIReceiverFI", ErrInvalidProperty, fwm.FIReceiverFI)
	}

	if fwm.LocalInstrument != nil {
		if fwm.LocalInstrument.LocalInstrumentCode == SequenceBCoverPaymentStructured {
			if fwm.Charges != nil {
				return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
			}
			if fwm.InstructedAmount != nil {
				return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
			}
			if fwm.ExchangeRate != nil {
				return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
			}
		}
		if fwm.LocalInstrument.LocalInstrumentCode != SequenceBCoverPaymentStructured {
			if err := fwm.invalidCoverPaymentTags(); err != nil {
				return err
			}
		}
	}

	// ToDo: From the spec - Certain {7xxx} tags & {8xxx} tags may not be permitted depending upon value of {3610}.  I'm not sure how to code this yet
	return nil
}

// checkPreviousMessageIdentifier returns an error if ReversalTransfer or ReversalPriorDayTransfer options are set and PreviousMessageIdentifier is missing
func (fwm *FEDWireMessage) checkPreviousMessageIdentifier() error {
	if fwm.TypeSubType == nil || fwm.BusinessFunctionCode == nil {
		return nil
	}

	switch fwm.TypeSubType.SubTypeCode {
	case ReversalTransfer, ReversalPriorDayTransfer:
		if fwm.PreviousMessageIdentifier == nil {
			return fieldError("PreviousMessageIdentifier", ErrFieldRequired)
		}
	}
	return nil
}

// validateCheckSameDaySettlement validates the CheckSameDaySettlement business function code
func (fwm *FEDWireMessage) validateCheckSameDaySettlement() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !cksTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return fwm.checkSharedProhibitedTags()
}

// validateDepositSendersAccount validates the DepositSendersAccount business function code
func (fwm *FEDWireMessage) validateDepositSendersAccount() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !depTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return fwm.checkSharedProhibitedTags()
}

// validateFEDFundsReturned validates the FEDFundsReturned business function code
func (fwm *FEDWireMessage) validateFEDFundsReturned() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !ffrTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return fwm.checkSharedProhibitedTags()
}

// validateFEDFundsSold validates the FEDFundsSold business function code
func (fwm *FEDWireMessage) validateFEDFundsSold() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !ffsTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	return fwm.checkSharedProhibitedTags()
}

// validateDrawdownResponse validates the DrawdownResponse business function code
func (fwm *FEDWireMessage) validateDrawdownResponse() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !drwTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	if err := fwm.checkMandatoryDrawdownResponseTags(); err != nil {
		return err
	}
	return fwm.checkSharedProhibitedTags()
}

// checkMandatoryDrawdownResponseTags checks for the tags required by DrawdownResponse in addition to the standard mandatoryFields
// Additional mandatory fields: Beneficiary, Originator
func (fwm *FEDWireMessage) checkMandatoryDrawdownResponseTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.Originator == nil {
		return fieldError("Originator", ErrFieldRequired)
	}
	return nil
}

// validateBankDrawdownRequest validates the BankDrawDownRequest business function code
func (fwm *FEDWireMessage) validateBankDrawdownRequest() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !drbTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	if err := fwm.checkMandatoryBankDrawdownRequestTags(); err != nil {
		return err
	}
	return fwm.checkSharedProhibitedTags()
}

// checkMandatoryBankDrawdownRequestTags checks for the tags required by BankDrawDownRequest in addition to the standard mandatoryFields
// Additional mandatory fields: AccountDebitedDrawdown, AccountCreditedDrawdown
func (fwm *FEDWireMessage) checkMandatoryBankDrawdownRequestTags() error {
	if fwm.AccountDebitedDrawdown == nil {
		return fieldError("AccountDebitedDrawdown", ErrFieldRequired)
	}
	if fwm.AccountCreditedDrawdown == nil {
		return fieldError("AccountCreditedDrawdown", ErrFieldRequired)
	}
	return nil
}

// validateCustomerCorporateDrawdownRequest validates the CustomerCorporateDrawdownRequest business function code
func (fwm *FEDWireMessage) validateCustomerCorporateDrawdownRequest() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !drcTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	if err := fwm.checkMandatoryCustomerCorporateDrawdownRequestTags(); err != nil {
		return err
	}
	return fwm.checkSharedProhibitedTags()
}

// checkMandatoryCustomerCorporateDrawdownRequestTags checks for the tags required by CustomerCorporateDrawdownRequest in addition to the standard mandatoryFields
// Additional mandatory fields: Beneficiary, AccountDebitedDrawdown, AccountCreditedDrawdown
func (fwm *FEDWireMessage) checkMandatoryCustomerCorporateDrawdownRequestTags() error {
	if fwm.Beneficiary == nil {
		return fieldError("Beneficiary", ErrFieldRequired)
	}
	if fwm.AccountDebitedDrawdown == nil {
		return fieldError("AccountDebitedDrawdown", ErrFieldRequired)
	}
	if fwm.AccountCreditedDrawdown == nil {
		return fieldError("AccountCreditedDrawdown", ErrFieldRequired)
	}
	return nil
}

// validateServiceMessage validates the BFCServiceMessage business function code
func (fwm *FEDWireMessage) validateServiceMessage() error {
	typeSubType := fwm.TypeSubType.TypeCode + fwm.TypeSubType.SubTypeCode
	if !svcTypeSubTypes.Contains(typeSubType) {
		return fieldError("TypeSubType", NewErrBusinessFunctionCodeProperty("TypeSubType", typeSubType,
			fwm.BusinessFunctionCode.BusinessFunctionCode))
	}
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		return err
	}
	return nil
}

// checkProhibitedServiceMessageTags ensures there are no tags present in the message that are incompatible with the BFCServiceMessage code
// Tags NOT permitted:
//   BusinessFunctionCode.TransactionTypeCode, LocalInstrument, PaymentNotification, Charges, InstructedAmount, ExchangeRate,
//   Beneficiary Code = SWIFTBICORBEIANDAccountNumber, Originator Code = SWIFTBICORBEIANDAccountNumber, OriginatorOptionF,
//   any {7xxx} tag, any {8xxx} tag
func (fwm *FEDWireMessage) checkProhibitedServiceMessageTags() error {
	// BusinessFunctionCode.TransactionTypeCode (Element 02) is invalid
	if fwm.BusinessFunctionCode != nil {
		if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
			return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
		}
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.Charges != nil {
		return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
	}
	if fwm.InstructedAmount != nil {
		return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
	}
	if fwm.ExchangeRate != nil {
		return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
	}
	if fwm.Beneficiary != nil && fwm.Beneficiary.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
	}
	if fwm.Originator != nil && fwm.Originator.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}
	return nil
}

// checkSharedProhibitedTags uses case logic for BusinessFunctionCodes that have the same invalid tags.  If this were to change per
// BusinessFunctionCode, create function isInvalidBusinessFunctionCodeTag() with the specific invalid tags for that
// BusinessFunctionCode (e.g. checkProhibitedBankTransferTags)
func (fwm *FEDWireMessage) checkSharedProhibitedTags() error {
	// shared between CheckSameDaySettlement, DepositSendersAccount, FEDFundsReturned, FEDFundsSold, DrawdownResponse, BankDrawDownRequest, and CustomerCorporateDrawdownRequest
	if strings.TrimSpace(fwm.BusinessFunctionCode.TransactionTypeCode) != "" {
		return fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, fwm.BusinessFunctionCode.TransactionTypeCode)
	}
	if fwm.LocalInstrument != nil {
		return fieldError("LocalInstrument", ErrInvalidProperty, fwm.LocalInstrument)
	}
	if fwm.PaymentNotification != nil {
		return fieldError("PaymentNotification", ErrInvalidProperty, fwm.PaymentNotification)
	}
	if fwm.Charges != nil {
		return fieldError("Charges", ErrInvalidProperty, fwm.Charges)
	}
	if fwm.InstructedAmount != nil {
		return fieldError("InstructedAmount", ErrInvalidProperty, fwm.InstructedAmount)
	}
	if fwm.ExchangeRate != nil {
		return fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate)
	}
	if fwm.Beneficiary.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, fwm.Beneficiary.Personal.IdentificationCode)
	}
	if fwm.Originator.Personal.IdentificationCode == SWIFTBICORBEIANDAccountNumber {
		return fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, fwm.Originator.Personal.IdentificationCode)
	}
	if fwm.OriginatorOptionF != nil {
		return fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF)
	}
	if fwm.ServiceMessage != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "ServiceMessage")
	}
	if fwm.UnstructuredAddenda != nil {
		return fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda")
	}
	if err := fwm.invalidCoverPaymentTags(); err != nil {
		return err
	}
	if err := fwm.invalidRemittanceTags(); err != nil {
		return err
	}

	switch fwm.BusinessFunctionCode.BusinessFunctionCode {
	case CheckSameDaySettlement, DepositSendersAccount, FEDFundsReturned, FEDFundsSold:
		// unique exclusions: AccountDebitedDrawdown, AccountCreditedDrawdown, FIDrawdownDebitAccountAdvice
		if fwm.AccountDebitedDrawdown != nil {
			return fieldError("AccountDebitedDrawdown", ErrInvalidProperty, fwm.AccountDebitedDrawdown)
		}
		if fwm.AccountCreditedDrawdown != nil {
			return fieldError("AccountCreditedDrawdown", ErrInvalidProperty, fwm.AccountCreditedDrawdown)
		}
		if fwm.FIDrawdownDebitAccountAdvice != nil {
			return fieldError("FIDrawdownDebitAccountAdvice", ErrInvalidProperty, fwm.FIDrawdownDebitAccountAdvice)
		}
	case DrawdownResponse, BankDrawDownRequest, CustomerCorporateDrawdownRequest:
		// this group has no unique exclusions
	}
	return nil
}

// invalidRemittanceTags returns an error if certain {8xxx} range tags are present.
// The validity of these tags generally depends on the value of the LocalInstrument tag.
func (fwm *FEDWireMessage) invalidRemittanceTags() error {
	if fwm.RelatedRemittance != nil {
		return fieldError("RelatedRemittance", ErrInvalidProperty, "RelatedRemittance")
	}
	if fwm.RemittanceOriginator != nil {
		return fieldError("RemittanceOriginator", ErrInvalidProperty, "RemittanceOriginator")
	}
	if fwm.RemittanceBeneficiary != nil {
		return fieldError("RemittanceBeneficiary", ErrInvalidProperty, "RemittanceBeneficiary")
	}
	if fwm.PrimaryRemittanceDocument != nil {
		return fieldError("PrimaryRemittanceDocument", ErrInvalidProperty, "PrimaryRemittanceDocument")
	}
	if fwm.ActualAmountPaid != nil {
		return fieldError("ActualAmountPaid", ErrInvalidProperty, "ActualAmountPaid")
	}
	if fwm.GrossAmountRemittanceDocument != nil {
		return fieldError("GrossAmountRemittanceDocument", ErrInvalidProperty, "GrossAmountRemittanceDocument")
	}
	if fwm.AmountNegotiatedDiscount != nil {
		return fieldError("AmountNegotiatedDiscount", ErrInvalidProperty, "AmountNegotiatedDiscount")
	}
	if fwm.Adjustment != nil {
		return fieldError("Adjustment", ErrInvalidProperty, "Adjustment")
	}
	if fwm.DateRemittanceDocument != nil {
		return fieldError("DateRemittanceDocument", ErrInvalidProperty, "DateRemittanceDocument")
	}
	if fwm.SecondaryRemittanceDocument != nil {
		return fieldError("SecondaryRemittanceDocument", ErrInvalidProperty, "SecondaryRemittanceDocument")
	}
	if fwm.RemittanceFreeText != nil {
		return fieldError("RemittanceFreeText", ErrInvalidProperty, "RemittanceFreeText")
	}
	return nil
}

// invalidCoverPaymentTags returns an error if certain {7xxx} range tags are present.
// The validity of these tags generally depends on the value of the LocalInstrument tag.
func (fwm *FEDWireMessage) invalidCoverPaymentTags() error {
	if fwm.CurrencyInstructedAmount != nil {
		return fieldError("CurrencyInstructedAmount ", ErrInvalidProperty, fwm.CurrencyInstructedAmount)
	}
	if fwm.OrderingCustomer != nil {
		return fieldError("OrderingCustomer", ErrInvalidProperty, fwm.OrderingCustomer)
	}
	if fwm.OrderingInstitution != nil {
		return fieldError("OrderingInstitution", ErrInvalidProperty, fwm.OrderingInstitution)
	}
	if fwm.IntermediaryInstitution != nil {
		return fieldError("IntermediaryInstitution", ErrInvalidProperty, fwm.IntermediaryInstitution)
	}
	if fwm.InstitutionAccount != nil {
		return fieldError("InstitutionAccount", ErrInvalidProperty, fwm.InstitutionAccount)
	}
	if fwm.BeneficiaryCustomer != nil {
		return fieldError("BeneficiaryCustomer", ErrInvalidProperty, fwm.BeneficiaryCustomer)
	}
	if fwm.Remittance != nil {
		return fieldError("Remittance", ErrInvalidProperty, fwm.Remittance)
	}
	if fwm.SenderToReceiver != nil {
		return fieldError("SenderToReceiver", ErrInvalidProperty, fwm.SenderToReceiver)
	}
	return nil
}

// SetSenderSupplied appends a SenderSupplied to the FEDWireMessage
func (fwm *FEDWireMessage) SetSenderSupplied(ss *SenderSupplied) {
	fwm.SenderSupplied = ss
}

// GetSenderSupplied returns the current SenderSupplied
func (fwm *FEDWireMessage) GetSenderSupplied() *SenderSupplied {
	return fwm.SenderSupplied
}

// SetTypeSubType appends a TypeSubType to the FEDWireMessage
func (fwm *FEDWireMessage) SetTypeSubType(tst *TypeSubType) {
	fwm.TypeSubType = tst
}

// GetTypeSubType returns the current TypeSubType
func (fwm *FEDWireMessage) GetTypeSubType() *TypeSubType {
	return fwm.TypeSubType
}

// SetInputMessageAccountabilityData appends a InputMessageAccountabilityData to the FEDWireMessage
func (fwm *FEDWireMessage) SetInputMessageAccountabilityData(imad *InputMessageAccountabilityData) {
	fwm.InputMessageAccountabilityData = imad
}

// GetInputMessageAccountabilityData returns the current InputMessageAccountabilityData
func (fwm *FEDWireMessage) GetInputMessageAccountabilityData() *InputMessageAccountabilityData {
	return fwm.InputMessageAccountabilityData
}

// SetAmount appends a Amount to the FEDWireMessage
func (fwm *FEDWireMessage) SetAmount(amt *Amount) {
	fwm.Amount = amt
}

// GetAmount returns the current Amount
func (fwm *FEDWireMessage) GetAmount() *Amount {
	return fwm.Amount
}

// SetSenderDepositoryInstitution appends a SenderDepositoryInstitution to the FEDWireMessage
func (fwm *FEDWireMessage) SetSenderDepositoryInstitution(sdi *SenderDepositoryInstitution) {
	fwm.SenderDepositoryInstitution = sdi
}

// GetSenderDepositoryInstitution returns the current SenderDepositoryInstitution
func (fwm *FEDWireMessage) GetSenderDepositoryInstitution() *SenderDepositoryInstitution {
	return fwm.SenderDepositoryInstitution
}

// SetReceiverDepositoryInstitution appends a ReceiverDepositoryInstitution to the FEDWireMessage
func (fwm *FEDWireMessage) SetReceiverDepositoryInstitution(rdi *ReceiverDepositoryInstitution) {
	fwm.ReceiverDepositoryInstitution = rdi
}

// GetReceiverDepositoryInstitution returns the current ReceiverDepositoryInstitution
func (fwm *FEDWireMessage) GetReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	return fwm.ReceiverDepositoryInstitution
}

// SetBusinessFunctionCode appends a BusinessFunctionCode to the FEDWireMessage
func (fwm *FEDWireMessage) SetBusinessFunctionCode(bfc *BusinessFunctionCode) {
	fwm.BusinessFunctionCode = bfc
}

// GetBusinessFunctionCode returns the current BusinessFunctionCode
func (fwm *FEDWireMessage) GetBusinessFunctionCode() *BusinessFunctionCode {
	return fwm.BusinessFunctionCode
}

// SetSenderReference appends a SenderReference to the FEDWireMessage
func (fwm *FEDWireMessage) SetSenderReference(sr *SenderReference) {
	fwm.SenderReference = sr
}

// GetSenderReference returns the current SenderReference
func (fwm *FEDWireMessage) GetSenderReference() *SenderReference {
	return fwm.SenderReference
}

// SetPreviousMessageIdentifier appends a PreviousMessageIdentifier to the FEDWireMessage
func (fwm *FEDWireMessage) SetPreviousMessageIdentifier(pmi *PreviousMessageIdentifier) {
	fwm.PreviousMessageIdentifier = pmi
}

// GetPreviousMessageIdentifier returns the current PreviousMessageIdentifier
func (fwm *FEDWireMessage) GetPreviousMessageIdentifier() *PreviousMessageIdentifier {
	return fwm.PreviousMessageIdentifier
}

// SetLocalInstrument appends a LocalInstrument to the FEDWireMessage
func (fwm *FEDWireMessage) SetLocalInstrument(li *LocalInstrument) {
	fwm.LocalInstrument = li
}

// GetLocalInstrument returns the current LocalInstrument
func (fwm *FEDWireMessage) GetLocalInstrument() *LocalInstrument {
	return fwm.LocalInstrument
}

// SetPaymentNotification appends a PaymentNotification to the FEDWireMessage
func (fwm *FEDWireMessage) SetPaymentNotification(pn *PaymentNotification) {
	fwm.PaymentNotification = pn
}

// GetPaymentNotification returns the current PaymentNotification
func (fwm *FEDWireMessage) GetPaymentNotification() *PaymentNotification {
	return fwm.PaymentNotification
}

// SetCharges appends a Charges to the FEDWireMessage
func (fwm *FEDWireMessage) SetCharges(c *Charges) {
	fwm.Charges = c
}

// GetCharges returns the current Charges
func (fwm *FEDWireMessage) GetCharges() *Charges {
	return fwm.Charges
}

// SetInstructedAmount appends a InstructedAmount to the FEDWireMessage
func (fwm *FEDWireMessage) SetInstructedAmount(ia *InstructedAmount) {
	fwm.InstructedAmount = ia
}

// GetInstructedAmount returns the current Instructed Amount
func (fwm *FEDWireMessage) GetInstructedAmount() *InstructedAmount {
	return fwm.InstructedAmount
}

// SetExchangeRate appends a ExchangeRate to the FEDWireMessage
func (fwm *FEDWireMessage) SetExchangeRate(er *ExchangeRate) {
	fwm.ExchangeRate = er
}

// GetExchangeRate returns the current ExchangeRate
func (fwm *FEDWireMessage) GetExchangeRate() *ExchangeRate {
	return fwm.ExchangeRate
}

// SetBeneficiaryIntermediaryFI appends a BeneficiaryIntermediaryFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetBeneficiaryIntermediaryFI(bifi *BeneficiaryIntermediaryFI) {
	fwm.BeneficiaryIntermediaryFI = bifi
}

// GetBeneficiaryIntermediaryFI returns the current BeneficiaryIntermediaryFI
func (fwm *FEDWireMessage) GetBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	return fwm.BeneficiaryIntermediaryFI
}

// SetBeneficiaryFI appends a BeneficiaryFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetBeneficiaryFI(bfi *BeneficiaryFI) {
	fwm.BeneficiaryFI = bfi
}

// GetBeneficiaryFI returns the current BeneficiaryFI
func (fwm *FEDWireMessage) GetBeneficiaryFI() *BeneficiaryFI {
	return fwm.BeneficiaryFI
}

// SetBeneficiary appends a Beneficiary to the FEDWireMessage
func (fwm *FEDWireMessage) SetBeneficiary(ben *Beneficiary) {
	fwm.Beneficiary = ben
}

// GetBeneficiary returns the current Beneficiary
func (fwm *FEDWireMessage) GetBeneficiary() *Beneficiary {
	return fwm.Beneficiary
}

// SetBeneficiaryReference appends a BeneficiaryReference to the FEDWireMessage
func (fwm *FEDWireMessage) SetBeneficiaryReference(br *BeneficiaryReference) {
	fwm.BeneficiaryReference = br
}

// GetBeneficiaryReference returns the current BeneficiaryReference
func (fwm *FEDWireMessage) GetBeneficiaryReference() *BeneficiaryReference {
	return fwm.BeneficiaryReference
}

// SetAccountDebitedDrawdown appends a AccountDebitedDrawdown to the FEDWireMessage
func (fwm *FEDWireMessage) SetAccountDebitedDrawdown(debitDD *AccountDebitedDrawdown) {
	fwm.AccountDebitedDrawdown = debitDD
}

// GetAccountDebitedDrawdown returns the current AccountDebitedDrawdown
func (fwm *FEDWireMessage) GetAccountDebitedDrawdown() *AccountDebitedDrawdown {
	return fwm.AccountDebitedDrawdown
}

// SetOriginator appends a Originator to the FEDWireMessage
func (fwm *FEDWireMessage) SetOriginator(o *Originator) {
	fwm.Originator = o
}

// GetOriginator returns the current Originator
func (fwm *FEDWireMessage) GetOriginator() *Originator {
	return fwm.Originator
}

// SetOriginatorOptionF appends a OriginatorOptionF to the FEDWireMessage
func (fwm *FEDWireMessage) SetOriginatorOptionF(oof *OriginatorOptionF) {
	fwm.OriginatorOptionF = oof
}

// GetOriginatorOptionF returns the current OriginatorOptionF
func (fwm *FEDWireMessage) GetOriginatorOptionF() *OriginatorOptionF {
	return fwm.OriginatorOptionF
}

// SetOriginatorFI appends a OriginatorFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetOriginatorFI(ofi *OriginatorFI) {
	fwm.OriginatorFI = ofi
}

// GetOriginatorFI returns the current OriginatorFI
func (fwm *FEDWireMessage) GetOriginatorFI() *OriginatorFI {
	return fwm.OriginatorFI
}

// SetInstructingFI appends a InstructingFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetInstructingFI(ifi *InstructingFI) {
	fwm.InstructingFI = ifi
}

// GetInstructingFI returns the current InstructingFI
func (fwm *FEDWireMessage) GetInstructingFI() *InstructingFI {
	return fwm.InstructingFI
}

// SetAccountCreditedDrawdown appends a AccountCreditedDrawdown to the FEDWireMessage
func (fwm *FEDWireMessage) SetAccountCreditedDrawdown(creditDD *AccountCreditedDrawdown) {
	fwm.AccountCreditedDrawdown = creditDD
}

// GetAccountCreditedDrawdown returns the current AccountCreditedDrawdown
func (fwm *FEDWireMessage) GetAccountCreditedDrawdown() *AccountCreditedDrawdown {
	return fwm.AccountCreditedDrawdown
}

// SetOriginatorToBeneficiary appends a OriginatorToBeneficiary to the FEDWireMessage
func (fwm *FEDWireMessage) SetOriginatorToBeneficiary(ob *OriginatorToBeneficiary) {
	fwm.OriginatorToBeneficiary = ob
}

// GetOriginatorToBeneficiary returns the current OriginatorToBeneficiary
func (fwm *FEDWireMessage) GetOriginatorToBeneficiary() *OriginatorToBeneficiary {
	return fwm.OriginatorToBeneficiary
}

// SetFIReceiverFI appends a FIReceiverFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIReceiverFI(firfi *FIReceiverFI) {
	fwm.FIReceiverFI = firfi
}

// GetFIReceiverFI returns the current FIReceiverFI
func (fwm *FEDWireMessage) GetFIReceiverFI() *FIReceiverFI {
	return fwm.FIReceiverFI
}

// SetFIDrawdownDebitAccountAdvice appends a FIDrawdownDebitAccountAdvice to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIDrawdownDebitAccountAdvice(debitDDAdvice *FIDrawdownDebitAccountAdvice) {
	fwm.FIDrawdownDebitAccountAdvice = debitDDAdvice
}

// GetFIDrawdownDebitAccountAdvice returns the current FIDrawdownDebitAccountAdvice
func (fwm *FEDWireMessage) GetFIDrawdownDebitAccountAdvice() *FIDrawdownDebitAccountAdvice {
	return fwm.FIDrawdownDebitAccountAdvice
}

// SetFIIntermediaryFI appends a FIIntermediaryFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIIntermediaryFI(fiifi *FIIntermediaryFI) {
	fwm.FIIntermediaryFI = fiifi
}

// GetFIIntermediaryFI returns the current FIIntermediaryFI
func (fwm *FEDWireMessage) GetFIIntermediaryFI() *FIIntermediaryFI {
	return fwm.FIIntermediaryFI
}

// SetFIIntermediaryFIAdvice appends a FIIntermediaryFIAdvice to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIIntermediaryFIAdvice(fiifia *FIIntermediaryFIAdvice) {
	fwm.FIIntermediaryFIAdvice = fiifia
}

// GetFIIntermediaryFIAdvice returns the current FIIntermediaryFIAdvice
func (fwm *FEDWireMessage) GetFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	return fwm.FIIntermediaryFIAdvice
}

// SetFIBeneficiaryFI appends a FIBeneficiaryFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIBeneficiaryFI(fibfi *FIBeneficiaryFI) {
	fwm.FIBeneficiaryFI = fibfi
}

// GetFIBeneficiaryFI returns the current FIBeneficiaryFI
func (fwm *FEDWireMessage) GetFIBeneficiaryFI() *FIBeneficiaryFI {
	return fwm.FIBeneficiaryFI
}

// SetFIBeneficiaryFIAdvice appends a FIBeneficiaryFIAdvice to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIBeneficiaryFIAdvice(fibfia *FIBeneficiaryFIAdvice) {
	fwm.FIBeneficiaryFIAdvice = fibfia
}

// GetFIBeneficiaryFIAdvice returns the current FIBeneficiaryFIAdvice
func (fwm *FEDWireMessage) GetFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	return fwm.FIBeneficiaryFIAdvice
}

// SetFIBeneficiary appends a FIBeneficiary to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIBeneficiary(fib *FIBeneficiary) {
	fwm.FIBeneficiary = fib
}

// GetFIBeneficiary returns the current FIBeneficiary
func (fwm *FEDWireMessage) GetFIBeneficiary() *FIBeneficiary {
	return fwm.FIBeneficiary
}

// SetFIBeneficiaryAdvice appends a FIBeneficiaryAdviceto the FEDWireMessage
func (fwm *FEDWireMessage) SetFIBeneficiaryAdvice(fiba *FIBeneficiaryAdvice) {
	fwm.FIBeneficiaryAdvice = fiba
}

// GetFIBeneficiaryAdvice returns the current FIBeneficiaryAdvice
func (fwm *FEDWireMessage) GetFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	return fwm.FIBeneficiaryAdvice
}

// SetFIPaymentMethodToBeneficiary appends a FIPaymentMethodToBeneficiary to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIPaymentMethodToBeneficiary(pm *FIPaymentMethodToBeneficiary) {
	fwm.FIPaymentMethodToBeneficiary = pm
}

// GetFIPaymentMethodToBeneficiary returns the current FIPaymentMethodToBeneficiary
func (fwm *FEDWireMessage) GetFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	return fwm.FIPaymentMethodToBeneficiary
}

// SetFIAdditionalFIToFI appends a FIAdditionalFIToFI to the FEDWireMessage
func (fwm *FEDWireMessage) SetFIAdditionalFIToFI(fifi *FIAdditionalFIToFI) {
	fwm.FIAdditionalFIToFI = fifi
}

// GetFIAdditionalFIToFI returns the current FIAdditionalFIToFI
func (fwm *FEDWireMessage) GetFIAdditionalFIToFI() *FIAdditionalFIToFI {
	return fwm.FIAdditionalFIToFI
}

// SetCurrencyInstructedAmount appends a CurrencyInstructedAmount to the FEDWireMessage
func (fwm *FEDWireMessage) SetCurrencyInstructedAmount(cia *CurrencyInstructedAmount) {
	fwm.CurrencyInstructedAmount = cia
}

// GetCurrencyInstructedAmount returns the current CurrencyInstructedAmount
func (fwm *FEDWireMessage) GetCurrencyInstructedAmount() *CurrencyInstructedAmount {
	return fwm.CurrencyInstructedAmount
}

// SetOrderingCustomer appends a OrderingCustomer to the FEDWireMessage
func (fwm *FEDWireMessage) SetOrderingCustomer(oc *OrderingCustomer) {
	fwm.OrderingCustomer = oc
}

// GetOrderingCustomer returns the current OrderingCustomer
func (fwm *FEDWireMessage) GetOrderingCustomer() *OrderingCustomer {
	return fwm.OrderingCustomer
}

// SetOrderingInstitution appends a OrderingInstitution to the FEDWireMessage
func (fwm *FEDWireMessage) SetOrderingInstitution(oi *OrderingInstitution) {
	fwm.OrderingInstitution = oi
}

// GetOrderingInstitution returns the current OrderingInstitution
func (fwm *FEDWireMessage) GetOrderingInstitution() *OrderingInstitution {
	return fwm.OrderingInstitution
}

// SetIntermediaryInstitution appends a IntermediaryInstitution to the FEDWireMessage
func (fwm *FEDWireMessage) SetIntermediaryInstitution(ii *IntermediaryInstitution) {
	fwm.IntermediaryInstitution = ii
}

// GetIntermediaryInstitution returns the current IntermediaryInstitution
func (fwm *FEDWireMessage) GetIntermediaryInstitution() *IntermediaryInstitution {
	return fwm.IntermediaryInstitution
}

// SetInstitutionAccount appends a InstitutionAccount to the FEDWireMessage
func (fwm *FEDWireMessage) SetInstitutionAccount(iAccount *InstitutionAccount) {
	fwm.InstitutionAccount = iAccount
}

// GetInstitutionAccount returns the current InstitutionAccount
func (fwm *FEDWireMessage) GetInstitutionAccount() *InstitutionAccount {
	return fwm.InstitutionAccount
}

// SetBeneficiaryCustomer appends a BeneficiaryCustomer to the FEDWireMessage
func (fwm *FEDWireMessage) SetBeneficiaryCustomer(bc *BeneficiaryCustomer) {
	fwm.BeneficiaryCustomer = bc
}

// GetBeneficiaryCustomer returns the current BeneficiaryCustomer
func (fwm *FEDWireMessage) GetBeneficiaryCustomer() *BeneficiaryCustomer {
	return fwm.BeneficiaryCustomer
}

// SetRemittance appends a Remittance to the FEDWireMessage
func (fwm *FEDWireMessage) SetRemittance(ri *Remittance) {
	fwm.Remittance = ri
}

// GetRemittance returns the current Remittance
func (fwm *FEDWireMessage) GetRemittance() *Remittance {
	return fwm.Remittance
}

// SetSenderToReceiver appends a SenderToReceiver to the FEDWireMessage
func (fwm *FEDWireMessage) SetSenderToReceiver(str *SenderToReceiver) {
	fwm.SenderToReceiver = str
}

// GetSenderToReceiver  returns the current SenderToReceiver
func (fwm *FEDWireMessage) GetSenderToReceiver() *SenderToReceiver {
	return fwm.SenderToReceiver
}

// SetUnstructuredAddenda appends a UnstructuredAddenda to the FEDWireMessage
func (fwm *FEDWireMessage) SetUnstructuredAddenda(ua *UnstructuredAddenda) {
	fwm.UnstructuredAddenda = ua
}

// GetUnstructuredAddenda returns the current UnstructuredAddenda
func (fwm *FEDWireMessage) GetUnstructuredAddenda() *UnstructuredAddenda {
	return fwm.UnstructuredAddenda
}

// SetRelatedRemittance appends a RelatedRemittance to the FEDWireMessage
func (fwm *FEDWireMessage) SetRelatedRemittance(rr *RelatedRemittance) {
	fwm.RelatedRemittance = rr
}

// GetRelatedRemittance returns the current RelatedRemittance
func (fwm *FEDWireMessage) GetRelatedRemittance() *RelatedRemittance {
	return fwm.RelatedRemittance
}

// SetRemittanceOriginator appends a RemittanceOriginator to the FEDWireMessage
func (fwm *FEDWireMessage) SetRemittanceOriginator(ro *RemittanceOriginator) {
	fwm.RemittanceOriginator = ro
}

// GetRemittanceOriginator returns the current RemittanceOriginator
func (fwm *FEDWireMessage) GetRemittanceOriginator() *RemittanceOriginator {
	return fwm.RemittanceOriginator
}

// SetRemittanceBeneficiary appends a RemittanceBeneficiary to the FEDWireMessage
func (fwm *FEDWireMessage) SetRemittanceBeneficiary(rb *RemittanceBeneficiary) {
	fwm.RemittanceBeneficiary = rb
}

// GetRemittanceBeneficiary returns the current RemittanceBeneficiary
func (fwm *FEDWireMessage) GetRemittanceBeneficiary() *RemittanceBeneficiary {
	return fwm.RemittanceBeneficiary
}

// SetPrimaryRemittanceDocument appends a PrimaryRemittanceDocument to the FEDWireMessage
func (fwm *FEDWireMessage) SetPrimaryRemittanceDocument(prd *PrimaryRemittanceDocument) {
	fwm.PrimaryRemittanceDocument = prd
}

// GetPrimaryRemittanceDocument returns the current PrimaryRemittanceDocument
func (fwm *FEDWireMessage) GetPrimaryRemittanceDocument() *PrimaryRemittanceDocument {
	return fwm.PrimaryRemittanceDocument
}

// SetActualAmountPaid appends a ActualAmountPaid to the FEDWireMessage
func (fwm *FEDWireMessage) SetActualAmountPaid(aap *ActualAmountPaid) {
	fwm.ActualAmountPaid = aap
}

// GetActualAmountPaid returns the current ActualAmountPaid
func (fwm *FEDWireMessage) GetActualAmountPaid() *ActualAmountPaid {
	return fwm.ActualAmountPaid
}

// SetGrossAmountRemittanceDocument appends a GrossAmountRemittanceDocument to the FEDWireMessage
func (fwm *FEDWireMessage) SetGrossAmountRemittanceDocument(gard *GrossAmountRemittanceDocument) {
	fwm.GrossAmountRemittanceDocument = gard
}

// GetGrossAmountRemittanceDocument returns the current GrossAmountRemittanceDocument
func (fwm *FEDWireMessage) GetGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	return fwm.GrossAmountRemittanceDocument
}

// SetAmountNegotiatedDiscount appends a AmountNegotiatedDiscount to the FEDWireMessage
func (fwm *FEDWireMessage) SetAmountNegotiatedDiscount(nd *AmountNegotiatedDiscount) {
	fwm.AmountNegotiatedDiscount = nd
}

// GetAmountNegotiatedDiscount returns the current AmountNegotiatedDiscount
func (fwm *FEDWireMessage) GetAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	return fwm.AmountNegotiatedDiscount
}

// SetAdjustment appends a Adjustment to the FEDWireMessage
func (fwm *FEDWireMessage) SetAdjustment(adj *Adjustment) {
	fwm.Adjustment = adj
}

// GetAdjustment returns the current Adjustment
func (fwm *FEDWireMessage) GetAdjustment() *Adjustment {
	return fwm.Adjustment
}

// SetDateRemittanceDocument appends a DateRemittanceDocument to the FEDWireMessage
func (fwm *FEDWireMessage) SetDateRemittanceDocument(drd *DateRemittanceDocument) {
	fwm.DateRemittanceDocument = drd
}

// GetDateRemittanceDocument returns the current DateRemittanceDocument
func (fwm *FEDWireMessage) GetDateRemittanceDocument() *DateRemittanceDocument {
	return fwm.DateRemittanceDocument
}

// SetSecondaryRemittanceDocument appends a SecondaryRemittanceDocument to the FEDWireMessage
func (fwm *FEDWireMessage) SetSecondaryRemittanceDocument(srd *SecondaryRemittanceDocument) {
	fwm.SecondaryRemittanceDocument = srd
}

// GetSecondaryRemittanceDocument returns the current SecondaryRemittanceDocument
func (fwm *FEDWireMessage) GetSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	return fwm.SecondaryRemittanceDocument
}

// SetRemittanceFreeText appends a RemittanceFreeText to the FEDWireMessage
func (fwm *FEDWireMessage) SetRemittanceFreeText(rft *RemittanceFreeText) {
	fwm.RemittanceFreeText = rft
}

// GetRemittanceFreeText returns the current RemittanceFreeText
func (fwm *FEDWireMessage) GetRemittanceFreeText() *RemittanceFreeText {
	return fwm.RemittanceFreeText
}

// SetServiceMessage appends a ServiceMessage to the FEDWireMessage
func (fwm *FEDWireMessage) SetServiceMessage(sm *ServiceMessage) {
	fwm.ServiceMessage = sm
}

// GetServiceMessage returns the current ServiceMessage
func (fwm *FEDWireMessage) GetServiceMessage() *ServiceMessage {
	return fwm.ServiceMessage
}

// SetMessageDisposition appends a MessageDisposition to the FEDWireMessage
func (fwm *FEDWireMessage) SetMessageDisposition(md *MessageDisposition) {
	fwm.MessageDisposition = md
}

// GetMessageDisposition returns the current MessageDisposition
func (fwm *FEDWireMessage) GetMessageDisposition() *MessageDisposition {
	return fwm.MessageDisposition
}

// SetReceiptTimeStamp appends a ReceiptTimeStamp to the FEDWireMessage
func (fwm *FEDWireMessage) SetReceiptTimeStamp(rts *ReceiptTimeStamp) {
	fwm.ReceiptTimeStamp = rts
}

// GetReceiptTimeStamp returns the current ReceiptTimeStamp
func (fwm *FEDWireMessage) GetReceiptTimeStamp() *ReceiptTimeStamp {
	return fwm.ReceiptTimeStamp
}

// SetOutputMessageAccountabilityData appends a ErrorWire to the FEDWireMessage
func (fwm *FEDWireMessage) SetOutputMessageAccountabilityData(omad *OutputMessageAccountabilityData) {
	fwm.OutputMessageAccountabilityData = omad
}

// GetOutputMessageAccountabilityData returns the current OutputMessageAccountabilityData
func (fwm *FEDWireMessage) GetOutputMessageAccountabilityData() *OutputMessageAccountabilityData {
	return fwm.OutputMessageAccountabilityData
}

// SetErrorWire appends a ErrorWire to the FEDWireMessage
func (fwm *FEDWireMessage) SetErrorWire(ew *ErrorWire) {
	fwm.ErrorWire = ew
}

// GetErrorWire returns the current ErrorWire
func (fwm *FEDWireMessage) GetErrorWire() *ErrorWire {
	return fwm.ErrorWire
}

// Only allowed if BusinessFunctionCode is CustomerTransferPlus.
func (fwm *FEDWireMessage) validateLocalInstrumentCode() error {
	if fwm.LocalInstrument != nil {
		if fwm.BusinessFunctionCode.BusinessFunctionCode != CustomerTransferPlus {
			return fieldError("LocalInstrument", ErrLocalInstrumentNotPermitted)
		}
		return nil
	}
	return nil

}

// BusinessFunctionCode must be CustomerTransfer or CustomerTransferPlus. Not permitted if LocalInstrument Code is SequenceBCoverPaymentStructured.
func (fwm *FEDWireMessage) validateCharges() error {
	if fwm.Charges != nil {
		bfc := fwm.BusinessFunctionCode.BusinessFunctionCode
		if !(bfc == CustomerTransfer || bfc == CustomerTransferPlus) {
			return NewErrInvalidPropertyForProperty("BusinessFunctionCode", bfc, "Charges", fwm.Charges.String())
		}
		if fwm.LocalInstrument != nil && fwm.LocalInstrument.LocalInstrumentCode == SequenceBCoverPaymentStructured {
			return NewErrInvalidPropertyForProperty("LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode,
				"Charges", fwm.Charges.String())
		}
		return nil
	}
	return nil
}

// Mandatory if ExchangeRate is present.
// BusinessFunctionCode must be CustomerTransfer or CustomerTransferPlus.
// Not permitted if LocalInstrument Code is SequenceBCoverPaymentStructured.
func (fwm *FEDWireMessage) validateInstructedAmount() error {
	if fwm.ExchangeRate != nil && fwm.InstructedAmount == nil {
		return fieldError("InstructedAmount", ErrFieldRequired)

	}
	if fwm.InstructedAmount != nil {
		bfc := fwm.BusinessFunctionCode.BusinessFunctionCode
		if !(bfc == CustomerTransfer || bfc == CustomerTransferPlus) {
			return NewErrInvalidPropertyForProperty("BusinessFunctionCode", bfc, "InstructedAmount", fwm.InstructedAmount.String())
		}
		if fwm.LocalInstrument != nil && fwm.LocalInstrument.LocalInstrumentCode == SequenceBCoverPaymentStructured {
			return NewErrInvalidPropertyForProperty("LocalInstrumentCode",
				fwm.LocalInstrument.LocalInstrumentCode, "Instructed Amount", fwm.InstructedAmount.String())
		}
		return nil
	}
	return nil
}

// validateExchangeRate validates TagExchangeRate within a FEDWireMessage
// * If present, InstructedAmount is mandatory.
// * BusinessFunctionCode must be CustomerTransfer or CustomerTransferPlus.
// * Not permitted if LocalInstrument Code is SequenceBCoverPaymentStructured.
func (fwm *FEDWireMessage) validateExchangeRate() error {
	if fwm.ExchangeRate != nil {
		if fwm.InstructedAmount == nil {
			return fieldError("InstructedAmount", ErrFieldRequired)
		}
		bfc := fwm.BusinessFunctionCode.BusinessFunctionCode
		if !(bfc == CustomerTransfer || bfc == CustomerTransferPlus) {
			return NewErrInvalidPropertyForProperty("BusinessFunctionCode", bfc, "InstructedAmount", fwm.InstructedAmount.String())
		}
		if fwm.LocalInstrument != nil && fwm.LocalInstrument.LocalInstrumentCode == SequenceBCoverPaymentStructured {
			return NewErrInvalidPropertyForProperty("LocalInstrumentCode",
				fwm.LocalInstrument.LocalInstrumentCode, "ExchangeRate", fwm.ExchangeRate.ExchangeRate)
		}
		return nil
	}
	return nil
}

// If present, tags BeneficiaryFI and Beneficiary are mandatory.
func (fwm *FEDWireMessage) validateBeneficiaryIntermediaryFI() error {
	if fwm.BeneficiaryIntermediaryFI != nil {
		if fwm.BeneficiaryFI == nil {
			return fieldError("BeneficiaryFI", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// If present, the Beneficiary tag is mandatory.
func (fwm *FEDWireMessage) validateBeneficiaryFI() error {
	if fwm.BeneficiaryFI != nil {
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// If present, Originator (or OriginatorOptionF if BusinessFunctionCode is CustomerTransferPlus) is mandatory.
func (fwm *FEDWireMessage) validateOriginatorFI() error {
	if fwm.OriginatorFI != nil {
		switch fwm.BusinessFunctionCode.BusinessFunctionCode {
		case CustomerTransferPlus:
			if fwm.OriginatorOptionF == nil {
				return fieldError("OriginatorOptionF", ErrFieldRequired)
			}
		default:
			if fwm.Originator == nil {
				return fieldError("Originator", ErrFieldRequired)
			}
		}
		return nil
	}
	return nil
}

// If present, Originator (or OriginatorOptionF if BusinessFunctionCode is CustomerTransferPlus) and OriginatorFI are mandatory.
func (fwm *FEDWireMessage) validateInstructingFI() error {
	if fwm.InstructingFI != nil {
		switch fwm.BusinessFunctionCode.BusinessFunctionCode {
		case CustomerTransferPlus:
			if fwm.OriginatorOptionF == nil {
				return fieldError("OriginatorOptionF", ErrFieldRequired)
			}
		default:
			if fwm.Originator == nil {
				return fieldError("Originator", ErrFieldRequired)
			}
		}
		if fwm.OriginatorFI == nil {
			return fieldError("OriginatorFI", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// If present, Beneficiary and Originator (or OriginatorOptionF if BusinessFunctionCode is CustomerTransferPlus) are mandatory.
func (fwm *FEDWireMessage) validateOriginatorToBeneficiary() error {
	if fwm.OriginatorToBeneficiary != nil {
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		switch fwm.BusinessFunctionCode.BusinessFunctionCode {
		case CustomerTransferPlus:
			if fwm.OriginatorOptionF == nil {
				return fieldError("OriginatorOptionF", ErrFieldRequired)
			}
		default:
			if fwm.Originator == nil {
				return fieldError("Originator", ErrFieldRequired)
			}
		}
		return nil
	}
	return nil
}

// validateFIIntermediaryFI validates TagFIIntermediaryFI within a FEDWireMessage
// If present, BeneficiaryIntermediaryFI, BeneficiaryFI and Beneficiary are required.
func (fwm *FEDWireMessage) validateFIIntermediaryFI() error {
	if fwm.FIIntermediaryFI != nil {
		if fwm.BeneficiaryIntermediaryFI == nil {
			return fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired)
		}
		if fwm.BeneficiaryFI == nil {
			return fieldError("BeneficiaryFI", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIIntermediaryFIAdvice validates TagFIIntermediaryFIAdvice within a FEDWireMessage
// If present, BeneficiaryIntermediaryFI, BeneficiaryFI and Beneficiary are required.
func (fwm *FEDWireMessage) validateFIIntermediaryFIAdvice() error {
	if fwm.FIIntermediaryFIAdvice != nil {
		if fwm.BeneficiaryIntermediaryFI == nil {
			return fieldError("BeneficiaryIntermediaryFI", ErrFieldRequired)
		}
		if fwm.BeneficiaryFI == nil {
			return fieldError("BeneficiaryFI", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIBeneficiaryFI validates TagFIBeneficiaryFI within a FEDWireMessage
// If present, BeneficiaryFI and Beneficiary are required.
func (fwm *FEDWireMessage) validateFIBeneficiaryFI() error {
	if fwm.FIBeneficiaryFI != nil {
		if fwm.BeneficiaryFI == nil {
			return fieldError("BeneficiaryFI", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIBeneficiaryFIAdvice validates TagFIBeneficiaryFIAdvice within a FEDWireMessage
// If present, BeneficiaryFI and Beneficiary are required.
func (fwm *FEDWireMessage) validateFIBeneficiaryFIAdvice() error {
	if fwm.FIBeneficiaryFIAdvice != nil {
		if fwm.BeneficiaryFI == nil {
			return fieldError("BeneficiaryFI", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIBeneficiary validates TagFIBeneficiary within a FEDWireMessage
// If present, Beneficiary is required.
func (fwm *FEDWireMessage) validateFIBeneficiary() error {
	if fwm.FIBeneficiary != nil {
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIBeneficiaryAdvice validates TagFIBeneficiaryAdvice within a FEDWireMessage
// If present, Beneficiary is required.
func (fwm *FEDWireMessage) validateFIBeneficiaryAdvice() error {
	if fwm.FIBeneficiaryAdvice != nil {
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateFIPaymentMethodToBeneficiary validates TagFIPaymentMethodToBeneficiary within a FEDWireMessage
// If present, FIBeneficiaryAdvice and Beneficiary are required.
func (fwm *FEDWireMessage) validateFIPaymentMethodToBeneficiary() error {
	if fwm.FIPaymentMethodToBeneficiary != nil {
		if fwm.FIBeneficiaryAdvice == nil {
			return fieldError("FIBeneficiaryAdvice", ErrFieldRequired)
		}
		if fwm.Beneficiary == nil {
			return fieldError("Beneficiary", ErrFieldRequired)
		}
		return nil
	}
	return nil
}

// validateUnstructuredAddenda validates TagUnstructuredAddenda within a FEDWireMessage
// * Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument is ANSIX12format,
//    GeneralXMLformat, ISO20022XMLformat, NarrativeText, STP820format, SWIFTfield70 or UNEDIFACTformat;
//    otherwise not permitted.
// * If LocalInstrument is ANSIX12format or STP820format, only the X12 Character Set* is permitted in
//    Addenda Information element.
// * If LocalInstrument is GeneralXMLformat, ISO20022XMLformat, NarrativeText, SWIFTfield70 or
//    UNEDIFACTformat, only the SWIFT MX ISO 20022 Character Set* is permitted in Addenda Information
//    element.
func (fwm *FEDWireMessage) validateUnstructuredAddenda() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil {
		switch fwm.LocalInstrument.LocalInstrumentCode {
		case ANSIX12format, GeneralXMLformat, ISO20022XMLformat, NarrativeText, STP820format, SWIFTfield70, UNEDIFACTformat:
			if fwm.UnstructuredAddenda == nil {
				return fieldError("UnstructuredAddenda", ErrFieldRequired)
			}
		default:
			if fwm.UnstructuredAddenda != nil {
				return NewErrInvalidPropertyForProperty("UnstructuredAddenda", fwm.UnstructuredAddenda.String(),
					"LocalInstrumentCode", fwm.LocalInstrument.LocalInstrumentCode)
			}
		}
	} else {
		if fwm.UnstructuredAddenda != nil {
			return fieldError("UnstructuredAddenda", ErrNotPermitted)
		}
	}

	// TODO: if LocalInstrument is ANSIX12format or STP820format, make sure Addenda Information only contains charaters within the X12 character set
	// TODO: if LocalInstrument is any of the other permitted formats, make sure Addenda Information only contains charaters within the SWIFT MX ISO 20022 character set

	return nil
}

// validateRelatedRemittance validates TagRelatedRemittance within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument is
//  RelatedRemittanceInformation; otherwise not permitted.
func (fwm *FEDWireMessage) validateRelatedRemittance() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RelatedRemittanceInformation {
		if fwm.RelatedRemittance == nil {
			return fieldError("RelatedRemittance", ErrFieldRequired)
		}
	} else {
		if fwm.RelatedRemittance != nil {
			return fieldError("RelatedRemittance", ErrNotPermitted)
		}
	}

	return nil
}

// validateRemittanceOriginator validates TagRemittanceOriginator within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateRemittanceOriginator() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.RemittanceOriginator == nil {
			return fieldError("RemittanceOriginator", ErrFieldRequired)
		}
	} else {
		if fwm.RemittanceOriginator != nil {
			return fieldError("RemittanceOriginator", ErrNotPermitted)
		}
	}

	return nil
}

// validateRemittanceBeneficiary validates TagRemittanceBeneficiary within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateRemittanceBeneficiary() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.RemittanceBeneficiary == nil {
			return fieldError("RemittanceBeneficiary", ErrFieldRequired)
		}
	} else {
		if fwm.RemittanceBeneficiary != nil {
			return fieldError("RemittanceBeneficiary", ErrNotPermitted)
		}
	}

	return nil
}

// PrimaryRemittanceDocument validates TagPrimaryRemittanceDocument within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validatePrimaryRemittanceDocument() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.PrimaryRemittanceDocument == nil {
			return fieldError("PrimaryRemittanceDocument", ErrFieldRequired)
		}
	} else {
		if fwm.PrimaryRemittanceDocument != nil {
			return fieldError("PrimaryRemittanceDocument", ErrNotPermitted)
		}
	}

	return nil
}

// validateActualAmountPaid validates TagActualAmountPaid within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateActualAmountPaid() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.ActualAmountPaid == nil {
			return fieldError("ActualAmountPaid", ErrFieldRequired)
		}
	} else {
		if fwm.ActualAmountPaid != nil {
			return fieldError("ActualAmountPaid", ErrNotPermitted)
		}
	}

	return nil
}

// validateGrossAmountRemittanceDocument validates TagGrossAmountRemittanceDocument within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateGrossAmountRemittanceDocument() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.GrossAmountRemittanceDocument == nil {
			return fieldError("GrossAmountRemittanceDocument", ErrFieldRequired)
		}
	} else {
		if fwm.GrossAmountRemittanceDocument != nil {
			return fieldError("GrossAmountRemittanceDocument", ErrNotPermitted)
		}
	}

	return nil
}

// validateAdjustment validates TagAdjustment within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateAdjustment() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.Adjustment == nil {
			return fieldError("Adjustment", ErrFieldRequired)
		}
	} else {
		if fwm.Adjustment != nil {
			return fieldError("Adjustment", ErrNotPermitted)
		}
	}

	return nil
}

// validateDateRemittanceDocument validates TagDateRemittanceDocument within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateDateRemittanceDocument() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.DateRemittanceDocument == nil {
			return fieldError("DateRemittanceDocument", ErrFieldRequired)
		}
	} else {
		if fwm.DateRemittanceDocument != nil {
			return fieldError("DateRemittanceDocument", ErrNotPermitted)
		}
	}

	return nil
}

// validateSecondaryRemittanceDocument validates a TagSecondaryRemittanceDocument within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateSecondaryRemittanceDocument() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.SecondaryRemittanceDocument == nil {
			return fieldError("SecondaryRemittanceDocument", ErrFieldRequired)
		}
	} else {
		if fwm.SecondaryRemittanceDocument != nil {
			return fieldError("SecondaryRemittanceDocument", ErrNotPermitted)
		}
	}

	return nil
}

// validateRemittanceFreeText validates a TagRemittanceFreeText within a FEDWireMessage
// Must be present if BusinessFunctionCode is CustomerTransferPlus and LocalInstrument code
//  is RemittanceInformationStructured; otherwise not permitted.
func (fwm *FEDWireMessage) validateRemittanceFreeText() error {
	if fwm.BusinessFunctionCode.BusinessFunctionCode == CustomerTransferPlus && fwm.LocalInstrument != nil &&
		fwm.LocalInstrument.LocalInstrumentCode == RemittanceInformationStructured {
		if fwm.RemittanceFreeText == nil {
			return fieldError("RemittanceFreeText", ErrFieldRequired)
		}
	} else {
		if fwm.RemittanceFreeText != nil {
			return fieldError("RemittanceFreeText", ErrNotPermitted)
		}
	}

	return nil
}

func (fwm *FEDWireMessage) otherTransferInformation() error {
	if err := fwm.validateLocalInstrumentCode(); err != nil {
		return err
	}
	if err := fwm.validateCharges(); err != nil {
		return err
	}
	if err := fwm.validateInstructedAmount(); err != nil {
		return err
	}
	if err := fwm.validateExchangeRate(); err != nil {
		return err
	}
	return nil
}

func (fwm *FEDWireMessage) isRemittanceValid() error {
	if err := fwm.validateRemittanceOriginator(); err != nil {
		return err
	}
	if err := fwm.validateRemittanceBeneficiary(); err != nil {
		return err
	}
	if err := fwm.validatePrimaryRemittanceDocument(); err != nil {
		return err
	}
	if err := fwm.validateActualAmountPaid(); err != nil {
		return err
	}
	if err := fwm.validateGrossAmountRemittanceDocument(); err != nil {
		return err
	}
	if err := fwm.validateAdjustment(); err != nil {
		return err
	}
	if err := fwm.validateDateRemittanceDocument(); err != nil {
		return err
	}
	if err := fwm.validateRemittanceFreeText(); err != nil {
		return err
	}
	return nil
}
