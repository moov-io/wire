package wire

import (
	"errors"
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

	require.EqualError(t, err, r.parseError(fieldError("LineTwelve", ErrValidLength)).Error())
}

// TestParseServiceMessageReaderParseError parses a wrong ServiceMessage reader parse error
func TestParseServiceMessageReaderParseError(t *testing.T) {
	var line = "{9000}®ine One                           Line Two                           Line Three                         Line Four                          Line Five                          Line Six                           Line Seven                         Line Eight                         Line Nine                          Line Ten                           Line Eleven                        line Twelve                       "
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
	fwm.ServiceMessage = sm
	bfc := mockBusinessFunctionCode()
	bfc.TransactionTypeCode = "COV"
	fwm.BusinessFunctionCode = bfc

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("BusinessFunctionCode.TransactionTypeCode", ErrTransactionTypeCode, bfc.TransactionTypeCode).Error())
}

// TestInvalidLocalInstrumentForServiceMessage test an invalid LocalInstrument
func TestInvalidLocalInstrumentForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	li := mockLocalInstrument()
	fwm.LocalInstrument = li

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("LocalInstrument", ErrInvalidProperty, li).Error())
}

// TestInvalidPaymentNotificationForServiceMessage test an invalid PaymentNotification
func TestInvalidPaymentNotificationForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	pn := mockPaymentNotification()
	fwm.PaymentNotification = pn

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("PaymentNotification", ErrInvalidProperty, pn).Error())
}

// TestInvalidChargesForServiceMessage test an invalid Charges
func TestInvalidChargesForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	c := mockCharges()
	fwm.Charges = c

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Charges", ErrInvalidProperty, c).Error())
}

// TestInvalidInstructedAmountForServiceMessage test an invalid InstructedAmount
func TestInvalidInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	ia := mockInstructedAmount()
	fwm.InstructedAmount = ia

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("InstructedAmount", ErrInvalidProperty, ia).Error())
}

// TestInvalidExchangeRateForServiceMessage test an invalid ExchangeRate
func TestInvalidExchangeRateForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	eRate := mockExchangeRate()
	fwm.ExchangeRate = eRate

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("ExchangeRate", ErrInvalidProperty, fwm.ExchangeRate).Error())
}

// TestInvalidBeneficiaryIdentificationCodeForServiceMessage test an invalid BeneficiaryIdentificationCode
func TestInvalidBeneficiaryIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	ben := mockBeneficiary()
	ben.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.Beneficiary = ben

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Beneficiary.Personal.IdentificationCode", ErrInvalidProperty, ben.Personal.IdentificationCode).Error())
}

// TestInvalidOriginatorIdentificationCodeForServiceMessage test an invalid OriginatorIdentificationCode
func TestInvalidOriginatorIdentificationCodeForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	o := mockOriginator()
	o.Personal.IdentificationCode = SWIFTBICORBEIANDAccountNumber
	fwm.Originator = o

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("Originator.Personal.IdentificationCode", ErrInvalidProperty, o.Personal.IdentificationCode).Error())
}

// TestInvalidOriginatorOptionFForServiceMessage test an invalid OriginatorOptionF
func TestInvalidOriginatorOptionFForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	off := mockOriginatorOptionF()
	fwm.OriginatorOptionF = off

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("OriginatorOptionF", ErrInvalidProperty, fwm.OriginatorOptionF).Error())
}

// TestInvalidUnstructuredAddendaForServiceMessage test an invalid UnstructuredAddenda
func TestInvalidUnstructuredAddendaForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	ua := mockUnstructuredAddenda()
	fwm.UnstructuredAddenda = ua

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("BusinessFunctionCode", ErrInvalidProperty, "Unstructured Addenda").Error())
}

// TestInvalidCurrencyInstructedAmountForServiceMessage test an invalid CurrencyInstructedAmount
func TestInvalidCurrencyInstructedAmountForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	cia := mockCurrencyInstructedAmount()
	fwm.CurrencyInstructedAmount = cia

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("CurrencyInstructedAmount", ErrInvalidProperty, fwm.CurrencyInstructedAmount).Error())
}

// TestInvalidRelatedRemittanceForServiceMessage test an invalid RelatedRemittance
func TestInvalidRelatedRemittanceForServiceMessage(t *testing.T) {
	fwm := new(FEDWireMessage)
	sm := mockServiceMessage()
	fwm.ServiceMessage = sm
	rr := mockRelatedRemittance()
	fwm.RelatedRemittance = rr

	err := fwm.checkProhibitedServiceMessageTags()

	require.EqualError(t, err, fieldError("RelatedRemittance", ErrInvalidProperty, fwm.RelatedRemittance).Error())
}

// TestServiceMessageTagError validates a ServiceMessage tag
func TestServiceMessageTagError(t *testing.T) {
	sm := mockServiceMessage()
	sm.tag = "{9999}"

	require.EqualError(t, sm.Validate(), fieldError("tag", ErrValidTagForType, sm.tag).Error())
}

// TestStringServiceMessageVariableLength parses using variable length
func TestStringServiceMessageVariableLength(t *testing.T) {
	var line = "{9000}A*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseServiceMessage()
	require.Nil(t, err)

	line = "{9000}A                                                                                                                                                                                                                                                                                                                                                                                                                                   NNN"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseServiceMessage()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{9000}**************"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseServiceMessage()
	require.ErrorContains(t, err, r.parseError(NewTagMaxLengthErr(errors.New(""))).Error())

	line = "{9000}A*"
	r = NewReader(strings.NewReader(line))
	r.line = line

	err = r.parseServiceMessage()
	require.Equal(t, err, nil)
}

// TestStringServiceMessageOptions validates Format() formatted according to the FormatOptions
func TestStringServiceMessageOptions(t *testing.T) {
	var line = "{9000}A*"
	r := NewReader(strings.NewReader(line))
	r.line = line

	err := r.parseServiceMessage()
	require.Equal(t, err, nil)

	record := r.currentFEDWireMessage.ServiceMessage
	require.Equal(t, record.String(), "{9000}A                                                                                                                                                                                                                                                                                                                                                                                                                                   ")
	require.Equal(t, record.Format(FormatOptions{VariableLengthFields: true}), "{9000}A*")
	require.Equal(t, record.String(), record.Format(FormatOptions{VariableLengthFields: false}))
}
