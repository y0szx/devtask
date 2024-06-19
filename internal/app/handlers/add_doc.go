package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddDocument(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

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

		var unm model.Documents
		err = json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		doc := &model.Documents{
			ID:      keyInt,
			DocData: unm.DocData,
			DocName: unm.DocName,
		}
		id, err := service.AddDoc(req.Context(), *doc)
		fmt.Println(id, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := &model.Documents{
			ID:      keyInt,
			DocData: unm.DocData,
			DocName: unm.DocName,
		}
		infJson, _ := json.Marshal(resp)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}