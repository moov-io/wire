package wire

//  mockExchangeRate creates a ExchangeRate
func mockExchangeRate() *ExchangeRate {
	eRate := NewExchangeRate()
	eRate.ExchangeRate = "1,2345"
	return eRate
}
