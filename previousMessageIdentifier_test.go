package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockPreviousMessageIdentifier creates a PreviousMessageIdentifier
func mockPreviousMessageIdentifier() *PreviousMessageIdentifier {
	pmi := NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	return pmi
}

// TestMockPreviousMessageIdentifier validates mockPreviousMessageIdentifier
func TestMockPreviousMessageIdentifier(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	if err := pmi.Validate(); err != nil {
		t.Error("mockPreviousMessageIdentifier does not validate and will break other tests")
	}
}

// TestPreviousMessageIdentifierAlphaNumeric validates PreviousMessageIdentifier is alphanumeric
func TestPreviousMessageIdentifierAlphaNumeric(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "®"
	if err := pmi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePreviousMessageIdentifierWrongLength parses a wrong PreviousMessageIdentifier record length
func TestParsePreviousMessageIdentifierWrongLength(t *testing.T) {
	var line = "{3500}Previous"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)
	err := r.parsePreviousMessageIdentifier()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(28, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePreviousMessageIdentifierReaderParseError parses a wrong PreviousMessageIdentifier reader parse error
func TestParsePreviousMessageIdentifierReaderParseError(t *testing.T) {
	var line = "{3500}Previous®Message Ident"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	pmi := mockPreviousMessageIdentifier()
	fwm.SetPreviousMessageIdentifier(pmi)
	err := r.parsePreviousMessageIdentifier()
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

// TestPreviousMessageIdentifierTagError validates a PreviousMessageIdentifier tag
func TestPreviousMessageIdentifierTagError(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	pmi.tag = "{9999}"
	if err := pmi.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
