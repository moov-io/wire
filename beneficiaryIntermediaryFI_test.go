package wire

// mockBeneficiaryIntermediaryFI creates a BeneficiaryIntermediaryFI
func mockBeneficiaryIntermediaryFI() *BeneficiaryIntermediaryFI {
	bifi := NewBeneficiaryIntermediaryFI()
	bifi.FinancialInstitution.IdentificationCode = DemandDepositAccountNumber
	bifi.FinancialInstitution.Identifier = "123456789"
	bifi.FinancialInstitution.Name = "FI Name"
	bifi.FinancialInstitution.Address.AddressLineOne = "Address One"
	bifi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	bifi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return bifi
}
