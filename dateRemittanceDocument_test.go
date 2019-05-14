package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
	"time"
)

// DateRemittanceDocument creates a DateRemittanceDocument
func mockDateRemittanceDocument() *DateRemittanceDocument {
	drd := NewDateRemittanceDocument()
	drd.DateRemittanceDocument = time.Now().Format("20060102")
	return drd
}

// TestMockDateRemittanceDocument validates mockDateRemittanceDocument
func TestMockDateRemittanceDocument(t *testing.T) {
	drd := mockDateRemittanceDocument()
	if err := drd.Validate(); err != nil {
		t.Error("mockDateRemittanceDocument does not validate and will break other tests")
	}
}

// TestDateRemittanceDocumentRequired validates DateRemittanceDocument DateRemittanceDocument is required
func TestDateRemittanceDocumentDateRemittanceDocumentRequired(t *testing.T) {
	drd := mockDateRemittanceDocument()
	drd.DateRemittanceDocument = ""
	if err := drd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseDateRemittanceDocumentWrongLength parses a wrong DateRemittanceDocument record length
func TestParseDateRemittanceDocumentWrongLength(t *testing.T) {
	var line = "{8650}20190509  "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(rd)
	err := r.parseDateRemittanceDocument()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(14, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseDateRemittanceDocumentReaderParseError parses a wrong DateRemittanceDocument reader parse error
func TestParseDateRemittanceDocumentReaderParseError(t *testing.T) {
	var line = "{8650}14190509"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	rd := mockDateRemittanceDocument()
	fwm.SetDateRemittanceDocument(rd)
	err := r.parseDateRemittanceDocument()
	if err != nil {
		if !base.Match(err, ErrValidDate) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrValidDate) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestDateRemittanceDocumentTagError validates a DateRemittanceDocument tag
func TestDateRemittanceDocumentTagError(t *testing.T) {
	drd := mockDateRemittanceDocument()
	drd.tag = "{9999}"
	if err := drd.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
