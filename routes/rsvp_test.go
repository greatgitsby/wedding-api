package routes_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/greatgitsby/wedding-api/routes"
	"github.com/greatgitsby/wedding-api/test"
)

var ROUTER = test.GetRouter()

func TestMain(m *testing.M) {
	routes.Routes_RSVP(ROUTER, test.GetContext())
	os.Exit(m.Run())
}

func TestRSVPGetAllRoute(t *testing.T) {
	req, _ := http.NewRequest("GET", "/rsvp", nil)

	test.TestHTTPResponse(t, ROUTER, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusOK
	})
}
