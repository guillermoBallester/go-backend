package main

import (
	"github.com/google/wire"
	server "github.com/gsasso/go-backend/src/server/internal/server"
)

func InitializeApp() *server.LogisticServer {

	wire.Build(server.ServerProvider)
	return &server.LogisticServer{}
}
