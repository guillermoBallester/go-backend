package main

import (
	"github.com/gsasso/go-backend/src/server/internal/controller"
	server "github.com/gsasso/go-backend/src/server/internal/server"
	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

func main() {

	//s := ticker.Summary{}
	//go s.Tick()
	summary := ticker.Summary{}
	logisticCtlr := controller.NewLogisticController(summary)
	logisticServer := server.RunGRPCServer(logisticCtlr)
	go summary.Tick()
	logisticServer.Start()

	//app := InitializeApp()
	//if app != nil {
	//	app.Start()
	//}

}
