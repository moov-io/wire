package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockFIIntermediaryFI creates a FIIntermediaryFI
func mockFIIntermediaryFI() *FIIntermediaryFI {
	fiifi := NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	return fiifi
}

// TestMockFIIntermediaryFI validates mockFIIntermediaryFI
func TestMockFIIntermediaryFI(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	if err := fiifi.Validate(); err != nil {
		t.Error("mockFIIntermediaryFI does not validate and will break other tests")
	}
}

// TestFIIntermediaryFILineOneAlphaNumeric validates FIIntermediaryFI LineOne is alphanumeric
func TestFIIntermediaryFILineOneAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFILineTwoAlphaNumeric validates FIIntermediaryFI LineTwo is alphanumeric
func TestFIIntermediaryFILineTwoAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineTwo = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFILineThreeAlphaNumeric validates FIIntermediaryFI LineThree is alphanumeric
func TestFIIntermediaryFILineThreeAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineThree = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFILineFourAlphaNumeric validates FIIntermediaryFI LineFour is alphanumeric
func TestFIIntermediaryFILineFourAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineFour = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFILineFiveAlphaNumeric validates FIIntermediaryFI LineFive is alphanumeric
func TestFIIntermediaryFILineFiveAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineFive = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFILineSixAlphaNumeric validates FIIntermediaryFI LineSix is alphanumeric
func TestFIIntermediaryFILineSixAlphaNumeric(t *testing.T) {
	fiifi := mockFIIntermediaryFI()
	fiifi.FIToFI.LineSix = "®"
	if err := fiifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}