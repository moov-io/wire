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
	expected := fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, creditDD.DrawdownCreditAccountNumber).Error()

	require.EqualError(t, creditDD.Validate(), expected)
}

// TestAccountCreditedDrawdownNumberRequired validates AccountCreditedDrawdown is required
func TestDrawdownCreditAccountNumberRequired(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = ""
	expected := fieldError("DrawdownCreditAccountNumber", ErrFieldRequired).Error()

	require.EqualError(t, creditDD.Validate(), expected)
}

// TestParseAccountCreditedDrawdownWrongLength parses a wrong AccountCreditedDrawdown record length
func TestParseAccountCreditedDrawdownWrongLength(t *testing.T) {
	var line = "{5400}12345678"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountCreditedDrawdown()

	expected := r.parseError(fieldError("DrawdownCreditAccountNumber", ErrValidLengthSize)).Error()
	require.EqualError(t, err, expected)
}

// TestParseAccountCreditedDrawdownReaderParseError parses a wrong AccountCreditedDrawdown reader parse error
func TestParseAccountCreditedDrawdownReaderParseError(t *testing.T) {
	var line = "{5400}12345678Z"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountCreditedDrawdown()

	expected := r.parseError(fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, "12345678Z")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("DrawdownCreditAccountNumber", ErrNonNumeric, "12345678Z")).Error()
	require.EqualError(t, err, expected)
}

// TestAccountCreditedDrawdownTagError validates AccountCreditedDrawdown tag
func TestAccountCreditedDrawdownTagError(t *testing.T) {
	creditDD := mockAccountCreditedDrawdown()
	creditDD.tag = "{9999}"

	err := creditDD.Validate()

	expected := fieldError("tag", ErrValidTagForType, creditDD.tag).Error()
	require.EqualError(t, err, expected)
}

// TestStringAccountCreditedDrawdownVariableLength parses using variable length
func TestStringAccountCreditedDrawdownVariableLength(t *testing.T) {
	var line = "{5400}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountCreditedDrawdown()
	expected := r.parseError(NewTagMinLengthErr(7, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{5400}1234567890123"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountCreditedDrawdown()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{5400} *"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountCreditedDrawdown()
	expected = r.parseError(fieldError("DrawdownCreditAccountNumber", ErrFieldRequired)).Error()
	require.EqualError(t, err, expected)

	line = "{5400}1*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountCreditedDrawdown()
	require.Equal(t, err, nil)
}

// TestStringAccountCreditedDrawdownOptions validates string() with options
func TestStringAccountCreditedDrawdownOptions(t *testing.T) {
	var line = "{5400}1*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountCreditedDrawdown()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.AccountCreditedDrawdown.String()
	require.Equal(t, str, "{5400}1        ")

	str = r.currentFEDWireMessage.AccountCreditedDrawdown.String(true)
	require.Equal(t, str, "{5400}1*")
}
