package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/routes"
)

func root(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"version": "v1",
	})
}

func main() {

	var err error
	var port int
	var port_str string
	var port_exists bool

	port_str, port_exists = os.LookupEnv("PORT")

	if !port_exists {
		port_str = "8000"
	}

	if port, err = strconv.Atoi(port_str); err != nil {
		log.Fatalln("Invalid port")
	}

	// Setup server
	s := gin.Default()

	// Root route
	s.GET("/", root)

	// Register RSVP routes
	routes.RSVP(s)

	// Listen
	s.Run(fmt.Sprintf("localhost:%d", port))
}
