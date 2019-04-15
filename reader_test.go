package wire

import (
	"os"
	"testing"
)

// TestFedWireMessage validates reading an FedWireMessage file
func TestFedWireMessageFileRead(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if _, err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}
