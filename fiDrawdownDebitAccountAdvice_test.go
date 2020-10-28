package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIDrawdownDebitAccountAdvice creates a FIDrawdownDebitAccountAdvice
func mockFIDrawdownDebitAccountAdvice() *FIDrawdownDebitAccountAdvice {
	debitDDAdvice := NewFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.AdviceCode = AdviceCodeLetter
	debitDDAdvice.Advice.LineOne = "Line One"
	debitDDAdvice.Advice.LineTwo = "Line Two"
	debitDDAdvice.Advice.LineThree = "Line Three"
	debitDDAdvice.Advice.LineFour = "Line Four"
	debitDDAdvice.Advice.LineFive = "Line Five"
	debitDDAdvice.Advice.LineSix = "Line Six"
	return debitDDAdvice
}

// TestMockFIDrawdownDebitAccountAdvice validates mockFIDrawdownDebitAccountAdvice
func TestMockFIDrawdownDebitAccountAdvice(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()

	require.NoError(t, debitDDAdvice.Validate(), "mockFIDrawdownDebitAccountAdvice does not validate and will break other tests")
}

// TestFIDrawdownDebitAccountAdviceAdviceCodeValid validates FIDrawdownDebitAccountAdvice AdviceCode is alphanumeric
func TestFIDrawdownDebitAccountAdviceCodeValid(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.AdviceCode = "Z"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AdviceCode", ErrAdviceCode, debitDDAdvice.Advice.AdviceCode).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric validates FIDrawdownDebitAccountAdvice LineOne is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineOne = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineOne", ErrNonAlphanumeric, debitDDAdvice.Advice.LineOne).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric validates FIDrawdownDebitAccountAdvice LineTwo is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineTwo = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineTwo", ErrNonAlphanumeric, debitDDAdvice.Advice.LineTwo).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric validates FIDrawdownDebitAccountAdvice LineThree is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineThree = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineThree", ErrNonAlphanumeric, debitDDAdvice.Advice.LineThree).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFour is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFour = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFour", ErrNonAlphanumeric, debitDDAdvice.Advice.LineFour).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFive is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFive = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFive", ErrNonAlphanumeric, debitDDAdvice.Advice.LineFive).Error(), err.Error())
}

// TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric validates FIDrawdownDebitAccountAdvice LineSix is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineSix = "®"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineSix", ErrNonAlphanumeric, debitDDAdvice.Advice.LineSix).Error(), err.Error())
}

// TestParseFIDrawdownDebitAccountAdviceWrongLength parses a wrong FIDrawdownDebitAccountAdvice record length
func TestParseFIDrawdownDebitAccountAdviceWrongLength(t *testing.T) {
	var line = "{6110}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIDrawdownDebitAccountAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(200, len(r.line)).Error(), err.Error())
}

// TestParseFIDrawdownDebitAccountAdviceReaderParseError parses a wrong FIDrawdownDebitAccountAdvice reader parse error
func TestParseFIDrawdownDebitAccountAdviceReaderParseError(t *testing.T) {
	var line = "{6110}LTR®ine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIDrawdownDebitAccountAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestFIDrawdownDebitAccountAdviceTagError validates a FIDrawdownDebitAccountAdvice tag
func TestFIDrawdownDebitAccountAdviceTagError(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.tag = "{9999}"

	err := debitDDAdvice.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, debitDDAdvice.tag).Error(), err.Error())
}
