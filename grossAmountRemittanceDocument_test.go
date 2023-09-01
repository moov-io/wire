package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// GrossAmountRemittanceDocument creates a GrossAmountRemittanceDocument
func mockGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	gard := NewGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = "USD"
	gard.RemittanceAmount.Amount = "1234.56"
	return gard
}

// TestMockGrossAmountRemittanceDocument validates mockGrossAmountRemittanceDocument
func TestMockGrossAmountRemittanceDocument(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()

	require.NoError(t, gard.Validate(), "mockGrossAmountRemittanceDocument does not validate and will break other tests")
}

// TestGrossAmountRemittanceAmountRequired validates GrossAmountRemittance Amount is required
func TestGrossAmountRemittanceAmountRequired(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.Amount = ""

	require.EqualError(t, gard.Validate(), fieldError("Amount", ErrFieldRequired).Error())
}

// TestGrossAmountRemittanceCurrencyCodeRequired validates GrossAmountRemittance CurrencyCode is required
func TestGrossAmountRemittanceCurrencyCodeRequired(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = ""

	require.EqualError(t, gard.Validate(), fieldError("CurrencyCode", ErrFieldRequired).Error())
}

// TestGrossAmountRemittanceAmountValid validates Amount
func TestGrossAmountRemittanceAmountValid(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.Amount = "X,"

	require.EqualError(t, gard.Validate(), fieldError("Amount", ErrNonAmount, gard.RemittanceAmount.Amount).Error())
}

// TestGrossAmountRemittanceCurrencyCodeValid validates CurrencyCode
func TestGrossAmountRemittanceCurrencyCodeValid(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = "XZP"

	require.EqualError(t, gard.Validate(), fieldError("CurrencyCode", ErrNonCurrencyCode, gard.RemittanceAmount.CurrencyCode).Error())
}

// TestParseGrossAmountRemittanceWrongLength parses a wrong GrossAmountRemittance record length
func TestParseGrossAmountRemittanceWrongLength(t *testing.T) {
	var line = "{8500}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseGrossAmountRemittanceDocument()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrRequireDelimiter)).Error())
}

// TestParseGrossAmountRemittanceReaderParseError parses a wrong GrossAmountRemittance reader parse error
func TestParseGrossAmountRemittanceReaderParseError(t *testing.T) {
	var line = "{8500}USD1234.56Z           *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseGrossAmountRemittanceDocument()

	expected := r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)
}

// TestGrossAmountRemittanceTagError validates a GrossAmountRemittance tag
func TestGrossAmountRemittanceTagError(t *testing.T) {
	gard := mockGrossAmountRemittanceDocument()
	gard.tag = "{9999}"

	require.EqualError(t, gard.Validate(), fieldError("tag", ErrValidTagForType, gard.tag).Error())
}

// TestStringGrossAmountRemittanceDocumentVariableLength parses using variable length
func TestStringGrossAmountRemittanceDocumentVariableLength(t *testing.T) {
	var line = "{8500}USD1234.56*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseGrossAmountRemittanceDocument()
	require.Nil(t, err)

	line = "{8500}USD1234.56            NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseGrossAmountRemittanceDocument()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{8500}USD1234.56***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseGrossAmountRemittanceDocument()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8500}USD1234.56*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseGrossAmountRemittanceDocument()
	require.Equal(t, err, nil)
}

// TestStringGrossAmountRemittanceDocumentOptions validates Format() formatted according to the FormatOptions
func TestStringGrossAmountRemittanceDocumentOptions(t *testing.T) {
	var line = "{8500}USD1234.56*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseGrossAmountRemittanceDocument()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.GrossAmountRemittanceDocument
	require.Equal(t, record.String(), "{8500}USD1234.56            *")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{8500}USD1234.56*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
