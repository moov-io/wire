// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/moov-io/wire"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestServer__CreateFileEndpoint creates JSON from existing WIRE Files and submits them to our
// HTTP API. We do this to ensure all SEC codes can be submitted and created via the HTTP API.
func TestServer__CreateFileEndpoint(t *testing.T) {
	files := getTestFiles()

	if len(files) == 0 {
		fmt.Printf("got no test Wire files to process")
	}

	for _, file := range files {
		f, err := os.Open(file.WIREFilepath)
		if err != nil {
			log.Fatal(err)
		}

		wireFile, err := wire.NewReader(f).Read()
		if err != nil {
			fmt.Printf("Issue reading file: %+v \n", err)
		}

		// ensure we have a validated file structure
		if err := wireFile.Validate(); err != nil {
			t.Errorf("Could not validate entire read file: %v", err)
		}

		// If you trust the file but it's formatting is off building will probably resolve the malformed file.
		if err := wireFile.Create(); err != nil {
			t.Errorf("Could not build file with read properties: %v", err)
		}

		if err := f.Close(); err != nil {
			t.Errorf("Problem closing %s: %v", file.WIREFilepath, err)
		}

		// Marshal the wire.File into JSON for HTTP API submission
		bs, err := json.Marshal(wireFile)
		if err != nil {
			t.Fatalf("Problem converting %s to JSON: %v", file.WIREFilepath, err)
		}

		httpReq, err := http.NewRequest("POST", "/files/create", bytes.NewReader(bs))
		if err != nil {
			t.Fatal(err)
		}
		httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

		createFileReq, err := decodeCreateFileRequest(context.TODO(), httpReq)
		if err != nil {
			t.Error(string(bs))
			t.Fatalf("file %s had error against HTTP decode: %v", file.WIREFilepath, err)
		}

		repo := NewRepositoryInMemory(testTTLDuration, nil)
		s := NewService(repo)

		endpoint := createFileEndpoint(s, repo, nil) // nil logger

		resp, err := endpoint(context.TODO(), createFileReq)
		if err != nil {
			t.Fatalf("%s couldn't be created against our HTTP API: %v", file.WIREFilepath, err)
		}
		if resp == nil {
			t.Fatalf("resp == nil")
		}
		createFileResponse, ok := resp.(createFileResponse)
		if !ok {
			t.Fatalf("couldn't convert %#v to createFileResponse", resp)
		}
		if createFileResponse.ID == "" || createFileResponse.Err != nil {
			t.Fatalf("%s failed HTTP API creation: %v", file.WIREFilepath, createFileResponse.Err)
		}
	}
}

type testFile struct {
	WIREFilepath string
	Filename     string
}

func getTestFiles() []testFile {

	matches, err := filepath.Glob("../test/testdata/*.txt")
	if err != nil {
		return nil
	}

	var testFiles []testFile
	for i := range matches {
		filename := filepath.Base(matches[i])

		testFiles = append(testFiles, testFile{
			WIREFilepath: matches[i],
			Filename:     strings.TrimSuffix(filename, ".txt"),
		})
	}

	return testFiles
}
