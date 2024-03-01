package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCoucou(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/coucou", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
	mux := NewRouter()
	mux.ServeHTTP(recorder, request)

	resp := recorder.Result()
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

func TestNotFound(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
	mux := NewRouter()
	mux.ServeHTTP(recorder, request)

	resp := recorder.Result()
	expectedStatus := http.StatusNotFound
	if status := resp.StatusCode; status != expectedStatus {
		t.Errorf("unexpected response status: got [%v], wanted [%v]", status, expectedStatus)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
		return
	}

	expectedErrorResponseCode := "not-found"
	errorResponse := ErrorResponse{}

	if err := json.Unmarshal(body, &errorResponse); err != nil {
		t.Errorf("failed to parse json response body: %v", err)
		return
	}
	if actualResponseCode := errorResponse.Code; actualResponseCode != expectedErrorResponseCode {
		t.Errorf("unexpected response code: got [%v], wanted [%v]", actualResponseCode, expectedErrorResponseCode)
		return
	}
}
