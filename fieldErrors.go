// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"errors"
	"fmt"
)

var (
	// Errors specific to validation

	//ErrNonAlphanumeric is returned when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
	//ErrUpperAlpha is returned when a field is not in uppercase
	ErrUpperAlpha = errors.New("is not uppercase A-Z or 0-9")
	//ErrFieldInclusion is returned when a field is mandatory and has a default value
	ErrFieldInclusion = errors.New("is a mandatory field and has a default value")
	//ErrConstructor is returned when there's a mandatory field is not initialized correctly, and prompts to use the constructor
	ErrConstructor = errors.New("is a mandatory field and has a default value. Use the constructor")
	//ErrFieldRequired is returned when a field is required
	ErrFieldRequired = errors.New("is a required field")
	//ErrValidMonth is returned for an invalid month
	ErrValidMonth = errors.New("is an invalid month")
	//ErrValidDay is returned for an invalid day
	ErrValidDay = errors.New("is an invalid day")
	//ErrValidYear is returned for an invalid year
	ErrValidYear = errors.New("is an invalid year")
	//ErrValidCentury is returned for an invalid century
	ErrValidCentury = errors.New("is an invalid century")

	// SenderSupplied Tag {1500}

	// ErrFormatVersion is returned when there is an invalid FormatVersion
	ErrFormatVersion = errors.New("is not 30")
	// ErrTestProductionCode is returned when there is an invalid TestProductionCode
	ErrTestProductionCode = errors.New("is an invalid test production code")
	// ErrMessageDuplicationCode is returned when there is an invalid MessageDuplicationCode
	ErrMessageDuplicationCode = errors.New("is an invalid message duplication code")

	//ErrTypeCode is returned when there's an invalid TypeCode tag {1510}
	ErrTypeCode = errors.New("is an invalid type code")
	//ErrSubTypeCode is returned when there's an invalid SubTypeCode tag {1510}
	ErrSubTypeCode = errors.New("is an invalid sub type Code")

	// ErrLocalInstrumentCode is returned when there is an invalid LocalInstrumentCode tag {3610}
	ErrLocalInstrumentCode = errors.New("is an invalid local instrument Code")
	// ErrPaymentNotificationIndicator is returned when there is an invalid PaymentNotificationIndicator {3620}
	ErrPaymentNotificationIndicator = errors.New("is an invalid payment notification indicator")

	// ErrRecordSize is returned when there's an invalid record size
	ErrRecordSize = errors.New("is not 094")
)

// FieldError is returned for errors at a field level in a tag
type FieldError struct {
	FieldName string      // field name where error happened
	Value     interface{} // value that cause error
	Err       error       // context of the error.
	Msg       string      // deprecated
}

// Error message is constructed
// FieldName Msg Value
// ToDo:
// Example1:
// Example2:
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %v %s", e.FieldName, e.Value, e.Err)
}

// Unwrap implements the base.UnwrappableError interface for FieldError
func (e *FieldError) Unwrap() error {
	return e.Err
}

func fieldError(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*FieldError); ok {
		return err
	}
	fe := FieldError{
		FieldName: field,
		Err:       err,
	}
	// only the first value counts
	if len(values) > 0 {
		fe.Value = values[0]
	}
	return &fe
}
