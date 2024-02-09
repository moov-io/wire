package wire

import (
	"errors"
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

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode).Error())
}

// TestBeneficiaryIntermediaryFIIdentificationCodeFI validates BeneficiaryIntermediaryFI IdentificationCode is an FI code
func TestBeneficiaryIntermediaryFIIdentificationCodeFI(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = "1"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrIdentificationCode, bifi.FinancialInstitution.IdentificationCode).Error())
}

// TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric validates BeneficiaryIntermediaryFI Identifier is alphanumeric
func TestBeneficiaryIntermediaryFIIdentifierAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = "®"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrNonAlphanumeric, bifi.FinancialInstitution.Identifier).Error())
}

// TestBeneficiaryIntermediaryFINameAlphaNumeric validates BeneficiaryIntermediaryFI Name is alphanumeric
func TestBeneficiaryIntermediaryFINameAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Name = "®"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, bifi.FinancialInstitution.Name).Error())
}

// TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineOne is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineOneAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineOne = "®"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineOne).Error())
}

// TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineTwo is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineTwoAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineTwo = "®"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineTwo).Error())
}

// TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric validates BeneficiaryIntermediaryFI AddressLineThree is alphanumeric
func TestBeneficiaryIntermediaryFIAddressLineThreeAlphaNumeric(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Address.AddressLineThree = "®"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, bifi.FinancialInstitution.Address.AddressLineThree).Error())
}

// TestBeneficiaryIntermediaryFIIdentificationCodeRequired validates BeneficiaryIntermediaryFI IdentificationCode is required
func TestBeneficiaryIntermediaryFIIdentificationCodeRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = ""

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrFieldRequired).Error())
}

// TestBeneficiaryIntermediaryFIIdentifierRequired validates BeneficiaryIntermediaryFI Identifier is required
func TestBeneficiaryIntermediaryFIIdentifierRequired(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.Identifier = ""

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("Identifier", ErrFieldRequired).Error())
}

// TestParseBeneficiaryIntermediaryFIWrongLength parses a wrong BeneficiaryIntermediaryFI record length
func TestParseBeneficiaryIntermediaryFIWrongLength(t *testing.T) {
	var line = "{4000}D123456789                         FI Name                            Address One                        Address Two                        Address Three                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryIntermediaryFI()

	require.EqualError(t, err, r.parseError(fieldError("Identifier", ErrRequireDelimiter)).Error())
}

// TestParseBeneficiaryIntermediaryFIReaderParseError parses a wrong BeneficiaryIntermediaryFI reader parse error
func TestParseBeneficiaryIntermediaryFIReaderParseError(t *testing.T) {
	var line = "{4000}D123456789                         *F® Name                            *Address One                        *Address Two                        *Address Three                     *"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	bifi := mockBeneficiaryIntermediaryFI()
	fwm.BeneficiaryIntermediaryFI = bifi

	err := r.parseBeneficiaryIntermediaryFI()

	expected := r.parseError(fieldError("Name", ErrNonAlphanumeric, "F® Name")).Error()
	require.EqualError(t, err, expected)

	_, err = r.Read()

	expected = r.parseError(fieldError("Name", ErrNonAlphanumeric, "F® Name")).Error()
	require.EqualError(t, err, expected)
}

// TestBeneficiaryIntermediaryFITagError validates a BeneficiaryFI tag
func TestBeneficiaryIntermediaryFITagError(t *testing.T) {
	bifi := mockBeneficiaryIntermediaryFI()
	bifi.tag = "{9999}"

	err := bifi.Validate()

	require.EqualError(t, err, fieldError("tag", ErrValidTagForType, bifi.tag).Error())
}

// TestStringBeneficiaryIntermediaryFIVariableLength parses using variable length
func TestStringBeneficiaryIntermediaryFIVariableLength(t *testing.T) {
	var line = "{4000}"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryIntermediaryFI()
	expected := r.parseError(NewTagMinLengthErr(7, len(r.line))).Error()
	require.EqualError(t, err, expected)

	line = "{4000}D123456789                         *FI Name                            *Address One                        *Address Two                        *Address Three                    NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryIntermediaryFI()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{4000}D123456789*******"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryIntermediaryFI()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{4000}D123456789****"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseBeneficiaryIntermediaryFI()
	require.Equal(t, err, nil)
}

// TestStringBeneficiaryIntermediaryFIOptions validates Format() formatted according to the FormatOptions
func TestStringBeneficiaryIntermediaryFIOptions(t *testing.T) {
	var line = "{4000}D123456789*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseBeneficiaryIntermediaryFI()
	require.Equal(t, err, nil)

	bifi := r.currentFEDWireMessage.BeneficiaryIntermediaryFI
	require.Equal(t, bifi.String(), "{4000}D123456789                         *                                   *                                   *                                   *                                   *")
	require.Equal(t, bifi.Format(FormatOptions{VariableLengthFields: true}), "{4000}D123456789*")
	require.Equal(t, bifi.String(), bifi.Format(FormatOptions{VariableLengthFields: false}))
}
