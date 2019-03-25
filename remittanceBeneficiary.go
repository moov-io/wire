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
	IdentificationType string`json:"identificationType,omitempty"`
	// IdentificationCode  Organization Identification Codes  * `BANK` - Bank Party Identification * `CUST` - Customer Number * `DUNS` - Data Universal Number System (Dun & Bradstreet) * `EMPL` - Employer Identification Number * `GS1G` - Global Location Number * `PROP` - Proprietary Identification Number * `SWBB` - SWIFT BIC or BEI * `TXID` - Tax Identification Number  Private Identification Codes  * `ARNU` - Alien Registration Number * `CCPT` - Passport Number * `CUST` - Customer Number * `DPOB` - Date & Place of Birth * `DRLC` - Driver’s License Number * `EMPL` - Employee Identification Number * `NIDN` - National Identity Number * `PROP` - Proprietary Identification Number * `SOSE` - Social Security Number * `TXID` - Tax Identification Number
	IdentificationCode string `json:"identificationCode,omitempty"`
	// IdentificationNumber
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// IdentificationNumberIssuer
	IdentificationNumberIssuer string `json:"identificationNumberIssuer,omitempty"`
	// DateAndfBirthPlace
	DateAndBirthPlace string `json:"dateAndBirthPlace,omitempty"`
	RemittanceData RemittanceData `json:"remittanceData,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewRemittanceBeneficiary returns a new RemittanceBeneficiary
func NewRemittanceBeneficiary() RemittanceBeneficiary  {
	rb := RemittanceBeneficiary {
		tag: TagRemittanceBeneficiary,
	}
	return rb
}

// Parse takes the input string and parses the RemittanceBeneficiary values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (rb *RemittanceBeneficiary) Parse(record string) {
}

// String writes RemittanceBeneficiary
func (rb *RemittanceBeneficiary) String() string {
	var buf strings.Builder
	// ToDo: Separator
	buf.Grow(1108)
	buf.WriteString(rb.tag)
	return buf.String()
}

// Validate performs WIRE format rule checks on RemittanceBeneficiary and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (rb *RemittanceBeneficiary) Validate() error {
	if err := rb.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (rb *RemittanceBeneficiary) fieldInclusion() error {
	return nil
}
