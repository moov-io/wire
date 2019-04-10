package wire

//  mockOriginatorFI creates a OriginatorFI
func mockOriginatorFI() *OriginatorFI {
	ofi := NewOriginatorFI()
	ofi.FinancialInstitution.IdentificationCode = PassportNumber
	ofi.FinancialInstitution.Identifier = "123456789"
	ofi.FinancialInstitution.Name = "FI Name"
	ofi.FinancialInstitution.Address.AddressLineOne = "Address One"
	ofi.FinancialInstitution.Address.AddressLineTwo = "Address Two"
	ofi.FinancialInstitution.Address.AddressLineThree = "Address Three"
	return ofi
}