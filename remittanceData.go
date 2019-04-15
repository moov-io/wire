// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// RemittanceData is remittance data
type RemittanceData struct {
	// Name
	Name string `json:"name,omitempty"`
	// DateBirthPlace
	DateBirthPlace string `json:"dateBirthPlace,omitempty""`
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
}
