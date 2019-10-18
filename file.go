// Copyright 2019 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// File contains the structures of a parsed WIRE File.
type File struct {
	ID             string         `json:"id"`
	FEDWireMessage FEDWireMessage `json:"fedWireMessage"`
}

// NewFile constructs a file template
func NewFile() *File {
	return &File{}
}

// AddFEDWireMessage appends a FEDWireMessage to the File
func (f *File) AddFEDWireMessage(fwm FEDWireMessage) FEDWireMessage {
	f.FEDWireMessage = fwm
	return f.FEDWireMessage
}

// Create will tabulate and assemble an WIRE file into a valid state.
//
// Create implementations are free to modify computable fields in a file and should
// call the Validate() function at the end of their execution.
func (f *File) Create() error {
	return nil
}

// Validate will never modify the file.
func (f *File) Validate() error {
	if err := f.FEDWireMessage.verify(); err != nil {
		return err
	}
	return nil
}

// FileFromJSON attempts to return a *File object assuming the input is valid JSON.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and callers should confirm with Validate(). Invalid files may
// be rejected by other Financial Institutions or ACH tools.
func FileFromJSON(bs []byte) (*File, error) {
	if len(bs) == 0 {
		//return nil, errors.New("no JSON data provided")
		return nil, nil
	}

	file := NewFile()
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(file); err != nil {
		return nil, fmt.Errorf("problem reading File: %v", err)
	}
	return file, nil
}
