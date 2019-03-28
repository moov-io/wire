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
	BeneficiaryIntermediaryFI *FinancialInstitution `json:"beneficiaryIntermediaryFI,omitempty"`
	// BeneficiaryFI
	BeneficiaryFI *FinancialInstitution `json:"beneficiaryFI,omitempty"`
	// Beneficiary
	Beneficiary *Personal `json:"beneficiary,omitempty"`
	// BeneficiaryReference
	BeneficiaryReference *BeneficiaryReference `json:"beneficiaryReference,omitempty"`
	// AccountDebitedDrawdown
	AccountDebitedDrawdown *AccountDebitedDrawdown `json:"accountDebitedDrawdown,omitempty"`
	// Originator
	Originator *Personal `json:"originator,omitempty"`
	// OriginatorOptionF
	OriginatorOptionF *OriginatorOptionF `json:"originatorOptionF,omitempty"`
	// OriginatorFI
	OriginatorFI *FinancialInstitution `json:"originatorFI,omitempty"`
	// InstructingFI
	InstructingFI *FinancialInstitution `json:"instructingFI,omitempty"`
	// AccountCreditedDrawdown
	AccountCreditedDrawdown *AccountCreditedDrawdown `json:"accountCreditedDrawdown,omitempty"`
	// OriginatorToBeneficiary
	OriginatorToBeneficiary *OriginatorToBeneficiary `json:"originatorToBeneficiary,omitempty"`
	// FiReceiverFI
	FiReceiverFI *FiToFi `json:"fiReceiverFI,omitempty"`
	// FiDrawdownDebitAccountAdvice
	FiDrawdownDebitAccountAdvice *Advice `json:"fiDrawdownDebitAccountAdvice,omitempty"`
	// FiIntermediaryFI
	FiIntermediaryFI *FiToFi `json:"fiIntermediaryFI,omitempty"`
	// FiIntermediaryFIAdvice
	FiIntermediaryFIAdvice *Advice `json:"fiIntermediaryFIAdvice,omitempty"`
	// FiBeneficiaryFI
	FiBeneficiaryFI *FiToFi `json:"fiBeneficiaryFI,omitempty"`
	// FiBeneficiaryFIAdvice
	FiBeneficiaryFIAdvice *Advice `json:"fiBeneficiaryFIAdvice,omitempty"`
	// FiBeneficiary
	FiBeneficiary *FiToFi `json:"fiBeneficiary,omitempty"`
	// FiBeneficiaryAdvic
	FiBeneficiaryAdvice *Advice `json:"fiBeneficiaryAdvice,omitempty"`
	// FiPaymentMethodToBeneficiary
	FiPaymentMethodToBeneficiary *FIPaymentMethodToBeneficiary `json:"fiPaymentMethodToBeneficiary,omitempty"`
	// FiAdditionalFiToFi
	FiAdditionalFiToFi *AdditionalFiToFi `json:"fiAdditionalFiToFi,omitempty"`
	// CurrencyInstructedAmount
	CurrencyInstructedAmount *CoverPayment `json:"currencyInstructedAmount,omitempty"`
	// OrderingCustomer
	OrderingCustomer *CoverPayment `json:"orderingCustomer,omitempty"`
	// OrderingInstitution
	OrderingInstitution *CoverPayment `json:"orderingInstitution,omitempty"`
	// IntermediaryInstitution
	IntermediaryInstitution *CoverPayment `json:"intermediaryInstitution,omitempty"`
	// InstitutionAccount
	InstitutionAccount *CoverPayment `json:"institutionAccount,omitempty"`
	// BeneficiaryCustomer
	BeneficiaryCustomer *CoverPayment `json:"beneficiaryCustomer,omitempty"`
	// Remittance
	Remittance *CoverPayment `json:"remittance,omitempty"`
	// SenderToReceiver
	SenderToReceiver *CoverPayment `json:"senderToReceiver,omitempty"`
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
	ActualAmountPaid *RemittanceAmount `json:"actualAmountPaid,omitempty"`
	// GrossAmountRemittanceDocument
	GrossAmountRemittanceDocument *RemittanceAmount `json:"grossAmountRemittanceDocument,omitempty"`
	// AmountNegotiatedDiscount
	AmountNegotiatedDiscount *RemittanceAmount `json:"amountNegotiatedDiscount,omitempty"`
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
