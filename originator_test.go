package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// mockOriginator creates a Originator
func mockOriginator() *Originator {
	o := NewOriginator()
	o.Personal.IdentificationCode = PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address One"
	o.Personal.Address.AddressLineTwo = "Address Two"
	o.Personal.Address.AddressLineThree = "Address Three"
	return o
}

// TestMockOriginator validates mockOriginator
func TestMockOriginator(t *testing.T) {
	o := mockOriginator()

	require.NoError(t, o.Validate(), "mockOriginator does not validate and will break other tests")
}

// TestOriginatorIdentificationCodeValid validates Originator IdentificationCode
func TestOriginatorIdentificationCodeValid(t *testing.T) {
	o := mockOriginator()
	o.Personal.IdentificationCode = "Baseball Card ID"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrIdentificationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorIdentifierAlphaNumeric validates Originator Identifier is alphanumeric
func TestOriginatorIdentifierAlphaNumeric(t *testing.T) {
	o := mockOriginator()
	o.Personal.Identifier = "®"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorNameAlphaNumeric validates Originator Name is alphanumeric
func TestOriginatorNameAlphaNumeric(t *testing.T) {
	o := mockOriginator()
	o.Personal.Name = "®"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorAddressLineOneAlphaNumeric validates Originator AddressLineOne is alphanumeric
func TestOriginatorAddressLineOneAlphaNumeric(t *testing.T) {
	o := mockOriginator()
	o.Personal.Address.AddressLineOne = "®"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorAddressLineTwoAlphaNumeric validates Originator AddressLineTwo is alphanumeric
func TestOriginatorAddressLineTwoAlphaNumeric(t *testing.T) {
	o := mockOriginator()
	o.Personal.Address.AddressLineTwo = "®"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorAddressLineThreeAlphaNumeric validates Originator AddressLineThree is alphanumeric
func TestOriginatorAddressLineThreeAlphaNumeric(t *testing.T) {
	o := mockOriginator()
	o.Personal.Address.AddressLineThree = "®"
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorIdentificationCodeRequired validates Originator IdentificationCode is required
func TestOriginatorIdentificationCodeRequired(t *testing.T) {
	o := mockOriginator()
	o.Personal.IdentificationCode = ""
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestOriginatorIdentifierRequired validates Originator Identifier is required
func TestOriginatorIdentifierRequired(t *testing.T) {
	o := mockOriginator()
	o.Personal.Identifier = ""
	if err := o.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorWrongLength parses a wrong Originator record length
func TestParseOriginatorWrongLength(t *testing.T) {
	var line = "{5000}11234                              Name                               Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	o := mockOriginator()
	fwm.SetOriginator(o)
	err := r.parseOriginator()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(181, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseOriginatorReaderParseError parses a wrong Originator reader parse error
func TestParseOriginatorReaderParseError(t *testing.T) {
	var line = "{5000}11234                              ®ame                               Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	o := mockOriginator()
	fwm.SetOriginator(o)
	err := r.parseOriginator()
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

// TestOriginatorTagError validates a Originator tag
func TestOriginatorTagError(t *testing.T) {
	o := mockOriginator()
	o.tag = "{9999}"

	require.EqualError(t, o.Validate(), fieldError("tag", ErrValidTagForType, o.tag).Error())
}
