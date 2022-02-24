package service

import (
	"context"

	"github.com/pkg/errors"

	"storage/internal/app/repository"
	"storage/internal/app/types"
)

type UseCase struct {
	storage *repository.Storage
}

func New(storage *repository.Storage) *UseCase {
	return &UseCase{
		storage: storage,
	}
}

const (
	folderNameLength = 3
	copyName = "(copy)"
)

func (u *UseCase) CreateFolder(ctx context.Context, folder types.Folder) (uid string, err error) {
	if len(folder.Name) <= folderNameLength {
		return "", errors.Errorf(
			"file name: %s not allowed. Length must be more than %d characters",
			folder.Name,
			folderNameLength,
		)
	}

	exist, err := u.storage.Folder.NameExists(ctx, folder.Name, folder.Parent)
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