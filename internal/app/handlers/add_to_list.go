package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"net/http"
)

// Create handles HTTP requests to create a new entry in the storage service.
// It expects a JSON payload describing the entry in the request body and uses the provided service to add the entry.
func Create(service StorageInfo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.ListInfSysRequest
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create model.ListInfSys object from decoded request data
		infRepo := &model.ListInfSys{
			Name:     unm.Name,
			Owner:    unm.Owner,
			Admin:    unm.Admin,
			Contacts: unm.Contacts,
		}

		// Add the created entry using the service
		id, err := service.AddInfo(req.Context(), *infRepo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare JSON response with created entry details
		resp := &model.ListInfSys{
			ID:       id,
			Name:     infRepo.Name,
			Owner:    infRepo.Owner,
			Admin:    infRepo.Admin,
			Contacts: infRepo.Contacts,
		}
		infJson, _ := json.Marshal(resp)

		// Write the JSON response
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
