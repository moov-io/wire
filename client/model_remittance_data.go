/*
 * Wire API
 *
 * Moov Wire implements an HTTP API for creating, parsing, and validating Fedwire messages.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RemittanceData struct for RemittanceData
type RemittanceData struct {
	// Name
	Name string `json:"name,omitempty"`
	// AddressType  * `ADDR` - Complete Postal Address * `BIZZ` - Business Address * `DLVY` - Delivery Address * `HOME` - Home Address * `MLTO` - Mail Address * `PBOX` - Post Office Box
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
