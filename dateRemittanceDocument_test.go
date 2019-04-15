package wire

// DateRemittanceDocument creates a DateRemittanceDocument
func mockDateRemittanceDocument() *DateRemittanceDocument {
	drd := NewDateRemittanceDocument()
	// ToDo: Use date function
	drd.DateRemittanceDocument = "20190415"
	return drd
}
