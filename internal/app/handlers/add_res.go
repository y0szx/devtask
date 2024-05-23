package handlers

//
//import (
//	"devtask/internal/model"
//	"encoding/json"
//	"net/http"
//)
//
//func AddRes(service StoragePVZ) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		var unm model.Resources
//		err := json.NewDecoder(req.Body).Decode(&unm)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		pvzRepo := &model.Resources{
//			ID:  unm.ID,
//			Cpu: unm.Cpu,
//			Ram: unm.Ram,
//			Hdd: unm.Hdd,
//		}
//		id, err := service.ResAdd(req.Context(), *pvzRepo)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			return
//		}
//
//		resp := &model.Resources{
//			ID:  id,
//			Cpu: unm.Cpu,
//			Ram: unm.Ram,
//			Hdd: unm.Hdd,
//		}
//		pvzJson, _ := json.Marshal(resp)
//		_, err = w.Write(pvzJson)
//		if err != nil {
//			return
//		}
//	})
//}
