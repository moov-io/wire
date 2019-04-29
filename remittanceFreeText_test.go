package wire

import (
	"github.com/moov-io/base"
	"testing"
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
	if err := rft.Validate(); err != nil {
		t.Error("mockRemittanceFreeText does not validate and will break other tests")
	}
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
