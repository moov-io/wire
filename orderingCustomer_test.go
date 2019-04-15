package wire

//  OrderingCustomer creates a OrderingCustomer
func mockOrderingCustomer() *OrderingCustomer {
	oc := NewOrderingCustomer()
	oc.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	oc.CoverPayment.SwiftLineOne = "Swift Line One"
	oc.CoverPayment.SwiftLineTwo = "Swift Line Two"
	oc.CoverPayment.SwiftLineThree = "Swift Line Three"
	oc.CoverPayment.SwiftLineFour = "Swift Line Four"
	oc.CoverPayment.SwiftLineFive = "Swift Line Five"
	return oc
}
