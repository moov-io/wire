// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// ActualAmountPaid is the actual amount paid
type ActualAmountPaid struct {
	// tag
	tag string
	// RemittanceAmount is remittance amounts
	RemittanceAmount RemittanceAmount `json:"remittanceAmount,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewActualAmountPaid returns a new ActualAmountPaid
func NewActualAmountPaid() *ActualAmountPaid {
	aap := &ActualAmountPaid{
		tag: TagActualAmountPaid,
	}
	return aap
}

// Parse takes the input string and parses the ActualAmountPaid values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (aap *ActualAmountPaid) Parse(record string) error {
	if utf8.RuneCountInString(record) < 8 {
		return NewTagMinLengthErr(8, len(record))
	}

	aap.tag = record[:6]
	length := 6

	value, read, err := aap.parseVariableStringField(record[length:], 3)
	if err != nil {
		return fieldError("CurrencyCode", err)
	}
	aap.RemittanceAmount.CurrencyCode = value
	length += read

	value, read, err = aap.parseVariableStringField(record[length:], 19)
	if err != nil {
		return fieldError("Amount", err)
	}
	aap.RemittanceAmount.Amount = value
	length += read

	if !aap.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (aap *ActualAmountPaid) UnmarshalJSON(data []byte) error {
	type Alias ActualAmountPaid
	aux := struct {
		*Alias
	}{
		(*Alias)(aap),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	aap.tag = TagActualAmountPaid
	return nil
}

// String returns a fixed-width ActualAmountPaid record
func (aap *ActualAmountPaid) String() string {
	return aap.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns an ActualAmountPaid record formatted according to the FormatOptions
func (aap *ActualAmountPaid) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(28)

	buf.WriteString(aap.tag)
	buf.WriteString(aap.FormatCurrencyCode(options))
	buf.WriteString(aap.FormatAmount(options))

	return buf.String()
}

// Validate performs WIRE format rule checks on ActualAmountPaid and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// Currency Code and Amount are mandatory for each set of remittance data.
func (aap *ActualAmountPaid) Validate() error {
	if err := aap.fieldInclusion(); err != nil {
		return err
	}
	if aap.tag != TagActualAmountPaid {
		return fieldError("tag", ErrValidTagForType, aap.tag)
	}
	if err := aap.isCurrencyCode(aap.RemittanceAmount.CurrencyCode); err != nil {
		return fieldError("CurrencyCode", err, aap.RemittanceAmount.CurrencyCode)
	}
	if err := aap.isAmount(aap.RemittanceAmount.Amount); err != nil {
		return fieldError("Amount", err, aap.RemittanceAmount.Amount)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (aap *ActualAmountPaid) fieldInclusion() error {
	if aap.RemittanceAmount.Amount == "" {
		return fieldError("Amount", ErrFieldRequired)
	}
	if aap.RemittanceAmount.CurrencyCode == "" {
		return fieldError("CurrencyCode", ErrFieldRequired)

	}
	return nil
}

// CurrencyCodeField gets a string of the RemittanceAmount.CurrencyCode field
func (aap *ActualAmountPaid) CurrencyCodeField() string {
	return aap.alphaField(aap.RemittanceAmount.CurrencyCode, 3)
}

// AmountField gets a string of the RemittanceAmount.Amount field
func (aap *ActualAmountPaid) AmountField() string {
	return aap.alphaField(aap.RemittanceAmount.Amount, 19)
}

// FormatCurrencyCode returns RemittanceAmount.CurrencyCode formatted according to the FormatOptions
func (aap *ActualAmountPaid) FormatCurrencyCode(options FormatOptions) string {
	return aap.formatAlphaField(aap.RemittanceAmount.CurrencyCode, 3, options)
}

// FormatAmount returns RemittanceAmount.Amount formatted according to the FormatOptions
func (aap *ActualAmountPaid) FormatAmount(options FormatOptions) string {
	return aap.formatAlphaField(aap.RemittanceAmount.Amount, 19, options)
}
