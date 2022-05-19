package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// mockSenderSupplied creates a SenderSupplied
func mockSenderSupplied() *SenderSupplied {
	ss := NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = MessageDuplicationOriginal
	return ss
}

// TestMockSenderSupplied validates mockSenderSupplied
func TestMockSenderSupplied(t *testing.T) {
	ss := mockSenderSupplied()

	require.NoError(t, ss.Validate(), "mockSenderSupplied does not validate and will break other tests")
}

// TestSenderSuppliedUserRequestCorrelationAlphaNumeric validates SenderSupplied UserRequestCorrelation is alphanumeric
func TestSenderSuppliedUserRequestCorrelationAlphaNumeric(t *testing.T) {
	ss := mockSenderSupplied()
	ss.UserRequestCorrelation = "®"

	err := ss.Validate()

	require.EqualError(t, err, fieldError("UserRequestCorrelation", ErrNonAlphanumeric, ss.UserRequestCorrelation).Error())
}

// TestSenderSuppliedFormatVersionValid validates SenderSupplied FormatVersion
func TestSenderSuppliedFormatVersionValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.FormatVersion = "55"

	err := ss.Validate()

	if !base.Match(err, ErrFormatVersion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSenderSuppliedProductionCodeValid validates SenderSupplied ProductionCode
func TestSenderSuppliedProductionCodeValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.TestProductionCode = "Z"

	err := ss.Validate()

	if !base.Match(err, ErrTestProductionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSenderSuppliedMessageDuplicationCodeValid validates SenderSupplied MessageDuplicationCode
func TestSenderSuppliedMessageDuplicationCodeValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.MessageDuplicationCode = "Z"

	err := ss.Validate()

	if !base.Match(err, ErrMessageDuplicationCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSenderSuppliedUserRequestCorrelationRequired validates SenderSupplied UserRequestCorrelation is required
func TestSenderSuppliedUserRequestCorrelationRequired(t *testing.T) {
	ss := mockSenderSupplied()
	ss.UserRequestCorrelation = ""

	err := ss.Validate()

	require.EqualError(t, err, fieldError("UserRequestCorrelation", ErrFieldRequired, ss.UserRequestCorrelation).Error())
}

// TestParseSenderSuppliedWrongLength parses a wrong SenderSupplied record length
func TestParseSenderSuppliedWrongLength(t *testing.T) {
	var line = "{1500}30P"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderSupplied()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(10, len(r.line))).Error())
}

// TestParseSenderSuppliedReaderParseError parses a wrong SenderSupplied reader parse error
func TestParseSenderSuppliedReaderParseError(t *testing.T) {
	var line = "{1500}25User ReqP "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderSupplied()

	require.EqualError(t, err, r.parseError(fieldError("FormatVersion", ErrFormatVersion, "25")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("FormatVersion", ErrFormatVersion, "25")).Error())
}

// TestSenderSuppliedTagError validates a SenderSupplied tag
func TestSenderSuppliedTagError(t *testing.T) {
	ss := mockSenderSupplied()
	ss.tag = "{9999}"

	require.EqualError(t, ss.Validate(), fieldError("tag", ErrValidTagForType, ss.tag).Error())
}
