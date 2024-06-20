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

// UpdateISInfo returns an HTTP handler that updates information system details in InfSys table.
func UpdateISInfo(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Read the entire request body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Extract key from request parameters
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
		sysInfo, err := service.GetInfoIS(req.Context(), keyInt)
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Unmarshal the request body into TableInfSystems struct
		var unm model.TableInfSystems
		unm = *sysInfo
		if err = json.Unmarshal(body, &unm); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare updated information system data
		infRepo := &model.TableInfSystems{
			Name:               unm.Name,
			Owner:              unm.Owner,
			Vms:                unm.Vms,
			Cpu:                unm.Cpu,
			Ram:                unm.Ram,
			Hdd:                unm.Hdd,
			SoftwareUsed:       unm.SoftwareUsed,
			AdminName:          unm.AdminName,
			AdminEmail:         unm.AdminEmail,
			AdminTg:            unm.AdminTg,
			ResourceAssignment: unm.ResourceAssignment,
			Status:             unm.Status,
		}

		// Update information system details using service
		id, err := service.UpdateInfoIS(req.Context(), infRepo, keyInt)
		if err != nil {
			if errors.Is(err, model.ErrNoRowsInResultSet) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare response with updated information system details
		resp := &model.TableInfSystems{
			ID:                 id,
			Name:               unm.Name,
			Owner:              unm.Owner,
			Vms:                unm.Vms,
			Cpu:                unm.Cpu,
			Ram:                unm.Ram,
			Hdd:                unm.Hdd,
			SoftwareUsed:       unm.SoftwareUsed,
			AdminName:          unm.AdminName,
			AdminEmail:         unm.AdminEmail,
			AdminTg:            unm.AdminTg,
			ResourceAssignment: unm.ResourceAssignment,
			Status:             unm.Status,
		}
		infJson, _ := json.Marshal(resp)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
