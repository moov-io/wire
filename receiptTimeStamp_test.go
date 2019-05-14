package wire

import (
	"github.com/moov-io/base"
	"log"
	"strings"
	"testing"
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
	if err := rts.Validate(); err != nil {
		t.Error("mockReceiptTimeStamp does not validate and will break other tests")
	}
}

// TestParseReceiptTimeStamp parses a known ReceiptTimeStamp  record string
func TestParseReceiptTimeStamp(t *testing.T) {
	var line = "{1110}05021230A123"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rts := mockReceiptTimeStamp()
	fwm.SetReceiptTimeStamp(rts)
	err := r.parseReceiptTimeStamp()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.ReceiptTimeStamp

	if record.ReceiptDate != "0502" {
		t.Errorf("ReceiptDate Expected '0502' got: %v", record.ReceiptDate)
	}
	if record.ReceiptTime != "1230" {
		t.Errorf("ReceiptTime Expected '1230' got: %v", record.ReceiptTime)
	}
	if record.ReceiptApplicationIdentification != "A123" {
		t.Errorf("ReceiptApplicationIdentification Expected 'A123' got: %v", record.ReceiptApplicationIdentification)
	}
}

// TestWriteReceiptTimeStamp writes a ReceiptTimeStamp record string
func TestWriteReceiptTimeStamp(t *testing.T) {
	var line = "{1110}05021230A123"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rts := mockReceiptTimeStamp()
	fwm.SetReceiptTimeStamp(rts)
	err := r.parseReceiptTimeStamp()
	if err != nil {
		t.Errorf("%T: %s", err, err)
		log.Fatal(err)
	}
	record := r.currentFEDWireMessage.ReceiptTimeStamp
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestReceiptTimeStampTagError validates a ReceiptTimeStamp tag
func TestReceiptTimeStampTagError(t *testing.T) {
	rts := mockReceiptTimeStamp()
	rts.tag = "{9999}"
	if err := rts.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
