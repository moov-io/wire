package wire

//  CurrencyInstructedAmount creates a CurrencyInstructedAmount
func mockCurrencyInstructedAmount() *CurrencyInstructedAmount {
	cia := NewCurrencyInstructedAmount()
	cia.SwiftFieldTag = "Swift Field Tag"
	cia.Amount = "1500,49"

	return cia
}
