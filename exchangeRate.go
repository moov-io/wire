// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// ExchangeRate is the ExchangeRate of the wire
type ExchangeRate struct {
	// tag
	tag string
	// ExchangeRate is the exchange rate
	// Must contain at least one numeric character and only one decimal comma marker (e.g., an exchange rate of 1.2345 should be entered as 1,2345).
	ExchangeRate string `json:"exchangeRate,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewExchangeRate returns a new ExchangeRate
func NewExchangeRate() ExchangeRate  {
	er := ExchangeRate {
		tag: TagExchangeRate,
	}
	return er
}

// Parse takes the input string and parses the ExchangeRate values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (er *ExchangeRate) Parse(record string) {
}

// String writes ExchangeRate
func (er*ExchangeRate) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(12)
	buf.WriteString(er.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on ExchangeRate and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (er *ExchangeRate) Validate() error {
	if err := er.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (er *ExchangeRate) fieldInclusion() error {
	return nil
}
