package repo

import (
	"iot-access-management/internal/error/trace_error"
	repo_model "iot-access-management/internal/models/repo"
)

type IotRepo interface {
	CreateWhiteList(whitelist repo_model.WhiteList) *trace_error.TraceError
	DeleteWhiteList(whitelist repo_model.WhiteList) *trace_error.TraceError
	ListWhiteList() ([]repo_model.WhiteList, *trace_error.TraceError)
}
