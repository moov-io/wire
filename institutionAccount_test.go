package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// InstitutionAccount creates a InstitutionAccount
func mockInstitutionAccount() *InstitutionAccount {
	iAccount := NewInstitutionAccount()
	iAccount.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	iAccount.CoverPayment.SwiftLineOne = "Swift Line One"
	iAccount.CoverPayment.SwiftLineTwo = "Swift Line Two"
	iAccount.CoverPayment.SwiftLineThree = "Swift Line Three"
	iAccount.CoverPayment.SwiftLineFour = "Swift Line Four"
	iAccount.CoverPayment.SwiftLineFive = "Swift Line Five"
	return iAccount
}

// TestMockInstitutionAccount validates mockInstitutionAccount
func TestMockInstitutionAccount(t *testing.T) {
	iAccount := mockInstitutionAccount()

	require.NoError(t, iAccount.Validate(), "mockInstitutionAccount does not validate and will break other tests")
}

// TestInstitutionAccountSwiftFieldTagAlphaNumeric validates InstitutionAccount SwiftFieldTag is alphanumeric
func TestInstitutionAccountSwiftFieldTagAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftFieldTag = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftFieldTag).Error())
}

// TestInstitutionAccountSwiftLineOneAlphaNumeric validates InstitutionAccount SwiftLineOne is alphanumeric
func TestInstitutionAccountSwiftLineOneAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineOne = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftLineOne).Error())
}

// TestInstitutionAccountSwiftLineTwoAlphaNumeric validates InstitutionAccount SwiftLineTwo is alphanumeric
func TestInstitutionAccountSwiftLineTwoAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineTwo = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftLineTwo).Error())
}

// TestInstitutionAccountSwiftLineThreeAlphaNumeric validates InstitutionAccount SwiftLineThree is alphanumeric
func TestInstitutionAccountSwiftLineThreeAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineThree = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftLineThree).Error())
}

// TestInstitutionAccountSwiftLineFourAlphaNumeric validates InstitutionAccount SwiftLineFour is alphanumeric
func TestInstitutionAccountSwiftLineFourAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineFour = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftLineFour).Error())
}

// TestInstitutionAccountSwiftLineFiveAlphaNumeric validates InstitutionAccount SwiftLineFive is alphanumeric
func TestInstitutionAccountSwiftLineFiveAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineFive = "®"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, iAccount.CoverPayment.SwiftLineFive).Error())
}

// TestInstitutionAccountSwiftLineSixAlphaNumeric validates InstitutionAccount SwiftLineSix is an invalid property
func TestInstitutionAccountSwiftLineSixAlphaNumeric(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.CoverPayment.SwiftLineSix = "Test"

	err := iAccount.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrInvalidProperty, iAccount.CoverPayment.SwiftLineSix).Error())
}

// TestParseInstitutionAccountWrongLength parses a wrong InstitutionAccount record length
func TestParseInstitutionAccountWrongLength(t *testing.T) {
	var line = "{7057}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstitutionAccount()

	require.EqualError(t, err, r.parseError(fieldError("SwiftFieldTag", ErrRequireDelimiter)).Error())
}

// TestParseInstitutionAccountReaderParseError parses a wrong InstitutionAccount reader parse error
func TestParseInstitutionAccountReaderParseError(t *testing.T) {
	var line = "{7057}Swift*Swift ®ine One                     *Swift Line Two                     *Swift Line Three                   *Swift Line Four                    *Swift Line Five                   *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstitutionAccount()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "Swift ®ine One")).Error())
}

// TestInstitutionAccountTagError validates a InstitutionAccount tag
func TestInstitutionAccountTagError(t *testing.T) {
	iAccount := mockInstitutionAccount()
	iAccount.tag = "{9999}"

	require.EqualError(t, iAccount.Validate(), fieldError("tag", ErrValidTagForType, iAccount.tag).Error())
}

// TestStringInstitutionAccountVariableLength parses using variable length
func TestStringInstitutionAccountVariableLength(t *testing.T) {
	var line = "{7057}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstitutionAccount()
	require.NoError(t, err)

	line = "{7057}Swift                                                                                                                                                                               NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstitutionAccount()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{7057}********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstitutionAccount()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{7057}*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseInstitutionAccount()
	require.NoError(t, err)
}

// TestStringInstitutionAccountOptions validates Format() formatted according to the FormatOptions
func TestStringInstitutionAccountOptions(t *testing.T) {
	var line = "{7057}Swift*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstitutionAccount()
	require.NoError(t, err)

	record := r.currentFEDWireMessage.InstitutionAccount
	require.Equal(t, "{7057}Swift*                                   *                                   *                                   *                                   *                                   *", record.String())
	require.Equal(t, "{7057}Swift*", record.Format(FormatOptions{VariableLengthFields: true}))
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
