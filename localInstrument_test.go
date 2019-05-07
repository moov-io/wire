package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockLocalInstrument creates a LocalInstrument
func mockLocalInstrument() *LocalInstrument {
	li := NewLocalInstrument()
	li.LocalInstrumentCode = ANSIX12format
	li.ProprietaryCode = ""
	return li
}

// TestMockLocalInstrument validates mockLocalInstrument
func TestMockLocalInstrument(t *testing.T) {
	li := mockLocalInstrument()
	if err := li.Validate(); err != nil {
		t.Error("mockLocalInstrument does not validate and will break other tests")
	}
}

// TestLocalInstrumentCodeValid validates LocalInstrumentCode
func TestLocalInstrumentCodeValid(t *testing.T) {
	li := mockLocalInstrument()
	li.LocalInstrumentCode = "Chestnut"
	if err := li.Validate(); err != nil {
		if !base.Match(err, ErrLocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryCodeValid validates ProprietaryCode
func TestProprietaryCodeValid(t *testing.T) {
	li := mockLocalInstrument()
	li.ProprietaryCode = "Proprietary"
	if err := li.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryCodeAlphaNumeric validates ProprietaryCode is alphanumeric
func TestProprietaryCodeAlphaNumeric(t *testing.T) {
	li := mockLocalInstrument()
	li.LocalInstrumentCode = ProprietaryLocalInstrumentCode
	li.ProprietaryCode = "®"
	if err := li.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseLocalInstrumentWrongLength parses a wrong LocalInstrumente record length
func TestParseLocalInstrumentWrongLength(t *testing.T) {
	var line = "{3610}ANSI                                 "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	err := r.parseLocalInstrument()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(45, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseLocalInstrumentReaderParseError parses a wrong LocalInstrumente reader parse error
func TestParseLocalInstrumentReaderParseError(t *testing.T) {
	var line = "{3610}ABCD                                   "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	err := r.parseLocalInstrument()
	if err != nil {
		if !base.Match(err, ErrLocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrLocalInstrumentCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}