package wire

// mockSenderSupplied creates a SenderSupplied
func mockSenderSupplied() *SenderSupplied {
	ss := NewSenderSupplied()
	ss.UserRequestCorrelation = "User Req"
	ss.MessageDuplicationCode = MessageDuplicationOriginal
	return ss
}
