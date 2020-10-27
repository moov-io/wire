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

	require.NotNil(t, err)
	require.Equal(t, fieldError("Amount", ErrFieldRequired).Error(), err.Error())
}

// TestActualAmountPaidCurrencyCodeRequired validates ActualAmountPaid CurrencyCode is required
func TestCurrencyCodeRequired(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = ""

	err := aap.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("CurrencyCode", ErrFieldRequired).Error(), err.Error())
}

// TestActualAmountPaidAmountValid validates Amount
func TestActualAmountPaidAmountValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.Amount = "X,"

	err := aap.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Amount", ErrNonAmount, aap.RemittanceAmount.Amount).Error(), err.Error())
}

// TestActualAmountPaidCurrencyCodeValid validates Amount
func TestActualAmountPaidCurrencyCodeValid(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.RemittanceAmount.CurrencyCode = "XZP"

	err := aap.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("CurrencyCode", ErrNonCurrencyCode, aap.RemittanceAmount.CurrencyCode).Error(), err.Error())
}

// TestParseActualAmountPaidWrongLength parses a wrong ActualAmountPaid record length
func TestParseActualAmountPaidWrongLength(t *testing.T) {
	var line = "{8450}USD1234.56          "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)

	err := r.parseActualAmountPaid()

	require.NotNil(t, err)
	expected := NewTagWrongLengthErr(28, len(r.line))
	require.Contains(t, err.Error(), expected.Error())
}

// TestParseActualAmountPaidReaderParseError parses a wrong ActualAmountPaid reader parse error
func TestParseActualAmountPaidReaderParseError(t *testing.T) {
	var line = "{8450}USD1234.56Z           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	aap := mockActualAmountPaid()
	fwm.SetActualAmountPaid(aap)

	err := r.parseActualAmountPaid()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestActualAmountPaidTagError validates ActualAmountPaid tag
func TestActualAmountPaidTagError(t *testing.T) {
	aap := mockActualAmountPaid()
	aap.tag = "{9999}"

	err := aap.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, aap.tag).Error(), err.Error())
}
