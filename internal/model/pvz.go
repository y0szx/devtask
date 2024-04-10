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
	ID                  int64  `db:"id"`
	Name                string `db:"name"`
	Owner               string `db:"owner"`
	Vms                 string `db:"vms"`
	SoftwareUsed        string `db:"softwareUsed"`
	ResourceAssignment  string `db:"resourceAssignment"`
	Status              bool   `db:"status"`
	ResourcesConsumedID int64  `db:"resourcesConsumedID"`
	AdminID             int64  `db:"adminID"`
}

type Resources struct {
	ID  int64  `db:"id"`
	Cpu int64  `db:"cpu"`
	Ram int64  `db:"ram"`
	Hdd string `db:"hdd"`
}

type Admin struct {
	ID       int64  `db:"id"`
	Fio      string `db:"fio"`
	Email    string `db:"email"`
	Telegram string `db:"telegram"`
}
