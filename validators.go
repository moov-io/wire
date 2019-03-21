// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

//ToDo:  Review to see if we want something like model_*enum

package wire

import (
	"regexp"
)

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
)

// validator is common validation and formatting of golang types to WIRE type strings
type validator struct{}

// isTypeCode ensures tag {1510} TypeCode is valid
func (v *validator) isTypeCode(code string) error {
	switch code {
	case
		FundsTransfer,
		ForeignTransfer,
		SettlementTransfer:
		return nil
	}
	return ErrTypeCode
}

// isSubTypeCode ensures tag {1510} SubTypeCode is valid
func (v *validator) isSubTypeCode(code string) error {
	switch code {
	case
		BasicFundsTransfer,
		RequestReversal,
		ReversalTransfer,
		RequestReversalPriorDayTransfer,
		ReversalPriorDayTransfer,
		RequestCredit,
		FundsTransferRequestCredit,
		RefusalRequestCredit,
		ServiceMessage:
		return nil
	}
	return ErrSubTypeCode
}

func (v *validator) isLocalInstrumentCode(code string) error {
	switch code {
	case
		// ANSI X12 format
		"ANSI",
		// Sequence B Cover Payment Structured
		"COVS",
		// General XML format
		"GXML",
		// ISO 20022 XML formaT
		"IXML",
		// Narrative Text
		"NARR",
		// Proprietary Local Instrument Code
		"PROP",
		//  Remittance Information Structured
		"RMTS",
		// Related Remittance Information
		"RRMT",
		// STP 820 format
		"S820",
		// SWIFT field 70
		"SWIF",
		// UN/EDIFACT format
		"UEDI":
		return nil
	}
	return ErrLocalInstrumentCode
}

func (v *validator) isPaymentNotificationIndicator(code string) error {
	switch code {
	case
		// * `0 - 6` - Reserved for market practice conventions.
		"0", "1", "2", "3", "4", "5", "6",
		// * `7 - 9` - Reserved for bilateral agreements between Fedwire senders and receivers.
		"7", "8", "9":
		return nil
	}
	return ErrPaymentNotificationIndicator
}