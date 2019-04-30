package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// SecondaryRemittanceDocument creates a SecondaryRemittanceDocument
func mockSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	srd := NewSecondaryRemittanceDocument()
	srd.DocumentTypeCode = StatementAccount
	srd.ProprietaryDocumentTypeCode = ""
	srd.DocumentIdentificationNumber = "222222"
	srd.Issuer = "Issuer 2"
	return srd
}

// TestMockSecondaryRemittanceDocument validates mockSecondaryRemittanceDocument
func TestMockSecondaryRemittanceDocument(t *testing.T) {
	srd := mockSecondaryRemittanceDocument()
	if err := srd.Validate(); err != nil {
		t.Error("mockSecondaryRemittanceDocument does not validate and will break other tests")
	}
}

// TestSRDDocumentTypeCodeValid validates SecondaryRemittanceDocument DocumentTypeCode
func TestSRDDocumentTypeCodeValid(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = "ZZZZ"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrDocumentTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDProprietaryDocumentTypeCodeAlphaNumeric validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is alphanumeric
func TestSRDProprietaryDocumentTypeCodeAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDDocumentIdentificationNumberAlphaNumeric validates SecondaryRemittanceDocument DocumentIdentificationNumber is alphanumeric
func TestSRDDocumentIdentificationNumberAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentIdentificationNumber = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDIssuerAlphaNumeric validates SecondaryRemittanceDocument Issuer is alphanumeric
func TestSRDIssuerAlphaNumeric(t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.Issuer = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDProprietaryDocumentTypeCodeRequired validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is required
func TestSRDProprietaryDocumentTypeCodeRequired (t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = ""
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDDocumentIdentificationNumberRequired validates SecondaryRemittanceDocument DocumentIdentificationNumber is required
func TestSRDDocumentIdentificationNumberRequired (t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentIdentificationNumber = ""
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSRDProprietaryDocumentTypeCodeInvalid validates SecondaryRemittanceDocument ProprietaryDocumentTypeCode is invalid
func TestSRDProprietaryDocumentTypeCodeInvalid (t *testing.T) {
	prd := mockSecondaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = "Proprietary"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}