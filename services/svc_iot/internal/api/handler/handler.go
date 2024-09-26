package handler

import (
	"fmt"
	"iot-access-management/internal/models/api_to_core"
	"iot-access-management/internal/util/http_helper"
	"iot-access-management/services/svc_iot/internal/api/api_context"
	"iot-access-management/services/svc_iot/internal/core_manager"
	"net/http"
)

type IoTHandler struct {
	iotManager core_manager.IoTManager
}

func NewIoTHandler(iotm core_manager.IoTManager) *IoTHandler {
	return &IoTHandler{
		iotManager: iotm,
	}
}

func (ioth *IoTHandler) CreateWhiteList(w http.ResponseWriter, r *http.Request) {

	whiteListRequest, err := api_context.GetWhiteListCreatePayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", whiteListRequest, err))
		return
	}

	whiteList := api_to_core.ApiWhiteListCreateRequestToCoreWhiteList(*whiteListRequest)
	createErr := ioth.iotManager.CreateWhiteList(whiteList)

	if createErr != nil {
		http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to insert: [%v]", createErr))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ioth *IoTHandler) DeleteWhiteList(w http.ResponseWriter, r *http.Request) {

	whiteListRequest, err := api_context.GetWhiteListCreatePayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", whiteListRequest, err))
		return
	}

	whiteList := api_to_core.ApiWhiteListCreateRequestToCoreWhiteList(*whiteListRequest)
	deleteErr := ioth.iotManager.DeleteWhiteList(whiteList)
	if deleteErr != nil {
		switch {
		case deleteErr.Is(core_manager.ErrWhiteListNotFound):
			http_helper.RespondNotFoundWithError(w, deleteErr.Info())
			return
		default:
			http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to delete: [%v]", err))
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ioth *IoTHandler) RequestAccess(w http.ResponseWriter, r *http.Request) {

	accessRequest, err := api_context.GetAccessRequestPayloadFromBody(r)
	if err != nil {
		http_helper.RespondBadRequestWithError(w,
			fmt.Sprintf("json not valid, expected Json: [%+v]; error: [%v]", accessRequest, err))
		return
	}

	coreAccessRequest := api_to_core.ApiAccessRequestToCoreAccessRequest(*accessRequest)
	found, mgrErr := ioth.iotManager.RequestAccess(coreAccessRequest)
	if mgrErr != nil {
		http_helper.RespondUnauthorized(w, found)
		return
	}

	http_helper.RespondOkWithBody(w, found)
}

func (ioth *IoTHandler) ListWhiteList(w http.ResponseWriter, r *http.Request) {

	newCredential, err := ioth.iotManager.GetWhiteList()

	if err != nil {
		switch {
		case err.Is(core_manager.ErrWhiteListNotFound):
			http_helper.RespondNotFoundWithError(w, err.Info())
			return
		default:
			http_helper.RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to insert: [%v]", err))
			return
		}
	}

	http_helper.RespondWithStatusCreatedAndBody(w, newCredential)
}
