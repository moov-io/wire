package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiaryFI creates a FIBeneficiaryFI
func mockFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "Line One"
	fibfi.FIToFI.LineTwo = "Line Two"
	fibfi.FIToFI.LineThree = "Line Three"
	fibfi.FIToFI.LineFour = "Line Four"
	fibfi.FIToFI.LineFive = "Line Five"
	fibfi.FIToFI.LineSix = "Line Six"
	return fibfi
}

// TestMockFIBeneficiaryFI validates mockFIBeneficiaryFI
func TestMockFIBeneficiaryFI(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()

	require.NoError(t, fibfi.Validate(), "mockFIBeneficiaryFI does not validate and will break other tests")
}

// TestFIBeneficiaryFILineOneAlphaNumeric validates FIBeneficiaryFI LineOne is alphanumeric
func TestFIBeneficiaryFILineOneAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fibfi.FIToFI.LineOne).Error())
}

// TestFIBeneficiaryFILineTwoAlphaNumeric validates FIBeneficiaryFI LineTwo is alphanumeric
func TestFIBeneficiaryFILineTwoAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineTwo = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fibfi.FIToFI.LineTwo).Error())
}

// TestFIBeneficiaryFILineThreeAlphaNumeric validates FIBeneficiaryFI LineThree is alphanumeric
func TestFIBeneficiaryFILineThreeAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineThree = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fibfi.FIToFI.LineThree).Error())
}

// TestFIBeneficiaryFILineFourAlphaNumeric validates FIBeneficiaryFI LineFour is alphanumeric
func TestFIBeneficiaryFILineFourAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineFour = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fibfi.FIToFI.LineFour).Error())
}

// TestFIBeneficiaryFILineFiveAlphaNumeric validates FIBeneficiaryFI LineFive is alphanumeric
func TestFIBeneficiaryFILineFiveAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineFive = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fibfi.FIToFI.LineFive).Error())
}

// TestFIBeneficiaryFILineSixAlphaNumeric validates FIBeneficiaryFI LineSix is alphanumeric
func TestFIBeneficiaryFILineSixAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineSix = "®"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fibfi.FIToFI.LineSix).Error())
}

// TestParseFIBeneficiaryFIWrongLength parses a wrong FIBeneficiaryFI record length
func TestParseFIBeneficiaryFIWrongLength(t *testing.T) {
	var line = "{6300}Line One                      Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryFI()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(201, len(r.line))).Error())
}

// TestParseFIBeneficiaryFIReaderParseError parses a wrong FIBeneficiaryFI reader parse error
func TestParseFIBeneficiaryFIReaderParseError(t *testing.T) {
	var line = "{6300}Line ®ne                      Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiaryFI()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ne")).Error()
	require.EqualError(t, err, expected)
}

// TestFIBeneficiaryFITagError validates a FIBeneficiaryFI tag
func TestFIBeneficiaryFITagError(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.tag = "{9999}"

	err := fibfi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fibfi.tag).Error())
}
