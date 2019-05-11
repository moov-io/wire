package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockCharges creates a Charges
func mockCharges() *Charges {
	c := NewCharges()
	c.ChargeDetails = "B"
	c.SendersChargesOne = "USD0,99"
	c.SendersChargesTwo = "USD2,99"
	c.SendersChargesThree = "USD3,99"
	c.SendersChargesFour = "USD1,00"
	return c
}

// TestMockCharges validates mockCharges
func TestMockCharges(t *testing.T) {
	c := mockCharges()
	if err := c.Validate(); err != nil {
		t.Error("mockCharges does not validate and will break other tests")
	}
}

// TestParseChargesWrongLength parses a wrong Charges record length
func TestParseChargesWrongLength(t *testing.T) {
	var line = "{3700}BUSD0,99        USD2,99        USD3,99        USD1,00      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	c := mockCharges()
	fwm.SetCharges(c)
	if err := r.parseCharges(); err != nil {
		if !base.Match(err, NewTagWrongLengthErr(67, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseChargesReaderParseError parses a wrong Charges reader parse error
func TestParseChargesReaderParseError(t *testing.T) {
	var line = "{3700}BUSD0,99        USD2,99        USD3Z99        USD1,00        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	c := mockCharges()
	fwm.SetCharges(c)
	if err := r.parseCharges(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err := r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
