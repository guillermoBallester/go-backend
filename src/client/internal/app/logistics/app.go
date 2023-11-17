package logistics

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/config"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/model"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/services/client"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/services/operator"
	"github.com/coopnorge/interview-backend/src/client/internal/app/pkg/printer"
	apiv1 "github.com/coopnorge/interview-backend/src/generated/logistics/api/v1"
)

const (
	appName = "Coop Logistics Engine"

	maxWarehouses = 1<<8 - 1
	maxCargoUnits = 1 << 10
)

// ServiceInstance of application
type ServiceInstance struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	logisticsClient *client.APILogisticsClient
	worldOperator   *operator.WorldOperator

	maxMoveWaitNumber int
	reportTable       *printer.ASCIITablePrinter
	statistics        *model.Statistics
}

// NewServiceInstance constructor
func NewServiceInstance(lc *client.APILogisticsClient, wo *operator.WorldOperator, cfg *config.ClientAppConfig) (*ServiceInstance, error) {
	log.Printf("%s, initializing...\n", appName)

	serviceCtx, serviceCtxCancel := context.WithCancel(context.Background())
	connCtx, connCtxCancel := context.WithTimeout(serviceCtx, 30*time.Second)
	defer connCtxCancel()

	log.Printf("%s, trying to connect to API - %s...\n", appName, cfg.GetCombinedAddress())
	if connErr := lc.Connect(cfg.GetCombinedAddress(), connCtx); connErr != nil {
		serviceCtxCancel()
		err := errors.New(fmt.Sprintf(
			"%s, failed to connect to API (%s), error: %v",
			appName,
			cfg.GetCombinedAddress(),
			connErr,
		))

		return nil, err
	}

	service := &ServiceInstance{
		ctx:       serviceCtx,
		ctxCancel: serviceCtxCancel,

		logisticsClient: lc,
		worldOperator:   wo,

		maxMoveWaitNumber: 100,
		reportTable:       printer.NewASCIITablePrinter(),
		statistics: &model.Statistics{
			ExecTime: time.Now(),
			Operation: []*model.Operation{
				{Name: "MoveUnit"},
				{Name: "UnitReachedWarehouse"},
			},
		},
	}

	service.reportTable.AddHeader([]string{"Operation", "Count", "Errors"})
	worldPopulationErr := wo.Populate(
		uint32(rand.Intn(maxWarehouses-10+1)+10),
		uint32(rand.Intn(maxCargoUnits-10+1)+10),
	)
	if worldPopulationErr != nil {
		return nil, worldPopulationErr
	}

	return service, nil
}

// Run app
func (s *ServiceInstance) Run() error {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() { // Handle graceful shutdown
		<-signals // Wait for the signal

		log.Printf("%s, shutting down...\n", appName)

		s.ctxCancel()
		if s.logisticsClient != nil {
			_ = s.logisticsClient.Disconnect()
		}

		log.Printf("%s, stopped!\n", appName)

		os.Exit(0)
	}()

	deliveryUnits := s.worldOperator.GetDeliveryUnit()
	totalDeliveryUnits := len(deliveryUnits)

	for {
		var wg sync.WaitGroup
		unitsReachedObjective := 0

		// Check if all units reached goal
		for _, unit := range deliveryUnits {
			if unit.Metadata == true {
				unitsReachedObjective++
			}
		}

		if unitsReachedObjective == totalDeliveryUnits {
			log.Println("All delivery units reached warehouse...")
			break
		}

		for _, unit := range deliveryUnits {
			if unit.Metadata == true {
				continue
			}

			wg.Add(1)
			go s.processDelivery(unit, &wg)

		}

		wg.Wait()
	}

	for _, o := range s.statistics.Operation {
		s.reportTable.AddRow([]string{
			o.Name,
			strconv.FormatUint(o.A, 10),
			strconv.FormatUint(o.B, 10),
		})
	}

	fmt.Println("\nExecution time:", time.Since(s.statistics.ExecTime))
	fmt.Println(s.reportTable)

	return nil
}

func (s *ServiceInstance) processDelivery(unit *model.GraphNode, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(s.maxMoveWaitNumber) * time.Microsecond)
	s.maxMoveWaitNumber = rand.Intn(s.maxMoveWaitNumber+1) + 1
	if s.maxMoveWaitNumber >= 1 {
		s.maxMoveWaitNumber = s.maxMoveWaitNumber >> 1
	}

	oldCoordinate := *unit.Coordinate
	newCoordinate := s.worldOperator.MoveDeliveryUnitToNearestWarehouse(unit.ID)
	unitMessage := fmt.Sprintf("%s moving to - Latitude:%d, Longitude:%d", unit.Name, newCoordinate.X, newCoordinate.Y)

	log.Println(unitMessage)

	s.statistics.Operation[0].AddA()
	moveErr := s.logisticsClient.MoveUnit(
		s.ctx,
		&apiv1.MoveUnitRequest{
			CargoUnitId: int64(unit.ID),
			Location: &apiv1.Location{
				Latitude:  uint32(newCoordinate.X),
				Longitude: uint32(newCoordinate.Y),
			},
		},
	)
	if moveErr != nil {
		log.Printf("filed to send MoveUnit %s, API error: %v\n", unitMessage, moveErr)
		s.statistics.Operation[0].AddB()

		return
	} else if newCoordinate != oldCoordinate {
		return
	}

	announcement := fmt.Sprintf("%s - Reached Objective.", unitMessage)
	warehouse := s.worldOperator.FindEntityByCoordinate(newCoordinate, model.Warehouses)
	if warehouse == nil {
		log.Printf("Warehouses not found in coordinates Latitude:%d Longitude:%d", newCoordinate.X, newCoordinate.Y)
		return
	}

	s.statistics.Operation[1].AddA()
	reachErr := s.logisticsClient.UnitReachedWarehouse(
		s.ctx,
		&apiv1.UnitReachedWarehouseRequest{
			Location: &apiv1.Location{Latitude: uint32(newCoordinate.X), Longitude: uint32(newCoordinate.Y)},
			Announcement: &apiv1.WarehouseAnnouncement{
				CargoUnitId: int64(unit.ID),
				WarehouseId: int64(warehouse.ID),
				Message:     announcement,
			},
		},
	)
	if reachErr != nil {
		log.Printf("filed to send UnitReachedWarehouse %s, API error: %v\n", unitMessage, moveErr)
		s.statistics.Operation[1].AddB()
		return
	}

	log.Println(announcement)
	unit.Metadata = true // Unit reached Warehouse

	return
}
