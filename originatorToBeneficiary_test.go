package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
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
	if err := ob.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorToBeneficiaryLineTwoAlphaNumeric validates OriginatorToBeneficiary LineTwo is alphanumeric
func TestOriginatorToBeneficiaryLineTwoAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineTwo = "®"
	if err := ob.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorToBeneficiaryLineThreeAlphaNumeric validates OriginatorToBeneficiary LineThree is alphanumeric
func TestOriginatorToBeneficiaryLineThreeAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineThree = "®"
	if err := ob.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorToBeneficiaryLineFourAlphaNumeric validates OriginatorToBeneficiary LineFour is alphanumeric
func TestOriginatorToBeneficiaryLineFourAlphaNumeric(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.LineFour = "®"
	if err := ob.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorToBeneficiaryWrongLength parses a wrong OriginatorToBeneficiary record length
func TestParseOriginatorToBeneficiaryWrongLength(t *testing.T) {
	var line = "{6000}LineOne                            LineTwo                            LineThree                          LineFour                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	err := r.parseOriginatorToBeneficiary()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(146, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorToBeneficiaryReaderParseError parses a wrong OriginatorToBeneficiary reader parse error
func TestParseOriginatorToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6000}LineOne                            ®ineTwo                            LineThree                          LineFour                           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ob := mockOriginatorToBeneficiary()
	fwm.SetOriginatorToBeneficiary(ob)
	err := r.parseOriginatorToBeneficiary()
	if err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorToBeneficiaryTagError validates a OriginatorToBeneficiary tag
func TestOriginatorToBeneficiaryTagError(t *testing.T) {
	ob := mockOriginatorToBeneficiary()
	ob.tag = "{9999}"

	require.EqualError(t, ob.Validate(), fieldError("tag", ErrValidTagForType, ob.tag).Error())
}
