// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// ToDo: Evaluate use of {}
// To Do see if we want these constants

const (
	// TagSenderSuppliedInformation is SenderSuppliedInformation
	TagSenderSuppliedInformation  = "{1500}"
	// TagTypeSubType is TypeSubType
	TagTypeSubType = "{1510}"
	// TagInputMessageAccountabilityData is InputMessageAccountabilityData (IMAD)
	TagInputMessageAccountabilityData = "{1520}"
	// TagAmount is Amount
	TagAmount = "{2000}"
	// TagSenderDepositoryInstitution is SenderDepositoryInstitution
	TagSenderDepositoryInstitution = "{3100}"
	// TagReceiverDepositoryInstitution is ReceiverDepositoryInstitution
	TagReceiverDepositoryInstitution = "{3400}"
	// TagBusinessFunctionCode is BusinessFunctionCode
	TagBusinessFunctionCode = "{3600}"
	// TagBeneficiaryIntermediaryFI is BeneficiaryIntermediaryFI
	TagBeneficiaryIntermediaryFI = "{4000}"
	// TagBeneficiaryFI is BeneficiaryFI
	TagBeneficiaryFI = "{4100}"
	// TagBeneficiary is Beneficiary
	TagBeneficiary = "{4200}"
	// TagBeneficiaryReference  is BeneficiaryReference
	TagBeneficiaryReference = "{4320}"
	// TagTagAccountDebitedDrawdown is AccountDebitedDrawdown
	TagAccountDebitedDrawdown = "{4400}"
	// TagOriginator is Originator
	TagOriginator = "{5000}"

	// ToDo: Additional Tags
	// TagSenderReference is SenderReference
	// TagSenderReference = 3320
	// 3000
	// 4200
	// 5000
	// 6000, etc...

	// TagSenderSuppliedInformation

	// FormatVersion designates the FEDWIRE message format version
	FormatVersion = "30"
	// EnvironmentTest designates a test environment
	EnvironmentTest = "T"
	// EnvironmentProduction designates a production environment
	EnvironmentProduction = "P"
	// MessageDuplicationOriginal designates an original message
	MessageDuplicationOriginal = " "
	// MessageDuplicationResend designates a resend of a message
	MessageDuplicationResend = "P"

	// TagTypeSubType TypeCode

	// FundsTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer in which the
	// sender and/or receiver may be a bank or a third party (i.e., customer of a bank).
	FundsTransfer = "10"
	// ForeignTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer to
	// or from a foreign central bank or government or international organization with an account at the Federal
	// Reserve Bank of New York.
	ForeignTransfer = "15"
	// SettlementTransfer is SenderSuppliedInformation {1510} TypeCode which designates a funds transfer
	// between Fedwire Funds Service participants.
	SettlementTransfer = "30"

	// TagTypeSubType SubTypeCode

	// BasicFundsTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a basic value funds transfer.
	BasicFundsTransfer = "00"
	// RequestReversal is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value request for
	// reversal of a funds transfer
	// originated on the current business day.
	RequestReversal = "01"
	// ReversalTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a value reversal of a
	// funds transfer received on the current business day.  May be used in response to a subtype
	// code 01 Request for Reversal.
	ReversalTransfer = "02"
	// RequestReversalPriorDayTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value
	// request for a reversal of a funds transfer originated on a prior business day.
	RequestReversalPriorDayTransfer = "07"
	// ReversalPriorDayTransfer is SenderSuppliedInformation {1510} SubTypeCode which designates a value reversal of
	// a funds transfer received on a prior business day.  May be used in response to a subtype code 07 Request for
	// Reversal of a Prior Day Transfer.
	ReversalPriorDayTransfer = "08"
	// RequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value request for the
	// receiver to send a funds transfer to a designated party.
	RequestCredit = "31"
	// FundsTransferRequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a value funds
	// transfer honoring a subtype 31 request for credit.
	FundsTransferRequestCredit = "32"
	// RefusalRequestCredit is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value message
	// indicating refusal to honor a subtype 31 request for credit.
	RefusalRequestCredit = "33"
	// ServiceMessage is SenderSuppliedInformation {1510} SubTypeCode which designates a non-value message used to
	// communicate questions and information that is not covered by a specific subtype.
	ServiceMessage = "90"
)
