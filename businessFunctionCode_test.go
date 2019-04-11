package wire

//  mockBusinessFunctionCode creates a BusinessFunctionCode
func mockBusinessFunctionCode() *BusinessFunctionCode {
	bfc := NewBusinessFunctionCode()
	bfc.BusinessFunctionCode = BankTransfer
	bfc.TransactionTypeCode = "XYZ"
	return bfc
}
