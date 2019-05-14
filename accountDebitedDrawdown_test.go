package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockAccountDebitedDrawdown creates a AccountDebitedDrawdown
func mockAccountDebitedDrawdown() *AccountDebitedDrawdown {
	debitDD := NewAccountDebitedDrawdown()
	debitDD.IdentificationCode = DemandDepositAccountNumber
	debitDD.Identifier = "123456789"
	debitDD.Name = "debitDD Name"
	debitDD.Address.AddressLineOne = "Address One"
	debitDD.Address.AddressLineTwo = "Address Two"
	debitDD.Address.AddressLineThree = "Address Three"
	return debitDD
}

// TestMockAccountDebitedDrawdown validates mockAccountDebitedDrawdown
func TestMockAccountDebitedDrawdown(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	if err := debitDD.Validate(); err != nil {
		t.Error("mockAccountDebitedDrawdown does not validate and will break other tests")
	}
}

// TestADDIdentifierAlphaNumeric validates Name is alphanumeric
func TestADDIdentifierAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDNameAlphaNumeric validates Identifier is alphanumeric
func TestADDNameAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDAddressLineOneAlphaNumeric validates AddressLineOne is alphanumeric
func TestADDAddressLineOneAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineOne = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDAddressLineTwoAlphaNumeric validates AddressLineTwo is alphanumeric
func TestADDAddressLineTwoAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineTwo = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDAddressLineThreeAlphaNumeric validates AddressLineThree is alphanumeric
func TestADDAddressLineThreeAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineThree = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDIdentifierRequired validates Identifier is required
func TestADDIdentifierRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDNameRequired validates Name is required
func TestADDNameRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDIdentificationRequired validates IdentificationCode is required
func TestADDIdentificationCodeRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDIdentificationCodeValid validates IdentificationCode
func TestADDIdentificationCodeValid(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = TaxIdentificationNumber
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestADDIdentificationCodeBogus validates IdentificationCode if the IdentificationCode is bogus
func TestIdentificationCodeBogus(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = "Card ID"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAccountDebitedDrawdownWrongLength parses a wrong AccountDebitedDrawdown record length
func TestParseAccountDebitedDrawdownWrongLength(t *testing.T) {
	var line = "{4400}D123456789                         debitDD Name                       Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)
	err := r.parseAccountDebitedDrawdown()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseAccountDebitedDrawdownReaderParseError parses a wrong AccountDebitedDrawdown reader parse error
func TestParseAccountDebitedDrawdownReaderParseError(t *testing.T) {
	var line = "{4400}D123456789                         debitDD ®ame                       Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	debitDD := mockAccountDebitedDrawdown()
	fwm.SetAccountDebitedDrawdown(debitDD)
	err := r.parseAccountDebitedDrawdown()
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

// TestAccountDebitedDrawdownTagError validates AccountDebitedDrawdown tag
func TestAccountDebitedDrawdownTagError(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.tag = "{9999}"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
