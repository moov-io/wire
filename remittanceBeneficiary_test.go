package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// RemittanceBeneficiary creates a RemittanceBeneficiary
func mockRemittanceBeneficiary() *RemittanceBeneficiary {
	rb := NewRemittanceBeneficiary()
	rb.RemittanceData.Name = "Name"
	rb.IdentificationType = OrganizationID
	rb.IdentificationCode = OICCustomerNumber
	rb.IdentificationNumber = "111111"
	rb.IdentificationNumberIssuer = "Bank"
	rb.RemittanceData.DateBirthPlace = ""
	rb.RemittanceData.AddressType = CompletePostalAddress
	rb.RemittanceData.Department = "Department"
	rb.RemittanceData.SubDepartment = "Sub-Department"
	rb.RemittanceData.StreetName = "Street Name"
	rb.RemittanceData.BuildingNumber = "16"
	rb.RemittanceData.PostCode = "19405"
	rb.RemittanceData.TownName = "AnyTown"
	rb.RemittanceData.CountrySubDivisionState = "PA"
	rb.RemittanceData.Country = "UA"
	rb.RemittanceData.AddressLineOne = "Address Line One"
	rb.RemittanceData.AddressLineTwo = "Address Line Two"
	rb.RemittanceData.AddressLineThree = "Address Line Three"
	rb.RemittanceData.AddressLineFour = "Address Line Four"
	rb.RemittanceData.AddressLineFive = "Address Line Five"
	rb.RemittanceData.AddressLineSix = "Address Line Six"
	rb.RemittanceData.AddressLineSeven = "Address Line Seven"
	rb.RemittanceData.CountryOfResidence = "US"
	return rb
}

// TestMockRemittanceBeneficiary validates mockRemittanceBeneficiary
func TestMockRemittanceBeneficiary(t *testing.T) {
	rb := mockRemittanceBeneficiary()

	require.NoError(t, rb.Validate(), "mockRemittanceBeneficiary does not validate and will break other tests")
}

// TestRemittanceBeneficiaryIdentificationTypeValid validates RemittanceBeneficiary IdentificationType
func TestRemittanceBeneficiaryIdentificationTypeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationType = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationType", ErrIdentificationType, rb.IdentificationType).Error())
}

// TestRemittanceBeneficiaryIdentificationCodeValid validates RemittanceBeneficiary IdentificationCode
func TestRemittanceBeneficiaryIdentificationCodeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrOrganizationIdentificationCode, rb.IdentificationCode).Error())
}

// TestRemittanceBeneficiaryIdentificationCodeValid2 validates RemittanceBeneficiary IdentificationCode
func TestRemittanceBeneficiaryIdentificationCodeValid2(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationType = PrivateID
	rb.IdentificationCode = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrPrivateIdentificationCode, rb.IdentificationCode).Error())
}

// TestRemittanceBeneficiaryAddressTypeValid validates RemittanceBeneficiary AddressType
func TestRemittanceBeneficiaryAddressTypeValid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressType = "BBRB"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressType", ErrAddressType, rb.RemittanceData.AddressType).Error())
}

// TestRemittanceBeneficiaryNameAlphaNumeric validates RemittanceBeneficiary Name is alphanumeric
func TestRemittanceBeneficiaryNameAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Name = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, rb.RemittanceData.Name).Error())
}

// TestRemittanceBeneficiaryIdentificationNumberAlphaNumeric validates RemittanceBeneficiary IdentificationNumber is alphanumeric
func TestRemittanceBeneficiaryIdentificationNumberAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumber = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrNonAlphanumeric, rb.IdentificationNumber).Error())
}

// TestRemittanceBeneficiaryIdentificationNumberIssuerAlphaNumeric validates RemittanceBeneficiary IdentificationNumberIssuer is alphanumeric
func TestRemittanceBeneficiaryIdentificationNumberIssuerAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumberIssuer = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrNonAlphanumeric, rb.IdentificationNumberIssuer).Error())
}

// TestRemittanceBeneficiaryDepartmentAlphaNumeric validates RemittanceBeneficiary Department is alphanumeric
func TestRemittanceBeneficiaryDepartmentAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Department = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("Department", ErrNonAlphanumeric, rb.RemittanceData.Department).Error())
}

// TestRemittanceBeneficiarySubDepartmentAlphaNumeric validates RemittanceBeneficiary SubDepartment is alphanumeric
func TestRemittanceBeneficiarySubDepartmentAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.SubDepartment = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("SubDepartment", ErrNonAlphanumeric, rb.RemittanceData.SubDepartment).Error())
}

// TestRemittanceBeneficiaryStreetNameAlphaNumeric validates RemittanceBeneficiary StreetName is alphanumeric
func TestRemittanceBeneficiaryStreetNameAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.StreetName = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("StreetName", ErrNonAlphanumeric, rb.RemittanceData.StreetName).Error())
}

// TestRemittanceBeneficiaryBuildingNumberAlphaNumeric validates RemittanceBeneficiary BuildingNumber is alphanumeric
func TestRemittanceBeneficiaryBuildingNumberAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.BuildingNumber = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("BuildingNumber", ErrNonAlphanumeric, rb.RemittanceData.BuildingNumber).Error())
}

// TestRemittanceBeneficiaryPostCodeAlphaNumeric validates RemittanceBeneficiary PostCode is alphanumeric
func TestRemittanceBeneficiaryPostCodeAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.PostCode = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("PostCode", ErrNonAlphanumeric, rb.RemittanceData.PostCode).Error())
}

// TestRemittanceBeneficiaryTownNameAlphaNumeric validates RemittanceBeneficiary TownName is alphanumeric
func TestRemittanceBeneficiaryTownNameAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.TownName = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("TownName", ErrNonAlphanumeric, rb.RemittanceData.TownName).Error())
}

// TestRemittanceBeneficiaryCountrySubDivisionStateAlphaNumeric validates RemittanceBeneficiary CountrySubDivisionState
// is alphanumeric
func TestRemittanceBeneficiaryCountrySubDivisionStateAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.CountrySubDivisionState = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("CountrySubDivisionState", ErrNonAlphanumeric, rb.RemittanceData.CountrySubDivisionState).Error())
}

// TestRemittanceBeneficiaryCountryAlphaNumeric validates RemittanceBeneficiary Country is alphanumeric
func TestRemittanceBeneficiaryCountryAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Country = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("Country", ErrNonAlphanumeric, rb.RemittanceData.Country).Error())
}

// TestRemittanceBeneficiaryAddressLineOneAlphaNumeric validates RemittanceBeneficiary AddressLineOne is alphanumeric
func TestRemittanceBeneficiaryAddressLineOneAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineOne = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, rb.RemittanceData.AddressLineOne).Error())
}

// TestRemittanceBeneficiaryAddressLineTwoAlphaNumeric validates RemittanceBeneficiary AddressLineTwo is alphanumeric
func TestRemittanceBeneficiaryAddressLineTwoAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineTwo = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, rb.RemittanceData.AddressLineTwo).Error())
}

// TestRemittanceBeneficiaryAddressLineThreeAlphaNumeric validates RemittanceBeneficiary AddressLineThree is alphanumeric
func TestRemittanceBeneficiaryAddressLineThreeAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineThree = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, rb.RemittanceData.AddressLineThree).Error())
}

// TestRemittanceBeneficiaryAddressLineFourAlphaNumeric validates RemittanceBeneficiary AddressLineFour is alphanumeric
func TestRemittanceBeneficiaryAddressLineFourAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineFour = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineFour", ErrNonAlphanumeric, rb.RemittanceData.AddressLineFour).Error())
}

// TestRemittanceBeneficiaryAddressLineFiveAlphaNumeric validates RemittanceBeneficiary AddressLineFive is alphanumeric
func TestRemittanceBeneficiaryAddressLineFiveAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineFive = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineFive", ErrNonAlphanumeric, rb.RemittanceData.AddressLineFive).Error())
}

// TestRemittanceBeneficiaryAddressLineSixAlphaNumeric validates RemittanceBeneficiary AddressLineSix is alphanumeric
func TestRemittanceBeneficiaryAddressLineSixAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineSix = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineSix", ErrNonAlphanumeric, rb.RemittanceData.AddressLineSix).Error())
}

// TestRemittanceBeneficiaryAddressLineSevenAlphaNumeric validates RemittanceBeneficiary AddressLineSeven is alphanumeric
func TestRemittanceBeneficiaryAddressLineSevenAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.AddressLineSeven = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("AddressLineSeven", ErrNonAlphanumeric, rb.RemittanceData.AddressLineSeven).Error())
}

// TestRemittanceBeneficiaryCountryOfResidenceAlphaNumeric validates RemittanceBeneficiary CountryOfResidence is alphanumeric
func TestRemittanceBeneficiaryCountryOfResidenceAlphaNumeric(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.CountryOfResidence = "®"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("CountryOfResidence", ErrNonAlphanumeric, rb.RemittanceData.CountryOfResidence).Error())
}

// TestRemittanceBeneficiaryNameRequired validates RemittanceBeneficiary Name is required
func TestRemittanceBeneficiaryNameRequired(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.RemittanceData.Name = ""

	err := rb.Validate()

	require.EqualError(t, err, fieldError("Name", ErrFieldRequired).Error())
}

// TestRemittanceBeneficiaryIdentificationNumberInvalid validates RemittanceBeneficiary IdentificationNumber
func TestRemittanceBeneficiaryIdentificationNumberInvalid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICDateBirthPlace
	rb.IdentificationNumber = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrInvalidProperty, rb.IdentificationNumber).Error())
}

// TestIdentificationNumberIssuerInvalid_IdentificationNumber validates RemittanceBeneficiary IdentificationNumberIssuer
func TestIdentificationNumberIssuerInvalid_IdentificationNumber(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationNumber = ""
	rb.IdentificationNumberIssuer = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrInvalidProperty, rb.IdentificationNumberIssuer).Error())
}

// TestIdentificationNumberIssuerInvalid_PICDateBirthPlace validates RemittanceBeneficiary IdentificationNumber
func TestIdentificationNumberIssuerInvalid_PICDateBirthPlace(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICDateBirthPlace
	rb.IdentificationNumberIssuer = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrInvalidProperty, rb.IdentificationNumber).Error())
}

// TestIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI validates RemittanceBeneficiary IdentificationNumberIssuer
func TestIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = OICSWIFTBICORBEI
	rb.IdentificationNumberIssuer = "zz"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrInvalidProperty, rb.IdentificationNumberIssuer).Error())
}

// TestRemittanceBeneficiaryDateBirthPlaceInvalid validates RemittanceBeneficiary DateBirthPlace
func TestRemittanceBeneficiaryDateBirthPlaceInvalid(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.IdentificationCode = PICCustomerNumber
	rb.RemittanceData.DateBirthPlace = "Pottstown"

	err := rb.Validate()

	require.EqualError(t, err, fieldError("DateBirthPlace", ErrInvalidProperty, rb.RemittanceData.DateBirthPlace).Error())
}

// TestParseRemittanceBeneficiaryWrongLength parses a wrong RemittanceBeneficiary record length
func TestParseRemittanceBeneficiaryWrongLength(t *testing.T) {
	var line = "{8350}Name                                                                                                                                        OICUST111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                 US"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceBeneficiary()

	require.EqualError(t, err, r.parseError(fieldError("AddressLineSeven", ErrValidLength)).Error())
}

// TestParseRemittanceBeneficiaryReaderParseError parses a wrong RemittanceBeneficiary reader parse error
func TestParseRemittanceBeneficiaryReaderParseError(t *testing.T) {
	var line = "{8350}®ame                                                                                                                                        OICUST111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                   US"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceBeneficiary()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®ame")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("Name", ErrNonAlphanumeric, "®ame")).Error())
}

// TestRemittanceBeneficiaryTagError validates a RemittanceBeneficiary tag
func TestRemittanceBeneficiaryTagError(t *testing.T) {
	rb := mockRemittanceBeneficiary()
	rb.tag = "{9999}"

	require.EqualError(t, rb.Validate(), fieldError("tag", ErrValidTagForType, rb.tag).Error())
}

// TestStringRemittanceBeneficiaryVariableLength parses using variable length
func TestStringRemittanceBeneficiaryVariableLength(t *testing.T) {
	var line = "{8350}Name*PIARNU****ADDR*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceBeneficiary()
	require.Nil(t, err)

	line = "{8350}Name                                                                                                                                        PIARNU                                                                                                                                                        ADDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8350}Name*PIARNU***ADDR****************************"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceBeneficiary()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8350}Name*PIARNU****ADDR*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceBeneficiary()
	require.Equal(t, err, nil)
}

// TestStringRemittanceBeneficiaryOptions validates Format() formatted according to the FormatOptions
func TestStringRemittanceBeneficiaryOptions(t *testing.T) {
	var line = "{8350}Name*PIARNU****ADDR"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceBeneficiary()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.RemittanceBeneficiary
	require.Equal(t, record.String(), "{8350}Name                                                                                                                                        PIARNU                                                                                                                                                        ADDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{8350}Name*PIARNU***ADDR*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
