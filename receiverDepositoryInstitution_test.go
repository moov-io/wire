package wire

// mockReceiverDepositoryInstitution creates a ReceiverDepositoryInstitution
func mockReceiverDepositoryInstitution() *ReceiverDepositoryInstitution {
	rdi := NewReceiverDepositoryInstitution()
	rdi.ReceiverABANumber = "231380104"
	rdi.ReceiverShortName = "Citadel"
	return rdi
}
