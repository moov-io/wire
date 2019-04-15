package wire

//  IntermediaryInstitution creates a IntermediaryInstitution
func mockIntermediaryInstitution() *IntermediaryInstitution {
	ii := NewIntermediaryInstitution()
	ii.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ii.CoverPayment.SwiftLineOne = "Swift Line One"
	ii.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ii.CoverPayment.SwiftLineThree = "Swift Line Three"
	ii.CoverPayment.SwiftLineFour = "Swift Line Four"
	ii.CoverPayment.SwiftLineFive = "Swift Line Five"
	return ii
}
