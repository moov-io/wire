// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
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

func (c *converters) parseVariableStringField(r string, maxLen int) (s string, index int) {
	if delimiterIndex := strings.Index(r, "*"); delimiterIndex > 0 {
		index = delimiterIndex
	}
	if delimiterIndex := strings.Index(r, "{"); delimiterIndex > 0 && delimiterIndex < index {
		index = delimiterIndex
	}

	if index == 0 || index > maxLen {
		index = maxLen
	}

	if index < len(r) {
		index = len(r)
	}

	s = strings.TrimSpace(r[:index])
	return s, index
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaVariableField(s string, max uint, isVariable bool) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}

	if isVariable {
		s += "*"
	} else {
		s += strings.Repeat(" ", int(max-ln))
	}

	return s
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
