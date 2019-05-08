package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockFIIntermediaryFIAdvice creates a FIIntermediaryFIAdvice
func mockFIIntermediaryFIAdvice() *FIIntermediaryFIAdvice {
	fiifia := NewFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = AdviceCodeLetter
	fiifia.Advice.LineOne = "Line One"
	fiifia.Advice.LineTwo = "Line Two"
	fiifia.Advice.LineThree = "Line Three"
	fiifia.Advice.LineFour = "Line Four"
	fiifia.Advice.LineFive = "Line Five"
	fiifia.Advice.LineSix = "Line Six"
	return fiifia
}

// TestMockFIIntermediaryFIAdvice validates mockFIIntermediaryFIAdvice
func TestMockFIIntermediaryFIAdvice(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	if err := fiifia.Validate(); err != nil {
		t.Error("mockFIIntermediaryFIAdvice does not validate and will break other tests")
	}
}

// TestFIIntermediaryFIAdviceAdviceCodeValid validates FIIntermediaryFIAdvice AdviceCode is alphanumeric
func TestFIIntermediaryFIAdviceAdviceCodeValid(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.AdviceCode = "Z"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrAdviceCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineOneAlphaNumeric validates FIIntermediaryFIAdvice LineOne is alphanumeric
func TestFIIntermediaryFIAdviceLineOneAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineOne = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineTwoAlphaNumeric validates FIIntermediaryFIAdvice LineTwo is alphanumeric
func TestFIIntermediaryFIAdviceLineTwoAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineTwo = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineThreeAlphaNumeric validates FIIntermediaryFIAdvice LineThree is alphanumeric
func TestFIIntermediaryFIAdviceLineThreeAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineThree = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineFourAlphaNumeric validates FIIntermediaryFIAdvice LineFour is alphanumeric
func TestFIIntermediaryFIAdviceLineFourAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineFour = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineFiveAlphaNumeric validates FIIntermediaryFIAdvice LineFive is alphanumeric
func TestFIIntermediaryFIAdviceLineFiveAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineFive = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIIntermediaryFIAdviceLineSixAlphaNumeric validates FIIntermediaryFIAdvice LineSix is alphanumeric
func TestFIIntermediaryFIAdviceLineSixAlphaNumeric(t *testing.T) {
	fiifia := mockFIIntermediaryFIAdvice()
	fiifia.Advice.LineSix = "®"
	if err := fiifia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIIntermediaryFIAdviceWrongLength parses a wrong FIIntermediaryFIAdvice record length
func TestParseFIIntermediaryFIAdviceWrongLength(t *testing.T) {
	var line = "{6210}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	err := r.parseFIIntermediaryFIAdvice()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(200, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIIntermediaryFIAdviceReaderParseError parses a wrong FIIntermediaryFIAdvice reader parse error
func TestParseFIIntermediaryFIAdviceReaderParseError(t *testing.T) {
	var line = "{6210}LTRLine ®ne                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fiifia := mockFIIntermediaryFIAdvice()
	fwm.SetFIIntermediaryFIAdvice(fiifia)
	err := r.parseFIIntermediaryFIAdvice()
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
