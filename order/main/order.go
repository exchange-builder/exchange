package main

import (
	"order/bootstrap"
	"order/global"
	"order/mappings"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	mappings.CreateUrlMappings()
	mappings.Router.Run(":" + global.App.Config.App.Port)
}
