package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockMessageDisposition creates a MessageDisposition
func mockMessageDisposition() *MessageDisposition {
	md := NewMessageDisposition()
	md.FormatVersion = FormatVersion
	md.TestProductionCode = EnvironmentProduction
	md.MessageDuplicationCode = MessageDuplicationOriginal
	md.MessageStatusIndicator = "2"
	return md
}

// TestMockMessageDisposition validates mockMessageDisposition
func TestMockMessageDisposition(t *testing.T) {
	md := mockMessageDisposition()

	require.NoError(t, md.Validate(), "mockMessageDisposition does not validate and will break other tests")
}

// TestParseMessageDisposition parses a known MessageDisposition record string
func TestParseMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseMessageDisposition())

	record := r.currentFEDWireMessage.MessageDisposition
	require.Equal(t, "30", record.FormatVersion)
	require.Equal(t, "P", record.TestProductionCode)
	require.Empty(t, record.MessageDuplicationCode)
	require.Equal(t, "2", record.MessageStatusIndicator)
}

// TestWriteMessageDisposition writes a MessageDisposition record string
func TestWriteMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseMessageDisposition())

	record := r.currentFEDWireMessage.MessageDisposition
	require.Equal(t, line, record.String())
}

// TestMessageDispositionTagError validates a MessageDisposition tag
func TestMessageDispositionTagError(t *testing.T) {
	md := mockMessageDisposition()
	md.tag = "{9999}"

	require.EqualError(t, md.Validate(), fieldError("tag", ErrValidTagForType, md.tag).Error())
}

// TestStringMessageDispositionVariableLength parses using variable length
func TestStringMessageDispositionVariableLength(t *testing.T) {
	var line = "{1100}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseMessageDisposition()
	require.Nil(t, err)

	line = "{1100}     NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseMessageDisposition()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{1100}*******"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseMessageDisposition()
	require.ErrorContains(t, err, ErrValidLength.Error())

	line = "{1100}     *"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseMessageDisposition()
	require.Equal(t, err, nil)
}

// TestStringMessageDispositionOptions validates Format() formatted according to the FormatOptions
func TestStringMessageDispositionOptions(t *testing.T) {
	var line = "{1100}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseMessageDisposition()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.MessageDisposition
	require.Equal(t, record.String(), "{1100}     ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{1100}     ")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
