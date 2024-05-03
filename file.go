// Copyright 2020 The Moov Authors
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
	ID              string           `json:"id"`
	FEDWireMessages []FEDWireMessage `json:"fedWireMessages"`
	ValidateOptions ValidateOpts     `json:"validateOptions,omitempty"`
}

// NewFile constructs a file template
func NewFile(opts ...FilePropertyFunc) *File {
	f := &File{}

	for _, opt := range opts {
		opt(f)
	}

	return f
}

// SetValidation stores ValidateOpts on the FEDWireMessage's validation rules
func (f *File) SetValidation(opts ValidateOpts) {
	if f == nil {
		return
	}

	f.ValidateOptions = opts
}

// AddFEDWireMessage appends a FEDWireMessage to the File
func (f *File) AddFEDWireMessage(fwm FEDWireMessage) {
	if f != nil {
		f.FEDWireMessages = append(f.FEDWireMessages, fwm)
	}
}

// Validate will never modify the file.
func (f *File) Validate() error {
	if f == nil {
		return nil
	}

	if len(f.FEDWireMessages) == 0 {
		return fmt.Errorf("no FEDWireMessages")
	}

	for i := range f.FEDWireMessages {
		if err := f.FEDWireMessages[i].verifyWithOpts(f.ValidateOptions); err != nil {
			return err
		}
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
		// return nil, errors.New("no JSON data provided")
		return nil, nil
	}

	file := NewFile()
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(file); err != nil {
		return nil, fmt.Errorf("problem reading File: %v", err)
	}
	return file, nil
}

type FilePropertyFunc func(*File)

// OutgoingFile configures the FedWireMessage ValidationOpts for an outgoing file
func OutgoingFile() FilePropertyFunc {
	return func(f *File) {
		if f != nil {
			f.ValidateOptions.AllowMissingSenderSupplied = false
		}
	}
}

// IncomingFile configures the FedWireMessage ValidationOpts for an incoming file
func IncomingFile() FilePropertyFunc {
	return func(f *File) {
		if f != nil {
			f.ValidateOptions.AllowMissingSenderSupplied = true
		}
	}
}
