package main

import (
	"context"
	"fmt"
	"iot-access-management/internal/app"
	"iot-access-management/internal/util"
	"iot-access-management/services/svc_credential_manager/binding"
	"iot-access-management/services/svc_credential_manager/internal/api/route"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
	configPath := filepath.Join(util.GetEffectiveUserHomeFolder(), "services/svc_credential_manager/config")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	routeDefs := route.RoutesDef()
	bindedRouteDefs, _ := binding.BindDependencies(ctx, configPath,routeDefs)

	application := app.New(configPath, bindedRouteDefs)

	err := application.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
