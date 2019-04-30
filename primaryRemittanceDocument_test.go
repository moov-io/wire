package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// PrimaryRemittanceDocument creates a PrimaryRemittanceDocument
func mockPrimaryRemittanceDocument() *PrimaryRemittanceDocument {
	prd := NewPrimaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = ""
	prd.DocumentIdentificationNumber = "111111"
	prd.Issuer = "Issuer"
	return prd
}

// TestMockPrimaryRemittanceDocument validates mockPrimaryRemittanceDocument
func TestMockPrimaryRemittanceDocument(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	if err := prd.Validate(); err != nil {
		t.Error("mockPrimaryRemittanceDocument does not validate and will break other tests")
	}
}

// TestDocumentTypeCodeValid validates PrimaryRemittanceDocument DocumentTypeCode
func TestDocumentTypeCodeValid(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = "ZZZZ"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrDocumentTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryDocumentTypeCodeAlphaNumeric validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is alphanumeric
func TestProprietaryDocumentTypeCodeAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestDocumentIdentificationNumberAlphaNumeric validates PrimaryRemittanceDocument DocumentIdentificationNumber is alphanumeric
func TestDocumentIdentificationNumberAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentIdentificationNumber = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIssuerAlphaNumeric validates PrimaryRemittanceDocument Issuer is alphanumeric
func TestIssuerAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.Issuer = "®"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryDocumentTypeCodeRequired validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is required
func TestProprietaryDocumentTypeCodeRequired (t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = ""
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestDocumentIdentificationNumberRequired validates PrimaryRemittanceDocument DocumentIdentificationNumber is required
func TestDocumentIdentificationNumberRequired (t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentIdentificationNumber = ""
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryDocumentTypeCodeInvalid validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is invalid
func TestProprietaryDocumentTypeCodeInvalid (t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = "Proprietary"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}