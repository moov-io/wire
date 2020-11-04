package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockInstructingFI creates a InstructingFI
func mockInstructingFI() *InstructingFI {
	ifi := NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ifi
}

// TestMockInstructingFI validates mockInstructingFI
func TestMockInstructingFI(t *testing.T) {
	bfi := mockInstructingFI()

	require.NoError(t, bfi.Validate(), "mockInstructingFI does not validate and will break other tests")
}

// TestInstructingFIIdentificationCodeValid validates InstructingFI IdentificationCode
func TestInstructingFIIdentificationCodeValid(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = "Football Card ID"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bfi.FinancialInstitution.IdentificationCode).Error())
}

// TestInstructingFIIdentificationCodeFI validates InstructingFI IdentificationCode is an FI code
func TestInstructingFIIdentificationCodeFI(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = "1"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bfi.FinancialInstitution.IdentificationCode).Error())
}

// TestInstructingFIIdentifierAlphaNumeric validates InstructingFI Identifier is alphanumeric
func TestInstructingFIIdentifierAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Identifier = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, bfi.FinancialInstitution.Identifier).Error())
}

// TestInstructingFINameAlphaNumeric validates InstructingFI Name is alphanumeric
func TestInstructingFINameAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Name = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, bfi.FinancialInstitution.Name).Error())
}

// TestInstructingFIAddressLineOneAlphaNumeric validates InstructingFI AddressLineOne is alphanumeric
func TestInstructingFIAddressLineOneAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineOne = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineOne).Error())
}

// TestInstructingFIAddressLineTwoAlphaNumeric validates InstructingFI AddressLineTwo is alphanumeric
func TestInstructingFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineTwo = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineTwo).Error())
}

// TestInstructingFIAddressLineThreeAlphaNumeric validates InstructingFI AddressLineThree is alphanumeric
func TestInstructingFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Address.AddressLineThree = "®"

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, bfi.FinancialInstitution.Address.AddressLineThree).Error())
}

// TestInstructingFIIdentificationCodeRequired validates InstructingFI IdentificationCode is required
func TestInstructingFIIdentificationCodeRequired(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.IdentificationCode = ""

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestInstructingFIIdentifierRequired validates InstructingFI Identifier is required
func TestInstructingFIIdentifierRequired(t *testing.T) {
	bfi := mockInstructingFI()
	bfi.FinancialInstitution.Identifier = ""

	err := bfi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestParseInstructingFIWrongLength parses a wrong InstructingFI record length
func TestParseInstructingFIWrongLength(t *testing.T) {
	var line = "{5200}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructingFI()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(181, len(r.line))).Error())
}

// TestParseInstructingFIReaderParseError parses a wrong InstructingFI reader parse error
func TestParseInstructingFIReaderParseError(t *testing.T) {
	var line = "{5200}D123456789                         ®I Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInstructingFI()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®I Name")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®I Name")).Error())
}

// TestInstructingFITagError validates a InstructingFI tag
func TestInstructingFITagError(t *testing.T) {
	ifi := mockInstructingFI()
	ifi.tag = "{9999}"

	require.EqualError(t, ifi.Validate(), fieldError("tag", ErrValidTagForType, ifi.tag).Error())
}
