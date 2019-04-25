package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockTypeSubType creates a TypeSubType
func mockTypeSubType() *TypeSubType {
	tst := NewTypeSubType()
	tst.tag = TagTypeSubType
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = BasicFundsTransfer
	return tst
}

// TestTypeSubType validates mockTypeSubType
func TestMockTypeSubType(t *testing.T) {
	tst := mockTypeSubType()
	if err := tst.Validate(); err != nil {
		t.Error("mockTypeSubType does not validate and will break other tests")
	}
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