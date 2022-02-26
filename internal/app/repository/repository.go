package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"storage/internal/app/repository/pq"
	"storage/internal/app/types"
)

type Folder interface {
	Get(ctx context.Context, uid string) (types.Folder, error)
	Create(ctx context.Context, file types.Folder) (uid string, err error)
	NameExists(ctx context.Context, name string, level int) (exists bool, err error)
	GetDirectoryByOneLevel(ctx context.Context, uid string, level, before int) (folders []types.Folder, err error)
	GetDirectoryUids(ctx context.Context, uid string) (uids []string, err error)
	DeleteDirectory(ctx context.Context, uids []string) (err error)
}

type Storage struct {
	Folder Folder
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Folder: pq.NewFile(db),
	}
}
