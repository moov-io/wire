package wire

import (
	"github.com/moov-io/base"
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

// TestChargeDetailsValid validates ChargeDetails is valid
func TestPaymentNotificationIndicatorValid(t *testing.T) {
	c := mockCharges()
	c.ChargeDetails = "F"
	if err := c.Validate(); err != nil {
		if !base.Match(err, ErrChargeDetails) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
