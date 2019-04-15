package wire

// AmountNegotiatedDiscount creates a AmountNegotiatedDiscount
func mockAmountNegotiatedDiscount() *AmountNegotiatedDiscount {
	nd := NewAmountNegotiatedDiscount()
	nd.RemittanceAmount.CurrencyCode = "USD"
	nd.RemittanceAmount.Amount = "1234.56"
	return nd
}
