package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBusinessFunctionCode creates a BusinessFunctionCode
func mockBusinessFunctionCode() *BusinessFunctionCode {
	bfc := NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = CustomerTransfer
	bfc.TransactionTypeCode = "   "
	return bfc
}

// TestMockBusinessFunctionCode validates mockBusinessFunctionCode
func TestMockBusinessFunctionCode(t *testing.T) {
	bfc := mockBusinessFunctionCode()

	require.NoError(t, bfc.Validate(), "mockBusinessFunctionCode does not validate and will break other tests")
}

// TestBusinessFunctionCodeValid validates BusinessFunctionCode
func TestBusinessFunctionCodeValid(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = "ZZZ"

	err := bfc.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrBusinessFunctionCode.Error())
}

// TestBusinessFunctionCodeRequired validates BusinessFunctionCode is required
func TestBusinessFunctionCodeRequired(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.BusinessFunctionCode = ""

	err := bfc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("BusinessFunctionCode", ErrFieldRequired, bfc.BusinessFunctionCode).Error(), err.Error())
}

// TestParseBusinessFunctionCodeWrongLength parses a wrong BusinessFunctionCode record length
func TestParseBusinessFunctionCodeWrongLength(t *testing.T) {
	var line = "{3600}CT"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(12, len(r.line)).Error())
}

// TestParseBusinessFunctionCodeReaderParseError parses a wrong BusinessFunctionCode reader parse error
func TestParseBusinessFunctionCodeReaderParseError(t *testing.T) {
	var line = "{3600}CTAXXY"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBusinessFunctionCode()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrBusinessFunctionCode.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrBusinessFunctionCode.Error())
}

// TestBusinessFunctionCodeTagError validates a BusinessFunctionCode tag
func TestBusinessFunctionCodeTagError(t *testing.T) {
	bfc := mockBusinessFunctionCode()
	bfc.tag = "{9999}"

	err := bfc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, bfc.tag).Error(), err.Error())
}
