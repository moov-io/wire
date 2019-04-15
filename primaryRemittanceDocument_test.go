package wire

// PrimaryRemittanceDocument creates a PrimaryRemittanceDocument
func mockPrimaryRemittanceDocument() *PrimaryRemittanceDocument {
	prd := NewPrimaryRemittanceDocument()
	prd.DocumentTypeCode = AccountsReceivableOpenItem
	prd.ProprietaryDocumentTypeCode = ""
	prd.DocumentIdentificationNumber = "111111"
	prd.Issuer = "Issuer"
	return prd
}
