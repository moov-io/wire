// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestFile__FileFromJSON(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(bs) == 0 {
		t.Fatal("no bytes read")
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("file=%#v", file)

	if file.ID != "" {
		t.Error("id isn't set in JSON")
	}
	if file.FEDWireMessage.FIAdditionalFIToFI == nil {
		t.Error("FIAdditionalFIToFI shouldn't be nil")
	}
}
