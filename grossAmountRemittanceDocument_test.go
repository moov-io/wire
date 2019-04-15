package wire

// GrossAmountRemittanceDocument creates a GrossAmountRemittanceDocument
func mockGrossAmountRemittanceDocument() *GrossAmountRemittanceDocument {
	gard := NewGrossAmountRemittanceDocument()
	gard.RemittanceAmount.CurrencyCode = "USD"
	gard.RemittanceAmount.Amount = "1234.56"
	return gard
}
