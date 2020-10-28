package wire

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// DateRemittanceDocument creates a DateRemittanceDocument
func mockDateRemittanceDocument() *DateRemittanceDocument {
	drd := NewDateRemittanceDocument()
	drd.DateRemittanceDocument = time.Now().Format("20060102")
	return drd
}

// TestMockDateRemittanceDocument validates mockDateRemittanceDocument
func TestMockDateRemittanceDocument(t *testing.T) {
	drd := mockDateRemittanceDocument()

	require.NoError(t, drd.Validate(), "mockDateRemittanceDocument does not validate and will break other tests")
}

// TestDateRemittanceDocumentRequired validates DateRemittanceDocument DateRemittanceDocument is required
func TestDateRemittanceDocumentDateRemittanceDocumentRequired(t *testing.T) {
	drd := mockDateRemittanceDocument()
	drd.DateRemittanceDocument = ""

	err := drd.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("DateRemittanceDocument", ErrFieldRequired).Error(), err.Error())
}

// TestParseDateRemittanceDocumentWrongLength parses a wrong DateRemittanceDocument record length
func TestParseDateRemittanceDocumentWrongLength(t *testing.T) {
	var line = "{8650}20190509  "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseDateRemittanceDocument()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), NewTagWrongLengthErr(14, len(r.line)).Error())
}

// TestParseDateRemittanceDocumentReaderParseError parses a wrong DateRemittanceDocument reader parse error
func TestParseDateRemittanceDocumentReaderParseError(t *testing.T) {
	var line = "{8650}14190509"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseDateRemittanceDocument()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrValidDate.Error())

	_, err = r.Read()

	require.NotNil(t, err)
	require.Contains(t, err.Error(), ErrValidDate.Error())
}

// TestDateRemittanceDocumentTagError validates a DateRemittanceDocument tag
func TestDateRemittanceDocumentTagError(t *testing.T) {
	drd := mockDateRemittanceDocument()
	drd.tag = "{9999}"

	err := drd.Validate()

	require.NotNil(t, err)
	require.Equal(t, fieldError("tag", ErrValidTagForType, drd.tag).Error(), err.Error())
}
