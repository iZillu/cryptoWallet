package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host, Port, Username, Password, DBName, SSLMode string
}

const (
	usersTable             = "users"
	walletsTable           = "wallets"
	usersWalletsTable      = "users_wallets"
	transactionsTable      = "transactions"
	usersTransactionsTable = "users_transactions"
)

func NewPostgresDB(cnfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cnfg.Host, cnfg.Port, cnfg.Username, cnfg.DBName, cnfg.Password, cnfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
