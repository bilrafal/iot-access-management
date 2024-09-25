package binding

import (
	"context"
	"fmt"
	"iot-access-management/internal/app"
	"iot-access-management/internal/config"
	"iot-access-management/internal/db"
	"iot-access-management/internal/repo/repo_credential_simple"
	"iot-access-management/services/svc_credential_manager/internal/api/handler"
	"iot-access-management/services/svc_credential_manager/internal/core_manager/manager_implementation"
	"net/http"
)

func BindDependencies(ctx context.Context, cfgPath string, routesDef []app.RouteDef) ([]app.RouteDef, config.Config) {
	cfg := config.LoadConfig(cfgPath)
	handlers := make([]app.RouteDef, 0)
	repo := repo_credential_simple.NewRepoCredentialSimple(ctx, db.DbType(cfg.DbDef.DbType))
	manager := manager_implementation.NewCredentialManagerSimple(repo)
	handlerManager := handler.NewCredentialHandler(manager)

	for _, def := range routesDef {
		switch fmt.Sprintf("%p", def.Handler) {
		case fmt.Sprintf("%p", handlerManager.CreateUser):
			handlers = append(handlers, app.NewRouteDef(http.MethodPost, "create-user", handlerManager.CreateUser))
		case fmt.Sprintf("%p", handlerManager.GetUser):
			handlers = append(handlers, app.NewRouteDef(http.MethodGet, "get-user/{id}", handlerManager.GetUser))
		default:
			panic(fmt.Sprintf("Unsupported route definition type: %p", def.Handler))
		}
	}

	return handlers, cfg
}
