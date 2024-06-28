package controllers

import (
	"database/sql"
	"fmt"
	"github.com/antlabs/pcopy"
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/teris-io/shortid"
	"log"
	"net/http"
	"order/model"
	"os"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

var dbMap = initDb()

func init() {
	pcopy.Preheat(&model.OrderDto{}, &model.OrderReq{}) // 一对类型只要预热一次

}
func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/exchange")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(model.OrderDto{}, "ex_order").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
func PostOrder(c *gin.Context) {
	//Place an order
	body := model.OrderReq{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	log.Println(body)
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	generate, _ := sid.Generate()
	orderDto := &model.OrderDto{}
	err := pcopy.Copy(&orderDto, &body, pcopy.WithUsePreheat())
	if err != nil {
		return
	}
	orderDto.OrderCode = generate
	orderDto.CreateTime = time.Now()
	orderDto.UpdateTime = time.Now()
	log.Println(orderDto)
	err1 := dbMap.Insert(orderDto)
	checkErr(err1, "Insert failed")
	if err1 != nil {
		log.Println(err1.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "order created",
		"id":      generate,
	})

}

func GetOrderList(c *gin.Context) {
	//get orders
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

}