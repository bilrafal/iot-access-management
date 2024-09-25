package route

import (
	"iot-access-management/internal/app"
	handler_manager "iot-access-management/services/svc_credential_manager/internal/api/handler"
	"net/http"
)

func RoutesDef() (routeDefs []app.RouteDef) {
	hdl := &handler_manager.CredentialHandler{}

	routeDefs = append(routeDefs, app.NewRouteDef(http.MethodPost, "create-user", hdl.CreateUser))
	routeDefs = append(routeDefs, app.NewRouteDef(http.MethodPost, "get-user/{id}", hdl.GetUser))

	return
}
