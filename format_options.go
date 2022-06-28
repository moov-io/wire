package wire

// FormatOptions specify options for writing wire records to strings
type FormatOptions struct {
	VariableLengthFields bool   // set to true to use variable length fields instead of fixed-width
	NewlineCharacter     string // determines line endings - "\n" by default
}
