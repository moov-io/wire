package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockReceiverDepositoryInstitution creates a ReceiverDepositoryInstitution
func mockReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	rdi := NewReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	return rdi
}

// TestMockReceiverDepositoryInstitution validates mockReceiverDepositoryInstitution
func TestMockReceiverDepositoryInstitution(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	if err := rdi.Validate(); err != nil {
		t.Error("mockReceiverDepositoryInstitution does not validate and will break other tests")
	}
}

// TestReceiverABANumberAlphaNumeric validates ReceiverDepositoryInstitution ReceiverABANumber is alphanumeric
func TestReceiverABANumberAlphaNumeric(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "®"
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestReceiverShortNameAlphaNumeric validates ReceiverDepositoryInstitution ReceiverShortName is alphanumeric
func TestReceiverShortNameAlphaNumeric(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverShortName = "®"
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestReceiverABANumberRequired validates ReceiverDepositoryInstitution ReceiverABANumber is required
func TestReceiverABANumberRequired(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = ""
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestReceiverShortNameRequired validates ReceiverDepositoryInstitution ReceiverShortName is required
func TestReceiverShortNameRequired(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverShortName = ""
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseReceiverWrongLength parses a wrong Receiver record length
func TestParseReceiverWrongLength(t *testing.T) {
	var line = "{3400}00"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	err := r.parseReceiverDepositoryInstitution()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(15, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseReceiverReaderParseError parses a wrong Receiver reader parse error
func TestParseReceiverReaderParseError(t *testing.T) {
	var line = "{3400}2313Z0104Citadel           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rdi := mockReceiverDepositoryInstitution()
	fwm.SetReceiverDepositoryInstitution(rdi)
	err := r.parseReceiverDepositoryInstitution()
	if err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestReceiverDepositoryInstitutionTagError validates a ReceiverDepositoryInstitution tag
func TestReceiverDepositoryInstitutionTagError(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.tag = "{9999}"
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
