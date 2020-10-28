package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiaryFIAdvice creates a FIBeneficiaryFIAdvice
func mockFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := NewFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = AdviceCodeTelex
	fibfia.Advice.LineOne = "Line One"
	fibfia.Advice.LineTwo = "Line Two"
	fibfia.Advice.LineThree = "Line Three"
	fibfia.Advice.LineFour = "Line Four"
	fibfia.Advice.LineFive = "Line Five"
	fibfia.Advice.LineSix = "Line Six"
	return fibfia
}

// TestMockFIBeneficiaryFIAdvice validates mockFIBeneficiaryFIAdvice
func TestMockFIBeneficiaryFIAdvice(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()

	require.NoError(t, fibfia.Validate(), "mockFIBeneficiaryFIAdvice does not validate and will break other tests")
}

// TestFIBeneficiaryFIAdviceAdviceCodeValid validates FIBeneficiaryFIAdvice AdviceCode is alphanumeric
func TestFIBeneficiaryFIAdviceAdviceCodeValid(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = "Z"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AdviceCode", ErrAdviceCode, fibfia.Advice.AdviceCode).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineOneAlphaNumeric validates FIBeneficiaryFIAdvice LineOne is alphanumeric
func TestFIBeneficiaryFIAdviceLineOneAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineOne = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineOne", ErrNonAlphanumeric, fibfia.Advice.LineOne).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineTwoAlphaNumeric validates FIBeneficiaryFIAdvice LineTwo is alphanumeric
func TestFIBeneficiaryFIAdviceLineTwoAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineTwo = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineTwo", ErrNonAlphanumeric, fibfia.Advice.LineTwo).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineThreeAlphaNumeric validates FIBeneficiaryFIAdvice LineThree is alphanumeric
func TestFIBeneficiaryFIAdviceLineThreeAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineThree = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineThree", ErrNonAlphanumeric, fibfia.Advice.LineThree).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineFourAlphaNumeric validates FIBeneficiaryFIAdvice LineFour is alphanumeric
func TestFIBeneficiaryFIAdviceLineFourAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineFour = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFour", ErrNonAlphanumeric, fibfia.Advice.LineFour).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineFiveAlphaNumeric validates FIBeneficiaryFIAdvice LineFive is alphanumeric
func TestFIBeneficiaryFIAdviceLineFiveAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineFive = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFive", ErrNonAlphanumeric, fibfia.Advice.LineFive).Error(), err.Error())
}

// TestFIBeneficiaryFIAdviceLineSixAlphaNumeric validates FIBeneficiaryFIAdvice LineSix is alphanumeric
func TestFIBeneficiaryFIAdviceLineSixAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineSix = "®"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineSix", ErrNonAlphanumeric, fibfia.Advice.LineSix).Error(), err.Error())
}

// TestParseFIBeneficiaryFIAdviceWrongLength parses a wrong FIBeneficiaryFIAdvice record length
func TestParseFIBeneficiaryFIAdviceWrongLength(t *testing.T) {
	var line = "{6310}TLXLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryFIAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(200, len(r.line)).Error(), err.Error())
}

// TestParseFIBeneficiaryFIAdviceReaderParseError parses a wrong FIBeneficiaryFIAdvice reader parse error
func TestParseFIBeneficiaryFIAdviceReaderParseError(t *testing.T) {
	var line = "{6310}TLXLine ®ne                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryFIAdvice()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestFIBeneficiaryFIAdviceTagError validates a FIBeneficiaryFIAdvice tag
func TestFIBeneficiaryFIAdviceTagError(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.tag = "{9999}"

	err := fibfia.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, fibfia.tag).Error(), err.Error())
}
