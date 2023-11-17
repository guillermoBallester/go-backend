package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	serverApi "github.com/coopnorge/interview-backend/src/server/generated/proto"
)

type GRPCLogisticServer struct {
	service serverApi.CoopLogisticsEngineAPIClient
	serverApi.UnimplementedCoopLogisticsEngineAPIServer
}

func NewGRPCLogisticServer(service serverApi.CoopLogisticsEngineAPIClient) *GRPCLogisticServer {
	return &GRPCLogisticServer{
		service: service,
	}
}

func (s *GRPCLogisticServer) MoveUnit(ctx context.Context, req *serverApi.MoveUnitRequest) (*serverApi.DefaultResponse, error) {
	resp, err := s.service.MoveUnit(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.CargoUnitId)
	return resp, nil
}

func (s *GRPCLogisticServer) UnitReachedWarehouse(ctx context.Context, req *serverApi.UnitReachedWarehouseRequest) (*serverApi.DefaultResponse, error) {
	resp, err := s.service.UnitReachedWarehouse(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func RunGRPCServer(listenAddr string, service serverApi.CoopLogisticsEngineAPIClient) error {
	GRPCLogisticServer := NewGRPCLogisticServer(service)

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	serverApi.RegisterCoopLogisticsEngineAPIServer(server, GRPCLogisticServer)
	return server.Serve(listener)

}
