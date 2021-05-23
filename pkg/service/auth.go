package service

import (
	"fmt"
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/iZillu/cryptoWallet/pkg/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

const salt = "akSAhg242jGt38"

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user cryptoWallet.User) (int64, error) {
	var err error

	user.Password, err = createPasswordHash(user.Password)
	if err != nil {
		return 0, errors.Wrap(err, "CreateUser:")
	}

	return s.repo.CreateUser(user)
}

func createPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+salt), 8)
	if err != nil {
		return "", errors.Wrap(err, "createPasswordHash:")
	}

	return fmt.Sprintf("%x", hash), nil
}
