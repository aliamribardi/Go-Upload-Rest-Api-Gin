package Routes

import (
	"Go-Upload-Rest-Api-Gin/Controllers"
	"github.com/gin-gonic/gin"
)

func RouteApi() {
	Route := gin.Default()
	// Route.Static("/Assets/Upload/Image/", "./Assets/Upload/Image/")
	// Route.LoadHTMLGlob("Views/*")

	Route.MaxMultipartMemory = 8 << 20 

	Route.POST("/api", Controllers.UploadApi)

	Route.Run("localhost:8080")
}
