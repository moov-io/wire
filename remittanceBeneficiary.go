// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

var _ segment = &RemittanceBeneficiary{}

// RemittanceBeneficiary is remittance beneficiary
type RemittanceBeneficiary struct {
	// tag
	tag string
	// Name
	Name string `json:"name,omitempty"`
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
func (rb *RemittanceBeneficiary) Parse(record string) (int, error) {
	if utf8.RuneCountInString(record) < 29 {
		return 0, NewTagWrongLengthErr(29, utf8.RuneCountInString(record))
	}

	var err error
	var length, read int

	if rb.tag, read, err = rb.parseTag(record); err != nil {
		return 0, fieldError("RemittanceBeneficiary.Tag", err)
	}
	length += read

	if rb.Name, read, err = rb.parseVariableStringField(record[length:], 140); err != nil {
		return 0, fieldError("Name", err)
	}
	length += read

	if rb.IdentificationType, read, err = rb.parseVariableStringField(record[length:], 2); err != nil {
		return 0, fieldError("IdentificationType", err)
	}
	length += read

	if rb.IdentificationCode, read, err = rb.parseVariableStringField(record[length:], 4); err != nil {
		return 0, fieldError("IdentificationCode", err)
	}
	length += read

	if rb.IdentificationNumber, read, err = rb.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("IdentificationNumber", err)
	}
	length += read

	if rb.IdentificationNumberIssuer, read, err = rb.parseVariableStringField(record[length:], 35); err != nil {
		return 0, fieldError("IdentificationNumberIssuer", err)
	}
	length += read

	if read, err = rb.RemittanceData.ParseForRemittanceBeneficiary(record[length:]); err != nil {
		return 0, err
	}
	length += read

	return length, nil
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

	isCompressed := false
	if len(options) > 0 {
		isCompressed = options[0]
	}

	var buf strings.Builder
	buf.Grow(1114)

	buf.WriteString(rb.tag)
	buf.WriteString(rb.NameField(isCompressed))
	buf.WriteString(rb.IdentificationTypeField(isCompressed))
	buf.WriteString(rb.IdentificationCodeField(isCompressed))
	buf.WriteString(rb.IdentificationNumberField(isCompressed))
	buf.WriteString(rb.IdentificationNumberIssuerField(isCompressed))
	buf.WriteString(rb.RemittanceData.StringForRemittanceBeneficiary(isCompressed))

	return buf.String()
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
	if err := rb.isAlphanumeric(rb.Name); err != nil {
		return fieldError("Name", err, rb.Name)
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
	if rb.Name == "" {
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
func (rb *RemittanceBeneficiary) NameField(isCompressed bool) string {
	return rb.alphaVariableField(rb.Name, 140, isCompressed)
}

// IdentificationTypeField gets a string of the IdentificationType field
func (rb *RemittanceBeneficiary) IdentificationTypeField(isCompressed bool) string {
	return rb.alphaVariableField(rb.IdentificationType, 2, isCompressed)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (rb *RemittanceBeneficiary) IdentificationCodeField(isCompressed bool) string {
	return rb.alphaVariableField(rb.IdentificationCode, 4, isCompressed)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (rb *RemittanceBeneficiary) IdentificationNumberField(isCompressed bool) string {
	return rb.alphaVariableField(rb.IdentificationNumber, 35, isCompressed)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (rb *RemittanceBeneficiary) IdentificationNumberIssuerField(isCompressed bool) string {
	return rb.alphaVariableField(rb.IdentificationNumberIssuer, 35, isCompressed)
}
