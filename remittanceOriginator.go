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

	if err := ro.verifyDataWithReadLength(record, length); err != nil {
		return NewTagMaxLengthErr(err)
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

// String returns a fixed-width RemittanceOriginator record
func (ro *RemittanceOriginator) String() string {
	return ro.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a RemittanceOriginator record formatted according to the FormatOptions
func (ro *RemittanceOriginator) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(3442)

	buf.WriteString(ro.tag)
	buf.WriteString(ro.FormatIdentificationType(options))
	buf.WriteString(ro.FormatIdentificationCode(options))
	buf.WriteString(ro.FormatName(options))
	buf.WriteString(ro.FormatIdentificationNumber(options))
	buf.WriteString(ro.FormatIdentificationNumberIssuer(options))
	buf.WriteString(ro.FormatDateBirthPlace(options))
	buf.WriteString(ro.FormatAddressType(options))
	buf.WriteString(ro.FormatDepartment(options))
	buf.WriteString(ro.FormatSubDepartment(options))
	buf.WriteString(ro.FormatStreetName(options))
	buf.WriteString(ro.FormatBuildingNumber(options))
	buf.WriteString(ro.FormatPostCode(options))
	buf.WriteString(ro.FormatTownName(options))
	buf.WriteString(ro.FormatCountrySubDivisionState(options))
	buf.WriteString(ro.FormatCountry(options))
	buf.WriteString(ro.FormatAddressLineOne(options))
	buf.WriteString(ro.FormatAddressLineTwo(options))
	buf.WriteString(ro.FormatAddressLineThree(options))
	buf.WriteString(ro.FormatAddressLineFour(options))
	buf.WriteString(ro.FormatAddressLineFive(options))
	buf.WriteString(ro.FormatAddressLineSix(options))
	buf.WriteString(ro.FormatAddressLineSeven(options))
	buf.WriteString(ro.FormatCountryOfResidence(options))
	buf.WriteString(ro.FormatContactName(options))
	buf.WriteString(ro.FormatContactPhoneNumber(options))
	buf.WriteString(ro.FormatContactMobileNumber(options))
	buf.WriteString(ro.FormatContactFaxNumber(options))
	buf.WriteString(ro.FormatContactElectronicAddress(options))
	buf.WriteString(ro.FormatContactOther(options))

	if options.VariableLengthFields {
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
func (ro *RemittanceOriginator) IdentificationTypeField() string {
	return ro.alphaField(ro.IdentificationType, 2)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ro *RemittanceOriginator) IdentificationCodeField() string {
	return ro.alphaField(ro.IdentificationCode, 4)
}

// NameField gets a string of the Name field
func (ro *RemittanceOriginator) NameField() string {
	return ro.alphaField(ro.RemittanceData.Name, 140)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (ro *RemittanceOriginator) IdentificationNumberField() string {
	return ro.alphaField(ro.IdentificationNumber, 35)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (ro *RemittanceOriginator) IdentificationNumberIssuerField() string {
	return ro.alphaField(ro.IdentificationNumberIssuer, 35)
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (ro *RemittanceOriginator) DateBirthPlaceField() string {
	return ro.alphaField(ro.RemittanceData.DateBirthPlace, 82)
}

// AddressTypeField gets a string of the AddressType field
func (ro *RemittanceOriginator) AddressTypeField() string {
	return ro.alphaField(ro.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (ro *RemittanceOriginator) DepartmentField() string {
	return ro.alphaField(ro.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (ro *RemittanceOriginator) SubDepartmentField() string {
	return ro.alphaField(ro.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (ro *RemittanceOriginator) StreetNameField() string {
	return ro.alphaField(ro.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (ro *RemittanceOriginator) BuildingNumberField() string {
	return ro.alphaField(ro.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (ro *RemittanceOriginator) PostCodeField() string {
	return ro.alphaField(ro.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (ro *RemittanceOriginator) TownNameField() string {
	return ro.alphaField(ro.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (ro *RemittanceOriginator) CountrySubDivisionStateField() string {
	return ro.alphaField(ro.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (ro *RemittanceOriginator) CountryField() string {
	return ro.alphaField(ro.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (ro *RemittanceOriginator) AddressLineOneField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (ro *RemittanceOriginator) AddressLineTwoField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (ro *RemittanceOriginator) AddressLineThreeField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (ro *RemittanceOriginator) AddressLineFourField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (ro *RemittanceOriginator) AddressLineFiveField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (ro *RemittanceOriginator) AddressLineSixField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (ro *RemittanceOriginator) AddressLineSevenField() string {
	return ro.alphaField(ro.RemittanceData.AddressLineSeven, 70)
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (ro *RemittanceOriginator) CountryOfResidenceField() string {
	return ro.alphaField(ro.RemittanceData.CountryOfResidence, 2)
}

// ContactNameField gets a string of the ContactName field
func (ro *RemittanceOriginator) ContactNameField() string {
	return ro.alphaField(ro.ContactName, 140)
}

// ContactPhoneNumberField gets a string of the ContactPhoneNumber field
func (ro *RemittanceOriginator) ContactPhoneNumberField() string {
	return ro.alphaField(ro.ContactPhoneNumber, 35)
}

// ContactMobileNumberField gets a string of the ContactMobileNumber field
func (ro *RemittanceOriginator) ContactMobileNumberField() string {
	return ro.alphaField(ro.ContactMobileNumber, 35)
}

// ContactFaxNumberField gets a string of the ContactFaxNumber field
func (ro *RemittanceOriginator) ContactFaxNumberField() string {
	return ro.alphaField(ro.ContactFaxNumber, 35)
}

// ContactElectronicAddressField gets a string of the ContactElectronicAddress field
func (ro *RemittanceOriginator) ContactElectronicAddressField() string {
	return ro.alphaField(ro.ContactElectronicAddress, 2048)
}

// ContactOtherField gets a string of the ContactOther field
func (ro *RemittanceOriginator) ContactOtherField() string {
	return ro.alphaField(ro.ContactOther, 35)
}

// FormatIdentificationType returns IdentificationType formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatIdentificationType(options FormatOptions) string {
	return ro.formatAlphaField(ro.IdentificationType, 2, options)
}

// FormatIdentificationCode returns IdentificationCode formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatIdentificationCode(options FormatOptions) string {
	return ro.formatAlphaField(ro.IdentificationCode, 4, options)
}

// FormatName returns Name formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatName(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.Name, 140, options)
}

// FormatIdentificationNumber returns IdentificationNumber formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatIdentificationNumber(options FormatOptions) string {
	return ro.formatAlphaField(ro.IdentificationNumber, 35, options)
}

// FormatIdentificationNumberIssuer returns IdentificationNumberIssuer formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatIdentificationNumberIssuer(options FormatOptions) string {
	return ro.formatAlphaField(ro.IdentificationNumberIssuer, 35, options)
}

// FormatDateBirthPlace returns DateBirthPlace formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatDateBirthPlace(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.DateBirthPlace, 82, options)
}

// FormatAddressType returns AddressType formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressType(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressType, 4, options)
}

// FormatDepartment returns Department formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatDepartment(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.Department, 70, options)
}

// FormatSubDepartment returns SubDepartment formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatSubDepartment(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.SubDepartment, 70, options)
}

// FormatStreetName returns StreetName formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatStreetName(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.StreetName, 70, options)
}

// FormatBuildingNumber returns BuildingNumber formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatBuildingNumber(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.BuildingNumber, 16, options)
}

// FormatPostCode returns PostCode formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatPostCode(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.PostCode, 16, options)
}

// FormatTownName returns TownName formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatTownName(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.TownName, 35, options)
}

// FormatCountrySubDivisionState returns CountrySubDivisionState formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatCountrySubDivisionState(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.CountrySubDivisionState, 35, options)
}

// FormatCountry returns Country formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatCountry(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.Country, 2, options)
}

// FormatAddressLineOne returns AddressLineOne formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineOne(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineOne, 70, options)
}

// FormatAddressLineTwo returns AddressLineTwo formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineTwo(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineTwo, 70, options)
}

// FormatAddressLineThree returns AddressLineThree formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineThree(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineThree, 70, options)
}

// FormatAddressLineFour returns AddressLineFour formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineFour(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineOne, 70, options)
}

// FormatAddressLineFive returns AddressLineFive formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineFive(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineFive, 70, options)
}

// FormatAddressLineSix returns AddressLineSix formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineSix(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineSix, 70, options)
}

// FormatAddressLineSeven returns AddressLineSeven formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatAddressLineSeven(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.AddressLineSeven, 70, options)
}

// FormatCountryOfResidence returns CountryOfResidence formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatCountryOfResidence(options FormatOptions) string {
	return ro.formatAlphaField(ro.RemittanceData.CountryOfResidence, 2, options)
}

// FormatContactName returns ContactName formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactName(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactName, 140, options)
}

// FormatContactPhoneNumber returns ContactPhoneNumber formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactPhoneNumber(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactPhoneNumber, 35, options)
}

// FormatContactMobileNumber returns ContactMobileNumber formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactMobileNumber(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactMobileNumber, 35, options)
}

// FormatContactFaxNumber returns ContactFaxNumber formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactFaxNumber(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactFaxNumber, 35, options)
}

// FormatContactElectronicAddress returns ContactElectronicAddress formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactElectronicAddress(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactElectronicAddress, 2048, options)
}

// FormatContactOther returns ContactOther formatted according to the FormatOptions
func (ro *RemittanceOriginator) FormatContactOther(options FormatOptions) string {
	return ro.formatAlphaField(ro.ContactOther, 35, options)
}
