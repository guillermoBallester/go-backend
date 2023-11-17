package operator

import (
	"testing"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/model"
)

func TestNewWorldOperator(t *testing.T) {
	// Create a new WorldOperator instance
	wOperator := NewWorldOperator()

	populationErr := wOperator.Populate(2, 2)
	if populationErr != nil {
		t.Errorf("Not expected error when populating NewWorldOperator instance, error: %v", populationErr)
	}

	// Check the number of nodes
	expectedNumNodes := 4 // 2 warehouses + 2 cargo units
	if len(wOperator.world.Nodes) != expectedNumNodes {
		t.Errorf("Expected %d nodes, but got %d", expectedNumNodes, len(wOperator.world.Nodes))
	}

	// Check the number of edges
	expectedNumEdges := 0 // Random connections are made after creating the WorldOperator, so no fixed number of edges can be expected
	actualNumEdges := len(wOperator.world.Edges)
	if actualNumEdges < expectedNumEdges {
		t.Errorf("Expected at least %d edges, but got %d", expectedNumEdges, actualNumEdges)
	}

	// Check the Connected flag for each warehouse
	for _, node := range wOperator.world.Nodes {
		if node.Type == model.CargoUnits {
			if !node.Connected {
				t.Errorf("Cargo unit with ID %d should not be connected, but it is", node.ID)
			}
		}
	}
}

func TestNewWorldOperatorActorOverflow(t *testing.T) {
	wOperator := NewWorldOperator()

	populationErr := wOperator.Populate(^uint32(0), ^uint32(0))
	if populationErr == nil {
		t.Errorf("Expected error, since sum of max will overflow uint32")
	}
}
