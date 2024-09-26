package route

import (
	"iot-access-management/internal/router"
	handler_manager "iot-access-management/services/svc_iot/internal/api/handler"
	"net/http"
)

func RoutesDef() (routeDefs []router.RouteDef) {
	hdl := &handler_manager.IoTHandler{}

	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "white-list", hdl.CreateWhiteList))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodDelete, "white-list", hdl.DeleteWhiteList))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodPost, "access", hdl.RequestAccess))
	routeDefs = append(routeDefs, router.NewRouteDef(http.MethodGet, "white-list", hdl.ListWhiteList))

	return
}
