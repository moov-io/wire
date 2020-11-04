package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIIntermediaryFIAdvice creates a FIIntermediaryFIAdvice
func mockFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	fiifia := NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = AdviceCodeLetter
	fiifia.Advice.LineOne = "Line One"
	fiifia.Advice.LineTwo = "Line Two"
	fiifia.Advice.LineThree = "Line Three"
	fiifia.Advice.LineFour = "Line Four"
	fiifia.Advice.LineFive = "Line Five"
	fiifia.Advice.LineSix = "Line Six"
	return fiifia
}

// TestMockFIIntermediaryFIAdvice validates mockFIIntermediaryFIAdvice
func TestMockFIIntermediaryFIAdvice(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()

	require.NoError(t, fiifia.Validate(), "mockFIIntermediaryFIAdvice does not validate and will break other tests")
}

// TestFIIntermediaryFIAdviceAdviceCodeValid validates FIIntermediaryFIAdvice AdviceCode is alphanumeric
func TestFIIntermediaryFIAdviceAdviceCodeValid(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = "Z"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("AdviceCode", ErrAdviceCode, fiifia.Advice.AdviceCode).Error())
}

// TestFIIntermediaryFIAdviceLineOneAlphaNumeric validates FIIntermediaryFIAdvice LineOne is alphanumeric
func TestFIIntermediaryFIAdviceLineOneAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineOne = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fiifia.Advice.LineOne).Error())
}

// TestFIIntermediaryFIAdviceLineTwoAlphaNumeric validates FIIntermediaryFIAdvice LineTwo is alphanumeric
func TestFIIntermediaryFIAdviceLineTwoAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineTwo = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fiifia.Advice.LineTwo).Error())
}

// TestFIIntermediaryFIAdviceLineThreeAlphaNumeric validates FIIntermediaryFIAdvice LineThree is alphanumeric
func TestFIIntermediaryFIAdviceLineThreeAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineThree = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fiifia.Advice.LineThree).Error())
}

// TestFIIntermediaryFIAdviceLineFourAlphaNumeric validates FIIntermediaryFIAdvice LineFour is alphanumeric
func TestFIIntermediaryFIAdviceLineFourAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineFour = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fiifia.Advice.LineFour).Error())
}

// TestFIIntermediaryFIAdviceLineFiveAlphaNumeric validates FIIntermediaryFIAdvice LineFive is alphanumeric
func TestFIIntermediaryFIAdviceLineFiveAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineFive = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fiifia.Advice.LineFive).Error())
}

// TestFIIntermediaryFIAdviceLineSixAlphaNumeric validates FIIntermediaryFIAdvice LineSix is alphanumeric
func TestFIIntermediaryFIAdviceLineSixAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineSix = "®"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fiifia.Advice.LineSix).Error())
}

// TestParseFIIntermediaryFIAdviceWrongLength parses a wrong FIIntermediaryFIAdvice record length
func TestParseFIIntermediaryFIAdviceWrongLength(t *testing.T) {
	var line = "{6210}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIIntermediaryFIAdvice()
	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(200, len(r.line))).Error())
}

// TestParseFIIntermediaryFIAdviceReaderParseError parses a wrong FIIntermediaryFIAdvice reader parse error
func TestParseFIIntermediaryFIAdviceReaderParseError(t *testing.T) {
	var line = "{6210}LTRLine ®ne                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIIntermediaryFIAdvice()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)
}

// TestFIIntermediaryFIAdviceTagError validates a FIIntermediaryFIAdvice tag
func TestFIIntermediaryFIAdviceTagError(t *testing.T) {
	fiifia := mockFIIntermediaryFI()
	fiifia.tag = "{9999}"

	err := fiifia.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fiifia.tag).Error())
}
