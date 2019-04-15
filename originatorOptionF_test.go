package wire

// mockOriginatorOptionF creates a OriginatorOptionF
func mockOriginatorOptionF() *OriginatorOptionF {
	oof := NewOriginatorOptionF()
	oof.PartyIdentifier = "TXID/123-45-6789"
	oof.Name = "Name"
	oof.LineOne = "LineOne"
	oof.LineTwo = "LineTwo"
	oof.LineThree = "LineThree"
	return oof
}