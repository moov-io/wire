package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, pmi.Validate(), "mockPreviousMessageIdentifier does not validate and will break other tests")
}

// TestPreviousMessageIdentifierAlphaNumeric validates PreviousMessageIdentifier is alphanumeric
func TestPreviousMessageIdentifierAlphaNumeric(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "®"

	err := pmi.Validate()

	require.EqualError(t, err, fieldError("PreviousMessageIdentifier", ErrNonAlphanumeric, pmi.PreviousMessageIdentifier).Error())
}

// TestParsePreviousMessageIdentifierWrongLength parses a wrong PreviousMessageIdentifier record length
func TestParsePreviousMessageIdentifierWrongLength(t *testing.T) {
	var line = "{3500}Previous"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePreviousMessageIdentifier()

	require.EqualError(t, err, r.parseError(fieldError("PreviousMessageIdentifier", ErrValidLength)).Error())
}

// TestParsePreviousMessageIdentifierReaderParseError parses a wrong PreviousMessageIdentifier reader parse error
func TestParsePreviousMessageIdentifierReaderParseError(t *testing.T) {
	var line = "{3500}Previous®Message Iden"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePreviousMessageIdentifier()

	require.EqualError(t, err, r.parseError(fieldError("PreviousMessageIdentifier", ErrNonAlphanumeric, "Previous®Message Iden")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("PreviousMessageIdentifier", ErrNonAlphanumeric, "Previous®Message Iden")).Error())
}

// TestPreviousMessageIdentifierTagError validates a PreviousMessageIdentifier tag
func TestPreviousMessageIdentifierTagError(t *testing.T) {
	pmi := mockPreviousMessageIdentifier()
	pmi.tag = "{9999}"

	require.EqualError(t, pmi.Validate(), fieldError("tag", ErrValidTagForType, pmi.tag).Error())
}

// TestStringPreviousMessageIdentifierVariableLength parses using variable length
func TestStringPreviousMessageIdentifierVariableLength(t *testing.T) {
	var line = "{3500}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePreviousMessageIdentifier()
	require.Nil(t, err)

	line = "{3500}                      NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePreviousMessageIdentifier()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3500}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePreviousMessageIdentifier()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3500}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePreviousMessageIdentifier()
	require.Equal(t, err, nil)
}

// TestStringPreviousMessageIdentifierOptions validates string() with options
func TestStringPreviousMessageIdentifierOptions(t *testing.T) {
	var line = "{3500}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePreviousMessageIdentifier()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.PreviousMessageIdentifier.String()
	require.Equal(t, str, "{3500}                      ")

	str = r.currentFEDWireMessage.PreviousMessageIdentifier.String(true)
	require.Equal(t, str, "{3500}*")
}
