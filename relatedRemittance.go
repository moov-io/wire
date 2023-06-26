// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// RelatedRemittance is related remittance
type RelatedRemittance struct {
	// tag
	tag string
	// RemittanceIdentification is remittance identification
	RemittanceIdentification string `json:"remittanceIdentification,omitempty"`
	// RemittanceLocationMethod is  remittance location method
	RemittanceLocationMethod string `json:"remittanceLocationMethod,omitempty"`
	// RemittanceLocationElectronicAddress (E-mail or URL address)
	RemittanceLocationElectronicAddress string `json:"remittanceLocationElctronicAddress,omitempty"`
	// RemittanceData is RemittanceData
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRelatedRemittance returns a new RelatedRemittance
func NewRelatedRemittance() *RelatedRemittance {
	rr := &RelatedRemittance{
		tag: TagRelatedRemittance,
	}
	return rr
}

// Parse takes the input string and parses the RelatedRemittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rr *RelatedRemittance) Parse(record string) error {
	if utf8.RuneCountInString(record) < 18 {
		return NewTagMinLengthErr(18, len(record))
	}

	rr.tag = record[:6]
	length := 6

	value, read, err := rr.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("RemittanceIdentification", err)
	}
	rr.RemittanceIdentification = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("RemittanceLocationMethod", err)
	}
	rr.RemittanceLocationMethod = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 2048)
	if err != nil {
		return fieldError("RemittanceLocationElectronicAddress", err)
	}
	rr.RemittanceLocationElectronicAddress = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("RemittanceData", err)
	}
	rr.RemittanceData.Name = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("AddressType", err)
	}
	rr.RemittanceData.AddressType = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("Department", err)
	}
	rr.RemittanceData.Department = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("SubDepartment", err)
	}
	rr.RemittanceData.SubDepartment = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("StreetName", err)
	}
	rr.RemittanceData.StreetName = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("BuildingNumber", err)
	}
	rr.RemittanceData.BuildingNumber = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("PostCode", err)
	}
	rr.RemittanceData.PostCode = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("TownName", err)
	}
	rr.RemittanceData.TownName = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("CountrySubDivisionState", err)
	}
	rr.RemittanceData.CountrySubDivisionState = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("Country", err)
	}
	rr.RemittanceData.Country = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	rr.RemittanceData.AddressLineOne = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	rr.RemittanceData.AddressLineTwo = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	rr.RemittanceData.AddressLineThree = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFour", err)
	}
	rr.RemittanceData.AddressLineFour = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFive", err)
	}
	rr.RemittanceData.AddressLineFive = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSix", err)
	}
	rr.RemittanceData.AddressLineSix = value
	length += read

	value, read, err = rr.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSeven", err)
	}
	rr.RemittanceData.AddressLineSeven = value
	length += read

	if err := rr.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
	}

	return nil
}

func (rr *RelatedRemittance) UnmarshalJSON(data []byte) error {
	type Alias RelatedRemittance
	aux := struct {
		*Alias
	}{
		(*Alias)(rr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rr.tag = TagRelatedRemittance
	return nil
}

// String returns a fixed-width RelatedRemittance record
func (rr *RelatedRemittance) String() string {
	return rr.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a RelatedRemittance record formatted according to the FormatOptions
func (rr *RelatedRemittance) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(3041)

	buf.WriteString(rr.tag)
	buf.WriteString(rr.FormatRemittanceIdentification(options) + Delimiter)
	buf.WriteString(rr.FormatRemittanceLocationMethod(options) + Delimiter)
	buf.WriteString(rr.FormatRemittanceLocationElectronicAddress(options) + Delimiter)
	buf.WriteString(rr.FormatName(options) + Delimiter)
	buf.WriteString(rr.FormatAddressType(options) + Delimiter)
	buf.WriteString(rr.FormatDepartment(options) + Delimiter)
	buf.WriteString(rr.FormatSubDepartment(options) + Delimiter)
	buf.WriteString(rr.FormatStreetName(options) + Delimiter)
	buf.WriteString(rr.FormatBuildingNumber(options) + Delimiter)
	buf.WriteString(rr.FormatPostCode(options) + Delimiter)
	buf.WriteString(rr.FormatTownName(options) + Delimiter)
	buf.WriteString(rr.FormatCountrySubDivisionState(options) + Delimiter)
	buf.WriteString(rr.FormatCountry(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineOne(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineTwo(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineThree(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineFour(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineFive(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineSix(options) + Delimiter)
	buf.WriteString(rr.FormatAddressLineSeven(options) + Delimiter)

	if options.VariableLengthFields {
		return rr.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on RelatedRemittance and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rr *RelatedRemittance) Validate() error {
	if rr.tag != TagRelatedRemittance {
		return fieldError("tag", ErrValidTagForType, rr.tag)
	}
	if err := rr.fieldInclusion(); err != nil {
		return err
	}
	if err := rr.isAlphanumeric(rr.RemittanceIdentification); err != nil {
		return fieldError("RemittanceIdentification", err, rr.RemittanceIdentification)
	}
	if err := rr.isRemittanceLocationMethod(rr.RemittanceLocationMethod); err != nil {
		return fieldError("RemittanceLocationMethod", err, rr.RemittanceLocationMethod)
	}
	if err := rr.isAlphanumeric(rr.RemittanceLocationElectronicAddress); err != nil {
		return fieldError("RemittanceLocationElectronicAddress", err, rr.RemittanceLocationElectronicAddress)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Name); err != nil {
		return fieldError("Name", err, rr.RemittanceData.Name)
	}
	if err := rr.isAddressType(rr.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, rr.RemittanceData.AddressType)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Department); err != nil {
		return fieldError("Department", err, rr.RemittanceData.Department)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, rr.RemittanceData.SubDepartment)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, rr.RemittanceData.StreetName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, rr.RemittanceData.BuildingNumber)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, rr.RemittanceData.PostCode)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, rr.RemittanceData.TownName)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, rr.RemittanceData.CountrySubDivisionState)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.Country); err != nil {
		return fieldError("Country", err, rr.RemittanceData.Country)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, rr.RemittanceData.AddressLineOne)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, rr.RemittanceData.AddressLineTwo)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, rr.RemittanceData.AddressLineThree)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, rr.RemittanceData.AddressLineFour)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, rr.RemittanceData.AddressLineFive)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, rr.RemittanceData.AddressLineSix)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, rr.RemittanceData.AddressLineSeven)
	}
	if err := rr.isAlphanumeric(rr.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, rr.RemittanceData.CountryOfResidence)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rr *RelatedRemittance) fieldInclusion() error {
	if rr.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}
	return nil
}

// RemittanceIdentificationField gets a string of the RemittanceIdentification field
func (rr *RelatedRemittance) RemittanceIdentificationField() string {
	return rr.alphaField(rr.RemittanceIdentification, 35)
}

// RemittanceLocationMethodField gets a string of the RemittanceLocationMethod field
func (rr *RelatedRemittance) RemittanceLocationMethodField() string {
	return rr.alphaField(rr.RemittanceLocationMethod, 4)
}

// RemittanceLocationElectronicAddressField gets a string of the RemittanceLocationElectronicAddress field
func (rr *RelatedRemittance) RemittanceLocationElectronicAddressField() string {
	return rr.alphaField(rr.RemittanceLocationElectronicAddress, 2048)
}

// NameField gets a string of the Name field
func (rr *RelatedRemittance) NameField() string {
	return rr.alphaField(rr.RemittanceData.Name, 140)
}

// AddressTypeField gets a string of the AddressType field
func (rr *RelatedRemittance) AddressTypeField() string {
	return rr.alphaField(rr.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (rr *RelatedRemittance) DepartmentField() string {
	return rr.alphaField(rr.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (rr *RelatedRemittance) SubDepartmentField() string {
	return rr.alphaField(rr.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (rr *RelatedRemittance) StreetNameField() string {
	return rr.alphaField(rr.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rr *RelatedRemittance) BuildingNumberField() string {
	return rr.alphaField(rr.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (rr *RelatedRemittance) PostCodeField() string {
	return rr.alphaField(rr.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (rr *RelatedRemittance) TownNameField() string {
	return rr.alphaField(rr.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rr *RelatedRemittance) CountrySubDivisionStateField() string {
	return rr.alphaField(rr.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (rr *RelatedRemittance) CountryField() string {
	return rr.alphaField(rr.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rr *RelatedRemittance) AddressLineOneField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rr *RelatedRemittance) AddressLineTwoField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rr *RelatedRemittance) AddressLineThreeField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rr *RelatedRemittance) AddressLineFourField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rr *RelatedRemittance) AddressLineFiveField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rr *RelatedRemittance) AddressLineSixField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rr *RelatedRemittance) AddressLineSevenField() string {
	return rr.alphaField(rr.RemittanceData.AddressLineSeven, 70)
}

// FormatRemittanceIdentification returns rr.RemittanceIdentification formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatRemittanceIdentification(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceIdentification, 35, options)
}

// FormatRemittanceLocationMethod returns rr.RemittanceLocationMethod formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatRemittanceLocationMethod(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceLocationMethod, 4, options)
}

// FormatRemittanceLocationElectronicAddress returns rr.RemittanceLocationElectronicAddress formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatRemittanceLocationElectronicAddress(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceLocationElectronicAddress, 2048, options)
}

// FormatName returns RemittanceData.Name formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatName(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.Name, 140, options)
}

// FormatAddressType returns RemittanceData.AddressType formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressType(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressType, 4, options)
}

// FormatDepartment returns RemittanceData.Department formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatDepartment(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.Department, 70, options)
}

// FormatSubDepartment returns RemittanceData.SubDepartment formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatSubDepartment(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.SubDepartment, 70, options)
}

// FormatStreetName returns RemittanceData.StreetName formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatStreetName(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.StreetName, 70, options)
}

// FormatBuildingNumber returns RemittanceData.BuildingNumber formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatBuildingNumber(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.BuildingNumber, 16, options)
}

// FormatPostCode returns RemittanceData.PostCode formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatPostCode(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.PostCode, 16, options)
}

// FormatTownName returns RemittanceData.TownName formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatTownName(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.TownName, 35, options)
}

// FormatCountrySubDivisionState returns RemittanceData.CountrySubDivisionState formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatCountrySubDivisionState(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.CountrySubDivisionState, 35, options)
}

// FormatCountry returns RemittanceData.Country formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatCountry(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.Country, 2, options)
}

// FormatAddressLineOne returns RemittanceData.AddressLineOne formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineOne(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineOne, 70, options)
}

// FormatAddressLineTwo returns RemittanceData.AddressLineTwo formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineTwo(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineTwo, 70, options)
}

// FormatAddressLineThree returns RemittanceData.AddressLineThree formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineThree(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineThree, 70, options)
}

// FormatAddressLineFour returns RemittanceData.AddressLineFour formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineFour(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineFour, 70, options)
}

// FormatAddressLineFive returns RemittanceData.AddressLineFive formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineFive(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineFive, 70, options)
}

// FormatAddressLineSix returns RemittanceData.AddressLineSix formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineSix(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineSix, 70, options)
}

// FormatAddressLineSeven returns RemittanceData.AddressLineSeven formatted according to the FormatOptions
func (rr *RelatedRemittance) FormatAddressLineSeven(options FormatOptions) string {
	return rr.formatAlphaField(rr.RemittanceData.AddressLineSeven, 70, options)
}
