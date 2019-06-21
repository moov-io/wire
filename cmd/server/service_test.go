// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"github.com/moov-io/wire"
	"io/ioutil"
	"net/url"
	"strings"

	"testing"
	"unicode/utf8"
)

func testNextID(tb testing.TB) string {
	id := NextID()
	if utf8.RuneCountInString(id) != 16 {
		tb.Errorf("got other length %d for ID %s", len(id), id)
	}
	return id
}

func TestNextID(t *testing.T) {
	id := testNextID(t)

	_, err := url.Parse("https://moov.io/" + id)
	if err != nil {
		t.Fatalf("failed to parse url with id of %q: %v", id, err)
	}
}

func BenchmarkNextID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNextID(b)
	}
}

// test mocks are in mock_test.go

// CreateFile tests
func TestCreateFile(t *testing.T) {
	s := mockServiceInMemory()
	fwm := mockFEDWireMessage()
	id, err := s.CreateFile(fwm)
	if err != nil {
		t.Fatal(err.Error())
	}
	if id != "12345" {
		t.Errorf("expected %s received %s w/ error %s", "12345", id, err)
	}
}
func TestCreateFileIDExists(t *testing.T) {
	s := mockServiceInMemory()
	fwm := mockFEDWireMessage()
	id, err := s.CreateFile(fwm)

	fwm2 := wire.FEDWireMessage{ID: "12345"}
	id, err = s.CreateFile(&fwm2)
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %s", "ErrAlreadyExists", id, err)
	}
}

func TestCreateFileNoID(t *testing.T) {
	s := mockServiceInMemory()
	fwm := wire.NewFEDWireMessage()
	id, err := s.CreateFile(&fwm)
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %s", "NextID", id, err)
	}
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Service.GetFile tests

func TestGetFile(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s w/ error %s", "98765", f.ID, err)
	}
}

func TestGetFileNotFound(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("12345")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", f.ID, err)
	}
}

// Service.GetFiles tests

func TestGetFiles(t *testing.T) {
	s := mockServiceInMemory()
	files := s.GetFiles()
	if len(files) != 1 {
		t.Errorf("expected %s received %v", "1", len(files))
	}
}

// Service.DeleteFile tests

func TestDeleteFile(t *testing.T) {
	s := mockServiceInMemory()
	err := s.DeleteFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s", "nil", err)
	}
	_, err = s.GetFile("98765")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s", "ErrNotFound", err)
	}
}

// Service.GetFileContents tests

func TestGetFileContents(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFEDWireMessage())
	if err != nil {
		t.Fatal(err.Error())
	}

	// build file
	r, err := s.GetFileContents(id)
	if err != nil {
		if !strings.Contains(err.Error(), "required") {
			t.Fatal(err.Error())
		}
	}
	if r != nil {
		bs, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatal(err.Error())
		}

		if len(bs) == 0 {
			t.Fatal("expected to read fil")
		}
	}
}

// Service.ValidateFile tests

func TestValidateFile(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFEDWireMessage())
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := s.ValidateFile(id); err != nil {
		if !strings.Contains(err.Error(), "required ") {
			t.Fatal(err.Error())
		}
	}
}

func TestValidateFileMissing(t *testing.T) {
	s := mockServiceInMemory()
	err := s.ValidateFile("missing")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateFileBad(t *testing.T) {
	s := mockServiceInMemory()

	fId, _ := s.CreateFile(mockFEDWireMessage())

	// setup file, add batch
	f, err := s.GetFile(fId)
	if f == nil {
		t.Fatalf("couldn't get file: %v", err)
	}

	// validate
	if err := s.ValidateFile(fId); err == nil {
		t.Fatal("expected error")
	}
}
