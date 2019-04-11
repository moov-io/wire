package wire

//  mockBeneficiary creates a Beneficiary
func mockBeneficiary() *Beneficiary {
	ben := NewBeneficiary()
	ben.Personal.IdentificationCode = DriversLicenseNumber
	ben.Personal.Identifier = "1234"
	ben.Personal.Name = "Name"
	ben.Personal.Address.AddressLineOne = "Address One"
	ben.Personal.Address.AddressLineTwo = "Address Two"
	ben.Personal.Address.AddressLineThree = "Address Three"
	return ben
}
