package wire

// SecondaryRemittanceDocument creates a SecondaryRemittanceDocument
func mockSecondaryRemittanceDocument() *SecondaryRemittanceDocument {
	srd := NewSecondaryRemittanceDocument()
	srd.DocumentTypeCode = StatementAccount
	srd.ProprietaryDocumentTypeCode = ""
	srd.DocumentIdentificationNumber = "222222"
	srd.Issuer = "Issuer 2"
	return srd
}
