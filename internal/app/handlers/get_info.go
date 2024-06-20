package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetISTable returns an HTTP handler that retrieves information of IS by ID from InfSys table.
func GetISTable(service StorageInfo) http.Handler {
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

		// Retrieve information of an IS by ID using the service
		sysInfo, err := service.GetInfoIS(req.Context(), keyInt)
		if err != nil {
			// Handle specific error cases
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		
		// Serialize the information system to JSON
		infJson, _ := json.Marshal(sysInfo)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
