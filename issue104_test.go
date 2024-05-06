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

	"github.com/stretchr/testify/require"
)

func TestIssue104(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "fedWireMessage-BankTransfer.json"))
	require.NoError(t, err)

	file, err := FileFromJSON(bs)
	require.NoError(t, err)
	require.NotNil(t, file)

	var buf bytes.Buffer
	require.NoError(t, NewWriter(&buf).Write(file))

	// verify the output
	lines := strings.Split(buf.String(), "\n")
	require.Len(t, lines, 26)

	for i := range lines {
		if lines[i] == "" {
			continue
		}

		prefix := lines[i][0:6]
		// tags are 4 digits surrounded by {..} - example: {1500}
		require.Equal(t, "{", string(prefix[0]), "tag missing opening bracket")
		require.Equal(t, "}", string(prefix[5]), "tag missing closing bracket")
	}
}
