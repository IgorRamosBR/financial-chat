package main

import (
	"api-authorization/internal/application/routes"
	"api-authorization/internal/domain/services"
	"api-authorization/internal/infra/clients"
	"api-authorization/internal/infra/repositories"
	"log"
)

func main() {
	mongoClient := clients.NewMongoClient()
	defer clients.CloseMongoClient()

	userRepository := repositories.NewUserRepository(mongoClient)

	userService := services.NewUserService(userRepository)
	oauthService := services.NewOAuthService(userService)

	router := routes.Route(userService, oauthService)

	if err := router.Start(":8080"); err != nil {
		log.Fatalln("Error to starting application", err)
	}
}
