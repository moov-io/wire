package wire

//  mockInputMessageAccountabilityData creates a mockInputMessageAccountabilityData
func mockInputMessageAccountabilityData() *InputMessageAccountabilityData {
	imad := NewInputMessageAccountabilityData()
	imad.InputCycleDate = "20190410"
	imad.InputSource = "Source08"
	imad.InputSequenceNumber = "000001"
	return imad
}