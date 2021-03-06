package service

import (
	"context"

	"github.com/pkg/errors"

	"storage/internal/app/types"
)

const (
	folderNameLength = 3
	copyName         = "(copy)"
)

func (u *UseCase) CreateFolder(ctx context.Context, folder types.Folder) (uid string, err error) {
	if len(folder.Name) <= folderNameLength {
		return "", errors.Errorf(
			"file name: %s not allowed. Length must be more than %d characters",
			folder.Name,
			folderNameLength,
		)
	}

	exist, err := u.storage.Folder.NameExists(ctx, folder.Name, folder.Level)
	if err != nil {
		return "", errors.Wrap(
			err,
			"existing name",
		)
	}

	if exist {
		folder.Name += " " + copyName
	}

	uid, err = u.storage.Folder.Create(ctx, folder)
	if err != nil {
		return "", errors.Wrap(
			err,
			"creating folder",
		)
	}

	return uid, nil
}

func (u *UseCase) GetFolder(ctx context.Context, uid string) (folder types.Folder, err error) {
	folder, err = u.storage.Folder.Get(ctx, uid)
	if err != nil {
		return folder, errors.Wrap(
			err,
			"getting folder",
		)
	}

	if folder == (types.Folder{}) {
		return folder, errors.New("folder unknown")
	}

	return folder, nil
}

func (u *UseCase) GetFolderDirectory(ctx context.Context, uid string, level int) (
	folders []types.Folder,
	err error,
) {
	return u.storage.Folder.GetDirectoryByOneLevel(ctx, uid, level, level+1)
}

func (u *UseCase) DeleteFolder(ctx context.Context, uid string) error {
	uids, err := u.storage.Folder.GetDirectoryUids(ctx, uid)
	if err != nil {
		return err
	}

	return u.storage.Folder.DeleteDirectory(ctx, uids)
}
