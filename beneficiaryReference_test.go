package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBeneficiaryReference creates a BeneficiaryReference
func mockBeneficiaryReference() *BeneficiaryReference {
	br := NewBeneficiaryReference()
	br.BeneficiaryReference = "Reference"
	return br
}

// TestMockBeneficiary validates mockBeneficiaryReference
func TestMockBeneficiaryReference(t *testing.T) {
	br := mockBeneficiaryReference()
	if err := br.Validate(); err != nil {
		t.Error("mockBeneficiaryReference does not validate and will break other tests")
	}
}

// TestBeneficiaryReferenceAlphaNumeric validates BeneficiaryReference is alphanumeric
func TestBeneficiaryReferenceAlphaNumeric(t *testing.T) {
	br := mockBeneficiaryReference()
	br.BeneficiaryReference = "®"
	if err := br.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryReferenceWrongLength parses a wrong BeneficiaryReference record length
func TestParseBeneficiaryReferenceWrongLength(t *testing.T) {
	var line = "{4320}Reference      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	err := r.parseBeneficiaryReference()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(22, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBeneficiaryReferenceReaderParseError parses a wrong BeneficiaryReference reader parse error
func TestParseBeneficiaryReferenceReaderParseError(t *testing.T) {
	var line = "{4320}Reference®      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	br := mockBeneficiaryReference()
	fwm.SetBeneficiaryReference(br)
	err := r.parseBeneficiaryReference()
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