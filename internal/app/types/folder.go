package types

import "time"

type Folder struct {
	UID       string    `db:"uid"`
	Name      string    `db:"name"`
	Level     int       `db:"level"`
	CreatedAt time.Time `db:"created_at"`
}
