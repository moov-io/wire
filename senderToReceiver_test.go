package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// SenderToReceiver creates a SenderToReceiver
func mockSenderToReceiver() *SenderToReceiver {
	sr := NewSenderToReceiver()
	sr.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	sr.CoverPayment.SwiftLineOne = "Swift Line One"
	sr.CoverPayment.SwiftLineTwo = "Swift Line Two"
	sr.CoverPayment.SwiftLineThree = "Swift Line Three"
	sr.CoverPayment.SwiftLineFour = "Swift Line Four"
	sr.CoverPayment.SwiftLineFive = "Swift Line Five"
	sr.CoverPayment.SwiftLineSix = "Swift Line Six"
	return sr
}

// TestMockSenderToReceiver validates mockSenderToReceiver
func TestMockSenderToReceiver(t *testing.T) {
	sr := mockSenderToReceiver()

	require.NoError(t, sr.Validate(), "mockSenderToReceiver does not validate and will break other tests")
}

// TestSenderToReceiverSwiftFieldTagAlphaNumeric validates SenderToReceiver SwiftFieldTag is alphanumeric
func TestSenderToReceiverSwiftFieldTagAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftFieldTag = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftFieldTag", ErrNonAlphanumeric, sr.CoverPayment.SwiftFieldTag).Error())
}

// TestSenderToReceiverSwiftLineOneAlphaNumeric validates SenderToReceiver SwiftLineOne is alphanumeric
func TestSenderToReceiverSwiftLineOneAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineOne = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineOne", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineOne).Error())
}

// TestSenderToReceiverSwiftLineTwoAlphaNumeric validates SenderToReceiver SwiftLineTwo is alphanumeric
func TestSenderToReceiverSwiftLineTwoAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineTwo = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineTwo", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineTwo).Error())
}

// TestSenderToReceiverSwiftLineThreeAlphaNumeric validates SenderToReceiver SwiftLineThree is alphanumeric
func TestSenderToReceiverSwiftLineThreeAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineThree = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineThree", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineThree).Error())
}

// TestSenderToReceiverSwiftLineFourAlphaNumeric validates SenderToReceiver SwiftLineFour is alphanumeric
func TestSenderToReceiverSwiftLineFourAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineFour = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFour", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineFour).Error())
}

// TestSenderToReceiverSwiftLineFiveAlphaNumeric validates SenderToReceiver SwiftLineFive is alphanumeric
func TestSenderToReceiverSwiftLineFiveAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineFive = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineFive", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineFive).Error())
}

// TestSenderToReceiverSwiftLineSixAlphaNumeric validates SenderToReceiver SwiftLineSix is alphanumeric
func TestSenderToReceiverSwiftLineSixAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineSix = "®"

	err := sr.Validate()

	require.EqualError(t, err, fieldError("SwiftLineSix", ErrNonAlphanumeric, sr.CoverPayment.SwiftLineSix).Error())
}

// TestParseSenderToReceiverWrongLength parses a wrong SenderToReceiver record length
func TestParseSenderToReceiverWrongLength(t *testing.T) {
	var line = "{7072}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    Swift Line Six                   "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderToReceiver()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(221, len(r.line))).Error())
}

// TestParseSenderToReceiverReaderParseError parses a wrong SenderToReceiver reader parse error
func TestParseSenderToReceiverReaderParseError(t *testing.T) {
	var line = "{7072}Swift®wift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    Swift Line Six                     "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseSenderToReceiver()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "®wift Line One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("SwiftLineOne", ErrNonAlphanumeric, "®wift Line One")).Error())
}

// TestSenderToReceiverTagError validates a SenderToReceiver tag
func TestSenderToReceiverTagError(t *testing.T) {
	str := mockSenderToReceiver()
	str.tag = "{9999}"

	require.EqualError(t, str.Validate(), fieldError("tag", ErrValidTagForType, str.tag).Error())
}
