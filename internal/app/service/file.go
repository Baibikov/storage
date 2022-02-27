package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"

	"storage/internal/app/types"
	"storage/pkg/utils"
)

func (u *UseCase) SaveFile(ctx context.Context, uid string, file *multipart.FileHeader) (string, error) {
	folder, err := u.storage.Folder.Get(ctx, uid)
	if err != nil {
		return "", err
	}
	if folder == (types.Folder{}) {
		return "", errors.New("folder not found")
	}

	wg := errgroup.Group{}

	wg.Go(func() error {
		return u.makeFile(file, folder.UID)
	})

	createdUID := ""
	wg.Go(func() error {
		createdUID, err = u.storage.File.Create(ctx, types.File{
			FileName: file.Filename,
			Folder:   folder.UID,
		})
		return errors.Wrap(err, "storage creating file")
	})

	err = wg.Wait()
	if err != nil {
		e := u.removeFile(ctx, u.filePath(uid, file.Filename), createdUID)
		return "", multierr.Append(err, e)
	}

	return createdUID, nil
}

func (u *UseCase) makeFile(file *multipart.FileHeader, folderUID string) error {
	meta, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "opening file")
	}
	defer multierr.AppendInto(
		&err,
		errors.Wrap(meta.Close(), "closing file"),
	)

	fb := make([]byte, file.Size)
	_, err = meta.Read(fb)
	if err != nil {
		return errors.Wrap(err, "reading file")
	}

	dir := u.folderPath(folderUID)

	err = createDirIfNotExists(dir, os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.filePath(folderUID, file.Filename), fb, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "writing file")
	}

	return nil
}

func (u *UseCase) removeFile(ctx context.Context, filePath string, uid string) error {
	err := os.Remove(filePath)
	if err != nil {
		return errors.Wrap(err, "removing file")
	}

	err = u.storage.File.Remove(ctx, uid)
	return errors.Wrap(err, "remove by uid")
}

func createDirIfNotExists(dir string, mode os.FileMode) error {
	if !utils.FolderExists(dir) {
		err := os.Mkdir(dir, mode)
		if err != nil {
			return errors.Wrap(err, "making dir")
		}
	}

	return nil
}

func (u *UseCase) folderPath(uid string) string {
	return fmt.Sprintf("%s/%s", u.config.SRC, uid)
}

func (u *UseCase) filePath(uid string, name string) string {
	return fmt.Sprintf("%s/%s/%s", u.config.SRC, uid, name)
}
