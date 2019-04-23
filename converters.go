// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"strings"
)

// converters handles golang to WIRE type Converters
type converters struct{}

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
