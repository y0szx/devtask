package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"net/http"
)

func Create(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.ListInfSysRequest
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pvzRepo := &model.ListInfSys{
			Name:     unm.Name,
			Owner:    unm.Owner,
			Admin:    unm.Admin,
			Contacts: unm.Contacts,
		}
		id, err := service.AddInfo(req.Context(), *pvzRepo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := &model.ListInfSys{
			ID:       id,
			Name:     pvzRepo.Name,
			Owner:    pvzRepo.Owner,
			Admin:    pvzRepo.Admin,
			Contacts: pvzRepo.Contacts,
		}
		pvzJson, _ := json.Marshal(resp)
		_, err = w.Write(pvzJson)
		if err != nil {
			return
		}
	})
}
