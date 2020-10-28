package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// mockFIAdditionalFIToFI creates a FIAdditionalFIToFI
func mockFIAdditionalFIToFI() *FIAdditionalFIToFI {
	fifi := NewFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "Line One"
	fifi.AdditionalFIToFI.LineTwo = "Line Two"
	fifi.AdditionalFIToFI.LineThree = "Line Three"
	fifi.AdditionalFIToFI.LineFour = "Line Four"
	fifi.AdditionalFIToFI.LineFive = "Line Five"
	fifi.AdditionalFIToFI.LineSix = "Line Six"
	return fifi
}

// TestMockFIAdditionalFIToFI validates mockFIAdditionalFIToFI
func TestMockFIAdditionalFIToFI(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()

	require.NoError(t, fifi.Validate(), "mockFIAdditionalFIToFI does not validate and will break other tests")
}

// TestFIAdditionalFIToFILineOneAlphaNumeric validates FIAdditionalFIToFI LineOne is alphanumeric
func TestFIAdditionalFIToFILineOneAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineOne", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineOne).Error(), err.Error())
}

// TestFIAdditionalFIToFILineTwoAlphaNumeric validates FIAdditionalFIToFI LineTwo is alphanumeric
func TestFIAdditionalFIToFILineTwoAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineTwo = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineTwo", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineTwo).Error(), err.Error())
}

// TestFIAdditionalFIToFILineThreeAlphaNumeric validates FIAdditionalFIToFI LineThree is alphanumeric
func TestFIAdditionalFIToFILineThreeAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineThree = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineThree", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineThree).Error(), err.Error())
}

// TestFIAdditionalFIToFILineFourAlphaNumeric validates FIAdditionalFIToFI LineFour is alphanumeric
func TestFIAdditionalFIToFILineFourAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFour = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFour", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineFour).Error(), err.Error())
}

// TestFIAdditionalFIToFILineFiveAlphaNumeric validates FIAdditionalFIToFI LineFive is alphanumeric
func TestFIAdditionalFIToFILineFiveAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFive = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFive", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineFive).Error(), err.Error())
}

// TestFIAdditionalFIToFILineSixAlphaNumeric validates FIAdditionalFIToFI LineSix is alphanumeric
func TestFIAdditionalFIToFILineSixAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineSix = "®"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineSix", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineSix).Error(), err.Error())
}

// TestParseFIAdditionalFIToFIWrongLength parses a wrong FIAdditionalFIToFI record length
func TestParseFIAdditionalFIToFIWrongLength(t *testing.T) {
	var line = "{6500}Line One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(216, len(r.line)).Error())
}

// TestParseFIAdditionalFIToFIReaderParseError parses a wrong FIAdditionalFIToFI reader parse error
func TestParseFIAdditionalFIToFIReaderParseError(t *testing.T) {
	var line = "{6500}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFIAdditionalFIToFITagError validates a FIAdditionalFIToFI tag
func TestFIAdditionalFIToFITagError(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.tag = "{9999}"

	err := fifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, fifi.tag).Error(), err.Error())
}
