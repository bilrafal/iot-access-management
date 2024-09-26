package core_manager

import (
	"iot-access-management/internal/error/trace_error"
	"iot-access-management/internal/models/core"
)

var (
	ErrWhiteListNotFound = trace_error.NewTraceError("NO_WHITELIST_FOUND")
)

type IoTManager interface {
	CreateWhiteList(whitelist core.WhiteList) *trace_error.TraceError
	DeleteWhiteList(whitelist core.WhiteList) *trace_error.TraceError
	GetWhiteList() ([]core.WhiteList, *trace_error.TraceError)
	RequestAccess(accessRequest core.AccessRequest) (bool, *trace_error.TraceError)
}
