package wire

// mockTypeSubType creates a TypeSubType
func mockTypeSubType() *TypeSubType {
	tst := NewTypeSubType()
	tst.tag = TagTypeSubType
	tst.TypeCode = FundsTransfer
	tst.SubTypeCode = BasicFundsTransfer
	return tst
}