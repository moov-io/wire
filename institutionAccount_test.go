package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

//  InstitutionAccount creates a InstitutionAccount
func mockInstitutionAccount() *InstitutionAccount {
	iAccount := NewInstitutionAccount()
	iAccount.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	iAccount.CoverPayment.SwiftLineOne = "Swift Line One"
	iAccount.CoverPayment.SwiftLineTwo = "Swift Line Two"
	iAccount.CoverPayment.SwiftLineThree = "Swift Line Three"
	iAccount.CoverPayment.SwiftLineFour = "Swift Line Four"
	iAccount.CoverPayment.SwiftLineFive = "Swift Line Five"
	return iAccount
}

// TestMockInstitutionAccount validates mockInstitutionAccount
func TestMockInstitutionAccount(t *testing.T) {
	iAccount := mockInstitutionAccount()
	if err := iAccount.Validate(); err != nil {
		t.Error("mockInstitutionAccount does not validate and will break other tests")
	}
}

// TestInstitutionAccountSwiftFieldTagAlphaNumeric validates InstitutionAccount SwiftFieldTag is alphanumeric
func TestInstitutionAccountSwiftFieldTagAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftFieldTag = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineOneAlphaNumeric validates InstitutionAccount SwiftLineOne is alphanumeric
func TestInstitutionAccountSwiftLineOneAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineOne = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineTwoAlphaNumeric validates InstitutionAccount SwiftLineTwo is alphanumeric
func TestInstitutionAccountSwiftLineTwoAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineTwo = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineThreeAlphaNumeric validates InstitutionAccount SwiftLineThree is alphanumeric
func TestInstitutionAccountSwiftLineThreeAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineThree = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineFourAlphaNumeric validates InstitutionAccount SwiftLineFour is alphanumeric
func TestInstitutionAccountSwiftLineFourAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineFour = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineFiveAlphaNumeric validates InstitutionAccount SwiftLineFive is alphanumeric
func TestInstitutionAccountSwiftLineFiveAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineFive = "®"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstitutionAccountSwiftLineSixAlphaNumeric validates InstitutionAccount SwiftLineSix is alphanumeric
func TestInstitutionAccountSwiftLineSixAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineSix = "Test"
	if err := iAccount.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInstitutionAccountWrongLength parses a wrong InstitutionAccount record length
func TestParseInstitutionAccountWrongLength(t *testing.T) {
	var line = "{7057}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	iAccount := mockInstitutionAccount()
	fwm.SetInstitutionAccount(iAccount)
	err := r.parseInstitutionAccount()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(186, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInstitutionAccountReaderParseError parses a wrong InstitutionAccount reader parse error
func TestParseInstitutionAccountReaderParseError(t *testing.T) {
	var line = "{7057}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	iAccount := mockInstitutionAccount()
	fwm.SetInstitutionAccount(iAccount)
	err := r.parseInstitutionAccount()
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
