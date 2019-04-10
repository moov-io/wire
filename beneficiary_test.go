package wire

//  mockBeneficiary creates a Beneficiary
func mockBeneficiary() *Beneficiary {
	b := NewBeneficiary()
	b.Personal.IdentificationCode = DriversLicenseNumber
	b.Personal.Identifier = "1234"
	b.Personal.Name = "Name"
	b.Personal.Address.AddressLineOne = "Address One"
	b.Personal.Address.AddressLineTwo = "Address Two"
	b.Personal.Address.AddressLineThree = "Address Three"
	return b
}