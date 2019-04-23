package wire

// mockFIBeneficiaryFI creates a FIBeneficiaryFI
func mockFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := NewFIBeneficiaryFI()
	fibfi.FIToFI.LineOne = "Line One"
	fibfi.FIToFI.LineTwo = "Line Two"
	fibfi.FIToFI.LineThree = "Line Three"
	fibfi.FIToFI.LineFour = "Line Four"
	fibfi.FIToFI.LineFive = "Line Five"
	fibfi.FIToFI.LineSix = "Line Six"
	return fibfi
}
