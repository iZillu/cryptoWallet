package service

import (
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/iZillu/cryptoWallet/pkg/repository"
)

type Authorization interface {
	CreateUser(user cryptoWallet.User) (int64, error)
}

type Service struct {
	Authorization
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Authorization),
	}
}
