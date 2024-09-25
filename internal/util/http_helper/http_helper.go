package http_helper

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func RespondOkWithBody(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RespondBadRequestWithError(w http.ResponseWriter, errMsg string) {
	w.Write([]byte(errMsg))
	w.WriteHeader(http.StatusBadRequest)
}

func RespondInternalServerErrorWithError(w http.ResponseWriter, errMsg string) {
	slog.Error(errMsg)
	w.WriteHeader(http.StatusInternalServerError)
}

func RespondWithStatusCreatedAndBody(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to marshal: [%v]", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(res)
	if err != nil {
		RespondInternalServerErrorWithError(w, fmt.Sprintf("failed to write reponse: [%v]", err))
		return
	}
}
