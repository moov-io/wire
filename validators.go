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
