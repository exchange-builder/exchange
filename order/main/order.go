package main

import (
	"fmt"
	"net/http"
	"order/model"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//Place an order
	r.POST("/order", func(c *gin.Context) {
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
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
