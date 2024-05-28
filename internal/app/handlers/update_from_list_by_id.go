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

func Update(service StorageInfo) http.Handler {
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

		sysInfo, err := service.GetInfo(req.Context(), keyInt)
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
		infRepo := &model.ListInfSys{
			Name:     unm.Name,
			Owner:    unm.Owner,
			Admin:    unm.Admin,
			Contacts: unm.Contacts,
		}
		id, err := service.UpdateInfo(req.Context(), infRepo, keyInt)
		if err != nil {
			if errors.Is(err, model.ErrNoRowsInResultSet) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
