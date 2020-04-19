package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLogin mocks the login API endpoint and tests the functionality
func TestLogin(t *testing.T) {

	uri := "/api/v1/login"

	server := newTestServer(uri, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	defer server.Close()

	statusCode := Login(server.URL, "username", "password")
	if statusCode != 200 {
		t.Errorf("Test Login failed")
	}
}

// newTestServer creates a new multiplex server to handle the API endpoints
func newTestServer(path string, h func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	mux.HandleFunc(path, h)

	return server
}
