package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockErrorWire creates a ErrorWire
func mockErrorWire() *ErrorWire {
	ew := NewErrorWire()
	ew.ErrorCategory = "E"
	ew.ErrorCode = "XYZ"
	ew.ErrorDescription = "Data Error"
	return ew
}

// TestMockErrorWire validates mockErrorWire
func TestMockErrorWire(t *testing.T) {
	ew := mockErrorWire()

	require.NoError(t, ew.Validate(), "mockErrorWire does not validate and will break other tests")
}

// TestParseErrorWire parses a known ErrorWire  record string
func TestParseErrorWire(t *testing.T) {
	var line = "{1130}1XYZData Error                         *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseErrorWire())
	record := r.currentFEDWireMessage.ErrorWire

	assert.Equal(t, "1", record.ErrorCategory)
	assert.Equal(t, "XYZ", record.ErrorCode)
	assert.Equal(t, "Data Error", record.ErrorDescription)
}

// TestWriteErrorWire writes a ErrorWire record string
func TestWriteErrorWire(t *testing.T) {
	var line = "{1130}1XYZData Error                         *"
	r := NewReader(strings.NewReader(line))
	r.line = line
	require.NoError(t, r.parseErrorWire())
	record := r.currentFEDWireMessage.ErrorWire

	assert.Equal(t, line, record.String())
}

// TestStringErrorWireAmountVariableLength parses using variable length
func TestStringErrorWireAmountVariableLength(t *testing.T) {
	var line = "{1130}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseErrorWire()
	require.Nil(t, err)

	line = "{1130}1XYZData Error                         NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseErrorWire()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{1130}1XYZData Error***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseErrorWire()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{1130}1XYZData Error*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseErrorWire()
	require.Equal(t, err, nil)
}

// TestStringErrorWireOptions validates Format() formatted according to the FormatOptions
func TestStringErrorWireOptions(t *testing.T) {
	var line = "{1130}1XYZData Error*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseErrorWire()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.ErrorWire
	require.Equal(t, record.String(), "{1130}1XYZData Error                         *")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{1130}1XYZData Error*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
