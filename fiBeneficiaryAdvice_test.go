package wire

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
