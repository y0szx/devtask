package info

import (
	"context"
	"devtask/internal/model"
)

// storageInfo interface defines the methods for interacting with the storage layer
type storageInfo interface {
	AddToList(ctx context.Context, info *model.ListInfSys) (int64, error)
	GetFromList(ctx context.Context, id int64) (*model.ListInfSys, error)
	UpdateInList(ctx context.Context, info *model.ListInfSys, id int64) (int64, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]model.ListInfSys, error)
	GetISTable(ctx context.Context, id int64) (*model.TableInfSystems, error)
	AddISInfo(ctx context.Context, info *model.TableInfSystems) (int64, error)
	UpdateISInfo(ctx context.Context, info *model.TableInfSystems, id int64) (int64, error)
	AddImage(ctx context.Context, info *model.Images) (int64, error)
	GetImages(ctx context.Context, id int64) ([]model.Images, error)
	AddDocument(ctx context.Context, info *model.Documents) (int64, error)
	GetDocuments(ctx context.Context, id int64) ([]model.Documents, error)
}

// Service struct encapsulates the storageInfo interface
type Service struct {
	storage storageInfo
}

// NewService creates a new Service instance with the provided storage
func NewService(storage storageInfo) *Service {
	return &Service{
		storage: storage,
	}
}

// GetInfo retrieves a ListInfSys record by ID
func (s Service) GetInfo(ctx context.Context, id int64) (model.ListInfSys, error) {
	inf, err := s.storage.GetFromList(ctx, id)
	return *inf, err
}

// AddInfo adds a new ListInfSys record and returns the inserted ID
func (s Service) AddInfo(ctx context.Context, inf model.ListInfSys) (int64, error) {
	id, err := s.storage.AddToList(ctx, &inf)
	return id, err
}

// UpdateInfo updates an existing ListInfSys record by ID
func (s Service) UpdateInfo(ctx context.Context, inf *model.ListInfSys, id int64) (int64, error) {
	id, err := s.storage.UpdateInList(ctx, inf, id)
	return id, err
}

// DeleteInfo deletes a ListInfSys record by ID
func (s Service) DeleteInfo(ctx context.Context, id int64) error {
	err := s.storage.Delete(ctx, id)
	return err
}

// ListInfo retrieves all ListInfSys records
func (s Service) ListInfo(ctx context.Context) ([]model.ListInfSys, error) {
	info, err := s.storage.List(ctx)
	return info, err
}

// GetInfoIS retrieves a TableInfSystems record by ID
func (s Service) GetInfoIS(ctx context.Context, id int64) (*model.TableInfSystems, error) {
	inf, err := s.storage.GetISTable(ctx, id)
	return inf, err
}

// AddInfoIS adds a new TableInfSystems record and returns the inserted ID
func (s Service) AddInfoIS(ctx context.Context, inf model.TableInfSystems) (int64, error) {
	id, err := s.storage.AddISInfo(ctx, &inf)
	return id, err
}

// UpdateInfoIS updates an existing TableInfSystems record by ID
func (s Service) UpdateInfoIS(ctx context.Context, inf *model.TableInfSystems, id int64) (int64, error) {
	id, err := s.storage.UpdateISInfo(ctx, inf, id)
	return id, err
}

// GetImg retrieves all Images records by ID
func (s Service) GetImg(ctx context.Context, id int64) ([]model.Images, error) {
	inf, err := s.storage.GetImages(ctx, id)
	return inf, err
}

// AddImg adds a new Images record and returns the inserted image ID
func (s Service) AddImg(ctx context.Context, inf model.Images) (int64, error) {
	id, err := s.storage.AddImage(ctx, &inf)
	return id, err
}

// GetDocs retrieves all Documents records by ID
func (s Service) GetDocs(ctx context.Context, id int64) ([]model.Documents, error) {
	inf, err := s.storage.GetDocuments(ctx, id)
	return inf, err
}

// AddDoc adds a new Documents record and returns the inserted ID
func (s Service) AddDoc(ctx context.Context, inf model.Documents) (int64, error) {
	id, err := s.storage.AddDocument(ctx, &inf)
	return id, err
}
