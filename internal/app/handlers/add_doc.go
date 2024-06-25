package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// AddDocument handles HTTP requests to add a document to the storage service.
// It expects a document JSON payload in the request body and uses the provided service to add the document.
func AddDocument(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		key, ok := mux.Vars(req)[QueryParamKey]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		keyInt, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Decode the request body into a model.Documents struct
		var unm model.Documents
		err = json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create a model.Documents object with the parsed ID and decoded data
		doc := &model.Documents{
			ID:      keyInt,
			DocData: unm.DocData,
			DocName: unm.DocName,
		}

		// Call the service to add the document
		_, err = service.AddDoc(req.Context(), *doc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare the response JSON with the added document details
		resp := &model.Documents{
			ID:      keyInt,
			DocData: unm.DocData,
			DocName: unm.DocName,
		}
		infJson, _ := json.Marshal(resp)

		// Write the JSON response
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
