package wire

import (
	"github.com/moov-io/base"
	"strings"
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

// TestParseFIAdditionalFIToFIWrongLength parses a wrong FIAdditionalFIToFI record length
func TestParseFIAdditionalFIToFIWrongLength(t *testing.T) {
	var line = "{6500}Line One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)
	err := r.parseFIAdditionalFIToFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(216, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIAdditionalFIToFIReaderParseError parses a wrong FIAdditionalFIToFI reader parse error
func TestParseFIAdditionalFIToFIReaderParseError(t *testing.T) {
	var line = "{6500}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fifi := mockFIAdditionalFIToFI()
	fwm.SetFIAdditionalFIToFI(fifi)
	err := r.parseFIAdditionalFIToFI()
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

// TestFIAdditionalFIToFITagError validates a FIAdditionalFIToFI tag
func TestFIAdditionalFIToFITagError(t *testing.T) {
	fifi := mockFIAdditionalFIToFI()
	fifi.tag = "{9999}"
	if err := fifi.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
