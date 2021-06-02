// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverters_parseNumField(t *testing.T) {
	c := &converters{}
	assert.Equal(t, 123, c.parseNumField("123"))
	assert.Equal(t, 0, c.parseNumField(""))
	assert.Equal(t, 0, c.parseNumField("ABC"))
	assert.Equal(t, 123, c.parseNumField(" 123 "))
}

func TestConverters_parseStringField(t *testing.T) {
	c := &converters{}
	assert.Equal(t, "123", c.parseStringField("123"))
	assert.Equal(t, "", c.parseStringField(""))
	assert.Equal(t, "ABC", c.parseStringField("ABC"))
	assert.Equal(t, "ABC", c.parseStringField(" ABC "))
}

func TestConverters_alphaField(t *testing.T) {
	c := &converters{}
	assert.Equal(t, "AB", c.alphaField("ABC", 2))
	assert.Equal(t, "ABC", c.alphaField("ABC", 3))
	assert.Equal(t, "ABC ", c.alphaField("ABC", 4))
}

func TestConverters_numericStringField(t *testing.T) {
	c := &converters{}
	assert.Equal(t, "23", c.numericStringField("123", 2))
	assert.Equal(t, "123", c.numericStringField("123", 3))
	assert.Equal(t, "0123", c.numericStringField("123", 4))
}
