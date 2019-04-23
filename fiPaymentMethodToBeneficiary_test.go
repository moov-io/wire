package wire

// mockFIPaymentMethodToBeneficiary creates a FIPaymentMethodToBeneficiary
func mockFIPaymentMethodToBeneficiary() *FIPaymentMethodToBeneficiary {
	pm := NewFIPaymentMethodToBeneficiary()
	pm.PaymentMethod = "CHECK"
	pm.AdditionalInformation = "Additional Information"
	return pm
}
