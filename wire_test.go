// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestWire__ReadCrashers will attempt to parse files which have previously been reported
// as crashing. These files are typically generated via fuzzing, but might also be reported
// by users.
func TestWire__ReadCrashers(t *testing.T) {
	root := filepath.Join("test", "testdata", "crashers")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if (err != nil && err != filepath.SkipDir) || info == nil || info.IsDir() {
			return nil // Ignore SkipDir and directories
		}
		if strings.HasSuffix(path, ".output") {
			return nil // go-fuzz makes these which contain the panic's trace
		}

		fd, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("problem opening %s: %v", path, err)
		}

		// Read out test file and ensure we don't panic
		NewReader(fd).Read()
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
