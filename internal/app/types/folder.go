package types

import "time"

type Folder struct {
	UID 		string 		`db:"uid"`
	Name 		string 		`db:"name"`
	Parent      int 		`db:"parent"`
	CreatedAt 	time.Time 	`db:"createdAt"`
}
