package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"net/http"
)

func List(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		pvzInfo, err := service.ListInfo(req.Context())
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		pvzJson, _ := json.Marshal(pvzInfo)
		_, err = w.Write(pvzJson)
		if err != nil {
			return
		}
	})
}
