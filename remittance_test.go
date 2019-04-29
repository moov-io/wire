package wire

import (
	"github.com/moov-io/base"
	"testing"
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
	if err := ri.Validate(); err != nil {
		t.Error("mockRemittance does not validate and will break other tests")
	}
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
