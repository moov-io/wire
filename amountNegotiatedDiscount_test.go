package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// AmountNegotiatedDiscount creates a AmountNegotiatedDiscount
func mockAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	nd := NewAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "USD"
	nd.RemittanceAmount.Amount = "1234.56"
	return nd
}

// TestMockAmountNegotiatedDiscount validates mockAmountNegotiatedDiscount
func TestMockAmountNegotiatedDiscount(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()

	require.NoError(t, nd.Validate(), "mockAmountNegotiatedDiscount does not validate and will break other tests")
}

// TestAmountNegotiatedDiscountAmountValid validates AmountNegotiatedDiscount Amount
func TestAmountNegotiatedDiscountValid(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.Amount = "X,"

	err := nd.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, nd.RemittanceAmount.Amount).Error())
}

// TestAmountNegotiatedDiscountCurrencyCodeValid validates AmountNegotiatedDiscount CurrencyCode
func TestAmountNegotiatedDiscountCurrencyCodeValid(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "XZP"

	err := nd.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrNonCurrencyCode, nd.RemittanceAmount.CurrencyCode).Error())
}

// TestAmountNegotiatedDiscountAmountRequired validates AmountNegotiatedDiscount Amount is required
func TestAmountNegotiatedDiscountRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.Amount = ""

	err := nd.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

// TestAmountNegotiatedDiscountCurrencyCodeRequired validates AmountNegotiatedDiscount CurrencyCode is required
func TestAmountNegotiatedDiscountCurrencyCodeRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = ""

	err := nd.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrFieldRequired).Error())
}

// TestParseAmountNegotiatedDiscountWrongLength parses a wrong AmountNegotiatedDiscount record length
func TestParseAmountNegotiatedDiscountWrongLength(t *testing.T) {
	var line = "{8550}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrRequireDelimiter)).Error())
}

// TestParseAmountNegotiatedDiscountReaderParseError parses a wrong AmountNegotiatedDiscount reader parse error
func TestParseAmountNegotiatedDiscountReaderParseError(t *testing.T) {
	var line = "{8550}USD1234.56Z           *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()

	expected := r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)
}

// TestAmountNegotiatedDiscountTagError validates AmountNegotiatedDiscount tag
func TestAmountNegotiatedDiscountTagError(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.tag = "{9999}"

	err := nd.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, nd.tag).Error())
}

// TestStringAmountNegotiatedDiscountVariableLength parses using variable length
func TestStringAmountNegotiatedDiscountVariableLength(t *testing.T) {
	var line = "{8600}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()
	expected := r.parseError(NewTagMinLengthErr(8, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{8550}USD1234.56          NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAmountNegotiatedDiscount()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{8550}USD1234.56***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAmountNegotiatedDiscount()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8550}USD1234.56*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAmountNegotiatedDiscount()
	require.Equal(t, err, nil)
}

// TestStringAmountNegotiatedDiscountOptions validates Format() formatted according to the FormatOptions
func TestStringAmountNegotiatedDiscountOptions(t *testing.T) {
	var line = "{8550}USD1234.56*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()
	require.Equal(t, err, nil)

	and := r.currentFEDWireMessage.AmountNegotiatedDiscount
	require.Equal(t, and.String(), "{8550}USD1234.56            *")
	require.Equal(t, and.Format(FormatOptions{VariableLengthFields: true}), "{8550}USD1234.56*")
	require.Equal(t, and.String(), and.Format(FormatOptions{VariableLengthFields: false}))
}
