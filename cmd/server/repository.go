// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/moov-io/wire"
	"sync"
	"time"
)

// Repository is the Service storage mechanism abstraction
type Repository interface {
	StoreFile(file *wire.File) error
	FindFile(id string) (*wire.File, error)
	FindAllFiles() []*wire.File
	DeleteFile(id string) error
}

type repositoryInMemory struct {
	mtx   sync.RWMutex
	files map[string]*wire.File

	ttl time.Duration

	logger log.Logger
}

// NewRepositoryInMemory is an in memory wire storage repository for files
func NewRepositoryInMemory(ttl time.Duration, logger log.Logger) Repository {
	repo := &repositoryInMemory{
		files:  make(map[string]*wire.File),
		ttl:    ttl,
		logger: logger,
	}

	if ttl <= 0*time.Second {
		// Don't run the cleanup if we've disabled the TTL
		return repo
	}

	// Run our anon goroutine to cleanup old ACH files
	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			repo.cleanupOldFiles()
		}
	}()

	return repo
}

func (r *repositoryInMemory) StoreFile(f *wire.File) error {
	if f == nil {
		return errors.New("nil FEDWIRE Message provided")
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.files[f.ID]; ok {
		return ErrAlreadyExists
	}
	r.files[f.ID] = f
	return nil
}

// FindFile retrieves a wire.File based on the supplied ID
func (r *repositoryInMemory) FindFile(id string) (*wire.File, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.files[id]; ok {
		return val, nil
	}
	return nil, ErrNotFound
}

// FindAllFiles returns all files that have been saved in memory
func (r *repositoryInMemory) FindAllFiles() []*wire.File {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	files := make([]*wire.File, 0, len(r.files))
	for i := range r.files {
		files = append(files, r.files[i])
	}
	return files
}

func (r *repositoryInMemory) DeleteFile(id string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.files, id)
	return nil
}

// cleanupOldFiles will iterate through r.files and delete entries which are older than
// the environmental variable FEDWIREMESSAGE_FILE_TTL (parsed as a time.Duration).
func (r *repositoryInMemory) cleanupOldFiles() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	removed := 0
	tooOld := time.Now().Add(-1 * r.ttl)
	tooOldStr := tooOld.Format("060102") // YYMMDD
	tooOldStr = tooOldStr[2:6]

	for i := range r.files {
		if r.files[i].FEDWireMessage.ReceiptTimeStamp != nil {
			if r.files[i].FEDWireMessage.ReceiptTimeStamp.ReceiptDate < tooOldStr {
				removed++
				delete(r.files, i)
			}
		}
	}

	if r.logger != nil {
		r.logger.Log("files", fmt.Sprintf("removed %d FEDWireMessage File files older than %v", removed, tooOld.Format(time.RFC3339)))
	}
}
