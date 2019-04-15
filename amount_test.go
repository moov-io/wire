package wire

// mockAmount creates an a Amount
func mockAmount() *Amount {
	a := NewAmount()
	a.Amount = "000001234567"
	return a
}