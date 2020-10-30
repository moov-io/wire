package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// RemittanceFreeText creates a RemittanceFreeText
func mockRemittanceFreeText() *RemittanceFreeText {
	rft := NewRemittanceFreeText()
	rft.LineOne = "Remittance Free Text Line One"
	rft.LineTwo = "Remittance Free Text Line Two"
	rft.LineThree = "Remittance Free Text Line Three"
	return rft
}

// TestMockRemittanceFreeText validates mockRemittanceFreeText
func TestMockRemittanceFreeText(t *testing.T) {
	rft := mockRemittanceFreeText()

	require.NoError(t, rft.Validate(), "mockRemittanceFreeText does not validate and will break other tests")
}

// TestRemittanceFreeTextLineOneAlphaNumeric validates RemittanceFreeText LineOne is alphanumeric
func TestRemittanceFreeTextLineOneAlphaNumeric(t *testing.T) {
	rft := mockRemittanceFreeText()
	rft.LineOne = "®"
	if err := rft.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceFreeTextLineTwoAlphaNumeric validates RemittanceFreeText LineTwo is alphanumeric
func TestRemittanceFreeTextLineTwoAlphaNumeric(t *testing.T) {
	rft := mockRemittanceFreeText()
	rft.LineTwo = "®"
	if err := rft.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceFreeTextLineThreeAlphaNumeric validates RemittanceFreeText LineThree is alphanumeric
func TestRemittanceFreeTextLineThreeAlphaNumeric(t *testing.T) {
	rft := mockRemittanceFreeText()
	rft.LineThree = "®"
	if err := rft.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceFreeTextWrongLength parses a wrong RemittanceFreeText record length
func TestParseRemittanceFreeTextWrongLength(t *testing.T) {
	var line = "{8750}Re®ittance Free Text Line One                                                                                                               Remittance Free Text Line Two                                                                                                               Remittance Free Text Line Three                                                                                                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)
	err := r.parseRemittanceFreeText()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(426, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceFreeTextReaderParseError parses a wrong RemittanceFreeText reader parse error
func TestParseRemittanceFreeTextReaderParseError(t *testing.T) {
	var line = "{8750}Re®ittance Free Text Line One                                                                                                               Remittance Free Text Line Two                                                                                                               Remittance Free Text Line Three                                                                                                             "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rft := mockRemittanceFreeText()
	fwm.SetRemittanceFreeText(rft)
	err := r.parseRemittanceFreeText()
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

// TestRemittanceFreeTextTagError validates a RemittanceFreeText tag
func TestRemittanceFreeTextTagError(t *testing.T) {
	rft := mockRemittanceFreeText()
	rft.tag = "{9999}"

	require.EqualError(t, rft.Validate(), fieldError("tag", ErrValidTagForType, rft.tag).Error())
}
