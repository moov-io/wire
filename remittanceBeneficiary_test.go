package wire

// RemittanceBeneficiary creates a RemittanceBeneficiary
func mockRemittanceBeneficiary() *RemittanceBeneficiary {
	rb := NewRemittanceBeneficiary()
	rb.RemittanceData.Name = "Name"
	rb.IdentificationType = OrganizationID
	rb.IdentificationCode = OICCustomerNumber
	rb.IdentificationNumber = "111111"
	rb.IdentificationNumberIssuer = "Bank"
	rb.RemittanceData.DateBirthPlace = "03062013 AnyTown"
	rb.RemittanceData.AddressType = CompletePostalAddress
	rb.RemittanceData.Department = "Department"
	rb.RemittanceData.SubDepartment = "Sub-Department"
	rb.RemittanceData.BuildingNumber = "16"
	rb.RemittanceData.PostCode = "19405"
	rb.RemittanceData.TownName = "AnyTown"
	rb.RemittanceData.CountrySubDivisionState = "PA"
	rb.RemittanceData.Country = "UA"
	rb.RemittanceData.AddressLineOne = "Address Line One"
	rb.RemittanceData.AddressLineTwo = "Address Line Two"
	rb.RemittanceData.AddressLineThree = "Address Line Three"
	rb.RemittanceData.AddressLineFour = "Address Line Four"
	rb.RemittanceData.AddressLineFive = "Address Line Five"
	rb.RemittanceData.AddressLineSix = "Address Line Six"
	rb.RemittanceData.AddressLineSeven = "Address Line Seven"
	rb.CountryOfResidence = "US"
	return rb
}
