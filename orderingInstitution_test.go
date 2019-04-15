package wire

//  OrderingInstitution creates a OrderingInstitution
func mockOrderingInstitution() *OrderingInstitution {
	oi := NewOrderingInstitution()
	oi.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oi.CoverPayment.SwiftLineOne = "Swift Line One"
	oi.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oi.CoverPayment.SwiftLineThree = "Swift Line Three"
	oi.CoverPayment.SwiftLineFour = "Swift Line Four"
	oi.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oi
}
