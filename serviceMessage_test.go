package wire

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
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

// TestTransactionTypeCodeForServiceMessage test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	bfc := mockBusinessFunctionCode()
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrTransactionTypeCode) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidLocalInstrumentForServiceMessage test an invalid LocalInstrument
func TestInvalidLocalInstrumentForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidPaymentNotificationForServiceMessage test an invalid PaymentNotification
func TestInvalidPaymentNotificationForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidChargesForServiceMessage test an invalid Charges
func TestInvalidChargesForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	c := mockCharges()
	fwm.SetCharges(c)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidInstructedAmountForServiceMessage test an invalid InstructedAmount
func TestInvalidInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidExchangeRateForServiceMessage test an invalid ExchangeRate
func TestInvalidExchangeRateForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidBeneficiaryIdentificationCodeForServiceMessage test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetBeneficiary(ben)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorIdentificationCodeForServiceMessage test an invalid OriginatorIdentificationCode
func TestInvalidOriginatorIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetOriginator(o)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidOriginatorOptionFForServiceMessage test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	off := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(off)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidUnstructuredAddendaForServiceMessage test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidCurrencyInstructedAmountForServiceMessage test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestInvalidRelatedRemittanceForServiceMessage test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)
	if err := fwm.checkProhibitedServiceMessageTags(); err != nil {
		if !base.Match(err, ErrInvalidProperty) {
			t.Errorf("%T: %s", err, err)
		}

	}
}

// TestServiceMessageTagError validates a ServiceMessage tag
func TestServiceMessageTagError(t *testing.T) {
	sm := mockServiceMessage()
	sm.tag = "{9999}"
	if err := sm.Validate(); err != nil {
		if !base.Match(err, ErrValidTagForType) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
