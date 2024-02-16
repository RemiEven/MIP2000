package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCoucou(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	w := httptest.NewRecorder()

	Coucou(w, request)

	resp := w.Result()
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("unexpected response status: got [%v], wanted [%v]", status, http.StatusOK)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
		return
	}

	if string(body) != "hello, coucou" {
		t.Errorf("unexpected response body: got [%v], wanted [%v]", string(body), "hello, coucou")
	}
}
