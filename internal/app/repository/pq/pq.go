package pq

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

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

func (f *File) Get(ctx context.Context, uid string) (file types.Folder, err error) {
	query := `
		select 
		    uid,
		    name,
		    created_at
		from file.folders
		where uid = $1
	`
	err = f.db.GetContext(ctx, &file, query, uid)
	if errors.Is(err, sql.ErrNoRows) {
		return file, nil
	}

	return file, sqlerr.WithSql(
		err,
		query,
		uid,
	)
}


func (f *File) Create(ctx context.Context, file types.Folder) (uid string, err error) {
	query := `
		insert into file.folders(name, parent) 
		values($1, $2)
		returning uid
	`

	return uid, sqlerr.WithSql(
		f.db.GetContext(ctx, &uid, query, file.Name, file.Parent),
		query,
		file.Name,
		file.Parent,
	)
}

func (f *File) NameExists(ctx context.Context, name string, parent int) (exists bool, err error) {
	query := `
		select exists(
		    select 
		    from file.folders 
		    where 
		          name = $1 and parent = $2
		)
	`
	return exists, sqlerr.WithSql(
		f.db.GetContext(ctx, &exists, query, name, parent),
		query,
		name,
		parent,
	)
}