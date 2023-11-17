package server

import (
	"net"

	"google.golang.org/grpc"

	controller "github.com/gsasso/go-backend/src/server/internal/controller"
	serverApi "github.com/gsasso/go-backend/src/server/internal/generated/proto"
)

func RunGRPCServer(listenAddr string, ctlr *controller.LogisticCtlr) error {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	serverApi.RegisterCoopLogisticsEngineAPIServer(server, ctlr)

	return server.Serve(listener)

}
