package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockReceiverDepositoryInstitution creates a ReceiverDepositoryInstitution
func mockReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	rdi := NewReceiverDepositoryInstitution(false)
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	return rdi
}

// TestMockReceiverDepositoryInstitution validates mockReceiverDepositoryInstitution
func TestMockReceiverDepositoryInstitution(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()

	require.NoError(t, rdi.Validate(), "mockReceiverDepositoryInstitution does not validate and will break other tests")
}

// TestReceiverABANumberAlphaNumeric validates ReceiverDepositoryInstitution ReceiverABANumber is alphanumeric
func TestReceiverABANumberAlphaNumeric(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "®"

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("ReceiverABANumber", ErrNonNumeric, rdi.ReceiverABANumber).Error())
}

// TestReceiverShortNameAlphaNumeric validates ReceiverDepositoryInstitution ReceiverShortName is alphanumeric
func TestReceiverShortNameAlphaNumeric(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverShortName = "®"

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("ReceiverShortName", ErrNonAlphanumeric, rdi.ReceiverShortName).Error())
}

// TestReceiverABANumberRequired validates ReceiverDepositoryInstitution ReceiverABANumber is required
func TestReceiverABANumberRequired(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = ""

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("ReceiverABANumber", ErrFieldRequired, rdi.ReceiverABANumber).Error())
}

// TestReceiverShortNameRequired validates ReceiverDepositoryInstitution ReceiverShortName is required
func TestReceiverShortNameRequired(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.ReceiverShortName = ""

	err := rdi.Validate()

	require.EqualError(t, err, fieldError("ReceiverShortName", ErrFieldRequired, rdi.ReceiverShortName).Error())
}

// TestParseReceiverWrongLength parses a wrong Receiver record length
func TestParseReceiverWrongLength(t *testing.T) {
	var line = "{3400}00"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseReceiverDepositoryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("ReceiverABANumber", ErrValidLengthSize)).Error())
}

// TestParseReceiverReaderParseError parses a wrong Receiver reader parse error
func TestParseReceiverReaderParseError(t *testing.T) {
	var line = "{3400}2313Z0104Citadel           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseReceiverDepositoryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("ReceiverABANumber", ErrNonNumeric, "2313Z0104")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("ReceiverABANumber", ErrNonNumeric, "2313Z0104")).Error())
}

// TestReceiverDepositoryInstitutionTagError validates a ReceiverDepositoryInstitution tag
func TestReceiverDepositoryInstitutionTagError(t *testing.T) {
	rdi := mockReceiverDepositoryInstitution()
	rdi.tag = "{9999}"

	require.EqualError(t, rdi.Validate(), fieldError("tag", ErrValidTagForType, rdi.tag).Error())
}
