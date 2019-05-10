package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestParseFIReceiverFIWrongLength parses a wrong FIReceiverFI record length
func TestParseFIReceiverFIWrongLength(t *testing.T) {
	var line = "{6100}Line Six                                                                                                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	firfi := mockFIReceiverFI()
	fwm.SetFIReceiverFI(firfi)
	err := r.parseFIReceiverFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(201, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIReceiverFIReaderParseError parses a wrong FIReceiverFI reader parse error
func TestParseFIReceiverFIReaderParseError(t *testing.T) {
	var line = "{6100}Line Si®                                                                                                                                                                                           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	firfi := mockFIReceiverFI()
	fwm.SetFIReceiverFI(firfi)
	err := r.parseFIReceiverFI()
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
