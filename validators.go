// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

//ToDo:  Review to see if we want something like model_*enum

package wire

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var (
	upperAlphanumericRegex = regexp.MustCompile(`[^ A-Z0-9!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
	alphanumericRegex      = regexp.MustCompile(`[^ \w!"#$%&'()*+,-.\\/:;<>=?@\[\]^_{}|~]+`)
)

// validator is common validation and formatting of golang types to WIRE type strings
type validator struct{}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	if alphanumericRegex.MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return ErrNonAlphanumeric
	}
	return nil
}

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
		SSIServiceMessage:
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
		// Reserved for market practice conventions.
		"0", "1", "2", "3", "4", "5", "6",
		// Reserved for bilateral agreements between Fedwire senders and receivers.
		"7", "8", "9":
		return nil
	}
	return ErrPaymentNotificationIndicator
}

func (v *validator) isTestProductionCode(code string) error {
	switch code {
	case
		EnvironmentTest,
		EnvironmentProduction:
		return nil
	}
	return ErrTestProductionCode
}

func (v *validator) isMessageDuplicationCode(code string) error {
	switch code {
	case
		MessageDuplicationOriginal,
		MessageDuplicationResend:
		return nil
	}
	return ErrMessageDuplicationCode
}

// isCentury validates a 2 digit century 20-29
func (v *validator) isCentury(s string) error {
	if s < "20" || s > "29" {
		return ErrValidCentury
	}
	return nil
}

// isYear validates a 2 digit year 00-99
func (v *validator) isYear(s string) error {
	if s < "00" || s > "99" {
		return ErrValidYear
	}
	return nil
}

// isMonth validates a 2 digit month 01-12
func (v *validator) isMonth(s string) error {
	switch s {
	case
		"01", "02", "03", "04", "05", "06",
		"07", "08", "09", "10", "11", "12":
		return nil
	}
	return ErrValidMonth
}

// isDay validates a 2 digit day based on a 2 digit month
// month 01-12 day 01-31 based on month
func (v *validator) isDay(m string, d string) error {
	// ToDo: Future Consideration Leap Year - not sure if cards actually have 0229
	switch m {
	// February
	case "02":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29":
			return nil
		}
	// April, June, September, November
	case "04", "06", "09", "11":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30":
			return nil
		}
	// January, March, May, July, August, October, December
	case "01", "03", "05", "07", "08", "10", "12":
		switch d {
		case
			"01", "02", "03", "04", "05", "06",
			"07", "08", "09", "10", "11", "12",
			"13", "14", "15", "16", "17", "18",
			"19", "20", "21", "22", "23", "24",
			"25", "26", "27", "28", "29", "30", "31":
			return nil
		}
	}
	return ErrValidDay
}

// validateDate will return the incoming string only if it matches a valid CCYYMMDD
// date format. (C=Century, Y=Year, M=Month, D=Day)
func (v *validator) validateDate(s string) string {
	if length := utf8.RuneCountInString(s); length != 8 {
		return ""
	}
	cc, yy, mm, dd := s[:2], s[2:4], s[4:6], s[4:8]

	if err := v.isCentury(cc); err != nil {
		return ""
	}
	if err := v.isYear(yy); err != nil {
		return ""
	}
	if err := v.isMonth(mm); err != nil {
		return ""
	}
	if err := v.isDay(mm, dd); err != nil {
		return ""
	}
	return fmt.Sprintf("%s%s%s%s", cc, yy, mm, dd)
}
