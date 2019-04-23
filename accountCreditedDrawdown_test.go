package wire

// mockAccountCreditedDrawdown creates a AccountCreditedDrawdown
func mockAccountCreditedDrawdown() *AccountCreditedDrawdown {
	creditDD := NewAccountCreditedDrawdown()
	creditDD.DrawdownCreditAccountNumber = "123456789"
	return creditDD
}
