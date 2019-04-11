package wire

//  mockInstructedAmount creates a InstructedAmount
func mockInstructedAmount() *InstructedAmount {
	ia := NewInstructedAmount()
	ia.CurrencyCode = "USD"
	ia.Amount = "4567,89"
	return ia
}