// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"regexp"
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

func (c *converters) parseTag(r string) (s string, index int, err error) {

	if len(r) < 6 {
		err = ErrValidTagForType
		return
	}

	expectTag := r[:6]

	tagRegexString := `^{([0-9]{4})}$`
	reg := regexp.MustCompile(tagRegexString)
	if !reg.MatchString(expectTag) {
		err = ErrValidTagForType
		return
	}

	s = expectTag
	index = 6

	return
}

func (c *converters) parseVariableStringField(r string, maxLen int) (s string, read int, err error) {

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
		if max-ln > 0 {
			s += "*"
		}
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
