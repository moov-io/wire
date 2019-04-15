package wire

// RelatedRemittance creates a RelatedRemittance
func mockRelatedRemittance() *RelatedRemittance {
	rr := NewRelatedRemittance()
	rr.RemittanceIdentification = "Remittance Identification"
	rr.RemittanceLocationMethod = RLMElectronicDataExchange
	rr.RemittanceLocationElectronicAddress = "http://moov.io"
	rr.RemittanceData.Name = "Name"
	rr.RemittanceData.AddressType = CompletePostalAddress
	rr.RemittanceData.Department = "Department"
	rr.RemittanceData.SubDepartment = "Sub-Department"
	rr.RemittanceData.BuildingNumber = "16"
	rr.RemittanceData.PostCode = "19405"
	rr.RemittanceData.TownName = "AnyTown"
	rr.RemittanceData.CountrySubDivisionState = "PA"
	rr.RemittanceData.Country = "UA"
	rr.RemittanceData.AddressLineOne = "Address Line One"
	rr.RemittanceData.AddressLineTwo = "Address Line Two"
	rr.RemittanceData.AddressLineThree = "Address Line Three"
	rr.RemittanceData.AddressLineFour = "Address Line Four"
	rr.RemittanceData.AddressLineFive = "Address Line Five"
	rr.RemittanceData.AddressLineSix = "Address Line Six"
	rr.RemittanceData.AddressLineSeven = "Address Line Seven"
	return rr
}
