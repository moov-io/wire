// Copyright 2020 The Moov Authors
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
	Message   string
	TagLength int
	Length    int
}

// NewTagWrongLengthErr creates a new error of the TagWrongLengthErr type
func NewTagWrongLengthErr(tagLength, length int) TagWrongLengthErr {
	return TagWrongLengthErr{
		Message:   fmt.Sprintf("must be minimum %d characters and found %d", tagLength, length),
		TagLength: tagLength,
		Length:    length,
	}
}

func (e TagWrongLengthErr) Error() string {
	return e.Message
}

// ErrInvalidTag is the error given when a tag is invalid
type ErrInvalidTag struct {
	Message string
	Type    string
}

// NewErrInvalidTag creates a new error of the ErrInvalidTag type
func NewErrInvalidTag(tag string) ErrInvalidTag {
	return ErrInvalidTag{
		Message: fmt.Sprintf("%s is an invalid tag", tag),
		Type:    tag,
	}
}

func (e ErrInvalidTag) Error() string {
	return e.Message
}
