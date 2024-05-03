// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile__FileFromJSON(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(bs) == 0 {
		t.Fatal("no bytes read")
	}

	file, err := FileFromJSON(bs)

	require.NoError(t, err)
	require.Empty(t, file.ID, "id should not have been set")
	require.NotNil(t, file.FEDWireMessages[0].FIAdditionalFIToFI, "FIAdditionalFIToFI shouldn't be nil")
}
