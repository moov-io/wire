package wire

// RemittanceFreeText creates a RemittanceFreeText
func mockRemittanceFreeText() *RemittanceFreeText {
	rft := NewRemittanceFreeText()
	rft.LineOne = "Remittance Free Text Line One"
	rft.LineTwo = "Remittance Free Text Line Two"
	rft.LineThree = "Remittance Free Text Line Three"
	return rft
}
