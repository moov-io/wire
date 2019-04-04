// Copyright 2019 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"io"
)

// A Writer writes an x9.file to an encoded file.
//
// As returned by NewWriter, a Writer writes x9file structs into
// x9 formatted files.

// Writer struct
type Writer struct {
	w       *bufio.Writer
	lineNum int //current line being written
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Writer writes a single x9.file record to w
func (w *Writer) Write(file *File) error {
	if err := file.Validate(); err != nil {
		return err
	}
	w.lineNum = 0
	// Iterate over all records in the file
	if err := w.writeFedWireMessage(file); err != nil {
		return err
	}
	w.lineNum++

	return w.w.Flush()
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

func (w *Writer) writeFedWireMessage(file *File) error {
	for range file.FedWireMessages {

	}
	return nil
}