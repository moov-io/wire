package wire

//  mockSenderDepositoryInstitution creates a SenderDepositoryInstitution
func mockSenderDepositoryInstitution() *SenderDepositoryInstitution {
	sdi := NewSenderDepositoryInstitution()
	sdi.SenderABANumber = "121042882"
	sdi.SenderShortName = "Wells Fargo NA"
	return sdi
}
