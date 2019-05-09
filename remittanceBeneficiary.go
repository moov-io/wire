// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
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
	if utf8.RuneCountInString(record) != 1114 {
		return NewTagWrongLengthErr(1114, len(record))
	}
	rb.tag = record[:6]
	rb.RemittanceData.Name = rb.parseStringField(record[6:146])
	rb.IdentificationType = rb.parseStringField(record[146:148])
	rb.IdentificationCode = rb.parseStringField(record[148:152])
	rb.IdentificationNumber = rb.parseStringField(record[152:187])
	rb.IdentificationNumberIssuer = rb.parseStringField(record[187:222])
	rb.RemittanceData.DateBirthPlace = rb.parseStringField(record[222:304])
	rb.RemittanceData.AddressType = rb.parseStringField(record[304:308])
	rb.RemittanceData.Department = rb.parseStringField(record[308:378])
	rb.RemittanceData.SubDepartment = rb.parseStringField(record[378:448])
	rb.RemittanceData.StreetName = rb.parseStringField(record[448:518])
	rb.RemittanceData.BuildingNumber = rb.parseStringField(record[518:534])
	rb.RemittanceData.PostCode = rb.parseStringField(record[534:550])
	rb.RemittanceData.TownName = rb.parseStringField(record[550:585])
	rb.RemittanceData.CountrySubDivisionState = rb.parseStringField(record[585:620])
	rb.RemittanceData.Country = rb.parseStringField(record[620:622])
	rb.RemittanceData.AddressLineOne = rb.parseStringField(record[622:692])
	rb.RemittanceData.AddressLineTwo = rb.parseStringField(record[692:762])
	rb.RemittanceData.AddressLineThree = rb.parseStringField(record[762:832])
	rb.RemittanceData.AddressLineFour = rb.parseStringField(record[832:902])
	rb.RemittanceData.AddressLineFive = rb.parseStringField(record[902:972])
	rb.RemittanceData.AddressLineSix = rb.parseStringField(record[972:1042])
	rb.RemittanceData.AddressLineSeven = rb.parseStringField(record[1042:1112])
	rb.RemittanceData.CountryOfResidence = rb.parseStringField(record[1112:1114])
	return nil
}

// String writes RemittanceBeneficiary
func (rb *RemittanceBeneficiary) String() string {
	var buf strings.Builder
	buf.Grow(1114)
	buf.WriteString(rb.tag)
	buf.WriteString(rb.NameField())
	buf.WriteString(rb.IdentificationTypeField())
	buf.WriteString(rb.IdentificationCodeField())
	buf.WriteString(rb.IdentificationNumberField())
	buf.WriteString(rb.IdentificationNumberIssuerField())
	buf.WriteString(rb.DateBirthPlaceField())
	buf.WriteString(rb.AddressTypeField())
	buf.WriteString(rb.DepartmentField())
	buf.WriteString(rb.SubDepartmentField())
	buf.WriteString(rb.StreetNameField())
	buf.WriteString(rb.BuildingNumberField())
	buf.WriteString(rb.PostCodeField())
	buf.WriteString(rb.TownNameField())
	buf.WriteString(rb.CountrySubDivisionStateField())
	buf.WriteString(rb.CountryField())
	buf.WriteString(rb.AddressLineOneField())
	buf.WriteString(rb.AddressLineTwoField())
	buf.WriteString(rb.AddressLineThreeField())
	buf.WriteString(rb.AddressLineFourField())
	buf.WriteString(rb.AddressLineFiveField())
	buf.WriteString(rb.AddressLineSixField())
	buf.WriteString(rb.AddressLineSevenField())
	buf.WriteString(rb.CountryOfResidenceField())
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rb *RemittanceBeneficiary) Validate() error {
	if err := rb.fieldInclusion(); err != nil {
		return err
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
		return fieldError("AddressLineSeven", err, rb.RemittanceData.CountryOfResidence)
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
