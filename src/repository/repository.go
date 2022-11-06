package repository

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

const keyDB = "gorm_db"

type Repository struct {
	Db *gorm.DB
}

func (r *Repository) GetDB(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyDB).(*gorm.DB)
	if !ok {
		return r.Db
	}
	return val
}

func (r *Repository) BeginTransaction(ctx context.Context, opts ...*sql.TxOptions) context.Context {
	db := r.GetDB(ctx)
	return context.WithValue(ctx, keyDB, db.Begin(opts...))
}

func (r *Repository) Commit(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Commit()
}

func (r *Repository) Rollback(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Rollback()
}
