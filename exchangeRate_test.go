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

	require.NotNil(t, err)
	require.Equal(t, fieldError("ExchangeRate", ErrNonAmount, eRate.ExchangeRate).Error(), err.Error())
}

// TestParseExchangeRateWrongLength parses a wrong ExchangeRate record length
func TestParseExchangeRateWrongLength(t *testing.T) {
	var line = "{3720}1,2345"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseExchangeRate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(18, len(r.line)).Error())
}

// TestParseExchangeRateReaderParseError parses a wrong ExchangeRate reader parse error
func TestParseExchangeRateReaderParseError(t *testing.T) {
	var line = "{3720}1,2345Z     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseExchangeRate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestExchangeRateTagError validates a ExchangeRate tag
func TestExchangeRateTagError(t *testing.T) {
	eRate := mockCurrencyInstructedAmount()
	eRate.tag = "{9999}"

	err := eRate.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, eRate.tag).Error(), err.Error())
}
