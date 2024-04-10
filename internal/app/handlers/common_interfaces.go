package handlers

import (
	"context"
	"devtask/internal/model"
)

const QueryParamKey = "key"

type StoragePVZ interface {
	GetInfo(ctx context.Context, id int64) (model.ListInfSys, error)
	AddInfo(ctx context.Context, pvz model.ListInfSys) (int64, error)
	UpdateInfo(ctx context.Context, pvz *model.ListInfSys, id int64) (int64, error)
	DeleteInfo(ctx context.Context, id int64) error
	ListInfo(ctx context.Context) ([]model.ListInfSys, error)
	GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddInfoIS(ctx context.Context, pvz model.TableInfSystems) (int64, error)
	UpdateInfoIS(ctx context.Context, pvz *model.TableInfSystems, id int64) (int64, error)
	ResAdd(ctx context.Context, pvz model.Resources) (int64, error)
	ResUpdate(ctx context.Context, pvz *model.Resources, id int64) (int64, error)
	ResGet(ctx context.Context, id int64) (model.Resources, error)
	AdmAdd(ctx context.Context, pvz model.Admin) (int64, error)
	AdmUpdate(ctx context.Context, pvz *model.Admin, id int64) (int64, error)
	AdmGet(ctx context.Context, id int64) (model.Admin, error)
}
