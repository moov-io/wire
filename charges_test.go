package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, c.Validate(), "mockCharges does not validate and will break other tests")
}

// TestChargeDetailsValid validates ChargeDetails is valid
func TestPaymentNotificationIndicatorValid(t *testing.T) {
	c := mockCharges()
	c.ChargeDetails = "F"

	err := c.Validate()

	require.EqualError(t, err, fieldError("ChargeDetails", ErrChargeDetails, c.ChargeDetails).Error())
}

func TestChargesCrash(t *testing.T) {
	c := &Charges{}
	c.Parse("{3700}") // invalid, caused a fuzz crash

	require.Empty(t, c.tag)
	require.Empty(t, c.ChargeDetails)
}

// TestSenderChargesOneAlphaNumeric validates SenderChargesOne is alphanumeric
func TestSenderChargesOneAlphaNumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesOne = "®"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesOne", ErrSendersCharges, c.SendersChargesOne).Error())
}

// TestSenderChargesTwoAlphaNumeric validates SenderChargesTwo is alphanumeric
func TestSenderChargesTwoAlphaNumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesTwo = "®"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesTwo", ErrSendersCharges, c.SendersChargesTwo).Error())
}

// TestSenderChargesThreeAlphaNumeric validates SenderChargesThree is alphanumeric
func TestSenderChargesThreeAlphaNumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesThree = "®"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesThree", ErrSendersCharges, c.SendersChargesThree).Error())
}

// TestSenderChargesFourAlphaNumeric validates SenderChargesFour is alphanumeric
func TestSenderChargesFourAlphaNumeric(t *testing.T) {
	c := mockCharges()
	c.SendersChargesFour = "®"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesFour", ErrSendersCharges, c.SendersChargesFour).Error())
}

// TestSendersChargesOneFormat validates SendersChargesOne begins with a valid currency code
func TestSendersChargesOneFormat(t *testing.T) {
	c := mockCharges()
	c.SendersChargesOne = "100,00"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesOne", ErrSendersCharges, c.SendersChargesOne).Error())
}

// TestSendersChargesTwoFormat validates SendersChargesTwo begins with a valid currency code
func TestSendersChargesTwoFormat(t *testing.T) {
	c := mockCharges()
	c.SendersChargesTwo = "100,00"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesTwo", ErrSendersCharges, c.SendersChargesTwo).Error())
}

// TestSendersChargesThreeFormat validates SendersChargesThree begins with a valid currency code
func TestSendersChargesThreeFormat(t *testing.T) {
	c := mockCharges()
	c.SendersChargesThree = "100,00"

	require.EqualError(t, c.Validate(), fieldError("SendersChargesThree", ErrSendersCharges, c.SendersChargesThree).Error())
}

// TestSendersChargesFourFormat validates SendersChargesFour begins with a valid currency code
func TestSendersChargesFourFormat(t *testing.T) {
	c := mockCharges()
	c.SendersChargesFour = "100,00"
	require.EqualError(t, c.Validate(), fieldError("SendersChargesFour", ErrSendersCharges, c.SendersChargesFour).Error())
}