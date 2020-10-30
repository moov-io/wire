package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIReceiverFI creates a FIReceiverFI
func mockFIReceiverFI() *FIReceiverFI {
	firfi := NewFIReceiverFI()
	firfi.FIToFI.LineOne = "Line One"
	firfi.FIToFI.LineOne = "Line Two"
	firfi.FIToFI.LineOne = "Line Three"
	firfi.FIToFI.LineOne = "Line Four"
	firfi.FIToFI.LineOne = "Line Five"
	firfi.FIToFI.LineOne = "Line Six"
	return firfi
}

// TestMockFIReceiverFI validates mockFIReceiverFI
func TestMockFIReceiverFI(t *testing.T) {
	firfi := mockFIReceiverFI()

	require.NoError(t, firfi.Validate(), "mockFIReceiverFI does not validate and will break other tests")
}

// TestFIReceiverFILineOneAlphaNumeric validates FIReceiverFI LineOne is alphanumeric
func TestFIReceiverFILineOneAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineOne = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, firfi.FIToFI.LineOne).Error())
}

// TestFIReceiverFILineTwoAlphaNumeric validates FIReceiverFI LineTwo is alphanumeric
func TestFIReceiverFILineTwoAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineTwo = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, firfi.FIToFI.LineTwo).Error())
}

// TestFIReceiverFILineThreeAlphaNumeric validates FIReceiverFI LineThree is alphanumeric
func TestFIReceiverFILineThreeAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineThree = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, firfi.FIToFI.LineThree).Error())
}

// TestFIReceiverFILineFourAlphaNumeric validates FIReceiverFI LineFour is alphanumeric
func TestFIReceiverFILineFourAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineFour = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, firfi.FIToFI.LineFour).Error())
}

// TestFIReceiverFILineFiveAlphaNumeric validates FIReceiverFI LineFive is alphanumeric
func TestFIReceiverFILineFiveAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineFive = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, firfi.FIToFI.LineFive).Error())
}

// TestFIReceiverFILineSixAlphaNumeric validates FIReceiverFI LineSix is alphanumeric
func TestFIReceiverFILineSixAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineSix = "®"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, firfi.FIToFI.LineSix).Error())
}

// TestParseFIReceiverFIWrongLength parses a wrong FIReceiverFI record length
func TestParseFIReceiverFIWrongLength(t *testing.T) {
	var line = "{6100}Line Six                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIReceiverFI()
	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(201, len(r.line))).Error())
}

// TestParseFIReceiverFIReaderParseError parses a wrong FIReceiverFI reader parse error
func TestParseFIReceiverFIReaderParseError(t *testing.T) {
	var line = "{6100}Line Si®                                                                                                                                                                                           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIReceiverFI()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line Si®")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line Si®")).Error()
	require.EqualError(t, err, expected)
}

// TestFIReceiverFITagError validates a FIReceiverFI tag
func TestFIReceiverFITagError(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.tag = "{9999}"

	err := firfi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, firfi.tag).Error())
}
