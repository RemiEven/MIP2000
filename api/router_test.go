package api

import (
	"encoding/json"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-test/deep"
)

func TestImagePath(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/image", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
	fileSystem := os.DirFS("./testdata/")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		t.Errorf("failed to open image directory")
		return
	}
	mux := NewRouter(readDirFileSystem)
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

	expectedFile, err := os.Open("./testdata/images/1000000845.jpg")
	if err != nil {
		t.Errorf("failed to open expected file: %v", err)
		return
	}
	expected, err := io.ReadAll(expectedFile)

	if err != nil {
		t.Errorf("failed to read expected file: %v", err)
		return
	}

	if diff := deep.Equal(body, expected); diff != nil {
		t.Errorf("unexpected response body: %v", diff)
		return
	}

}

func TestInternalServerError(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/image", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
	fileSystem := os.DirFS("./testdata/emptydirectory/")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		t.Errorf("failed to open image directory")
		return
	}
	mux := NewRouter(readDirFileSystem)
	mux.ServeHTTP(recorder, request)

	resp := recorder.Result()
	if status := resp.StatusCode; status != http.StatusInternalServerError {
		t.Errorf("unexpected response status: got [%v], wanted [%v]", status, http.StatusInternalServerError)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
		return
	}

	expectedErrorResponseCode := "internal-server-error"
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

func TestNotFound(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Errorf("failed to create request: %v", err)
		return
	}
	recorder := httptest.NewRecorder()
	fileSystem := os.DirFS("..")
	readDirFileSystem, ok := fileSystem.(fs.ReadDirFS)
	if !ok {
		t.Errorf("failed to open image directory")
		return
	}
	mux := NewRouter(readDirFileSystem)
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
