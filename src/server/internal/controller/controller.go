package controller

import (
	"context"
	"fmt"

	"demo/go/pkg/mod/google.golang.org/protobuf@v1.31.0/runtime/protoimpl"

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
	fmt.Printf("Cargo Unit moved: %s", cargoId)
	return protoimpl.X.MessageStringOf(cargoId), nil

}

func (s *LogisticCtlr) UnitReachedWarehouse(ctx context.Context, req *serverApi.UnitReachedWarehouseRequest) (*serverApi.DefaultResponse, error) {
	announcement := req.GetAnnouncement()
	fmt.Printf("Announcement: %s", announcement)
	return protoimpl.X.MessageStringOf(announcement), nil
}
