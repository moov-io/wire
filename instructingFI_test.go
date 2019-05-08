package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockInstructingFI creates a InstructingFI
func mockInstructingFI() *InstructingFI {
	ifi := NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ifi
}

// TestMockInstructingFI validates mockInstructingFI
func TestMockInstructingFI(t *testing.T) {
	bfi := mockInstructingFI()
	if err := bfi.Validate(); err != nil {
		t.Error("mockInstructingFI does not validate and will break other tests")
	}
}

// TestInstructingFIIdentificationCodeValid validates InstructingFI IdentificationCode
func TestInstructingFIIdentificationCodeValid(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = "Football Card ID"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIIdentificationCodeFI validates InstructingFI IdentificationCode is an FI code
func TestInstructingFIIdentificationCodeFI(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = "1"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIIdentifierAlphaNumeric validates InstructingFI Identifier is alphanumeric
func TestInstructingFIIdentifierAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Identifier = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFINameAlphaNumeric validates InstructingFI Name is alphanumeric
func TestInstructingFINameAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Name = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIAddressLineOneAlphaNumeric validates InstructingFI AddressLineOne is alphanumeric
func TestInstructingFIAddressLineOneAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineOne = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIAddressLineTwoAlphaNumeric validates InstructingFI AddressLineTwo is alphanumeric
func TestInstructingFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineTwo = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIAddressLineThreeAlphaNumeric validates InstructingFI AddressLineThree is alphanumeric
func TestInstructingFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineThree = "®"
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIIdentificationCodeRequired validates InstructingFI IdentificationCode is required
func TestInstructingFIIdentificationCodeRequired(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = ""
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInstructingFIIdentifierRequired validates InstructingFI Identifier is required
func TestInstructingFIIdentifierRequired(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Identifier = ""
	if err := bfi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInstructingFIWrongLength parses a wrong InstructingFI record length
func TestParseInstructingFIWrongLength(t *testing.T) {
	var line = "{5200}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ofi := mockInstructingFI()
	fwm.SetInstructingFI(ofi)
	err := r.parseInstructingFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInstructingFIReaderParseError parses a wrong InstructingFI reader parse error
func TestParseInstructingFIReaderParseError(t *testing.T) {
	var line = "{5200}D123456789                         ®I Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ofi := mockInstructingFI()
	fwm.SetInstructingFI(ofi)
	err := r.parseInstructingFI()
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
