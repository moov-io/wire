package wire

import (
	"github.com/moov-io/base"
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
