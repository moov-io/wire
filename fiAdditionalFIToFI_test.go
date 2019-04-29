package wire

import (
	"github.com/moov-io/base"
	"testing"
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
	if err := fifi.Validate(); err != nil {
		t.Error("mockFIAdditionalFIToFI does not validate and will break other tests")
	}
}

// TestFIAdditionalFIToFILineOneAlphaNumeric validates FIAdditionalFIToFI LineOne is alphanumeric
func TestFIAdditionalFIToFILineOneAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIAdditionalFIToFILineTwoAlphaNumeric validates FIAdditionalFIToFI LineTwo is alphanumeric
func TestFIAdditionalFIToFILineTwoAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineTwo = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIAdditionalFIToFILineThreeAlphaNumeric validates FIAdditionalFIToFI LineThree is alphanumeric
func TestFIAdditionalFIToFILineThreeAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineThree = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIAdditionalFIToFILineFourAlphaNumeric validates FIAdditionalFIToFI LineFour is alphanumeric
func TestFIAdditionalFIToFILineFourAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFour = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIAdditionalFIToFILineFiveAlphaNumeric validates FIAdditionalFIToFI LineFive is alphanumeric
func TestFIAdditionalFIToFILineFiveAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineFive = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIAdditionalFIToFILineSixAlphaNumeric validates FIAdditionalFIToFI LineSix is alphanumeric
func TestFIAdditionalFIToFILineSixAlphaNumeric(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineSix = "®"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
