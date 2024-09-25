package main

import (
	"context"
	"fmt"
	"iot-access-management/internal/app"
	"iot-access-management/services/svc_credential_manager/binding"
	"iot-access-management/services/svc_credential_manager/internal/api/route"
	"os"
	"os/signal"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	binder := &binding.CredentialManagerBinder{}
	routeDefs := route.RoutesDef()

	application := app.New(ctx, binder, routeDefs)

	err := application.Start()
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
