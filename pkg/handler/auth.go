package handler

import (
	cryptoWallet "github.com/iZillu/cryptoWallet"
	"github.com/labstack/echo"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	var input cryptoWallet.User

	if err := c.Bind(&input); err != nil {
		log.Error().Err(err).Msg("signUp: binding user")
		return err
	}
	input.IP = getIP(c.Request())
	input.UserAgent = c.Request().UserAgent()

	id, err := h.Service.Authorization.CreateUser(input)
	if err != nil {
		log.Error().Err(err).Msg("signUp:")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c echo.Context) error {

	return nil
}

func (h *Handler) verifyEmail(c echo.Context) error {

	return nil
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
