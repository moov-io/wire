package main

import (
	"errors"
	"github.com/moov-io/wire"
	"sync"
)

type WireFileRepository interface {
	getFiles() ([]*wire.File, error)
	getFile(fileId string) (*wire.File, error)

	saveFile(file *wire.File) error
	deleteFile(fileId string) error
}

type memoryWireFileRepository struct {
	mu    sync.Mutex
	files map[string]*wire.File
}

func (r *memoryWireFileRepository) getFiles() ([]*wire.File, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var out []*wire.File
	for _, v := range r.files {
		f := *v
		out = append(out, &f)
	}
	return out, nil
}

func (r *memoryWireFileRepository) getFile(fileId string) (*wire.File, error) {
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

func (r *memoryWireFileRepository) saveFile(file *wire.File) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if file.ID == "" {
		return errors.New("empty Wire File ID")
	}
	r.files[file.ID] = file
	return nil
}

func (r *memoryWireFileRepository) deleteFile(fileId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if fileId == "" {
		return errors.New("empty Wire File Id")
	}

	delete(r.files, fileId)

	return nil
}