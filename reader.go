// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// line is the current line being parsed from the input r
	// ToDo: is line needed?
	line string
	// currentFedWireMessage is the current FedWireMessage being parsed
	currentFedWireMessage FedWireMessage
	// lineNum is the line number of the file being parsed
	// ToDo: is lineNum needed?
	lineNum int
	// tagName holds the current tag being parsed.
	tagName string
}

// Read reads each line of the FED Wire file and defines which parser to use based
// on the first character of each line. It also enforces FED Wire formatting rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++

		/*		lineLength := len(line)

				// ToDo: Adjust below stump code
				if lineLength < 5 {
					msg := fmt.Sprintf(msgRecordLength, lineLength)
					err := &FileError{FieldName: "RecordLength", Value: strconv.Itoa(lineLength), Msg: msg}
					return r.File, r.error(err)
				}*/

		// ToDo: parseLine or parseTag?

		r.line = line
		if err := r.parseLine(); err != nil {
			return r.File, err
		}
	}

	return r.File, nil
}

func (r *Reader) parseLine() error {

	// ToDo:  For each type check length based on the tag

	return nil
}

func (r *Reader) parseTag() error {

	// ToDo:  For each type check length based on the tag

	return nil
}
