package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockSenderReference creates a SenderReference
func mockSenderReference() *SenderReference {
	sr := NewSenderReference()
	sr.SenderReference = "Sender Reference"
	return sr
}

// TestMockSenderReference validates mockSenderReference
func TestMockSenderReference(t *testing.T) {
	sr := mockSenderReference()
	if err := sr.Validate(); err != nil {
		t.Error("mockSenderReference does not validate and will break other tests")
	}
}

// TestSenderReferenceAlphaNumeric validates SenderReference is alphanumeric
func TestSenderReferenceAlphaNumeric(t *testing.T) {
	sr := mockSenderReference()
	sr.SenderReference = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderReferenceWrongLength parses a wrong SenderReference record length
func TestParseSenderReferenceWrongLength(t *testing.T) {
	var line = "{3320}Se"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	err := r.parseSenderReference()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(22, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderReferenceReaderParseError parses a wrong SenderReference reader parse error
func TestParseSenderReferenceReaderParseError(t *testing.T) {
	var line = "{3320}Sender®Reference"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sr := mockSenderReference()
	fwm.SetSenderReference(sr)
	err := r.parseSenderReference()
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