package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddImage(service StorageInfo) http.Handler {
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

		var unm model.Images
		err = json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		image := &model.Images{
			ID:        keyInt,
			ImageID:   unm.ImageID,
			ImageData: unm.ImageData,
			ImageName: unm.ImageName,
		}
		image_id, err := service.AddImg(req.Context(), *image)
		fmt.Println(image_id, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := &model.Images{
			ID:        keyInt,
			ImageID:   image_id,
			ImageData: unm.ImageData,
			ImageName: unm.ImageName,
		}
		infJson, _ := json.Marshal(resp)
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
