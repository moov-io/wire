// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"github.com/go-kit/kit/log"
	"github.com/moov-io/wire"
	"testing"
	"time"
)

var (
	testTTLDuration = 0 * time.Second // disable TTL expiry
)

func TestRepositoryFiles(t *testing.T) {
	r := NewRepositoryInMemory(testTTLDuration, nil)

	if v := len(r.FindAllFiles()); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	fwm := mockFEDWireMessage()
	f := &wire.File{
		ID:             NextID(),
		FEDWireMessage: *fwm,
	}
	if err := r.StoreFile(f); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	found, err := r.FindFile(f.ID)
	if err != nil || found == nil {
		t.Errorf("found=%v, err=%v", found, err)
	}

	if v := len(r.FindAllFiles()); v != 1 {
		t.Errorf("unexpected length: %d", v)
	}

	if err := r.DeleteFile(f.ID); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRepository__cleanupOldFiles(t *testing.T) {
	r := NewRepositoryInMemory(testTTLDuration, nil)
	if repo, ok := r.(*repositoryInMemory); !ok {
		t.Fatalf("unexpected repository: %T %#v", r, r)
	} else {
		// write a file and later verify it's cleaned up
		file := wire.NewFile()
		fwm := wire.NewFEDWireMessage()
		rts := wire.NewReceiptTimeStamp()
		rd := time.Now().Add(-1 * 24 * time.Hour).Format("060102")
		rts.ReceiptDate = rd[2:6]
		fwm.SetReceiptTimeStamp(rts)
		file.AddFEDWireMessage(fwm)
		repo.StoreFile(file)
		if n := len(repo.FindAllFiles()); n != 1 {
			t.Errorf("got %d WIRE files", n)
		}
		repo.cleanupOldFiles() // make sure we don't panic
		if n := len(repo.FindAllFiles()); n != 0 {
			t.Errorf("got %d WIRE files", n)
		}
	}

	// Create a repo with a logger
	r = NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	if repo, ok := r.(*repositoryInMemory); !ok {
		t.Fatalf("unexpected repository: %T %#v", r, r)
	} else {
		repo.cleanupOldFiles() // make sure we don't panic
	}
}
