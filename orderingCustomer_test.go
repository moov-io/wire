package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// OrderingCustomer creates a OrderingCustomer
func mockOrderingCustomer() *OrderingCustomer {
	oc := NewOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oc.CoverPayment.SwiftLineOne = "Swift Line One"
	oc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oc.CoverPayment.SwiftLineThree = "Swift Line Three"
	oc.CoverPayment.SwiftLineFour = "Swift Line Four"
	oc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oc
}

// TestMockOrderingCustomer validates mockOrderingCustomer
func TestMockOrderingCustomer(t *testing.T) {
	oc := mockOrderingCustomer()

	require.NoError(t, oc.Validate(), "mockOrderingCustomer does not validate and will break other tests")
}

// TestOrderingCustomerSwiftFieldTagAlphaNumeric validates OrderingCustomer SwiftFieldTag is alphanumeric
func TestOrderingCustomerSwiftFieldTagAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, oc.CoverPayment.SwiftFieldTag).Error())
}

// TestOrderingCustomerSwiftLineOneAlphaNumeric validates OrderingCustomer SwiftLineOne is alphanumeric
func TestOrderingCustomerSwiftLineOneAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineOne = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, oc.CoverPayment.SwiftLineOne).Error())
}

// TestOrderingCustomerSwiftLineTwoAlphaNumeric validates OrderingCustomer SwiftLineTwo is alphanumeric
func TestOrderingCustomerSwiftLineTwoAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineTwo = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, oc.CoverPayment.SwiftLineTwo).Error())
}

// TestOrderingCustomerSwiftLineThreeAlphaNumeric validates OrderingCustomer SwiftLineThree is alphanumeric
func TestOrderingCustomerSwiftLineThreeAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineThree = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, oc.CoverPayment.SwiftLineThree).Error())
}

// TestOrderingCustomerSwiftLineFourAlphaNumeric validates OrderingCustomer SwiftLineFour is alphanumeric
func TestOrderingCustomerSwiftLineFourAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineFour = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, oc.CoverPayment.SwiftLineFour).Error())
}

// TestOrderingCustomerSwiftLineFiveAlphaNumeric validates OrderingCustomer SwiftLineFive is alphanumeric
func TestOrderingCustomerSwiftLineFiveAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineFive = "®"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, oc.CoverPayment.SwiftLineFive).Error())
}

// TestOrderingCustomerSwiftLineSixAlphaNumeric validates OrderingCustomer SwiftLineSix is alphanumeric
func TestOrderingCustomerSwiftLineSixAlphaNumeric(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.CoverPayment.SwiftLineSix = "Test"

	err := oc.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, oc.CoverPayment.SwiftLineSix).Error())
}

// TestParseOrderingCustomerWrongLength parses a wrong OrderingCustomer record length
func TestParseOrderingCustomerWrongLength(t *testing.T) {
	var line = "{7050}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingCustomer()

	require.EqualError(t, err, r.parseError(fieldError("SwiftFieldTag", ErrRequireDelimiter)).Error())
}

// TestParseOrderingCustomerReaderParseError parses a wrong OrderingCustomer reader parse error
func TestParseOrderingCustomerReaderParseError(t *testing.T) {
	var line = "{7050}Swift*Swift ®ine One                     *Swift Line Two                     *Swift Line Three                   *Swift Line Four                    *Swift Line Five                   *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingCustomer()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())
}

// TestOrderingCustomerTagError validates a OrderingCustomer tag
func TestOrderingCustomerTagError(t *testing.T) {
	oc := mockOrderingCustomer()
	oc.tag = "{9999}"

	require.EqualError(t, oc.Validate(), fieldError("tag", ErrValidTagForType, oc.tag).Error())
}

// TestStringOrderingCustomerVariableLength parses using variable length
func TestStringOrderingCustomerVariableLength(t *testing.T) {
	var line = "{7050}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingCustomer()
	require.Nil(t, err)

	line = "{7050}                                                                                                                                                                                    NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOrderingCustomer()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{7050}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOrderingCustomer()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{7050}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseOrderingCustomer()
	require.Equal(t, err, nil)
}

// TestStringOrderingCustomerOptions validates Format() formatted according to the FormatOptions
func TestStringOrderingCustomerOptions(t *testing.T) {
	var line = "{7050}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingCustomer()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.OrderingCustomer
	require.Equal(t, record.String(), "{7050}     *                                   *                                   *                                   *                                   *                                   *")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{7050}*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
