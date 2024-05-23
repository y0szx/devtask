package handlers

//
//import (
//	"devtask/internal/model"
//	"encoding/json"
//	"errors"
//	"github.com/gorilla/mux"
//	"net/http"
//	"strconv"
//)
//
//func GetAdm(service StoragePVZ) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
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
//		pvzInfo, err := service.AdmGet(req.Context(), keyInt)
//		if err != nil {
//			if errors.Is(err, model.ErrObjectNotFound) {
//				w.WriteHeader(http.StatusNotFound)
//				return
//			}
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//		pvzJson, _ := json.Marshal(pvzInfo)
//		_, err = w.Write(pvzJson)
//		if err != nil {
//			return
//		}
//	})
//}
