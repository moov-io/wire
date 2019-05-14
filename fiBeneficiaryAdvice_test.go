package wire

import (
	"github.com/moov-io/base"
	"strings"
	"testing"
)

// mockFIBeneficiaryAdvice creates a FIBeneficiaryAdvice
func mockFIBeneficiaryAdvice() *FIBeneficiaryAdvice {
	fiba := NewFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = AdviceCodeLetter
	fiba.Advice.LineOne = "Line One"
	fiba.Advice.LineTwo = "Line Two"
	fiba.Advice.LineThree = "Line Three"
	fiba.Advice.LineFour = "Line Four"
	fiba.Advice.LineFive = "Line Five"
	fiba.Advice.LineSix = "Line Six"
	return fiba
}

// TestMockFIBeneficiaryAdvice validates mockFIBeneficiaryAdvice
func TestMockFIBeneficiaryAdvice(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	if err := fiba.Validate(); err != nil {
		t.Error("mockFIBeneficiary does not validate and will break other tests")
	}
}

// TestFIBeneficiaryAdviceCodeValid validates FIBeneficiaryAdvice AdviceCode is alphanumeric
func TestFIBeneficiaryAdviceCodeValid(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.AdviceCode = "Z"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrAdviceCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineOneAlphaNumeric validates FIBeneficiaryAdvice LineOne is alphanumeric
func TestFIBeneficiaryAdviceLineOneAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineOne = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineTwoAlphaNumeric validates FIBeneficiaryAdvice LineTwo is alphanumeric
func TestFIBeneficiaryAdviceLineTwoAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineTwo = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineThreeAlphaNumeric validates FIBeneficiaryAdvice LineThree is alphanumeric
func TestFIBeneficiaryAdviceLineThreeAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineThree = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineFourAlphaNumeric validates FIBeneficiaryAdvice LineFour is alphanumeric
func TestFIBeneficiaryAdviceLineFourAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFour = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineFiveAlphaNumeric validates FIBeneficiaryAdvice LineFive is alphanumeric
func TestFIBeneficiaryAdviceLineFiveAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineFive = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryAdviceLineSixAlphaNumeric validates FIBeneficiaryAdvice LineSix is alphanumeric
func TestFIBeneficiaryAdviceLineSixAlphaNumeric(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.Advice.LineSix = "®"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIBeneficiaryAdviceWrongLength parses a wrong FIBeneficiaryAdvice record length
func TestParseFIBeneficiaryAdviceWrongLength(t *testing.T) {
	var line = "{6410}LTRLine One                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	err := r.parseFIBeneficiaryAdvice()
	if err != nil {
		if !base.Match(err, NewTagWrongLengthErr(200, len(r.line))) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestParseFIBeneficiaryAdviceReaderParseError parses a wrong FIBeneficiaryAdvice reader parse error
func TestParseFIBeneficiaryAdviceReaderParseError(t *testing.T) {
	var line = "{6410}LTRLine ®ne                  Line Two                         Line Three                       Line Four                        Line Five                        Line Six                         "
	r := NewReader(strings.NewReader(line))
	r.line = line
	fwm := new(FEDWireMessage)
	fiba := mockFIBeneficiaryAdvice()
	fwm.SetFIBeneficiaryAdvice(fiba)
	err := r.parseFIBeneficiaryAdvice()
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

// TestFIBeneficiaryAdviceTagError validates a FIBeneficiaryAdvice tag
func TestFIBeneficiaryAdviceTagError(t *testing.T) {
	fiba := mockFIBeneficiaryAdvice()
	fiba.tag = "{9999}"
	if err := fiba.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
