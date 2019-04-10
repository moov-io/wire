package wire

//  mockSenderSupplied creates a SenderSupplied
func mockSenderSupplied() *SenderSupplied {
	ss := NewSenderSupplied()
	ss.tag = TagSenderSupplied
	ss.FormatVersion = FormatVersion
	ss.TestProductionCode = EnvironmentTest
	ss.MessageDuplicationCode = MessageDuplicationOriginal
	return ss
}
