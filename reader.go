// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"bufio"
	"github.com/moov-io/base"
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// line is the current line being parsed from the input r
	line string
	// currentFedWireMessage is the current FedWireMessage being parsed
	currentFedWireMessage FedWireMessage
	// lineNum is the line number of the file being parsed
	lineNum int
	// tagName holds the current tag name being parsed.
	tagName string
	// errors holds each error encountered when attempting to parse the file
	errors base.ErrorList
}

// error returns a new ParseError based on err
func (r *Reader) parseError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*base.ParseError); ok {
		return err
	}
	return &base.ParseError{
		Line:   r.lineNum,
		Record: r.tagName,
		Err:    err,
	}
}

// addCurrentFedWireMessage creates the current FedWireMessage for the file being read. A successful
// current FedWireMessage will be added to r.File once parsed.
func (r *Reader) addCurrentFedWireMessage(fwm FedWireMessage) {
	r.currentFedWireMessage = FedWireMessage{}
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
		// ToDo: Line length check?
		// ToDo: File length Check?
		r.line = line
		if err := r.parseLine(); err != nil {
			return r.File, err
		}
	}
	return r.File, nil
}

func (r *Reader) parseLine() error {
	switch r.line[:6] {
	case TagSenderSupplied:
		if err := r.parseSenderSupplied(); err != nil {
			return err
		}
	case TagTypeSubType:
		if err := r.parseTypeSubType(); err != nil {
			return err
		}
	case TagInputMessageAccountabilityData:
		if err := r.parseInputMessageAccountabilityData(); err != nil {
			return err
		}
	case TagAmount:
		if err := r.parseAmount(); err != nil {
			return err
		}
	case TagSenderDepositoryInstitution:
		if err := r.parseSenderDepositoryInstitution(); err != nil {
			return err
		}
	case TagReceiverDepositoryInstitution:
		if err := r.parseReceiverDepositoryInstitution(); err != nil {
			return err
		}
	case TagBusinessFunctionCode:
		if err := r.parseBusinessFunctionCode(); err != nil {
			return err
		}
	default:
		return NewErrInvalidTag(r.line[:6])
	}
	return nil
}

// ToDo:  For each type check length based on the tag

/*func (r *Reader) parse() error {
	return nil
}*/

func (r *Reader) parseSenderSupplied() error {
	r.tagName = "SenderSupplied"
	if len(r.line) != 18 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(18, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.SenderSupplied.Parse(r.line)
	if err := r.File.FedWireMessage.SenderSupplied.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseTypeSubType() error {
	r.tagName = "TypeSubType"
	if len(r.line) != 10 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(10, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.TypeSubType.Parse(r.line)
	if err := r.File.FedWireMessage.TypeSubType.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseInputMessageAccountabilityData() error {
	r.tagName = "InputMessageAccountabilityData"
	if len(r.line) != 22 {
		r.errors.Add(r.parseError(NewTagWrongLengthErr(22, len(r.line))))
		return r.errors
	}
	r.File.FedWireMessage.InputMessageAccountabilityData.Parse(r.line)
	if err := r.File.FedWireMessage.InputMessageAccountabilityData.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseAmount() error {
	r.tagName = "Amount"
	r.File.FedWireMessage.Amount.Parse(r.line)
	if err := r.File.FedWireMessage.Amount.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseSenderDepositoryInstitution() error {
	r.tagName = "SenderDepositoryInstitution"
	r.File.FedWireMessage.SenderDepositoryInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.SenderDepositoryInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseReceiverDepositoryInstitution() error {
	r.tagName = "ReceiverDepositoryInstitution"
	r.File.FedWireMessage.ReceiverDepositoryInstitution.Parse(r.line)
	if err := r.File.FedWireMessage.ReceiverDepositoryInstitution.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

func (r *Reader) parseBusinessFunctionCode() error {
	r.tagName = "BusinessFunctionCode"
	r.File.FedWireMessage.BusinessFunctionCode.Parse(r.line)
	if err := r.File.FedWireMessage.BusinessFunctionCode.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}
