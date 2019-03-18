// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// FedWireMessage is a FedWire Message
type FedWireMessage struct {
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
	ReceiverDI ReceiverDepositoryInstitution `json:"ReceiverDepositoryInstitution"`
	// BusinessFunctionCode
	BusinessFunctionCode BusinessFunctionCode `json:"BusinessFunctionCode"`
}

// NewFedWireMessage returns a new FedWireMessage
func NewFedWireMessage() FedWireMessage {
	fwm := FedWireMessage{}
	return fwm
}
