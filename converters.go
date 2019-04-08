// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strconv"
	"strings"
	"time"
)

// converters handles golang to WIRE type Converters
type converters struct{}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

// formatYYYYMMDDDate takes a time.Time and returns a string of YYYYMMDD
func (c *converters) formatYYYYMMDDDate(t time.Time) string {
	return t.Format("20060102")
}

// parseYYYMMDDDate returns a time.Time when passed time as YYYYMMDD
func (c *converters) parseYYYYMMDDDate(s string) time.Time {
	t, _ := time.Parse("20060102", s)
	return t
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

// numericField right-justified, unsigned, and zero filled
func (c *converters) numericField(n int, max uint) string {
	s := strconv.Itoa(n)
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
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

// FormatTag adds {} to a tag
func (c *converters) FormatTag(s string) string {
	s = "{" + s + "}"
	return s
}
