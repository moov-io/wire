package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Remittance creates a Remittance
func mockRemittance() *Remittance {
	ri := NewRemittance()
	ri.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ri.CoverPayment.SwiftLineOne = "Swift Line One"
	ri.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ri.CoverPayment.SwiftLineThree = "Swift Line Three"
	ri.CoverPayment.SwiftLineFour = "Swift Line Four"
	return ri
}

// TestMockRemittance validates mockRemittance
func TestMockRemittance(t *testing.T) {
	ri := mockRemittance()

	require.NoError(t, ri.Validate(), "mockRemittance does not validate and will break other tests")
}

// TestRemittanceSwiftFieldTagAlphaNumeric validates Remittance SwiftFieldTag is alphanumeric
func TestRemittanceSwiftFieldTagAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftFieldTag = "®"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, ri.CoverPayment.SwiftFieldTag).Error())
}

// TestRemittanceSwiftLineOneAlphaNumeric validates Remittance SwiftLineOne is alphanumeric
func TestRemittanceSwiftLineOneAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineOne = "®"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, ri.CoverPayment.SwiftLineOne).Error())
}

// TestRemittanceSwiftLineTwoAlphaNumeric validates Remittance SwiftLineTwo is alphanumeric
func TestRemittanceSwiftLineTwoAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineTwo = "®"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, ri.CoverPayment.SwiftLineTwo).Error())
}

// TestRemittanceSwiftLineThreeAlphaNumeric validates Remittance SwiftLineThree is alphanumeric
func TestRemittanceSwiftLineThreeAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineThree = "®"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, ri.CoverPayment.SwiftLineThree).Error())
}

// TestRemittanceSwiftLineFourAlphaNumeric validates Remittance SwiftLineFour is alphanumeric
func TestRemittanceSwiftLineFourAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineFour = "®"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, ri.CoverPayment.SwiftLineFour).Error())
}

// TestRemittanceSwiftLineFiveAlphaNumeric validates Remittance SwiftLineFive is an invalid property
func TestRemittanceSwiftLineFiveAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineFive = "Test"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrInvalidProperty, ri.CoverPayment.SwiftLineFive).Error())
}

// TestRemittanceSwiftLineSixAlphaNumeric validates Remittance SwiftLineSix is an invalid property
func TestRemittanceSwiftLineSixAlphaNumeric(t *testing.T) {
	ri := mockRemittance()
	ri.CoverPayment.SwiftLineSix = "Test"

	err := ri.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, ri.CoverPayment.SwiftLineSix).Error())
}

// TestParseRemittanceWrongLength parses a wrong Remittance record length
func TestParseRemittanceWrongLength(t *testing.T) {
	var line = "{7070}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittance()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineFour", ErrValidLengthSize)).Error())
}

// TestParseRemittanceReaderParseError parses a wrong Remittance reader parse error
func TestParseRemittanceReaderParseError(t *testing.T) {
	var line = "{7070}Swift®wift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                   "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittance()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "®wift Line One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "®wift Line One")).Error())
}

// TestRemittanceTagError validates a Remittance tag
func TestRemittanceTagError(t *testing.T) {
	ri := mockRemittance()
	ri.tag = "{9999}"

	require.EqualError(t, ri.Validate(), fieldError("tag", ErrValidTagForType, ri.tag).Error())
}

// TestStringRemittanceVariableLength parses using variable length
func TestStringRemittanceVariableLength(t *testing.T) {
	var line = "{7070}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittance()
	require.Nil(t, err)

	line = "{7070}                                                                                                                                                 NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittance()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{7070}************"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittance()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{7070}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittance()
	require.Equal(t, err, nil)
}

// TestStringRemittanceOptions validates string() with options
func TestStringRemittanceOptions(t *testing.T) {
	var line = "{7070}*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittance()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.Remittance.String()
	require.Equal(t, str, "{7070}                                                                                                                                                 ")

	str = r.currentFEDWireMessage.Remittance.String(true)
	require.Equal(t, str, "{7070}*")
}
