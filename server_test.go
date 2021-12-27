package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}
