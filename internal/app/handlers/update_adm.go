package handlers

//
//import (
//	"devtask/internal/model"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"github.com/gorilla/mux"
//	"io"
//	"net/http"
//	"strconv"
//)
//
//func UpdateAdm(service StoragePVZ) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		body, err := io.ReadAll(req.Body)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		key, ok := mux.Vars(req)[QueryParamKey]
//		if !ok {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		keyInt, err := strconv.ParseInt(key, 10, 64)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		var unm model.Admin
//		if err = json.Unmarshal(body, &unm); err != nil {
//			fmt.Println(err)
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		pvzRepo := &model.Admin{
//			Fio:      unm.Fio,
//			Email:    unm.Email,
//			Telegram: unm.Telegram,
//		}
//		id, err := service.AdmUpdate(req.Context(), pvzRepo, keyInt)
//		if err != nil {
//			if errors.Is(err, model.ErrNoRowsInResultSet) {
//				w.WriteHeader(http.StatusNotFound)
//				return
//			}
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//
//		resp := &model.Admin{
//			ID:       id,
//			Fio:      unm.Fio,
//			Email:    unm.Email,
//			Telegram: unm.Telegram,
//		}
//		pvzJson, _ := json.Marshal(resp)
//		_, err = w.Write(pvzJson)
//		if err != nil {
//			return
//		}
//	})
//}
