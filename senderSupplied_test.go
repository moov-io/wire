package wire

import (
	"errors"
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

// TestSenderSuppliedUserRequestCorrelation validates SenderSupplied UserRequestCorrelation is optional
func TestSenderSuppliedUserRequestCorrelation(t *testing.T) {
	ss := mockSenderSupplied()
	ss.UserRequestCorrelation = ""

	err := ss.Validate()
	require.NoError(t, err)

	var line = "{1500}30        T "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderSupplied()
	require.NoError(t, err)
}

// TestParseSenderSuppliedWrongLength parses a wrong SenderSupplied record length
func TestParseSenderSuppliedWrongLength(t *testing.T) {
	var line = "{1500}30P"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderSupplied()

	require.EqualError(t, err, r.parseError(NewTagMinLengthErr(11, len(r.line))).Error())
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

// TestStringSenderSuppliedVariableLength parses using variable length
func TestStringSenderSuppliedVariableLength(t *testing.T) {
	var line = "{1500}301*T "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderSupplied()
	require.ErrorContains(t, err, ErrValidLength.Error())

	line = "{1500}301       T NNN "
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderSupplied()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{1500}301*T** "
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseSenderSupplied()
	require.ErrorContains(t, err, ErrValidLength.Error())
}

// TestStringSenderSuppliedOptions validates Format() formatted according to the FormatOptions
func TestStringSenderSuppliedOptions(t *testing.T) {
	var line = "{1500}301       T "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderSupplied()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.SenderSupplied
	require.Equal(t, record.String(), "{1500}301       T ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{1500}301       T ")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
