package wire

import (
	"github.com/moov-io/base"
	"strings"
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
		if !base.Match(err, ErrNonNumeric) {
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

// TestParseInputMessageAccountabilityDataWrongLength parses a wrong InputMessageAccountabilityData record length
func TestParseInputMessageAccountabilityDataWrongLength(t *testing.T) {
	var line = "{1510}1"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	err := r.parseInputMessageAccountabilityData()
	if err != nil {

		if !base.Match(err, NewTagWrongLengthErr(28, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseInputMessageAccountabilityDataReaderParseError parses a wrong InputMessageAccountabilityData reader parse error
func TestParseInputMessageAccountabilityDataReaderParseError(t *testing.T) {
	var line = "{1520}20190507Source0800000Z"
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	imad := mockInputMessageAccountabilityData()
	fwm.SetInputMessageAccountabilityData(imad)
	err := r.parseInputMessageAccountabilityData()
	if err != nil {
		if !base.Match(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonNumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}