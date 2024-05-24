package postgres

import (
	"context"
	"devtask/internal/model"
	"devtask/internal/pkg/db"
	"errors"
	"fmt"
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
	err := r.db.Get(ctx, &a, `SELECT id,name,owner,vms,cpu,ram,hdd,software_used,admin_name,admin_email,admin_tg,resource_assignment,status FROM infsys where id=$1`, id)
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
	err := r.db.ExecQueryRow(ctx, `INSERT INTO infsys(id,name,owner,vms,cpu,ram,hdd,software_used,admin_name,admin_email,admin_tg,resource_assignment,status) VALUES ($1,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id;`, pvz.ID, pvz.Name, pvz.Owner, pvz.Vms, pvz.Cpu, pvz.Ram, pvz.Hdd, pvz.SoftwareUsed, pvz.AdminName, pvz.AdminEmail, pvz.AdminTg, pvz.ResourceAssignment, pvz.Status).Scan(&id)
	return id, err
}

func (r *PVZRepo) UpdateISInfo(ctx context.Context, pvz *model.TableInfSystems, id int64) (int64, error) {
	cT, err := r.db.Exec(ctx, `UPDATE infsys SET name=$1, owner=$2, vms=$3, cpu=$4, ram=$5, hdd=$6, software_used=$7, admin_name=$8, admin_email=$9, admin_tg=$10, resource_assignment=$11, status=$12 WHERE id=$13;`, pvz.Name, pvz.Owner, pvz.Vms, pvz.Cpu, pvz.Ram, pvz.Hdd, pvz.SoftwareUsed, pvz.AdminName, pvz.AdminEmail, pvz.AdminTg, pvz.ResourceAssignment, pvz.Status, id)
	fmt.Println(cT, err)
	if cT.RowsAffected() == 0 {
		return 0, model.ErrNoRowsInResultSet
	}
	return id, err
}
