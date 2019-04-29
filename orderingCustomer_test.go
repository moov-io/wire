package wire

import (
	"github.com/moov-io/base"
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
