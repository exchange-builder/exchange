package main

import (
	"order/web"
)

func main() {
	r := web.SetupRouter()
	r = web.PostOrder(r)
	r.Run(":8080")
}
