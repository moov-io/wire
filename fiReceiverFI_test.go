package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockFIReceiverFI creates a FIReceiverFI
func mockFIReceiverFI() *FIReceiverFI {
	firfi := NewFIReceiverFI()
	firfi.FIToFI.LineOne = "Line One"
	firfi.FIToFI.LineOne = "Line Two"
	firfi.FIToFI.LineOne = "Line Three"
	firfi.FIToFI.LineOne = "Line Four"
	firfi.FIToFI.LineOne = "Line Five"
	firfi.FIToFI.LineOne = "Line Six"
	return firfi
}

// TestMockFIReceiverFI validates mockFIReceiverFI
func TestMockFIReceiverFI(t *testing.T) {
	firfi := mockFIReceiverFI()
	if err := firfi.Validate(); err != nil {
		t.Error("mockFIReceiverFI does not validate and will break other tests")
	}
}

// TestFIReceiverFILineOneAlphaNumeric validates FIReceiverFI LineOne is alphanumeric
func TestFIReceiverFILineOneAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineOne = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIReceiverFILineTwoAlphaNumeric validates FIReceiverFI LineTwo is alphanumeric
func TestFIReceiverFILineTwoAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineTwo = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIReceiverFILineThreeAlphaNumeric validates FIReceiverFI LineThree is alphanumeric
func TestFIReceiverFILineThreeAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineThree = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIReceiverFILineFourAlphaNumeric validates FIReceiverFI LineFour is alphanumeric
func TestFIReceiverFILineFourAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineFour = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIReceiverFILineFiveAlphaNumeric validates FIReceiverFI LineFive is alphanumeric
func TestFIReceiverFILineFiveAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineFive = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIReceiverFILineSixAlphaNumeric validates FIReceiverFI LineSix is alphanumeric
func TestFIReceiverFILineSixAlphaNumeric(t *testing.T) {
	firfi := mockFIReceiverFI()
	firfi.FIToFI.LineSix = "®"
	if err := firfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}