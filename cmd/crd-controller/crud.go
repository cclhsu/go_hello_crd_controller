package main

import (
	resources "go_hello_crd_controller/cmd/crd-controller/db/sqlite3"

	"github.com/gin-gonic/gin"
)

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	}
// }

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// router.Use(Cors())

	v1 := router.Group("api/v1")
	{
		v1.POST("/table", resources.CreateResourceTable)
		v1.DELETE("/table", resources.DropResourceTable)

		v1.GET("/resources", resources.GetResources)
		v1.GET("/resources/:id", resources.GetResource)
		v1.POST("/resources", resources.PostResource)
		v1.PUT("/resources/:id", resources.UpdateResource)
		v1.DELETE("/resources/:id", resources.DeleteResource)

		v1.GET("/resources/:id/health", resources.GetResourceHealth)
		v1.PUT("/resources/:id/health", resources.UpdateResourceHealth)
	}

	return router
}

func main() {
	router := SetupRouter()
	router.Run() // listen and server on 0.0.0.0:8080
}
