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

	require.EqualError(t, err, fieldError("AdviceCode", ErrAdviceCode, debitDDAdvice.Advice.AdviceCode).Error())
}

// TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric validates FIDrawdownDebitAccountAdvice LineOne is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineOne = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, debitDDAdvice.Advice.LineOne).Error())
}

// TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric validates FIDrawdownDebitAccountAdvice LineTwo is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineTwo = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, debitDDAdvice.Advice.LineTwo).Error())
}

// TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric validates FIDrawdownDebitAccountAdvice LineThree is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineThree = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, debitDDAdvice.Advice.LineThree).Error())
}

// TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFour is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFour = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, debitDDAdvice.Advice.LineFour).Error())
}

// TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFive is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFive = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, debitDDAdvice.Advice.LineFive).Error())
}

// TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric validates FIDrawdownDebitAccountAdvice LineSix is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineSix = "®"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, debitDDAdvice.Advice.LineSix).Error())
}

// TestParseFIDrawdownDebitAccountAdviceWrongLength parses a wrong FIDrawdownDebitAccountAdvice record length
func TestParseFIDrawdownDebitAccountAdviceWrongLength(t *testing.T) {
	var line = "{6110}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIDrawdownDebitAccountAdvice()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(200, len(r.line))).Error())
}

// TestParseFIDrawdownDebitAccountAdviceReaderParseError parses a wrong FIDrawdownDebitAccountAdvice reader parse error
func TestParseFIDrawdownDebitAccountAdviceReaderParseError(t *testing.T) {
	var line = "{6110}LTR®ine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIDrawdownDebitAccountAdvice()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error()
	require.EqualError(t, err, expected)
}

// TestFIDrawdownDebitAccountAdviceTagError validates a FIDrawdownDebitAccountAdvice tag
func TestFIDrawdownDebitAccountAdviceTagError(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.tag = "{9999}"

	err := debitDDAdvice.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, debitDDAdvice.tag).Error())
}
