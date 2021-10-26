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
//	defer clients.CloseMongoClient(ctx)

	roomRepository := repositories.NewRoomRepository(mongoClient)

	roomService := services.NewRoomService(roomRepository)

	router := routes.NewRouter(roomService)

	err := router.Run(":1323")
	if err != nil {
		panic(err)
	}
}
