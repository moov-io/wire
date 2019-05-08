package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockFIDrawdownDebitAccountAdvice creates a FIDrawdownDebitAccountAdvice
func mockFIDrawdownDebitAccountAdvice() *FIDrawdownDebitAccountAdvice {
	debitDDAdvice := NewFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.AdviceCode = AdviceCodeLetter
	debitDDAdvice.Advice.LineOne = "Line One"
	debitDDAdvice.Advice.LineTwo = "Line Two"
	debitDDAdvice.Advice.LineThree = "Line Three"
	debitDDAdvice.Advice.LineFour = "Line Four"
	debitDDAdvice.Advice.LineFive = "Line Five"
	debitDDAdvice.Advice.LineSix = "Line Six"
	return debitDDAdvice
}

// TestMockFIDrawdownDebitAccountAdvice validates mockFIDrawdownDebitAccountAdvice
func TestMockFIDrawdownDebitAccountAdvice(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	if err := debitDDAdvice.Validate(); err != nil {
		t.Error("mockFIDrawdownDebitAccountAdvice does not validate and will break other tests")
	}
}

// TestFIDrawdownDebitAccountAdviceAdviceCodeValid validates FIDrawdownDebitAccountAdvice AdviceCode is alphanumeric
func TestFIDrawdownDebitAccountAdviceCodeValid(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.AdviceCode = "Z"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrAdviceCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric validates FIDrawdownDebitAccountAdvice LineOne is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineOneAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineOne = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric validates FIDrawdownDebitAccountAdvice LineTwo is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineTwoAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineTwo = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric validates FIDrawdownDebitAccountAdvice LineThree is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineThreeAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineThree = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFour is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFourAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFour = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric validates FIDrawdownDebitAccountAdvice LineFive is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineFiveAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineFive = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric validates FIDrawdownDebitAccountAdvice LineSix is alphanumeric
func TestFIDrawdownDebitAccountAdviceLineSixAlphaNumeric(t *testing.T) {
	debitDDAdvice := mockFIDrawdownDebitAccountAdvice()
	debitDDAdvice.Advice.LineSix = "®"
	if err := debitDDAdvice.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIDrawdownDebitAccountAdviceWrongLength parses a wrong FIDrawdownDebitAccountAdvice record length
func TestParseFIDrawdownDebitAccountAdviceWrongLength(t *testing.T) {
	var line = "{6110}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fibfia := mockFIDrawdownDebitAccountAdvice()
	fwm.SetFIDrawdownDebitAccountAdvice(fibfia)
	err := r.parseFIDrawdownDebitAccountAdvice()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(200, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIDrawdownDebitAccountAdviceReaderParseError parses a wrong FIDrawdownDebitAccountAdvice reader parse error
func TestParseFIDrawdownDebitAccountAdviceReaderParseError(t *testing.T) {
	var line = "{6110}LTR®ine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fibfia := mockFIDrawdownDebitAccountAdvice()
	fwm.SetFIDrawdownDebitAccountAdvice(fibfia)
	err := r.parseFIDrawdownDebitAccountAdvice()
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
