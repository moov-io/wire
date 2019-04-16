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

	fwmFile, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err = fwmFile.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

}