package types

import "time"

type File struct {
	UID      string    `db:"uid"`
	FileName string    `db:"file_name"`
	Folder   string    `db:"folder"`
	AddedAt  time.Time `db:"added_at"`
}
