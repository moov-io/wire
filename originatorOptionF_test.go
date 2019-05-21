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
	oof.Name = "1/Name"
	oof.LineOne = "1/1234"
	oof.LineTwo = "2/1000 Colonial Farm Rd"
	oof.LineThree = "5/Pottstown"
	return oof
}

// TestMockOriginatorOptionF validates mockOriginatorOptionF
func TestMockOriginatorOptionF(t *testing.T) {
	oof := mockOriginatorOptionF()
	if err := oof.Validate(); err != nil {
		t.Error("mockOriginatorOptionF does not validate and will break other tests")
	}
}

// TestOriginatorOptionFPartyIdentifier validates OriginatorOptionF PartyIdentifier is valid
func TestOriginatorOptionFPartyIdentifier(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "®®sdaasd"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierNull validates OriginatorOptionF PartyIdentifier is not null
func TestOriginatorOptionFPartyIdentifierNull(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = ""
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierCount validates OriginatorOptionF PartyIdentifier count is > 2
func TestOriginatorOptionFPartyIdentifierCount(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "X"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierUIDCount validates OriginatorOptionF PartyIdentifier unique ID has the
// correct count
func TestOriginatorOptionFPartyIdentifierUIDCount(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "B/C"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierUIDSlash validates OriginatorOptionF PartyIdentifier unique ID has the
// returns an error if '/' is not in the correct spot
func TestOriginatorOptionFPartyIdentifierSlash(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "CCPTFGH/"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierUIDInvalid validates OriginatorOptionF PartyIdentifier unique ID is valid
func TestOriginatorOptionFPartyIdentifierUIDInvalid(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "ZZZZFGH/"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierInvalidAccountSpace validates OriginatorOptionF PartyIdentifier with an invalid
// empty string in piece 2 for an Account
func TestOriginatorOptionFPartyInvalidAccountSpace(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "/ 1"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFPartyIdentifierInvalidSpace validates OriginatorOptionF PartyIdentifier with an invalid
// empty string in piece 6
func TestOriginatorOptionFPartyInvalidSpace(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.PartyIdentifier = "CCPT/ BDF"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrPartyIdentifier) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFName validates OriginatorOptionF Name is valid
func TestOriginatorOptionFName(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.Name = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFName) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFNameNull validates OriginatorOptionF Name is not null
func TestOriginatorOptionFNameNull(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.Name = ""
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFName) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineValid validates OriginatorOptionF Line* is valid
func TestOriginatorOptionFLineValid(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "Z/123"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFLine) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineSlash validates OriginatorOptionF Line* slash is valid
func TestOriginatorOptionFLineSlash(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "11123"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFLine) {
			t.Errorf("%T: %s", err, err)
		}
	}
}


// TestOriginatorOptionFLineOne validates OriginatorOptionF LineOne is valid
func TestOriginatorOptionFLineOne(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineOne = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFLine) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineTwo validates OriginatorOptionF LineTwo is valid
func TestOriginatorOptionFLineTwo(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineTwo = "®"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFLine) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorOptionFLineThree validates OriginatorOptionF LineThree is valid
func TestOriginatorOptionFLineThree(t *testing.T) {
	oof := mockOriginatorOptionF()
	oof.LineThree = "1/B"
	if err := oof.Validate(); err != nil {
		if !base.Match(err, ErrOptionFLine) {
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
		if !base.Match(err, ErrOptionFName) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrOptionFName) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
