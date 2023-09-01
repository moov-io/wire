package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// ActualAmountPaid creates a ActualAmountPaid
func mockActualAmountPaid() *ActualAmountPaid {
	aap := NewActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = "USD"
	aap.RemittanceAmount.Amount = "1234.56"
	return aap
}

// TestMockActualAmountPaid validates mockActualAmountPaid
func TestMockActualAmountPaid(t *testing.T) {
	aap := mockActualAmountPaid()

	require.NoError(t, aap.Validate(), "mockActualAmountPaid does not validate and will break other tests")
}

// TestActualAmountPaidAmountRequired validates ActualAmountPaid Amount is required
func TestActualAmountPaidAmountRequired(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.Amount = ""

	err := aap.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

// TestActualAmountPaidCurrencyCodeRequired validates ActualAmountPaid CurrencyCode is required
func TestCurrencyCodeRequired(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = ""

	err := aap.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrFieldRequired).Error())
}

// TestActualAmountPaidAmountValid validates Amount
func TestActualAmountPaidAmountValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.Amount = "X,"

	err := aap.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, aap.RemittanceAmount.Amount).Error())
}

// TestActualAmountPaidCurrencyCodeValid validates Amount
func TestActualAmountPaidCurrencyCodeValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = "XZP"

	err := aap.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrNonCurrencyCode, aap.RemittanceAmount.CurrencyCode).Error())
}

// TestParseActualAmountPaidWrongLength parses a wrong ActualAmountPaid record length
func TestParseActualAmountPaidWrongLength(t *testing.T) {
	var line = "{8450}USD1234.56          *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseActualAmountPaid()
	require.NoError(t, err)
}

// TestParseActualAmountPaidReaderParseError parses a wrong ActualAmountPaid reader parse error
func TestParseActualAmountPaidReaderParseError(t *testing.T) {
	var line = "{8450}USD1234.56Z           *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseActualAmountPaid()

	expected := r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)
}

// TestActualAmountPaidTagError validates ActualAmountPaid tag
func TestActualAmountPaidTagError(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.tag = "{9999}"

	err := aap.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, aap.tag).Error())
}

// TestStringActualAmountPaidVariableLength parses using variable length
func TestStringActualAmountPaidVariableLength(t *testing.T) {
	var line = "{8450}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseActualAmountPaid()
	expected := r.parseError(NewTagMinLengthErr(8, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{8450}USD1234.56            NNN*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseActualAmountPaid()
	require.ErrorContains(t, err, ErrNonAmount.Error())

	line = "{8450}****"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseActualAmountPaid()
	require.ErrorContains(t, err, ErrValidLength.Error())

	line = "{8450}USD*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseActualAmountPaid()
	expected = r.parseError(fieldError("Amount", ErrFieldRequired)).Error()
	require.EqualError(t, err, expected)

	line = "{8450}USD1234.56*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseActualAmountPaid()
	require.Equal(t, err, nil)
}

// TestStringActualAmountPaidOptions validates Format() formatted according to the FormatOptions
func TestStringActualAmountPaidOptions(t *testing.T) {
	var line = "{8450}USD1234.56*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseActualAmountPaid()
	require.Equal(t, err, nil)

	aap := r.currentFEDWireMessage.ActualAmountPaid
	require.Equal(t, aap.String(), "{8450}USD1234.56            *")
	require.Equal(t, aap.Format(FormatOptions{VariableLengthFields: true}), "{8450}USD1234.56*")
	require.Equal(t, aap.String(), aap.Format(FormatOptions{VariableLengthFields: false}))
}
