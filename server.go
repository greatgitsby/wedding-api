package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/routes"
	"github.com/jackc/pgx/v4/pgxpool"
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
	var db *pgxpool.Pool
	var db_url string
	var db_url_exists bool

	port_str, port_exists = os.LookupEnv("PORT")
	db_url, db_url_exists = os.LookupEnv("DATABASE_URL")

	if !db_url_exists {
		db_url = "postgres://localhost/test"
	}

	if !port_exists {
		port_str = "8000"
	}

	if port, err = strconv.Atoi(port_str); err != nil {
		log.Fatalln("Invalid port")
	}

	// Setup DB connection
	db, err = pgxpool.Connect(context.Background(), db_url)

	if err != nil {
		log.Fatalln("DB error:", err)
	}

	defer db.Close()

	// Setup server
	s := gin.Default()

	// Root route
	s.GET("/", root)

	// Register RSVP routes
	routes.Routes_RSVP(s, db)

	// Listen
	s.Run(fmt.Sprintf("0.0.0.0:%d", port))
}
