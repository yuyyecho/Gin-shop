package routers

import (
	"Gin-shop/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//gin框架创建
	router := gin.Default()

	//注册中间件
	//鉴权
	//日志
	//跨域

	//静态资源路径设置

	//绑定路由规则
	//注册用户
	router.GET("/register", controllers.ShowReg)
	router.POST("/register", controllers.HandleReg)
	////用户激活
	//router.GET("/active", controllers.ShowActive)
	////用户登录
	//router.GET("/login", controllers.ShowLogin)
	//router.POST("/login", controllers.HandleLogin)
	//
	////首页
	//router.GET("/", controllers.ShowIndex)
	//
	////退出登录
	//router.GET("/user/logout", controllers.ShowLogout)
	//
	////可以使用路由分组
	////用户中心信息页
	//router.GET("/user/userCenterInfo", controllers.ShowUserCenterInfo)
	////用户中心订单页
	//router.GET("/user/userCenterOrder", controllers.ShowUserCenterOrder)
	////用户中心地址页
	//router.GET("/user/userCenterSite", controllers.ShowUserCenterSite)
	//router.POST("/user/userCenterSite", controllers.HandleUserCenterSite)
	//
	////商品详情页
	//router.GET("/goodsDetail", controllers.ShowGoodsDetail)
	////商品列表页
	//router.GET("/goodsList", controllers.ShowGoodsList)
	////商品搜索页
	//router.POST("/goodsSearch", controllers.HandleGoodsSearch)
	//
	////添加购物车
	//router.POST("/user/addCart", controllers.HandleAddCart)
	////购物车展示
	//router.GET("/user/cart", controllers.ShowCart)
	////更新购物车数据
	//router.POST("/user/UpdateCart", controllers.HandleUpdateCart)
	////删除购物车数据
	//router.POST("/user/DeleteCart", controllers.HandleDeleteCart)
	//
	////展示订单页面
	//router.POST("/user/ShowOrder", controllers.HandleShowOrder)
	////添加订单
	//router.POST("/user/AddOrder", controllers.HandleAddOrder)
	////处理支付
	//router.GET("/user/pay", controllers.Pay)
	////支付成功
	//router.GET("/user/payok", controllers.PayOk)

	return router
}
