package model

import (
	"errors"
)

// ListInfSys represents the structure of informational system list entry in the database
type ListInfSys struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Owner    string `db:"owner"`
	Admin    string `db:"admin"`
	Contacts string `db:"contacts"`
}

// ListInfSysRequest represents the structure of a request to create or update informational systems list entry
type ListInfSysRequest struct {
	Name     string `db:"name"`
	Owner    string `db:"owner"`
	Admin    string `db:"admin"`
	Contacts string `db:"contacts"`
}

// Common errors for the model package
var ErrObjectNotFound = errors.New("not found")
var ErrNoRowsInResultSet = errors.New("no rows in result set")

// TableInfSystems represents the structure of detailed information about informational system entry in the database
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

// Images represents the structure of an image entry in the database
type Images struct {
	ImageID   int64  `db:"image_id"`
	ID        int64  `db:"id"`
	ImageData []byte `db:"image_data"`
	ImageName string `db:"image_name"`
}

// Documents represents the structure of a document entry in the database
type Documents struct {
	ID      int64  `db:"id"`
	DocData []byte `db:"doc_data"`
	DocName string `db:"doc_name"`
}
