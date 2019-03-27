// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// AccountCreditedDrawdown is the account which is credited in a drawdown
type AccountCreditedDrawdown struct {
	// tag
	tag string
	// DrawdownCreditAccountNumber  9 character ABA
	DrawdownCreditAccountNumber string `json:"drawdownCreditAccountNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewAccountCreditedDrawdown returns a new AccountCreditedDrawdown
func NewAccountCreditedDrawdown() AccountCreditedDrawdown {
	creditDD := AccountCreditedDrawdown{
		tag: TagAccountCreditedDrawdown,
	}
	return creditDD
}

// Parse takes the input string and parses the AccountCreditedDrawdown values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (creditDD *AccountCreditedDrawdown) Parse(record string) {
}

// String writes AccountCreditedDrawdown
func (creditDD *AccountCreditedDrawdown) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(175)
	buf.WriteString(creditDD.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on AccountCreditedDrawdown and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (creditDD *AccountCreditedDrawdown) Validate() error {
	if err := creditDD.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (creditDD *AccountCreditedDrawdown) fieldInclusion() error {
	return nil
}
