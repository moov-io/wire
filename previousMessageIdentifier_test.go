package wire

//  mockPreviousMessageIdentifier creates a PreviousMessageIdentifier
func mockPreviousMessageIdentifier() *PreviousMessageIdentifier {
	pmi := NewPreviousMessageIdentifier()
	pmi.PreviousMessageIdentifier = "Previous Message Ident"
	return pmi
}
