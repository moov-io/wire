// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
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

	require.False(t, c.parseFirstOption([]bool{}))
	require.False(t, c.parseFirstOption([]bool{false, false}))
	require.False(t, c.parseFirstOption([]bool{false, true}))
	require.True(t, c.parseFirstOption([]bool{true, false}))
	require.True(t, c.parseFirstOption([]bool{true, true}))
}

func TestConverters__parseVariableStringField(t *testing.T) {

	tests := []struct {
		input     string
		maxLength int
		want      string
		gotSize   int
		err       error
	}{
		{
			input:     "1234{0000}56789",
			maxLength: 3,
			want:      "",
			gotSize:   0,
			err:       ErrRequireDelimiter,
		},
		{
			input:     "1234{0000}56789",
			maxLength: 4,
			want:      "",
			gotSize:   0,
			err:       ErrRequireDelimiter,
		},
		{
			input:     "1234{0000}56789",
			maxLength: 5,
			want:      "",
			gotSize:   0,
			err:       ErrRequireDelimiter,
		},
		{
			input:     "123456789",
			maxLength: 5,
			want:      "",
			gotSize:   0,
			err:       ErrRequireDelimiter,
		},
		{
			input:     "1234*56789",
			maxLength: 5,
			want:      "1234",
			gotSize:   5,
			err:       nil,
		},
		{
			input:     "1234*56789",
			maxLength: 4,
			want:      "1234",
			gotSize:   5,
			err:       nil,
		},
		{
			input:     "1234*56789",
			maxLength: 3,
			want:      "123",
			gotSize:   5,
			err:       nil,
		},
		{
			input:     "*123456789",
			maxLength: 3,
			want:      "",
			gotSize:   1,
			err:       nil,
		},
	}
	c := &converters{}

	for index, tt := range tests {
		t.Run(fmt.Sprintf("sub_%d", index), func(t *testing.T) {
			got, size, err := c.parseVariableStringField(tt.input, tt.maxLength)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.gotSize, size)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestConverters__parseFixedStringField(t *testing.T) {

	tests := []struct {
		input     string
		maxLength int
		want      string
		gotSize   int
		err       error
	}{
		{
			input:     "1234{0000}56789",
			maxLength: 3,
			want:      "123",
			gotSize:   3,
			err:       nil,
		},
		{
			input:     "1234{0000}56789",
			maxLength: 4,
			want:      "1234",
			gotSize:   4,
			err:       nil,
		},
		{
			input:     "1234{0000}56789",
			maxLength: 5,
			want:      "",
			gotSize:   0,
			err:       ErrValidLength,
		},
		{
			input:     "1234*56789",
			maxLength: 5,
			want:      "",
			gotSize:   0,
			err:       ErrValidLength,
		},
		{
			input:     "1234*56789",
			maxLength: 4,
			want:      "1234",
			gotSize:   4,
			err:       nil,
		},
		{
			input:     "1234*56789",
			maxLength: 3,
			want:      "123",
			gotSize:   3,
			err:       nil,
		},
		{
			input:     "*123456789",
			maxLength: 3,
			want:      "",
			gotSize:   0,
			err:       ErrValidLength,
		},
	}
	c := &converters{}

	for index, tt := range tests {
		t.Run(fmt.Sprintf("sub_%d", index), func(t *testing.T) {
			got, size, err := c.parseFixedStringField(tt.input, tt.maxLength)
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.gotSize, size)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestConverters_alphaVariableField(t *testing.T) {
	tests := []struct {
		input          string
		variableLength bool
		maxLength      uint
		want           string
	}{
		{
			input:          "{0000}1234  ",
			variableLength: false,
			maxLength:      10,
			want:           "{0000}1234",
		},
		{
			input:          "{0000}1234  ",
			variableLength: true,
			maxLength:      10,
			want:           "{0000}1234",
		},
		{
			input:          "{0000}12",
			variableLength: false,
			maxLength:      10,
			want:           "{0000}12  ",
		},
		{
			input:          "{0000}12",
			variableLength: true,
			maxLength:      10,
			want:           "{0000}12",
		},
	}
	c := &converters{}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			opts := FormatOptions{
				VariableLengthFields: tt.variableLength,
			}
			require.Equal(t, tt.want, c.formatAlphaField(tt.input, tt.maxLength, opts))
		})
	}
}
