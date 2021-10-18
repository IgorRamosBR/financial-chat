package handlers

import (
	"api-authorization/internal/application/dto"
	"api-authorization/internal/domain/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type OAuthHandler struct {
	oauthService services.OAuthService
}

func NewOAuthHandler(oauthService services.OAuthService) OAuthHandler {
	return OAuthHandler{
		oauthService: oauthService,
	}
}

func (h OAuthHandler) AuthenticateUser(c echo.Context) error {
	auth, err := bindRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = auth.IsValid()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	h.oauthService.Authorize(auth)

	return nil
}

func bindRequest(c echo.Context) (dto.Auth, error) {
	tokenRequest := dto.Auth{}
	err := c.Bind(&tokenRequest)
	if err != nil {
		return dto.Auth{}, err

	}
	return tokenRequest, nil
}
