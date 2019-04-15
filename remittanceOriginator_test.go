package wire

// RemittanceOriginator creates a RemittanceOriginator
func mockRemittanceOriginator() *RemittanceOriginator {
	ro := NewRemittanceOriginator()
	ro.IdentificationType = OrganizationID
	ro.IdentificationCode = OICCustomerNumber
	ro.IdentificationNumber = "111111"
	ro.IdentificationNumberIssuer = "Bank"
	ro.RemittanceData.DateBirthPlace = "12072008 AnyTown"
	ro.RemittanceData.Name = "Name"
	ro.RemittanceData.AddressType = CompletePostalAddress
	ro.RemittanceData.Department = "Department"
	ro.RemittanceData.SubDepartment = "Sub-Department"
	ro.RemittanceData.BuildingNumber = "16"
	ro.RemittanceData.PostCode = "19405"
	ro.RemittanceData.TownName = "AnyTown"
	ro.RemittanceData.CountrySubDivisionState = "PA"
	ro.RemittanceData.Country = "UA"
	ro.RemittanceData.AddressLineOne = "Address Line One"
	ro.RemittanceData.AddressLineTwo = "Address Line Two"
	ro.RemittanceData.AddressLineThree = "Address Line Three"
	ro.RemittanceData.AddressLineFour = "Address Line Four"
	ro.RemittanceData.AddressLineFive = "Address Line Five"
	ro.RemittanceData.AddressLineSix = "Address Line Six"
	ro.RemittanceData.AddressLineSeven = "Address Line Seven"
	ro.CountryOfResidence = "US"
	ro.ContactName = "Contact Name"
	ro.ContactPhoneNumber = "5551231212"
	ro.ContactMobileNumber = "5551231212"
	ro.ContactFaxNumber = "5551231212"
	ro.ContactElectronicAddress = "http://www.moov.io"
	ro.ContactOther = "Contact Other"
	return ro
}
