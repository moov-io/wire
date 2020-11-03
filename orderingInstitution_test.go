package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//  OrderingInstitution creates a OrderingInstitution
func mockOrderingInstitution() *OrderingInstitution {
	oi := NewOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oi.CoverPayment.SwiftLineOne = "Swift Line One"
	oi.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oi.CoverPayment.SwiftLineThree = "Swift Line Three"
	oi.CoverPayment.SwiftLineFour = "Swift Line Four"
	oi.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oi
}

// TestMockOrderingInstitution validates mockOrderingInstitution
func TestMockOrderingInstitution(t *testing.T) {
	oi := mockOrderingInstitution()

	require.NoError(t, oi.Validate(), "mockOrderingInstitution does not validate and will break other tests")
}

// TestOrderingInstitutionSwiftFieldTagAlphaNumeric validates OrderingInstitution SwiftFieldTag is alphanumeric
func TestOrderingInstitutionSwiftFieldTagAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, oi.CoverPayment.SwiftFieldTag).Error())
}

// TestOrderingInstitutionSwiftLineOneAlphaNumeric validates OrderingInstitution SwiftLineOne is alphanumeric
func TestOrderingInstitutionSwiftLineOneAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineOne = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, oi.CoverPayment.SwiftLineOne).Error())
}

// TestOrderingInstitutionSwiftLineTwoAlphaNumeric validates OrderingInstitution SwiftLineTwo is alphanumeric
func TestOrderingInstitutionSwiftLineTwoAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineTwo = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, oi.CoverPayment.SwiftLineTwo).Error())
}

// TestOrderingInstitutionSwiftLineThreeAlphaNumeric validates OrderingInstitution SwiftLineThree is alphanumeric
func TestOrderingInstitutionSwiftLineThreeAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineThree = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, oi.CoverPayment.SwiftLineThree).Error())
}

// TestOrderingInstitutionSwiftLineFourAlphaNumeric validates OrderingInstitution SwiftLineFour is alphanumeric
func TestOrderingInstitutionSwiftLineFourAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineFour = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, oi.CoverPayment.SwiftLineFour).Error())
}

// TestOrderingInstitutionSwiftLineFiveAlphaNumeric validates OrderingInstitution SwiftLineFive is alphanumeric
func TestOrderingInstitutionSwiftLineFiveAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineFive = "®"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, oi.CoverPayment.SwiftLineFive).Error())
}

// TestOrderingInstitutionSwiftLineSixAlphaNumeric validates OrderingInstitution SwiftLineSix is alphanumeric
func TestOrderingInstitutionSwiftLineSixAlphaNumeric(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.CoverPayment.SwiftLineSix = "Test"

	err := oi.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, oi.CoverPayment.SwiftLineSix).Error())
}

// TestParseOrderingInstitutionWrongLength parses a wrong OrderingInstitution record length
func TestParseOrderingInstitutionWrongLength(t *testing.T) {
	var line = "{7052}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingInstitution()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(186, len(r.line))).Error())
}

// TestParseOrderingInstitutionReaderParseError parses a wrong OrderingInstitution reader parse error
func TestParseOrderingInstitutionReaderParseError(t *testing.T) {
	var line = "{7052}SwiftSwift ®ine One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOrderingInstitution()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())
}

// TestOrderingInstitutionTagError validates a OrderingInstitution tag
func TestOrderingInstitutionTagError(t *testing.T) {
	oi := mockOrderingInstitution()
	oi.tag = "{9999}"

	require.EqualError(t, oi.Validate(), fieldError("tag", ErrValidTagForType, oi.tag).Error())
}
