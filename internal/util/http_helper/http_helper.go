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

func RespondNotFoundWithError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(errMsg))
}

func RespondBadRequestWithError(w http.ResponseWriter, errMsg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(errMsg))
}

func RespondInternalServerErrorWithError(w http.ResponseWriter, errMsg string) {
	slog.Error(errMsg)
	w.WriteHeader(http.StatusInternalServerError)
}

func RespondUnauthorized(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	_, err = w.Write(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
