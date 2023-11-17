package main

import (
	"fmt"
	"log"

	controller "github.com/gsasso/go-backend/src/server/internal/controller"
	server "github.com/gsasso/go-backend/src/server/internal/server"
)

func main() {
	fmt.Println("Starting server")

	LogicticController := controller.NewLogisticController()

	err := server.RunGRPCServer(":50051", LogicticController)
	if err != nil {
		log.Panic("Server run failure")
	}

}
