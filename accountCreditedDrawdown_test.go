package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockAccountCreditedDrawdown creates a AccountCreditedDrawdown
func mockAccountCreditedDrawdown() *AccountCreditedDrawdown {
	creditDD := NewAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "123456789"
	return creditDD
}

// TestMockAccountCreditedDrawdown validates mockAccountCreditedDrawdown
func TestMockAccountCreditedDrawdown(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()

	require.NoError(t, creditDD.Validate(), "mockAccountCreditedDrawdown does not validate and will break other tests")
}

// TestAccountCreditedDrawDownNumberAlphaNumeric validates AccountCreditedDrawdown is alphanumeric
func TestDrawdownCreditAccountNumberAlphaNumeric(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "®"

	err := creditDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, creditDD.DrawdownCreditAccountNumber).Error()
	require.Equal(t, expected, err.Error())
}

// TestAccountCreditedDrawdownNumberRequired validates AccountCreditedDrawdown is required
func TestDrawdownCreditAccountNumberRequired(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = ""

	err := creditDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("DrawdownCreditAccountNumber", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

// TestParseAccountCreditedDrawdownWrongLength parses a wrong AccountCreditedDrawdown record length
func TestParseAccountCreditedDrawdownWrongLength(t *testing.T) {
	var line = "{5400}12345678"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountCreditedDrawdown()

	require.NotNil(t, err)
	expected := r.parseError(NewTagWrongLengthErr(15, len(r.line))).Error()
	require.Equal(t, expected, err.Error())
}

// TestParseAccountCreditedDrawdownReaderParseError parses a wrong AccountCreditedDrawdown reader parse error
func TestParseAccountCreditedDrawdownReaderParseError(t *testing.T) {
	var line = "{5400}12345678Z"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	crediDD := mockAccountCreditedDrawdown()
	fwm.SetAccountCreditedDrawdown(crediDD)

	err := r.parseAccountCreditedDrawdown()

	require.NotNil(t, err)
	expected := r.parseError(fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, "12345678Z")).Error()
	require.Equal(t, expected, err.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	expected = r.parseError(fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, "12345678Z")).Error()
	require.Equal(t, expected, err.Error())
}

// TestAccountCreditedDrawdownTagError validates AccountCreditedDrawdown tag
func TestAccountCreditedDrawdownTagError(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.tag = "{9999}"

	err := creditDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("tag", ErrValidTagForType, creditDD.tag).Error()
	require.Equal(t, expected, err.Error())
}
