package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiary creates a Beneficiary
func mockBeneficiary() *Beneficiary {
	ben := NewBeneficiary(false)
	ben.Personal.IdentificationCode = DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	return ben
}

// TestMockBeneficiary validates mockBeneficiary
func TestMockBeneficiary(t *testing.T) {
	ben := mockBeneficiary()

	require.NoError(t, ben.Validate(), "mockBeneficiary does not validate and will break other tests")
}

// TestBeneficiaryIdentificationCodeValid validates Beneficiary IdentificationCode
func TestBeneficiaryIdentificationCodeValid(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = "Baseball Card ID"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, ben.Personal.IdentificationCode).Error())
}

// TestBeneficiaryIdentifierAlphaNumeric validates Beneficiary Identifier is alphanumeric
func TestBeneficiaryIdentifierAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Identifier = "®"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, ben.Personal.Identifier).Error())
}

// TestBeneficiaryNameAlphaNumeric validates Beneficiary Name is alphanumeric
func TestBeneficiaryNameAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Name = "®"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, ben.Personal.Name).Error())
}

// TestBeneficiaryAddressLineOneAlphaNumeric validates Beneficiary AddressLineOne is alphanumeric
func TestBeneficiaryAddressLineOneAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineOne = "®"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, ben.Personal.Address.AddressLineOne).Error())
}

// TestBeneficiaryAddressLineTwoAlphaNumeric validates Beneficiary AddressLineTwo is alphanumeric
func TestBeneficiaryAddressLineTwoAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineTwo = "®"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, ben.Personal.Address.AddressLineTwo).Error())
}

// TestBeneficiaryAddressLineThreeAlphaNumeric validates Beneficiary AddressLineThree is alphanumeric
func TestBeneficiaryAddressLineThreeAlphaNumeric(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Address.AddressLineThree = "®"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, ben.Personal.Address.AddressLineThree).Error())
}

// TestBeneficiaryIdentificationCodeRequired validates Beneficiary IdentificationCode is required
func TestBeneficiaryIdentificationCodeRequired(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = ""

	err := ben.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestBeneficiaryIdentifierRequired validates Beneficiary Identifier is required
func TestBeneficiaryIdentifierRequired(t *testing.T) {
	ben := mockBeneficiary()
	ben.Personal.Identifier = ""

	err := ben.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestParseBeneficiaryWrongLength parses a wrong Beneficiary record length
func TestParseBeneficiaryWrongLength(t *testing.T) {
	var line = "{4200}31234                              Name                               Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiary()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(181, len(r.line))).Error())
}

// TestParseBeneficiaryReaderParseError parses a wrong Beneficiary reader parse error
func TestParseBeneficiaryReaderParseError(t *testing.T) {
	var line = "{4200}31234                              Na®e                               Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiary()

	expected := r.parseError(fieldError("Name", ErrNonAlphanumeric, "Na®e")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Name", ErrNonAlphanumeric, "Na®e")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryTagError validates Beneficiary tag
func TestBeneficiaryTagError(t *testing.T) {
	ben := mockBeneficiary()
	ben.tag = "{9999}"

	err := ben.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, ben.tag).Error())
}
