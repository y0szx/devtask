package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// AddImage handles HTTP requests to add an image to the storage service.
// It expects an image JSON payload in the request body and uses the provided service to add the image.
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

		// Decode the request body into a model.Images struct
		var unm model.Images
		err = json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create a model.Images object with the parsed ID and decoded data
		image := &model.Images{
			ID:        keyInt,
			ImageID:   unm.ImageID,
			ImageData: unm.ImageData,
			ImageName: unm.ImageName,
		}

		// Call the service to add the image
		image_id, err := service.AddImg(req.Context(), *image)
		fmt.Println(image_id, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare the response JSON with the added image details
		resp := &model.Images{
			ID:        keyInt,
			ImageID:   image_id,
			ImageData: unm.ImageData,
			ImageName: unm.ImageName,
		}
		infJson, _ := json.Marshal(resp)

		// Write the JSON response
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
