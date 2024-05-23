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

func UpdateISInfo(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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

		pvzInfo, err := service.GetInfoIS(req.Context(), keyInt)
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var unm model.TableInfSystems
		unm = *pvzInfo
		if err = json.Unmarshal(body, &unm); err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		pvzRepo := &model.TableInfSystems{
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
		id, err := service.UpdateInfoIS(req.Context(), pvzRepo, keyInt)
		if err != nil {
			if errors.Is(err, model.ErrNoRowsInResultSet) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
		pvzJson, _ := json.Marshal(resp)
		_, err = w.Write(pvzJson)
		if err != nil {
			return
		}
	})
}
