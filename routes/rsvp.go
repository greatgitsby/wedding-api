package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func rsvp_new(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{})
}

func rsvp_get_all(c *gin.Context) {
	rsvps := []string{}
	c.JSON(http.StatusOK, rsvps)
}

func rsvp_update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func RSVP(server *gin.Engine) {

	// Create group of routes regarding RSVPs
	server.GET("/rsvp", rsvp_get_all)
	server.POST("/rsvp", rsvp_new)
	server.PUT("/rsvp", rsvp_update)
}
