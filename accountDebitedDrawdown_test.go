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

// TestADDIdentifierAlphaNumeric validates Name is alphanumeric
func TestADDIdentifierAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = "®"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestADDNameAlphaNumeric validates Identifier is alphanumeric
func TestADDNameAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = "®"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestADDAddressLineOneAlphaNumeric validates AddressLineOne is alphanumeric
func TestADDAddressLineOneAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineOne = "®"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestADDAddressLineTwoAlphaNumeric validates AddressLineTwo is alphanumeric
func TestADDAddressLineTwoAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineTwo = "®"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestADDAddressLineThreeAlphaNumeric validates AddressLineThree is alphanumeric
func TestADDAddressLineThreeAlphaNumeric(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Address.AddressLineThree = "®"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestADDIdentifierRequired validates Identifier is required
func TestADDIdentifierRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Identifier = ""

	err := debitDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("Identifier", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

// TestADDNameRequired validates Name is required
func TestADDNameRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.Name = ""

	err := debitDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("Name", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

// TestADDIdentificationRequired validates IdentificationCode is required
func TestADDIdentificationCodeRequired(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = ""

	err := debitDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("IdentificationCode", ErrFieldRequired).Error()
	require.Equal(t, expected, err.Error())
}

// TestADDIdentificationCodeValid validates IdentificationCode
func TestADDIdentificationCodeValid(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = TaxIdentificationNumber

	err := debitDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode).Error()
	require.Equal(t, expected, err.Error())
}

// TestADDIdentificationCodeBogus validates IdentificationCode if the IdentificationCode is bogus
func TestIdentificationCodeBogus(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.IdentificationCode = "Card ID"

	err := debitDD.Validate()

	require.NotNil(t, err)
	expected := fieldError("IdentificationCode", ErrIdentificationCode, debitDD.IdentificationCode).Error()
	require.Equal(t, expected, err.Error())
}

// TestParseAccountDebitedDrawdownWrongLength parses a wrong AccountDebitedDrawdown record length
func TestParseAccountDebitedDrawdownWrongLength(t *testing.T) {
	var line = "{4400}D123456789                         debitDD Name                       Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()

	require.NotNil(t, err)
	expected := r.parseError(NewTagWrongLengthErr(181, len(r.line))).Error()
	require.Equal(t, expected, err.Error())
}

// TestParseAccountDebitedDrawdownReaderParseError parses a wrong AccountDebitedDrawdown reader parse error
func TestParseAccountDebitedDrawdownReaderParseError(t *testing.T) {
	var line = "{4400}D123456789                         debitDD ®ame                       Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseAccountDebitedDrawdown()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	fwm := mockCustomerTransferData()
	fwm.AccountDebitedDrawdown = &AccountDebitedDrawdown{}
	if err := fwm.AccountDebitedDrawdown.Parse(line); err != nil {
		t.Fatal(err)
	}
	fwm.Beneficiary = mockBeneficiary()
	fwm.Originator = mockOriginator()
	r.currentFEDWireMessage = fwm

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestAccountDebitedDrawdownTagError validates AccountDebitedDrawdown tag
func TestAccountDebitedDrawdownTagError(t *testing.T) {
	debitDD := mockAccountDebitedDrawdown()
	debitDD.tag = "{9999}"

	err := debitDD.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, debitDD.tag).Error(), err.Error())
}
