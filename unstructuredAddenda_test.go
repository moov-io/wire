package wire

// UnstructuredAddenda creates a UnstructuredAddenda
func mockUnstructuredAddenda() *UnstructuredAddenda {
	ua := NewUnstructuredAddenda()
	ua.AddendaLength = "0014"
	ua.Addenda = "This is a test"
	return ua
}
