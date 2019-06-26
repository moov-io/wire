// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"

	"encoding/json"

	"fmt"


	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/wire"
)

var (
	errNoFileId       = errors.New("no File ID found")
	errNoFEDWireMessageID = errors.New("No Wire ID found")
)

func addFileRoutes(logger log.Logger, r *mux.Router, repo WIREFileRepository) {
	r.Methods("GET").Path("/files").HandlerFunc(getFiles(logger, repo))
	r.Methods("POST").Path("/files/create").HandlerFunc(createFile(logger, repo))
	r.Methods("GET").Path("/files/{fileId}").HandlerFunc(getFile(logger, repo))
	r.Methods("POST").Path("/files/{fileId}").HandlerFunc(updateFileHeader(logger, repo))
	r.Methods("DELETE").Path("/files/{fileId}").HandlerFunc(deleteFile(logger, repo))

	r.Methods("GET").Path("/files/{fileId}/contents").HandlerFunc(getFileContents(logger, repo))
	r.Methods("GET").Path("/files/{fileId}/validate").HandlerFunc(validateFile(logger, repo))

	r.Methods("POST").Path("/files/{fileId}/FEDWireMessage").HandlerFunc(addFEDWireMEssageToFile(logger, repo))
	r.Methods("DELETE").Path("/files/{fileId}/cashLetters/{FEDWireMessage_Id}").HandlerFunc(deleteFEDWireMessageFromFile(logger, repo))
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

func getFiles(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		files, err := repo.getFiles() // TODO(adam): implement soft and hard limits
		if err != nil {
			logger.Log("files", fmt.Sprintf("error getting ICL files: %v", err), "requestId", moovhttp.GetRequestId(r))
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("found %d files", len(files)), "requestId", requestId)
		}

		w.Header().Set("X-Total-Count", fmt.Sprintf("%d", len(files)))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(files)
	}
}

func createFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req wire.File
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if req.ID == "" {
			req.ID = base.ID()
		}
		if err := repo.saveFile(&req); err != nil {
			logger.Log("files", fmt.Sprintf("problem saving file %s: %v", req.ID, err), "requestId", moovhttp.GetRequestId(r))
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("creatd file=%s", req.ID), "requestId", requestId)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(req)
	}
}

func getFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			logger.Log("files", fmt.Sprintf("problem reading file=%s: %v", fileId, err), "requestId", moovhttp.GetRequestId(r))
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("rendering file=%s", fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

func updateFileHeader(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req wire.FEDWireMessage
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		file.FEDWireMessage = req
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("updating FileHeader for file=%s", fileId), "requestId", requestId)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(file)
	}
}

func deleteFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		if err := repo.deleteFile(fileId); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("deleted file=%s", fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func getFileContents(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("rendering file=%s contents", fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "text/plain")
		if err := wire.NewWriter(w).Write(file); err != nil {
			logger.Log("files", fmt.Sprintf("problem rendering file=%s contents: %v", fileId, err))
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func validateFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if err := file.Create(); err != nil { // Create calls Validate
			if requestId := moovhttp.GetRequestId(r); requestId != "" {
				logger.Log("files", fmt.Sprintf("file=%s was invalid: %v", fileId, err), "requestId", requestId)
			}
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("validated file=%s", fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}

func addFEDWireMessageToFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req wire.FEDWireMessage
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		fileId := getFileId(w, r)
		if fileId == "" {
			return
		}
		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		//file.FEDWireMessage = append(file.AddFEDWireMessage(file.FEDWireMessage), req)
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("added CashLetter=%s to file=%s", req.ID, fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(file)
	}
}

// Not sure we need this

func deleteFEDWireMessageFromFile(logger log.Logger, repo WIREFileRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		fileId, FEDWireMessageID := getFileId(w, r), getFEDWireMessageID(w, r)
		if fileId == "" || FEDWireMessageID == "" {
			return
		}

		file, err := repo.getFile(fileId)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

			if file.FEDWireMessage.ID == FEDWireMessageID {
				// ToDo: Remove FEDWireMessage
			}
		if err := repo.saveFile(file); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if requestId := moovhttp.GetRequestId(r); requestId != "" {
			logger.Log("files", fmt.Sprintf("removed CashLetter=%s to file=%s", cashLetterId, fileId), "requestId", requestId)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(`{"error": null}`)
	}
}


