package postgres

import (
	"context"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	"errors"
	"github.com/jackc/pgx/v4"
)

type PVZRepo struct {
	db *db.Database
}

func NewPVZs(database *db.Database) *PVZRepo {
	return &PVZRepo{db: database}
}

func (r *PVZRepo) AddPVZ(ctx context.Context, pvz *model.ListInfSys) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO listinfsys(name,owner,admin,contacts) VALUES ($1,$2, $3, $4) RETURNING id;`, pvz.Name, pvz.Owner, pvz.Admin, pvz.Contacts).Scan(&id)
	return id, err
}

func (r *PVZRepo) Update(ctx context.Context, pvz *model.ListInfSys, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE listinfsys SET name=$1, owner=$2, admin=$3, contacts=$4 WHERE id=$5;`, pvz.Name, pvz.Owner, pvz.Admin, pvz.Contacts, id)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

func (r *PVZRepo) GetPVZ(ctx context.Context, id int64) (*model.ListInfSys, error) {
	var a model.ListInfSys
	err := r.db.Get(ctx, &a, `SELECT id,name,owner,admin,contacts FROM listinfsys where id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (r *PVZRepo) Delete(ctx context.Context, id int64) error {
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

func (r *PVZRepo) List(ctx context.Context) ([]model.ListInfSys, error) {
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

func (r *PVZRepo) GetISTable(ctx context.Context, id int64) (*model.TableInfSystems, error) {
	var a model.TableInfSystems
	err := r.db.Get(ctx, &a, `SELECT id,name,owner,vms,softwareUsed,resourceAssignment,status,resourcesConsumedID,adminID FROM infsys where id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (r *PVZRepo) AddISInfo(ctx context.Context, pvz *model.TableInfSystems) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO infsys(id, name,owner,vms,softwareUsed,resourceAssignment,status,resourcesConsumedID,adminID) VALUES ($1,$2, $3, $4, $5, $6, $7,$1,$1) RETURNING id;`, pvz.ID, pvz.Name, pvz.Owner, pvz.Vms, pvz.SoftwareUsed, pvz.ResourceAssignment, pvz.Status).Scan(&id)
	return id, err
}

func (r *PVZRepo) UpdateISInfo(ctx context.Context, pvz *model.TableInfSystems, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE infsys SET name=$1, owner=$2, vms=$3, softwareUsed=$4, resourceAssignment=$5, status=$6 WHERE id=$7;`, pvz.Name, pvz.Owner, pvz.Vms, pvz.SoftwareUsed, pvz.SoftwareUsed, pvz.ResourceAssignment, pvz.Status, id)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

func (r *PVZRepo) AddRes(ctx context.Context, pvz *model.Resources) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO resources(id,cpu,ram,hdd) VALUES ($1,$2, $3, $4) RETURNING id;`, pvz.ID, pvz.Cpu, pvz.Ram, pvz.Hdd).Scan(&id)
	return id, err
}

func (r *PVZRepo) GetRes(ctx context.Context, id int64) (*model.Resources, error) {
	var a model.Resources
	err := r.db.Get(ctx, &a, `SELECT id,cpu,ram,hdd FROM resources where id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (r *PVZRepo) UpdateRes(ctx context.Context, pvz *model.Resources, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE resources SET cpu=$1, ram=$2, hdd=$3 WHERE id=$4;`, pvz.Cpu, pvz.Ram, pvz.Hdd, id)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}

func (r *PVZRepo) AddAdm(ctx context.Context, pvz *model.Admin) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, `INSERT INTO admin(id,fio,email,telegram) VALUES ($1, $2, $3, $4) RETURNING id;`, pvz.ID, pvz.Fio, pvz.Email, pvz.Telegram).Scan(&id)
	return id, err
}

func (r *PVZRepo) GetAdm(ctx context.Context, id int64) (*model.Admin, error) {
	var a model.Admin
	err := r.db.Get(ctx, &a, `SELECT id,fio,email,telegram FROM admin where id=$1`, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrObjectNotFound
		}
		return nil, err
	}
	return &a, nil
}

func (r *PVZRepo) UpdateAdm(ctx context.Context, pvz *model.Admin, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE admin SET fio=$1, email=$2, telegram=$3 WHERE id=$4;`, pvz.Fio, pvz.Email, pvz.Telegram, id)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}