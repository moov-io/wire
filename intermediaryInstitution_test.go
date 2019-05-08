package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

//  IntermediaryInstitution creates a IntermediaryInstitution
func mockIntermediaryInstitution() *IntermediaryInstitution {
	ii := NewIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ii.CoverPayment.SwiftLineOne = "Swift Line One"
	ii.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ii.CoverPayment.SwiftLineThree = "Swift Line Three"
	ii.CoverPayment.SwiftLineFour = "Swift Line Four"
	ii.CoverPayment.SwiftLineFive = "Swift Line Five"
	return ii
}

// TestMockIntermediaryInstitution validates mockIntermediaryInstitution
func TestMockIntermediaryInstitution(t *testing.T) {
	ii := mockIntermediaryInstitution()
	if err := ii.Validate(); err != nil {
		t.Error("mockIntermediaryInstitution does not validate and will break other tests")
	}
}

// TestIntermediaryInstitutionSwiftFieldTagAlphaNumeric validates IntermediaryInstitution SwiftFieldTag is alphanumeric
func TestIntermediaryInstitutionSwiftFieldTagAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineOneAlphaNumeric validates IntermediaryInstitution SwiftLineOne is alphanumeric
func TestIntermediaryInstitutionSwiftLineOneAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineOne = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineTwoAlphaNumeric validates IntermediaryInstitution SwiftLineTwo is alphanumeric
func TestIntermediaryInstitutionSwiftLineTwoAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineTwo = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineThreeAlphaNumeric validates IntermediaryInstitution SwiftLineThree is alphanumeric
func TestIntermediaryInstitutionSwiftLineThreeAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineThree = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineFourAlphaNumeric validates IntermediaryInstitution SwiftLineFour is alphanumeric
func TestIntermediaryInstitutionSwiftLineFourAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineFour = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineFiveAlphaNumeric validates IntermediaryInstitution SwiftLineFive is alphanumeric
func TestIntermediaryInstitutionSwiftLineFiveAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineFive = "®"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIntermediaryInstitutionSwiftLineSixAlphaNumeric validates IntermediaryInstitution SwiftLineSix is alphanumeric
func TestIntermediaryInstitutionSwiftLineSixAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineSix = "Test"
	if err := ii.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseIntermediaryInstitutionWrongLength parses a wrong IntermediaryInstitution record length
func TestParseIntermediaryInstitutionWrongLength(t *testing.T) {
	var line = "{7056}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ii := mockIntermediaryInstitution()
	fwm.SetIntermediaryInstitution(ii)
	err := r.parseIntermediaryInstitution()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseIntermediaryInstitutionReaderParseError parses a wrong IntermediaryInstitution reader parse error
func TestParseIntermediaryInstitutionReaderParseError(t *testing.T) {
	var line = "{7056}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ii := mockIntermediaryInstitution()
	fwm.SetIntermediaryInstitution(ii)
	err := r.parseIntermediaryInstitution()
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
