package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockAmount creates an a Amount
func mockAmount() *Amount {
	a := NewAmount()
	a.Amount = "000001234567"
	return a
}

// TestMockAmount validates mockAmount
func TestMockAmount(t *testing.T) {
	a := mockAmount()

	require.NoError(t, a.Validate(), "mockAmount does not validate and will break other tests")
}

// TestAmountValid validates Amount
func TestAmountValid(t *testing.T) {
	a := mockAmount()
	a.Amount = "X,"

	err := a.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAmountRequired validates Amount is required
func TestAmountRequired(t *testing.T) {
	a := mockAmount()
	a.Amount = ""

	err := a.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Amount", ErrFieldRequired).Error(), err.Error())
}

// TestParseAmountWrongLength parses a wrong Amount record length
func TestParseAmountWrongLength(t *testing.T) {
	var line = "{2000}00"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmount()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(18, len(r.line)).Error())
}

// TestParseAmountReaderParseError parses a wrong Amount reader parse error
func TestParseAmountReaderParseError(t *testing.T) {
	var line = "{2000}00000Z030022"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmount()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAmountTagError validates Amount tag
func TestAmountTagError(t *testing.T) {
	a := mockAmount()
	a.tag = "{9999}"

	err := a.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, a.tag).Error(), err.Error())
}
