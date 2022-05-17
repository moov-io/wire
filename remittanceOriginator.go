// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &RemittanceOriginator{}

// RemittanceOriginator is remittance originator
type RemittanceOriginator struct {
	// tag
	tag string
	// is variable length
	isVariableLength bool
	// IdentificationType is identification type
	IdentificationType string `json:"identificationType,omitempty"`
	// IdentificationCode  Organization Identification Codes  * `BANK` - Bank Party Identification * `CUST` - Customer Number * `DUNS` - Data Universal Number System (Dun & Bradstreet) * `EMPL` - Employer Identification Number * `GS1G` - Global Location Number * `PROP` - Proprietary Identification Number * `SWBB` - SWIFT BIC or BEI * `TXID` - Tax Identification Number  Private Identification Codes  * `ARNU` - Alien Registration Number * `CCPT` - Passport Number * `CUST` - Customer Number * `DPOB` - Date & Place of Birth * `DRLC` - Driver’s License Number * `EMPL` - Employee Identification Number * `NIDN` - National Identity Number * `PROP` - Proprietary Identification Number * `SOSE` - Social Security Number * `TXID` - Tax Identification Number
	IdentificationCode string `json:"identificationCode,omitempty"`
	// Name
	Name string `json:"name,omitempty"`
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
func NewRemittanceOriginator(isVariable bool) *RemittanceOriginator {
	ro := &RemittanceOriginator{
		tag:              TagRemittanceOriginator,
		isVariableLength: isVariable,
	}
	return ro
}

// Parse takes the input string and parses the RemittanceOriginator values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (ro *RemittanceOriginator) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 29 {
		return 0, NewTagWrongLengthErr(29, utf8.RuneCountInString(record))
	}

	ro.tag = record[:6]

	length := 6
	read := 0

	ro.IdentificationType, read = ro.parseVariableStringField(record[length:], 2)
	length += read

	ro.IdentificationCode, read = ro.parseVariableStringField(record[length:], 4)
	length += read

	ro.Name, read = ro.parseVariableStringField(record[length:], 140)
	length += read

	ro.IdentificationNumber, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	ro.IdentificationNumberIssuer, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	read = ro.RemittanceData.ParseForRemittanceBeneficiary(record[length:])
	length += read

	ro.ContactName, read = ro.parseVariableStringField(record[length:], 140)
	length += read

	ro.ContactPhoneNumber, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	ro.ContactMobileNumber, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	ro.ContactFaxNumber, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	ro.ContactElectronicAddress, read = ro.parseVariableStringField(record[length:], 2048)
	length += read

	ro.ContactOther, read = ro.parseVariableStringField(record[length:], 35)
	length += read

	return length, nil
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
func (ro *RemittanceOriginator) String() string {
	var buf strings.Builder
	buf.Grow(3442)

	buf.WriteString(ro.tag)
	buf.WriteString(ro.IdentificationTypeField())
	buf.WriteString(ro.IdentificationCodeField())
	buf.WriteString(ro.NameField())
	buf.WriteString(ro.IdentificationNumberField())
	buf.WriteString(ro.IdentificationNumberIssuerField())
	buf.WriteString(ro.RemittanceData.StringForRemittanceBeneficiary(ro.isVariableLength))
	buf.WriteString(ro.ContactNameField())
	buf.WriteString(ro.ContactPhoneNumberField())
	buf.WriteString(ro.ContactMobileNumberField())
	buf.WriteString(ro.ContactFaxNumberField())
	buf.WriteString(ro.ContactElectronicAddressField())
	buf.WriteString(ro.ContactOtherField())

	return buf.String()
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
	return ro.alphaVariableField(ro.IdentificationType, 2, ro.isVariableLength)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ro *RemittanceOriginator) IdentificationCodeField() string {
	return ro.alphaVariableField(ro.IdentificationCode, 4, ro.isVariableLength)
}

// NameField gets a string of the Name field
func (ro *RemittanceOriginator) NameField() string {
	return ro.alphaVariableField(ro.Name, 140, ro.isVariableLength)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (ro *RemittanceOriginator) IdentificationNumberField() string {
	return ro.alphaVariableField(ro.IdentificationNumber, 35, ro.isVariableLength)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (ro *RemittanceOriginator) IdentificationNumberIssuerField() string {
	return ro.alphaVariableField(ro.IdentificationNumberIssuer, 35, ro.isVariableLength)
}

// ContactNameField gets a string of the ContactName field
func (ro *RemittanceOriginator) ContactNameField() string {
	return ro.alphaVariableField(ro.ContactName, 140, ro.isVariableLength)
}

// ContactPhoneNumberField gets a string of the ContactPhoneNumber field
func (ro *RemittanceOriginator) ContactPhoneNumberField() string {
	return ro.alphaVariableField(ro.ContactPhoneNumber, 35, ro.isVariableLength)
}

// ContactMobileNumberField gets a string of the ContactMobileNumber field
func (ro *RemittanceOriginator) ContactMobileNumberField() string {
	return ro.alphaVariableField(ro.ContactMobileNumber, 35, ro.isVariableLength)
}

// ContactFaxNumberField gets a string of the ContactFaxNumber field
func (ro *RemittanceOriginator) ContactFaxNumberField() string {
	return ro.alphaVariableField(ro.ContactFaxNumber, 35, ro.isVariableLength)
}

// ContactElectronicAddressField gets a string of the ContactElectronicAddress field
func (ro *RemittanceOriginator) ContactElectronicAddressField() string {
	return ro.alphaVariableField(ro.ContactElectronicAddress, 2048, ro.isVariableLength)
}

// ContactOtherField gets a string of the ContactOther field
func (ro *RemittanceOriginator) ContactOtherField() string {
	return ro.alphaVariableField(ro.ContactOther, 35, ro.isVariableLength)
}
