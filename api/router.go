package api

import "net/http"

func Coucou(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("hello, coucou"))
}
