// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import "strings"

// BusinessFunctionCode {3600}
type BusinessFunctionCode struct {
	// tag
	tag string
	// BusinessFunctionCode BTR: Bank Transfer (Beneficiary is a bank) DRC: Customer or Corporate Drawdown Request CKS: Check Same Day Settlement DRW: Drawdown Payment CTP: Customer Transfer Plus FFR: Fed Funds Returned CTR: Customer Transfer (Beneficiary is a not a bank) FFS: Fed Funds Sold DEP: Deposit to Sender’s Account SVC: Service Message DRB: Bank-to-Bank Drawdown Request
	BusinessFunctionCode string `json:"businessFunctionCode"`
	// TransactionTypeCode If {3600} is CTR, an optional Transaction Type Code element is permitted; however, the Transaction Type Code 'COV' is not permitted.
	// ToDo: Research This, I don't understand it yet
	TransactionTypeCode string `json:"transactionTypeCode,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewBusinessFunctionCode returns a new BusinessFunctionCode
func NewBusinessFunctionCode() *BusinessFunctionCode {
	bfc := &BusinessFunctionCode{
		tag: TagBusinessFunctionCode,
	}
	return bfc
}

// Parse takes the input string and parses the BusinessFunctionCode values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (bfc *BusinessFunctionCode) Parse(record string) {
	bfc.tag = record[:6]
	bfc.BusinessFunctionCode = bfc.parseStringField(record[6:9])
	bfc.TransactionTypeCode = record[9:12]
}

// String writes BusinessFunctionCode
func (bfc *BusinessFunctionCode) String() string {
	var buf strings.Builder
	buf.Grow(12)
	buf.WriteString(bfc.tag)
	buf.WriteString(bfc.BusinessFunctionCode)
	buf.WriteString(bfc.TransactionTypeCode)
	return buf.String()
}

// Validate performs WIRE format rule checks on BusinessFunctionCode and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bfc *BusinessFunctionCode) Validate() error {
	if err := bfc.fieldInclusion(); err != nil {
		return err
	}
	if err := bfc.isBusinessFunctionCode(bfc.BusinessFunctionCode); err != nil {
		return fieldError("BusinessFunctionCode", err, bfc.BusinessFunctionCode)
	}
	// transactionCodeType does not seem to be defined in the spec

	/*	if err := bfc.isTransactionTypeCode(bfc.TransactionTypeCode); err != nil {
		return fieldError("TransactionTypeCode", err, bfc.TransactionTypeCode)
	}*/
	return nil
}

// fieldInclusion validate mandatory fields. If fields are
// invalid the WIRE will return an error.
func (bfc *BusinessFunctionCode) fieldInclusion() error {

	// only element 01 (BusinessFunctionCode) is required
	if bfc.BusinessFunctionCode == "" {
		return fieldError("BusinessFunctionCode", ErrFieldRequired, bfc.BusinessFunctionCode)
	}
	return nil
}

// BusinessFunctionCodeField gets a string of the BusinessFunctionCode field
func (bfc *BusinessFunctionCode) BusinessFunctionCodeField() string {
	return bfc.alphaField(bfc.BusinessFunctionCode, 3)
}

// TransactionTypeCodeField gets a string of the TransactionTypeCode field
func (bfc *BusinessFunctionCode) TransactionTypeCodeField() string {
	return bfc.alphaField(bfc.TransactionTypeCode, 3)
}
