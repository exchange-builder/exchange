package main

import (
	"order/mappings"
)

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}
