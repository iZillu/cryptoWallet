package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/iZillu/cryptoWallet/pkg/repository"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	repo repository.Authorization
}

const (
	salt       = "akSAhg242jGt38"
	signingKey = "wkogj2pio3589p938aksodp"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
}

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

func (s *AuthService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.Wrap(errors.New("invalid signing method"), "parseToken:")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.Wrap(errors.New("token.Claims is not of type *tokenClaims"), "parseToken:")
	}

	return claims.UserID, nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	passHash, err := createPasswordHash(password)
	if err != nil {
		return "", errors.Wrap(err, "GenerateToken:")
	}

	user, err := s.repo.GetUser(email, passHash)
	if err != nil {
		return "", errors.Wrap(err, "GenerateToken:")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func createPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+salt), 8)
	if err != nil {
		return "", errors.Wrap(err, "createPasswordHash:")
	}

	return fmt.Sprintf("%x", hash), nil
}
