package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RSVP struct {
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name"`
	Email                     string `json:"email"`
	AdditionalGuestsAttending int    `json:"additional_guests_attending"`
	AdditionalGuestsMax       int    `json:"additional_guests_max"`
	Attending                 bool   `json:"attending"`
}

type NewRSVP struct {
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name"`
	Email                     string `json:"email"`
	AdditionalGuestsAttending int    `json:"additional_guests_attending"`
	Attending                 bool   `json:"attending"`
}

func rsvp_new(c *gin.Context, db *pgxpool.Pool) {
	c.JSON(http.StatusCreated, gin.H{})
}

func rsvp_get_all(c *gin.Context, db *pgxpool.Pool) {
	rsvps := []RSVP{}

	if rows, err := db.Query(context.Background(), "select * from rsvps"); err == nil {
		for rows.Next() {
			var r RSVP

			// Load row into RSVP struct
			err := rows.Scan(
				&r.LastName,
				&r.FirstName,
				&r.Email,
				&r.AdditionalGuestsAttending,
				&r.AdditionalGuestsMax,
				&r.Attending)

			// Exit if there is an error, responding as such
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error reading database",
				})

				return
			}

			// Add RSVP to list of RSVPs
			rsvps = append(rsvps, r)
		}

		c.JSON(http.StatusOK, rsvps)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading database",
		})
	}
}

func rsvp_update(c *gin.Context, db *pgxpool.Pool) {
	c.JSON(http.StatusOK, gin.H{})
}

func handler(handle func(c *gin.Context, db *pgxpool.Pool), db *pgxpool.Pool) gin.HandlerFunc {

	// Pass db to handle func
	closure := func(c *gin.Context) {
		handle(c, db)
	}

	return gin.HandlerFunc(closure)
}

func Routes_RSVP(server *gin.Engine, db *pgxpool.Pool) {

	// Create group of routes regarding RSVPs
	g := server.Group("/rsvp")
	{
		g.GET("", handler(rsvp_get_all, db))
		g.POST("", handler(rsvp_new, db))
		g.PUT("", handler(rsvp_update, db))

	}
}
