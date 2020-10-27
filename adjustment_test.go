package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// MockAdjustment creates a Adjustment
func mockAdjustment() *Adjustment {
	adj := NewAdjustment()
	adj.AdjustmentReasonCode = PricingError
	adj.CreditDebitIndicator = CreditIndicator
	adj.RemittanceAmount.CurrencyCode = "USD"
	adj.RemittanceAmount.Amount = "1234.56"
	adj.AdditionalInfo = " Adjustment Additional Information"
	return adj
}

// TestMockAdjustment validates mockAdjustment
func TestMockAdjustment(t *testing.T) {
	adj := mockAdjustment()

	require.NoError(t, adj.Validate(), "mockAdjustment does not validate and will break other tests")
}

// TestAdjustmentReasonCodeValid validates Adjustment AdjustmentReasonCode
func TestAdjustmentReasonCodeValid(t *testing.T) {
	adj := mockAdjustment()
	adj.AdjustmentReasonCode = "ZZ"

	err := adj.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrAdjustmentReasonCode.Error())
}

// TestCreditDebitIndicatorValid validates Adjustment CreditDebitIndicator
func TestCreditDebitIndicatorValid(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = "ZZZZ"

	err := adj.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrCreditDebitIndicator.Error())
}

// TestAdjustmentAmountValid validates Adjustment Amount
func TestAdjustmentAmountValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = "X,"

	err := adj.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAdjustmentCurrencyCodeValid validates Adjustment CurrencyCode
func TestAdjustmentCurrencyCodeValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = "XZP"

	err := adj.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonCurrencyCode.Error())
}

// TestAdjustmentReasonCodeRequired validates Adjustment AdjustmentReasonCode is required
func TestAdjustmentReasonCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.AdjustmentReasonCode = ""

	err := adj.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AdjustmentReasonCode", ErrFieldRequired).Error(), err.Error())
}

// TestCreditDebitIndicatorRequired validates Adjustment CreditDebitIndicator is required
func TestCreditDebitIndicatorRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = ""

	err := adj.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("CreditDebitIndicator", ErrFieldRequired).Error(), err.Error())
}

// TestAdjustmentAmountRequired validates Adjustment Amount is required
func TestAdjustmentAmountRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = ""

	err := adj.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Amount", ErrFieldRequired).Error(), err.Error())
}

// TestAdjustmentCurrencyCodeRequired validates Adjustment CurrencyCode is required
func TestAdjustmentCurrencyCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = ""

	err := adj.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("CurrencyCode", ErrFieldRequired).Error(), err.Error())
}

// TestParseAdjustmentWrongLength parses a wrong Adjustment record length
func TestParseAdjustmentWrongLength(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(174, len(r.line)).Error())
}

// TestParseAdjustmentReaderParseError parses a wrong Adjustment reader parse error
func TestParseAdjustmentReaderParseError(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAmount.Error())
}

// TestAdjustmentTagError validates Adjustment tag
func TestAdjustmentTagError(t *testing.T) {
	adj := mockAdjustment()
	adj.tag = "{9999}"

	err := adj.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, adj.tag).Error(), err.Error())
}
