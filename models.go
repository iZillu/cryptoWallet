package cryptoWallet

import "time"

type User struct {
	ID         int64  `json:"-"`
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"` // (pass + salt) hashed by bcrypt
	IP         string `json:"ip"`
	Role       string `json:"role"`
	UserAgent  string `json:"userAgent"`
	IsVerified bool   `json:"isVerified"`
}

type Wallet struct {
	ID      int64   `json:"-"`
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID              int64     `json:"-"`
	Amount          float64   `json:"amount"`
	Commission      float64   `json:"commission"`
	CoinName        string    `json:"coinName"`
	SenderAddress   string    `json:"senderAddress"`
	ReceiverAddress string    `json:"receiverAddress"`
	UnixActionTime  int64     `json:"unixActionTime"`
	ActionTime      time.Time `json:"actionTime"`
}
