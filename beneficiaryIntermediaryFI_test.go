package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockBeneficiaryIntermediaryFI creates a BeneficiaryIntermediaryFI
func mockBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	bifi := NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bifi
}

// TestMockBeneficiaryIntermediaryFI validates mockBeneficiaryIntermediaryFI
func TestMockBeneficiaryIntermediaryFI(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()

	require.NoError(t, bifi.Validate(), "mockBeneficiaryIntermediaryFI does not validate and will break other tests")
}

// TestBeneficiaryIntermediaryFIIdentificationCodeValid validates BeneficiaryIntermediaryFI IdentificationCode
func TestBeneficiaryIntermediaryFIIdentificationCodeValid(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = "Football Card ID"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIIdentificationCodeFI validates BeneficiaryIntermediaryFI IdentificationCode is an FI code
func TestBeneficiaryIntermediaryFIIdentificationCodeFI(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = "1"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric validates BeneficiaryIntermediaryFI Identifier is alphanumeric
func TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = "®"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Identifier", ErrNonAlphanumeric, bifi.FinancialInstitution.Identifier).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFINameAlphaNumeric validates BeneficiaryIntermediaryFI Name is alphanumeric
func TestBeneficiaryIntermediaryFINameAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Name = "®"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("Name", ErrNonAlphanumeric, bifi.FinancialInstitution.Name).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineOne is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineOne = "®"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AddressLineOne", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineOne).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineTwo is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineTwo = "®"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AddressLineTwo", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineTwo).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineThree is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineThree = "®"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("AddressLineThree", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineThree).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIIdentificationCodeRequired validates BeneficiaryIntermediaryFI IdentificationCode is required
func TestBeneficiaryIntermediaryFIIdentificationCodeRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = ""

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.IdentificationCode", ErrFieldRequired).Error(), err.Error())
}

// TestBeneficiaryIntermediaryFIIdentifierRequired validates BeneficiaryIntermediaryFI Identifier is required
func TestBeneficiaryIntermediaryFIIdentifierRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = ""

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("BeneficiaryIntermediaryFI.FinancialInstitution.Identifier", ErrFieldRequired).Error(), err.Error())
}

// TestParseBeneficiaryIntermediaryFIWrongLength parses a wrong BeneficiaryIntermediaryFI record length
func TestParseBeneficiaryIntermediaryFIWrongLength(t *testing.T) {
	var line = "{4000}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryIntermediaryFI()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(181, len(r.line)).Error())
}

// TestParseBeneficiaryIntermediaryFIReaderParseError parses a wrong BeneficiaryIntermediaryFI reader parse error
func TestParseBeneficiaryIntermediaryFIReaderParseError(t *testing.T) {
	var line = "{4000}D123456789                         F® Name                            Address One                        Address Two                        Address Three                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.SetBeneficiaryIntermediaryFI(bifi)

	err := r.parseBeneficiaryIntermediaryFI()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrNonAlphanumeric.Error())
}

// TestBeneficiaryIntermediaryFITagError validates a BeneficiaryFI tag
func TestBeneficiaryIntermediaryFITagError(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.tag = "{9999}"

	err := bifi.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, bifi.tag).Error(), err.Error())
}
