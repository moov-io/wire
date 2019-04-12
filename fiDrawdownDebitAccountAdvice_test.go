package wire

//  mockFIDrawdownDebitAccountAdvice creates a FIDrawdownDebitAccountAdvice
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