package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
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
	if err := tst.Validate(); err != nil {
		if !base.Match(err, ErrTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSubTypeCodeValid validates TypeSubType SubTypeCode
func TestSubTypeCodeValid(t *testing.T) {
	tst := mockTypeSubType()
	tst.SubTypeCode = "ZZ"
	if err := tst.Validate(); err != nil {
		if !base.Match(err, ErrSubTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestTypeCodeRequired validates TypeSubType TypeCode is required
func TestTypeCodeCodeRequired(t *testing.T) {
	tst := mockTypeSubType()
	tst.TypeCode = ""
	if err := tst.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSubTypeCodeRequired validates TypeSubType SubTypeCode is required
func TestSubTypeCodeRequired(t *testing.T) {
	tst := mockTypeSubType()
	tst.SubTypeCode = ""
	if err := tst.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseTypeSubTypeWrongLength parses a wrong TypeSubType record length
func TestParseTypeSubTypeWrongLength(t *testing.T) {
	var line = "{1510}1"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	err := r.parseTypeSubType()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(18, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseTypeSubTypeReaderParseError parses a wrong TypeSubType reader parse error
func TestParseTypeSubTypeReaderParseError(t *testing.T) {
	var line = "{1510}100Z"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	tst := mockTypeSubType()
	fwm.SetTypeSubType(tst)
	err := r.parseTypeSubType()
	if err != nil {
		if !base.Match(err, ErrSubTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrSubTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestTypeSubTypeTagError validates a TypeSubType tag
func TestTypeSubTypeTagError(t *testing.T) {
	tst := mockTypeSubType()
	tst.tag = "{9999}"

	require.EqualError(t, tst.Validate(), fieldError("tag", ErrValidTagForType, tst.tag).Error())
}
