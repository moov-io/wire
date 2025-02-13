package wire

import (
	"errors"
	"strings"
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

// TestStringChargesVariableLength parses using variable length
func TestStringChargesVariableLength(t *testing.T) {
	var line = "{3700}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCharges()
	expected := r.parseError(NewTagMinLengthErr(7, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{3700}B                                                            NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseCharges()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{3700}B******"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseCharges()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{3700}B*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseCharges()
	require.NoError(t, err)
}

// TestStringChargesOptions validates Format() formatted according to the FormatOptions
func TestStringChargesOptions(t *testing.T) {
	var line = "{3700}B*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseCharges()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.Charges
	require.Equal(t, "{3700}B               *               *               *               *", record.String())
	require.Equal(t, "{3700}B*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
