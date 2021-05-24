package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) getWallets(c echo.Context) error {
	id := c.Get("userID")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userID": id,
	})
}

func (h *Handler) makeTransaction(c echo.Context) error {
	return nil
}

func (h *Handler) getTransactions(c echo.Context) error {
	return nil
}
