package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockOriginatorFI creates a OriginatorFI
func mockOriginatorFI() *OriginatorFI {
	ofi := NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ofi
}

// TestMockOriginatorFI validates mockOriginatorFI
func TestMockOriginatorFI(t *testing.T) {
	ofi := mockOriginatorFI()
	if err := ofi.Validate(); err != nil {
		t.Error("mockOriginatorFI does not validate and will break other tests")
	}
}

// TestOriginatorFIIdentificationCodeValid validates OriginatorFI IdentificationCode
func TestOriginatorFIIdentificationCodeValid(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = "Football Card ID"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIIdentificationCodeFI validates OriginatorFI IdentificationCode is an FI code
func TestOriginatorFIIdentificationCodeFI(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = "1"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIIdentifierAlphaNumeric validates OriginatorFI Identifier is alphanumeric
func TestOriginatorFIIdentifierAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Identifier = "®"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFINameAlphaNumeric validates OriginatorFI Name is alphanumeric
func TestOriginatorFINameAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Name = "®"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIAddressLineOneAlphaNumeric validates OriginatorFI AddressLineOne is alphanumeric
func TestOriginatorFIAddressLineOneAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineOne = "®"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIAddressLineTwoAlphaNumeric validates OriginatorFI AddressLineTwo is alphanumeric
func TestOriginatorFIAddressLineTwoAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineTwo = "®"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIAddressLineThreeAlphaNumeric validates OriginatorFI AddressLineThree is alphanumeric
func TestOriginatorFIAddressLineThreeAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineThree = "®"
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIIdentificationCodeRequired validates OriginatorFI IdentificationCode is required
func TestOriginatorFIIdentificationCodeRequired(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = ""
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorFIIdentifierRequired validates OriginatorFI Identifier is required
func TestOriginatorFIIdentifierRequired(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Identifier = ""
	if err := ofi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorFIWrongLength parses a wrong OriginatorFI record length
func TestParseOriginatorFIWrongLength(t *testing.T) {
	var line = "{5100}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	err := r.parseOriginatorFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorFIReaderParseError parses a wrong OriginatorFI reader parse error
func TestParseOriginatorFIReaderParseError(t *testing.T) {
	var line = "{5100}D123456789                         ®I Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ofi := mockOriginatorFI()
	fwm.SetOriginatorFI(ofi)
	err := r.parseOriginatorFI()
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
