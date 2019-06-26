// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"sync"

	"github.com/moov-io/wire"
)

type WIREFileRepository interface {
	getFiles() ([]*wire.File, error)
	getFile(fileId string) (*wire.File, error)

	saveFile(file *wire.File) error
	deleteFile(fileId string) error
}

type memoryWIREFileRepository struct {
	mu    sync.Mutex
	files map[string]*wire.File
}



func (r *memoryWIREFileRepository) getFile(fileId string) (*wire.File, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.files {
		if r.files[i].ID == fileId {
			f := *r.files[i]
			return &f, nil
		}
	}
	return nil, nil
}

func (r *memoryWIREFileRepository) saveFile(file *wire.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if file.ID == "" {
		return errors.New("empty WIRE File ID")
	}
	r.files[file.ID] = file
	return nil
}

func (r *memoryWIREFileRepository) deleteFile(fileId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if fileId == "" {
		return errors.New("empty Wire File Id")
	}

	delete(r.files, fileId)

	return nil
}
