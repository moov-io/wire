package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBeneficiaryCustomer creates a BeneficiaryCustomer
func mockBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := NewBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	bc.CoverPayment.SwiftLineOne = "Swift Line One"
	bc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	bc.CoverPayment.SwiftLineThree = "Swift Line Three"
	bc.CoverPayment.SwiftLineFour = "Swift Line Four"
	bc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return bc
}

// TestMockBeneficiaryCustomer validates mockBeneficiaryCustomer
func TestMockBeneficiaryCustomer(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	if err := bc.Validate(); err != nil {
		t.Error("mockBeneficiaryCustomer does not validate and will break other tests")
	}
}

// TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric validates BeneficiaryCustomer SwiftFieldTag is alphanumeric
func TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineOneAlphaNumeric validates BeneficiaryCustomer SwiftLineOne is alphanumeric
func TestBeneficiaryCustomerSwiftLineOneAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineOne = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric validates BeneficiaryCustomer SwiftLineTwo is alphanumeric
func TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineTwo = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric validates BeneficiaryCustomer SwiftLineThree is alphanumeric
func TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineThree = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineFourAlphaNumeric validates BeneficiaryCustomer SwiftLineFour is alphanumeric
func TestBeneficiaryCustomerSwiftLineFourAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFour = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric validates BeneficiaryCustomer SwiftLineFive is alphanumeric
func TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFive = "®"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryCustomerSwiftLineSixAlphaNumeric validates BeneficiaryCustomer SwiftLineSix is alphanumeric
func TestBeneficiaryCustomerSwiftLineSixAlphaNumeric(t *testing.T) {
	sr := mockBeneficiaryCustomer()
	sr.CoverPayment.SwiftLineSix = "Test"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryCustomerWrongLength parses a wrong BeneficiaryCustomer record length
func TestParseBeneficiaryCustomerWrongLength(t *testing.T) {
	var line = "{7059}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bc := mockBeneficiaryCustomer()
	fwm.SetBeneficiaryCustomer(bc)
	err := r.parseBeneficiaryCustomer()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryCustomerReaderParseError parses a wrong BeneficiaryCustomer reader parse error
func TestParseBeneficiaryCustomerReaderParseError(t *testing.T) {
	var line = "{7059}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bc := mockBeneficiaryCustomer()
	fwm.SetBeneficiaryCustomer(bc)
	err := r.parseBeneficiaryCustomer()
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

// TestBeneficiaryCustomerTagError validates a BeneficiaryCustomer tag
func TestBeneficiaryCustomerTagError(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.tag = "{9999}"
	if err := bc.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
