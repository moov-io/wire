package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockSenderSupplied creates a SenderSupplied
func mockSenderSupplied() *SenderSupplied {
	ss := NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = MessageDuplicationOriginal
	return ss
}

// TestMockSenderSupplied validates mockSenderSupplied
func TestMockSenderSupplied(t *testing.T) {
	ss := mockSenderSupplied()
	if err := ss.Validate(); err != nil {
		t.Error("mockSenderSupplied does not validate and will break other tests")
	}
}

// TestSenderSuppliedUserRequestCorrelationAlphaNumeric validates SenderSupplied UserRequestCorrelation is alphanumeric
func TestSenderSuppliedUserRequestCorrelationAlphaNumeric(t *testing.T) {
	ss := mockSenderSupplied()
	ss.UserRequestCorrelation = "®"
	if err := ss.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderSuppliedFormatVersionValid validates SenderSupplied FormatVersion
func TestSenderSuppliedFormatVersionValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.FormatVersion = "55"
	if err := ss.Validate(); err != nil {
		if !base.Match(err, ErrFormatVersion) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderSuppliedProductionCodeValid validates SenderSupplied ProductionCode
func TestSenderSuppliedProductionCodeValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.TestProductionCode = "Z"
	if err := ss.Validate(); err != nil {
		if !base.Match(err, ErrTestProductionCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderSuppliedMessageDuplicationCodeValid validates SenderSupplied MessageDuplicationCode
func TestSenderSuppliedMessageDuplicationCodeValid(t *testing.T) {
	ss := mockSenderSupplied()
	ss.MessageDuplicationCode = "Z"
	if err := ss.Validate(); err != nil {
		if !base.Match(err, ErrMessageDuplicationCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderSuppliedUserRequestCorrelationRequired validates SenderSupplied UserRequestCorrelation is required
func TestSenderSuppliedUserRequestCorrelationRequired(t *testing.T) {
	ss := mockSenderSupplied()
	ss.UserRequestCorrelation = ""
	if err := ss.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
