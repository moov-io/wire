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

	require.EqualError(t, err, r.parseError(fieldError("PreviousMessageIdentifier", ErrValidLengthSize)).Error())
}

// TestParsePreviousMessageIdentifierReaderParseError parses a wrong PreviousMessageIdentifier reader parse error
func TestParsePreviousMessageIdentifierReaderParseError(t *testing.T) {
	var line = "{3500}Previous®Message Ident"
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
