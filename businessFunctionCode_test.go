package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockBusinessFunctionCode creates a BusinessFunctionCode
func mockBusinessFunctionCode() *BusinessFunctionCode {
	bfc := NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "   "
	return bfc
}

// TestMockBusinessFunctionCode validates mockBusinessFunctionCode
func TestMockBusinessFunctionCode(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	if err := bfc.Validate(); err != nil {
		t.Error("mockBusinessFunctionCode does not validate and will break other tests")
	}
}

// TestBusinessFunctionCodeValid validates BusinessFunctionCode
func TestBusinessFunctionCodeValid(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = "ZZZ"
	if err := bfc.Validate(); err != nil {
		if !base.Match(err, ErrBusinessFunctionCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBusinessFunctionCodeRequired validates BusinessFunctionCode is required
func TestBusinessFunctionCodeRequired(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = ""
	if err := bfc.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBusinessFunctionCodeWrongLength parses a wrong BusinessFunctionCode record length
func TestParseBusinessFunctionCodeWrongLength(t *testing.T) {
	var line = "{3600}CT"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)
	if err := r.parseBusinessFunctionCode(); err != nil {
		if !base.Match(err, NewTagWrongLengthErr(12, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseBusinessFunctionCodeReaderParseError parses a wrong BusinessFunctionCode reader parse error
func TestParseBusinessFunctionCodeReaderParseError(t *testing.T) {
	var line = "{3600}CTAXXY"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bfc := mockBusinessFunctionCode()
	fwm.SetBusinessFunctionCode(bfc)
	if err := r.parseBusinessFunctionCode(); err != nil {
		if !base.Match(err, ErrBusinessFunctionCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err := r.Read()
	if err != nil {
		if !base.Has(err, ErrBusinessFunctionCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
