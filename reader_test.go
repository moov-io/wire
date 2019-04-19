package wire

import (
	"os"
	"testing"
)

// TestFedWireMessageCustomerTransfer_FileREAD validates reading an CustomerTransfer FedWireMessage
func TestFedWireMessageCustomerTransfer_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-CustomerTransfer.txt")
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

// TestFedWireMessageBankTransfer_FileREAD validates reading an BankTransfer FedWireMessage
func TestFedWireMessageBankTransfer_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-BankTransfer.txt")
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

// TestFedWireMessageCustomerTransferPlus_FileREAD validates reading an CustomerTransferPlus FedWireMessage
func TestFedWireMessageCustomerTransferPlus_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-CustomerTransferPlus.txt")
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