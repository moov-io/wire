package wire

//  mockInstructingFI creates a InstructingFI
func mockInstructingFI() *InstructingFI {
	ifi := NewInstructingFI()
	ifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	ifi.FinancialInstitution.Identifier = "123456789"
	ifi.FinancialInstitution.Name = "FI Name"
	ifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ifi
}