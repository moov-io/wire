package wire

// mockPaymentNotification creates a PaymentNotification
func mockPaymentNotification() *PaymentNotification {
	pn := NewPaymentNotification()
	pn.PaymentNotificationIndicator = "1"
	pn.ContactNotificationElectronicAddress = "http://moov.io"
	pn.ContactName = "Contact Name"
	pn.ContactPhoneNumber = "5555551212"
	pn.ContactMobileNumber = "5551231212"
	pn.ContactFaxNumber = "5554561212"
	pn.EndToEndIdentification = "End To End Identification"
	return pn
}
