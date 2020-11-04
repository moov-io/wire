package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockExchangeRate creates a ExchangeRate
func mockExchangeRate() *ExchangeRate {
	eRate := NewExchangeRate()
	eRate.ExchangeRate = "1,2345"
	return eRate
}

// TestMockExchangeRate validates mockExchangeRate
func TestMockExchangeRate(t *testing.T) {
	eRate := mockExchangeRate()

	require.NoError(t, eRate.Validate(), "mockExchangeRate does not validate and will break other tests")
}

// TestExchangeRate validates ExchangeRate
func TestExchangeRateNumeric(t *testing.T) {
	eRate := mockExchangeRate()
	eRate.ExchangeRate = "1,--0.00"

	err := eRate.Validate()

	require.EqualError(t, err, fieldError("ExchangeRate", ErrNonAmount, eRate.ExchangeRate).Error())
}

// TestParseExchangeRateWrongLength parses a wrong ExchangeRate record length
func TestParseExchangeRateWrongLength(t *testing.T) {
	var line = "{3720}1,2345"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseExchangeRate()

	expected := r.parseError(NewTagWrongLengthErr(18, len(r.line))).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(NewTagWrongLengthErr(18, len(r.line))).Error()
	require.EqualError(t, err, expected)
}

// TestParseExchangeRateReaderParseError parses a wrong ExchangeRate reader parse error
func TestParseExchangeRateReaderParseError(t *testing.T) {
	var line = "{3720}1,2345Z     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseExchangeRate()

	require.EqualError(t, err, r.parseError(fieldError("ExchangeRate", ErrNonAmount, "1,2345Z")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("ExchangeRate", ErrNonAmount, "1,2345Z")).Error())
}

// TestExchangeRateTagError validates a ExchangeRate tag
func TestExchangeRateTagError(t *testing.T) {
	eRate := mockCurrencyInstructedAmount()
	eRate.tag = "{9999}"

	err := eRate.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, eRate.tag).Error())
}
