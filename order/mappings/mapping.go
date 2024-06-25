package mappings

import (
	"github.com/gin-gonic/gin"
	"order/controllers"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// v1 of the API
	v1 := Router.Group("/v1")
	{
		//v1.GET("/order/:id", controllers.GetUserDetail)
		v1.GET("/order/list", controllers.GetOrderList)
		v1.POST("/order/add", controllers.PostOrder)
	}
}
