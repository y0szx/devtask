package handlers

import (
	"devtask/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddISInfo(service StoragePVZ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var unm model.TableInfSystems
		err := json.NewDecoder(req.Body).Decode(&unm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pvzRepo := &model.TableInfSystems{
			ID:                 unm.ID,
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
		id, err := service.AddInfoIS(req.Context(), *pvzRepo)
		fmt.Println(id, err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
		pvzJson, _ := json.Marshal(resp)
		_, err = w.Write(pvzJson)
		if err != nil {
			return
		}
	})
}
