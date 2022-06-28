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

// String returns a fixed-width RemittanceBeneficiary record
func (rb *RemittanceBeneficiary) String() string {
	return rb.Format(FormatOptions{
		VariableLengthFields: false,
	})
}

// Format returns a RemittanceBeneficiary record formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) Format(options FormatOptions) string {
	var buf strings.Builder
	buf.Grow(1114)

	buf.WriteString(rb.tag)
	buf.WriteString(rb.FormatName(options))
	buf.WriteString(rb.FormatIdentificationType(options))
	buf.WriteString(rb.FormatIdentificationCode(options))
	buf.WriteString(rb.FormatIdentificationNumber(options))
	buf.WriteString(rb.FormatIdentificationNumberIssuer(options))
	buf.WriteString(rb.FormatDateBirthPlace(options))
	buf.WriteString(rb.FormatAddressType(options))
	buf.WriteString(rb.FormatDepartment(options))
	buf.WriteString(rb.FormatSubDepartment(options))
	buf.WriteString(rb.FormatStreetName(options))
	buf.WriteString(rb.FormatBuildingNumber(options))
	buf.WriteString(rb.FormatPostCode(options))
	buf.WriteString(rb.FormatTownName(options))
	buf.WriteString(rb.FormatCountrySubDivisionState(options))
	buf.WriteString(rb.FormatCountry(options))
	buf.WriteString(rb.FormatAddressLineOne(options))
	buf.WriteString(rb.FormatAddressLineTwo(options))
	buf.WriteString(rb.FormatAddressLineThree(options))
	buf.WriteString(rb.FormatAddressLineFour(options))
	buf.WriteString(rb.FormatAddressLineFive(options))
	buf.WriteString(rb.FormatAddressLineSix(options))
	buf.WriteString(rb.FormatAddressLineSeven(options))
	buf.WriteString(rb.FormatCountryOfResidence(options))

	if options.VariableLengthFields {
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
func (rb *RemittanceBeneficiary) NameField() string {
	return rb.alphaField(rb.RemittanceData.Name, 140)
}

// IdentificationTypeField gets a string of the IdentificationType field
func (rb *RemittanceBeneficiary) IdentificationTypeField() string {
	return rb.alphaField(rb.IdentificationType, 2)
}

// IdentificationCodeField gets a string of the IdentificationCode field
func (rb *RemittanceBeneficiary) IdentificationCodeField() string {
	return rb.alphaField(rb.IdentificationCode, 4)
}

// IdentificationNumberField gets a string of the IdentificationNumber field
func (rb *RemittanceBeneficiary) IdentificationNumberField() string {
	return rb.alphaField(rb.IdentificationNumber, 35)
}

// IdentificationNumberIssuerField gets a string of the IdentificationNumberIssuer field
func (rb *RemittanceBeneficiary) IdentificationNumberIssuerField() string {
	return rb.alphaField(rb.IdentificationNumberIssuer, 35)
}

// DateBirthPlaceField gets a string of the DateBirthPlace field
func (rb *RemittanceBeneficiary) DateBirthPlaceField() string {
	return rb.alphaField(rb.RemittanceData.DateBirthPlace, 82)
}

// AddressTypeField gets a string of the AddressType field
func (rb *RemittanceBeneficiary) AddressTypeField() string {
	return rb.alphaField(rb.RemittanceData.AddressType, 4)
}

// DepartmentField gets a string of the Department field
func (rb *RemittanceBeneficiary) DepartmentField() string {
	return rb.alphaField(rb.RemittanceData.Department, 70)
}

// SubDepartmentField gets a string of the SubDepartment field
func (rb *RemittanceBeneficiary) SubDepartmentField() string {
	return rb.alphaField(rb.RemittanceData.SubDepartment, 70)
}

// StreetNameField gets a string of the StreetName field
func (rb *RemittanceBeneficiary) StreetNameField() string {
	return rb.alphaField(rb.RemittanceData.StreetName, 70)
}

// BuildingNumberField gets a string of the BuildingNumber field
func (rb *RemittanceBeneficiary) BuildingNumberField() string {
	return rb.alphaField(rb.RemittanceData.BuildingNumber, 16)
}

// PostCodeField gets a string of the PostCode field
func (rb *RemittanceBeneficiary) PostCodeField() string {
	return rb.alphaField(rb.RemittanceData.PostCode, 16)
}

// TownNameField gets a string of the TownName field
func (rb *RemittanceBeneficiary) TownNameField() string {
	return rb.alphaField(rb.RemittanceData.TownName, 35)
}

// CountrySubDivisionStateField gets a string of the CountrySubDivisionState field
func (rb *RemittanceBeneficiary) CountrySubDivisionStateField() string {
	return rb.alphaField(rb.RemittanceData.CountrySubDivisionState, 35)
}

// CountryField gets a string of the Country field
func (rb *RemittanceBeneficiary) CountryField() string {
	return rb.alphaField(rb.RemittanceData.Country, 2)
}

// AddressLineOneField gets a string of the AddressLineOne field
func (rb *RemittanceBeneficiary) AddressLineOneField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineOne, 70)
}

// AddressLineTwoField gets a string of the AddressLineTwo field
func (rb *RemittanceBeneficiary) AddressLineTwoField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineTwo, 70)
}

// AddressLineThreeField gets a string of the AddressLineThree field
func (rb *RemittanceBeneficiary) AddressLineThreeField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineThree, 70)
}

// AddressLineFourField gets a string of the AddressLineFour field
func (rb *RemittanceBeneficiary) AddressLineFourField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineFour, 70)
}

// AddressLineFiveField gets a string of the AddressLineFive field
func (rb *RemittanceBeneficiary) AddressLineFiveField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineFive, 70)
}

// AddressLineSixField gets a string of the AddressLineSix field
func (rb *RemittanceBeneficiary) AddressLineSixField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineSix, 70)
}

// AddressLineSevenField gets a string of the AddressLineSeven field
func (rb *RemittanceBeneficiary) AddressLineSevenField() string {
	return rb.alphaField(rb.RemittanceData.AddressLineSeven, 70)
}

// CountryOfResidenceField gets a string of the CountryOfResidence field
func (rb *RemittanceBeneficiary) CountryOfResidenceField() string {
	return rb.alphaField(rb.RemittanceData.CountryOfResidence, 2)
}

// FormatName returns Name formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatName(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.Name, 140, options)
}

// FormatIdentificationType returns IdentificationType formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatIdentificationType(options FormatOptions) string {
	return rb.formatAlphaField(rb.IdentificationType, 2, options)
}

// FormatIdentificationCode returns IdentificationCode formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatIdentificationCode(options FormatOptions) string {
	return rb.formatAlphaField(rb.IdentificationCode, 4, options)
}

// FormatIdentificationNumber returns IdentificationNumber formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatIdentificationNumber(options FormatOptions) string {
	return rb.formatAlphaField(rb.IdentificationNumber, 35, options)
}

// FormatIdentificationNumberIssuer returns IdentificationNumberIssuer formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatIdentificationNumberIssuer(options FormatOptions) string {
	return rb.formatAlphaField(rb.IdentificationNumberIssuer, 35, options)
}

// FormatDateBirthPlace returns DateBirthPlace formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatDateBirthPlace(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.DateBirthPlace, 82, options)
}

// FormatAddressType returns AddressType formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressType(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressType, 4, options)
}

// FormatDepartment returns Department formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatDepartment(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.Department, 70, options)
}

// FormatSubDepartment returns SubDepartment formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatSubDepartment(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.SubDepartment, 70, options)
}

// FormatStreetName returns StreetName formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatStreetName(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.StreetName, 70, options)
}

// FormatBuildingNumber returns BuildingNumber formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatBuildingNumber(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.BuildingNumber, 16, options)
}

// FormatPostCode returns PostCode formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatPostCode(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.PostCode, 16, options)
}

// FormatTownName returns TownName formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatTownName(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.TownName, 35, options)
}

// FormatCountrySubDivisionState returns CountrySubDivisionState formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatCountrySubDivisionState(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.CountrySubDivisionState, 35, options)
}

// FormatCountry returns Country formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatCountry(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.Country, 2, options)
}

// FormatAddressLineOne returns AddressLineOne formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineOne(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineOne, 70, options)
}

// FormatAddressLineTwo returns AddressLineTwo formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineTwo(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineTwo, 70, options)
}

// FormatAddressLineThree returns AddressLineThree formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineThree(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineThree, 70, options)
}

// FormatAddressLineFour returns AddressLineFour formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineFour(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineFour, 70, options)
}

// FormatAddressLineFive returns AddressLineFive formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineFive(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineFive, 70, options)
}

// FormatAddressLineSix returns AddressLineSix formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineSix(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineSix, 70, options)
}

// FormatAddressLineSeven returns AddressLineSeven formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatAddressLineSeven(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.AddressLineSeven, 70, options)
}

// FormatCountryOfResidence returns CountryOfResidence formatted according to the FormatOptions
func (rb *RemittanceBeneficiary) FormatCountryOfResidence(options FormatOptions) string {
	return rb.formatAlphaField(rb.RemittanceData.CountryOfResidence, 2, options)
}
