// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// RemittanceData is remittance data
type RemittanceData struct {
	// Name
	Name string `json:"name,omitempty"`
	// DateBirthPlace
	DateBirthPlace string `json:"dateBirthPlace,omitempty"`
	// AddressType
	AddressType string `json:"addressType,omitempty"`
	// Department
	Department string `json:"department,omitempty"`
	// SubDepartment
	SubDepartment string `json:"subDepartment,omitempty"`
	// StreetName
	StreetName string `json:"streetName,omitempty"`
	// BuildingNumber
	BuildingNumber string `json:"buildingNumber,omitempty"`
	// PostCode
	PostCode string `json:"postCode,omitempty"`
	// TownName
	TownName string `json:"townName,omitempty"`
	// CountrySubDivisionState
	CountrySubDivisionState string `json:"countrySubDivisionState,omitempty"`
	// Country
	Country string `json:"country,omitempty"`
	// AddressLineOne
	AddressLineOne string `json:"addressLineOne,omitempty"`
	// AddressLineTwo
	AddressLineTwo string `json:"addressLineTwo,omitempty"`
	// AddressLineThree
	AddressLineThree string `json:"addressLineThree,omitempty"`
	// AddressLineFour
	AddressLineFour string `json:"addressLineFour,omitempty"`
	// AddressLineFive
	AddressLineFive string `json:"addressLineFive,omitempty"`
	// AddressLineSix
	AddressLineSix string `json:"addressLineSix,omitempty"`
	// AddressLineSeven
	AddressLineSeven string `json:"addressLineSeven,omitempty"`
	// CountryOfResidence
	CountryOfResidence string `json:"countryOfResidence,omitempty"`

	// converters is composed for WIRE to GoLang Converters
	converters
}

// Parse takes the input string and parses the RemittanceData values
func (r *RemittanceData) ParseForRelatedRemittance(record string) int {

	length := 6
	read := 0

	r.Name, read = r.parseVariableStringField(record[length:], 140)
	length += read

	r.AddressType, read = r.parseVariableStringField(record[length:], 4)
	length += read

	r.Department, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.SubDepartment, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.StreetName, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.BuildingNumber, read = r.parseVariableStringField(record[length:], 16)
	length += read

	r.PostCode, read = r.parseVariableStringField(record[length:], 16)
	length += read

	r.TownName, read = r.parseVariableStringField(record[length:], 35)
	length += read

	r.CountrySubDivisionState, read = r.parseVariableStringField(record[length:], 35)
	length += read

	r.Country, read = r.parseVariableStringField(record[length:], 2)
	length += read

	r.AddressLineOne, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineTwo, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineThree, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineFour, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineFive, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineSix, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineSeven, read = r.parseVariableStringField(record[length:], 70)
	length += read

	return length
}

// Parse takes the input string and parses the RemittanceData values
func (r *RemittanceData) ParseForRemittanceBeneficiary(record string) int {
	length := 6
	read := 0

	r.DateBirthPlace, read = r.parseVariableStringField(record[length:], 82)
	length += read

	r.AddressType, read = r.parseVariableStringField(record[length:], 4)
	length += read

	r.Department, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.SubDepartment, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.StreetName, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.BuildingNumber, read = r.parseVariableStringField(record[length:], 16)
	length += read

	r.PostCode, read = r.parseVariableStringField(record[length:], 16)
	length += read

	r.TownName, read = r.parseVariableStringField(record[length:], 35)
	length += read

	r.CountrySubDivisionState, read = r.parseVariableStringField(record[length:], 35)
	length += read

	r.Country, read = r.parseVariableStringField(record[length:], 2)
	length += read

	r.AddressLineOne, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineTwo, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineThree, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineFour, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineFive, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineSix, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.AddressLineSeven, read = r.parseVariableStringField(record[length:], 70)
	length += read

	r.CountryOfResidence, read = r.parseVariableStringField(record[length:], 2)
	length += read

	return length
}

// String writes RelatedRemittance
func (r *RemittanceData) StringForRelatedRemittance(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(954)

	buf.WriteString(r.NameField(isVariable))
	buf.WriteString(r.AddressTypeField(isVariable))
	buf.WriteString(r.DepartmentField(isVariable))
	buf.WriteString(r.SubDepartmentField(isVariable))
	buf.WriteString(r.StreetNameField(isVariable))
	buf.WriteString(r.BuildingNumberField(isVariable))
	buf.WriteString(r.PostCodeField(isVariable))
	buf.WriteString(r.TownNameField(isVariable))
	buf.WriteString(r.CountrySubDivisionStateField(isVariable))
	buf.WriteString(r.CountryField(isVariable))
	buf.WriteString(r.AddressLineOneField(isVariable))
	buf.WriteString(r.AddressLineTwoField(isVariable))
	buf.WriteString(r.AddressLineThreeField(isVariable))
	buf.WriteString(r.AddressLineFourField(isVariable))
	buf.WriteString(r.AddressLineFiveField(isVariable))
	buf.WriteString(r.AddressLineSixField(isVariable))
	buf.WriteString(r.AddressLineSevenField(isVariable))

	return buf.String()
}

// String writes RelatedRemittance
func (r *RemittanceData) StringForRemittanceBeneficiary(isVariable bool) string {
	var buf strings.Builder
	buf.Grow(892)

	buf.WriteString(r.DateBirthPlaceField(isVariable))
	buf.WriteString(r.AddressTypeField(isVariable))
	buf.WriteString(r.DepartmentField(isVariable))
	buf.WriteString(r.SubDepartmentField(isVariable))
	buf.WriteString(r.StreetNameField(isVariable))
	buf.WriteString(r.BuildingNumberField(isVariable))
	buf.WriteString(r.PostCodeField(isVariable))
	buf.WriteString(r.TownNameField(isVariable))
	buf.WriteString(r.CountrySubDivisionStateField(isVariable))
	buf.WriteString(r.CountryField(isVariable))
	buf.WriteString(r.AddressLineOneField(isVariable))
	buf.WriteString(r.AddressLineTwoField(isVariable))
	buf.WriteString(r.AddressLineThreeField(isVariable))
	buf.WriteString(r.AddressLineFourField(isVariable))
	buf.WriteString(r.AddressLineFiveField(isVariable))
	buf.WriteString(r.AddressLineSixField(isVariable))
	buf.WriteString(r.AddressLineSevenField(isVariable))
	buf.WriteString(r.CountryOfResidenceField(isVariable))

	return buf.String()
}

// NameField gets a string of the Name field
func (r *RemittanceData) NameField(isVariable bool) string {
	return r.alphaVariableField(r.Name, 140, isVariable)
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (r *RemittanceData) DateBirthPlaceField(isVariable bool) string {
	return r.alphaVariableField(r.DateBirthPlace, 82, isVariable)
}

// AddressTypeField gets a string of the AddressType field
func (r *RemittanceData) AddressTypeField(isVariable bool) string {
	return r.alphaVariableField(r.AddressType, 4, isVariable)
}

// DepartmentField gets a string of the Department field
func (r *RemittanceData) DepartmentField(isVariable bool) string {
	return r.alphaVariableField(r.Department, 70, isVariable)
}

// SubDepartmentField gets a string of the SubDepartment field
func (r *RemittanceData) SubDepartmentField(isVariable bool) string {
	return r.alphaVariableField(r.SubDepartment, 70, isVariable)
}

// StreetNameField gets a string of the StreetName field
func (r *RemittanceData) StreetNameField(isVariable bool) string {
	return r.alphaVariableField(r.StreetName, 70, isVariable)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (r *RemittanceData) BuildingNumberField(isVariable bool) string {
	return r.alphaVariableField(r.BuildingNumber, 16, isVariable)
}

// PostCodeField gets a string of the PostCode field
func (r *RemittanceData) PostCodeField(isVariable bool) string {
	return r.alphaVariableField(r.PostCode, 16, isVariable)
}

// TownNameField gets a string of the TownName field
func (r *RemittanceData) TownNameField(isVariable bool) string {
	return r.alphaVariableField(r.TownName, 35, isVariable)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (r *RemittanceData) CountrySubDivisionStateField(isVariable bool) string {
	return r.alphaVariableField(r.CountrySubDivisionState, 35, isVariable)
}

// CountryField gets a string of the Country field
func (r *RemittanceData) CountryField(isVariable bool) string {
	return r.alphaVariableField(r.Country, 2, isVariable)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (r *RemittanceData) AddressLineOneField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineOne, 70, isVariable)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (r *RemittanceData) AddressLineTwoField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineTwo, 70, isVariable)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (r *RemittanceData) AddressLineThreeField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineThree, 70, isVariable)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (r *RemittanceData) AddressLineFourField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineFour, 70, isVariable)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (r *RemittanceData) AddressLineFiveField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineFive, 70, isVariable)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (r *RemittanceData) AddressLineSixField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineSix, 70, isVariable)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (r *RemittanceData) AddressLineSevenField(isVariable bool) string {
	return r.alphaVariableField(r.AddressLineSeven, 70, isVariable)
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (r *RemittanceData) CountryOfResidenceField(isVariable bool) string {
	return r.alphaVariableField(r.CountryOfResidence, 2, isVariable)
}
