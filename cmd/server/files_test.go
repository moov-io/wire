// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	return f, err
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

		require.Equal(t, http.StatusCreated, w.Code, w.Body)
		var resp wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
		assert.NotEmpty(t, resp.ID)
		assert.NotNil(t, resp.FEDWireMessages[0].FIAdditionalFIToFI)
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
		assert.NotNil(t, resp.FEDWireMessages[0].FIAdditionalFIToFI)
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
		assert.NotEmpty(t, resp.FEDWireMessages)
		assert.Empty(t, resp.ValidateOptions)
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

func TestFiles_createFile_missingSenderSupplied(t *testing.T) {
	repo := &testWireFileRepository{}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	// set up a message with no SenderSupplied field
	fwm := mockFEDWireMessage()
	fwm.SenderSupplied = nil
	file := wire.NewFile()
	file.AddFEDWireMessage(fwm)

	// create from JSON, without validation options, should fail without sender supplied
	resp, _ := routerUploadJSON(t, router, file)
	require.Equal(t, http.StatusBadRequest, resp.Code, resp.Body)
	require.Contains(t, resp.Body.String(), "SenderSupplied")

	// create from JSON, using validation options, should succeed without sender supplied
	file.ValidateOptions = wire.ValidateOpts{
		AllowMissingSenderSupplied: true,
	}
	resp, uploaded := routerUploadJSON(t, router, file)
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	assert.NotEmpty(t, uploaded.ID)
	assert.Nil(t, uploaded.FEDWireMessages[0].SenderSupplied)

	// make sure the file was saved
	resp, found := routerGetFile(t, router, uploaded.ID)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)
	assert.Equal(t, uploaded.ID, found.ID)
	assert.Nil(t, found.FEDWireMessages[0].SenderSupplied)
	assert.True(t, found.ValidateOptions.AllowMissingSenderSupplied)

	// get file contents calls Validate()
	// if isIncoming was passed properly, then the file should be valid
	resp = routerGetFileContents(t, router, uploaded.ID)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)
	assert.NotNil(t, resp.Body)

	// upload raw should succeed without sender supplied
	resp, rawUpload := routerUploadRaw(t, router, resp.Body,
		setQueryParam("allowMissingSenderSupplied", "true"),
	)
	require.Equal(t, http.StatusCreated, resp.Code, resp.Body)
	assert.NotEmpty(t, rawUpload.ID)
	assert.Nil(t, rawUpload.FEDWireMessages[0].SenderSupplied)
	assert.True(t, rawUpload.ValidateOptions.AllowMissingSenderSupplied)

	// get new file
	resp, found = routerGetFile(t, router, rawUpload.ID)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)
	assert.Equal(t, rawUpload.ID, found.ID)
	assert.Nil(t, found.FEDWireMessages[0].SenderSupplied)

	// get new file contents
	resp = routerGetFileContents(t, router, rawUpload.ID)
	require.Equal(t, http.StatusOK, resp.Code, resp.Body)
	assert.NotNil(t, resp.Body)
}

func setQueryParam(key, value string) func(values url.Values) url.Values {
	return func(values url.Values) url.Values {
		values.Set(key, value)
		return values
	}
}

func routerUploadJSON(t *testing.T, router *mux.Router, file *wire.File) (*httptest.ResponseRecorder, *wire.File) {
	bs, err := json.Marshal(file)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/files/create", bytes.NewReader(bs))
	req.Header.Set("content-type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	var resp *wire.File
	if w.Code == http.StatusCreated {
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	}
	return w, resp
}

func routerUploadRaw(t *testing.T, router *mux.Router, raw io.Reader, queryOpts ...func(values url.Values) url.Values) (*httptest.ResponseRecorder, *wire.File) {
	req := httptest.NewRequest("POST", "/files/create", raw)
	req.Header.Set("content-type", "text/plain")

	query := req.URL.Query()
	for _, opt := range queryOpts {
		query = opt(query)
	}
	req.URL.RawQuery = query.Encode()

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	var resp *wire.File
	if w.Code == http.StatusCreated {
		_ = json.NewDecoder(w.Body).Decode(&resp)
	}
	return w, resp
}

func routerGetFile(t *testing.T, router *mux.Router, id string) (*httptest.ResponseRecorder, *wire.File) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/"+id, nil)

	router.ServeHTTP(w, req)
	w.Flush()

	var file *wire.File
	if w.Code == http.StatusOK {
		_ = json.NewDecoder(w.Body).Decode(&file)
	}
	return w, file
}

func routerGetFileContents(t *testing.T, router *mux.Router, id string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/"+id+"/contents", nil)

	router.ServeHTTP(w, req)
	w.Flush()

	return w
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
		assert.Contains(t, w.Body.String(), `{"error":null}`)
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
			ID:              base.ID(),
			FEDWireMessages: []wire.FEDWireMessage{fwm},
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
			ID:              base.ID(),
			FEDWireMessages: []wire.FEDWireMessage{fwm},
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

		assert.Contains(t, w.Body.String(), `{"error":null}`)
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
	require.ErrorContains(t, err, "file validation failed")
	fwm := mockFEDWireMessage()
	repo := &testWireFileRepository{file: f}
	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)

	t.Run("adds message to file", func(t *testing.T) {
		w := httptest.NewRecorder()
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(fwm))
		req := httptest.NewRequest("POST", "/files/foo/messages", &buf)

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusOK, w.Code, w.Body)
		var out wire.File
		require.NoError(t, json.NewDecoder(w.Body).Decode(&out))
		assert.NotNil(t, out.FEDWireMessages[0].SenderSupplied)
	})

	t.Run("repo error", func(t *testing.T) {
		w := httptest.NewRecorder()
		repo.err = errors.New("bad error")
		var buf bytes.Buffer
		require.NoError(t, json.NewEncoder(&buf).Encode(fwm))
		req := httptest.NewRequest("POST", "/files/foo/messages", &buf)

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
		req := httptest.NewRequest("POST", "/files/foo/messages", &buf)

		router.ServeHTTP(w, req)
		w.Flush()

		assert.Equal(t, http.StatusNotFound, w.Code, w.Body)
	})
}

// func TestFiles_removeFEDWireMessageFromFile(t *testing.T) {
// 	f, err := readFile("fedWireMessage-CustomerTransfer.txt")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	repo := &testWireFileRepository{file: f}
//
// 	FEDWireMessageID := base.ID()
// 	repo.file.FEDWireMessage[0].ID = FEDWireMessageID
//
// 	w := httptest.NewRecorder()
// 	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/foo/FEDWireMessage/%s", FEDWireMessageID), nil)
//
// 	router := mux.NewRouter()
// 	addFileRoutes(log.NewNopLogger(), router, repo)
// 	router.ServeHTTP(w, req)
// 	w.Flush()
//
// 	if w.Code != http.StatusOK {
// 		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
// 	}
//
// 	// error case
// 	repo.err = errors.New("bad error")
//
// 	w = httptest.NewRecorder()
// 	router.ServeHTTP(w, req)
// 	w.Flush()
//
// 	if w.Code != http.StatusBadRequest {
// 		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
// 	}
// }
