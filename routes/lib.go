package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
)

func handler(handle func(c *gin.Context, ctx *api.Context), ctx *api.Context) gin.HandlerFunc {

	// Pass db to handle func
	closure := func(c *gin.Context) {
		handle(c, ctx)
	}

	return gin.HandlerFunc(closure)
}
