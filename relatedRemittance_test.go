package wire

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// RelatedRemittance creates a RelatedRemittance
func mockRelatedRemittance() *RelatedRemittance {
	rr := NewRelatedRemittance()
	rr.RemittanceIdentification = "Remittance Identification"
	rr.RemittanceLocationMethod = RLMElectronicDataExchange
	rr.RemittanceLocationElectronicAddress = "http://moov.io"
	rr.RemittanceData.Name = "Name"
	rr.RemittanceData.AddressType = CompletePostalAddress
	rr.RemittanceData.Department = "Department"
	rr.RemittanceData.SubDepartment = "Sub-Department"
	rr.RemittanceData.BuildingNumber = "16"
	rr.RemittanceData.PostCode = "19405"
	rr.RemittanceData.TownName = "AnyTown"
	rr.RemittanceData.CountrySubDivisionState = "PA"
	rr.RemittanceData.Country = "UA"
	rr.RemittanceData.AddressLineOne = "Address Line One"
	rr.RemittanceData.AddressLineTwo = "Address Line Two"
	rr.RemittanceData.AddressLineThree = "Address Line Three"
	rr.RemittanceData.AddressLineFour = "Address Line Four"
	rr.RemittanceData.AddressLineFive = "Address Line Five"
	rr.RemittanceData.AddressLineSix = "Address Line Six"
	rr.RemittanceData.AddressLineSeven = "Address Line Seven"
	return rr
}

// TestMockRelatedRemittance validates mockRelatedRemittance
func TestMockRelatedRemittance(t *testing.T) {
	rr := mockRelatedRemittance()

	require.NoError(t, rr.Validate(), "mockRelatedRemittance does not validate and will break other tests")
}

// TestRelatedRemittanceLocationMethodValid validates RelatedRemittance RemittanceLocationMethod
func TestRelatedRemittanceLocationMethodValid(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceLocationMethod = "BBRB"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("RemittanceLocationMethod", ErrRemittanceLocationMethod, rr.RemittanceLocationMethod).Error())
}

// TestRelatedRemittanceLocationElectronicAddressAlphaNumeric validates RelatedRemittance
// RemittanceLocationElectronicAddressAlphaNumeric
func TestRelatedRemittanceLocationElectronicAddressAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceLocationElectronicAddress = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("RemittanceLocationElectronicAddress", ErrNonAlphanumeric, rr.RemittanceLocationElectronicAddress).Error())
}

// TestRelatedRemittanceAddressTypeValid validates RelatedRemittance AddressType
func TestRelatedRemittanceAddressTypeValid(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressType = "BBRB"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressType", ErrAddressType, rr.RemittanceData.AddressType).Error())
}

// TestRelatedRemittanceNameAlphaNumeric validates RelatedRemittance Name is alphanumeric
func TestRelatedRemittanceNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Name = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("Name", ErrNonAlphanumeric, rr.RemittanceData.Name).Error())
}

// TestRelatedRemittanceDepartmentAlphaNumeric validates RelatedRemittance Department is alphanumeric
func TestRelatedRemittanceDepartmentAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Department = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("Department", ErrNonAlphanumeric, rr.RemittanceData.Department).Error())
}

// TestRelatedRemittanceSubDepartmentAlphaNumeric validates RelatedRemittance SubDepartment is alphanumeric
func TestRelatedRemittanceSubDepartmentAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.SubDepartment = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("SubDepartment", ErrNonAlphanumeric, rr.RemittanceData.SubDepartment).Error())
}

// TestRelatedRemittanceStreetNameAlphaNumeric validates RelatedRemittance StreetName is alphanumeric
func TestRelatedRemittanceStreetNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.StreetName = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("StreetName", ErrNonAlphanumeric, rr.RemittanceData.StreetName).Error())
}

// TestRelatedRemittanceBuildingNumberAlphaNumeric validates RelatedRemittance BuildingNumber is alphanumeric
func TestRelatedRemittanceBuildingNumberAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.BuildingNumber = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("BuildingNumber", ErrNonAlphanumeric, rr.RemittanceData.BuildingNumber).Error())
}

// TestRelatedRemittancePostCodeAlphaNumeric validates RelatedRemittance PostCode is alphanumeric
func TestRelatedRemittancePostCodeAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.PostCode = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("PostCode", ErrNonAlphanumeric, rr.RemittanceData.PostCode).Error())
}

// TestRelatedRemittanceTownNameAlphaNumeric validates RelatedRemittance TownName is alphanumeric
func TestRelatedRemittanceTownNameAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.TownName = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("TownName", ErrNonAlphanumeric, rr.RemittanceData.TownName).Error())
}

// TestRelatedRemittanceCountrySubDivisionStateAlphaNumeric validates RelatedRemittance CountrySubDivisionState
// is alphanumeric
func TestRelatedRemittanceCountrySubDivisionStateAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.CountrySubDivisionState = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("CountrySubDivisionState", ErrNonAlphanumeric, rr.RemittanceData.CountrySubDivisionState).Error())
}

// TestRelatedRemittanceCountryAlphaNumeric validates RelatedRemittance Country is alphanumeric
func TestRelatedRemittanceCountryAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Country = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("Country", ErrNonAlphanumeric, rr.RemittanceData.Country).Error())
}

// TestRelatedRemittanceAddressLineOneAlphaNumeric validates RelatedRemittance AddressLineOne is alphanumeric
func TestRelatedRemittanceAddressLineOneAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineOne = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineOne", ErrNonAlphanumeric, rr.RemittanceData.AddressLineOne).Error())
}

// TestRelatedRemittanceAddressLineTwoAlphaNumeric validates RelatedRemittance AddressLineTwo is alphanumeric
func TestRelatedRemittanceAddressLineTwoAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineTwo = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineTwo", ErrNonAlphanumeric, rr.RemittanceData.AddressLineTwo).Error())
}

// TestRelatedRemittanceAddressLineThreeAlphaNumeric validates RelatedRemittance AddressLineThree is alphanumeric
func TestRelatedRemittanceAddressLineThreeAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineThree = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineThree", ErrNonAlphanumeric, rr.RemittanceData.AddressLineThree).Error())
}

// TestRelatedRemittanceAddressLineFourAlphaNumeric validates RelatedRemittance AddressLineFour is alphanumeric
func TestRelatedRemittanceAddressLineFourAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineFour = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineFour", ErrNonAlphanumeric, rr.RemittanceData.AddressLineFour).Error())
}

// TestRelatedRemittanceAddressLineFiveAlphaNumeric validates RelatedRemittance AddressLineFive is alphanumeric
func TestRelatedRemittanceAddressLineFiveAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineFive = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineFive", ErrNonAlphanumeric, rr.RemittanceData.AddressLineFive).Error())
}

// TestRelatedRemittanceAddressLineSixAlphaNumeric validates RelatedRemittance AddressLineSix is alphanumeric
func TestRelatedRemittanceAddressLineSixAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineSix = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineSix", ErrNonAlphanumeric, rr.RemittanceData.AddressLineSix).Error())
}

// TestRelatedRemittanceAddressLineSevenAlphaNumeric validates RelatedRemittance AddressLineSeven is alphanumeric
func TestRelatedRemittanceAddressLineSevenAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.AddressLineSeven = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("AddressLineSeven", ErrNonAlphanumeric, rr.RemittanceData.AddressLineSeven).Error())
}

// TestRelatedRemittanceCountryOfResidenceAlphaNumeric validates RelatedRemittance CountryOfResidence is alphanumeric
func TestRelatedRemittanceCountryOfResidenceAlphaNumeric(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.CountryOfResidence = "®"

	err := rr.Validate()

	require.EqualError(t, err, fieldError("CountryOfResidence", ErrNonAlphanumeric, rr.RemittanceData.CountryOfResidence).Error())
}

// TestRelatedRemittanceNameRequired validates RelatedRemittance Name is required
func TestRelatedRemittanceNameRequired(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.RemittanceData.Name = ""

	err := rr.Validate()

	require.EqualError(t, err, fieldError("Name", ErrFieldRequired).Error())
}

// TestParseRelatedRemittanceWrongLength parses a wrong RelatedRemittance record length
func TestParseRelatedRemittanceWrongLength(t *testing.T) {
	var line = "{8250}Remittance Identification          EDIChttp://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  Name                                                                                                                                        ADDRDepartment                                                            Sub-Department                                                                                                                              16              19405           AnyTown                            PA                                 *UAAddress Line One                                                      *Address Line Two                                                      *Address Line Three                                                    *Address Line Four                                                     *Address Line Five                                                     *Address Line Six                                                      *Address Line Seven                                                  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRelatedRemittance()

	require.EqualError(t, err, r.parseError(fieldError("StreetName", ErrRequireDelimiter)).Error())
}

// TestParseRelatedRemittanceReaderParseError parses a wrong RelatedRemittance reader parse error
func TestParseRelatedRemittanceReaderParseError(t *testing.T) {
	var line = "{8250}Remittance ®dentification          *EDIC*http://moov.io                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  *Name                                                                                                                                        *ADDRDepartment                                                            *Sub-Department                                                                                                                              *16              *19405           *AnyTown                            *PA                                 *UAAddress Line One                                                      *Address Line Two                                                      *Address Line Three                                                    *Address Line Four                                                     *Address Line Five                                                     *Address Line Six                                                      *Address Line Seven                                                   *"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRelatedRemittance()

	require.EqualError(t, err, r.parseError(fieldError("RemittanceIdentification", ErrNonAlphanumeric, "Remittance ®dentification")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("RemittanceIdentification", ErrNonAlphanumeric, "Remittance ®dentification")).Error())
}

// TestRelatedRemittanceTagError validates a RelatedRemittance tag
func TestRelatedRemittanceTagError(t *testing.T) {
	rr := mockRelatedRemittance()
	rr.tag = "{9999}"

	require.EqualError(t, rr.Validate(), fieldError("tag", ErrValidTagForType, rr.tag).Error())
}

// TestStringRelatedRemittanceVariableLength parses using variable length
func TestStringRelatedRemittanceVariableLength(t *testing.T) {
	var line = "{8250}*EDIC**A*ADDR*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRelatedRemittance()
	require.Nil(t, err)

	line = "{8250}                                   EDIC                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                A                                                                                                                                           ADDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRelatedRemittance()
	require.ErrorContains(t, err, ErrRequireDelimiter.Error())

	line = "{8250}*EDIC*A*ADDR***************************"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRelatedRemittance()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{8250}*EDIC**A*ADDR*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseRelatedRemittance()
	require.Equal(t, err, nil)
}

// TestStringRelatedRemittanceOptions validates Format() formatted according to the FormatOptions
func TestStringRelatedRemittanceOptions(t *testing.T) {
	var line = "{8250}*EDIC**A*ADDR*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseRelatedRemittance()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.RelatedRemittance
	require.Equal(t, record.String(), "{8250}                                   *EDIC*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                *A                                                                                                                                           *ADDR*                                                                      *                                                                      *                                                                      *                *                *                                   *                                   *  *                                                                      *                                                                      *                                                                      *                                                                      *                                                                      *                                                                      *                                                                      *")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{8250}*EDIC**A*ADDR*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
