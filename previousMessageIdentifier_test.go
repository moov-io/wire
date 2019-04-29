package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockPreviousMessageIdentifier creates a PreviousMessageIdentifier
func mockPreviousMessageIdentifier() *PreviousMessageIdentifier {
	pmi := NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	return pmi
}

// TestMockPreviousMessageIdentifier validates mockPreviousMessageIdentifier
func TestMockPreviousMessageIdentifier(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	if err := pmi.Validate(); err != nil {
		t.Error("mockPreviousMessageIdentifier does not validate and will break other tests")
	}
}

// TestPreviousMessageIdentifierAlphaNumeric validates PreviousMessageIdentifier is alphanumeric
func TestPreviousMessageIdentifierAlphaNumeric(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "®"
	if err := pmi.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
