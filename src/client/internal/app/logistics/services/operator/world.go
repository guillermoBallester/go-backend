package operator

import (
	"errors"
	"math"
	"math/rand"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/model"
	"github.com/coopnorge/interview-backend/src/client/internal/app/pkg/generator"
	"github.com/google/wire"
)

// ServiceSetForOperator providers
var ServiceSetForOperator = wire.NewSet(NewWorldOperator)

// WorldOperator that handles world and units movements
type WorldOperator struct {
	world *model.Graph
}

// NewWorldOperator instance
func NewWorldOperator() *WorldOperator {
	return &WorldOperator{
		world: model.NewGraph(),
	}
}

func (wo *WorldOperator) Populate(maxWarehouses, maxCargoUnits uint32) error {
	if uint64(maxWarehouses)+uint64(maxCargoUnits) >= 1<<32-1 {
		return errors.New("world actor count overflow")
	}

	generator.AddNewActors(model.Warehouses, wo.world, uint(maxWarehouses), 0)
	generator.AddNewActors(model.CargoUnits, wo.world, uint(maxCargoUnits), uint(maxWarehouses))

	var warehouseIDs []uint
	var deliveryUnitIDs []uint
	for _, node := range wo.world.Nodes {
		if node.Type == model.Warehouses {
			warehouseIDs = append(warehouseIDs, node.ID)
		} else if node.Type == model.CargoUnits {
			deliveryUnitIDs = append(deliveryUnitIDs, node.ID)
		}
	}

	for _, warehouseID := range warehouseIDs {
		if len(deliveryUnitIDs) == 0 {
			break
		}

		rand.Shuffle(len(deliveryUnitIDs), func(i, j int) {
			deliveryUnitIDs[i], deliveryUnitIDs[j] = deliveryUnitIDs[j], deliveryUnitIDs[i]
		})

		numDeliveryUnits := rand.Intn(len(deliveryUnitIDs)) + 1 // Random number of units to connect (at least 1)
		for i := 0; i < numDeliveryUnits; i++ {
			unitID := deliveryUnitIDs[i]

			wo.world.AddEdge(model.GraphEdge{Source: unitID, Target: warehouseID})
		}

		deliveryUnitIDs = deliveryUnitIDs[numDeliveryUnits:]
	}

	return nil
}

// GetDeliveryUnit from the world
func (wo *WorldOperator) GetDeliveryUnit() []*model.GraphNode {
	return wo.world.GetNodesByType(model.CargoUnits)
}

// FindEntityByCoordinate in the world
func (wo *WorldOperator) FindEntityByCoordinate(coordinate model.Coordinate, entityType model.ActorType) *model.GraphNode {
	return wo.world.FindNodesByLocation(coordinate, entityType)
}

// MoveDeliveryUnitToNearestWarehouse moves the given unit to the nearest connected warehouse based on their X and Y locations
func (wo *WorldOperator) MoveDeliveryUnitToNearestWarehouse(unitID uint) model.Coordinate {
	deliveryUnitNode := wo.world.GetNodeByID(unitID)
	unitX := deliveryUnitNode.X
	unitY := deliveryUnitNode.Y

	connectedWarehouses := wo.world.GetConnectedNodes(unitID, model.Warehouses)

	// Initialize variables for tracking the nearest warehouse
	minDistance := math.MaxFloat64
	nearestWarehouseID := uint(0)

	for _, warehouseNode := range connectedWarehouses {
		warehouseX := warehouseNode.X
		warehouseY := warehouseNode.Y

		distance := math.Sqrt(math.Pow(float64(unitX-warehouseX), 2) + math.Pow(float64(unitY-warehouseY), 2))

		// Update nearest warehouse if distance is smaller
		if distance < minDistance {
			minDistance = distance
			nearestWarehouseID = warehouseNode.ID
		}
	}

	// Move unit to goal
	if unitX < wo.world.GetNodeByID(nearestWarehouseID).X {
		deliveryUnitNode.X++
	} else if unitX > wo.world.GetNodeByID(nearestWarehouseID).X {
		deliveryUnitNode.X--
	}
	if unitY < wo.world.GetNodeByID(nearestWarehouseID).Y {
		deliveryUnitNode.Y++
	} else if unitY > wo.world.GetNodeByID(nearestWarehouseID).Y {
		deliveryUnitNode.Y--
	}

	return model.Coordinate{X: deliveryUnitNode.X, Y: deliveryUnitNode.Y}
}
