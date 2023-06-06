// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"strconv"
	"strings"
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
	r += strings.Repeat(" ", int(max-ln))
	return r
}

// numericStringField right-justified zero filled
func (c *converters) numericStringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

// alphaField returns the input formatted as a fixed-width alphanumeric string.
// If the length of s exceeds max, s will be truncated to max.
func (c *converters) alphaField(s string, max uint) string {
	return c.formatAlphaField(s, max, FormatOptions{})
}

// formatAlphaField returns the input formatted according to the FormatOptions.
// If the length of s exceeds max, s will be truncated to max. If options.VariableLengthFields
// is set, any trailing whitespace is replaced with a single asterisk (*) character.
func (c *converters) formatAlphaField(s string, max uint, options FormatOptions) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}

	if options.VariableLengthFields {
		if max-ln > 0 {
			s += "*"
		}
	} else {
		s += strings.Repeat(" ", int(max-ln))
	}

	return s
}

func (c *converters) parseVariableStringField(r string, maxLen int) (got string, size int, err error) {
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	// Omit field?
	if len(r) == 0 {
		return
	}

	endIndex := maxLen
	delimiterIndex := maxLen

	if index := strings.Index(r, "*"); index > -1 {
		delimiterIndex = index
	}
	if index := strings.Index(r, "{"); index > -1 {
		endIndex = index
	}

	hasDelimiter := false
	size = min(endIndex, delimiterIndex)
	if size >= maxLen {
		size = maxLen
	} else if size < maxLen {
		if delimiterIndex == size {
			hasDelimiter = true
		}
	}

	if size > len(r) {
		size = 0
		err = ErrValidLength
		return
	}

	if got = strings.TrimSpace(r[:size]); got == "*" {
		got = ""
	}

	if hasDelimiter {
		size++
	}

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

		if inspectLetter1 != "*" || inspectLetter2 != "*" || i == 6 {
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
	if n > expected && data[expected:] == "*" {
		return nil
	}
	return fmt.Errorf("found data of %d length but expected %d", n, expected)
}
