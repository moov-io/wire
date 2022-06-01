// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConverters__stripDelimiters(t *testing.T) {
	c := &converters{}

	require.Equal(t, "{0000}", c.stripDelimiters("{0000}"))
	require.Equal(t, "{0000}123", c.stripDelimiters("{0000}123"))
	require.Equal(t, "{0000}*", c.stripDelimiters("{0000}****"))
	require.Equal(t, "{000***", c.stripDelimiters("{000******"))
	require.Equal(t, "{0000}****A*", c.stripDelimiters("{0000}****A*****"))
	require.Equal(t, "{0000}****1*", c.stripDelimiters("{0000}****1*****"))
	require.Equal(t, "{0000}**** *", c.stripDelimiters("{0000}**** *****"))
}

func TestConverters__parseFirstOption(t *testing.T) {
	c := &converters{}

	require.Equal(t, false, c.parseFirstOption([]bool{}))
	require.Equal(t, false, c.parseFirstOption([]bool{false, false}))
	require.Equal(t, false, c.parseFirstOption([]bool{false, true}))
	require.Equal(t, true, c.parseFirstOption([]bool{true, false}))
	require.Equal(t, true, c.parseFirstOption([]bool{true, true}))
}

func TestConverters__parseVariableStringField(t *testing.T) {
	c := &converters{}

	got, size, err := c.parseVariableStringField("1234{0000}56789", 3)
	require.Equal(t, "123", got)
	require.Equal(t, 3, size)
	require.Nil(t, err)

	got, size, err = c.parseVariableStringField("1234{0000}56789", 4)
	require.Equal(t, "1234", got)
	require.Equal(t, 4, size)
	require.Nil(t, err)

	got, size, err = c.parseVariableStringField("1234{0000}56789", 5)
	require.Equal(t, "1234", got)
	require.Equal(t, 4, size)
	require.Nil(t, err)

	got, size, err = c.parseVariableStringField("1234*56789", 7)
	require.Equal(t, "1234", got)
	require.Equal(t, 5, size)
	require.Nil(t, err)

	got, size, err = c.parseVariableStringField("1234*56789", 3)
	require.Equal(t, "123", got)
	require.Equal(t, 3, size)
	require.Nil(t, err)
}

func TestConverters__alphaVariableField(t *testing.T) {
	c := &converters{}

	require.Equal(t, "{0000}1234", c.alphaVariableField("{0000}1234  ", 10, false))
	require.Equal(t, "{0000}1234", c.alphaVariableField("{0000}1234  ", 10, true))
	require.Equal(t, "{0000}12  ", c.alphaVariableField("{0000}12", 10, false))
	require.Equal(t, "{0000}12*", c.alphaVariableField("{0000}12", 10, true))
}
