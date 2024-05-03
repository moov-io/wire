package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/wire"
)

func FuzzReaderWriterWire(f *testing.F) {
	populateCorpus(f, true)

	f.Fuzz(func(t *testing.T, contents string) {
		// Read the sample file
		r := wire.NewReader(strings.NewReader(contents))
		// r.SetValidation(&ach.ValidateOpts{
		// 	SkipAll: true,
		// })
		file, err := r.Read()
		if err != nil {
			t.Skip()
		}

		// Write the file
		wire.NewWriter(io.Discard).Write(file)

		// Remove Validation override
		// file.SetValidation(&ach.ValidateOpts{
		// 	SkipAll: false,
		// })
		file.Validate()
	})
}

func populateCorpus(f *testing.F, wire bool) {
	f.Helper()

	err := filepath.Walk(filepath.Join("..", ".."), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(strings.ToLower(path))
		if (ext == ".txt" && wire) || (ext == ".json" && !wire) {
			bs, err := os.ReadFile(path)
			if err != nil {
				f.Fatal(err)
			}
			f.Add(string(bs))
		}
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}
