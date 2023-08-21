package Routes

import (
	"github.com/gin-gonic/gin"
)

func RouteApi() {
	Route := gin.Default()

	Route.MaxMultipartMemory = 8 << 20 

	Route.Run()
}
