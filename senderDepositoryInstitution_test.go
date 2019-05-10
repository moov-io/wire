package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockSenderDepositoryInstitution creates a SenderDepositoryInstitution
func mockSenderDepositoryInstitution() *SenderDepositoryInstitution {
	sdi := NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	return sdi
}

// TestMockSenderDepositoryInstitution validates mockSenderDepositoryInstitution
func TestMockSenderDepositoryInstitution(t *testing.T) {
	sdi := mockSenderDepositoryInstitution()
	if err := sdi.Validate(); err != nil {
		t.Error("mockSenderDepositoryInstitution does not validate and will break other tests")
	}
}

// TestSenderABANumberAlphaNumeric validates SenderDepositoryInstitution SenderABANumber is alphanumeric
func TestSenderABANumberAlphaNumeric(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderABANumber = "®"
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderShortNameAlphaNumeric validates SenderDepositoryInstitution SenderShortName is alphanumeric
func TestSenderShortNameAlphaNumeric(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderShortName = "®"
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderABANumberRequired validates SenderDepositoryInstitution SenderABANumber is required
func TestSenderABANumberRequired(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderABANumber = ""
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderShortNameRequired validates SenderDepositoryInstitution SenderShortName is required
func TestSenderShortNameRequired(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderShortName = ""
	if err := rdi.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderWrongLength parses a wrong Sender record length
func TestParseSenderWrongLength(t *testing.T) {
	var line = "{3100}0012"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	err := r.parseSenderDepositoryInstitution()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(15, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderReaderParseError parses a wrong Sender reader parse error
func TestParseSenderReaderParseError(t *testing.T) {
	var line = "{3100}1210Z2882Wells Fargo NA    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sdi := mockSenderDepositoryInstitution()
	fwm.SetSenderDepositoryInstitution(sdi)
	err := r.parseSenderDepositoryInstitution()
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
