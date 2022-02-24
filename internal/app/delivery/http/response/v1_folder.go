package response

import "time"

type V1FolderGet struct {
	Name string `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
}

type V1FolderPost struct {
	UID string `json:"uid"`
}