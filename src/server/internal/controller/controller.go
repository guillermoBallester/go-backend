package controller

import (
	"context"
	"fmt"

	serverApi "github.com/gsasso/go-backend/src/server/internal/generated/proto"
)

type LogisticCtlr struct {
	serverApi.UnimplementedCoopLogisticsEngineAPIServer
}

func NewLogisticController() *LogisticCtlr {
	return &LogisticCtlr{}
}

func (ctlr *LogisticCtlr) MoveUnit(ctx context.Context, req *serverApi.MoveUnitRequest) (*serverApi.DefaultResponse, error) {
	cargoId := req.GetCargoUnitId()
	fmt.Println("Cargo Unit moved: %s", cargoId)
	resp := &serverApi.DefaultResponse{}
	return resp, nil

}

func (s *LogisticCtlr) UnitReachedWarehouse(ctx context.Context, req *serverApi.UnitReachedWarehouseRequest) (*serverApi.DefaultResponse, error) {
	announcement := req.GetAnnouncement()
	fmt.Println("Announcement: %s", announcement)
	resp := &serverApi.DefaultResponse{}
	return resp, nil
}
