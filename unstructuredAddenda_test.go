package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// UnstructuredAddenda creates a UnstructuredAddenda
func mockUnstructuredAddenda() *UnstructuredAddenda {
	ua := NewUnstructuredAddenda()
	ua.AddendaLength = "0020"
	ua.Addenda = "Unstructured Addenda"
	return ua
}

// TestMockUnstructuredAddenda validates mockUnstructuredAddenda
func TestMockUnstructuredAddenda(t *testing.T) {
	ua := mockUnstructuredAddenda()

	require.NoError(t, ua.Validate(), "mockUnstructuredAddenda does not validate and will break other tests")
}

// TestUnstructuredAddendaLengthNumeric validates UnstructuredAddenda Length is numeric
func TestAddendaLengthNumeric(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.AddendaLength = "09T4"
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestUnstructuredAddendaAlphaNumeric validates UnstructuredAddenda Addenda is alphanumeric
func TestAddendaAlphaNumeric(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.Addenda = "®"
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestUnstructuredAddendaLengthRequired validates UnstructuredAddenda Length is required
func TestAddendaLengthRequired(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.AddendaLength = ""
	if err := ua.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseUnstructuredAddendaWrongLength parses a wrong Addenda record length
func TestParseAddendaWrongLength(t *testing.T) {
	var line = "{8200}0020Unstructured Addenda  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	err := r.parseUnstructuredAddenda()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(30, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseUnstructuredAddendaReaderParseError parses a wrong Addenda reader parse error
func TestParseUnstructuredAddendaReaderParseError(t *testing.T) {
	var line = "{8200}0020®nstructured Addenda"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	err := r.parseUnstructuredAddenda()
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

// TestUnstructuredAddendaTagError validates a UnstructuredAddenda tag
func TestUnstructuredAddendaTagError(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.tag = "{9999}"

	require.EqualError(t, ua.Validate(), fieldError("tag", ErrValidTagForType, ua.tag).Error())
}
