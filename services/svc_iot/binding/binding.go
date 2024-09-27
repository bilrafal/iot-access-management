package binding

import (
	"context"
	"fmt"
	"iot-access-management/internal/client"
	"iot-access-management/internal/config"
	"iot-access-management/internal/db"
	"iot-access-management/internal/repo/repo_iot_simple"
	"iot-access-management/internal/router"
	"iot-access-management/services/svc_iot/internal/api/handler"
	"iot-access-management/services/svc_iot/internal/core_manager/iot_implementation"
	"path/filepath"
)

type IoTManagerBinder struct {
}

// GetConfig method loads the http server and app config from the yaml file
func (b *IoTManagerBinder) GetConfig() config.Config {
	configPath := filepath.Join("services", "svc_iot", "config")
	return config.LoadConfig(configPath)
}

// BindDependencies method binds the necessary elements (i.e. repos, http routes) based on the app config
func (b *IoTManagerBinder) BindDependencies(ctx context.Context, routesDef []router.RouteDef) []router.RouteDef {
	//Get app config from file
	cfg := b.GetConfig()

	//Fill all necessary elements: repo, core functionality manager and handler manager
	repo := repo_iot_simple.NewRepoIotSimple(ctx, db.DbType(cfg.DbDef.DbType))
	c := client.NewClient(cfg.ClientConfig.ClientHost, cfg.ClientConfig.ClientPort, cfg.ClientConfig.Timeout)
	manager := iot_implementation.NewCredentialManagerSimple(repo, *c)
	handlerManager := handler.NewIoTHandler(manager)

	handlers := make([]router.RouteDef, 0)
	for _, def := range routesDef {
		switch fmt.Sprintf("%p", def.Handler) {
		case fmt.Sprintf("%p", handlerManager.CreateWhiteList):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.CreateWhiteList))
		case fmt.Sprintf("%p", handlerManager.DeleteWhiteList):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.DeleteWhiteList))
		case fmt.Sprintf("%p", handlerManager.RequestAccess):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.RequestAccess))
		case fmt.Sprintf("%p", handlerManager.ListWhiteList):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.ListWhiteList))
		default:
			panic(fmt.Sprintf("Unsupported route definition type: %p", def.Handler))
		}
	}

	return handlers
}
