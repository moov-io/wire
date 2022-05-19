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

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, a.Amount).Error())
}

// TestAmountRequired validates Amount is required
func TestAmountRequired(t *testing.T) {
	a := mockAmount()
	a.Amount = ""

	err := a.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

// TestParseAmountWrongLength parses a wrong Amount record length
func TestParseAmountWrongLength(t *testing.T) {
	var line = "{2000}00"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmount()

	require.EqualError(t, err, r.parseError(fieldError("Amount", ErrValidLengthSize)).Error())
}

// TestParseAmountReaderParseError parses a wrong Amount reader parse error
func TestParseAmountReaderParseError(t *testing.T) {
	var line = "{2000}00000Z030022"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmount()

	expected := r.parseError(fieldError("Amount", ErrNonAmount, "00000Z030022")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Amount", ErrNonAmount, "00000Z030022")).Error()
	require.EqualError(t, err, expected)
}

// TestAmountTagError validates Amount tag
func TestAmountTagError(t *testing.T) {
	a := mockAmount()
	a.tag = "{9999}"

	err := a.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, a.tag).Error())
}
