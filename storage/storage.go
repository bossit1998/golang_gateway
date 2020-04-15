package storage

import (
	"github.com/jmoiron/sqlx"

	"bitbucket.org/alien_soft/api_gateway/storage/postgres"
	"bitbucket.org/alien_soft/api_gateway/storage/repo"
)

// Sql Storage ...
type StorageI interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
