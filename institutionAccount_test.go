package wire

//  InstitutionAccount creates a InstitutionAccount
func mockInstitutionAccount() *InstitutionAccount {
	ia := NewInstitutionAccount()
	ia.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	ia.CoverPayment.SwiftLineOne = "Swift Line One"
	ia.CoverPayment.SwiftLineTwo = "Swift Line Two"
	ia.CoverPayment.SwiftLineThree = "Swift Line Three"
	ia.CoverPayment.SwiftLineFour = "Swift Line Four"
	ia.CoverPayment.SwiftLineFive = "Swift Line Five"
	return ia
}
