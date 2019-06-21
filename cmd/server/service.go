// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/moov-io/wire"
	"io"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// Service is a REST interface for interacting with ACH file structures
// TODO: Add ctx to function parameters to pass the client security token
type Service interface {
	// CreateFile creates a new ach file record and returns a resource ID
	CreateFile(fwm *wire.FEDWireMessage) (string, error)
	// AddFile retrieves a file based on the File id
	GetFile(id string) (*wire.File, error)
	// GetFiles retrieves all files accessible from the client.
	GetFiles() []*wire.File
	// DeleteFile takes a file resource ID and deletes it from the store
	DeleteFile(id string) error
	// GetFileContents creates a valid plaintext file in memory assuming it has a FileHeader and at least one Batch record.
	GetFileContents(id string) (io.Reader, error)
	// ValidateFile
	ValidateFile(id string) error
}

// service a concrete implementation of the service.
type service struct {
	store Repository
}

// NewService creates a new concrete service
func NewService(r Repository) Service {
	return &service{
		store: r,
	}
}

// CreateFile add a file to storage
// TODO(adam): the HTTP endpoint accepts malformed bodies (and missing data)
func (s *service) CreateFile(fwm *wire.FEDWireMessage) (string, error) {
	// create a new file
	f := wire.NewFile()
	f.ID = fwm.ID
	if f.ID == "" {
		id := NextID()
		f.ID = id
	}
	if err := s.store.StoreFile(f); err != nil {
		return "", err
	}
	return f.ID, nil
}

// GetFile returns a files based on the supplied id
func (s *service) GetFile(id string) (*wire.File, error) {
	f, err := s.store.FindFile(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return f, nil
}

func (s *service) GetFiles() []*wire.File {
	return s.store.FindAllFiles()
}

func (s *service) DeleteFile(id string) error {
	return s.store.DeleteFile(id)
}

func (s *service) GetFileContents(id string) (io.Reader, error) {
	f, err := s.GetFile(id)
	if err != nil {
		return nil, fmt.Errorf("problem reading file %s: %v", id, err)
	}
	if err := f.Create(); err != nil {
		return nil, fmt.Errorf("problem creating file %s: %v", id, err)
	}

	var buf bytes.Buffer
	w := wire.NewWriter(&buf)
	if err := w.Write(f); err != nil {
		return nil, fmt.Errorf("problem writing plaintext file %s: %v", id, err)
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, errors.New("empty ACH file contents")
	}

	return &buf, nil
}

func (s *service) ValidateFile(id string) error {
	f, err := s.GetFile(id)
	if err != nil {
		return fmt.Errorf("problem reading file %s: %v", id, err)
	}
	return f.Validate()
}
