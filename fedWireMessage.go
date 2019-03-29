// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

//ToDo: omitEmpty

// FedWireMessage is a FedWire Message
type FedWireMessage struct {
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
	// FiReceiverFI
	FIReceiverFI *FIReceiverFI `json:"fiReceiverFI,omitempty"`
	// FiDrawdownDebitAccountAdvice
	FIDrawdownDebitAccountAdvice *FIDrawdownDebitAccountAdvice `json:"fiDrawdownDebitAccountAdvice,omitempty"`
	// FiIntermediaryFI
	FIIntermediaryFI *FIIntermediaryFI `json:"fiIntermediaryFI,omitempty"`
	// FiIntermediaryFIAdvice
	FIIntermediaryFIAdvice *FIIntermediaryFIAdvice `json:"fiIntermediaryFIAdvice,omitempty"`
	// FiBeneficiaryFI
	FIBeneficiaryFI *FIBeneficiaryFI `json:"fiBeneficiaryFI,omitempty"`
	// FiBeneficiaryFIAdvice
	FIBeneficiaryFIAdvice *FIBeneficiaryFIAdvice `json:"fiBeneficiaryFIAdvice,omitempty"`
	// FiBeneficiary
	FIBeneficiary *FIBeneficiary `json:"fiBeneficiary,omitempty"`
	// FiBeneficiaryAdvice
	FIBeneficiaryAdvice *FIBeneficiaryAdvice `json:"fiBeneficiaryAdvice,omitempty"`
	// FiPaymentMethodToBeneficiary
	FIPaymentMethodToBeneficiary *FIPaymentMethodToBeneficiary `json:"fiPaymentMethodToBeneficiary,omitempty"`
	// FiAdditionalFiToFi
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
	// ToDo: Remove Information
	// Remittance
	Remittance *RemittanceInformation `json:"remittance,omitempty"`
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

// NewFedWireMessage returns a new FedWireMessage
func NewFedWireMessage() FedWireMessage {
	fwm := FedWireMessage{}
	return fwm
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (fwm *FedWireMessage) verify() error {
	// ToDo: Add errors

	if fwm.SenderSupplied == nil {
		return nil
	}

	if fwm.TypeSubType == nil {
		return nil
	}

	if fwm.InputMessageAccountabilityData == nil {
		return nil
	}

	if fwm.Amount == nil {
		return nil
	}

	if fwm.SenderDepositoryInstitution == nil {
		return nil
	}

	if fwm.ReceiverDepositoryInstitution == nil {
		return nil
	}

	if fwm.BusinessFunctionCode == nil {
		return nil
	}

	return nil
}
