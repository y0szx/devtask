package postgres

import (
	"context"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
)

// InfRepo struct contains a reference to the database
type InfRepo struct {
	db *db.Database
}

// NewInfo initializes a new InfRepo instance
func NewInfo(database *db.Database) *InfRepo {
	return &InfRepo{db: database}
}

// AddToList inserts a new record into the listinfsys table and returns the inserted ID
func (r *InfRepo) AddToList(ctx context.Context, inf *model.ListInfSys) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO listinfsys(name,owner,admin,contacts) VALUES ($1,$2, $3, $4) RETURNING id;`, inf.Name, inf.Owner, inf.Admin, inf.Contacts).Scan(&id)
	fmt.Println(err)
	return id, err
}

// UpdateInList updates an existing record in the listinfsys table by ID
func (r *InfRepo) UpdateInList(ctx context.Context, inf *model.ListInfSys, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE listinfsys SET name=$1, owner=$2, admin=$3, contacts=$4 WHERE id=$5;`, inf.Name, inf.Owner, inf.Admin, inf.Contacts, id)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

// GetFromList retrieves a record from the listinfsys table by ID
func (r *InfRepo) GetFromList(ctx context.Context, id int64) (*model.ListInfSys, error) {
	var a model.ListInfSys
	err := r.db.Get(ctx, &a, `SELECT id,name,owner,admin,contacts FROM listinfsys WHERE id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

// Delete removes a record from the listinfsys table by ID
func (r *InfRepo) Delete(ctx context.Context, id int64) error {
	cT, err := r.db.Exec(ctx, `DELETE FROM listinfsys WHERE id = ($1);`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.ErrObjectNotFound
		}
		return err
	}
	if cT.RowsAffected() == 0 {
		return model.ErrNoRowsInResultSet
	}
	return nil
}

// List retrieves all records from the listinfsys table
func (r *InfRepo) List(ctx context.Context) ([]model.ListInfSys, error) {
	var a []model.ListInfSys
	err := r.db.Select(ctx, &a, `SELECT id,name,owner,admin,contacts FROM listinfsys`)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}

	if a == nil {
		return nil, model.ErrObjectNotFound
	}
	return a, nil
}

// GetISTable retrieves a record from the infsys table by ID
func (r *InfRepo) GetISTable(ctx context.Context, id int64) (*model.TableInfSystems, error) {
	var a model.TableInfSystems
	err := r.db.Get(ctx, &a, `SELECT id,name,owner,vms,cpu,ram,hdd,software_used,admin_name,admin_email,admin_tg,resource_assignment,status FROM infsys WHERE id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

// AddISInfo inserts a new record into the infsys table and returns the inserted ID
func (r *InfRepo) AddISInfo(ctx context.Context, inf *model.TableInfSystems) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO infsys(id,name,owner,vms,cpu,ram,hdd,software_used,admin_name,admin_email,admin_tg,resource_assignment,status) VALUES ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;`, inf.ID, inf.Name, inf.Owner, inf.Vms, inf.Cpu, inf.Ram, inf.Hdd, inf.SoftwareUsed, inf.AdminName, inf.AdminEmail, inf.AdminTg, inf.ResourceAssignment, inf.Status).Scan(&id)
	return id, err
}

// UpdateISInfo updates an existing record in the infsys table by ID
func (r *InfRepo) UpdateISInfo(ctx context.Context, inf *model.TableInfSystems, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE infsys SET name=$1, owner=$2, vms=$3, cpu=$4, ram=$5, hdd=$6, software_used=$7, admin_name=$8, admin_email=$9, admin_tg=$10, resource_assignment=$11, status=$12 WHERE id=$13;`, inf.Name, inf.Owner, inf.Vms, inf.Cpu, inf.Ram, inf.Hdd, inf.SoftwareUsed, inf.AdminName, inf.AdminEmail, inf.AdminTg, inf.ResourceAssignment, inf.Status, id)
	fmt.Println(cT, err)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

// AddImage inserts a new image record into the images table and returns the inserted image ID
func (r *InfRepo) AddImage(ctx context.Context, inf *model.Images) (int64, error) {
	var image_id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO images(id,image_data,image_name) VALUES ($1,$2, $3) RETURNING image_id;`, inf.ID, inf.ImageData, inf.ImageName).Scan(&image_id)
	return image_id, err
}

// GetImages retrieves all image records from the images table by ID
func (r *InfRepo) GetImages(ctx context.Context, id int64) ([]model.Images, error) {
	var a []model.Images
	err := r.db.Select(ctx, &a, `SELECT image_id, id, image_data,image_name FROM images WHERE id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return a, nil
}

// AddDocument inserts a new document record into the docs table and returns the inserted ID
func (r *InfRepo) AddDocument(ctx context.Context, inf *model.Documents) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO docs(id,doc_data,doc_name) VALUES ($1,$2, $3) RETURNING id;`, inf.ID, inf.DocData, inf.DocName).Scan(&id)
	return id, err
}

// GetDocuments retrieves all document records from the docs table by ID
func (r *InfRepo) GetDocuments(ctx context.Context, id int64) ([]model.Documents, error) {
	var a []model.Documents
	err := r.db.Select(ctx, &a, `SELECT id,doc_data,doc_name FROM docs WHERE id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return a, nil
}
