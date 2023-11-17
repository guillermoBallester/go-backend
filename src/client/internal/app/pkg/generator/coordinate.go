package generator

import (
	"math/rand"

	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/model"
)

// NewCoordinates with unique placement
func NewCoordinates(numCoordinates, xRange, yRange int) []model.Coordinate {
	coordinates := make([]model.Coordinate, numCoordinates)

	for i := 0; i < numCoordinates; i++ {
		x := rand.Intn(xRange)
		y := rand.Intn(yRange)

		coordinates[i] = model.Coordinate{X: x, Y: y}
	}

	return coordinates
}
