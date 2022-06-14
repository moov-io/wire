package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockAccountDebitedDrawdown creates a AccountDebitedDrawdown
func mockAccountDebitedDrawdown() *AccountDebitedDrawdown {
	debitDD := NewAccountDebitedDrawdown()
	debitDD.IdentificationCode = DemandDepositAccountNumber
	debitDD.Identifier = "123456789"
	debitDD.Name = "debitDD Name"
	debitDD.Address.AddressLineOne = "Address One"
	debitDD.Address.AddressLineTwo = "Address Two"
	debitDD.Address.AddressLineThree = "Address Three"
	return debitDD
}

// TestMockAccountDebitedDrawdown validates mockAccountDebitedDrawdown
func TestMockAccountDebitedDrawdown(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()

	require.NoError(t, debitDD.Validate(), "mockAccountDebitedDrawdown does not validate and will break other tests")
}

// TestADDIdentifierAlphaNumeric validates Identifier is alphanumeric
func TestADDIdentifierAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = "®"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, debitDD.Identifier).Error())
}

// TestADDNameAlphaNumeric validates Name is alphanumeric
func TestADDNameAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = "®"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, debitDD.Name).Error())
}

// TestADDAddressLineOneAlphaNumeric validates AddressLineOne is alphanumeric
func TestADDAddressLineOneAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineOne = "®"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, debitDD.Address.AddressLineOne).Error())
}

// TestADDAddressLineTwoAlphaNumeric validates AddressLineTwo is alphanumeric
func TestADDAddressLineTwoAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineTwo = "®"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, debitDD.Address.AddressLineTwo).Error())
}

// TestADDAddressLineThreeAlphaNumeric validates AddressLineThree is alphanumeric
func TestADDAddressLineThreeAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineThree = "®"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, debitDD.Address.AddressLineThree).Error())
}

// TestADDIdentifierRequired validates Identifier is required
func TestADDIdentifierRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = ""

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestADDNameRequired validates Name is required
func TestADDNameRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = ""

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("Name", ErrFieldRequired).Error())
}

// TestADDIdentificationRequired validates IdentificationCode is required
func TestADDIdentificationCodeRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = ""

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestADDIdentificationCodeValid validates IdentificationCode
func TestADDIdentificationCodeValid(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = TaxIdentificationNumber

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode).Error())
}

// TestADDIdentificationCodeBogus validates IdentificationCode if the IdentificationCode is bogus
func TestIdentificationCodeBogus(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = "Card ID"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode).Error())
}

// TestParseAccountDebitedDrawdownWrongLength parses a wrong AccountDebitedDrawdown record length
func TestParseAccountDebitedDrawdownWrongLength(t *testing.T) {
	var line = "{4400}D123456789                         debitDD Name                       Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()

	require.EqualError(t, err, r.parseError(fieldError("AddressLineThree", ErrValidLength)).Error())
}

// TestParseAccountDebitedDrawdownReaderParseError parses a wrong AccountDebitedDrawdown reader parse error
func TestParseAccountDebitedDrawdownReaderParseError(t *testing.T) {
	var line = "{4400}D123456789                         debitDD ®ame                       Address One                        Address Two                        Address Three                     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()

	expected := r.parseError(fieldError("Name", ErrNonAlphanumeric, "debitDD ®ame")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Name", ErrNonAlphanumeric, "debitDD ®ame")).Error()
	require.EqualError(t, err, expected)
}

// TestAccountDebitedDrawdownTagError validates AccountDebitedDrawdown tag
func TestAccountDebitedDrawdownTagError(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.tag = "{9999}"

	err := debitDD.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, debitDD.tag).Error())
}

// TestStringDebitedDrawdownVariableLength parses using variable length
func TestStringAccountDebitedDrawdownVariableLength(t *testing.T) {
	var line = "{4400}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()
	expected := r.parseError(NewTagMinLengthErr(9, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{4400}D123456789                         debitDD Name                       Address One                        Address Two                        Address Three                    NNNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountDebitedDrawdown()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{4400}***"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountDebitedDrawdown()
	expected = r.parseError(fieldError("Identifier", ErrFieldRequired)).Error()
	require.EqualError(t, err, expected)

	line = "{4400}D2*3*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseAccountDebitedDrawdown()
	require.Equal(t, err, nil)
}

// TestStringDebitedDrawdownOptions validates string() with options
func TestStringAccountDebitedDrawdownOptions(t *testing.T) {
	var line = "{4400}D2*3*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.AccountDebitedDrawdown.String()
	require.Equal(t, str, "{4400}D2                                 3                                                                                                                                           ")

	str = r.currentFEDWireMessage.AccountDebitedDrawdown.String(true)
	require.Equal(t, str, "{4400}D2*3*")
}
