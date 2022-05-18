package wire

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	t.Run("DrawdownResponse", testRead(filepath.Join("test", "testdata", "fedWireMessage-DrawdownResponse.txt")))
	t.Run("BankDrawDownRequest", testRead(filepath.Join("test", "testdata", "fedWireMessage-BankDrawDownRequest.txt")))
	t.Run("CustomerCorporateDrawDownRequest", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerCorporateDrawDownRequest.txt")))
	t.Run("ServiceMessage", testRead(filepath.Join("test", "testdata", "fedWireMessage-ServiceMessage.txt")))
	t.Run("CustomerTransferPlusCOVS", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusCOVS.txt")))
	t.Run("CustomerTransferPlusUnstructuredAddenda", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusUnstructuredAddenda.txt")))
	t.Run("CustomerTransferPlusStructuredRemittance", testRead(filepath.Join("test", "testdata", "fedWireMessage-CustomerTransferPlusStructuredRemittance.txt")))
	t.Run("FedAppendedTags", testRead(filepath.Join("test", "testdata", "fedWireMessage-FedAppendedTags.txt")))
}

func testRead(filePathName string) func(t *testing.T) {
	return func(t *testing.T) {
		f, err := os.Open(filePathName)
		if err != nil {
			t.Fatalf("%T: %s", err, err)
		}
		defer f.Close()
		r := NewReader(f)

		fwmFile, err := r.Read()
		if err != nil {
			t.Fatalf("%T: %s", err, err)
		}

		// ensure we have a validated file structure
		require.NoError(t, fwmFile.Validate())
	}
}

func TestReadInvalidTag(t *testing.T) {
	f, err := os.Open("./test/testdata/fedWireMessage-InvalidTag.txt")
	if err != nil {
		t.Fatalf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)

	_, err = r.Read()

	require.Error(t, err)
	require.Contains(t, err.Error(), NewErrInvalidTag("{1599}").Error())
}

func TestReadShortLine(t *testing.T) {
	f, err := NewReader(strings.NewReader("00")).Read()

	require.Errorf(t, err, "expected an error - file: %v", f)
}

// hits the case where each individual tag is well-formed, but it's missing tags required for this Business Function Code
func TestRead_missingTag(t *testing.T) {
	f, err := os.Open(path.Join("test", "testdata", "fedWireMessage-MissingRequiredTag.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := NewReader(f)

	_, err = r.Read()

	require.EqualError(t, err, "file validation failed: FIBeneficiaryAdvice is a required field")
}
