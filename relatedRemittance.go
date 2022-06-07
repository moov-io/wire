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

	var err error
	length := 6
	read := 0

	if rr.RemittanceIdentification, read, err = rr.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("RemittanceIdentification", err)
	}
	length += read

	if rr.RemittanceLocationMethod, read, err = rr.parseVariableStringField(record[length:], 4); err != nil {
		return fieldError("RemittanceLocationMethod", err)
	}
	length += read

	if rr.RemittanceLocationElectronicAddress, read, err = rr.parseVariableStringField(record[length:], 2048); err != nil {
		return fieldError("RemittanceLocationElectronicAddress", err)
	}
	length += read

	if rr.RemittanceData.Name, read, err = rr.parseVariableStringField(record[length:], 140); err != nil {
		return fieldError("Name", err)
	}
	length += read

	if rr.RemittanceData.AddressType, read, err = rr.parseVariableStringField(record[length:], 4); err != nil {
		return fieldError("AddressType", err)
	}
	length += read

	if rr.RemittanceData.Department, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("Department", err)
	}
	length += read

	if rr.RemittanceData.SubDepartment, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("SubDepartment", err)
	}
	length += read

	if rr.RemittanceData.StreetName, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("StreetName", err)
	}
	length += read

		if rr.RemittanceData.BuildingNumber, read, err = rr.parseVariableStringField(record[length:], 16); err != nil {
		return fieldError("BuildingNumber", err)
	}
	length += read

	if rr.RemittanceData.PostCode, read, err = rr.parseVariableStringField(record[length:], 16); err != nil {
		return fieldError("PostCode", err)
	}
	length += read

	if rr.RemittanceData.TownName, read, err = rr.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("TownName", err)
	}
	length += read

	if rr.RemittanceData.CountrySubDivisionState, read, err = rr.parseVariableStringField(record[length:], 35); err != nil {
		return fieldError("CountrySubDivisionState", err)
	}
	length += read

	if rr.RemittanceData.Country, read, err = rr.parseVariableStringField(record[length:], 2); err != nil {
		return fieldError("Country", err)
	}
	length += read

	if rr.RemittanceData.AddressLineOne, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineOne", err)
	}
	length += read

	if rr.RemittanceData.AddressLineTwo, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineTwo", err)
	}
	length += read

	if rr.RemittanceData.AddressLineThree, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineThree", err)
	}
	length += read

	if rr.RemittanceData.AddressLineFour, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineFour", err)
	}
	length += read

	if rr.RemittanceData.AddressLineFive, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineFive", err)
	}
	length += read

	if rr.RemittanceData.AddressLineSix, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineSix", err)
	}
	length += read

	if rr.RemittanceData.AddressLineSeven, read, err = rr.parseVariableStringField(record[length:], 70); err != nil {
		return fieldError("AddressLineSeven", err)
	}
	length += read

	if len(record) != length {
		return NewTagMaxLengthErr()
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

// String writes RelatedRemittance
func (rr *RelatedRemittance) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(3041)

	buf.WriteString(rr.tag)
	buf.WriteString(rr.RemittanceIdentificationField(options...))
	buf.WriteString(rr.RemittanceLocationMethodField(options...))
	buf.WriteString(rr.RemittanceLocationElectronicAddressField(options...))
	buf.WriteString(rr.NameField(options...))
	buf.WriteString(rr.AddressTypeField(options...))
	buf.WriteString(rr.DepartmentField(options...))
	buf.WriteString(rr.SubDepartmentField(options...))
	buf.WriteString(rr.StreetNameField(options...))
	buf.WriteString(rr.BuildingNumberField(options...))
	buf.WriteString(rr.PostCodeField(options...))
	buf.WriteString(rr.TownNameField(options...))
	buf.WriteString(rr.CountrySubDivisionStateField(options...))
	buf.WriteString(rr.CountryField(options...))
	buf.WriteString(rr.AddressLineOneField(options...))
	buf.WriteString(rr.AddressLineTwoField(options...))
	buf.WriteString(rr.AddressLineThreeField(options...))
	buf.WriteString(rr.AddressLineFourField(options...))
	buf.WriteString(rr.AddressLineFiveField(options...))
	buf.WriteString(rr.AddressLineSixField(options...))
	buf.WriteString(rr.AddressLineSevenField(options...))

	if rr.parseFirstOption(options) {
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
func (rr *RelatedRemittance) RemittanceIdentificationField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceIdentification, 35, rr.parseFirstOption(options))
}

// RemittanceLocationMethodField gets a string of the RemittanceLocationMethod field
func (rr *RelatedRemittance) RemittanceLocationMethodField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceLocationMethod, 4, rr.parseFirstOption(options))
}

// RemittanceLocationElectronicAddressField gets a string of the RemittanceLocationElectronicAddress field
func (rr *RelatedRemittance) RemittanceLocationElectronicAddressField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceLocationElectronicAddress, 2048, rr.parseFirstOption(options))
}

// NameField gets a string of the Name field
func (rr *RelatedRemittance) NameField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.Name, 140, rr.parseFirstOption(options))
}

// AddressTypeField gets a string of the AddressType field
func (rr *RelatedRemittance) AddressTypeField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressType, 4, rr.parseFirstOption(options))
}

// DepartmentField gets a string of the Department field
func (rr *RelatedRemittance) DepartmentField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.Department, 70, rr.parseFirstOption(options))
}

// SubDepartmentField gets a string of the SubDepartment field
func (rr *RelatedRemittance) SubDepartmentField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.SubDepartment, 70, rr.parseFirstOption(options))
}

// StreetNameField gets a string of the StreetName field
func (rr *RelatedRemittance) StreetNameField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.StreetName, 70, rr.parseFirstOption(options))
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rr *RelatedRemittance) BuildingNumberField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.BuildingNumber, 16, rr.parseFirstOption(options))
}

// PostCodeField gets a string of the PostCode field
func (rr *RelatedRemittance) PostCodeField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.PostCode, 16, rr.parseFirstOption(options))
}

// TownNameField gets a string of the TownName field
func (rr *RelatedRemittance) TownNameField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.TownName, 35, rr.parseFirstOption(options))
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rr *RelatedRemittance) CountrySubDivisionStateField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.CountrySubDivisionState, 35, rr.parseFirstOption(options))
}

// CountryField gets a string of the Country field
func (rr *RelatedRemittance) CountryField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.Country, 2, rr.parseFirstOption(options))
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rr *RelatedRemittance) AddressLineOneField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineOne, 70, rr.parseFirstOption(options))
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rr *RelatedRemittance) AddressLineTwoField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineTwo, 70, rr.parseFirstOption(options))
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rr *RelatedRemittance) AddressLineThreeField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineThree, 70, rr.parseFirstOption(options))
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rr *RelatedRemittance) AddressLineFourField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineFour, 70, rr.parseFirstOption(options))
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rr *RelatedRemittance) AddressLineFiveField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineFive, 70, rr.parseFirstOption(options))
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rr *RelatedRemittance) AddressLineSixField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineSix, 70, rr.parseFirstOption(options))
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rr *RelatedRemittance) AddressLineSevenField(options ...bool) string {
	return rr.alphaVariableField(rr.RemittanceData.AddressLineSeven, 70, rr.parseFirstOption(options))
}
