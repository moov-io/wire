// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestIssue104(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.json"))
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}
	if file == nil {
		t.Fatal("nil File")
	}

	var buf bytes.Buffer
	if err := NewWriter(&buf).Write(file); err != nil {
		t.Fatal(err)
	}

	// verify the output
	lines := strings.Split(buf.String(), "\n")
	if n := len(lines); n != 27 {
		t.Errorf("got %d lines", n)
	}
	for i := range lines {
		if lines[i] == "" {
			continue
		}

		prefix := string(lines[i][0:6])
		// tags are 4 digits surrounded by {..} - example: {1500}
		if prefix[0] != '{' || prefix[5] != '}' {
			t.Errorf("index #%d - missing tag: %q", i, prefix)
		}
	}
}
