package main

import (
	"github.com/gsasso/go-backend/src/server/internal/controller"
	server "github.com/gsasso/go-backend/src/server/internal/server"
	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

func main() {
	logisticCtlr := controller.NewLogisticController(&ticker.SummaryService{})
	logisticServer := server.RunGRPCServer(logisticCtlr)
	logisticServer.Start()
}
