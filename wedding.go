package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
	"github.com/greatgitsby/wedding-api/routes"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
func main() {

	var api_ctx api.Context
	var err error
	var port int
	var port_str string
	var port_exists bool
	var db *pgxpool.Pool
	var db_url string
	var db_url_exists bool
	var sess *session.Session

	LoadEnv()

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
	if db, err = api.GetDBConn(db_url); err != nil {
		log.Fatalln("DB error:", err)
	}

	if sess, err = api.GetAWSSession(); err != nil {
		log.Fatalln("AWS error:", err)
	}

	defer db.Close()

	gin.SetMode(gin.DebugMode)

	// Setup server
	s := gin.Default()

	api_ctx.DBPool = db
	api_ctx.AWSSession = sess

	// Register root routes
	routes.Routes_Root(s, &api_ctx)

	// Register RSVP routes
	routes.Routes_RSVP(s, &api_ctx)

	// Register image upload routes
	routes.Routes_Images(s, &api_ctx)

	// Listen
	s.Run(fmt.Sprintf(":%d", port))
}
