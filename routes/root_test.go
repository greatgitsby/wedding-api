package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/greatgitsby/wedding-api/test"
)

func TestRootRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	test.TestHTTPResponse(t, ROUTER, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusOK
	})
}
