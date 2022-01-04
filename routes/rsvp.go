package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
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
	FirstName                 string `json:"first_name" binding:"required"`
	LastName                  string `json:"last_name" binding:"required"`
	Email                     string `json:"email" binding:"required"`
	AdditionalGuestsAttending int    `json:"additional_guests_attending" binding:"required"`
	Attending                 bool   `json:"attending" binding:"required"`
}

func rsvp_new(resp *gin.Context, ctx *api.Context) {
	var new_rsvp NewRSVP

	if err := resp.ShouldBindJSON(&new_rsvp); err != nil {
		resp.AbortWithError(http.StatusBadRequest, err)
	} else {
		resp.JSON(http.StatusCreated, new_rsvp)
	}
}

func rsvp_get_all(resp *gin.Context, ctx *api.Context) {
	rsvps := []RSVP{}

	if rows, err := ctx.DBPool.Query(context.Background(), "select * from rsvps"); err == nil {
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
				resp.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error reading database",
				})

				return
			}

			// Add RSVP to list of RSVPs
			rsvps = append(rsvps, r)
		}

		resp.JSON(http.StatusOK, rsvps)
	} else {
		resp.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading database",
		})
	}
}

func rsvp_update(resp *gin.Context, ctx *api.Context) {
	resp.JSON(http.StatusOK, gin.H{})
}

func Routes_RSVP(server *gin.Engine, ctx *api.Context) {

	// Create group of routes regarding RSVPs
	g := server.Group("/rsvp")
	{
		g.GET("", handler(rsvp_get_all, ctx))
		g.POST("", handler(rsvp_new, ctx))
		g.PUT("", handler(rsvp_update, ctx))
	}
}
