package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"net/http"
	"order/model"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func PostOrder(r *gin.Engine) *gin.Engine {
	//Place an order
	r.POST("/order/add", func(c *gin.Context) {
		body := model.OrderReq{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
		generate, _ := sid.Generate()
		c.JSON(http.StatusOK, gin.H{
			"message": "order created",
			"id":      generate,
		})
	})
	return r
}
