package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"net/http"
)

func Create(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.ListInfSysRequest
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		infRepo := &model.ListInfSys{
			Name:     unm.Name,
			Owner:    unm.Owner,
			Admin:    unm.Admin,
			Contacts: unm.Contacts,
		}
		id, err := service.AddInfo(req.Context(), *infRepo)
		if err != nil {
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
