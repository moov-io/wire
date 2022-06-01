package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, br.Validate(), "mockBeneficiaryReference does not validate and will break other tests")
}

// TestBeneficiaryReferenceAlphaNumeric validates BeneficiaryReference is alphanumeric
func TestBeneficiaryReferenceAlphaNumeric(t *testing.T) {
	br := mockBeneficiaryReference()
	br.BeneficiaryReference = "®"

	err := br.Validate()

	require.EqualError(t, err, fieldError("BeneficiaryReference", ErrNonAlphanumeric, br.BeneficiaryReference).Error())
}

// TestParseBeneficiaryReferenceWrongLength parses a wrong BeneficiaryReference record length
func TestParseBeneficiaryReferenceWrongLength(t *testing.T) {
	var line = "{4320}Reference      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()

	require.EqualError(t, err, r.parseError(fieldError("BeneficiaryReference", ErrValidLengthSize)).Error())
}

// TestParseBeneficiaryReferenceReaderParseError parses a wrong BeneficiaryReference reader parse error
func TestParseBeneficiaryReferenceReaderParseError(t *testing.T) {
	var line = "{4320}Reference®     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()

	expected := r.parseError(fieldError("BeneficiaryReference", ErrNonAlphanumeric, "Reference®")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("BeneficiaryReference", ErrNonAlphanumeric, "Reference®")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryReferenceTagError validates a BeneficiaryReference tag
func TestBeneficiaryReferenceTagError(t *testing.T) {
	br := mockBeneficiaryReference()
	br.tag = "{9999}"

	err := br.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, br.tag).Error())
}

// TestStringBeneficiaryReferenceVariableLength parses using variable length
func TestStringBeneficiaryReferenceVariableLength(t *testing.T) {
	var line = "{4320}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()
	require.Nil(t, err)

	line = "{4320}Reference       NN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryReference()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{4320}**"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryReference()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{4320}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryReference()
	require.Equal(t, err, nil)
}

// TestStringBeneficiaryReferenceOptions validates string() with options
func TestStringBeneficiaryReferenceOptions(t *testing.T) {
	var line = "{4320}Reference*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryReference()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.BeneficiaryReference.String()
	require.Equal(t, str, "{4320}Reference       ")

	str = r.currentFEDWireMessage.BeneficiaryReference.String(true)
	require.Equal(t, str, "{4320}Reference*")
}

