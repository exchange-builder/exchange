package main

import (
	"order/bootstrap"
	"order/controllers"
	"order/global"
	"order/mappings"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	controllers.InitDb()
	mappings.CreateUrlMappings()
	mappings.Router.Run(":" + global.App.Config.App.Port)
}
