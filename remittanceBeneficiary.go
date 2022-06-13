// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// RemittanceBeneficiary is remittance beneficiary
type RemittanceBeneficiary struct {
	// tag
	tag string
	// IdentificationType is identification type
	IdentificationType string `json:"identificationType,omitempty"`
	// IdentificationCode  Organization Identification Codes  * `BANK` - Bank Party Identification * `CUST` - Customer Number * `DUNS` - Data Universal Number System (Dun & Bradstreet) * `EMPL` - Employer Identification Number * `GS1G` - Global Location Number * `PROP` - Proprietary Identification Number * `SWBB` - SWIFT BIC or BEI * `TXID` - Tax Identification Number  Private Identification Codes  * `ARNU` - Alien Registration Number * `CCPT` - Passport Number * `CUST` - Customer Number * `DPOB` - Date & Place of Birth * `DRLC` - Driver’s License Number * `EMPL` - Employee Identification Number * `NIDN` - National Identity Number * `PROP` - Proprietary Identification Number * `SOSE` - Social Security Number * `TXID` - Tax Identification Number
	IdentificationCode string `json:"identificationCode,omitempty"`
	// IdentificationNumber
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// IdentificationNumberIssuer
	IdentificationNumberIssuer string `json:"identificationNumberIssuer,omitempty"`
	// RemittanceData
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceBeneficiary returns a new RemittanceBeneficiary
func NewRemittanceBeneficiary() *RemittanceBeneficiary {
	rb := &RemittanceBeneficiary{
		tag: TagRemittanceBeneficiary,
	}
	return rb
}

// Parse takes the input string and parses the RemittanceBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rb *RemittanceBeneficiary) Parse(record string) error {
	if utf8.RuneCountInString(record) < 24 {
		return NewTagMinLengthErr(24, len(record))
	}

	rb.tag = record[:6]
	length := 6

	value, read, err := rb.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("Name", err)
	}
	rb.RemittanceData.Name = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("IdentificationType", err)
	}
	rb.IdentificationType = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("IdentificationCode", err)
	}
	rb.IdentificationCode = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("IdentificationNumber", err)
	}
	rb.IdentificationNumber = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("IdentificationNumberIssuer", err)
	}
	rb.IdentificationNumberIssuer = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 82)
	if err != nil {
		return fieldError("DateBirthPlace", err)
	}
	rb.RemittanceData.DateBirthPlace = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("AddressType", err)
	}
	rb.RemittanceData.AddressType = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("Department", err)
	}
	rb.RemittanceData.Department = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("SubDepartment", err)
	}
	rb.RemittanceData.SubDepartment = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("StreetName", err)
	}
	rb.RemittanceData.StreetName = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("BuildingNumber", err)
	}
	rb.RemittanceData.BuildingNumber = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("PostCode", err)
	}
	rb.RemittanceData.PostCode = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("TownName", err)
	}
	rb.RemittanceData.TownName = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("CountrySubDivisionState", err)
	}
	rb.RemittanceData.CountrySubDivisionState = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("Country", err)
	}
	rb.RemittanceData.Country = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	rb.RemittanceData.AddressLineOne = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	rb.RemittanceData.AddressLineTwo = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	rb.RemittanceData.AddressLineThree = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFour", err)
	}
	rb.RemittanceData.AddressLineFour = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFive", err)
	}
	rb.RemittanceData.AddressLineFive = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSix", err)
	}
	rb.RemittanceData.AddressLineSix = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSeven", err)
	}
	rb.RemittanceData.AddressLineSeven = value
	length += read

	value, read, err = rb.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("CountryOfResidence", err)
	}
	rb.RemittanceData.CountryOfResidence = value
	length += read

	if !rb.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (rb *RemittanceBeneficiary) UnmarshalJSON(data []byte) error {
	type Alias RemittanceBeneficiary
	aux := struct {
		*Alias
	}{
		(*Alias)(rb),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	rb.tag = TagRemittanceBeneficiary
	return nil
}

// String writes RemittanceBeneficiary
func (rb *RemittanceBeneficiary) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(1114)

	buf.WriteString(rb.tag)
	buf.WriteString(rb.NameField(options...))
	buf.WriteString(rb.IdentificationTypeField(options...))
	buf.WriteString(rb.IdentificationCodeField(options...))
	buf.WriteString(rb.IdentificationNumberField(options...))
	buf.WriteString(rb.IdentificationNumberIssuerField(options...))
	buf.WriteString(rb.DateBirthPlaceField(options...))
	buf.WriteString(rb.AddressTypeField(options...))
	buf.WriteString(rb.DepartmentField(options...))
	buf.WriteString(rb.SubDepartmentField(options...))
	buf.WriteString(rb.StreetNameField(options...))
	buf.WriteString(rb.BuildingNumberField(options...))
	buf.WriteString(rb.PostCodeField(options...))
	buf.WriteString(rb.TownNameField(options...))
	buf.WriteString(rb.CountrySubDivisionStateField(options...))
	buf.WriteString(rb.CountryField(options...))
	buf.WriteString(rb.AddressLineOneField(options...))
	buf.WriteString(rb.AddressLineTwoField(options...))
	buf.WriteString(rb.AddressLineThreeField(options...))
	buf.WriteString(rb.AddressLineFourField(options...))
	buf.WriteString(rb.AddressLineFiveField(options...))
	buf.WriteString(rb.AddressLineSixField(options...))
	buf.WriteString(rb.AddressLineSevenField(options...))
	buf.WriteString(rb.CountryOfResidenceField(options...))

	if rb.parseFirstOption(options) {
		return rb.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on RemittanceBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// * Name is mandatory.
// * Identification Number
//   * Not permitted unless Identification Type and Identification Code are present.
//   * Not permitted for Identification Code PICDateBirthPlace.
// * Identification Number Issuer
//   * Not permitted unless Identification Type, Identification Code and Identification Number are present.
//   * Not permitted for Identification Code SWBB and PICDateBirthPlace.
// * Date & Place of Birth is only permitted for Identification Code PICDateBirthPlace.
func (rb *RemittanceBeneficiary) Validate() error {
	if err := rb.fieldInclusion(); err != nil {
		return err
	}
	if rb.tag != TagRemittanceBeneficiary {
		return fieldError("tag", ErrValidTagForType, rb.tag)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Name); err != nil {
		return fieldError("Name", err, rb.RemittanceData.Name)
	}
	if err := rb.isIdentificationType(rb.IdentificationType); err != nil {
		return fieldError("IdentificationType", err, rb.IdentificationType)
	}
	switch rb.IdentificationType {
	case OrganizationID:
		if err := rb.isOrganizationIdentificationCode(rb.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, rb.IdentificationCode)
		}
	case PrivateID:
		if err := rb.isPrivateIdentificationCode(rb.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, rb.IdentificationCode)
		}
	}
	if err := rb.isAlphanumeric(rb.IdentificationNumber); err != nil {
		return fieldError("IdentificationNumber", err, rb.IdentificationNumber)
	}
	if err := rb.isAlphanumeric(rb.IdentificationNumberIssuer); err != nil {
		return fieldError("IdentificationNumberIssuer", err, rb.IdentificationNumberIssuer)
	}
	if err := rb.isAddressType(rb.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, rb.RemittanceData.AddressType)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Department); err != nil {
		return fieldError("Department", err, rb.RemittanceData.Department)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, rb.RemittanceData.SubDepartment)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, rb.RemittanceData.StreetName)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, rb.RemittanceData.BuildingNumber)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, rb.RemittanceData.PostCode)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, rb.RemittanceData.TownName)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, rb.RemittanceData.CountrySubDivisionState)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.Country); err != nil {
		return fieldError("Country", err, rb.RemittanceData.Country)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, rb.RemittanceData.AddressLineOne)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, rb.RemittanceData.AddressLineTwo)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, rb.RemittanceData.AddressLineThree)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, rb.RemittanceData.AddressLineFour)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, rb.RemittanceData.AddressLineFive)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, rb.RemittanceData.AddressLineSix)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, rb.RemittanceData.AddressLineSeven)
	}
	if err := rb.isAlphanumeric(rb.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, rb.RemittanceData.CountryOfResidence)
	}

	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rb *RemittanceBeneficiary) fieldInclusion() error {
	if rb.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}

	if rb.IdentificationCode == PICDateBirthPlace {
		if rb.IdentificationNumber != "" {
			return fieldError("IdentificationNumber", ErrInvalidProperty, rb.IdentificationNumber)
		}
	}
	if rb.IdentificationNumber == "" || rb.IdentificationCode == OICSWIFTBICORBEI ||
		rb.IdentificationCode == PICDateBirthPlace {
		if rb.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty, rb.IdentificationNumberIssuer)
		}
	}
	if rb.IdentificationCode != PICDateBirthPlace {
		if rb.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty, rb.RemittanceData.DateBirthPlace)
		}
	}

	return nil
}

// NameField gets a string of the Name field
func (rb *RemittanceBeneficiary) NameField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.Name, 140, rb.parseFirstOption(options))
}

// IdentificationTypeField gets a string of the IdentificationType field
func (rb *RemittanceBeneficiary) IdentificationTypeField(options ...bool) string {
	return rb.alphaVariableField(rb.IdentificationType, 2, rb.parseFirstOption(options))
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (rb *RemittanceBeneficiary) IdentificationCodeField(options ...bool) string {
	return rb.alphaVariableField(rb.IdentificationCode, 4, rb.parseFirstOption(options))
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (rb *RemittanceBeneficiary) IdentificationNumberField(options ...bool) string {
	return rb.alphaVariableField(rb.IdentificationNumber, 35, rb.parseFirstOption(options))
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (rb *RemittanceBeneficiary) IdentificationNumberIssuerField(options ...bool) string {
	return rb.alphaVariableField(rb.IdentificationNumberIssuer, 35, rb.parseFirstOption(options))
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (rb *RemittanceBeneficiary) DateBirthPlaceField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.DateBirthPlace, 82, rb.parseFirstOption(options))
}

// AddressTypeField gets a string of the AddressType field
func (rb *RemittanceBeneficiary) AddressTypeField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressType, 4, rb.parseFirstOption(options))
}

// DepartmentField gets a string of the Department field
func (rb *RemittanceBeneficiary) DepartmentField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.Department, 70, rb.parseFirstOption(options))
}

// SubDepartmentField gets a string of the SubDepartment field
func (rb *RemittanceBeneficiary) SubDepartmentField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.SubDepartment, 70, rb.parseFirstOption(options))
}

// StreetNameField gets a string of the StreetName field
func (rb *RemittanceBeneficiary) StreetNameField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.StreetName, 70, rb.parseFirstOption(options))
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rb *RemittanceBeneficiary) BuildingNumberField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.BuildingNumber, 16, rb.parseFirstOption(options))
}

// PostCodeField gets a string of the PostCode field
func (rb *RemittanceBeneficiary) PostCodeField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.PostCode, 16, rb.parseFirstOption(options))
}

// TownNameField gets a string of the TownName field
func (rb *RemittanceBeneficiary) TownNameField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.TownName, 35, rb.parseFirstOption(options))
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rb *RemittanceBeneficiary) CountrySubDivisionStateField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.CountrySubDivisionState, 35, rb.parseFirstOption(options))
}

// CountryField gets a string of the Country field
func (rb *RemittanceBeneficiary) CountryField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.Country, 2, rb.parseFirstOption(options))
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rb *RemittanceBeneficiary) AddressLineOneField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineOne, 70, rb.parseFirstOption(options))
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rb *RemittanceBeneficiary) AddressLineTwoField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineTwo, 70, rb.parseFirstOption(options))
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rb *RemittanceBeneficiary) AddressLineThreeField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineThree, 70, rb.parseFirstOption(options))
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rb *RemittanceBeneficiary) AddressLineFourField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineFour, 70, rb.parseFirstOption(options))
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rb *RemittanceBeneficiary) AddressLineFiveField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineFive, 70, rb.parseFirstOption(options))
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rb *RemittanceBeneficiary) AddressLineSixField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineSix, 70, rb.parseFirstOption(options))
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rb *RemittanceBeneficiary) AddressLineSevenField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.AddressLineSeven, 70, rb.parseFirstOption(options))
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (rb *RemittanceBeneficiary) CountryOfResidenceField(options ...bool) string {
	return rb.alphaVariableField(rb.RemittanceData.CountryOfResidence, 2, rb.parseFirstOption(options))
}
