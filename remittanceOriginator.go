// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
	// CountryOfResidence
	CountryOfResidence string `json:"countryOfResidence,omitempty"`
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
func (ro *RemittanceOriginator) Parse(record string) {
	ro.tag = record[:6]
	ro.IdentificationType = ro.parseStringField(record[6:8])
	ro.IdentificationCode = ro.parseStringField(record[8:12])
	ro.RemittanceData.Name = ro.parseStringField(record[12:152])
	ro.IdentificationNumber = ro.parseStringField(record[152:187])
	ro.IdentificationNumberIssuer = ro.parseStringField(record[187:222])
	ro.RemittanceData.DateBirthPlace = ro.parseStringField(record[222:304])
	ro.RemittanceData.AddressType = ro.parseStringField(record[304:308])
	ro.RemittanceData.Department = ro.parseStringField(record[308:374])
	ro.RemittanceData.SubDepartment = ro.parseStringField(record[374:444])
	ro.RemittanceData.StreetName = ro.parseStringField(record[444:514])
	ro.RemittanceData.BuildingNumber = ro.parseStringField(record[514:530])
	ro.RemittanceData.PostCode = ro.parseStringField(record[530:546])
	ro.RemittanceData.TownName = ro.parseStringField(record[546:581])
	ro.RemittanceData.CountrySubDivisionState = ro.parseStringField(record[581:616])
	ro.RemittanceData.Country = ro.parseStringField(record[616:618])
	ro.RemittanceData.AddressLineOne = ro.parseStringField(record[618:688])
	ro.RemittanceData.AddressLineTwo = ro.parseStringField(record[688:758])
	ro.RemittanceData.AddressLineThree = ro.parseStringField(record[758:828])
	ro.RemittanceData.AddressLineFour = ro.parseStringField(record[828:898])
	ro.RemittanceData.AddressLineFive = ro.parseStringField(record[898:968])
	ro.RemittanceData.AddressLineSix = ro.parseStringField(record[968:1038])
	ro.RemittanceData.AddressLineSeven = ro.parseStringField(record[1038:1108])
	ro.CountryOfResidence = ro.parseStringField(record[1108:1110])
	ro.ContactName = ro.parseStringField(record[1110:1250])
	ro.ContactPhoneNumber = ro.parseStringField(record[1250:1285])
	ro.ContactMobileNumber = ro.parseStringField(record[1285:1320])
	ro.ContactFaxNumber = ro.parseStringField(record[1320:1355])
	ro.ContactElectronicAddress = ro.parseStringField(record[1355:3403])
	ro.ContactOther = ro.parseStringField(record[3403:3438])
}

// String writes RemittanceOriginator
func (ro *RemittanceOriginator) String() string {
	var buf strings.Builder
	buf.Grow(3438)
	buf.WriteString(ro.tag)
	buf.WriteString(ro.IdentificationTypeField())
	buf.WriteString(ro.IdentificationCodeField())
	buf.WriteString(ro.NameField())
	buf.WriteString(ro.IdentificationNumberField())
	buf.WriteString(ro.IdentificationNumberIssuerField())
	buf.WriteString(ro.DateBirthPlaceField())
	buf.WriteString(ro.AddressTypeField())
	buf.WriteString(ro.DepartmentField())
	buf.WriteString(ro.SubDepartmentField())
	buf.WriteString(ro.StreetNameField())
	buf.WriteString(ro.BuildingNumberField())
	buf.WriteString(ro.PostCodeField())
	buf.WriteString(ro.TownNameField())
	buf.WriteString(ro.CountrySubDivisionStateField())
	buf.WriteString(ro.CountryField())
	buf.WriteString(ro.AddressLineOneField())
	buf.WriteString(ro.AddressLineTwoField())
	buf.WriteString(ro.AddressLineThreeField())
	buf.WriteString(ro.AddressLineFourField())
	buf.WriteString(ro.AddressLineFiveField())
	buf.WriteString(ro.AddressLineSixField())
	buf.WriteString(ro.AddressLineSevenField())
	buf.WriteString(ro.CountryOfResidenceField())
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
func (ro *RemittanceOriginator) Validate() error {
	if err := ro.fieldInclusion(); err != nil {
		return err
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

	if err := ro.isAlphanumeric(ro.CountryOfResidence); err != nil {
		return fieldError("CountryOfResidence", err, ro.CountryOfResidence)
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
	switch ro.IdentificationCode {
	case PICDateBirthPlace:
		if ro.IdentificationNumber != "" {
			return fieldError("IdentificationNumber", ErrInvalidProperty)
		}
		if ro.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty)
		}
	case OICSWIFTBICORBEI:
		if ro.IdentificationNumberIssuer != "" {
			return fieldError("IdentificationNumberIssuer", ErrInvalidProperty)
		}
		if ro.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty)
		}
	default:
		if ro.RemittanceData.DateBirthPlace != "" {
			return fieldError("DateBirthPlace", ErrInvalidProperty)
		}
		if ro.IdentificationNumber == "" {
			return fieldError("IdentificationNumber", ErrFieldRequired)
		}
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (ro *RemittanceOriginator) fieldInclusion() error {
	if ro.IdentificationType == "" {
		return fieldError("IdentificationType", ErrFieldRequired)
	}
	if ro.IdentificationCode == "" {
		return fieldError("IdentificationCode", ErrFieldRequired)
	}
	if ro.RemittanceData.Name == "" {
		return fieldError("Name", ErrFieldRequired)
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
	return ro.alphaField(ro.CountryOfResidence, 2)
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
