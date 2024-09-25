package http_helper

import (
	"encoding/json"
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
