// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

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
	// DateAndfBirthPlace
	DateAndBirthPlace string         `json:"dateAndBirthPlace,omitempty"`
	RemittanceData    RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceBeneficiary returns a new RemittanceBeneficiary
func NewRemittanceBeneficiary() RemittanceBeneficiary {
	rb := RemittanceBeneficiary{
		tag: TagRemittanceBeneficiary,
	}
	return rb
}

// Parse takes the input string and parses the RemittanceBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rb *RemittanceBeneficiary) Parse(record string) {
	rb.tag = record[:6]
	rb.RemittanceData.Name = rb.parseStringField(record[6:146])
	rb.IdentificationType = rb.parseStringField(record[146:148])
	rb.IdentificationCode = rb.parseStringField(record[148:154])
	rb.IdentificationNumber = rb.parseStringField(record[154:189])
	rb.IdentificationNumberIssuer = rb.parseStringField(record[189:224])
	//rb.RemittanceData.DateBirthPlace = rb.parseStringField(record[224:306])
	rb.RemittanceData.AddressType = rb.parseStringField(record[306:310])
	rb.RemittanceData.Department = rb.parseStringField(record[310:380])
	rb.RemittanceData.SubDepartment = rb.parseStringField(record[380:450])
	rb.RemittanceData.StreetName = rb.parseStringField(record[450:520])
	rb.RemittanceData.BuildingNumber = rb.parseStringField(record[520:536])
	rb.RemittanceData.PostCode = rb.parseStringField(record[536:552])
	rb.RemittanceData.TownName = rb.parseStringField(record[552:587])
	rb.RemittanceData.CountrySubDivisionState = rb.parseStringField(record[587:622])
	rb.RemittanceData.Country = rb.parseStringField(record[622:624])
	rb.RemittanceData.AddressLineOne = rb.parseStringField(record[624:694])
	rb.RemittanceData.AddressLineTwo = rb.parseStringField(record[694:764])
	rb.RemittanceData.AddressLineThree = rb.parseStringField(record[764:834])
	rb.RemittanceData.AddressLineFour = rb.parseStringField(record[834:904])
	rb.RemittanceData.AddressLineFive = rb.parseStringField(record[904:974])
	rb.RemittanceData.AddressLineSix = rb.parseStringField(record[974:1044])
	rb.RemittanceData.AddressLineSeven = rb.parseStringField(record[1044:1114])
}

// String writes RemittanceBeneficiary
func (rb *RemittanceBeneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(1114)
	buf.WriteString(rb.tag)
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
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rb *RemittanceBeneficiary) fieldInclusion() error {
	return nil
}
