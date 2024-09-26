package route

import (
	"iot-access-management/internal/router"
	handler_manager "iot-access-management/services/svc_credential_manager/internal/api/handler"
	"net/http"
)

func RoutesDef() (routeDefs []router.RouteDef) {
	hdl := &handler_manager.CredentialHandler{}

	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "user", hdl.CreateUser))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodGet, "user/{id}", hdl.GetUser))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "credential", hdl.CreateCredential))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodGet, "credential/{code}", hdl.GetCredential))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "user-credential", hdl.AssignCredentialToUser))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodGet, "user-credential/{id}", hdl.GetUserCredentials))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "authorize/{door-id}/{credential-id}", hdl.AuthorizeUserOnDoor))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodDelete, "authorize/{door-id}/{credential-id}", hdl.RevokeAuthorization))

	return
}
