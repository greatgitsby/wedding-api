package routes_test

import (
	"os"
	"testing"

	"github.com/greatgitsby/wedding-api/routes"
	"github.com/greatgitsby/wedding-api/test"
)

var ROUTER = test.GetRouter()

func TestMain(m *testing.M) {
	routes.Routes_Root(ROUTER, test.GetContext())
	routes.Routes_RSVP(ROUTER, test.GetContext())
	os.Exit(m.Run())
}
