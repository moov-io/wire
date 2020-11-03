package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockOriginatorFI creates a OriginatorFI
func mockOriginatorFI() *OriginatorFI {
	ofi := NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ofi
}

// TestMockOriginatorFI validates mockOriginatorFI
func TestMockOriginatorFI(t *testing.T) {
	ofi := mockOriginatorFI()

	require.NoError(t, ofi.Validate(), "mockOriginatorFI does not validate and will break other tests")
}

// TestOriginatorFIIdentificationCodeValid validates OriginatorFI IdentificationCode
func TestOriginatorFIIdentificationCodeValid(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = "Football Card ID"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, ofi.FinancialInstitution.IdentificationCode).Error())
}

// TestOriginatorFIIdentificationCodeFI validates OriginatorFI IdentificationCode is an FI code
func TestOriginatorFIIdentificationCodeFI(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = "1"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, ofi.FinancialInstitution.IdentificationCode).Error())
}

// TestOriginatorFIIdentifierAlphaNumeric validates OriginatorFI Identifier is alphanumeric
func TestOriginatorFIIdentifierAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Identifier = "®"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, ofi.FinancialInstitution.Identifier).Error())
}

// TestOriginatorFINameAlphaNumeric validates OriginatorFI Name is alphanumeric
func TestOriginatorFINameAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Name = "®"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, ofi.FinancialInstitution.Name).Error())
}

// TestOriginatorFIAddressLineOneAlphaNumeric validates OriginatorFI AddressLineOne is alphanumeric
func TestOriginatorFIAddressLineOneAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineOne = "®"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, ofi.FinancialInstitution.Address.AddressLineOne).Error())
}

// TestOriginatorFIAddressLineTwoAlphaNumeric validates OriginatorFI AddressLineTwo is alphanumeric
func TestOriginatorFIAddressLineTwoAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineTwo = "®"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, ofi.FinancialInstitution.Address.AddressLineTwo).Error())
}

// TestOriginatorFIAddressLineThreeAlphaNumeric validates OriginatorFI AddressLineThree is alphanumeric
func TestOriginatorFIAddressLineThreeAlphaNumeric(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Address.AddressLineThree = "®"

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, ofi.FinancialInstitution.Address.AddressLineThree).Error())
}

// TestOriginatorFIIdentificationCodeRequired validates OriginatorFI IdentificationCode is required
func TestOriginatorFIIdentificationCodeRequired(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = ""

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestOriginatorFIIdentifierRequired validates OriginatorFI Identifier is required
func TestOriginatorFIIdentifierRequired(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.FinancialInstitution.Identifier = ""

	err := ofi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestParseOriginatorFIWrongLength parses a wrong OriginatorFI record length
func TestParseOriginatorFIWrongLength(t *testing.T) {
	var line = "{5100}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorFI()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(181, len(r.line))).Error())
}

// TestParseOriginatorFIReaderParseError parses a wrong OriginatorFI reader parse error
func TestParseOriginatorFIReaderParseError(t *testing.T) {
	var line = "{5100}D123456789                         ®I Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseOriginatorFI()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®I Name")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®I Name")).Error())
}

// TestOriginatorFITagError validates a OriginatorFI tag
func TestOriginatorFITagError(t *testing.T) {
	ofi := mockOriginatorFI()
	ofi.tag = "{9999}"

	require.EqualError(t, ofi.Validate(), fieldError("tag", ErrValidTagForType, ofi.tag).Error())
}
