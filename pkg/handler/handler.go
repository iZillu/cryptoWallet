package handler

import (
	"github.com/iZillu/cryptoWallet/pkg/service"
	"github.com/labstack/echo"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *echo.Echo {
	server := echo.New()

	auth := server.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/verify-email", h.verifyEmail)
	}

	api := server.Group("/api", h.UserIdentity)
	{
		api.GET("/get-wallets", h.getWallets)
		api.POST("/make-transaction", h.makeTransaction)
		api.GET("/get-transactions", h.getTransactions)
	}

	return server
}
