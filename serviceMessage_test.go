package wire

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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

	require.NoError(t, sm.Validate(), "mockServiceMessage does not validate and will break other tests")
}

// TestLineOneAlphaNumeric validates ServiceMessage LineOne is alphanumeric
func TestLineOneAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineOne = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrNonAlphanumeric, sm.LineOne).Error())
}

// TestLineTwoAlphaNumeric validates ServiceMessage LineTwo is alphanumeric
func TestLineTwoAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTwo = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineTwo", ErrNonAlphanumeric, sm.LineTwo).Error())
}

// TestLineThreeAlphaNumeric validates ServiceMessage LineThree is alphanumeric
func TestLineThreeAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineThree = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineThree", ErrNonAlphanumeric, sm.LineThree).Error())
}

// TestLineFourAlphaNumeric validates ServiceMessage LineFour is alphanumeric
func TestLineFourAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineFour = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineFour", ErrNonAlphanumeric, sm.LineFour).Error())
}

// TestLineFiveAlphaNumeric validates ServiceMessage LineFive is alphanumeric
func TestLineFiveAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineFive = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineFive", ErrNonAlphanumeric, sm.LineFive).Error())
}

// TestLineSixAlphaNumeric validates ServiceMessage LineSix is alphanumeric
func TestLineSixAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineSix = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineSix", ErrNonAlphanumeric, sm.LineSix).Error())
}

// TestLineSevenAlphaNumeric validates ServiceMessage LineSeven is alphanumeric
func TestLineSevenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineSeven = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineSeven", ErrNonAlphanumeric, sm.LineSeven).Error())
}

// TestLineEightAlphaNumeric validates ServiceMessage LineEight is alphanumeric
func TestLineEightAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineEight = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineEight", ErrNonAlphanumeric, sm.LineEight).Error())
}

// TestLineNineAlphaNumeric validates ServiceMessage LineNine is alphanumeric
func TestLineNineAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineNine = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineNine", ErrNonAlphanumeric, sm.LineNine).Error())
}

// TestLineTenAlphaNumeric validates ServiceMessage LineTen is alphanumeric
func TestLineTenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTen = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineTen", ErrNonAlphanumeric, sm.LineTen).Error())
}

// TestLineElevenAlphaNumeric validates ServiceMessage LineEleven is alphanumeric
func TestLineElevenAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineEleven = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineEleven", ErrNonAlphanumeric, sm.LineEleven).Error())
}

// TestLineTwelveAlphaNumeric validates ServiceMessage LineTwelve is alphanumeric
func TestLineTwelveAlphaNumeric(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineTwelve = "®"

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineTwelve", ErrNonAlphanumeric, sm.LineTwelve).Error())
}

// TestLineOneRequired validates ServiceMessage LineOne is required
func TestLineOneRequired(t *testing.T) {
	sm := mockServiceMessage()
	sm.LineOne = ""

	err := sm.Validate()

	require.EqualError(t, err, fieldError("LineOne", ErrFieldRequired).Error())
}

// TestParseServiceMessageWrongLength parses a wrong ServiceMessage record length
func TestParseServiceMessageWrongLength(t *testing.T) {
	var line = "{9000}Line One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           Line Seven                         Line Eight                         Line Nine                          Line Ten                           Line Eleven                        line Twelve                      "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseServiceMessage()

	require.EqualError(t, err, r.parseError(NewTagWrongLengthErr(426, len(r.line))).Error())
}

// TestParseServiceMessageReaderParseError parses a wrong ServiceMessage reader parse error
func TestParseServiceMessageReaderParseError(t *testing.T) {
	var line = "{9000}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           Line Seven                         Line Eight                         Line Nine                          Line Ten                           Line Eleven                        line Twelve                        "
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseServiceMessage()

	require.EqualError(t, err, r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error())

	_, err = r.Read()

	require.EqualError(t, err, r.parseError(fieldError("LineOne", ErrNonAlphanumeric, "®ine One")).Error())
}

// TestTransactionTypeCodeForServiceMessage test an invalid TransactionTypeCode
func TestInvalidTransactionTypeCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	bfc := mockBusinessFunctionCode()
	bfc.TransactionTypeCode = "COV"
	fwm.SetBusinessFunctionCode(bfc)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, bfc.TransactionTypeCode).Error())
}

// TestInvalidLocalInstrumentForServiceMessage test an invalid LocalInstrument
func TestInvalidLocalInstrumentForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	li := mockLocalInstrument()
	fwm.SetLocalInstrument(li)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("LocalInstrument", ErrInvalidProperty, li).Error())
}

// TestInvalidPaymentNotificationForServiceMessage test an invalid PaymentNotification
func TestInvalidPaymentNotificationForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	pn := mockPaymentNotification()
	fwm.SetPaymentNotification(pn)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("PaymentNotification", ErrInvalidProperty, pn).Error())
}

// TestInvalidChargesForServiceMessage test an invalid Charges
func TestInvalidChargesForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	c := mockCharges()
	fwm.SetCharges(c)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Charges", ErrInvalidProperty, c).Error())
}

// TestInvalidInstructedAmountForServiceMessage test an invalid InstructedAmount
func TestInvalidInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ia := mockInstructedAmount()
	fwm.SetInstructedAmount(ia)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("InstructedAmount", ErrInvalidProperty, ia).Error())
}

// TestInvalidExchangeRateForServiceMessage test an invalid ExchangeRate
func TestInvalidExchangeRateForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	eRate := mockExchangeRate()
	fwm.SetExchangeRate(eRate)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate).Error())
}

// TestInvalidBeneficiaryIdentificationCodeForServiceMessage test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetBeneficiary(ben)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, ben.Personal.IdentificationCode).Error())
}

// TestInvalidOriginatorIdentificationCodeForServiceMessage test an invalid OriginatorIdentificationCode
func TestInvalidOriginatorIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.SetOriginator(o)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, o.Personal.IdentificationCode).Error())
}

// TestInvalidOriginatorOptionFForServiceMessage test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	off := mockOriginatorOptionF()
	fwm.SetOriginatorOptionF(off)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error())
}

// TestInvalidUnstructuredAddendaForServiceMessage test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	ua := mockUnstructuredAddenda()
	fwm.SetUnstructuredAddenda(ua)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda").Error())
}

// TestInvalidCurrencyInstructedAmountForServiceMessage test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	cia := mockCurrencyInstructedAmount()
	fwm.SetCurrencyInstructedAmount(cia)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error())
}

// TestInvalidRelatedRemittanceForServiceMessage test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.SetServiceMessage(sm)
	rr := mockRelatedRemittance()
	fwm.SetRelatedRemittance(rr)

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error())
}

// TestServiceMessageTagError validates a ServiceMessage tag
func TestServiceMessageTagError(t *testing.T) {
	sm := mockServiceMessage()
	sm.tag = "{9999}"

	require.EqualError(t, sm.Validate(), fieldError("tag", ErrValidTagForType, sm.tag).Error())
}
