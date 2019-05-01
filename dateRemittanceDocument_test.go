package wire

import (
	"github.com/moov-io/base"
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

// TestDateRemittanceDocumentAmountRequired validates DateRemittanceDocument Amount is required
func TestDateRemittanceDocumentAmountRequired(t *testing.T) {
	drd := mockDateRemittanceDocument()
	drd.DateRemittanceDocument = ""
	if err := drd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}