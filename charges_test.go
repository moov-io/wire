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

// TestChargesChargeDetails validates Charges SendersChargeDetails
func TestChargesChargeDetails(t *testing.T) {
	c := mockCharges()
	c.ChargeDetails = "Z"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrChargeDetails) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesSendersChargesOne validates Charges SendersChargesOne currency code
func TestChargesSendersChargesOne(t *testing.T) {
	c := mockCharges()
	c.SendersChargesOne = "ZZZ0,99"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesCurrencyCodeTwo validates Charges SendersChargesTwo currency code
func TestChargesCurrencyCodeTwo(t *testing.T) {
	c := mockCharges()
	c.SendersChargesTwo = "ZZZ0,99"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesCurrencyCodeThree validates Charges SendersChargesThree currency code
func TestChargesCurrencyCodeThree(t *testing.T) {
	c := mockCharges()
	c.SendersChargesThree = "ZZZ0,99"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesCurrencyCodeFour validates Charges SendersChargesFour currency code
func TestChargesCurrencyCodeFour(t *testing.T) {
	c := mockCharges()
	c.SendersChargesFour = "ZZZ0,99"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonCurrencyCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesSendersChargesOneAlphanumeric validates Charges SendersChargesOneAlphanumeric
func TestChargesCurrencyCodeOneAlphanumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesOne = "®"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesSendersChargesTwoAlphanumeric validates Charges SendersChargesTwoAlphanumeric
func TestChargesCurrencyCodeTwoAlphanumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesTwo = "®"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesSendersChargesThreeAlphanumeric validates Charges SendersChargesThreeAlphanumeric
func TestChargesCurrencyCodeThreeAlphanumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesThree = "®"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesSendersChargesFourAlphanumeric validates Charges SendersChargesFourAlphanumeric
func TestChargesCurrencyCodeFourAlphanumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesFour = "®"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesAmountOne validates Charges SendersChargesOne amount
func TestChargesAmountOne(t *testing.T) {
	c := mockCharges()
	c.SendersChargesOne = "USD0,Z9"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesAmountOne validates Charges SendersChargesTwo amount
func TestChargesAmountTwo(t *testing.T) {
	c := mockCharges()
	c.SendersChargesTwo = "USD0,Z9"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestChargesAmountFour validates Charges SendersChargesFour amount
func TestChargesAmountFour(t *testing.T) {
	c := mockCharges()
	c.SendersChargesFour = "USD0,Z9"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrNonAmount) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
