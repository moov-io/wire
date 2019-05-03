package wire

import (
	"log"
	"strings"
	"testing"
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
	if err := md.Validate(); err != nil {
		t.Error("mockMessageDisposition does not validate and will break other tests")
	}
}

// TestParseMessageDisposition parses a known MessageDisposition  record string
func TestParseMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	md := mockMessageDisposition()
	fwm.SetMessageDisposition(md)
	err := r.parseMessageDisposition()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.MessageDisposition

	if record.FormatVersion != "30" {
		t.Errorf("FormatVersion Expected '30' got: %v", record.FormatVersion)
	}
	if record.TestProductionCode != "P" {
		t.Errorf("TestProductionCode Expected 'P' got: %v", record.TestProductionCode)
	}
	if record.MessageDuplicationCode != "" {
		t.Errorf("MessageDuplicationCode Expected '' got: %v", record.MessageDuplicationCode)
	}
	if record.MessageStatusIndicator != "2" {
		t.Errorf("MessageStatusIndicator Expected '2' got: %v", record.MessageStatusIndicator)
	}
}

// TestWriteMessageDisposition writes a MessageDisposition record string
func TestWriteMessageDisposition(t *testing.T) {
	var line = "{1100}30P 2"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	md := mockMessageDisposition()
	fwm.SetMessageDisposition(md)
	err := r.parseMessageDisposition()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.MessageDisposition
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}
