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

// TestFedWireMessageCheckSameDaySettlement_FileREAD validates reading an CheckSameDaySettlement FedWireMessage
func TestFedWireMessageCheckSameDaySettlement_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-CheckSameDaySettlement.txt")
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

// TestFedWireMessageDepositSendersAccount_FileREAD validates reading an DepositSendersAccount FedWireMessage
func TestFedWireMessageDepositSendersAccount_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-DepositSendersAccount.txt")
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

// TestFedWireMessageFEDFundsReturned_FileREAD validates reading an FEDFundsReturned FedWireMessage
func TestFedWireMessageFEDFundsReturned_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-FEDFundsReturned.txt")
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

// TestFedWireMessageFEDFundsSold_FileREAD validates reading an FEDFundsSold FedWireMessage
func TestFedWireMessageFEDFundsSold_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-FEDFundsSold.txt")
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

// TestFedWireMessageDrawdownRequest_FileREAD validates reading an DrawdownRequest FedWireMessage
func TestFedWireMessageDrawdownRequest_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-DrawdownRequest.txt")
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

// TestFedWireMessageBankDrawdownRequest_FileREAD validates reading an BankDrawdownRequest FedWireMessage
func TestFedWireMessageBankDrawdownRequest_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-BankDrawdownRequest.txt")
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

// TestFedWireMessageCustomerCorporateDrawdownRequest_FileREAD validates reading an CustomerCorporateDrawdownRequest FedWireMessage
func TestFedWireMessageCustomerCorporateDrawdownRequest_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-BankDrawdownRequest.txt")
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

// TestFedWireMessageServiceMessage_FileREAD validates reading an ServiceMessage FedWireMessage
func TestFedWireMessageServiceMessage_FileREAD(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-BankDrawdownRequest.txt")
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