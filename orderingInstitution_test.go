package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

//  OrderingInstitution creates a OrderingInstitution
func mockOrderingInstitution() *OrderingInstitution {
	oi := NewOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oi.CoverPayment.SwiftLineOne = "Swift Line One"
	oi.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oi.CoverPayment.SwiftLineThree = "Swift Line Three"
	oi.CoverPayment.SwiftLineFour = "Swift Line Four"
	oi.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oi
}

// TestMockOrderingInstitution validates mockOrderingInstitution
func TestMockOrderingInstitution(t *testing.T) {
	oi := mockOrderingInstitution()
	if err := oi.Validate(); err != nil {
		t.Error("mockOrderingInstitution does not validate and will break other tests")
	}
}

// TestOrderingInstitutionSwiftFieldTagAlphaNumeric validates OrderingInstitution SwiftFieldTag is alphanumeric
func TestOrderingInstitutionSwiftFieldTagAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineOneAlphaNumeric validates OrderingInstitution SwiftLineOne is alphanumeric
func TestOrderingInstitutionSwiftLineOneAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineOne = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineTwoAlphaNumeric validates OrderingInstitution SwiftLineTwo is alphanumeric
func TestOrderingInstitutionSwiftLineTwoAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineTwo = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineThreeAlphaNumeric validates OrderingInstitution SwiftLineThree is alphanumeric
func TestOrderingInstitutionSwiftLineThreeAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineThree = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineFourAlphaNumeric validates OrderingInstitution SwiftLineFour is alphanumeric
func TestOrderingInstitutionSwiftLineFourAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineFour = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineFiveAlphaNumeric validates OrderingInstitution SwiftLineFive is alphanumeric
func TestOrderingInstitutionSwiftLineFiveAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineFive = "®"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingInstitutionSwiftLineSixAlphaNumeric validates OrderingInstitution SwiftLineSix is alphanumeric
func TestOrderingInstitutionSwiftLineSixAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineSix = "Test"
	if err := oi.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOrderingInstitutionWrongLength parses a wrong OrderingInstitution record length
func TestParseOrderingInstitutionWrongLength(t *testing.T) {
	var line = "{7052}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oi := mockOrderingInstitution()
	fwm.SetOrderingInstitution(oi)
	err := r.parseOrderingInstitution()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOrderingInstitutionReaderParseError parses a wrong OrderingInstitution reader parse error
func TestParseOrderingInstitutionReaderParseError(t *testing.T) {
	var line = "{7052}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oi := mockOrderingInstitution()
	fwm.SetOrderingInstitution(oi)
	err := r.parseOrderingInstitution()
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
