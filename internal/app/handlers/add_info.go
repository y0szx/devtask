package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// AddISInfo handles HTTP requests to add information about information systems (IS) to the storage service.
// It expects an IS information JSON payload in the request body and uses the provided service to add the IS information.
func AddISInfo(service StorageInfo) http.Handler {
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

		// Decode the request body into a model.TableInfSystems struct
		var unm model.TableInfSystems
		err = json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Create a model.TableInfSystems object with the parsed ID and decoded data
		infRepo := &model.TableInfSystems{
			ID:                 keyInt,
			Name:               unm.Name,
			Owner:              unm.Owner,
			Vms:                unm.Vms,
			Cpu:                unm.Cpu,
			Ram:                unm.Ram,
			Hdd:                unm.Hdd,
			SoftwareUsed:       unm.SoftwareUsed,
			AdminName:          unm.AdminName,
			AdminEmail:         unm.AdminEmail,
			AdminTg:            unm.AdminTg,
			ResourceAssignment: unm.ResourceAssignment,
			Status:             unm.Status,
		}

		// Call the service to add the IS information
		id, err := service.AddInfoIS(req.Context(), *infRepo)
		fmt.Println(id, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Prepare the response JSON with the added IS information details
		resp := &model.TableInfSystems{
			ID:                 id,
			Name:               unm.Name,
			Owner:              unm.Owner,
			Vms:                unm.Vms,
			Cpu:                unm.Cpu,
			Ram:                unm.Ram,
			Hdd:                unm.Hdd,
			SoftwareUsed:       unm.SoftwareUsed,
			AdminName:          unm.AdminName,
			AdminEmail:         unm.AdminEmail,
			AdminTg:            unm.AdminTg,
			ResourceAssignment: unm.ResourceAssignment,
			Status:             unm.Status,
		}
		infJson, _ := json.Marshal(resp)

		// Write the JSON response
		_, err = w.Write(infJson)
		if err != nil {
			return
		}
	})
}
