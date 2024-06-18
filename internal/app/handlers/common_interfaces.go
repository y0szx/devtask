package handlers

import (
	"context"
	"devtask/internal/model"
)

const QueryParamKey = "key"

type StorageInfo interface {
	GetInfo(ctx context.Context, id int64) (model.ListInfSys, error)
	AddInfo(ctx context.Context, info model.ListInfSys) (int64, error)
	UpdateInfo(ctx context.Context, info *model.ListInfSys, id int64) (int64, error)
	DeleteInfo(ctx context.Context, id int64) error
	ListInfo(ctx context.Context) ([]model.ListInfSys, error)
	GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddInfoIS(ctx context.Context, info model.TableInfSystems) (int64, error)
	UpdateInfoIS(ctx context.Context, info *model.TableInfSystems, id int64) (int64, error)
	AddImg(ctx context.Context, info model.Images) (int64, error)
	GetImg(ctx context.Context, id int64) ([]model.Images, error)
}
