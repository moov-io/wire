package wire

// SenderToReceiver creates a SenderToReceiver
func mockSenderToReceiver() *SenderToReceiver {
	sr := NewSenderToReceiver()
	sr.CoverPayment.SwiftFieldTag = "Swift Field Tag"
	sr.CoverPayment.SwiftLineOne = "Swift Line One"
	sr.CoverPayment.SwiftLineTwo = "Swift Line Two"
	sr.CoverPayment.SwiftLineThree = "Swift Line Three"
	sr.CoverPayment.SwiftLineFour = "Swift Line Four"
	sr.CoverPayment.SwiftLineFive = "Swift Line Five"
	sr.CoverPayment.SwiftLineSix = "Swift Line Six"
	return sr
}
