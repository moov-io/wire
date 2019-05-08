package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBeneficiaryIntermediaryFI creates a BeneficiaryIntermediaryFI
func mockBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	bifi := NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bifi
}

// TestMockBeneficiaryIntermediaryFI validates mockBeneficiaryIntermediaryFI
func TestMockBeneficiaryIntermediaryFI(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	if err := bifi.Validate(); err != nil {
		t.Error("mockBeneficiaryIntermediaryFI does not validate and will break other tests")
	}
}

// TestBeneficiaryIntermediaryFIIdentificationCodeValid validates BeneficiaryIntermediaryFI IdentificationCode
func TestBeneficiaryIntermediaryFIIdentificationCodeValid(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = "Football Card ID"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIIdentificationCodeFI validates BeneficiaryIntermediaryFI IdentificationCode is an FI code
func TestBeneficiaryIntermediaryFIIdentificationCodeFI(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = "1"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric validates BeneficiaryIntermediaryFI Identifier is alphanumeric
func TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = "®"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFINameAlphaNumeric validates BeneficiaryIntermediaryFI Name is alphanumeric
func TestBeneficiaryIntermediaryFINameAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Name = "®"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineOne is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineOne = "®"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineTwo is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineTwo = "®"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineThree is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineThree = "®"
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIIdentificationCodeRequired validates BeneficiaryIntermediaryFI IdentificationCode is required
func TestBeneficiaryIntermediaryFIIdentificationCodeRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = ""
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBeneficiaryIntermediaryFIIdentifierRequired validates BeneficiaryIntermediaryFI Identifier is required
func TestBeneficiaryIntermediaryFIIdentifierRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = ""
	if err := bifi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryIntermediaryFIWrongLength parses a wrong BeneficiaryIntermediaryFI record length
func TestParseBeneficiaryIntermediaryFIWrongLength(t *testing.T) {
	var line = "{4000}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	err := r.parseBeneficiaryIntermediaryFI()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryIntermediaryFIReaderParseError parses a wrong BeneficiaryIntermediaryFI reader parse error
func TestParseBeneficiaryIntermediaryFIReaderParseError(t *testing.T) {
	var line = "{4000}D123456789                         F® Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)
	err := r.parseBeneficiaryIntermediaryFI()
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
