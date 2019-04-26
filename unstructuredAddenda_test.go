package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// UnstructuredAddenda creates a UnstructuredAddenda
func mockUnstructuredAddenda() *UnstructuredAddenda {
	ua := NewUnstructuredAddenda()
	ua.AddendaLength = "0014"
	ua.Addenda = "This is a test"
	return ua
}

// TestMockUnstructuredAddenda validates mockUnstructuredAddenda
func TestMockUnstructuredAddenda(t *testing.T) {
	ua := mockUnstructuredAddenda()
	if err := ua.Validate(); err != nil {
		t.Error("mockUnstructuredAddenda does not validate and will break other tests")
	}
}

// TestAddendaLengthNumeric validates UnstructuredAddenda Length is numeric
func TestAddendaLengthNumeric(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.AddendaLength = "09T4"
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddendaAlphaNumeric validates UnstructuredAddenda Addenda is alphanumeric
func TestAddendaAlphaNumeric(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.Addenda = "®"
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddendaLengthRequired validates UnstructuredAddenda Length is required
func TestAddendaLengthRequired(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.AddendaLength = ""
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
