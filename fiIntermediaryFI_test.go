package wire

// mockFIIntermediaryFI creates a FIIntermediaryFI
func mockFIIntermediaryFI() *FIIntermediaryFI {
	fiifi := NewFIIntermediaryFI()
	fiifi.FIToFI.LineOne = "Line One"
	fiifi.FIToFI.LineOne = "Line Two"
	fiifi.FIToFI.LineOne = "Line Three"
	fiifi.FIToFI.LineOne = "Line Four"
	fiifi.FIToFI.LineOne = "Line Five"
	fiifi.FIToFI.LineOne = "Line Six"
	return fiifi
}
