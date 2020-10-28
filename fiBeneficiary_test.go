package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIBeneficiary creates a FIBeneficiary
func mockFIBeneficiary() *FIBeneficiary {
	fib := NewFIBeneficiary()
	fib.FIToFI.LineOne = "Line One"
	fib.FIToFI.LineTwo = "Line Two"
	fib.FIToFI.LineThree = "Line Three"
	fib.FIToFI.LineFour = "Line Four"
	fib.FIToFI.LineFive = "Line Five"
	fib.FIToFI.LineSix = "Line Six"
	return fib
}

// TestMockFIBeneficiary validates mockFIBeneficiary
func TestMockFIBeneficiary(t *testing.T) {
	fib := mockFIBeneficiary()

	require.NoError(t, fib.Validate(), "mockFIBeneficiary does not validate and will break other tests")
}

// TestFIBeneficiaryLineOneAlphaNumeric validates FIBeneficiary LineOne is alphanumeric
func TestFIBeneficiaryLineOneAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineOne = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineOne", ErrNonAlphanumeric, fib.FIToFI.LineOne).Error(), err.Error())
}

// TestFIBeneficiaryLineTwoAlphaNumeric validates FIBeneficiary LineTwo is alphanumeric
func TestFIBeneficiaryLineTwoAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineTwo = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineTwo", ErrNonAlphanumeric, fib.FIToFI.LineTwo).Error(), err.Error())
}

// TestFIBeneficiaryLineThreeAlphaNumeric validates FIBeneficiary LineThree is alphanumeric
func TestFIBeneficiaryLineThreeAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineThree = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineThree", ErrNonAlphanumeric, fib.FIToFI.LineThree).Error(), err.Error())
}

// TestFIBeneficiaryLineFourAlphaNumeric validates FIBeneficiary LineFour is alphanumeric
func TestFIBeneficiaryLineFourAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFour = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFour", ErrNonAlphanumeric, fib.FIToFI.LineFour).Error(), err.Error())
}

// TestFIBeneficiaryLineFiveAlphaNumeric validates FIBeneficiary LineFive is alphanumeric
func TestFIBeneficiaryLineFiveAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFive = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineFive", ErrNonAlphanumeric, fib.FIToFI.LineFive).Error(), err.Error())
}

// TestFIBeneficiaryLineSixAlphaNumeric validates FIBeneficiary LineSix is alphanumeric
func TestFIBeneficiaryLineSixAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineSix = "®"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("LineSix", ErrNonAlphanumeric, fib.FIToFI.LineSix).Error(), err.Error())
}

// TestParseFIBeneficiaryWrongLength parses a wrong FIBeneficiary record length
func TestParseFIBeneficiaryWrongLength(t *testing.T) {
	var line = "{6400}Line Six                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(201, len(r.line)).Error())
}

// TestParseFIBeneficiaryReaderParseError parses a wrong FIBeneficiary reader parse error
func TestParseFIBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6400}Line Si®                                                                                                                                                                                           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIBeneficiary()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestFIBeneficiaryTagError validates a FIBeneficiary tag
func TestFIBeneficiaryTagError(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.tag = "{9999}"

	err := fib.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, fib.tag).Error(), err.Error())
}
