package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockInstructedAmount creates a InstructedAmount
func mockInstructedAmount() *InstructedAmount {
	ia := NewInstructedAmount()
	ia.CurrencyCode = "USD"
	ia.Amount = "4567,89"
	return ia
}

// TestMockInstructedAmount validates mockInstructedAmount
func TestMockInstructedAmount(t *testing.T) {
	ia := mockInstructedAmount()

	require.NoError(t, ia.Validate(), "mockInstructedAmount does not validate and will break other tests")
}

// TestInstructedAmountAmountRequired validates InstructedAmount Amount is required
func TestInstructedAmountRequired(t *testing.T) {
	ia := mockInstructedAmount()
	ia.Amount = ""

	err := ia.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

// TestInstructedAmountCurrencyCodeRequired validates InstructedAmount CurrencyCode is required
func TestInstructedAmountCurrencyCodeRequired(t *testing.T) {
	ia := mockInstructedAmount()
	ia.CurrencyCode = ""

	err := ia.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrFieldRequired).Error())
}

// TestInstructedAmountAmountValid validates Amount
func TestInstructedAmountValid(t *testing.T) {
	ia := mockInstructedAmount()
	ia.Amount = "X,"

	err := ia.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, ia.Amount).Error())
}

// TestInstructedAmountCurrencyCodeValid validates Amount
func TestInstructedAmountCurrencyCodeValid(t *testing.T) {
	ia := mockInstructedAmount()
	ia.CurrencyCode = "XZP"

	err := ia.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrNonCurrencyCode, ia.CurrencyCode).Error())
}

// TestParseInstructedAmountWrongLength parses a wrong InstructedAmount record length
func TestParseInstructedAmountWrongLength(t *testing.T) {
	var line = "{3710}USD4567,89"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructedAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrRequireDelimiter)).Error())
}

// TestParseInstructedAmountReaderParseError parses a wrong InstructedAmount reader parse error
func TestParseInstructedAmountReaderParseError(t *testing.T) {
	var line = "{3710}USD000000004567Z89*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructedAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "000000004567Z89")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrNonAmount, "000000004567Z89")).Error())
}

// TestInstructedAmountTagError validates a InstructedAmount tag
func TestInstructedAmountTagError(t *testing.T) {
	ia := mockInstructedAmount()
	ia.tag = "{9999}"

	require.EqualError(t, ia.Validate(), fieldError("tag", ErrValidTagForType, ia.tag).Error())
}

// TestStringInstructedAmountVariableLength parses using variable length
func TestStringInstructedAmountVariableLength(t *testing.T) {
	var line = "{3710}USD4567,89*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructedAmount()
	require.NoError(t, err)

	line = "{3710}USD4567,89        NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstructedAmount()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{3710}USD4567,89***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstructedAmount()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{3710}USD4567,89*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstructedAmount()
	require.NoError(t, err)
}

// TestStringInstructedAmountOptions validates Format() formatted according to the FormatOptions
func TestStringInstructedAmountOptions(t *testing.T) {
	var line = "{3710}USD4567,89*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructedAmount()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.InstructedAmount
	require.Equal(t, "{3710}USD4567,89        *", record.String())
	require.Equal(t, "{3710}USD4567,89*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
