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
func (r *RemittanceData) ParseForRelatedRemittance(record string) (length int, err error) {

	var read int

	if r.Name, read, err = r.parseVariableStringField(record[length:], 140); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if r.AddressType, read, err = r.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("AddressType", err)
	}
	length += read

	if r.Department, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("Department", err)
	}
	length += read

	if r.SubDepartment, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("SubDepartment", err)
	}
	length += read

	if r.StreetName, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("StreetName", err)
	}
	length += read

	if r.BuildingNumber, read, err = r.parseVariableStringField(record[length:], 16); err != nil {
		return 0, fieldError("BuildingNumber", err)
	}
	length += read

	if r.PostCode, read, err = r.parseVariableStringField(record[length:], 16); err != nil {
		return 0, fieldError("PostCode", err)
	}
	length += read

	if r.TownName, read, err = r.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("TownName", err)
	}
	length += read

	if r.CountrySubDivisionState, read, err = r.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("CountrySubDivisionState", err)
	}
	length += read

	if r.Country, read, err = r.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("Country", err)
	}
	length += read

	if r.AddressLineOne, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineOne", err)
	}
	length += read

	if r.AddressLineTwo, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineTwo", err)
	}
	length += read

	if r.AddressLineThree, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineThree", err)
	}
	length += read

	if r.AddressLineFour, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineFour", err)
	}
	length += read

	if r.AddressLineFive, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineFive", err)
	}
	length += read

	if r.AddressLineSix, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineSix", err)
	}
	length += read

	if r.AddressLineSeven, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineSeven", err)
	}
	length += read

	return
}

// Parse takes the input string and parses the RemittanceData values
func (r *RemittanceData) ParseForRemittanceBeneficiary(record string) (length int, err error) {
	var read int

	if r.DateBirthPlace, read, err = r.parseVariableStringField(record[length:], 82); err != nil {
		return 0, fieldError("DateBirthPlace", err)
	}
	length += read

	if r.AddressType, read, err = r.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("AddressType", err)
	}
	length += read

	if r.Department, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("Department", err)
	}
	length += read

	if r.SubDepartment, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("SubDepartment", err)
	}
	length += read

	if r.StreetName, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("StreetName", err)
	}
	length += read

	if r.BuildingNumber, read, err = r.parseVariableStringField(record[length:], 16); err != nil {
		return 0, fieldError("BuildingNumber", err)
	}
	length += read

	if r.PostCode, read, err = r.parseVariableStringField(record[length:], 16); err != nil {
		return 0, fieldError("PostCode", err)
	}
	length += read

	if r.TownName, read, err = r.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("TownName", err)
	}
	length += read

	if r.CountrySubDivisionState, read, err = r.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("CountrySubDivisionState", err)
	}
	length += read

	if r.Country, read, err = r.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("Country", err)
	}
	length += read

	if r.AddressLineOne, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineOne", err)
	}
	length += read

	if r.AddressLineTwo, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineTwo", err)
	}
	length += read

	if r.AddressLineThree, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineThree", err)
	}
	length += read

	if r.AddressLineFour, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineFour", err)
	}
	length += read

	if r.AddressLineFive, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineFive", err)
	}
	length += read

	if r.AddressLineSix, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineSix", err)
	}
	length += read

	if r.AddressLineSeven, read, err = r.parseVariableStringField(record[length:], 70); err != nil {
		return 0, fieldError("AddressLineSeven", err)
	}
	length += read

	if r.CountryOfResidence, read, err = r.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("CountryOfResidence", err)
	}
	length += read

	return
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
