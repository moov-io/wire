package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
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
	if err := sr.Validate(); err != nil {
		t.Error("mockSenderToReceiver does not validate and will break other tests")
	}
}

// TestSenderToReceiverSwiftFieldTagAlphaNumeric validates SenderToReceiver SwiftFieldTag is alphanumeric
func TestSenderToReceiverSwiftFieldTagAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftFieldTag = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineOneAlphaNumeric validates SenderToReceiver SwiftLineOne is alphanumeric
func TestSenderToReceiverSwiftLineOneAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineOne = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineTwoAlphaNumeric validates SenderToReceiver SwiftLineTwo is alphanumeric
func TestSenderToReceiverSwiftLineTwoAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineTwo = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineThreeAlphaNumeric validates SenderToReceiver SwiftLineThree is alphanumeric
func TestSenderToReceiverSwiftLineThreeAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineThree = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineFourAlphaNumeric validates SenderToReceiver SwiftLineFour is alphanumeric
func TestSenderToReceiverSwiftLineFourAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineFour = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineFiveAlphaNumeric validates SenderToReceiver SwiftLineFive is alphanumeric
func TestSenderToReceiverSwiftLineFiveAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineFive = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSenderToReceiverSwiftLineSixAlphaNumeric validates SenderToReceiver SwiftLineSix is alphanumeric
func TestSenderToReceiverSwiftLineSixAlphaNumeric(t *testing.T) {
	sr := mockSenderToReceiver()
	sr.CoverPayment.SwiftLineSix = "®"
	if err := sr.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderToReceiverWrongLength parses a wrong SenderToReceiver record length
func TestParseSenderToReceiverWrongLength(t *testing.T) {
	var line = "{7072}SwiftSwift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    Swift Line Six                   "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sr := mockSenderToReceiver()
	fwm.SetSenderToReceiver(sr)
	err := r.parseSenderToReceiver()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(221, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseSenderToReceiverReaderParseError parses a wrong SenderToReceiver reader parse error
func TestParseSenderToReceiverReaderParseError(t *testing.T) {
	var line = "{7072}Swift®wift Line One                     Swift Line Two                     Swift Line Three                   Swift Line Four                    Swift Line Five                    Swift Line Six                     "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sr := mockSenderToReceiver()
	fwm.SetSenderToReceiver(sr)
	err := r.parseSenderToReceiver()
	if err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
	_, err = r.Read()
	if err != nil {
		if !base.Has(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
