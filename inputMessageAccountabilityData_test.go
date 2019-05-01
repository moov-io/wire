package wire

import (
	"github.com/moov-io/base"
	"testing"
	"time"
)

// mockInputMessageAccountabilityData creates a mockInputMessageAccountabilityData
func mockInputMessageAccountabilityData() *InputMessageAccountabilityData {
	imad := NewInputMessageAccountabilityData()
	imad.InputCycleDate = time.Now().Format("20060102")
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	return imad
}

// TestMockInputMessageAccountabilityData validates mockInputMessageAccountabilityData
func TestMockInputMessageAccountabilityData(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	if err := imad.Validate(); err != nil {
		t.Error("mockInputMessageAccountabilityData does not validate and will break other tests")
	}
}

// TestInputMessageAccountabilityDataInputCycleDateRequired validates InputMessageAccountabilityData InputCycleDate is required
func TestInputMessageAccountabilityDataInputCycleDateRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputCycleDate = ""
	if err := imad.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInputMessageAccountabilityDataInputSourceAlphaNumeric validates InputMessageAccountabilityData InputSource is
// AlphaNumeric
func TestInputMessageAccountabilityDataInputSourceAlphaNumeric(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSource = "®"
	if err := imad.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInputMessageAccountabilityDataInputSequenceNumberAlphaNumeric validates InputMessageAccountabilityData InputSequenceNumber is
// AlphaNumeric
func TestInputMessageAccountabilityDataInputSequenceNumberAlphaNumeric(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSequenceNumber = "®"
	if err := imad.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInputMessageAccountabilityDataInputSourceRequired validates InputMessageAccountabilityData InputSource is required
func TestInputMessageAccountabilityDataInputSourceRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSource = ""
	if err := imad.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInputMessageAccountabilityDataInputSequenceNumberRequired validates InputMessageAccountabilityData
// InputSequenceNumber is required
func TestInputMessageAccountabilityDataInputSequenceNumberRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSequenceNumber = ""
	if err := imad.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
