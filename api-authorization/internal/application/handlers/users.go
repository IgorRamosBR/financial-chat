package handlers

import (
	"api-authorization/internal/application/dto"
	"api-authorization/internal/domain/models"
	"api-authorization/internal/domain/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}
func (h UserHandler) CreateUser(c echo.Context) error {
	request, err := bindUserRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := models.ToUser(request)
	err = h.userService.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return err
	}
	return nil
}

func bindUserRequest(c echo.Context) (dto.UserRequest, error) {
	tokenRequest := dto.UserRequest{}
	err := c.Bind(&tokenRequest)
	if err != nil {
		return dto.UserRequest{}, err

	}
	return tokenRequest, nil
}
