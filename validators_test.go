// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidators__validateOptionFName(t *testing.T) {
	v := &validator{}

	require.NoError(t, v.validateOptionFName("1/SMITH JOHN"))
	require.Error(t, v.validateOptionFName("1/"))
	require.Error(t, v.validateOptionFName("1"))
	require.Error(t, v.validateOptionFName(""))
	require.Error(t, v.validateOptionFName(" /"))
}
