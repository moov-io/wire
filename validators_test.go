// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidSize(t *testing.T) {
	require.True(t, validSizeInt(10))
	require.True(t, validSizeInt(1e7))

	require.False(t, validSizeInt(1e8+1))
	require.False(t, validSizeInt(1e9))
	require.False(t, validSizeInt(math.MaxInt))

	t.Run("don't grow", func(t *testing.T) {
		ua := &UnstructuredAddenda{}
		ua.AddendaLength = fmt.Sprintf("%0.0f", 1e9)
		expected := "1000"
		require.Equal(t, expected, ua.String())
	})

	t.Run("int", func(t *testing.T) {
		require.False(t, validSizeInt(int(1e9)))
	})

	t.Run("uint", func(t *testing.T) {
		a := uint(100)
		b := uint(201)

		require.False(t, validSizeUint(a-b))
		require.True(t, validSizeUint(b-a))
	})
}

func TestValidators__validateOptionFName(t *testing.T) {
	v := &validator{}

	require.NoError(t, v.validateOptionFName("1/SMITH JOHN"))
	require.Error(t, v.validateOptionFName("1/"))
	require.Error(t, v.validateOptionFName("1"))
	require.Error(t, v.validateOptionFName(""))
	require.Error(t, v.validateOptionFName(" /"))
}

func TestValidators__isAlphanumeric(t *testing.T) {
	v := &validator{}

	require.NoError(t, v.isAlphanumeric("Telepathic Bank (U.K.) / Acct #12345-ABC"))
	require.Error(t, v.isAlphanumeric("{1100}"))
	require.Error(t, v.isAlphanumeric("*"))
}
