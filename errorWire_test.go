package wire

import (
	"log"
	"strings"
	"testing"
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
	if err := ew.Validate(); err != nil {
		t.Error("mockErrorWire does not validate and will break other tests")
	}
}

// TestParseErrorWire parses a known ErrorWire  record string
func TestParseErrorWire(t *testing.T) {
	var line = "{1130}1XYZData Error                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ew := mockErrorWire()
	fwm.SetErrorWire(ew)
	err := r.parseErrorWire()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.ErrorWire

	if record.ErrorCategory != "1" {
		t.Errorf("ErrorCategory Expected '1' got: %v", record.ErrorCategory)
	}
	if record.ErrorCode != "XYZ" {
		t.Errorf("ErrorCode  Expected 'XYZ' got: %v", record.ErrorCode)
	}
	if record.ErrorDescription != "Data Error" {
		t.Errorf("ErrorDescription Expected 'Data Error' got: %v", record.ErrorDescription)
	}
}

// TestWriteErrorWire writes a ErrorWire record string
func TestWriteErrorWire(t *testing.T) {
	var line = "{1130}1XYZData Error                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	ew := mockErrorWire()
	fwm.SetErrorWire(ew)
	err := r.parseErrorWire()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.ErrorWire
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}
