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
func (ro *RemittanceOriginator) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 29 {
		return 0, NewTagWrongLengthErr(29, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if ro.tag, read, err = ro.parseTag(record); err != nil {
		return 0, fieldError("RemittanceOriginator.Tag", err)
	}
	length += read

	if ro.IdentificationType, read, err = ro.parseVariableStringField(record[length:], 2); err != nil {
		fieldError("IdentificationType", err)
	}
	length += read

	if ro.IdentificationCode, read, err = ro.parseVariableStringField(record[length:], 4); err != nil {
		fieldError("IdentificationCode", err)
	}
	length += read

	if ro.Name, read, err = ro.parseVariableStringField(record[length:], 140); err != nil {
		fieldError("Name", err)
	}
	length += read

	if ro.IdentificationNumber, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("IdentificationNumber", err)
	}
	length += read

	if ro.IdentificationNumberIssuer, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("IdentificationNumberIssuer", err)
	}
	length += read

	if read, err = ro.RemittanceData.ParseForRemittanceBeneficiary(record[length:]); err != nil {
		return 0, err
	}
	length += read

	if ro.ContactName, read, err = ro.parseVariableStringField(record[length:], 140); err != nil {
		fieldError("ContactName", err)
	}
	length += read

	if ro.ContactPhoneNumber, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("ContactPhoneNumber", err)
	}
	length += read

	if ro.ContactMobileNumber, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("ContactMobileNumber", err)
	}
	length += read

	if ro.ContactFaxNumber, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("ContactFaxNumber", err)
	}
	length += read

	if ro.ContactElectronicAddress, read, err = ro.parseVariableStringField(record[length:], 2048); err != nil {
		fieldError("ContactElectronicAddress", err)
	}
	length += read

	if ro.ContactOther, read, err = ro.parseVariableStringField(record[length:], 35); err != nil {
		fieldError("ContactOther", err)
	}
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
func (ro *RemittanceOriginator) String(options ...bool) string {

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(3442)

	buf.WriteString(ro.tag)
	buf.WriteString(ro.IdentificationTypeField(isCompressed))
	buf.WriteString(ro.IdentificationCodeField(isCompressed))
	buf.WriteString(ro.NameField(isCompressed))
	buf.WriteString(ro.IdentificationNumberField(isCompressed))
	buf.WriteString(ro.IdentificationNumberIssuerField(isCompressed))
	buf.WriteString(ro.RemittanceData.StringForRemittanceBeneficiary(isCompressed))
	buf.WriteString(ro.ContactNameField(isCompressed))
	buf.WriteString(ro.ContactPhoneNumberField(isCompressed))
	buf.WriteString(ro.ContactMobileNumberField(isCompressed))
	buf.WriteString(ro.ContactFaxNumberField(isCompressed))
	buf.WriteString(ro.ContactElectronicAddressField(isCompressed))
	buf.WriteString(ro.ContactOtherField(isCompressed))

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
	if err := ro.isAlphanumeric(ro.Name); err != nil {
		return fieldError("Name", err, ro.Name)
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
	if ro.Name == "" {
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
func (ro *RemittanceOriginator) IdentificationTypeField(isCompressed bool) string {
	return ro.alphaVariableField(ro.IdentificationType, 2, isCompressed)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (ro *RemittanceOriginator) IdentificationCodeField(isCompressed bool) string {
	return ro.alphaVariableField(ro.IdentificationCode, 4, isCompressed)
}

// NameField gets a string of the Name field
func (ro *RemittanceOriginator) NameField(isCompressed bool) string {
	return ro.alphaVariableField(ro.Name, 140, isCompressed)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (ro *RemittanceOriginator) IdentificationNumberField(isCompressed bool) string {
	return ro.alphaVariableField(ro.IdentificationNumber, 35, isCompressed)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (ro *RemittanceOriginator) IdentificationNumberIssuerField(isCompressed bool) string {
	return ro.alphaVariableField(ro.IdentificationNumberIssuer, 35, isCompressed)
}

// ContactNameField gets a string of the ContactName field
func (ro *RemittanceOriginator) ContactNameField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactName, 140, isCompressed)
}

// ContactPhoneNumberField gets a string of the ContactPhoneNumber field
func (ro *RemittanceOriginator) ContactPhoneNumberField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactPhoneNumber, 35, isCompressed)
}

// ContactMobileNumberField gets a string of the ContactMobileNumber field
func (ro *RemittanceOriginator) ContactMobileNumberField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactMobileNumber, 35, isCompressed)
}

// ContactFaxNumberField gets a string of the ContactFaxNumber field
func (ro *RemittanceOriginator) ContactFaxNumberField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactFaxNumber, 35, isCompressed)
}

// ContactElectronicAddressField gets a string of the ContactElectronicAddress field
func (ro *RemittanceOriginator) ContactElectronicAddressField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactElectronicAddress, 2048, isCompressed)
}

// ContactOtherField gets a string of the ContactOther field
func (ro *RemittanceOriginator) ContactOtherField(isCompressed bool) string {
	return ro.alphaVariableField(ro.ContactOther, 35, isCompressed)
}
