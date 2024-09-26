package main

import (
	"context"
	"fmt"
	"iot-access-management/internal/app"
	"iot-access-management/services/svc_iot/binding"
	"iot-access-management/services/svc_iot/internal/api/route"
	"os"
	"os/signal"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	binder := &binding.IoTManagerBinder{}
	routeDefs := route.RoutesDef()

	application := app.New(ctx, binder, routeDefs)

	err := application.Start()
	if err != nil {
		fmt.Println("failed to start IoT app:", err)
	}

}
