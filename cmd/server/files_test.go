// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/moov-io/base"
	"github.com/moov-io/base/log"
	"github.com/moov-io/wire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	assert.Empty(t, getFileId(w, req))
	assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
}

func TestFEDWireMessageID(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	assert.Empty(t, getFEDWireMessageID(w, req))
	assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
}

func TestFiles_getFiles(t *testing.T) {
	repo := &testWireFileRepository{
		file: &wire.File{
			ID: base.ID(),
		},
	}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	req := httptest.NewRequest("GET", "/files", nil)

	t.Run("retrieves file", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		var files []*wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&files))
		require.Len(t, files, 1)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func readFile(filename string) (*wire.File, error) {
	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	if err != nil {
		return nil, err
	}
	f, err := wire.NewReader(fd).Read()
	return &f, err
}

func TestFiles_createWithInterfaceData(t *testing.T) {
	router := mux.NewRouter()
	repo := &testWireFileRepository{}
	addFileRoutes(log.NewTestLogger(), router, repo)

	w := httptest.NewRecorder()
	raw := `FTI0811 XFT811  {1500}30        T {1510}1000{1520}20220128DOVTAL3C000001{2000}000000010000{3100}123456789DOVETAIL BANK US F*{3320}XX22012800000051*{3400}021000089CITIBANK NYC*{3600}CTP{3620}3*3AC4C307-0FFB-4028-BD8E-53D55BDB90E1*{3700}SUSD0,*{4200}D000100002*{5000}T000100011*DRESDEFFXXX*`
	req, err := http.NewRequest(http.MethodPost, "/files/create", bytes.NewReader([]byte(raw)))
	require.NoError(t, err)

	// create the file
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code, w.Body)
	var created wire.File
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &created))
	require.NotEmpty(t, created.ID)

	// retrieve the file
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/files/%s", created.ID), nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, w.Body)
	var file wire.File
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &file))
	require.Equal(t, created.ID, file.ID)

	// retrieve the file contents
	w = httptest.NewRecorder()
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("/files/%s/contents", created.ID), nil)
	require.NoError(t, err)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, w.Body)
	contents := w.Body.String()
	expectedTags := []string{"{1500}", "{1510}", "{1520}", "{2000}", "{3100}", "{3320}", "{3400}", "{3600}", "{3620}", "{3700}", "{4200}", "{5000}"}
	for _, tag := range expectedTags {
		require.Contains(t, contents, tag)
	}
}

func TestFiles_createFile(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "fedWireMessage-CustomerTransfer.txt"))
	require.NoError(t, err)
	req := httptest.NewRequest("POST", "/files/create", bytes.NewReader(bs))
	repo := &testWireFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("creates file", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusCreated, w.Code, w.Body)
		var resp wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		assert.NotEmpty(t, resp.ID)
		assert.NotNil(t, resp.FEDWireMessage.FIAdditionalFIToFI)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_createFileJSON(t *testing.T) {
	repo := &testWireFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("creates file from JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		bs, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "fedWireMessage-BankTransfer.json"))
		require.NoError(t, err)
		req := httptest.NewRequest("POST", "/files/create", bytes.NewReader(bs))
		req.Header.Set("content-type", "application/json")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusCreated, w.Code, w.Body)

		var resp wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		assert.NotEmpty(t, resp.ID)
		assert.NotNil(t, resp.FEDWireMessage.FIAdditionalFIToFI)
	})

	t.Run("creates file from JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		bs, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "fedWireMessage-CustomerTransfer.json"))
		require.NoError(t, err)
		req := httptest.NewRequest("POST", "/files/create", bytes.NewReader(bs))
		req.Header.Set("content-type", "application/json")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusCreated, w.Code, w.Body)

		var resp wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		assert.NotEmpty(t, resp.ID)
		assert.NotEmpty(t, resp.FEDWireMessage)
	})

	t.Run("invalid JSON", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/files/create", strings.NewReader(`{...invalid-json`))
		req.Header.Set("content-type", "application/json")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_getFile(t *testing.T) {
	req := httptest.NewRequest("GET", "/files/foo", nil)
	repo := &testWireFileRepository{
		file: &wire.File{
			ID: base.ID(),
		},
	}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("gets file", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		var file wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&file))
		assert.NotEmpty(t, file.ID)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = nil

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})
}

func TestFiles_deleteFile(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/files/foo", nil)
	repo := &testWireFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("deletes file", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_getFileContents(t *testing.T) {
	req := httptest.NewRequest("GET", "/files/foo/contents", nil)
	fwm := mockFEDWireMessage()
	repo := &testWireFileRepository{
		file: &wire.File{
			ID:             base.ID(),
			FEDWireMessage: fwm,
		},
	}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = nil

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})
}

func TestFiles_getFileContentsWithFormatAndNewLineQueryParams(t *testing.T) {
	fwm := mockFEDWireMessage()
	repo := &testWireFileRepository{
		file: &wire.File{
			ID:             base.ID(),
			FEDWireMessage: fwm,
		},
	}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	// test with no format no newline=false
	req := httptest.NewRequest("GET", "/files/foo/contents", nil)
	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "\n")
	})

	// test with format=variable and newline=true
	req = httptest.NewRequest("GET", "/files/foo/contents?format=variable&newline=false", nil)
	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
		// file does not contain "\n"
		assert.NotContains(t, w.Body.String(), "\n")
	})

	// test with format=variable param and no `newline`
	req = httptest.NewRequest("GET", "/files/foo/contents?format=variable", nil)
	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "\n")
	})

	// test with no format and newline=false
	req = httptest.NewRequest("GET", "/files/foo/contents?newline=false", nil)
	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
		assert.NotContains(t, w.Body.String(), "\n")
	})

	// test with no format and newline=false
	req = httptest.NewRequest("GET", "/files/foo/contents?newline=test", nil)
	t.Run("gets file contents", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})
}

func TestFiles_validateFile(t *testing.T) {
	req := httptest.NewRequest("GET", "/files/foo/validate", nil)
	f, err := readFile("fedWireMessage-CustomerTransfer.txt")
	require.NoError(t, err)
	repo := &testWireFileRepository{file: f}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("validates file", func(t *testing.T) {
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		assert.Contains(t, w.Body.String(), `"{\"error\": null}"`)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = nil

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})
}

func TestFiles_addFEDWireMessageToFile(t *testing.T) {
	f, err := readFile("fedWireMessage-NoMessage.txt")
	require.Contains(t, err.Error(), "file validation failed")
	fwm := mockFEDWireMessage()
	repo := &testWireFileRepository{file: f}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("adds message to file", func(t *testing.T) {
		w := httptest.NewRecorder()
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(fwm))
		req := httptest.NewRequest("POST", "/files/foo/FEDWireMessage", &buf)

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		var out wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&out))
		assert.NotNil(t, out.FEDWireMessage.SenderSupplied)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(fwm))
		req := httptest.NewRequest("POST", "/files/foo/FEDWireMessage", &buf)

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body)
	})

	t.Run("file not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.file = nil
		repo.err = nil
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(fwm))
		req := httptest.NewRequest("POST", "/files/foo/FEDWireMessage", &buf)

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})
}

/*func TestFiles_removeFEDWireMessageFromFile(t *testing.T) {
	f, err := readFile("fedWireMessage-CustomerTransfer.txt")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testWireFileRepository{file: f}

	FEDWireMessageID := base.ID()
	repo.file.FEDWireMessage.ID = FedWireMessageID

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/foo/FEDWireMessage/%s", FEDWireMessageID), nil)

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}*/
