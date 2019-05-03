package wire

import (
	"log"
	"strings"
	"testing"
)

// mockOutputMessageAccountabilityData creates a mockOutputMessageAccountabilityData
func mockOutputMessageAccountabilityData() *OutputMessageAccountabilityData {
	omad := NewOutputMessageAccountabilityData()
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
	if err := omad.Validate(); err != nil {
		t.Error("mockOutputMessageAccountabilityData does not validate and will break other tests")
	}
}

// TestParseOutputMessageAccountabilityData parses a known OutputMessageAccountabilityData  record string
func TestParseOutputMessageAccountabilityData(t *testing.T) {
	var line = "{1120}20190502Source0800000105021230B123"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	omad := mockOutputMessageAccountabilityData()
	fwm.SetOutputMessageAccountabilityData(omad)
	err := r.parseOutputMessageAccountabilityData()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.OutputMessageAccountabilityData

	if record.OutputCycleDate != "20190502" {
		t.Errorf("OutputCycleDate Expected '20190502' got: %v", record.OutputCycleDate)
	}
	if record.OutputDestinationID != "Source08" {
		t.Errorf("OutputDestinationID Expected 'Source08' got: %v", record.OutputDestinationID)
	}
	if record.OutputSequenceNumber != "000001" {
		t.Errorf("OutputSequenceNumber Expected 'Source08' got: %v", record.OutputSequenceNumber)
	}
	if record.OutputDate != "0502" {
		t.Errorf("OutputDate Expected '0502' got: %v", record.OutputDate)
	}
	if record.OutputTime != "1230" {
		t.Errorf("OutputTime Expected '1230' got: %v", record.OutputTime)
	}
	if record.OutputFRBApplicationIdentification != "B123" {
		t.Errorf("OutputFRBApplicationIdentification 'B123' got: %v", record.OutputFRBApplicationIdentification)
	}
}

// TestWriteOutputMessageAccountabilityData writes a OutputMessageAccountabilityData record string
func TestWriteOutputMessageAccountabilityData(t *testing.T) {
	var line = "{1120}20190502Source0800000105021230B123"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	omad := mockOutputMessageAccountabilityData()
	fwm.SetOutputMessageAccountabilityData(omad)
	err := r.parseOutputMessageAccountabilityData()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.OutputMessageAccountabilityData
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}
