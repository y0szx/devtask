package info

import (
	"context"
	"devtask/internal/model"
	"fmt"
)

type storageInfo interface {
	AddToList(ctx context.Context, info *model.ListInfSys) (int64, error)
	GetFromList(ctx context.Context, id int64) (*model.ListInfSys, error)
	UpdateInList(ctx context.Context, info *model.ListInfSys, id int64) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]model.ListInfSys, error)
	GetISTable(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddISInfo(ctx context.Context, info *model.TableInfSystems) (int64, error)
	UpdateISInfo(ctx context.Context, info *model.TableInfSystems, id int64) (int64, error)
}

type Service struct {
	storage storageInfo
}

func NewService(storage storageInfo) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetInfo(ctx context.Context, id int64) (model.ListInfSys, error) {
	inf, err := s.storage.GetFromList(ctx, id)
	return *inf, err
}

func (s Service) AddInfo(ctx context.Context, inf model.ListInfSys) (int64, error) {
	id, err := s.storage.AddToList(ctx, &inf)
	return id, err
}

func (s Service) UpdateInfo(ctx context.Context, inf *model.ListInfSys, id int64) (int64, error) {
	id, err := s.storage.UpdateInList(ctx, inf, id)
	return id, err
}

func (s Service) DeleteInfo(ctx context.Context, id int64) error {
	err := s.storage.Delete(ctx, id)
	fmt.Println(err)
	return err
}

func (s Service) ListInfo(ctx context.Context) ([]model.ListInfSys, error) {
	info, err := s.storage.List(ctx)
	return info, err
}

func (s Service) GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error) {
	inf, err := s.storage.GetISTable(ctx, id)
	fmt.Println(err)
	return inf, err
}

func (s Service) AddInfoIS(ctx context.Context, inf model.TableInfSystems) (int64, error) {
	id, err := s.storage.AddISInfo(ctx, &inf)
	return id, err
}

func (s Service) UpdateInfoIS(ctx context.Context, inf *model.TableInfSystems, id int64) (int64, error) {
	id, err := s.storage.UpdateISInfo(ctx, inf, id)
	return id, err
}
