// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"errors"
	"fmt"
)

var (
	// ErrFileTooLong is the error given when a file exceeds the maximum possible length
	ErrFileTooLong = errors.New("file exceeds maximum possible number of lines")
)

// TagWrongLengthErr is the error given when a Tag is the wrong length
type TagWrongLengthErr struct {
	Message string
	Length  int
}

// NewTagWrongLengthErr creates a new error of the TagWrongLengthErr type
func NewTagWrongLengthErr(length int) TagWrongLengthErr {
	return TagWrongLengthErr{
		Message: fmt.Sprintf("must be %d characters and found %d", length),
		Length:  length,
	}
}

func (e TagWrongLengthErr) Error() string {
	return e.Message
}
