//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/config"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/services/client"
	"github.com/coopnorge/interview-backend/src/client/internal/app/logistics/services/operator"

	"github.com/google/wire"
)

// newWire create new DI
func newWire(cfg *config.ClientAppConfig) (*logistics.ServiceInstance, func(), error) {
	panic(wire.Build(
		client.ServiceSetForClient,
		operator.ServiceSetForOperator,
		logistics.NewServiceInstance,
	))
}
