package wire

import (
	"strings"
	"testing"

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

	err := ua.Validate()

	require.EqualError(t, err, fieldError("AddendaLength", ErrNonNumeric, ua.AddendaLength).Error())
}

// TestUnstructuredAddendaAlphaNumeric validates UnstructuredAddenda Addenda is alphanumeric
func TestAddendaAlphaNumeric(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.Addenda = "®"

	err := ua.Validate()

	require.EqualError(t, err, fieldError("Addenda", ErrNonAlphanumeric, ua.Addenda).Error())
}

// TestUnstructuredAddendaLengthRequired validates UnstructuredAddenda Length is required
func TestAddendaLengthRequired(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.AddendaLength = ""

	err := ua.Validate()

	require.EqualError(t, err, fieldError("AddendaLength", ErrFieldRequired).Error())
}

// TestParseUnstructuredAddendaWrongLength parses a wrong Addenda record length
func TestParseAddendaWrongLength(t *testing.T) {
	var line = "{8200}0020Unstructured Addenda  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseUnstructuredAddenda()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(30, len(r.line))).Error())
}

// TestParseUnstructuredAddendaReaderParseError parses a wrong Addenda reader parse error
func TestParseUnstructuredAddendaReaderParseError(t *testing.T) {
	var line = "{8200}0020®nstructured Addenda"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseUnstructuredAddenda()

	require.EqualError(t, err, r.parseError(fieldError("Addenda", ErrNonAlphanumeric, "®nstructured Addend")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Addenda", ErrNonAlphanumeric, "®nstructured Addend")).Error())
}

// TestUnstructuredAddendaTagError validates a UnstructuredAddenda tag
func TestUnstructuredAddendaTagError(t *testing.T) {
	ua := mockUnstructuredAddenda()
	ua.tag = "{9999}"

	require.EqualError(t, ua.Validate(), fieldError("tag", ErrValidTagForType, ua.tag).Error())
}
