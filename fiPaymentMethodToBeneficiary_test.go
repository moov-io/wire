package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockFIPaymentMethodToBeneficiary creates a FIPaymentMethodToBeneficiary
func mockFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	return pm
}

// TestMockFIPaymentMethodToBeneficiary validates mockFIPaymentMethodToBeneficiary
func TestMockFIPaymentMethodToBeneficiary(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()

	require.NoError(t, pm.Validate(), "mockFIPaymentMethodToBeneficiary does not validate and will break other tests")
}

// TestPaymentMethodValid validates FIPaymentMethodToBeneficiary PaymentMethod
func TestPaymentMethodValid(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = ""

	err := pm.Validate()

	require.EqualError(t, err, fieldError("PaymentMethod", ErrFieldInclusion, pm.PaymentMethod).Error())
}

// TestAdditionalInformationAlphaNumeric validates FIPaymentMethodToBeneficiary AdditionalInformation is alphanumeric
func TestAdditionalInformationAlphaNumeric(t *testing.T) {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.AdditionalInformation = "®"

	err := pm.Validate()

	require.EqualError(t, err, fieldError("AdditionalInformation", ErrNonAlphanumeric, pm.AdditionalInformation).Error())
}

// TestParseFIPaymentMethodToBeneficiaryWrongLength parses a wrong FIPaymentMethodToBeneficiary record length
func TestParseFIPaymentMethodToBeneficiaryWrongLength(t *testing.T) {
	var line = "{6420}CHECKAdditional Information      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIPaymentMethodToBeneficiary()
	require.EqualError(t, err, r.parseError(fieldError("AdditionalInformation", ErrValidLength)).Error())
}

// TestParseFIPaymentMethodToBeneficiaryReaderParseError parses a wrong FIPaymentMethodToBeneficiary reader parse error
func TestParseFIPaymentMethodToBeneficiaryReaderParseError(t *testing.T) {
	var line = "{6420}CHECK®dditional Information       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIPaymentMethodToBeneficiary()

	expected := r.parseError(fieldError("AdditionalInformation", ErrNonAlphanumeric, "®dditional Information")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("AdditionalInformation", ErrNonAlphanumeric, "®dditional Information")).Error()
	require.EqualError(t, err, expected)
}

// TestFIPaymentMethodToBeneficiaryTagError validates a FIPaymentMethodToBeneficiary tag
func TestFIPaymentMethodToBeneficiaryTagError(t *testing.T) {
	pm := mockFIPaymentMethodToBeneficiary()
	pm.tag = "{9999}"

	err := pm.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, pm.tag).Error())
}

// TestStringFIPaymentMethodToBeneficiaryVariableLength parses using variable length
func TestStringFIPaymentMethodToBeneficiaryVariableLength(t *testing.T) {
	var line = "{6420}CHECK"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIPaymentMethodToBeneficiary()
	require.Nil(t, err)

	line = "{6420}CHECK                              NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIPaymentMethodToBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{6420}CHECK***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIPaymentMethodToBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{6420}CHECK*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseFIPaymentMethodToBeneficiary()
	require.Equal(t, err, nil)
}

// TestStringFIPaymentMethodToBeneficiaryOptions validates Format() formatted according to the FormatOptions
func TestStringFIPaymentMethodToBeneficiaryOptions(t *testing.T) {
	var line = "{6420}CHECK*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseFIPaymentMethodToBeneficiary()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.FIPaymentMethodToBeneficiary
	require.Equal(t, record.String(), "{6420}CHECK                              ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{6420}CHECK*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
