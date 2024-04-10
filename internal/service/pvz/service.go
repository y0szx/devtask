package pvz

import (
	"context"
	"devtask/internal/model"
)

type storagePVZ interface {
	AddPVZ(ctx context.Context, pvz *model.ListInfSys) (int64, error)
	GetPVZ(ctx context.Context, id int64) (*model.ListInfSys, error)
	Update(ctx context.Context, pvz *model.ListInfSys, id int64) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]model.ListInfSys, error)
	GetISTable(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddISInfo(ctx context.Context, pvz *model.TableInfSystems) (int64, error)
	UpdateISInfo(ctx context.Context, pvz *model.TableInfSystems, id int64) (int64, error)
	AddRes(ctx context.Context, pvz *model.Resources) (int64, error)
	GetRes(ctx context.Context, id int64) (*model.Resources, error)
	UpdateRes(ctx context.Context, pvz *model.Resources, id int64) (int64, error)
	AddAdm(ctx context.Context, pvz *model.Admin) (int64, error)
	GetAdm(ctx context.Context, id int64) (*model.Admin, error)
	UpdateAdm(ctx context.Context, pvz *model.Admin, id int64) (int64, error)
}

type Service struct {
	storage storagePVZ
}

func NewService(storage storagePVZ) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) GetInfo(ctx context.Context, id int64) (model.ListInfSys, error) {
	pvz, err := s.storage.GetPVZ(ctx, id)
	return *pvz, err
}

func (s Service) AddInfo(ctx context.Context, pvz model.ListInfSys) (int64, error) {
	id, err := s.storage.AddPVZ(ctx, &pvz)
	return id, err
}

func (s Service) UpdateInfo(ctx context.Context, pvz *model.ListInfSys, id int64) (int64, error) {
	id, err := s.storage.Update(ctx, pvz, id)
	return id, err
}

func (s Service) DeleteInfo(ctx context.Context, id int64) error {
	err := s.storage.Delete(ctx, id)
	return err
}

func (s Service) ListInfo(ctx context.Context) ([]model.ListInfSys, error) {
	pvzs, err := s.storage.List(ctx)
	return pvzs, err
}

func (s Service) GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error) {
	pvz, err := s.storage.GetISTable(ctx, id)
	return pvz, err
}

func (s Service) AddInfoIS(ctx context.Context, pvz model.TableInfSystems) (int64, error) {
	id, err := s.storage.AddISInfo(ctx, &pvz)
	return id, err
}

func (s Service) UpdateInfoIS(ctx context.Context, pvz *model.TableInfSystems, id int64) (int64, error) {
	id, err := s.storage.UpdateISInfo(ctx, pvz, id)
	return id, err
}

func (s Service) ResAdd(ctx context.Context, pvz model.Resources) (int64, error) {
	id, err := s.storage.AddRes(ctx, &pvz)
	return id, err
}

func (s Service) ResGet(ctx context.Context, id int64) (model.Resources, error) {
	pvz, err := s.storage.GetRes(ctx, id)
	return *pvz, err
}

func (s Service) ResUpdate(ctx context.Context, pvz *model.Resources, id int64) (int64, error) {
	id, err := s.storage.UpdateRes(ctx, pvz, id)
	return id, err
}

func (s Service) AdmAdd(ctx context.Context, pvz model.Admin) (int64, error) {
	id, err := s.storage.AddAdm(ctx, &pvz)
	return id, err
}

func (s Service) AdmGet(ctx context.Context, id int64) (model.Admin, error) {
	pvz, err := s.storage.GetAdm(ctx, id)
	return *pvz, err
}

func (s Service) AdmUpdate(ctx context.Context, pvz *model.Admin, id int64) (int64, error) {
	id, err := s.storage.UpdateAdm(ctx, pvz, id)
	return id, err
}
