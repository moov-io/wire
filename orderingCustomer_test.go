package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

//  OrderingCustomer creates a OrderingCustomer
func mockOrderingCustomer() *OrderingCustomer {
	oc := NewOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oc.CoverPayment.SwiftLineOne = "Swift Line One"
	oc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oc.CoverPayment.SwiftLineThree = "Swift Line Three"
	oc.CoverPayment.SwiftLineFour = "Swift Line Four"
	oc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oc
}

// TestMockOrderingCustomer validates mockOrderingCustomer
func TestMockOrderingCustomer(t *testing.T) {
	oc := mockOrderingCustomer()
	if err := oc.Validate(); err != nil {
		t.Error("mockOrderingCustomer does not validate and will break other tests")
	}
}

// TestOrderingCustomerSwiftFieldTagAlphaNumeric validates OrderingCustomer SwiftFieldTag is alphanumeric
func TestOrderingCustomerSwiftFieldTagAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineOneAlphaNumeric validates OrderingCustomer SwiftLineOne is alphanumeric
func TestOrderingCustomerSwiftLineOneAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineOne = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineTwoAlphaNumeric validates OrderingCustomer SwiftLineTwo is alphanumeric
func TestOrderingCustomerSwiftLineTwoAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineTwo = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineThreeAlphaNumeric validates OrderingCustomer SwiftLineThree is alphanumeric
func TestOrderingCustomerSwiftLineThreeAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineThree = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineFourAlphaNumeric validates OrderingCustomer SwiftLineFour is alphanumeric
func TestOrderingCustomerSwiftLineFourAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineFour = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineFiveAlphaNumeric validates OrderingCustomer SwiftLineFive is alphanumeric
func TestOrderingCustomerSwiftLineFiveAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineFive = "®"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOrderingCustomerSwiftLineSixAlphaNumeric validates OrderingCustomer SwiftLineSix is alphanumeric
func TestOrderingCustomerSwiftLineSixAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineSix = "Test"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOrderingCustomerWrongLength parses a wrong OrderingCustomer record length
func TestParseOrderingCustomerWrongLength(t *testing.T) {
	var line = "{7050}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oc := mockOrderingCustomer()
	fwm.SetOrderingCustomer(oc)
	err := r.parseOrderingCustomer()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOrderingCustomerReaderParseError parses a wrong OrderingCustomer reader parse error
func TestParseOrderingCustomerReaderParseError(t *testing.T) {
	var line = "{7050}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	oc := mockOrderingCustomer()
	fwm.SetOrderingCustomer(oc)
	err := r.parseOrderingCustomer()
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

// TestOrderingCustomerTagError validates a OrderingCustomer tag
func TestOrderingCustomerTagError(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.tag = "{9999}"
	if err := oc.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
