package Gin_shop

import (
	"Gin-shop/config"
	"Gin-shop/routers"
)

func main() {
	//实例化路由
	router := routers.NewRouter()
	//启动监听并执行GIN
	router.Run(config.Port)

}
