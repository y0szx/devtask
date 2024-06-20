package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"net/http"
)

// List returns an HTTP handler that lists all information systems.
func List(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Retrieve list of information systems using the service
		sysInfo, err := service.ListInfo(req.Context())
		if err != nil {
			// Handle specific error cases
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Serialize the list of information systems to JSON
		infJson, _ := json.Marshal(sysInfo)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
