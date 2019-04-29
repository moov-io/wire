package wire

import (
	"github.com/moov-io/base"
	"testing"
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
	if err := fib.Validate(); err != nil {
		t.Error("mockFIBeneficiary does not validate and will break other tests")
	}
}

// TestFIBeneficiaryLineOneAlphaNumeric validates FIBeneficiary LineOne is alphanumeric
func TestFIBeneficiaryLineOneAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineOne = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryLineTwoAlphaNumeric validates FIBeneficiary LineTwo is alphanumeric
func TestFIBeneficiaryLineTwoAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineTwo = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryLineThreeAlphaNumeric validates FIBeneficiary LineThree is alphanumeric
func TestFIBeneficiaryLineThreeAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineThree = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryLineFourAlphaNumeric validates FIBeneficiary LineFour is alphanumeric
func TestFIBeneficiaryLineFourAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFour = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryLineFiveAlphaNumeric validates FIBeneficiary LineFive is alphanumeric
func TestFIBeneficiaryLineFiveAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineFive = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryLineSixAlphaNumeric validates FIBeneficiary LineSix is alphanumeric
func TestFIBeneficiaryLineSixAlphaNumeric(t *testing.T) {
	fib := mockFIBeneficiary()
	fib.FIToFI.LineSix = "®"
	if err := fib.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
