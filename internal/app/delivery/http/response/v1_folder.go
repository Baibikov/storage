package response

import "time"

type V1FolderGet struct {
	UID         string    `json:"uid"`
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	CreatedDate time.Time `json:"createdDate"`
}

type V1FolderPost struct {
	UID string `json:"uid"`
}

type V1FolderDirectoryGet []V1FolderGet
