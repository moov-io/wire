package wire

// Remittance creates a Remittance
func mockRemittance() *Remittance {
	ri := NewRemittance()
	ri.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ri.CoverPayment.SwiftLineOne = "Swift Line One"
	ri.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ri.CoverPayment.SwiftLineThree = "Swift Line Three"
	ri.CoverPayment.SwiftLineFour = "Swift Line Four"
	return ri
}
