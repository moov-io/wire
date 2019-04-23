package wire

// mockOriginator creates a Originator
func mockOriginator() *Originator {
	o := NewOriginator()
	o.Personal.IdentificationCode = PassportNumber
	o.Personal.Identifier = "1234"
	o.Personal.Name = "Name"
	o.Personal.Address.AddressLineOne = "Address One"
	o.Personal.Address.AddressLineTwo = "Address Two"
	o.Personal.Address.AddressLineThree = "Address Three"
	return o
}
