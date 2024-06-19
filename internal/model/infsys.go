package model

import (
	"errors"
)

type ListInfSys struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Owner    string `db:"owner"`
	Admin    string `db:"admin"`
	Contacts string `db:"contacts"`
}

type ListInfSysRequest struct {
	Name     string `db:"name"`
	Owner    string `db:"owner"`
	Admin    string `db:"admin"`
	Contacts string `db:"contacts"`
}

var ErrObjectNotFound = errors.New("not found")
var ErrNoRowsInResultSet = errors.New("no rows in result set")

type TableInfSystems struct {
	ID                 int64  `db:"id"`
	Name               string `db:"name"`
	Owner              string `db:"owner"`
	Vms                string `db:"vms"`
	Cpu                int    `db:"cpu"`
	Ram                int    `db:"ram"`
	Hdd                string `db:"hdd"`
	SoftwareUsed       string `db:"software_used"`
	AdminName          string `db:"admin_name"`
	AdminEmail         string `db:"admin_email"`
	AdminTg            string `db:"admin_tg"`
	ResourceAssignment string `db:"resource_assignment"`
	Status             bool   `db:"status"`
}

type Images struct {
	ImageID   int64  `db:"image_id"`
	ID        int64  `db:"id"`
	ImageData []byte `db:"image_data"`
	ImageName string `db:"image_name"`
}

type Documents struct {
	ID      int64  `db:"id"`
	DocData []byte `db:"doc_data"`
	DocName string `db:"doc_name"`
}
