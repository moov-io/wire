package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// mockSenderDepositoryInstitution creates a SenderDepositoryInstitution
func mockSenderDepositoryInstitution() *SenderDepositoryInstitution {
	sdi := NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	return sdi
}

// TestMockSenderDepositoryInstitution validates mockSenderDepositoryInstitution
func TestMockSenderDepositoryInstitution(t *testing.T) {
	sdi := mockSenderDepositoryInstitution()

	require.NoError(t, sdi.Validate(), "mockSenderDepositoryInstitution does not validate and will break other tests")
}

// TestSenderABANumberAlphaNumeric validates SenderDepositoryInstitution SenderABANumber is alphanumeric
func TestSenderABANumberAlphaNumeric(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderABANumber = "®"

	err := rdi.Validate()

	if !base.Match(err, ErrNonNumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSenderShortNameAlphaNumeric validates SenderDepositoryInstitution SenderShortName is alphanumeric
func TestSenderShortNameAlphaNumeric(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderShortName = "®"

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("SenderShortName", ErrNonAlphanumeric, rdi.SenderShortName).Error())
}

// TestSenderABANumberRequired validates SenderDepositoryInstitution SenderABANumber is required
func TestSenderABANumberRequired(t *testing.T) {
	rdi := mockSenderDepositoryInstitution()
	rdi.SenderABANumber = ""

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("SenderABANumber", ErrFieldRequired, rdi.SenderABANumber).Error())
}

// TestParseSenderWrongLength parses a wrong Sender record length
func TestParseSenderWrongLength(t *testing.T) {
	var line = "{3100}0012"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderDepositoryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("SenderABANumber", ErrValidLength)).Error())
}

// TestParseSenderReaderParseError parses a wrong Sender reader parse error
func TestParseSenderReaderParseError(t *testing.T) {
	var line = "{3100}1210Z2882Wells Fargo NA    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderDepositoryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("SenderABANumber", ErrNonNumeric, "1210Z2882")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SenderABANumber", ErrNonNumeric, "1210Z2882")).Error())
}

// TestSenderDepositoryInstitutionTagError validates a SenderDepositoryInstitution tag
func TestSenderDepositoryInstitutionTagError(t *testing.T) {
	sdi := mockSenderDepositoryInstitution()
	sdi.tag = "{9999}"

	require.EqualError(t, sdi.Validate(), fieldError("tag", ErrValidTagForType, sdi.tag).Error())
}

// TestStringSenderDepositoryInstitutionVariableLength parses using variable length
func TestStringSenderDepositoryInstitutionVariableLength(t *testing.T) {
	var line = "{3100}1*A*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderDepositoryInstitution()
	require.Nil(t, err)

	line = "{3100}1        A                 NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderDepositoryInstitution()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3100}1*A***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderDepositoryInstitution()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{3100}1*A*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderDepositoryInstitution()
	require.Equal(t, err, nil)
}

// TestStringSenderDepositoryInstitutionOptions validates Format() formatted according to the FormatOptions
func TestStringSenderDepositoryInstitutionOptions(t *testing.T) {
	var line = "{3100}1*A*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderDepositoryInstitution()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.SenderDepositoryInstitution
	require.Equal(t, record.String(), "{3100}1        A                 ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{3100}1*A*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))

	line = "{3100}1*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderDepositoryInstitution()
	require.Equal(t, err, nil)

	record = r.currentFEDWireMessage.SenderDepositoryInstitution
	require.Equal(t, record.String(), "{3100}1                          ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{3100}1*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))

	line = "{3100}111111111*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderDepositoryInstitution()
	require.Equal(t, err, nil)

	record = r.currentFEDWireMessage.SenderDepositoryInstitution
	require.Equal(t, record.String(), "{3100}111111111                  ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{3100}111111111")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
