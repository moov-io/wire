package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// CurrencyInstructedAmount creates a CurrencyInstructedAmount
func mockCurrencyInstructedAmount() *CurrencyInstructedAmount {
	cia := NewCurrencyInstructedAmount()
	cia.SwiftFieldTag = "Swift Field Tag"
	cia.Amount = "1500,49"
	return cia
}

// TestMockCurrencyInstructedAmount validates mockCurrencyInstructedAmount
func TestMockCurrencyInstructedAmount(t *testing.T) {
	cia := mockCurrencyInstructedAmount()

	require.NoError(t, cia.Validate(), "mockCurrencyInstructedAmount does not validate and will break other tests")
}

// TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric validates CurrencyInstructedAmount SwiftFieldTag is alphanumeric
func TestCurrencyInstructedAmountSwiftFieldTagAlphaNumeric(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.SwiftFieldTag = "®"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, cia.SwiftFieldTag).Error())
}

// TestCurrencyInstructedAmountValid validates CurrencyInstructedAmount Amount is valid
func TestCurrencyInstructedAmountValid(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.Amount = "1-0"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, cia.Amount).Error())
}

// TestParseCurrencyInstructedAmountWrongLength parses a wrong CurrencyInstructedAmount record length
func TestParseCurrencyInstructedAmountWrongLength(t *testing.T) {
	var line = "{7033}Swift*000000000001500,4"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCurrencyInstructedAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrRequireDelimiter)).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrRequireDelimiter)).Error())
}

// TestParseCurrencyInstructedAmountReaderParseError parses a wrong CurrencyInstructedAmount reader parse error
func TestParseCurrencyInstructedAmountReaderParseError(t *testing.T) {
	var line = "{7033}Swift*00000000Z001500,49*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCurrencyInstructedAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "00000000Z001500,49")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "00000000Z001500,49")).Error())
}

// TestCurrencyInstructedAmountTagError validates a CurrencyInstructedAmount tag
func TestCurrencyInstructedAmountTagError(t *testing.T) {
	cia := mockCurrencyInstructedAmount()
	cia.tag = "{9999}"

	err := cia.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, cia.tag).Error())
}

// TestStringCurrencyInstructedAmountVariableLength parses using variable length
func TestStringCurrencyInstructedAmountVariableLength(t *testing.T) {
	var line = "{7033}*000000000001500,49*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCurrencyInstructedAmount()
	require.NoError(t, err)

	line = "{7033}B                                                            NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseCurrencyInstructedAmount()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())
}

// TestStringCurrencyInstructedAmountOptions validates Format() formatted according to the FormatOptions
func TestStringCurrencyInstructedAmountOptions(t *testing.T) {
	var line = "{7033}*000000000001500,49*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCurrencyInstructedAmount()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.CurrencyInstructedAmount
	require.Equal(t, "{7033}     *000000000001500,49*", record.String())
	require.Equal(t, "{7033}*000000000001500,49*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
