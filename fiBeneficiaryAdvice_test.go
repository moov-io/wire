package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiaryAdvice creates a FIBeneficiaryAdvice
func mockFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	fiba := NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	return fiba
}

// TestMockFIBeneficiaryAdvice validates mockFIBeneficiaryAdvice
func TestMockFIBeneficiaryAdvice(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()

	require.NoError(t, fiba.Validate(), "mockFIBeneficiaryAdvice does not validate and will break other tests")
}

// TestFIBeneficiaryAdviceCodeValid validates FIBeneficiaryAdvice AdviceCode is alphanumeric
func TestFIBeneficiaryAdviceCodeValid(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = "Z"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AdviceCode", ErrAdviceCode, fiba.Advice.AdviceCode).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineOneAlphaNumeric validates FIBeneficiaryAdvice LineOne is alphanumeric
func TestFIBeneficiaryAdviceLineOneAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineOne = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineOne", ErrNonAlphanumeric, fiba.Advice.LineOne).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineTwoAlphaNumeric validates FIBeneficiaryAdvice LineTwo is alphanumeric
func TestFIBeneficiaryAdviceLineTwoAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineTwo = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineTwo", ErrNonAlphanumeric, fiba.Advice.LineTwo).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineThreeAlphaNumeric validates FIBeneficiaryAdvice LineThree is alphanumeric
func TestFIBeneficiaryAdviceLineThreeAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineThree = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineThree", ErrNonAlphanumeric, fiba.Advice.LineThree).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineFourAlphaNumeric validates FIBeneficiaryAdvice LineFour is alphanumeric
func TestFIBeneficiaryAdviceLineFourAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFour = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFour", ErrNonAlphanumeric, fiba.Advice.LineFour).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineFiveAlphaNumeric validates FIBeneficiaryAdvice LineFive is alphanumeric
func TestFIBeneficiaryAdviceLineFiveAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFive = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFive", ErrNonAlphanumeric, fiba.Advice.LineFive).Error(), err.Error())
}

// TestFIBeneficiaryAdviceLineSixAlphaNumeric validates FIBeneficiaryAdvice LineSix is alphanumeric
func TestFIBeneficiaryAdviceLineSixAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineSix = "®"

	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineSix", ErrNonAlphanumeric, fiba.Advice.LineSix).Error(), err.Error())
}

// TestParseFIBeneficiaryAdviceWrongLength parses a wrong FIBeneficiaryAdvice record length
func TestParseFIBeneficiaryAdviceWrongLength(t *testing.T) {
	var line = "{6410}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(200, len(r.line)).Error(), err.Error())
}

// TestParseFIBeneficiaryAdviceReaderParseError parses a wrong FIBeneficiaryAdvice reader parse error
func TestParseFIBeneficiaryAdviceReaderParseError(t *testing.T) {
	var line = "{6410}LTRLine ®ne                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestFIBeneficiaryAdviceTagError validates a FIBeneficiaryAdvice tag
func TestFIBeneficiaryAdviceTagError(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.tag = "{9999}"
	err := fiba.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, fiba.tag).Error(), err.Error())
}
