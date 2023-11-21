package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	serverApi "github.com/gsasso/go-backend/src/server/internal/generated/proto"
)

type LogisticCtlr struct {
	serverApi.UnimplementedCoopLogisticsEngineAPIServer
}

func NewLogisticController() *LogisticCtlr {
	return &LogisticCtlr{}
}

func (ctlr *LogisticCtlr) MoveUnit(ctx context.Context, req *serverApi.MoveUnitRequest) (*serverApi.DefaultResponse, error) {

	// var wg sync.WaitGroup
	// var ReceiveUnitsCh = make(chan int64)
	// var MakeSummaryUnitsCh = make(chan string)

	// go ReceiveUnits(ReceiveUnitsCh, req.GetCargoUnitId())
	// go MakeSummaryUnits(ReceiveUnitsCh, MakeSummaryUnitsCh)

	// wg.Add(1)
	// go func(MakeSummaryUnitsCh <-chan string) {
	// 	summary := <-MakeSummaryUnitsCh
	// 	fmt.Println(summary)
	// 	wg.Done()
	// }(MakeSummaryUnitsCh)
	// wg.Wait()

	//ticker.ReceiveUnitsCh <- req.GetCargoUnitId()

	resp := &serverApi.DefaultResponse{}
	return resp, nil

}

func (s *LogisticCtlr) UnitReachedWarehouse(ctx context.Context, req *serverApi.UnitReachedWarehouseRequest) (*serverApi.DefaultResponse, error) {
	resp := &serverApi.DefaultResponse{}
	return resp, nil
}

func ReceiveUnits(units chan int64, cargoId int64) chan int64 {
	units <- cargoId
	return units
}

func MakeSummaryUnits(units chan int64, summary chan<- string) {
	var counter int
	time.Sleep(1 * time.Second)
	//close(units)
	for i := range units {
		fmt.Println("Unit ", i, " received")
		counter = counter + 1
	}
	str := strconv.Itoa(counter)
	message := "Total orders received" + str
	summary <- message
}
