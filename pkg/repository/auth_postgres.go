package repository

import (
	"fmt"
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user cryptoWallet.User) (int64, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (name, surname, email, password, ip, role, userAgent, isVerified) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Password, user.IP, user.Role, user.UserAgent, user.IsVerified)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
