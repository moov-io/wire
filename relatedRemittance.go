// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &RelatedRemittance{}

// RelatedRemittance is related remittance
type RelatedRemittance struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
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
func NewRelatedRemittance(isVariable bool) *RelatedRemittance {
	rr := &RelatedRemittance{
		tag:              TagRelatedRemittance,
		isVariableLength: isVariable,
	}
	return rr
}

// Parse takes the input string and parses the RelatedRemittance values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rr *RelatedRemittance) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 26 {
		return 0, NewTagWrongLengthErr(26, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if rr.tag, read, err = rr.parseTag(record); err != nil {
		return 0, fieldError("RelatedRemittance.Tag", err)
	}
	length += read

	if rr.RemittanceIdentification, read, err = rr.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("RemittanceIdentification", err)
	}
	length += read

	if rr.RemittanceLocationMethod, read, err = rr.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("RemittanceLocationMethod", err)
	}
	length += read

	if rr.RemittanceLocationElectronicAddress, read, err = rr.parseVariableStringField(record[length:], 2048); err != nil {
		return 0, fieldError("RemittanceLocationElectronicAddress", err)
	}
	length += read

	if read, err = rr.RemittanceData.ParseForRelatedRemittance(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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
func (rr *RelatedRemittance) String() string {
	var buf strings.Builder
	buf.Grow(3041)

	buf.WriteString(rr.tag)
	buf.WriteString(rr.RemittanceIdentificationField())
	buf.WriteString(rr.RemittanceLocationMethodField())
	buf.WriteString(rr.RemittanceLocationElectronicAddressField())
	buf.WriteString(rr.RemittanceData.StringForRelatedRemittance(rr.isVariableLength))

	return buf.String()
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
	return rr.alphaVariableField(rr.RemittanceIdentification, 35, rr.isVariableLength)
}

// RemittanceLocationMethodField gets a string of the RemittanceLocationMethod field
func (rr *RelatedRemittance) RemittanceLocationMethodField() string {
	return rr.alphaVariableField(rr.RemittanceLocationMethod, 4, rr.isVariableLength)
}

// RemittanceLocationElectronicAddressField gets a string of the RemittanceLocationElectronicAddress field
func (rr *RelatedRemittance) RemittanceLocationElectronicAddressField() string {
	return rr.alphaVariableField(rr.RemittanceLocationElectronicAddress, 2048, rr.isVariableLength)
}
