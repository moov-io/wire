package wire

import (
	"github.com/moov-io/base"
	"strings"
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
func TestProprietaryDocumentTypeCodeRequired(t *testing.T) {
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
func TestDocumentIdentificationNumberRequired(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentIdentificationNumber = ""
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestProprietaryDocumentTypeCodeInvalid validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is invalid
func TestProprietaryDocumentTypeCodeInvalid(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = "Proprietary"
	if err := prd.Validate(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePrimaryRemittanceDocumentWrongLength parses a wrong PrimaryRemittanceDocument record length
func TestParsePrimaryRemittanceDocumentWrongLength(t *testing.T) {
	var line = "{8400}AROI                                   111111                             Issuer                           "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	err := r.parsePrimaryRemittanceDocument()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(115, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParsePrimaryRemittanceDocumentReaderParseError parses a wrong PrimaryRemittanceDocument reader parse error
func TestParsePrimaryRemittanceDocumentReaderParseError(t *testing.T) {
	var line = "{8400}ZZZZ                                   111111                             Issuer                             "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	prd := mockPrimaryRemittanceDocument()
	fwm.SetPrimaryRemittanceDocument(prd)
	err := r.parsePrimaryRemittanceDocument()
	if err != nil {
		if !base.Match(err, ErrDocumentTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrDocumentTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
