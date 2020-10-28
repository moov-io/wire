package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiaryCustomer creates a BeneficiaryCustomer
func mockBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := NewBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	bc.CoverPayment.SwiftLineOne = "Swift Line One"
	bc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	bc.CoverPayment.SwiftLineThree = "Swift Line Three"
	bc.CoverPayment.SwiftLineFour = "Swift Line Four"
	bc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return bc
}

// TestMockBeneficiaryCustomer validates mockBeneficiaryCustomer
func TestMockBeneficiaryCustomer(t *testing.T) {
	bc := mockBeneficiaryCustomer()

	require.NoError(t, bc.Validate(), "mockBeneficiaryCustomer does not validate and will break other tests")
}

// TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric validates BeneficiaryCustomer SwiftFieldTag is alphanumeric
func TestBeneficiaryCustomerSwiftFieldTagAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftFieldTag", ErrNonAlphanumeric, bc.CoverPayment.SwiftFieldTag).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineOneAlphaNumeric validates BeneficiaryCustomer SwiftLineOne is alphanumeric
func TestBeneficiaryCustomerSwiftLineOneAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineOne = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineOne", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineOne).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric validates BeneficiaryCustomer SwiftLineTwo is alphanumeric
func TestBeneficiaryCustomerSwiftLineTwoAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineTwo = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineTwo", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineTwo).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric validates BeneficiaryCustomer SwiftLineThree is alphanumeric
func TestBeneficiaryCustomerSwiftLineThreeAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineThree = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineThree", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineThree).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineFourAlphaNumeric validates BeneficiaryCustomer SwiftLineFour is alphanumeric
func TestBeneficiaryCustomerSwiftLineFourAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFour = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineFour", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineFour).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric validates BeneficiaryCustomer SwiftLineFive is alphanumeric
func TestBeneficiaryCustomerSwiftLineFiveAlphaNumeric(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.CoverPayment.SwiftLineFive = "®"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineFive", ErrNonAlphanumeric, bc.CoverPayment.SwiftLineFive).Error(), err.Error())
}

// TestBeneficiaryCustomerSwiftLineSixAlphaNumeric validates BeneficiaryCustomer SwiftLineSix is alphanumeric
func TestBeneficiaryCustomerSwiftLineSixAlphaNumeric(t *testing.T) {
	sr := mockBeneficiaryCustomer()
	sr.CoverPayment.SwiftLineSix = "Test"

	err := sr.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("SwiftLineSix", ErrInvalidProperty, sr.CoverPayment.SwiftLineSix).Error(), err.Error())
}

// TestParseBeneficiaryCustomerWrongLength parses a wrong BeneficiaryCustomer record length
func TestParseBeneficiaryCustomerWrongLength(t *testing.T) {
	var line = "{7059}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryCustomer()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(186, len(r.line)).Error())
}

// TestParseBeneficiaryCustomerReaderParseError parses a wrong BeneficiaryCustomer reader parse error
func TestParseBeneficiaryCustomerReaderParseError(t *testing.T) {
	var line = "{7059}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryCustomer()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestBeneficiaryCustomerTagError validates a BeneficiaryCustomer tag
func TestBeneficiaryCustomerTagError(t *testing.T) {
	bc := mockBeneficiaryCustomer()
	bc.tag = "{9999}"

	err := bc.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, bc.tag).Error(), err.Error())
}
