package repository

import (
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user cryptoWallet.User) (int64, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
