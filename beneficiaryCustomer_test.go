package wire

//  mockBeneficiaryCustomer creates a BeneficiaryCustomer
func mockBeneficiaryCustomer() *BeneficiaryCustomer {
	bc := NewBeneficiaryCustomer()
	bc.CoverPayment.SwiftFieldTag = ""
	bc.CoverPayment.SwiftLineOne = ""
	bc.CoverPayment.SwiftLineTwo = ""
	bc.CoverPayment.SwiftLineThree= ""
	bc.CoverPayment.SwiftLineFour = ""
	bc.CoverPayment.SwiftLineFive = ""
	bc.CoverPayment.SwiftLineSix = ""
	return bc
}