package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func coucou(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.Write([]byte("hello, coucou"))
}

type ErrorResponse struct {
	Code string `json:"code"`
}

func notFound(responseWriter http.ResponseWriter, _ *http.Request) {
	responseWriter.WriteHeader(http.StatusNotFound)
	encoder := json.NewEncoder(responseWriter)
	errorResponse := ErrorResponse{
		Code: "not-found",
	}
	if err := encoder.Encode(errorResponse); err != nil {
		slog.Error("failed to encode error response", err)
	}
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/coucou", coucou)
	mux.HandleFunc("/", notFound)
	return mux
}
