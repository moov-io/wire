package main

import (
	"github.com/moov-io/base"
	"github.com/moov-io/wire"
	"testing"
)

type testWireFileRepository struct {
	err error

	file *wire.File
}

func (r *testWireFileRepository) getFiles() ([]*wire.File, error) {
	if r.err != nil {
		return nil, r.err
	}
	return []*wire.File{r.file}, nil
}

func (r *testWireFileRepository) getFile(fileId string) (*wire.File, error) {
	if r.err != nil {
		return nil, r.err
	}
	return r.file, nil
}

func (r *testWireFileRepository) saveFile(file *wire.File) error {
	if r.err == nil { // only persist if we're not error'ing
		r.file = file
	}
	return r.err
}

func (r *testWireFileRepository) deleteFile(fileId string) error {
	return r.err
}

func TestMemoryStorage(t *testing.T) {
	repo := &memoryWireFileRepository{
		files: make(map[string]*wire.File),
	}

	files, err := repo.getFiles()
	if err != nil || len(files) != 0 {
		t.Errorf("files=%#v error=%v", files, err)
	}

	f, err := readFile("fedWireMessage-CustomerTransfer.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.ID = base.ID()

	if err := repo.saveFile(f); err != nil {
		t.Fatal(err)
	}

	files, err = repo.getFiles()
	if err != nil || len(files) != 1 {
		t.Errorf("files=%#v error=%v", files, err)
	}

	file, err := repo.getFile(f.ID)
	if err != nil {
		t.Error(err)
	}
	if file.ID != f.ID {
		t.Errorf("file mis-match")
	}

	if err := repo.deleteFile(f.ID); err != nil {
		t.Error(err)
	}
	files, err = repo.getFiles()
	if err != nil || len(files) != 0 {
		t.Errorf("files=%#v error=%v", files, err)
	}
}
