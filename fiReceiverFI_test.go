package wire

// mockFIReceiverFI creates a FIReceiverFI
func mockFIReceiverFI() *FIReceiverFI {
	firfi := NewFIReceiverFI()
	firfi.FIToFI.LineOne = "Line One"
	firfi.FIToFI.LineOne = "Line Two"
	firfi.FIToFI.LineOne = "Line Three"
	firfi.FIToFI.LineOne = "Line Four"
	firfi.FIToFI.LineOne = "Line Five"
	firfi.FIToFI.LineOne = "Line Six"
	return firfi
}
