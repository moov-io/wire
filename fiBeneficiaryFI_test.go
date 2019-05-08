package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockFIBeneficiaryFI creates a FIBeneficiaryFI
func mockFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "Line One"
	fibfi.FIToFI.LineTwo = "Line Two"
	fibfi.FIToFI.LineThree = "Line Three"
	fibfi.FIToFI.LineFour = "Line Four"
	fibfi.FIToFI.LineFive = "Line Five"
	fibfi.FIToFI.LineSix = "Line Six"
	return fibfi
}

// TestMockFIBeneficiaryFI validates mockFIBeneficiaryFI
func TestMockFIBeneficiaryFI(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	if err := fibfi.Validate(); err != nil {
		t.Error("mockFIBeneficiaryFI does not validate and will break other tests")
	}
}

// TestFIBeneficiaryFILineOneAlphaNumeric validates FIBeneficiaryFI LineOne is alphanumeric
func TestFIBeneficiaryFILineOneAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFILineTwoAlphaNumeric validates FIBeneficiaryFI LineTwo is alphanumeric
func TestFIBeneficiaryFILineTwoAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineTwo = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFILineThreeAlphaNumeric validates FIBeneficiaryFI LineThree is alphanumeric
func TestFIBeneficiaryFILineThreeAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineThree = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFILineFourAlphaNumeric validates FIBeneficiaryFI LineFour is alphanumeric
func TestFIBeneficiaryFILineFourAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineFour = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFILineFiveAlphaNumeric validates FIBeneficiaryFI LineFive is alphanumeric
func TestFIBeneficiaryFILineFiveAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineFive = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFILineSixAlphaNumeric validates FIBeneficiaryFI LineSix is alphanumeric
func TestFIBeneficiaryFILineSixAlphaNumeric(t *testing.T) {
	fibfi := mockFIBeneficiaryFI()
	fibfi.FIToFI.LineSix = "®"
	if err := fibfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIBeneficiaryFIWrongLength parses a wrong FIBeneficiaryFI record length
func TestParseFIBeneficiaryFIWrongLength(t *testing.T) {
	var line = "{6100}Line Six                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	err := r.parseFIBeneficiaryFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(201, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIBeneficiaryFIReaderParseError parses a wrong FIBeneficiaryFI reader parse error
func TestParseFIBeneficiaryFIReaderParseError(t *testing.T) {
	var line = "{6100}Line Si®                                                                                                                                                                                           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fibfi := mockFIBeneficiaryFI()
	fwm.SetFIBeneficiaryFI(fibfi)
	err := r.parseFIBeneficiaryFI()
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
