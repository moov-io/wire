package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// GrossAmountRemittanceDocument creates a GrossAmountRemittanceDocument
func mockGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	gard := NewGrossAmountRemittanceDocument(false)
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

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(28, len(r.line))).Error())
}

// TestParseGrossAmountRemittanceReaderParseError parses a wrong GrossAmountRemittance reader parse error
func TestParseGrossAmountRemittanceReaderParseError(t *testing.T) {
	var line = "{8500}USD1234.56Z           "
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
