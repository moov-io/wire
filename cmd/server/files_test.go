// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"github.com/moov-io/wire"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/base"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func TestFileId(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	if v := getFileId(w, req); v != "" {
		t.Errorf("unexpected fileId=%s", v)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected HTTP status: %d", w.Code)
	}
}

/*func TestFEDWireMessageID(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	if v := getFEDWireMessageID(w, req); v != "" {
		t.Errorf("unexpected fileId=%s", v)
	}
	if w.Code != http.StatusBadRequest {
		t.Errorf("unexpected HTTP status: %d", w.Code)
	}
}
*/
func TestFiles__getFiles(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files", nil)

	repo := &testWireFileRepository{
		file: &wire.File{
			ID: base.ID(),
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
	var files []*wire.File
	if err := json.NewDecoder(w.Body).Decode(&files); err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Errorf("unexpected %d ICL files: %#v", len(files), files)
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func readFile(filename string) (*wire.File, error) {
	fd, err := os.Open(filepath.Join("..", "..", "test", "testdata", filename))
	if err != nil {
		return nil, err
	}
	f, err := wire.NewReader(fd).Read()
	return &f, err
}

/*func TestFiles__createFile(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(f); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", &buf)

	repo := &testWireFileRepository{}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusCreated {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
	var resp wire.File
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}


	// error case
	repo.err = errors.New("bad error")
	if err := json.NewEncoder(&buf).Encode(f); err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}*/

func TestFiles__getFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo", nil)

	repo := &testWireFileRepository{
		file: &wire.File{
			ID: base.ID(),
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	var file wire.File
	if err := json.NewDecoder(w.Body).Decode(&file); err != nil {
		t.Fatal(err)
	}
	if file.ID == "" {
		t.Errorf("unexpected ICL file: %#v", file)
	}

	// error case
	repo.err = errors.New("bad error")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}


func TestFiles__deleteFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/files/foo", nil)

	repo := &testWireFileRepository{}

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
}

/*func TestFiles__getFileContents(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo/contents", nil)

	repo := &testWireFileRepository{
		file: &wire.File{
			ID: base.ID(),
		},
	}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	if v := w.Header().Get("Content-Type"); v != "text/plain" {
		t.Errorf("unexpected Content-Type: %s", v)
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

/*func TestFiles__validateFile(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/files/foo/validate", nil)

	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testWireFileRepository{file: f}

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	if !strings.Contains(w.Body.String(), `"{\"error\": null}"`) {
		t.Errorf("unexpected body: %v", w.Body.String())
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

/*func TestFiles__addFEDWireMessageToFile(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	fwm := f.FEDWireMessage
	repo := &testWireFileRepository{file: f}

	// encode our CashLetter into JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(fwm); err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/foo/cashLetters", &buf)

	router := mux.NewRouter()
	addFileRoutes(log.NewNopLogger(), router, repo)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
	var out wire.File
	if err := json.NewDecoder(w.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}
	if len(out.FEDWireMessage) != 1 {
		t.Errorf("CashLetters: %#v", out.CashLetters)
	}

	// error case
	repo.err = errors.New("bad error")
	if err := json.NewEncoder(&buf).Encode(cashLetter); err != nil {
		t.Fatal(err)
	}

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %v", w.Code, w.Body.String())
	}
}

func TestFiles__removeCashLetterFromFile(t *testing.T) {
	f, err := readFile("BNK20180905121042882-A.icl")
	if err != nil {
		t.Fatal(err)
	}
	repo := &testWireFileRepository{file: f}

	cashLetterId := base.ID()
	repo.file.CashLetters[0].ID = cashLetterId

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/foo/cashLetters/%s", cashLetterId), nil)

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
}
*/