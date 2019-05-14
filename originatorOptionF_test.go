package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockOriginatorOptionF creates a OriginatorOptionF
func mockOriginatorOptionF() *OriginatorOptionF {
	oof := NewOriginatorOptionF()
	oof.PartyIdentifier = "TXID/123-45-6789"
	oof.Name = "Name"
	oof.LineOne = "LineOne"
	oof.LineTwo = "LineTwo"
	oof.LineThree = "LineThree"
	return oof
}

// TestMockOriginatorOptionF validates mockOriginatorOptionF
func TestMockOriginatorOptionF(t *testing.T) {
	oof := mockOriginatorOptionF()
	if err := oof.Validate(); err != nil {
		t.Error("mockOriginatorOptionF does not validate and will break other tests")
	}
}

// TestOriginatorOptionFPartyIdentifierAlphaNumeric validates OriginatorOptionF PartyIdentifier is alphanumeric
func TestOriginatorOptionFPartyIdentifierAlphaNumeric(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFNameAlphaNumeric validates OriginatorOptionF Name is alphanumeric
func TestOriginatorOptionFNameAlphaNumeric(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.Name = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineOneAlphaNumeric validates OriginatorOptionF LineOne is alphanumeric
func TestOriginatorOptionFLineOneAlphaNumeric(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineTwoAlphaNumeric validates OriginatorOptionF LineTwo is alphanumeric
func TestOriginatorOptionFLineTwoAlphaNumeric(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineTwo = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineThreeAlphaNumeric validates OriginatorOptionF LineThree is alphanumeric
func TestOriginatorOptionFLineThreeAlphaNumeric(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineThree = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorOptionFWrongLength parses a wrong OriginatorOptionF record length
func TestParseOriginatorOptionFWrongLength(t *testing.T) {
	var line = "{5010}TXID/123-45-6789                   Name                               LineOne                            LineTwo                            LineThree                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oof := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(oof)
	err := r.parseOriginatorOptionF()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorOptionFReaderParseError parses a wrong OriginatorOptionF reader parse error
func TestParseOriginatorOptionFReaderParseError(t *testing.T) {
	var line = "{5010}TXID/123-45-6789                   ®ame                               LineOne                            LineTwo                            LineThree                          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oof := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(oof)
	err := r.parseOriginatorOptionF()
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

// TestOriginatorOptionFTagError validates a OriginatorOptionF tag
func TestOriginatorOptionFTagError(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.tag = "{9999}"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
