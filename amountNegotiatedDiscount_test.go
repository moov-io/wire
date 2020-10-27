package wire

import (
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

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAmountNegotiatedDiscountCurrencyCodeValid validates AmountNegotiatedDiscount CurrencyCode
func TestAmountNegotiatedDiscountCurrencyCodeValid(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "XZP"

	err := nd.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonCurrencyCode.Error())
}

// TestAmountNegotiatedDiscountAmountRequired validates AmountNegotiatedDiscount Amount is required
func TestAmountNegotiatedDiscountRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.Amount = ""

	err := nd.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Amount", ErrFieldRequired).Error(), err.Error())
}

// TestAmountNegotiatedDiscountCurrencyCodeRequired validates AmountNegotiatedDiscount CurrencyCode is required
func TestAmountNegotiatedDiscountCurrencyCodeRequired(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = ""

	err := nd.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("CurrencyCode", ErrFieldRequired).Error(), err.Error())
}

// TestParseAmountNegotiatedDiscountWrongLength parses a wrong AmountNegotiatedDiscount record length
func TestParseAmountNegotiatedDiscountWrongLength(t *testing.T) {
	var line = "{8550}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(28, len(r.line)).Error())
}

// TestParseAmountNegotiatedDiscountReaderParseError parses a wrong AmountNegotiatedDiscount reader parse error
func TestParseAmountNegotiatedDiscountReaderParseError(t *testing.T) {
	var line = "{8550}USD1234.56Z           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmountNegotiatedDiscount()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAmountNegotiatedDiscountTagError validates AmountNegotiatedDiscount tag
func TestAmountNegotiatedDiscountTagError(t *testing.T) {
	nd := mockAmountNegotiatedDiscount()
	nd.tag = "{9999}"

	err := nd.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, nd.tag).Error(), err.Error())
}
