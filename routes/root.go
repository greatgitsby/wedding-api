package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
)

func root(resp *gin.Context, ctx *api.Context) {
	resp.JSON(http.StatusOK, gin.H{
		"contact": "hello@owen2moen.com",
		"version": "v1",
	})
}

func Routes_Root(server *gin.Engine, ctx *api.Context) {
	g := server.Group("/")
	{
		g.GET("", handler(root, ctx))
	}
}
