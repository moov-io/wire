package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBusinessFunctionCode creates a BusinessFunctionCode
func mockBusinessFunctionCode() *BusinessFunctionCode {
	bfc := NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "   "
	return bfc
}

// TestMockBusinessFunctionCode validates mockBusinessFunctionCode
func TestMockBusinessFunctionCode(t *testing.T) {
	bfc := mockBusinessFunctionCode()

	require.NoError(t, bfc.Validate(), "mockBusinessFunctionCode does not validate and will break other tests")
}

// TestBusinessFunctionCodeValid validates BusinessFunctionCode
func TestBusinessFunctionCodeValid(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = "ZZZ"

	err := bfc.Validate()

	require.EqualError(t, err, fieldError("BusinessFunctionCode", ErrBusinessFunctionCode, bfc.BusinessFunctionCode).Error())
}

// TestBusinessFunctionCodeRequired validates BusinessFunctionCode is required
func TestBusinessFunctionCodeRequired(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = ""

	err := bfc.Validate()

	require.EqualError(t, err, fieldError("BusinessFunctionCode", ErrFieldRequired, bfc.BusinessFunctionCode).Error())
}

// TestParseBusinessFunctionCodeWrongLength parses a wrong BusinessFunctionCode record length
func TestParseBusinessFunctionCodeWrongLength(t *testing.T) {
	var line = "{3600}CT"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()

	require.EqualError(t, err, r.parseError(NewTagMinLengthErr(9, len(r.line))).Error())
}

// TestParseBusinessFunctionCodeReaderParseError parses a wrong BusinessFunctionCode reader parse error
func TestParseBusinessFunctionCodeReaderParseError(t *testing.T) {
	var line = "{3600}CTAXXY*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()

	expected := r.parseError(fieldError("BusinessFunctionCode", ErrBusinessFunctionCode, "CTA")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("BusinessFunctionCode", ErrBusinessFunctionCode, "CTA")).Error()
	require.EqualError(t, err, expected)
}

// TestBusinessFunctionCodeTagError validates a BusinessFunctionCode tag
func TestBusinessFunctionCodeTagError(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.tag = "{9999}"

	err := bfc.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, bfc.tag).Error())
}

func TestBusinessFunctionCodeDelimiter(t *testing.T) {
	bfc := NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	require.Equal(t, "{3600}CTR", bfc.String())
	require.Equal(t, "{3600}CTR", bfc.Format(FormatOptions{VariableLengthFields: true}))

	bfc.TransactionTypeCode = "23"
	require.Equal(t, "{3600}CTR23 *", bfc.String())
	require.Equal(t, "{3600}CTR23*", bfc.Format(FormatOptions{VariableLengthFields: true}))
}

// TestStringBusinessFunctionCodeVariableLength parses using variable length
func TestStringBusinessFunctionCodeVariableLength(t *testing.T) {
	var line = "{3600}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()
	expected := r.parseError(NewTagMinLengthErr(9, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{3600}BTR   NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBusinessFunctionCode()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{3600}BTR***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBusinessFunctionCode()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{3600}BTR*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBusinessFunctionCode()
	require.NoError(t, err)
}

// TestStringBusinessFunctionCodeOptions validates Format() formatted according to the FormatOptions
func TestStringBusinessFunctionCodeOptions(t *testing.T) {
	var line = "{3600}BTR"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()
	require.NoError(t, err)

	bfc := r.currentFEDWireMessage.BusinessFunctionCode
	require.Equal(t, "{3600}BTR", bfc.String())
	require.Equal(t, "{3600}BTR", bfc.Format(FormatOptions{VariableLengthFields: true}))

	require.Equal(t, "{3600}BTR", bfc.Format(FormatOptions{VariableLengthFields: false}))
	require.Equal(t, bfc.Format(FormatOptions{VariableLengthFields: false}), bfc.String())
}
