package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIIntermediaryFI creates a FIIntermediaryFI
func mockFIIntermediaryFI() *FIIntermediaryFI {
	fiifi := NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	return fiifi
}

// TestMockFIIntermediaryFI validates mockFIIntermediaryFI
func TestMockFIIntermediaryFI(t *testing.T) {
	fiifi := mockFIIntermediaryFI()

	require.NoError(t, fiifi.Validate(), "mockFIIntermediaryFI does not validate and will break other tests")
}

// TestFIIntermediaryFILineOneAlphaNumeric validates FIIntermediaryFI LineOne is alphanumeric
func TestFIIntermediaryFILineOneAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fiifi.FIToFI.LineOne).Error())
}

// TestFIIntermediaryFILineTwoAlphaNumeric validates FIIntermediaryFI LineTwo is alphanumeric
func TestFIIntermediaryFILineTwoAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineTwo = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fiifi.FIToFI.LineTwo).Error())
}

// TestFIIntermediaryFILineThreeAlphaNumeric validates FIIntermediaryFI LineThree is alphanumeric
func TestFIIntermediaryFILineThreeAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineThree = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fiifi.FIToFI.LineThree).Error())
}

// TestFIIntermediaryFILineFourAlphaNumeric validates FIIntermediaryFI LineFour is alphanumeric
func TestFIIntermediaryFILineFourAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineFour = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fiifi.FIToFI.LineFour).Error())
}

// TestFIIntermediaryFILineFiveAlphaNumeric validates FIIntermediaryFI LineFive is alphanumeric
func TestFIIntermediaryFILineFiveAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineFive = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fiifi.FIToFI.LineFive).Error())
}

// TestFIIntermediaryFILineSixAlphaNumeric validates FIIntermediaryFI LineSix is alphanumeric
func TestFIIntermediaryFILineSixAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineSix = "®"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fiifi.FIToFI.LineSix).Error())
}

// TestParseFIIntermediaryFIWrongLength parses a wrong FIIntermediaryFI record length
func TestParseFIIntermediaryFIWrongLength(t *testing.T) {
	var line = "{6200}Line ®ix                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIIntermediaryFI()
	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(201, len(r.line))).Error())
}

// TestParseFIIntermediaryFIReaderParseError parses a wrong FIIntermediaryFI reader parse error
func TestParseFIIntermediaryFIReaderParseError(t *testing.T) {
	var line = "{6200}Line ®ix                                                                                                                                                                                           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIIntermediaryFI()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ix")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "Line ®ix")).Error()
	require.EqualError(t, err, expected)
}

// TestFIIntermediaryFITagError validates a FIIntermediaryFI tag
func TestFIIntermediaryFITagError(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.tag = "{9999}"

	err := fiifi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fiifi.tag).Error())
}
