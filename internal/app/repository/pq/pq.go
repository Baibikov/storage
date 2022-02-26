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
		    level,
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
		insert into file.folders(name, level) 
		values($1, $2)
		returning uid
	`

	return uid, sqlerr.WithSql(
		f.db.GetContext(ctx, &uid, query, file.Name, file.Level),
		query,
		file.Name,
		file.Level,
	)
}

func (f *File) NameExists(ctx context.Context, name string, level int) (exists bool, err error) {
	query := `
		select exists(
		    select 
		    from file.folders 
		    where 
		          name = $1 and level = $2
		)
	`
	return exists, sqlerr.WithSql(
		f.db.GetContext(ctx, &exists, query, name, level),
		query,
		name,
		level,
	)
}

func (f *File) GetDirectoryByOneLevel(ctx context.Context, uid string, level, before int) (folders []types.Folder, err error) {
	query := `
		with recursive directory_tree as (
			select
				   uid,
				   name,
				   level
			from file.folders
			where
				  uid = $1
			union all
			select
				   fld.uid,
				   fld.name,
				   fld.level
			from file.folders fld
				inner join directory_tree on true
				inner join file.folder_directory fd
					on fd.uid_parent = directory_tree.uid and fd.uid_child = fld.uid
			where fld.level = directory_tree.level+1
			  and fld.level <= $3
		)
		select
			   uid,
			   name,
			   level
		from directory_tree
		where level != $2;
	`
	return folders, sqlerr.WithSql(
		f.db.SelectContext(ctx, &folders, query, uid, level, before),
		query,
		uid,
		level,
		before,
	)
}
