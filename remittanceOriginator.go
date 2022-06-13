// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// RemittanceOriginator is remittance originator
type RemittanceOriginator struct {
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
	// ContactName
	ContactName string `json:"contactName,omitempty"`
	// ContactPhoneNumber
	ContactPhoneNumber string `json:"contactPhoneNumber,omitempty"`
	// ContactMobileNumber
	ContactMobileNumber string `json:"contactMobileNumber,omitempty"`
	// ContactFaxNumber
	ContactFaxNumber string `json:"contactFaxNumber,omitempty"`
	// ContactElectronicAddress ( i.e., E-mail or URL address)
	ContactElectronicAddress string `json:"contactElectronicAddress,omitempty"`
	// ContactOther
	ContactOther string `json:"contactOther,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceOriginator returns a new RemittanceOriginator
func NewRemittanceOriginator() *RemittanceOriginator {
	ro := &RemittanceOriginator{
		tag: TagRemittanceOriginator,
	}
	return ro
}

// Parse takes the input string and parses the RemittanceOriginator values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ro *RemittanceOriginator) Parse(record string) error {
	if utf8.RuneCountInString(record) < 21 {
		return NewTagMinLengthErr(21, len(record))
	}

	ro.tag = record[:6]
	length := 6

	value, read, err := ro.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("IdentificationType", err)
	}
	ro.IdentificationType = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("IdentificationCode", err)
	}
	ro.IdentificationCode = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("Name", err)
	}
	ro.RemittanceData.Name = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("IdentificationNumber", err)
	}
	ro.IdentificationNumber = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("IdentificationNumberIssuer", err)
	}
	ro.IdentificationNumberIssuer = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 82)
	if err != nil {
		return fieldError("DateBirthPlace", err)
	}
	ro.RemittanceData.DateBirthPlace = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 4)
	if err != nil {
		return fieldError("AddressType", err)
	}
	ro.RemittanceData.AddressType = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("Department", err)
	}
	ro.RemittanceData.Department = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("SubDepartment", err)
	}
	ro.RemittanceData.SubDepartment = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("StreetName", err)
	}
	ro.RemittanceData.StreetName = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("BuildingNumber", err)
	}
	ro.RemittanceData.BuildingNumber = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 16)
	if err != nil {
		return fieldError("PostCode", err)
	}
	ro.RemittanceData.PostCode = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("TownName", err)
	}
	ro.RemittanceData.TownName = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("CountrySubDivisionState", err)
	}
	ro.RemittanceData.CountrySubDivisionState = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("Country", err)
	}
	ro.RemittanceData.Country = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineOne", err)
	}
	ro.RemittanceData.AddressLineOne = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineTwo", err)
	}
	ro.RemittanceData.AddressLineTwo = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineThree", err)
	}
	ro.RemittanceData.AddressLineThree = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFour", err)
	}
	ro.RemittanceData.AddressLineFour = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineFive", err)
	}
	ro.RemittanceData.AddressLineFive = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSix", err)
	}
	ro.RemittanceData.AddressLineSix = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 70)
	if err != nil {
		return fieldError("AddressLineSeven", err)
	}
	ro.RemittanceData.AddressLineSeven = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 2)
	if err != nil {
		return fieldError("CountryOfResidence", err)
	}
	ro.RemittanceData.CountryOfResidence = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 140)
	if err != nil {
		return fieldError("ContactName", err)
	}
	ro.ContactName = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactPhoneNumber", err)
	}
	ro.ContactPhoneNumber = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactMobileNumber", err)
	}
	ro.ContactMobileNumber = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactFaxNumber", err)
	}
	ro.ContactFaxNumber = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 2048)
	if err != nil {
		return fieldError("ContactElectronicAddress", err)
	}
	ro.ContactElectronicAddress = value
	length += read

	value, read, err = ro.parseVariableStringField(record[length:], 35)
	if err != nil {
		return fieldError("ContactOther", err)
	}
	ro.ContactOther = value
	length += read

	if !ro.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (ro *RemittanceOriginator) UnmarshalJSON(data []byte) error {
	type Alias RemittanceOriginator
	aux := struct {
		*Alias
	}{
		(*Alias)(ro),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	ro.tag = TagRemittanceOriginator
	return nil
}

// String writes RemittanceOriginator
func (ro *RemittanceOriginator) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(3442)

	buf.WriteString(ro.tag)
	buf.WriteString(ro.IdentificationTypeField(options...))
	buf.WriteString(ro.IdentificationCodeField(options...))
	buf.WriteString(ro.NameField(options...))
	buf.WriteString(ro.IdentificationNumberField(options...))
	buf.WriteString(ro.IdentificationNumberIssuerField(options...))
	buf.WriteString(ro.DateBirthPlaceField(options...))
	buf.WriteString(ro.AddressTypeField(options...))
	buf.WriteString(ro.DepartmentField(options...))
	buf.WriteString(ro.SubDepartmentField(options...))
	buf.WriteString(ro.StreetNameField(options...))
	buf.WriteString(ro.BuildingNumberField(options...))
	buf.WriteString(ro.PostCodeField(options...))
	buf.WriteString(ro.TownNameField(options...))
	buf.WriteString(ro.CountrySubDivisionStateField(options...))
	buf.WriteString(ro.CountryField(options...))
	buf.WriteString(ro.AddressLineOneField(options...))
	buf.WriteString(ro.AddressLineTwoField(options...))
	buf.WriteString(ro.AddressLineThreeField(options...))
	buf.WriteString(ro.AddressLineFourField(options...))
	buf.WriteString(ro.AddressLineFiveField(options...))
	buf.WriteString(ro.AddressLineSixField(options...))
	buf.WriteString(ro.AddressLineSevenField(options...))
	buf.WriteString(ro.CountryOfResidenceField(options...))
	buf.WriteString(ro.ContactNameField(options...))
	buf.WriteString(ro.ContactPhoneNumberField(options...))
	buf.WriteString(ro.ContactMobileNumberField(options...))
	buf.WriteString(ro.ContactFaxNumberField(options...))
	buf.WriteString(ro.ContactElectronicAddressField(options...))
	buf.WriteString(ro.ContactOtherField(options...))

	if ro.parseFirstOption(options) {
		return ro.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on RemittanceOriginator and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
// * Identification Type, Identification Code and Name are mandatory.
// * Identification Number is mandatory for all Identification Codes except PICDateBirthPlace.
// * Identification Number is not permitted for Identification Code PICDateBirthPlace.
// * Identification Number Issuer is not permitted for Identification Code OICSWIFTBICORBEI and PICDateBirthPlace.
// * Date & Place of Birth is only permitted for Identification Code PICDateBirthPlace.
func (ro *RemittanceOriginator) Validate() error { //nolint:gocyclo
	if err := ro.fieldInclusion(); err != nil {
		return err
	}
	if ro.tag != TagRemittanceOriginator {
		return fieldError("tag", ErrValidTagForType, ro.tag)
	}
	if err := ro.isIdentificationType(ro.IdentificationType); err != nil {
		return fieldError("IdentificationType", err, ro.IdentificationType)
	}

	switch ro.IdentificationType {
	case OrganizationID:
		if err := ro.isOrganizationIdentificationCode(ro.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, ro.IdentificationCode)
		}

	case PrivateID:
		if err := ro.isPrivateIdentificationCode(ro.IdentificationCode); err != nil {
			return fieldError("IdentificationCode", err, ro.IdentificationCode)
		}
	}

	if err := ro.isAlphanumeric(ro.IdentificationNumber); err != nil {
		return fieldError("IdentificationNumber", err, ro.IdentificationNumber)
	}
	if err := ro.isAlphanumeric(ro.IdentificationNumberIssuer); err != nil {
		return fieldError("IdentificationNumberIssuer", err, ro.IdentificationNumberIssuer)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Name); err != nil {
		return fieldError("Name", err, ro.RemittanceData.Name)
	}
	if err := ro.isAddressType(ro.RemittanceData.AddressType); err != nil {
		return fieldError("AddressType", err, ro.RemittanceData.AddressType)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Department); err != nil {
		return fieldError("Department", err, ro.RemittanceData.Department)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.SubDepartment); err != nil {
		return fieldError("SubDepartment", err, ro.RemittanceData.SubDepartment)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.StreetName); err != nil {
		return fieldError("StreetName", err, ro.RemittanceData.StreetName)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.BuildingNumber); err != nil {
		return fieldError("BuildingNumber", err, ro.RemittanceData.BuildingNumber)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.PostCode); err != nil {
		return fieldError("PostCode", err, ro.RemittanceData.PostCode)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.TownName); err != nil {
		return fieldError("TownName", err, ro.RemittanceData.TownName)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.CountrySubDivisionState); err != nil {
		return fieldError("CountrySubDivisionState", err, ro.RemittanceData.CountrySubDivisionState)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.Country); err != nil {
		return fieldError("Country", err, ro.RemittanceData.Country)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineOne); err != nil {
		return fieldError("AddressLineOne", err, ro.RemittanceData.AddressLineOne)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineTwo); err != nil {
		return fieldError("AddressLineTwo", err, ro.RemittanceData.AddressLineTwo)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineThree); err != nil {
		return fieldError("AddressLineThree", err, ro.RemittanceData.AddressLineThree)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineFour); err != nil {
		return fieldError("AddressLineFour", err, ro.RemittanceData.AddressLineFour)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineFive); err != nil {
		return fieldError("AddressLineFive", err, ro.RemittanceData.AddressLineFive)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineSix); err != nil {
		return fieldError("AddressLineSix", err, ro.RemittanceData.AddressLineSix)
	}
	if err := ro.isAlphanumeric(ro.RemittanceData.AddressLineSeven); err != nil {
		return fieldError("AddressLineSeven", err, ro.RemittanceData.AddressLineSeven)
	}

	if err := ro.isAlphanumeric(ro.RemittanceData.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, ro.RemittanceData.CountryOfResidence)
	}
	if err := ro.isAlphanumeric(ro.ContactName); err != nil {
		return fieldError("ContactName", err, ro.ContactName)
	}
	if err := ro.isAlphanumeric(ro.ContactPhoneNumber); err != nil {
		return fieldError("ContactPhoneNumber", err, ro.ContactPhoneNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactMobileNumber); err != nil {
		return fieldError("ContactMobileNumber", err, ro.ContactMobileNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactFaxNumber); err != nil {
		return fieldError("ContactFaxNumber", err, ro.ContactFaxNumber)
	}
	if err := ro.isAlphanumeric(ro.ContactElectronicAddress); err != nil {
		return fieldError("ContactElectronicAddress", err, ro.ContactElectronicAddress)
	}
	if err := ro.isAlphanumeric(ro.ContactOther); err != nil {
		return fieldError("ContactOther", err, ro.ContactOther)
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ro *RemittanceOriginator) fieldInclusion() error {
	if ro.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
	}
	if ro.IdentificationCode == PICDateBirthPlace {
		if ro.IdentificationNumber != "" {
			return fieldError("IdentificationNumber", ErrInvalidProperty, ro.IdentificationNumber)
		}
	}

	if ro.IdentificationNumber == "" || ro.IdentificationCode == OICSWIFTBICORBEI ||
		ro.IdentificationCode == PICDateBirthPlace {
		if ro.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty, ro.IdentificationNumberIssuer)
		}
	}
	if ro.IdentificationCode != PICDateBirthPlace {
		if ro.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty, ro.RemittanceData.DateBirthPlace)
		}
	}
	return nil
}

// IdentificationTypeField gets a string of the IdentificationType field
func (ro *RemittanceOriginator) IdentificationTypeField(options ...bool) string {
	return ro.alphaVariableField(ro.IdentificationType, 2, ro.parseFirstOption(options))
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ro *RemittanceOriginator) IdentificationCodeField(options ...bool) string {
	return ro.alphaVariableField(ro.IdentificationCode, 4, ro.parseFirstOption(options))
}

// NameField gets a string of the Name field
func (ro *RemittanceOriginator) NameField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.Name, 140, ro.parseFirstOption(options))
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (ro *RemittanceOriginator) IdentificationNumberField(options ...bool) string {
	return ro.alphaVariableField(ro.IdentificationNumber, 35, ro.parseFirstOption(options))
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (ro *RemittanceOriginator) IdentificationNumberIssuerField(options ...bool) string {
	return ro.alphaVariableField(ro.IdentificationNumberIssuer, 35, ro.parseFirstOption(options))
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (ro *RemittanceOriginator) DateBirthPlaceField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.DateBirthPlace, 82, ro.parseFirstOption(options))
}

// AddressTypeField gets a string of the AddressType field
func (ro *RemittanceOriginator) AddressTypeField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressType, 4, ro.parseFirstOption(options))
}

// DepartmentField gets a string of the Department field
func (ro *RemittanceOriginator) DepartmentField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.Department, 70, ro.parseFirstOption(options))
}

// SubDepartmentField gets a string of the SubDepartment field
func (ro *RemittanceOriginator) SubDepartmentField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.SubDepartment, 70, ro.parseFirstOption(options))
}

// StreetNameField gets a string of the StreetName field
func (ro *RemittanceOriginator) StreetNameField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.StreetName, 70, ro.parseFirstOption(options))
}

// BuildingNumberField gets a string of the BuildingNumber field
func (ro *RemittanceOriginator) BuildingNumberField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.BuildingNumber, 16, ro.parseFirstOption(options))
}

// PostCodeField gets a string of the PostCode field
func (ro *RemittanceOriginator) PostCodeField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.PostCode, 16, ro.parseFirstOption(options))
}

// TownNameField gets a string of the TownName field
func (ro *RemittanceOriginator) TownNameField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.TownName, 35, ro.parseFirstOption(options))
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (ro *RemittanceOriginator) CountrySubDivisionStateField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.CountrySubDivisionState, 35, ro.parseFirstOption(options))
}

// CountryField gets a string of the Country field
func (ro *RemittanceOriginator) CountryField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.Country, 2, ro.parseFirstOption(options))
}

// AddressLineOneField gets a string of the AddressLineOne field
func (ro *RemittanceOriginator) AddressLineOneField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineOne, 70, ro.parseFirstOption(options))
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (ro *RemittanceOriginator) AddressLineTwoField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineTwo, 70, ro.parseFirstOption(options))
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (ro *RemittanceOriginator) AddressLineThreeField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineThree, 70, ro.parseFirstOption(options))
}

// AddressLineFourField gets a string of the AddressLineFour field
func (ro *RemittanceOriginator) AddressLineFourField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineOne, 70, ro.parseFirstOption(options))
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (ro *RemittanceOriginator) AddressLineFiveField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineFive, 70, ro.parseFirstOption(options))
}

// AddressLineSixField gets a string of the AddressLineSix field
func (ro *RemittanceOriginator) AddressLineSixField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineSix, 70, ro.parseFirstOption(options))
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (ro *RemittanceOriginator) AddressLineSevenField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.AddressLineSeven, 70, ro.parseFirstOption(options))
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (ro *RemittanceOriginator) CountryOfResidenceField(options ...bool) string {
	return ro.alphaVariableField(ro.RemittanceData.CountryOfResidence, 2, ro.parseFirstOption(options))
}

// ContactNameField gets a string of the ContactName field
func (ro *RemittanceOriginator) ContactNameField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactName, 140, ro.parseFirstOption(options))
}

// ContactPhoneNumberField gets a string of the ContactPhoneNumber field
func (ro *RemittanceOriginator) ContactPhoneNumberField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactPhoneNumber, 35, ro.parseFirstOption(options))
}

// ContactMobileNumberField gets a string of the ContactMobileNumber field
func (ro *RemittanceOriginator) ContactMobileNumberField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactMobileNumber, 35, ro.parseFirstOption(options))
}

// ContactFaxNumberField gets a string of the ContactFaxNumber field
func (ro *RemittanceOriginator) ContactFaxNumberField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactFaxNumber, 35, ro.parseFirstOption(options))
}

// ContactElectronicAddressField gets a string of the ContactElectronicAddress field
func (ro *RemittanceOriginator) ContactElectronicAddressField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactElectronicAddress, 2048, ro.parseFirstOption(options))
}

// ContactOtherField gets a string of the ContactOther field
func (ro *RemittanceOriginator) ContactOtherField(options ...bool) string {
	return ro.alphaVariableField(ro.ContactOther, 35, ro.parseFirstOption(options))
}
