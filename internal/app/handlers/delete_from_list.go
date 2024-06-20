package handlers

import (
	"devtask/internal/model"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Delete returns an HTTP handler that deletes an entity identified by key from the storage service.
func Delete(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Extract key from URL path parameters
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

		// Call DeleteInfo method on the service to delete the entity
		err = service.DeleteInfo(req.Context(), keyInt)
		if err != nil {
			// Handle specific error cases
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			if errors.Is(err, model.ErrNoRowsInResultSet) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
