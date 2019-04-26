package wire

import (
	"github.com/moov-io/base"
	"testing"
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
	if err := ob.Validate(); err != nil {
		t.Error("mockOriginatorToBeneficiary does not validate and will break other tests")
	}
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
