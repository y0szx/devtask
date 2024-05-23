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
	//ResourcesConsumedID int64  `db:"resourcesConsumedID"`
	//AdminID             int64  `db:"adminID"`
}

//type Resources struct {
//	ID  int64  `db:"id"`
//	Cpu int64  `db:"cpu"`
//	Ram int64  `db:"ram"`
//	Hdd string `db:"hdd"`
//}
//
//type Admin struct {
//	ID       int64  `db:"id"`
//	Fio      string `db:"fio"`
//	Email    string `db:"email"`
//	Telegram string `db:"telegram"`
//}
