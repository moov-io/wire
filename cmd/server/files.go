// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"
	"github.com/moov-io/wire"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	filesCreated = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "wire_files_created",
		Help: "The number of WIRE files created",
	}, nil) // TODO(adam): add key/value pairs []string{"destination", "origin"}

	filesDeleted = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "wire_files_deleted",
		Help: "The number of WIRE files deleted",
	}, nil)

	errNoFileId           = errors.New("no File ID found")
	errNoFEDWireMessageID = errors.New("no FEDWireMessage ID found")
)

func addFileRoutes(logger log.Logger, r *mux.Router, repo WireFileRepository) {
	r.Methods("GET").Path("/files").HandlerFunc(getFiles(logger, repo))
	r.Methods("POST").Path("/files/create").HandlerFunc(createFile(logger, repo))
	r.Methods("GET").Path("/files/{fileId}").HandlerFunc(getFile(logger, repo))
	r.Methods("DELETE").Path("/files/{fileId}").HandlerFunc(deleteFile(logger, repo))
	r.Methods("GET").Path("/files/{fileId}/contents").HandlerFunc(getFileContents(logger, repo))
	r.Methods("GET").Path("/files/{fileId}/validate").HandlerFunc(validateFile(logger, repo))
	r.Methods("POST").Path("/files/{fileId}/FEDWireMessage").HandlerFunc(addFEDWireMessageToFile(logger, repo))
}

func getFileId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["fileId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoFileId)
		return ""
	}
	return v
}

func getFEDWireMessageID(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["FEDWireMessageID"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoFEDWireMessageID)
		return ""
	}
	return v
}

func getFiles(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		files, err := repo.getFiles() // TODO(adam): implement soft and hard limits
		if err != nil {
			err = logger.LogErrorf("error retrieving files: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Logf("found %d files", len(files))

		w.Header().Set("X-Total-Count", fmt.Sprintf("%d", len(files)))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(files)
	}
}

func createFile(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		file := wire.NewFile()
		if strings.Contains(r.Header.Get("Content-Type"), "application/json") {
			if err := json.NewDecoder(r.Body).Decode(file); err != nil {
				err = logger.LogErrorf("error reading request body: %v", err).Err()
				moovhttp.Problem(w, err)
				return
			}

			if err := file.Validate(); err != nil {
				err = logger.LogErrorf("file validation failed: %v", err).Err()
				moovhttp.Problem(w, err)
				return
			}
		} else {
			f, err := wire.NewReader(r.Body).ReadWithOpts(validateOptsFromQuery(r.URL.Query()))
			if err != nil {
				err = logger.LogErrorf("error reading file: %v", err).Err()
				moovhttp.Problem(w, err)
				return
			}
			file = &f
		}

		if file.ID == "" {
			file.ID = base.ID()
		}
		logger = logger.Set("fileID", log.String(file.ID))

		if err := repo.saveFile(file); err != nil {
			err = logger.LogErrorf("problem saving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Log("created file")

		// record a metric for files created
		filesCreated.Add(1) // TODO(adam): add key/value pairs (like in ACH)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(file)
	}
}

func getFile(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Log("file not found")
			http.NotFound(w, r)
			return
		}

		logger.Log("rendering file")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func deleteFile(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		if err := repo.deleteFile(fileId); err != nil {
			err = logger.LogErrorf("error deleting file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		logger.Log("deleted file")

		filesDeleted.Add(1)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		type response struct {
			Error error `json:"error"`
		}

		json.NewEncoder(w).Encode(&response{Error: nil})
	}
}

func getFileContents(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		if file == nil {
			logger.Log("file not found")
			http.NotFound(w, r)
			return
		}
		logger.Log("rendering file contents")

		writer, err := GetWriter(w, r)
		if err != nil {
			err = logger.LogErrorf("problem getting writer: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		if err := writer.Write(file); err != nil {
			err = logger.LogErrorf("problem rendering file contents: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validateFile(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Log("file not found")
			http.NotFound(w, r)
			return
		}

		if err := file.Create(); err != nil { // Create calls Validate
			err = logger.LogErrorf("file was invalid: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("validated file")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		type response struct {
			Error error `json:"error"`
		}

		json.NewEncoder(w).Encode(&response{Error: nil})
	}
}

func addFEDWireMessageToFile(logger log.Logger, repo WireFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestID := moovhttp.GetRequestID(r); requestID != "" {
			logger = logger.Set("requestID", log.String(requestID))
		}

		w = wrapResponseWriter(logger, w, r)

		var req wire.FEDWireMessage
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			err = logger.LogErrorf("error reading request body: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			logger.LogError(errNoFileId)
			return
		}
		logger = logger.Set("fileID", log.String(fileId))

		file, err := repo.getFile(fileId)
		if err != nil {
			err = logger.LogErrorf("error retrieving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		if file == nil {
			logger.Log("file not found")
			http.NotFound(w, r)
			return
		}

		file.FEDWireMessage = file.AddFEDWireMessage(req)
		if err := repo.saveFile(file); err != nil {
			err = logger.LogErrorf("error saving file: %v", err).Err()
			moovhttp.Problem(w, err)
			return
		}

		logger.Log("added FEDWireMessage to file")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

// GetWriter returns a new Writer based on request param `type` that writes to w.
// query param `format`=variable - we set VariableLengthFields to `true`
// query param `newline`=false - we set NewlineCharacter to ""
// no query param - writer defaults to fixed-length fields and use "\n" for NewlineCharacter.
func GetWriter(w io.Writer, r *http.Request) (*wire.Writer, error) {
	// default writer
	writer := wire.NewWriter(w)
	queryParams := r.URL.Query()
	// if the query params are not set then return the default writer
	if queryParams.Get("format") == "" &&
		queryParams.Get("newline") == "" {
		return writer, nil
	}

	// default format options
	lengthFormatOption := wire.VariableLengthFields(writer.FormatOptions.VariableLengthFields)
	newLineFormatOption := wire.NewlineCharacter(writer.FormatOptions.NewlineCharacter)

	// check for query param `type`. if "variable" then update VariableLengthFields OptionFunc to true
	fileType := queryParams.Get("format")
	if fileType == "variable" {
		lengthFormatOption = wire.VariableLengthFields(true)
	}

	// check for query param `newline`. if "false" then update NewlineCharacter OptionFunc to false
	newlineStr := queryParams.Get("newline")
	if newlineStr != "" {
		newline, err := strconv.ParseBool(newlineStr)
		if err != nil {
			return nil, err
		}
		if !newline {
			newLineFormatOption = wire.NewlineCharacter("")
		}
	}

	// new writer based on lengthFormatOption and newLineFormatOption
	writer = wire.NewWriter(w, lengthFormatOption, newLineFormatOption)
	return writer, nil
}

// validateOptsFromQuery returns a ValidateOpts struct based on the query params.
// If no validation query params were provided, opts will be nil.
func validateOptsFromQuery(query url.Values) (opts *wire.ValidateOpts) {
	if len(query) == 0 {
		return opts
	}

	const (
		skipMandatoryIMAD          = "skipMandatoryIMAD"
		allowMissingSenderSupplied = "allowMissingSenderSupplied"
	)

	validationNames := []string{
		skipMandatoryIMAD,
		allowMissingSenderSupplied,
	}

	for _, param := range validationNames {
		if set, _ := strconv.ParseBool(query.Get(param)); set {
			if opts == nil {
				opts = &wire.ValidateOpts{}
			}

			switch param {
			case skipMandatoryIMAD:
				opts.SkipMandatoryIMAD = true
			case allowMissingSenderSupplied:
				opts.AllowMissingSenderSupplied = true
			}
		}
	}

	return opts
}
