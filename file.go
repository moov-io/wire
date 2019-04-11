// Copyright 2019 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

// ToDo:  Do we want tag to be exportable
// ToDo:  Do we want ID for each type?

// File contains the structures of a parsed WIRE File.
type File struct {
	ID             string         `json:"id"`
	FedWireMessage FedWireMessage `json:"fedWireMessage"`
}

// NewFile constructs a file template
func NewFile() *File {
	return &File{}
}

// AddFedWireMessage appends a FedWireMessage to the File
func (f *File) AddFedWireMessage(fwm FedWireMessage) FedWireMessage {
	f.FedWireMessage = fwm
	return f.FedWireMessage
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
	return nil
}
