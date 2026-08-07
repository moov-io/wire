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
		if len(contents) > 1<<20 {
			t.Skip()
		}

		r := wire.NewReader(strings.NewReader(contents))
		file, err := r.Read()
		if err != nil {
			// Still exercise Validate/Write on partial files when possible.
			_ = file.Validate()
			_ = wire.NewWriter(io.Discard).Write(&file)
			return
		}

		_ = wire.NewWriter(io.Discard).Write(&file)
		_ = file.Validate()
	})
}

func FuzzReaderWriterJSON(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		file, err := wire.FileFromJSON([]byte(contents))
		if err != nil || file == nil {
			return
		}

		_ = file.Validate()
		_ = wire.NewWriter(io.Discard).Write(file)
	})
}

func populateCorpus(f *testing.F, asWire bool) {
	f.Helper()

	f.Add("")
	f.Add("{}")

	// Prefer real wire samples under test/ and examples/.
	roots := []string{
		filepath.Join("..", "..", "test"),
		filepath.Join("..", "..", "examples"),
	}

	for _, root := range roots {
		_ = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil || info == nil || info.IsDir() {
				return nil
			}
			// Skip huge / binary-ish paths
			if strings.Contains(path, "fuzz") || strings.Contains(path, "crashers") {
				return nil
			}

			ext := filepath.Ext(strings.ToLower(path))
			name := strings.ToLower(filepath.Base(path))
			isWireText := ext == ".txt" || strings.Contains(name, "wire") || strings.HasPrefix(name, "fed")
			isJSON := ext == ".json"

			if (asWire && isWireText && !isJSON) || (!asWire && isJSON) {
				bs, err := os.ReadFile(path)
				if err != nil {
					return nil
				}
				// Cap seed size
				if len(bs) > 256*1024 {
					return nil
				}
				f.Add(string(bs))
			}
			return nil
		})
	}

	// Also include any previously found crashers as regression seeds.
	crasherDir := filepath.Join("crashers")
	_ = filepath.Walk(crasherDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".txt" || filepath.Ext(path) == "" {
			bs, err := os.ReadFile(path)
			if err == nil && len(bs) < 256*1024 {
				f.Add(string(bs))
			}
		}
		return nil
	})
}
