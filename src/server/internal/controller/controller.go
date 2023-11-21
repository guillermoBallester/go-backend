package controller

import (
	"context"

	serverApi "github.com/gsasso/go-backend/src/server/internal/generated/proto"
	"github.com/gsasso/go-backend/src/server/internal/ticker"
)

type LogisticCtlr struct {
	serverApi.UnimplementedCoopLogisticsEngineAPIServer
	summary ticker.Summary
}

func NewLogisticController(summary ticker.Summary) *LogisticCtlr {
	return &LogisticCtlr{
		summary: summary,
	}
}

func (ctlr *LogisticCtlr) MoveUnit(ctx context.Context, req *serverApi.MoveUnitRequest) (*serverApi.DefaultResponse, error) {
	ctlr.summary.IncreaseTotalUnits()
	resp := &serverApi.DefaultResponse{}
	return resp, nil
}

func (ctlr *LogisticCtlr) UnitReachedWarehouse(ctx context.Context, req *serverApi.UnitReachedWarehouseRequest) (*serverApi.DefaultResponse, error) {
	ctlr.summary.IncreaseTotalReached()
	resp := &serverApi.DefaultResponse{}
	return resp, nil
}
