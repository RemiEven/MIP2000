package api

import (
	"encoding/json"
	"io"
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/remieven/mip2000/memeselector"
)

func image(readDirFileSystem fs.ReadDirFS) func(responseWriter http.ResponseWriter, _ *http.Request) {
	return func(responseWriter http.ResponseWriter, _ *http.Request) {
		meme, err := memeselector.GetRandomMeme(readDirFileSystem)
		if err != nil {
			responseWriter.WriteHeader(http.StatusInternalServerError)
			responseWriter.Write([]byte(err.Error()))
			return
		}
		io.Copy(responseWriter, meme)
	}
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

func NewRouter(readDirFileSystem fs.ReadDirFS) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/image", image(readDirFileSystem))
	mux.HandleFunc("/", notFound)
	return mux
}
