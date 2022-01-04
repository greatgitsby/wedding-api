package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
)

var DB_URL string = "postgres://localhost/test"
var API_CTX *api.Context = nil

func GetContext() *api.Context {
	if API_CTX == nil {
		db, err := api.GetDBConn(DB_URL)

		if err != nil {
			log.Fatalln("Could not get DB pool")
		}

		API_CTX = &api.Context{
			DBPool: db,
		}
	}

	return API_CTX
}

// Helper function to create a router during testing
func GetRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	return gin.Default()
}

// Helper function to process a request and test its response
func TestHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// This is a helper function that allows us to reuse some code in the above
// test methods
func TestMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Process the request and test the response
	TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}
