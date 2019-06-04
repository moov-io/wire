package wire

import (
	"github.com/moov-io/base"
	"os"
	"path/filepath"
	"testing"
)

// TestRead reads wire Files with different BusinessFunctionCodes
func TestRead(t *testing.T) {
	t.Run("BankTransfer", testRead(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.txt")))
	t.Run("CustomerTransfer", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransfer.txt")))
	t.Run("CustomerTransferPlus", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlus.txt")))
	t.Run("CheckSameDaySettlement", testRead(filepath.Join("test", "testdata", "fedWireMessage-CheckSameDaySettlement.txt")))
	t.Run("DepositSendersAccount", testRead(filepath.Join("test", "testdata", "fedWireMessage-DepositSendersAccount.txt")))
	t.Run("FEDFundsReturned", testRead(filepath.Join("test", "testdata", "fedWireMessage-FEDFundsReturned.txt")))
	t.Run("FEDFundsSold", testRead(filepath.Join("test", "testdata", "fedWireMessage-FEDFundsSold.txt")))
	t.Run("DrawDownRequest", testRead(filepath.Join("test", "testdata", "fedWireMessage-DrawDownRequest.txt")))
	t.Run("BankDrawDownRequest", testRead(filepath.Join("test", "testdata", "fedWireMessage-BankDrawDownRequest.txt")))
	t.Run("CustomerCorporateDrawDownRequest", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerCorporateDrawDownRequest.txt")))
	t.Run("ServiceMessage", testRead(filepath.Join("test", "testdata", "fedWireMessage-ServiceMessage.txt")))
	t.Run("CustomerTransferPlusCOVS", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusCOVS.txt")))
	t.Run("CustomerTransferPlusUnstructuredAddenda", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusUnstructuredAddenda.txt")))
	t.Run("CustomerTransferPlusStructuredRemittance", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusStructuredRemittance.txt")))
}

func testRead(filePathName string) func(t *testing.T) {
	return func(t *testing.T) {
		f, err := os.Open(filePathName)
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
}

func TestReadInvalidTag(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-InvalidTag.txt")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)

	_, err = r.Read()
	if err != nil {
		if !base.Has(err, NewErrInvalidTag(r.line[:6])) {
			t.Errorf("%T: %s", err, err)
		}
	}
}
