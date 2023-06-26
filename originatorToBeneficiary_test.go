package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockOriginatorToBeneficiary creates a OriginatorToBeneficiary
func mockOriginatorToBeneficiary() *OriginatorToBeneficiary {
	ob := NewOriginatorToBeneficiary()
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	return ob
}

// TestMockOriginatorToBeneficiary validates mockOriginatorToBeneficiary
func TestMockOriginatorToBeneficiary(t *testing.T) {
	ob := mockOriginatorToBeneficiary()

	require.NoError(t, ob.Validate(), "mockOriginatorToBeneficiary does not validate and will break other tests")
}

// TestOriginatorToBeneficiaryLineOneAlphaNumeric validates OriginatorToBeneficiary LineOne is alphanumeric
func TestOriginatorToBeneficiaryLineOneAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineOne = "®"

	err := ob.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, ob.LineOne).Error())
}

// TestOriginatorToBeneficiaryLineTwoAlphaNumeric validates OriginatorToBeneficiary LineTwo is alphanumeric
func TestOriginatorToBeneficiaryLineTwoAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineTwo = "®"

	err := ob.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, ob.LineTwo).Error())
}

// TestOriginatorToBeneficiaryLineThreeAlphaNumeric validates OriginatorToBeneficiary LineThree is alphanumeric
func TestOriginatorToBeneficiaryLineThreeAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineThree = "®"

	err := ob.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, ob.LineThree).Error())
}

// TestOriginatorToBeneficiaryLineFourAlphaNumeric validates OriginatorToBeneficiary LineFour is alphanumeric
func TestOriginatorToBeneficiaryLineFourAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineFour = "®"

	err := ob.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, ob.LineFour).Error())
}

// TestParseOriginatorToBeneficiaryWrongLength parses a wrong OriginatorToBeneficiary record length
func TestParseOriginatorToBeneficiaryWrongLength(t *testing.T) {
	var line = "{6000}LineOne                            LineTwo                            LineThree                          LineFour                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()

	require.EqualError(t, err, r.parseError(fieldError("LineOne", ErrRequireDelimiter)).Error())
}

// TestParseOriginatorToBeneficiaryReaderParseError parses a wrong OriginatorToBeneficiary reader parse error
func TestParseOriginatorToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6000}LineOne                            *®ineTwo                            *LineThree                          *LineFour                          *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()

	require.EqualError(t, err, r.parseError(fieldError("LineTwo", ErrNonAlphanumeric, "®ineTwo")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("LineTwo", ErrNonAlphanumeric, "®ineTwo")).Error())
}

// TestOriginatorToBeneficiaryTagError validates a OriginatorToBeneficiary tag
func TestOriginatorToBeneficiaryTagError(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.tag = "{9999}"

	require.EqualError(t, ob.Validate(), fieldError("tag", ErrValidTagForType, ob.tag).Error())
}

// TestStringOriginatorToBeneficiaryVariableLength parses using variable length
func TestStringOriginatorToBeneficiaryVariableLength(t *testing.T) {
	var line = "{6000}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()
	require.Nil(t, err)

	line = "{6000}                                                                                                                                            NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOriginatorToBeneficiary()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{6000}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOriginatorToBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{6000}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOriginatorToBeneficiary()
	require.Equal(t, err, nil)
}

// TestStringOriginatorToBeneficiaryOptions validates Format() formatted according to the FormatOptions
func TestStringOriginatorToBeneficiaryOptions(t *testing.T) {
	var line = "{6000}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.OriginatorToBeneficiary
	require.Equal(t, record.String(), "{6000}                                   *                                   *                                   *                                   *")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{6000}*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}

// TestStringOriginatorToBeneficiaryOptionsWithExtraLength validates the length of each line if it exceeds the limit of 35 chars per line
func TestStringOriginatorToBeneficiaryOptionsWithExtraLength(t *testing.T) {
	var line = "{6000}Lorem ipsum dolor sit amet, co WOODEN*adipiscing elit, sed do eiusmod 54F*022/FROM MYCOMPY INC. VIA MYBANKNIF*SUC INTERNAL BANK TO BANK TRANSFER,*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()
	require.Equal(t, err, nil)

	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineOne, "Lorem ipsum dolor sit amet, co WOOD")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineTwo, "adipiscing elit, sed do eiusmod 54F")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineThree, "022/FROM MYCOMPY INC. VIA MYBANKNIF")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineFour, "SUC INTERNAL BANK TO BANK TRANSFER,")
}

// TestStringOriginatorToBeneficiaryOptionsWithSmallerLength // TestStringOriginatorToBeneficiaryOptionsWithExtraLength validates the length of each line if it is less than the limit of 35 chars per line
func TestStringOriginatorToBeneficiaryOptionsWithSmallerLength(t *testing.T) {
	var line = "{6000}Lorem ipsum dolor sit amet, co W*adipiscing elit, sed do eiusmod 54F*022/FROM INC. VIA MYBANKNIF*SUC INTERNAL BANK TO TRANSFER,*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorToBeneficiary()
	require.Equal(t, err, nil)

	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineOne, "Lorem ipsum dolor sit amet, co W")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineTwo, "adipiscing elit, sed do eiusmod 54F")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineThree, "022/FROM INC. VIA MYBANKNIF")
	require.Equal(t, r.currentFEDWireMessage.OriginatorToBeneficiary.LineFour, "SUC INTERNAL BANK TO TRANSFER,")
}
