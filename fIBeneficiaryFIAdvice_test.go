package wire

import (
	"github.com/moov-io/base"
	"testing"
)

// mockFIBeneficiaryFIAdvice creates a FIBeneficiaryFIAdvice
func mockFIBeneficiaryFIAdvice() *FIBeneficiaryFIAdvice {
	fibfia := NewFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = AdviceCodeTelex
	fibfia.Advice.LineOne = "Line One"
	fibfia.Advice.LineTwo = "Line Two"
	fibfia.Advice.LineThree = "Line Three"
	fibfia.Advice.LineFour = "Line Four"
	fibfia.Advice.LineFive = "Line Five"
	fibfia.Advice.LineSix = "Line Six"
	return fibfia
}

// TestMockFIBeneficiaryFIAdvice validates mockFIBeneficiaryFIAdvice
func TestMockFIBeneficiaryFIAdvice(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	if err := fibfia.Validate(); err != nil {
		t.Error("mockFIBeneficiaryFIAdvice does not validate and will break other tests")
	}
}

// TestFIBeneficiaryFIAdviceAdviceCodeValid validates FIBeneficiaryFIAdvice AdviceCode is alphanumeric
func TestFIBeneficiaryFIAdviceAdviceCodeValid(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.AdviceCode = "Z"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrAdviceCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineOneAlphaNumeric validates FIBeneficiaryFIAdvice LineOne is alphanumeric
func TestFIBeneficiaryFIAdviceLineOneAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineOne = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineTwoAlphaNumeric validates FIBeneficiaryFIAdvice LineTwo is alphanumeric
func TestFIBeneficiaryFIAdviceLineTwoAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineTwo = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineThreeAlphaNumeric validates FIBeneficiaryFIAdvice LineThree is alphanumeric
func TestFIBeneficiaryFIAdviceLineThreeAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineThree = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineFourAlphaNumeric validates FIBeneficiaryFIAdvice LineFour is alphanumeric
func TestFIBeneficiaryFIAdviceLineFourAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineFour = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineFiveAlphaNumeric validates FIBeneficiaryFIAdvice LineFive is alphanumeric
func TestFIBeneficiaryFIAdviceLineFiveAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineFive = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFIBeneficiaryFIAdviceLineSixAlphaNumeric validates FIBeneficiaryFIAdvice LineSix is alphanumeric
func TestFIBeneficiaryFIAdviceLineSixAlphaNumeric(t *testing.T) {
	fibfia := mockFIBeneficiaryFIAdvice()
	fibfia.Advice.LineSix = "®"
	if err := fibfia.Validate(); err != nil {
		if !base.Match(err, ErrNonAlphanumeric) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
