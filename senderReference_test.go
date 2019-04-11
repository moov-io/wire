package wire

//  mockSenderReference creates a SenderReference
func mockSenderReference() *SenderReference {
	sr := NewSenderReference()
	sr.SenderReference = "Sender Reference"
	return sr
}
