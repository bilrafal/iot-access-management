package binder

import (
	"context"
	"iot-access-management/internal/config"
	"iot-access-management/internal/router"
)

type Binder interface {
	BindDependencies(ctx context.Context, routesDef []router.RouteDef) []router.RouteDef
	GetConfig() config.Config
}
