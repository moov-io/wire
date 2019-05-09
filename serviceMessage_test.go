package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockServiceMessage creates a ServiceMessage
func mockServiceMessage() *ServiceMessage {
	sm := NewServiceMessage()
	sm.LineOne = "Line One"
	sm.LineTwo = "Line Two"
	sm.LineThree = "Line Three"
	sm.LineFour = "Line Four"
	sm.LineFive = "Line Five"
	sm.LineSix = "Line Six"
	sm.LineSeven = "Line Seven"
	sm.LineEight = "Line Eight"
	sm.LineNine = "Line Nine"
	sm.LineTen = "Line Ten"
	sm.LineEleven = "Line Eleven"
	sm.LineTwelve = "line Twelve"
	return sm
}

// TestMockServiceMessage validates mockServiceMessage
func TestMockServiceMessage(t *testing.T) {
	sm := mockServiceMessage()
	if err := sm.Validate(); err != nil {
		t.Error("mockServiceMessage does not validate and will break other tests")
	}
}

// TestLineOneAlphaNumeric validates ServiceMessage LineOne is alphanumeric
func TestLineOneAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineOne = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineTwoAlphaNumeric validates ServiceMessage LineTwo is alphanumeric
func TestLineTwoAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTwo = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineThreeAlphaNumeric validates ServiceMessage LineThree is alphanumeric
func TestLineThreeAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineThree = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineFourAlphaNumeric validates ServiceMessage LineFour is alphanumeric
func TestLineFourAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineFour = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineFiveAlphaNumeric validates ServiceMessage LineFive is alphanumeric
func TestLineFiveAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineFive = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineSixAlphaNumeric validates ServiceMessage LineSix is alphanumeric
func TestLineSixAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineSix = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineSevenAlphaNumeric validates ServiceMessage LineSeven is alphanumeric
func TestLineSevenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineSeven = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineEightAlphaNumeric validates ServiceMessage LineEight is alphanumeric
func TestLineEightAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineEight = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineNineAlphaNumeric validates ServiceMessage LineNine is alphanumeric
func TestLineNineAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineNine = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineTenAlphaNumeric validates ServiceMessage LineTen is alphanumeric
func TestLineTenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTen = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineElevenAlphaNumeric validates ServiceMessage LineEleven is alphanumeric
func TestLineElevenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineEleven = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineTwelveAlphaNumeric validates ServiceMessage LineTwelve is alphanumeric
func TestLineTwelveAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTwelve = "®"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestLineOneRequired validates ServiceMessage LineOne is required
func TestLineOneRequired(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineOne = ""
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrFieldRequired) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseServiceMessageWrongLength parses a wrong ServiceMessage record length
func TestParseServiceMessageWrongLength(t *testing.T) {
	var line = "{9000}Line One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           Line Seven                         Line Eight                         Line Nine                          Line Ten                           Line Eleven                        line Twelve                      "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	err := r.parseServiceMessage()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(426, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseServiceMessageReaderParseError parses a wrong ServiceMessage reader parse error
func TestParseServiceMessageReaderParseError(t *testing.T) {
	var line = "{9000}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           Line Seven                         Line Eight                         Line Nine                          Line Ten                           Line Eleven                        line Twelve                        "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	err := r.parseServiceMessage()
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
