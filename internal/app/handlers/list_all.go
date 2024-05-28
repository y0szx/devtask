package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"errors"
	"net/http"
)

func List(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		sysInfo, err := service.ListInfo(req.Context())
		if err != nil {
			if errors.Is(err, model.ErrObjectNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		infJson, _ := json.Marshal(sysInfo)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
