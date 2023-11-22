package server

import (
	"net"

	"github.com/google/wire"
	"google.golang.org/grpc"

	controller "github.com/gsasso/go-backend/src/server/internal/controller"
	serverApi "github.com/gsasso/go-backend/src/server/internal/generated/proto"
	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

type LogisticServer struct {
	server *grpc.Server
}

var ServerProvider = wire.NewSet(wire.Struct(new(ticker.Summary)), controller.NewLogisticController, RunGRPCServer)

func RunGRPCServer(ctlr *controller.LogisticCtlr) *LogisticServer {

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	serverApi.RegisterCoopLogisticsEngineAPIServer(server, ctlr)
	//ctlr.svc
	return &LogisticServer{server: server}

}

func (my *LogisticServer) Start() error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	my.server.Serve(listener)

	return nil

}
