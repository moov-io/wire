package wire

import (
	"github.com/moov-io/base"
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

// TestIdentifierAlphaNumeric validates Name is alphanumeric
func TestIdentifierAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestNameAlphaNumeric validates Identifier is alphanumeric
func TestNameAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddressLineOneAlphaNumeric validates AddressLineOne is alphanumeric
func TestAddressLineOneAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineOne = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddressLineTwoAlphaNumeric validates AddressLineTwo is alphanumeric
func TestAddressLineTwoAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineTwo = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddressLineThreeAlphaNumeric validates AddressLineThree is alphanumeric
func TestAddressLineThreeAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineThree = "®"
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentifierRequired validates Identifier is required
func TestIdentifierRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestNameRequired validates Name is required
func TestNameRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentificationRequired validates IdentificationCode is required
func TestIdentificationCodeNull(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = ""
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIdentificationCodeValid validates IdentificationCode
func TestIdentificationCodeRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = TaxIdentificationNumber
	if err := debitDD.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
