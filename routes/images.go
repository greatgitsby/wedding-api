package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/greatgitsby/wedding-api/api"
)

func upload_images(resp *gin.Context, ctx *api.Context) {
	uploader := s3manager.NewUploader(ctx.AWSSession)

	file, header, err := resp.Request.FormFile("file")

	if err != nil {
		resp.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	aws_key := fmt.Sprintf("user-upload/images/%d-%s", time.Now().Unix(), header.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("owen2moen"),
		Key:    &aws_key,
		Body:   file,
	})

	if err != nil {
		log.Println(err)
		resp.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"contact": "hello@owen2moen.com",
		})
	} else {
		resp.JSON(http.StatusAccepted, gin.H{
			"success": true,
		})
	}
}

func Routes_Images(server *gin.Engine, ctx *api.Context) {
	g := server.Group("/images")
	{
		g.POST("", handler(upload_images, ctx))
	}
}
