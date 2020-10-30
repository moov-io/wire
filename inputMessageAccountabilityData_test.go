package wire

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, imad.Validate(), "mockInputMessageAccountabilityData does not validate and will break other tests")
}

// TestInputMessageAccountabilityDataInputCycleDateRequired validates InputMessageAccountabilityData InputCycleDate is required
func TestInputMessageAccountabilityDataInputCycleDateRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputCycleDate = ""

	require.EqualError(t, imad.Validate(), fieldError("InputCycleDate", ErrFieldRequired, imad.InputCycleDate).Error())
}

// TestInputMessageAccountabilityDataInputSourceAlphaNumeric validates InputMessageAccountabilityData InputSource is
// AlphaNumeric
func TestInputMessageAccountabilityDataInputSourceAlphaNumeric(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSource = "®"

	require.EqualError(t, imad.Validate(), fieldError("InputSource", ErrNonAlphanumeric, imad.InputSource).Error())
}

// TestInputMessageAccountabilityDataInputSequenceNumberAlphaNumeric validates InputMessageAccountabilityData InputSequenceNumber is
// AlphaNumeric
func TestInputMessageAccountabilityDataInputSequenceNumberAlphaNumeric(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSequenceNumber = "®"

	require.EqualError(t, imad.Validate(), fieldError("InputSequenceNumber", ErrNonNumeric, imad.InputSequenceNumber).Error())
}

// TestInputMessageAccountabilityDataInputSourceRequired validates InputMessageAccountabilityData InputSource is required
func TestInputMessageAccountabilityDataInputSourceRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSource = ""

	require.EqualError(t, imad.Validate(), fieldError("InputSource", ErrFieldRequired, imad.InputSource).Error())
}

// TestInputMessageAccountabilityDataInputSequenceNumberRequired validates InputMessageAccountabilityData
// InputSequenceNumber is required
func TestInputMessageAccountabilityDataInputSequenceNumberRequired(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputSequenceNumber = ""

	require.EqualError(t, imad.Validate(), fieldError("InputSequenceNumber", ErrFieldRequired, imad.InputSequenceNumber).Error())
}

// TestParseInputMessageAccountabilityDataWrongLength parses a wrong InputMessageAccountabilityData record length
func TestParseInputMessageAccountabilityDataWrongLength(t *testing.T) {
	var line = "{1510}1"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInputMessageAccountabilityData()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(28, len(r.line))).Error())
}

// TestParseInputMessageAccountabilityDataReaderParseError parses a wrong InputMessageAccountabilityData reader parse error
func TestParseInputMessageAccountabilityDataReaderParseError(t *testing.T) {
	var line = "{1520}20190507Source0800000Z"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseInputMessageAccountabilityData()

	require.EqualError(t, err, r.parseError(fieldError("InputSequenceNumber", ErrNonNumeric, "00000Z")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("InputSequenceNumber", ErrNonNumeric, "00000Z")).Error())
}

// TestInputMessageAccountabilityDataTagError validates a InputMessageAccountabilityData tag
func TestInputMessageAccountabilityDataTagError(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.tag = "{9999}"

	require.EqualError(t, imad.Validate(), fieldError("tag", ErrValidTagForType, imad.tag).Error())
}

// TestInputMessageAccountabilityDataInputCycleDateError validates a InputMessageAccountabilityData InputCycleDate
func TestInputMessageAccountabilityDataInputCycleDateError(t *testing.T) {
	imad := mockInputMessageAccountabilityData()
	imad.InputCycleDate = "02010101"

	require.EqualError(t, imad.Validate(), fieldError("InputCycleDate", ErrValidDate, imad.InputCycleDate).Error())
}
