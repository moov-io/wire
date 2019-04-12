package wire

//  mockFIAdditionalFIToFI creates a FIAdditionalFIToFI
func mockFIAdditionalFIToFI() *FIAdditionalFIToFI {
	fifi := NewFIAdditionalFIToFI()
	fifi.AdditionalFIToFI.LineOne = "Line One"
	fifi.AdditionalFIToFI.LineTwo = "Line Two"
	fifi.AdditionalFIToFI.LineThree = "Line Three"
	fifi.AdditionalFIToFI.LineFour = "Line Four"
	fifi.AdditionalFIToFI.LineFive = "Line Five"
	fifi.AdditionalFIToFI.LineSix = "Line Six"
	return fifi
}
