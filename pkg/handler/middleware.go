package handler

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"strings"
)

func (h *Handler) UserIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if err := errors.New("empty auth header"); header == "" {
			log.Error().Err(err).Msg("userIdentity:")
			return err
		}

		headerParts := strings.Split(header, " ")
		if err := errors.New("invalid auth header"); len(headerParts) != 2 { // rы
			log.Error().Err(err).Msg("userIdentity:")
			return err
		}

		userID, err := h.Service.Authorization.ParseToken(headerParts[1])
		if err != nil {
			log.Error().Err(err)
			return err
		}

		c.Set("userID", userID)
		return next(c)
	}
}
