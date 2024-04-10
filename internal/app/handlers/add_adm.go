package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"net/http"
)

func AddAdm(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.Admin
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pvzRepo := &model.Admin{
			ID:       unm.ID,
			Fio:      unm.Fio,
			Email:    unm.Email,
			Telegram: unm.Telegram,
		}
		id, err := service.AdmAdd(req.Context(), *pvzRepo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := &model.Admin{
			ID:       id,
			Fio:      unm.Fio,
			Email:    unm.Email,
			Telegram: unm.Telegram,
		}
		pvzJson, _ := json.Marshal(resp)
		_, err = w.Write(pvzJson)
		if err != nil {
			return
		}
	})
}
