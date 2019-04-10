package wire

//  mockOriginatorToBeneficiary creates a OriginatorToBeneficiary
func mockOriginatorToBeneficiary() *OriginatorToBeneficiary {
	ob := NewOriginatorToBeneficiary()
	ob.LineOne = "LineOne"
	ob.LineTwo = "LineTwo"
	ob.LineThree = "LineThree"
	ob.LineFour = "LineFour"
	return ob
}