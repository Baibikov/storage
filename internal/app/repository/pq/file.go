package pq

import (
	"context"

	"github.com/jmoiron/sqlx"

	"storage/internal/app/types"
	"storage/pkg/sqlerr"
)

type File struct {
	db *sqlx.DB
}

func NewFile(db *sqlx.DB) *File {
	return &File{
		db: db,
	}
}

func (f File) Create(ctx context.Context, file types.File) (uid string, err error) {
	query := `
		insert into file.files(file_name, folder) values ($1, $2)
		returning uid
	`
	return uid, sqlerr.WithSql(
		f.db.GetContext(ctx, &uid, query, file.FileName, file.Folder),
		query,
		file.FileName,
		file.Folder,
	)
}

func (f File) Remove(ctx context.Context, uid string) (err error) {
	query := `
		delete from file.files where uid = $1
	`
	_, err = f.db.ExecContext(ctx, query, uid)
	return sqlerr.WithSql(err, query, uid)
}
