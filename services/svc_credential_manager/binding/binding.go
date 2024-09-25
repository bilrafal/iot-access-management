package binding

import (
	"context"
	"fmt"
	"iot-access-management/internal/config"
	"iot-access-management/internal/db"
	"iot-access-management/internal/repo/repo_credential_simple"
	"iot-access-management/internal/router"
	"iot-access-management/internal/util"
	"iot-access-management/services/svc_credential_manager/internal/api/handler"
	"iot-access-management/services/svc_credential_manager/internal/core_manager/manager_implementation"
	"path/filepath"
)

type CredentialManagerBinder struct {
}

func (b *CredentialManagerBinder) GetConfig() config.Config {
	configPath := filepath.Join(util.GetEffectiveUserHomeFolder(), "services", "svc_credential_manager", "config")
	return config.LoadConfig(configPath)
}
func (b *CredentialManagerBinder) BindDependencies(ctx context.Context, routesDef []router.RouteDef) []router.RouteDef {
	cfg := b.GetConfig()
	handlers := make([]router.RouteDef, 0)
	repo := repo_credential_simple.NewRepoCredentialSimple(ctx, db.DbType(cfg.DbDef.DbType))
	manager := manager_implementation.NewCredentialManagerSimple(repo)
	handlerManager := handler.NewCredentialHandler(manager)

	for _, def := range routesDef {
		switch fmt.Sprintf("%p", def.Handler) {
		case fmt.Sprintf("%p", handlerManager.CreateUser):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.CreateUser))
		case fmt.Sprintf("%p", handlerManager.GetUser):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.GetUser))
		case fmt.Sprintf("%p", handlerManager.CreateCredential):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.CreateCredential))
		default:
			panic(fmt.Sprintf("Unsupported route definition type: %p", def.Handler))
		}
	}

	return handlers
}
