package Routes

import (
	"Go-Upload-Rest-Api-Gin/Controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Route() {
	Route := gin.Default()
	Route.Static("/Assets/Upload/Image/", "./Assets/Upload/Image/")
	Route.LoadHTMLGlob("Views/*")

	Route.MaxMultipartMemory = 8 << 20

	Route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	Route.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	Route.POST("/upload", Controllers.Upload)

	Route.Run()
}