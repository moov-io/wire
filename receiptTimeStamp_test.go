package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockReceiptTimeStamp creates a ReceiptTimeStamp
func mockReceiptTimeStamp() *ReceiptTimeStamp {
	rts := NewReceiptTimeStamp()
	rts.ReceiptDate = "0502"
	rts.ReceiptTime = "1230"
	rts.ReceiptApplicationIdentification = "A123"
	return rts
}

// TestMockReceiptTimeStamp validates mockReceiptTimeStamp
func TestMockReceiptTimeStamp(t *testing.T) {
	rts := mockReceiptTimeStamp()

	require.NoError(t, rts.Validate(), "mockReceiptTimeStamp does not validate and will break other tests")
}

// TestParseReceiptTimeStamp parses a known ReceiptTimeStamp  record string
func TestParseReceiptTimeStamp(t *testing.T) {
	var line = "{1110}05021230A123"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseReceiptTimeStamp())

	record := r.currentFEDWireMessage.ReceiptTimeStamp
	require.Equal(t, "0502", record.ReceiptDate)
	require.Equal(t, "1230", record.ReceiptTime)
	require.Equal(t, "A123", record.ReceiptApplicationIdentification)
}

// TestWriteReceiptTimeStamp writes a ReceiptTimeStamp record string
func TestWriteReceiptTimeStamp(t *testing.T) {
	var line = "{1110}05021230A123"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseReceiptTimeStamp())

	record := r.currentFEDWireMessage.ReceiptTimeStamp
	require.Equal(t, line, record.String())
}

// TestReceiptTimeStampTagError validates a ReceiptTimeStamp tag
func TestReceiptTimeStampTagError(t *testing.T) {
	rts := mockReceiptTimeStamp()
	rts.tag = "{9999}"

	require.EqualError(t, rts.Validate(), fieldError("tag", ErrValidTagForType, rts.tag).Error())
}

// TestStringReceiptTimeStampVariableLength parses using variable length
func TestStringReceiptTimeStampVariableLength(t *testing.T) {
	var line = "{1110}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseReceiptTimeStamp()
	require.Nil(t, err)

	line = "{1110}            NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseReceiptTimeStamp()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{1110}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseReceiptTimeStamp()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{1110}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseReceiptTimeStamp()
	require.Equal(t, err, nil)
}

// TestStringReceiptTimeStampOptions validates string() with options
func TestStringReceiptTimeStampOptions(t *testing.T) {
	var line = "{1110}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseReceiptTimeStamp()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.ReceiptTimeStamp.String()
	require.Equal(t, str, "{1110}            ")

	str = r.currentFEDWireMessage.ReceiptTimeStamp.String(true)
	require.Equal(t, str, "{1110}*")
}
