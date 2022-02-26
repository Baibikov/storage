package pq

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"

	"storage/internal/app/types"
	"storage/pkg/dbutils"
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
				  uid = $1::uuid
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

func (f *File) GetDirectoryUids(ctx context.Context, uid string) (uids []string, err error) {
	query := `
		with recursive directory_tree as (
			select
				uid,
				level
			from file.folders
			where uid = $1::uuid
			union all
			select
				fld.uid,
				fld.level
			from file.folders fld
					 inner join directory_tree on true
					 inner join file.folder_directory fd
								on fd.uid_parent = directory_tree.uid and fd.uid_child = fld.uid
			where fld.level = directory_tree.level+1
		)
		select array_agg(uid) from directory_tree
	`
	var uidStringArray pq.StringArray
	err = sqlerr.WithSql(f.db.GetContext(ctx, &uidStringArray, query, uid), query, uid)
	if err != nil {
		return nil, err
	}

	uids = uidStringArray
	return uids, nil
}

func (f *File) DeleteDirectory(ctx context.Context, uids []string) (err error) {
	return dbutils.WrapTx(f.db, func(tx *sqlx.Tx) error {
		err = deleteFromFolderDirectory(ctx, tx, uids)
		if err != nil {
			return err
		}
		return deleteFromFoldersTx(ctx, tx, uids)
	})
}

func deleteFromFolderDirectory(ctx context.Context, tx sqlx.ExtContext, uids []string) error {
	query := `
		delete from file.folder_directory where uid_parent = any($1)
	`
	_, err := tx.ExecContext(ctx, query, pq.Array(uids))
	return sqlerr.WithSql(
		err,
		query,
		uids,
	)
}

func deleteFromFoldersTx(ctx context.Context, tx sqlx.ExtContext, uids []string) error {
	query := `
		delete from file.folders where uid = any($1)
	`
	_, err := tx.ExecContext(ctx, query, pq.Array(uids))
	return sqlerr.WithSql(
		err,
		query,
		uids,
	)
}
