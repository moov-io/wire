package wire

// mockBeneficiaryFI creates a BeneficiaryFI
func mockBeneficiaryFI() *BeneficiaryFI {
	bfi := NewBeneficiaryFI()
	bfi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bfi.FinancialInstitution.Identifier = "123456789"
	bfi.FinancialInstitution.Name = "FI Name"
	bfi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bfi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bfi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bfi
}
