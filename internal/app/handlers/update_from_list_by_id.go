package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// Update returns an HTTP handler that updates information system details in ListInfSys table.
func Update(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Read the entire request body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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

		// Retrieve existing information system details
		sysInfo, err := service.GetInfo(req.Context(), keyInt)
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Unmarshal the request body into ListInfSysRequest struct
		unm := &model.ListInfSysRequest{
			Name:     sysInfo.Name,
			Owner:    sysInfo.Owner,
			Admin:    sysInfo.Admin,
			Contacts: sysInfo.Contacts,
		}
		if err = json.Unmarshal(body, &unm); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare updated information system data
		infRepo := &model.ListInfSys{
			Name:     unm.Name,
			Owner:    unm.Owner,
			Admin:    unm.Admin,
			Contacts: unm.Contacts,
		}

		// Update information system details using service
		id, err := service.UpdateInfo(req.Context(), infRepo, keyInt)
		if err != nil {
			if errors.Is(err, model.ErrNoRowsInResultSet) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare response with updated information system details
		resp := &model.ListInfSys{
			ID:       id,
			Name:     infRepo.Name,
			Owner:    infRepo.Owner,
			Admin:    infRepo.Admin,
			Contacts: infRepo.Contacts,
		}
		infJson, _ := json.Marshal(resp)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
