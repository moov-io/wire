// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Delimiter = "*"
)

// converters handles golang to WIRE type Converters
type converters struct{}

func (c *converters) parseNumField(r string) (s int) {
	s, _ = strconv.Atoi(strings.TrimSpace(r))
	return s
}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

func (c *converters) parseAlphaField(r string, max uint) string {
	ln := uint(len(r))
	if ln > max {
		return r[ln-max:]
	}

	rem := max - ln
	if !validSizeUint(rem) {
		return ""
	} else {
		r += strings.Repeat(" ", int(rem)) //nolint:gosec
	}
	return r
}

// numericStringField right-justified zero filled
func (c *converters) numericStringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}

	rem := max - ln
	if !validSizeUint(rem) {
		return ""
	} else {
		s = strings.Repeat("0", int(rem)) + s //nolint:gosec
	}
	return s
}

// alphaField returns the input formatted as a fixed-width alphanumeric string.
// If the length of s exceeds max, s will be truncated to max.
func (c *converters) alphaField(s string, max uint) string {
	return c.formatAlphaField(s, max, FormatOptions{})
}

// formatAlphaField returns the input formatted according to the FormatOptions.
// If the length of s exceeds max, s will be truncated to max. If options.VariableLengthFields
// is set, any trailing whitespace is stripped
func (c *converters) formatAlphaField(s string, max uint, options FormatOptions) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	if !options.VariableLengthFields {
		rem := max - ln
		if !validSizeUint(rem) {
			return ""
		} else {
			s += strings.Repeat(" ", int(rem)) //nolint:gosec
		}
	}
	return s
}

// parseFixedStringField will use to parse for mandatory fixed length elements
//   - Has fixed length
//   - Not permitted Delimiter
func (c *converters) parseFixedStringField(r string, maxLen int) (got string, size int, err error) {
	max := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	// Omit field?
	if len(r) == 0 {
		return
	}

	endIndex := -1
	delimiterIndex := -1

	if index := strings.Index(r, Delimiter); index > -1 {
		delimiterIndex = index
	}
	if index := strings.Index(r, "{"); index > -1 {
		endIndex = index
	}

	size = max(endIndex, delimiterIndex)
	if size == -1 {
		size = len(r)
	}

	if size > maxLen {
		size = maxLen
	} else if size < maxLen {
		size = 0
		err = ErrValidLength
		return
	}

	got = strings.TrimSpace(r[:size])

	return
}

// parseVariableStringField will use to parse for mandatory variable length/optional fixed length/optional variable length elements
//   - Has variable length
//   - Always required delimiter
func (c *converters) parseVariableStringField(r string, maxLen int) (got string, size int, err error) {
	// Omit field?
	if len(r) == 0 {
		return
	}

	if index := strings.Index(r, Delimiter); index > -1 {
		size = index
	} else {
		size = 0
		err = ErrRequireDelimiter
		return
	}

	if got = strings.TrimSpace(r[:size]); got == Delimiter {
		got = ""
	}

	if len(got) > maxLen {
		got = got[:maxLen]
	}

	// skip delimiter
	size++

	return
}

// get first option from options
func (c *converters) parseFirstOption(options []bool) bool {

	firstOption := false

	if len(options) > 0 {
		firstOption = options[0]
	}

	return firstOption
}

// strip delimiters
func (c *converters) stripDelimiters(data string) string {

	index := len(data)

	for i := len(data) - 1; i > 5; i-- {

		inspectLetter1 := string(data[i])
		inspectLetter2 := string(data[i-1])

		if inspectLetter1 != Delimiter || inspectLetter2 != Delimiter || i == 6 {
			index = i + 1
			break
		}

	}

	return data[:index]
}

// verify input data with read length
func (c *converters) verifyDataWithReadLength(data string, expected int) error {
	n := len(data) // utf8.RuneCountInString(data)
	if n == expected {
		return nil
	}
	if n > expected && data[expected:] == Delimiter {
		return nil
	}
	return fmt.Errorf("found data of %d length but expected %d", n, expected)
}
