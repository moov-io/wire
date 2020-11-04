package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// mockTypeSubType creates a TypeSubType
func mockTypeSubType() *TypeSubType {
	tst := NewTypeSubType()
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = BasicFundsTransfer
	return tst
}

// TestTypeSubType validates mockTypeSubType
func TestMockTypeSubType(t *testing.T) {
	tst := mockTypeSubType()

	require.NoError(t, tst.Validate(), "mockTypeSubType does not validate and will break other tests")
}

// TestTypeCodeValid validates TypeSubType TypeCode
func TestTypeCodeValid(t *testing.T) {
	tst := mockTypeSubType()
	tst.TypeCode = "ZZ"

	err := tst.Validate()

	require.EqualError(t, err, fieldError("TypeCode", ErrTypeCode, tst.TypeCode).Error())
}

// TestSubTypeCodeValid validates TypeSubType SubTypeCode
func TestSubTypeCodeValid(t *testing.T) {
	tst := mockTypeSubType()
	tst.SubTypeCode = "ZZ"

	err := tst.Validate()

	require.EqualError(t, err, fieldError("SubTypeCode", ErrSubTypeCode, tst.SubTypeCode).Error())
}

// TestTypeCodeRequired validates TypeSubType TypeCode is required
func TestTypeCodeCodeRequired(t *testing.T) {
	tst := mockTypeSubType()
	tst.TypeCode = ""

	err := tst.Validate()

	require.EqualError(t, err, fieldError("TypeCode", ErrFieldRequired).Error())
}

// TestSubTypeCodeRequired validates TypeSubType SubTypeCode is required
func TestSubTypeCodeRequired(t *testing.T) {
	tst := mockTypeSubType()
	tst.SubTypeCode = ""

	err := tst.Validate()

	require.EqualError(t, err, fieldError("SubTypeCode", ErrFieldRequired).Error())
}

// TestParseTypeSubTypeWrongLength parses a wrong TypeSubType record length
func TestParseTypeSubTypeWrongLength(t *testing.T) {
	var line = "{1510}1"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseTypeSubType()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(10, len(r.line))).Error())
}

// TestParseTypeSubTypeReaderParseError parses a wrong TypeSubType reader parse error
func TestParseTypeSubTypeReaderParseError(t *testing.T) {
	var line = "{1510}100Z"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseTypeSubType()

	require.EqualError(t, err, r.parseError(fieldError("SubTypeCode", ErrSubTypeCode, "0Z")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SubTypeCode", ErrSubTypeCode, "0Z")).Error())
}

// TestTypeSubTypeTagError validates a TypeSubType tag
func TestTypeSubTypeTagError(t *testing.T) {
	tst := mockTypeSubType()
	tst.tag = "{9999}"

	require.EqualError(t, tst.Validate(), fieldError("tag", ErrValidTagForType, tst.tag).Error())
}
