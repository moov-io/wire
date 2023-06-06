package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// RemittanceOriginator creates a RemittanceOriginator
func mockRemittanceOriginator() *RemittanceOriginator {
	ro := NewRemittanceOriginator()
	ro.IdentificationType = OrganizationID
	ro.IdentificationCode = OICCustomerNumber
	ro.IdentificationNumber = "111111"
	ro.IdentificationNumberIssuer = "Bank"
	// ro.RemittanceData.DateBirthPlace = "12072008 AnyTown"
	ro.RemittanceData.Name = "Name"
	ro.RemittanceData.AddressType = CompletePostalAddress
	ro.RemittanceData.Department = "Department"
	ro.RemittanceData.SubDepartment = "Sub-Department"
	ro.RemittanceData.StreetName = "Street Name"
	ro.RemittanceData.BuildingNumber = "16"
	ro.RemittanceData.PostCode = "19405"
	ro.RemittanceData.TownName = "AnyTown"
	ro.RemittanceData.CountrySubDivisionState = "PA"
	ro.RemittanceData.Country = "UA"
	ro.RemittanceData.AddressLineOne = "Address Line One"
	ro.RemittanceData.AddressLineTwo = "Address Line Two"
	ro.RemittanceData.AddressLineThree = "Address Line Three"
	ro.RemittanceData.AddressLineFour = "Address Line Four"
	ro.RemittanceData.AddressLineFive = "Address Line Five"
	ro.RemittanceData.AddressLineSix = "Address Line Six"
	ro.RemittanceData.AddressLineSeven = "Address Line Seven"
	ro.RemittanceData.CountryOfResidence = "US"
	ro.ContactName = "Contact Name"
	ro.ContactPhoneNumber = "5551231212"
	ro.ContactMobileNumber = "5551231212"
	ro.ContactFaxNumber = "5551231212"
	ro.ContactElectronicAddress = "http://www.moov.io"
	ro.ContactOther = "Contact Other"
	return ro
}

// TestMockRemittanceOriginator validates mockRemittanceOriginator
func TestMockRemittanceOriginator(t *testing.T) {
	ro := mockRemittanceOriginator()

	require.NoError(t, ro.Validate(), "mockRemittanceOriginator does not validate and will break other tests")
}

// TestRemittanceOriginatorIdentificationTypeValid validates RemittanceOriginator IdentificationType
func TestRemittanceOriginatorIdentificationTypeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationType = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationType", ErrIdentificationType, ro.IdentificationType).Error())
}

// TestRemittanceOriginatorIdentificationCodeValid validates RemittanceOriginator IdentificationCode
func TestRemittanceOriginatorIdentificationCodeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrOrganizationIdentificationCode, ro.IdentificationCode).Error())
}

// TestRemittanceOriginatorIdentificationCodeValid2 validates RemittanceOriginator IdentificationCode
func TestRemittanceOriginatorIdentificationCodeValid2(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationType = PrivateID
	ro.IdentificationCode = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationCode", ErrPrivateIdentificationCode, ro.IdentificationCode).Error())
}

// TestRemittanceOriginatorAddressTypeValid validates RemittanceOriginator AddressType
func TestRemittanceOriginatorAddressTypeValid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressType = "BBRB"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressType", ErrAddressType, ro.RemittanceData.AddressType).Error())
}

// TestRemittanceOriginatorNameAlphaNumeric validates RemittanceOriginator Name is alphanumeric
func TestRemittanceOriginatorNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Name = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, ro.RemittanceData.Name).Error())
}

// TestRemittanceOriginatorIdentificationNumberAlphaNumeric validates RemittanceOriginator IdentificationNumber is alphanumeric
func TestRemittanceOriginatorIdentificationNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumber = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrNonAlphanumeric, ro.IdentificationNumber).Error())
}

// TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric validates RemittanceOriginator IdentificationNumberIssuer is alphanumeric
func TestRemittanceOriginatorIdentificationNumberIssuerAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumberIssuer = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrNonAlphanumeric, ro.IdentificationNumberIssuer).Error())
}

// TestRemittanceOriginatorDepartmentAlphaNumeric validates RemittanceOriginator Department is alphanumeric
func TestRemittanceOriginatorDepartmentAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Department = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("Department", ErrNonAlphanumeric, ro.RemittanceData.Department).Error())
}

// TestRemittanceOriginatorSubDepartmentAlphaNumeric validates RemittanceOriginator SubDepartment is alphanumeric
func TestRemittanceOriginatorSubDepartmentAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.SubDepartment = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("SubDepartment", ErrNonAlphanumeric, ro.RemittanceData.SubDepartment).Error())
}

// TestRemittanceOriginatorStreetNameAlphaNumeric validates RemittanceOriginator StreetName is alphanumeric
func TestRemittanceOriginatorStreetNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.StreetName = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("StreetName", ErrNonAlphanumeric, ro.RemittanceData.StreetName).Error())
}

// TestRemittanceOriginatorBuildingNumberAlphaNumeric validates RemittanceOriginator BuildingNumber is alphanumeric
func TestRemittanceOriginatorBuildingNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.BuildingNumber = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("BuildingNumber", ErrNonAlphanumeric, ro.RemittanceData.BuildingNumber).Error())
}

// TestRemittanceOriginatorPostCodeAlphaNumeric validates RemittanceOriginator PostCode is alphanumeric
func TestRemittanceOriginatorPostCodeAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.PostCode = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("PostCode", ErrNonAlphanumeric, ro.RemittanceData.PostCode).Error())
}

// TestRemittanceOriginatorTownNameAlphaNumeric validates RemittanceOriginator TownName is alphanumeric
func TestRemittanceOriginatorTownNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.TownName = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("TownName", ErrNonAlphanumeric, ro.RemittanceData.TownName).Error())
}

// TestRemittanceOriginatorCountrySubDivisionStateAlphaNumeric validates RemittanceOriginator CountrySubDivisionState
// is alphanumeric
func TestRemittanceOriginatorCountrySubDivisionStateAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountrySubDivisionState = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("CountrySubDivisionState", ErrNonAlphanumeric, ro.RemittanceData.CountrySubDivisionState).Error())
}

// TestRemittanceOriginatorCountryAlphaNumeric validates RemittanceOriginator Country is alphanumeric
func TestRemittanceOriginatorCountryAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Country = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("Country", ErrNonAlphanumeric, ro.RemittanceData.Country).Error())
}

// TestRemittanceOriginatorAddressLineOneAlphaNumeric validates RemittanceOriginator AddressLineOne is alphanumeric
func TestRemittanceOriginatorAddressLineOneAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineOne = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, ro.RemittanceData.AddressLineOne).Error())
}

// TestRemittanceOriginatorAddressLineTwoAlphaNumeric validates RemittanceOriginator AddressLineTwo is alphanumeric
func TestRemittanceOriginatorAddressLineTwoAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineTwo = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, ro.RemittanceData.AddressLineTwo).Error())
}

// TestRemittanceOriginatorAddressLineThreeAlphaNumeric validates RemittanceOriginator AddressLineThree is alphanumeric
func TestRemittanceOriginatorAddressLineThreeAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineThree = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, ro.RemittanceData.AddressLineThree).Error())
}

// TestRemittanceOriginatorAddressLineFourAlphaNumeric validates RemittanceOriginator AddressLineFour is alphanumeric
func TestRemittanceOriginatorAddressLineFourAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFour = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineFour", ErrNonAlphanumeric, ro.RemittanceData.AddressLineFour).Error())
}

// TestRemittanceOriginatorAddressLineFiveAlphaNumeric validates RemittanceOriginator AddressLineFive is alphanumeric
func TestRemittanceOriginatorAddressLineFiveAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineFive = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineFive", ErrNonAlphanumeric, ro.RemittanceData.AddressLineFive).Error())
}

// TestRemittanceOriginatorAddressLineSixAlphaNumeric validates RemittanceOriginator AddressLineSix is alphanumeric
func TestRemittanceOriginatorAddressLineSixAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSix = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineSix", ErrNonAlphanumeric, ro.RemittanceData.AddressLineSix).Error())
}

// TestRemittanceOriginatorAddressLineSevenAlphaNumeric validates RemittanceOriginator AddressLineSeven is alphanumeric
func TestRemittanceOriginatorAddressLineSevenAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.AddressLineSeven = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("AddressLineSeven", ErrNonAlphanumeric, ro.RemittanceData.AddressLineSeven).Error())
}

// TestRemittanceOriginatorCountryOfResidenceAlphaNumeric validates RemittanceOriginator CountryOfResidence is alphanumeric
func TestRemittanceOriginatorCountryOfResidenceAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.CountryOfResidence = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("CountryOfResidence", ErrNonAlphanumeric, ro.RemittanceData.CountryOfResidence).Error())
}

// TestRemittanceOriginatorContactNameAlphaNumeric validates RemittanceOriginator ContactName is alphanumeric
func TestRemittanceOriginatorContactNameAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactName = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactName", ErrNonAlphanumeric, ro.ContactName).Error())
}

// TestRemittanceOriginatorContactPhoneNumberAlphaNumeric validates RemittanceOriginator ContactPhoneNumber is alphanumeric
func TestRemittanceOriginatorContactPhoneNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactPhoneNumber = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactPhoneNumber", ErrNonAlphanumeric, ro.ContactPhoneNumber).Error())
}

// TestRemittanceOriginatorContactMobileNumberAlphaNumeric validates RemittanceOriginator ContactMobileNumber is alphanumeric
func TestRemittanceOriginatorContactMobileNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactMobileNumber = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactMobileNumber", ErrNonAlphanumeric, ro.ContactMobileNumber).Error())
}

// TestRemittanceOriginatorContactFaxNumberAlphaNumeric validates RemittanceOriginator ContactFaxNumber is alphanumeric
func TestRemittanceOriginatorContactFaxNumberAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactFaxNumber = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactFaxNumber", ErrNonAlphanumeric, ro.ContactFaxNumber).Error())
}

// TestRemittanceOriginatorContactElectronicAddressAlphaNumeric validates RemittanceOriginator ContactElectronicAddress
// is alphanumeric
func TestRemittanceOriginatorContactElectronicAddressAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactElectronicAddress = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactElectronicAddress", ErrNonAlphanumeric, ro.ContactElectronicAddress).Error())
}

// TestRemittanceOriginatorContactOtherAlphaNumeric validates RemittanceOriginator ContactOther
// is alphanumeric
func TestRemittanceOriginatorContactOtherAlphaNumeric(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.ContactOther = "®"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("ContactOther", ErrNonAlphanumeric, ro.ContactOther).Error())
}

// TestRemittanceOriginatorNameRequired validates RemittanceOriginator Name is required
func TestRemittanceOriginatorNameRequired(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.RemittanceData.Name = ""

	err := ro.Validate()

	require.EqualError(t, err, fieldError("Name", ErrFieldRequired).Error())
}

// TestRemittanceOriginatorIdentificationNumberInvalid validates RemittanceOriginator IdentificationNumber
func TestRemittanceOriginatorIdentificationNumberInvalid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICDateBirthPlace
	ro.IdentificationNumber = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrInvalidProperty, ro.IdentificationNumber).Error())
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_IdentificationNumber validates RemittanceOriginator IdentificationNumberIssuer
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_IdentificationNumber(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationNumber = ""
	ro.IdentificationNumberIssuer = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrInvalidProperty, ro.IdentificationNumberIssuer).Error())
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_PICDateBirthPlace validates RemittanceOriginator IdentificationNumber
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_PICDateBirthPlace(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICDateBirthPlace
	ro.IdentificationNumberIssuer = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumber", ErrInvalidProperty, ro.IdentificationNumber).Error())
}

// TestRemittanceOriginatorIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI validates RemittanceOriginator IdentificationNumberIssuer
func TestRemittanceOriginatorIdentificationNumberIssuerInvalid_OICSWIFTBICORBEI(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = OICSWIFTBICORBEI
	ro.IdentificationNumberIssuer = "zz"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("IdentificationNumberIssuer", ErrInvalidProperty, ro.IdentificationNumberIssuer).Error())
}

// TestRemittanceOriginatorDateBirthPlaceInvalid validates RemittanceOriginator DateBirthPlace
func TestRemittanceOriginatorDateBirthPlaceInvalid(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.IdentificationCode = PICCustomerNumber
	ro.RemittanceData.DateBirthPlace = "Pottstown"

	err := ro.Validate()

	require.EqualError(t, err, fieldError("DateBirthPlace", ErrInvalidProperty, ro.RemittanceData.DateBirthPlace).Error())
}

// TestParseRemittanceOriginatorWrongLength parses a wrong RemittanceOriginator record length
func TestParseRemittanceOriginatorWrongLength(t *testing.T) {
	var line = "{8300}OICUSTName                                                                                                                                        111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UAAddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                    USContact Name                                                                                                                                5551231212                         5551231212                         5551231212                         http://www.moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Contact Other                    "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceOriginator()

	require.EqualError(t, err, r.parseError(fieldError("ContactOther", ErrValidLength)).Error())
}

// TestParseRemittanceOriginatorReaderParseError parses a wrong RemittanceOriginator reader parse error
func TestParseRemittanceOriginatorReaderParseError(t *testing.T) {
	var line = "{8300}OICUSTName                                                                                                                                        111111                             Bank                                                                                                                 ADDRDepartment                                                            Sub-Department                                                        Street Name                                                           16              19405           AnyTown                            PA                                 UA®ddress Line One                                                      Address Line Two                                                      Address Line Three                                                    Address Line Four                                                     Address Line Five                                                     Address Line Six                                                      Address Line Seven                                                    USContact Name                                                                                                                                5551231212                         5551231212                         5551231212                         http://www.moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              Contact Other                     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceOriginator()

	require.EqualError(t, err, r.parseError(fieldError("AddressLineOne", ErrNonAlphanumeric, "®ddress Line One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("AddressLineOne", ErrNonAlphanumeric, "®ddress Line One")).Error())
}

// TestRemittanceOriginatorTagError validates a RemittanceOriginator tag
func TestRemittanceOriginatorTagError(t *testing.T) {
	ro := mockRemittanceOriginator()
	ro.tag = "{9999}"

	require.EqualError(t, ro.Validate(), fieldError("tag", ErrValidTagForType, ro.tag).Error())
}

// TestStringRemittanceOriginatorVariableLength parses using variable length
func TestStringRemittanceOriginatorVariableLength(t *testing.T) {
	var line = "{8300}OICUSTName****ADDR"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceOriginator()
	require.Nil(t, err)

	line = "{8300}OICUSTName                                                                                                                                                                                                                                                                                                ADDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceOriginator()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8300}OICUSTName****ADDR*****************************"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceOriginator()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8300}OICUSTName****ADDR*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRemittanceOriginator()
	require.Equal(t, err, nil)
}

// TestStringRemittanceOriginatorOptions validates Format() formatted according to the FormatOptions
func TestStringRemittanceOriginatorOptions(t *testing.T) {
	var line = "{8300}OICUSTName****ADDR"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRemittanceOriginator()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.RemittanceOriginator
	require.Equal(t, record.String(), "{8300}OICUSTName                                                                                                                                                                                                                                                                                                ADDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{8300}OICUSTName****ADDR*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
