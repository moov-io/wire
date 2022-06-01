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

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
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

// alphaVariableField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaVariableField(s string, max uint, isVariable bool) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}

	if isVariable {
		if max-ln > 0 {
			s += "*"
		}
	} else {
		s += strings.Repeat(" ", int(max-ln))
	}

	return s
}

func (c *converters) parseVariableStringField(r string, maxLen int) (s string, read int, err error) {

	// Omit field?
	if len(r) == 0 {
		return
	}

	read = maxLen

	if delimiterIndex := strings.Index(r, "*"); delimiterIndex > -1 {
		read = delimiterIndex
	} else if delimiterIndex := strings.Index(r, "{"); delimiterIndex > -1 {
		read = delimiterIndex
	}

	hasDelimiter := false
	if read > maxLen {
		read = maxLen
	} else if read < maxLen {
		hasDelimiter = true
	}

	if read > len(r) {
		read = 0
		err = ErrValidLengthSize
		return
	}

	if s = strings.TrimSpace(r[:read]); s == "*" {
		s = ""
	}

	if hasDelimiter {
		read++
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
		inspect1 := string(data[i])
		inspect2 := string(data[i-1])
		if inspect1 != "*" || inspect2 != "*" {
			index = i + 1
			break
		}
	}

	return data[:index]
}
