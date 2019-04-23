package wire

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
