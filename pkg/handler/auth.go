package handler

import (
	"github.com/iZillu/cryptoWallet"
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

// TODO: add ban for 3 minutes
func (h *Handler) signIn(c echo.Context) error {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&input); err != nil {
		log.Error().Err(err).Msg("signUp: binding user:")
		return err
	}

	token, err := h.Service.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		log.Error().Err(err).Msg("signUp:")
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
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
