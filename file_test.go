// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile__FileFromJSON(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.json"))
	require.NoError(t, err)

	file, err := FileFromJSON(bs)
	require.NoError(t, err)
	require.Empty(t, file.ID, "id should not have been set")
	require.NotNil(t, file.FEDWireMessages[0].FIAdditionalFIToFI, "FIAdditionalFIToFI shouldn't be nil")
}

func TestFile_readAndWriteBatch(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "batchFile.txt"))
	require.NoError(t, err)

	// Read the file in
	file, err := NewReader(bytes.NewReader(bs)).Read()
	require.NoError(t, err)

	require.Equal(t, 2, len(file.FEDWireMessages), "expected 2 FEDWireMessages")
	require.Equal(t, "BTR", file.FEDWireMessages[0].BusinessFunctionCode.BusinessFunctionCode)
	require.Equal(t, "CTR", file.FEDWireMessages[1].BusinessFunctionCode.BusinessFunctionCode)

	// Write the file back out
	var buf bytes.Buffer
	wOpts := []OptionFunc{
		VariableLengthFields(true),
		NewlineCharacter(""),
		MessageDelimiter("\n"),
	}
	require.NoError(t, NewWriter(&buf, wOpts...).Write(file))
	require.Equal(t, string(bs), buf.String())
}

func TestFile_readAndWriteBatch_incomingWithNewlines(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "batchFile_incomingWithNewlines.txt"))
	require.NoError(t, err)

	// Read the file in
	file, err := NewReader(bytes.NewReader(bs), IncomingFile()).Read()
	require.NoError(t, err)

	require.Equal(t, 3, len(file.FEDWireMessages), "expected 3 FEDWireMessages")
	require.Equal(t, "000000000001", file.FEDWireMessages[0].Amount.Amount)
	require.Equal(t, "000000000002", file.FEDWireMessages[1].Amount.Amount)
	require.Equal(t, "000000000003", file.FEDWireMessages[2].Amount.Amount)

	// Write the file back out
	var buf bytes.Buffer
	wOpts := []OptionFunc{
		VariableLengthFields(true),
		NewlineCharacter("\n"),
		MessageDelimiter("\n"),
	}
	require.NoError(t, NewWriter(&buf, wOpts...).Write(file))
	require.Equal(t, string(bs), buf.String())
}
