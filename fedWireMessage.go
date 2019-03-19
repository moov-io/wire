// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

//ToDo: omitEmpty

// FedWireMessage is a FedWire Message
type FedWireMessage struct {
	// ID
	ID string `json:"id"`
	// SenderSuppliedInformation
	SenderSuppliedInformation SenderSuppliedInformation `json:"senderSuppliedInformation"`
	// TypeSubType
	TypeSubType TypeSubType `json:"typeSubType"`
	// InputMessageAccountabilityData (IMAD)
	InputMessageAccountabilityData InputMessageAccountabilityData `json:"inputMessageAccountabilityData"`
	// Amount (up to a penny less than $10 billion)
	Amount Amount `json:"amount"`
	// SenderDepositoryInstitution
	SenderDI SenderDepositoryInstitution `json:"senderDepositoryInstitution"`
	// ReceiverDepositoryInstitution
	ReceiverDI ReceiverDepositoryInstitution `json:"receiverDepositoryInstitution"`
	// BusinessFunctionCode
	BusinessFunctionCode BusinessFunctionCode `json:"businessFunctionCode"`
	// BeneficiaryIntermediaryFI
	BeneficiaryIntermediaryFI BeneficiaryIntermediaryFI  `json:"beneficiaryIntermediaryFI"`
	// BeneficiaryFI
	BeneficiaryFI BeneficiaryFI  `json:"beneficiaryFI"`
	// Beneficiary
	Beneficiary Beneficiary  `json:"beneficiary, "`
	// BeneficiaryReference
	BeneficiaryReference BeneficiaryReference `json:"beneficiaryReference, omitEmpty"`
}

// NewFedWireMessage returns a new FedWireMessage
func NewFedWireMessage() FedWireMessage {
	fwm := FedWireMessage{}
	return fwm
}
