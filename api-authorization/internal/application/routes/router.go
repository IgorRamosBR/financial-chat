package routes

import (
	"api-authorization/internal/application/handlers"
	"api-authorization/internal/domain/services"
	"github.com/labstack/echo/v4"
)

func Route(userService services.UserService, oauthService services.OAuthService) *echo.Echo {
	router := echo.New()

	userHandler := handlers.NewUserHandler(userService)
	oauthHandler := handlers.NewOAuthHandler(oauthService)

	router.GET("/health", handlers.Health)
	router.POST("/oauth/token", oauthHandler.AuthenticateUser)
	router.POST("/users", userHandler.CreateUser)

	return router
}
