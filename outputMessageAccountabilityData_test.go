package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockOutputMessageAccountabilityData creates a mockOutputMessageAccountabilityData
func mockOutputMessageAccountabilityData() *OutputMessageAccountabilityData {
	omad := NewOutputMessageAccountabilityData(false)
	omad.OutputCycleDate = "20190502"
	omad.OutputDestinationID = "Source08"
	omad.OutputSequenceNumber = "000001"
	omad.OutputDate = "0502"
	omad.OutputTime = "1230"
	omad.OutputFRBApplicationIdentification = "B123"
	return omad
}

// TestMockOutputMessageAccountabilityData validates mockOutputMessageAccountabilityData
func TestMockOutputMessageAccountabilityData(t *testing.T) {
	omad := mockOutputMessageAccountabilityData()

	require.NoError(t, omad.Validate(), "mockOutputMessageAccountabilityData does not validate and will break other tests")
}

// TestParseOutputMessageAccountabilityData parses a known OutputMessageAccountabilityData  record string
func TestParseOutputMessageAccountabilityData(t *testing.T) {
	var line = "{1120}20190502Source0800000105021230B123"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseOutputMessageAccountabilityData())

	record := r.currentFEDWireMessage.OutputMessageAccountabilityData
	require.Equal(t, "20190502", record.OutputCycleDate)
	require.Equal(t, "Source08", record.OutputDestinationID)
	require.Equal(t, "000001", record.OutputSequenceNumber)
	require.Equal(t, "0502", record.OutputDate)
	require.Equal(t, "1230", record.OutputTime)
	require.Equal(t, "B123", record.OutputFRBApplicationIdentification)
}

// TestWriteOutputMessageAccountabilityData writes a OutputMessageAccountabilityData record string
func TestWriteOutputMessageAccountabilityData(t *testing.T) {
	var line = "{1120}20190502Source0800000105021230B123"
	r := NewReader(strings.NewReader(line))
	r.line = line

	require.NoError(t, r.parseOutputMessageAccountabilityData())

	record := r.currentFEDWireMessage.OutputMessageAccountabilityData
	require.Equal(t, line, record.String())
}

// TestOutputMessageAccountabilityDataTagError validates a OutputMessageAccountabilityData tag
func TestOutputMessageAccountabilityDataTagError(t *testing.T) {
	omad := mockOutputMessageAccountabilityData()
	omad.tag = "{9999}"

	require.EqualError(t, omad.Validate(), fieldError("tag", ErrValidTagForType, omad.tag).Error())
}
