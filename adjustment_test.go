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

	require.EqualError(t, err, fieldError("AdjustmentReasonCode", ErrAdjustmentReasonCode, adj.AdjustmentReasonCode).Error())
}

// TestCreditDebitIndicatorValid validates Adjustment CreditDebitIndicator
func TestCreditDebitIndicatorValid(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = "ZZZZ"

	err := adj.Validate()

	require.EqualError(t, err, fieldError("CreditDebitIndicator", ErrCreditDebitIndicator, adj.CreditDebitIndicator).Error())
}

// TestAdjustmentAmountValid validates Adjustment Amount
func TestAdjustmentAmountValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = "X,"

	err := adj.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrNonAmount, adj.RemittanceAmount.Amount).Error())
}

// TestAdjustmentCurrencyCodeValid validates Adjustment CurrencyCode
func TestAdjustmentCurrencyCodeValid(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = "XZP"

	err := adj.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrNonCurrencyCode, adj.RemittanceAmount.CurrencyCode).Error())
}

// TestAdjustmentReasonCodeRequired validates Adjustment AdjustmentReasonCode is required
func TestAdjustmentReasonCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.AdjustmentReasonCode = ""

	err := adj.Validate()

	require.EqualError(t, err, fieldError("AdjustmentReasonCode", ErrFieldRequired).Error())
}

// TestCreditDebitIndicatorRequired validates Adjustment CreditDebitIndicator is required
func TestCreditDebitIndicatorRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.CreditDebitIndicator = ""

	err := adj.Validate()

	require.EqualError(t, err, fieldError("CreditDebitIndicator", ErrFieldRequired).Error())
}

// TestAdjustmentAmountRequired validates Adjustment Amount is required
func TestAdjustmentAmountRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.Amount = ""

	err := adj.Validate()

	require.EqualError(t, err, fieldError("Amount", ErrFieldRequired).Error())
}

// TestAdjustmentCurrencyCodeRequired validates Adjustment CurrencyCode is required
func TestAdjustmentCurrencyCodeRequired(t *testing.T) {
	adj := mockAdjustment()
	adj.RemittanceAmount.CurrencyCode = ""

	err := adj.Validate()

	require.EqualError(t, err, fieldError("CurrencyCode", ErrFieldRequired).Error())
}

// TestParseAdjustmentWrongLength parses a wrong Adjustment record length
func TestParseAdjustmentWrongLength(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()

	require.EqualError(t, err, r.parseError(fieldError("AdditionalInfo", ErrValidLength)).Error())
}

// TestParseAdjustmentReaderParseError parses a wrong Adjustment reader parse error
func TestParseAdjustmentReaderParseError(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56Z             Adjustment Additional Information                                                                                                         "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()

	expected := r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Amount", ErrNonAmount, "1234.56Z")).Error()
	require.EqualError(t, err, expected)
}

// TestAdjustmentTagError validates Adjustment tag
func TestAdjustmentTagError(t *testing.T) {
	adj := mockAdjustment()
	adj.tag = "{9999}"

	err := adj.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, adj.tag).Error())
}

// TestStringAdjustmentVariableLength parses using variable length
func TestStringAdjustmentVariableLength(t *testing.T) {
	var line = "{8600}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()
	expected := r.parseError(NewTagMinLengthErr(10, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{8600}01CRDTUSD1234.56                                                                                                                                                        NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAdjustment()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{8600}01CRDTUSD1234.56****"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAdjustment()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{8600}01CRDTUSD1234.56*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAdjustment()
	require.Equal(t, err, nil)
}

// TestStringAdjustmentOptions validates string() with options
func TestStringAdjustmentOptions(t *testing.T) {
	var line = "{8600}01CRDTUSD1234.56            *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAdjustment()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.Adjustment.String()
	require.Equal(t, str, "{8600}01CRDTUSD1234.56                                                                                                                                                        ")

	str = r.currentFEDWireMessage.Adjustment.String(true)
	require.Equal(t, str, "{8600}01CRDTUSD1234.56*")
}
