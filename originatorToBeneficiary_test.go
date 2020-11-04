package wire

import (
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

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(146, len(r.line))).Error())
}

// TestParseOriginatorToBeneficiaryReaderParseError parses a wrong OriginatorToBeneficiary reader parse error
func TestParseOriginatorToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6000}LineOne                            ®ineTwo                            LineThree                          LineFour                           "
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
