package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, prd.Validate(), "mockPrimaryRemittanceDocument does not validate and will break other tests")
}

// TestDocumentTypeCodeValid validates PrimaryRemittanceDocument DocumentTypeCode
func TestDocumentTypeCodeValid(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = "ZZZZ"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentTypeCode", ErrDocumentTypeCode, prd.DocumentTypeCode).Error())
}

// TestProprietaryDocumentTypeCodeAlphaNumeric validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is alphanumeric
func TestProprietaryDocumentTypeCodeAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrNonAlphanumeric, prd.ProprietaryDocumentTypeCode).Error())
}

// TestDocumentIdentificationNumberAlphaNumeric validates PrimaryRemittanceDocument DocumentIdentificationNumber is alphanumeric
func TestDocumentIdentificationNumberAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentIdentificationNumber = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentIdentificationNumber", ErrNonAlphanumeric, prd.DocumentIdentificationNumber).Error())
}

// TestIssuerAlphaNumeric validates PrimaryRemittanceDocument Issuer is alphanumeric
func TestIssuerAlphaNumeric(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.Issuer = "®"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("Issuer", ErrNonAlphanumeric, prd.Issuer).Error())
}

// TestProprietaryDocumentTypeCodeRequired validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is required
func TestProprietaryDocumentTypeCodeRequired(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = ProprietaryDocumentType
	prd.ProprietaryDocumentTypeCode = ""

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrFieldRequired).Error())
}

// TestDocumentIdentificationNumberRequired validates PrimaryRemittanceDocument DocumentIdentificationNumber is required
func TestDocumentIdentificationNumberRequired(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentIdentificationNumber = ""

	err := prd.Validate()

	require.EqualError(t, err, fieldError("DocumentIdentificationNumber", ErrFieldRequired).Error())
}

// TestProprietaryDocumentTypeCodeInvalid validates PrimaryRemittanceDocument ProprietaryDocumentTypeCode is invalid
func TestProprietaryDocumentTypeCodeInvalid(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = "Proprietary"

	err := prd.Validate()

	require.EqualError(t, err, fieldError("ProprietaryDocumentTypeCode", ErrInvalidProperty, prd.ProprietaryDocumentTypeCode).Error())
}

// TestParsePrimaryRemittanceDocumentWrongLength parses a wrong PrimaryRemittanceDocument record length
func TestParsePrimaryRemittanceDocumentWrongLength(t *testing.T) {
	var line = "{8400}AROI                                   111111                             Issuer                           "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePrimaryRemittanceDocument()

	require.EqualError(t, err, r.parseError(fieldError("Issuer", ErrValidLength)).Error())
}

// TestParsePrimaryRemittanceDocumentReaderParseError parses a wrong PrimaryRemittanceDocument reader parse error
func TestParsePrimaryRemittanceDocumentReaderParseError(t *testing.T) {
	var line = "{8400}ZZZZ                                   111111                             Issuer                             "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePrimaryRemittanceDocument()

	require.EqualError(t, err, r.parseError(fieldError("DocumentTypeCode", ErrDocumentTypeCode, "ZZZZ")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("DocumentTypeCode", ErrDocumentTypeCode, "ZZZZ")).Error())
}

// TestPrimaryRemittanceDocumentTagError validates a PrimaryRemittanceDocument tag
func TestPrimaryRemittanceDocumentTagError(t *testing.T) {
	prd := mockPrimaryRemittanceDocument()
	prd.tag = "{9999}"

	require.EqualError(t, prd.Validate(), fieldError("tag", ErrValidTagForType, prd.tag).Error())
}

// TestStringPrimaryRemittanceDocumentVariableLength parses using variable length
func TestStringPrimaryRemittanceDocumentVariableLength(t *testing.T) {
	var line = "{8400}AROI*Issuer*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePrimaryRemittanceDocument()
	require.Nil(t, err)

	line = "{8400}AROI                                   Issuer                                                                NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePrimaryRemittanceDocument()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{8400}CMCN********"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePrimaryRemittanceDocument()
	require.EqualError(t, err, r.parseError(NewTagMaxLengthErr()).Error())

	line = "{8400}AROI*Issuer*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parsePrimaryRemittanceDocument()
	require.Equal(t, err, nil)
}

// TestStringPrimaryRemittanceDocumentOptions validates string() with options
func TestStringPrimaryRemittanceDocumentOptions(t *testing.T) {
	var line = "{8400}AROI*Issuer*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parsePrimaryRemittanceDocument()
	require.Equal(t, err, nil)

	str := r.currentFEDWireMessage.PrimaryRemittanceDocument.String()
	require.Equal(t, str, "{8400}AROI                                   Issuer                                                                ")

	str = r.currentFEDWireMessage.PrimaryRemittanceDocument.String(true)
	require.Equal(t, str, "{8400}AROI*Issuer*")
}
