package wire

//  mockAccountDebitedDrawdown creates a AccountDebitedDrawdown
func mockAccountDebitedDrawdown() *AccountDebitedDrawdown {
	debitDD := NewAccountDebitedDrawdown()
	debitDD.IdentificationCode = "DemandDepositAccountNumber"
	debitDD.Identifier = "123456789"
	debitDD.Name = "debitDD Name"
	debitDD.Address.AddressLineOne = "Address One"
	debitDD.Address.AddressLineTwo = "Address Two"
	debitDD.Address.AddressLineThree = "Address Three"
	return debitDD
}