package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// Remittance creates a Remittance
func mockRemittance() *Remittance {
	ri := NewRemittance()
	ri.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ri.CoverPayment.SwiftLineOne = "Swift Line One"
	ri.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ri.CoverPayment.SwiftLineThree = "Swift Line Three"
	ri.CoverPayment.SwiftLineFour = "Swift Line Four"
	return ri
}

// TestMockRemittance validates mockRemittance
func TestMockRemittance(t *testing.T) {
	ri := mockRemittance()

	require.NoError(t, ri.Validate(), "mockRemittance does not validate and will break other tests")
}

// TestRemittanceSwiftFieldTagAlphaNumeric validates Remittance SwiftFieldTag is alphanumeric
func TestRemittanceSwiftFieldTagAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftFieldTag = "®"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineOneAlphaNumeric validates Remittance SwiftLineOne is alphanumeric
func TestRemittanceSwiftLineOneAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineOne = "®"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineTwoAlphaNumeric validates Remittance SwiftLineTwo is alphanumeric
func TestRemittanceSwiftLineTwoAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineTwo = "®"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineThreeAlphaNumeric validates Remittance SwiftLineThree is alphanumeric
func TestRemittanceSwiftLineThreeAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineThree = "®"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineFourAlphaNumeric validates Remittance SwiftLineFour is alphanumeric
func TestRemittanceSwiftLineFourAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineFour = "®"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineFiveAlphaNumeric validates Remittance SwiftLineFive is alphanumeric
func TestRemittanceSwiftLineFiveAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineFive = "Test"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestRemittanceSwiftLineSixAlphaNumeric validates Remittance SwiftLineSix is alphanumeric
func TestRemittanceSwiftLineSixAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineSix = "Test"
	if err := ri.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceWrongLength parses a wrong Remittance record length
func TestParseRemittanceWrongLength(t *testing.T) {
	var line = "{7070}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ri := mockRemittance()
	fwm.SetRemittance(ri)
	err := r.parseRemittance()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseRemittanceReaderParseError parses a wrong Remittance reader parse error
func TestParseRemittanceReaderParseError(t *testing.T) {
	var line = "{7070}Swift®wift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ri := mockRemittance()
	fwm.SetRemittance(ri)
	err := r.parseRemittance()
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

// TestRemittanceTagError validates a Remittance tag
func TestRemittanceTagError(t *testing.T) {
	ri := mockRemittance()
	ri.tag = "{9999}"

	require.EqualError(t, ri.Validate(), fieldError("tag", ErrValidTagForType, ri.tag).Error())
}
