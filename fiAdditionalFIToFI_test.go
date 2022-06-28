package wire

import (
	"strings"
	"testing"

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

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineOne).Error())
}

// TestFIAdditionalFIToFILineTwoAlphaNumeric validates FIAdditionalFIToFI LineTwo is alphanumeric
func TestFIAdditionalFIToFILineTwoAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineTwo = "®"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineTwo).Error())
}

// TestFIAdditionalFIToFILineThreeAlphaNumeric validates FIAdditionalFIToFI LineThree is alphanumeric
func TestFIAdditionalFIToFILineThreeAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineThree = "®"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineThree).Error())
}

// TestFIAdditionalFIToFILineFourAlphaNumeric validates FIAdditionalFIToFI LineFour is alphanumeric
func TestFIAdditionalFIToFILineFourAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFour = "®"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineFour).Error())
}

// TestFIAdditionalFIToFILineFiveAlphaNumeric validates FIAdditionalFIToFI LineFive is alphanumeric
func TestFIAdditionalFIToFILineFiveAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFive = "®"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineFive).Error())
}

// TestFIAdditionalFIToFILineSixAlphaNumeric validates FIAdditionalFIToFI LineSix is alphanumeric
func TestFIAdditionalFIToFILineSixAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineSix = "®"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, fifi.AdditionalFIToFI.LineSix).Error())
}

// TestParseFIAdditionalFIToFIWrongLength parses a wrong FIAdditionalFIToFI record length
func TestParseFIAdditionalFIToFIWrongLength(t *testing.T) {
	var line = "{6500}Line One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()
	require.EqualError(t, err, r.parseError(fieldError("LineSix", ErrValidLength)).Error())
}

// TestParseFIAdditionalFIToFIReaderParseError parses a wrong FIAdditionalFIToFI reader parse error
func TestParseFIAdditionalFIToFIReaderParseError(t *testing.T) {
	var line = "{6500}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                          "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()

	expected := r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error()
	require.EqualError(t, err, expected)
}

// TestFIAdditionalFIToFITagError validates a FIAdditionalFIToFI tag
func TestFIAdditionalFIToFITagError(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.tag = "{9999}"

	err := fifi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, fifi.tag).Error())
}

// TestStringFIAdditionalFIToFIVariableLength parses using variable length
func TestStringFIAdditionalFIToFIVariableLength(t *testing.T) {
	var line = "{6500}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()
	require.Nil(t, err)

	line = "{6500}                                                                                                                                                                                                                  NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIAdditionalFIToFI()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{6500}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIAdditionalFIToFI()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{6500}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIAdditionalFIToFI()
	require.Equal(t, err, nil)
}

// TestStringFIAdditionalFIToFIOptions validates Format() formatted according to the FormatOptions
func TestStringFIAdditionalFIToFIOptions(t *testing.T) {
	var line = "{6500}*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIAdditionalFIToFI()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.FIAdditionalFIToFI
	require.Equal(t, record.String(), "{6500}                                                                                                                                                                                                                  ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{6500}*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
