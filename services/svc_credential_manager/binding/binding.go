package binding

import (
	"context"
	"fmt"
	"iot-access-management/internal/client"
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

// GetConfig method loads the http server and app config from the yaml file
func (b *CredentialManagerBinder) GetConfig() config.Config {
	configPath := filepath.Join(util.GetEffectiveUserHomeFolder(), "services", "svc_credential_manager", "config")
	return config.LoadConfig(configPath)
}

// BindDependencies method binds the necessary elements (i.e. repos, http routes) based on the app config
func (b *CredentialManagerBinder) BindDependencies(ctx context.Context, routesDef []router.RouteDef) []router.RouteDef {
	//Get app config from file
	cfg := b.GetConfig()

	//Fill all necessary elements: repo, core functionality manager and handler manager
	repo := repo_credential_simple.NewRepoCredentialSimple(ctx, db.DbType(cfg.DbDef.DbType))
	c := client.NewClient(cfg.ClientConfig.ClientHost, cfg.ClientConfig.ClientPort, cfg.ClientConfig.Timeout)
	manager := manager_implementation.NewCredentialManagerSimple(repo, *c)

	handlerManager := handler.NewCredentialHandler(manager)

	handlers := make([]router.RouteDef, 0)
	for _, def := range routesDef {
		switch fmt.Sprintf("%p", def.Handler) {
		case fmt.Sprintf("%p", handlerManager.CreateUser):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.CreateUser))
		case fmt.Sprintf("%p", handlerManager.GetUser):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.GetUser))
		case fmt.Sprintf("%p", handlerManager.CreateCredential):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.CreateCredential))
		case fmt.Sprintf("%p", handlerManager.AssignCredentialToUser):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.AssignCredentialToUser))
		case fmt.Sprintf("%p", handlerManager.GetUserCredentials):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.GetUserCredentials))
		case fmt.Sprintf("%p", handlerManager.AuthorizeUserOnDoor):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.AuthorizeUserOnDoor))
		case fmt.Sprintf("%p", handlerManager.RevokeAuthorization):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.RevokeAuthorization))
		case fmt.Sprintf("%p", handlerManager.GetCredential):
			handlers = append(handlers, router.NewRouteDef(def.HttpMethod, def.Pattern, handlerManager.GetCredential))
		default:
			panic(fmt.Sprintf("Unsupported route definition type: %p", def.Handler))
		}
	}

	return handlers
}
