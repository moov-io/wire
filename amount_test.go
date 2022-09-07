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

func TestAmount_Validate(t *testing.T) {
	tests := []struct {
		inAmount string
		wantErr  error
	}{
		{mockAmount().Amount, nil},
		{"", fieldError("Amount", ErrFieldRequired)},
		{"X,", fieldError("Amount", ErrNonAmount, "X,")},
		{"12.05", fieldError("Amount", ErrNonAmount, "12.05")},
		{"1,000.39", fieldError("Amount", ErrNonAmount, "1,000.39")},
	}

	for _, tt := range tests {
		t.Run(tt.inAmount, func(t *testing.T) {
			amt := NewAmount()
			amt.Amount = tt.inAmount

			got := amt.Validate()

			if tt.wantErr == nil {
				require.NoError(t, got)
			} else {
				require.Error(t, got)
				require.Equal(t, tt.wantErr, got)
			}
		})
	}
}

// TestParseAmountWrongLength parses a wrong Amount record length
func TestParseAmountWrongLength(t *testing.T) {
	var line = "{2000}00"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAmount()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(18, len(r.line))).Error())
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
