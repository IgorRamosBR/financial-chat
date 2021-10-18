package main

import (
	"api-chat/internal/application/routes"
	"api-chat/internal/domain/services"
	"api-chat/internal/infra/clients"
	"api-chat/internal/infra/repositories"
	"context"
)

func main() {
	ctx := context.Background()

	mongoClient := clients.NewMongoClient(ctx)
	defer clients.CloseMongoClient(ctx)

	roomRepository := repositories.NewRoomRepository(mongoClient)

	roomService := services.NewUserService(roomRepository)

	router := routes.NewRouter(roomService)

	router.Logger.Fatal(router.Start(":1323"))
}
