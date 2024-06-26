package handlers

import (
	"context"
	"devtask/internal/model"
)

const QueryParamKey = "key"

// StorageInfo defines the interface for interacting with the storage service.
type StorageInfo interface {
	// Methods for interacting with ListInfSys entities
	GetInfo(ctx context.Context, id int64) (model.ListInfSys, error)
	AddInfo(ctx context.Context, info model.ListInfSys) (int64, error)
	UpdateInfo(ctx context.Context, info *model.ListInfSys, id int64) (int64, error)
	DeleteInfo(ctx context.Context, id int64) error
	ListInfo(ctx context.Context) ([]model.ListInfSys, error)

	// Methods for interacting with TableInfSystems entities
	GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddInfoIS(ctx context.Context, info model.TableInfSystems) (int64, error)
	UpdateInfoIS(ctx context.Context, info *model.TableInfSystems, id int64) (int64, error)

	// Methods for interacting with Images entities
	AddImg(ctx context.Context, info model.Images) (int64, error)
	GetImg(ctx context.Context, id int64) ([]model.Images, error)

	// Methods for interacting with Documents entities
	AddDoc(ctx context.Context, info model.Documents) (int64, error)
	GetDocs(ctx context.Context, id int64) ([]model.Documents, error)
}
