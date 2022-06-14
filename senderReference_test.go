package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, sr.Validate(), "mockSenderReference does not validate and will break other tests")
}

// TestSenderReferenceAlphaNumeric validates SenderReference is alphanumeric
func TestSenderReferenceAlphaNumeric(t *testing.T) {
	sr := mockSenderReference()
	sr.SenderReference = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SenderReference", ErrNonAlphanumeric, sr.SenderReference).Error())
}

// TestParseSenderReferenceWrongLength parses a wrong SenderReference record length
func TestParseSenderReferenceWrongLength(t *testing.T) {
	var line = "{3320}Se"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderReference()

	require.EqualError(t, err, r.parseError(fieldError("SenderReference", ErrValidLength)).Error())
}

// TestParseSenderReferenceReaderParseError parses a wrong SenderReference reader parse error
func TestParseSenderReferenceReaderParseError(t *testing.T) {
	var line = "{3320}Sender®Referenc"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderReference()

	require.EqualError(t, err, r.parseError(fieldError("SenderReference", ErrNonAlphanumeric, "Sender®Referenc")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SenderReference", ErrNonAlphanumeric, "Sender®Referenc")).Error())
}

// TestSenderReferenceTagError validates a SenderReference tag
func TestSenderReferenceTagError(t *testing.T) {
	sr := mockSenderReference()
	sr.tag = "{9999}"

	require.EqualError(t, sr.Validate(), fieldError("tag", ErrValidTagForType, sr.tag).Error())
}

// TestStringSenderReferenceVariableLength parses using variable length
func TestStringSenderReferenceVariableLength(t *testing.T) {
	var line = "{3320}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderReference()
	require.Nil(t, err)

	line = "{3320}                NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderReference()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3320}***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderReference()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3320}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderReference()
	require.Equal(t, err, nil)
}

// TestStringSenderReferenceOptions validates string() with options
func TestStringSenderReferenceOptions(t *testing.T) {
	var line = "{3320}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderReference()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.SenderReference.String()
	require.Equal(t, str, "{3320}                ")

	str = r.currentFEDWireMessage.SenderReference.String(true)
	require.Equal(t, str, "{3320}*")
}
