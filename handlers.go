package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jamesonwilliams/dynago-docs/db"
	"github.com/jamesonwilliams/dynago-docs/model"
)

var database = db.InMemDatabase{}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func DocumentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	documents := database.RetrieveDocuments()
	if err := getEncoder(w).Encode(documents); err != nil {
		panic(err)
	}
}

func DocumentShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var documentId int
	var err error
	if documentId, err = strconv.Atoi(vars["documentId"]); err != nil {
		panic(err)
	}
	document := database.RetrieveDocument(documentId)
	if document.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := getEncoder(w).Encode(document); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	jsonErr := model.JsonErr{Code: http.StatusNotFound, Text: "Not Found"}
	if err := getEncoder(w).Encode(jsonErr); err != nil {
		panic(err)
	}

}

//
// Test with this curl command:
//
// curl \
//     -H "Content-Type: application/json" \
//     -d '{"name":"New Document"}' \
//     http://localhost:8080/documents
//
func DocumentCreate(w http.ResponseWriter, r *http.Request) {
	var document model.Document
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &document); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := getEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := database.StoreDocument(document)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := getEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func getEncoder(w http.ResponseWriter) *json.Encoder {
	enc := json.NewEncoder(w)
	enc.SetIndent("", " ")
	return enc
}
