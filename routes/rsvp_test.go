package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/greatgitsby/wedding-api/test"
)

func TestRSVPGetAllRoute(t *testing.T) {
	expectedStatusCode := http.StatusOK

	r := test.GetRouter()

	Routes_RSVP(r)

	req, _ := http.NewRequest("GET", "/rsvp", nil)

	test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		println(w.Code)
		return w.Code == expectedStatusCode
	})
}
