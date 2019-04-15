package wire

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