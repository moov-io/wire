package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// IntermediaryInstitution creates a IntermediaryInstitution
func mockIntermediaryInstitution() *IntermediaryInstitution {
	ii := NewIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ii.CoverPayment.SwiftLineOne = "Swift Line One"
	ii.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ii.CoverPayment.SwiftLineThree = "Swift Line Three"
	ii.CoverPayment.SwiftLineFour = "Swift Line Four"
	ii.CoverPayment.SwiftLineFive = "Swift Line Five"
	return ii
}

// TestMockIntermediaryInstitution validates mockIntermediaryInstitution
func TestMockIntermediaryInstitution(t *testing.T) {
	ii := mockIntermediaryInstitution()

	require.NoError(t, ii.Validate(), "mockIntermediaryInstitution does not validate and will break other tests")
}

// TestIntermediaryInstitutionSwiftFieldTagAlphaNumeric validates IntermediaryInstitution SwiftFieldTag is alphanumeric
func TestIntermediaryInstitutionSwiftFieldTagAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, ii.CoverPayment.SwiftFieldTag).Error())
}

// TestIntermediaryInstitutionSwiftLineOneAlphaNumeric validates IntermediaryInstitution SwiftLineOne is alphanumeric
func TestIntermediaryInstitutionSwiftLineOneAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineOne = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, ii.CoverPayment.SwiftLineOne).Error())
}

// TestIntermediaryInstitutionSwiftLineTwoAlphaNumeric validates IntermediaryInstitution SwiftLineTwo is alphanumeric
func TestIntermediaryInstitutionSwiftLineTwoAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineTwo = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, ii.CoverPayment.SwiftLineTwo).Error())
}

// TestIntermediaryInstitutionSwiftLineThreeAlphaNumeric validates IntermediaryInstitution SwiftLineThree is alphanumeric
func TestIntermediaryInstitutionSwiftLineThreeAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineThree = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, ii.CoverPayment.SwiftLineThree).Error())
}

// TestIntermediaryInstitutionSwiftLineFourAlphaNumeric validates IntermediaryInstitution SwiftLineFour is alphanumeric
func TestIntermediaryInstitutionSwiftLineFourAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineFour = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, ii.CoverPayment.SwiftLineFour).Error())
}

// TestIntermediaryInstitutionSwiftLineFiveAlphaNumeric validates IntermediaryInstitution SwiftLineFive is alphanumeric
func TestIntermediaryInstitutionSwiftLineFiveAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineFive = "®"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, ii.CoverPayment.SwiftLineFive).Error())
}

// TestIntermediaryInstitutionSwiftLineSixAlphaNumeric validates IntermediaryInstitution SwiftLineSix is alphanumeric
func TestIntermediaryInstitutionSwiftLineSixAlphaNumeric(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.CoverPayment.SwiftLineSix = "Test"

	err := ii.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, ii.CoverPayment.SwiftLineSix).Error())
}

// TestParseIntermediaryInstitutionWrongLength parses a wrong IntermediaryInstitution record length
func TestParseIntermediaryInstitutionWrongLength(t *testing.T) {
	var line = "{7056}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseIntermediaryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("SwiftFieldTag", ErrRequireDelimiter)).Error())
}

// TestParseIntermediaryInstitutionReaderParseError parses a wrong IntermediaryInstitution reader parse error
func TestParseIntermediaryInstitutionReaderParseError(t *testing.T) {
	var line = "{7056}Swift*Swift ®ine One                     *Swift Line Two                     *Swift Line Three                   *Swift Line Four                    *Swift Line Five                   *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseIntermediaryInstitution()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())
}

// TestIntermediaryInstitutionTagError validates a IntermediaryInstitution tag
func TestIntermediaryInstitutionTagError(t *testing.T) {
	ii := mockIntermediaryInstitution()
	ii.tag = "{9999}"

	require.EqualError(t, ii.Validate(), fieldError("tag", ErrValidTagForType, ii.tag).Error())
}

// TestStringIntermediaryInstitutionVariableLength parses using variable length
func TestStringIntermediaryInstitutionVariableLength(t *testing.T) {
	var line = "{7056}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseIntermediaryInstitution()
	require.NoError(t, err)

	line = "{7056}                                                                                                                                                                                    NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseIntermediaryInstitution()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{7056}***********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseIntermediaryInstitution()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{7056}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseIntermediaryInstitution()
	require.NoError(t, err)
}

// TestStringIntermediaryInstitutionOptions validates Format() formatted according to the FormatOptions
func TestStringIntermediaryInstitutionOptions(t *testing.T) {
	var line = "{7056}*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseIntermediaryInstitution()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.IntermediaryInstitution
	require.Equal(t, "{7056}     *                                   *                                   *                                   *                                   *                                   *", record.String())
	require.Equal(t, "{7056}*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
