package service

import (
	"storage/configs"
	"storage/internal/app/repository"
)

type UseCase struct {
	config  configs.Storage
	storage *repository.Storage
}

func New(storage *repository.Storage, config configs.Storage) *UseCase {
	return &UseCase{
		config:  config,
		storage: storage,
	}
}
