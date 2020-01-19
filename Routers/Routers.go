package Routers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("book", Controllers.ListBook)
		v1.POST("book", Controllers.AddNewBook)
		v1.GET("book/:id", Controllers.GetOneBook)
		v1.PUT("book/:id", Controllers.PutOneBook)
		v1.DELETE("book/:id", Controllers.DeleteBook)

		v1.GET("category/index", Controllers.ListCategory)
		v1.POST("category/getprice", Controllers.GetOneCategoryPrice)

		v1.POST("user/login", Controllers.UserLogin)
		v1.POST("user/modpwd", Controllers.UserChangePwd)
		v1.POST("order/add", Controllers.AddNewOrder) //添加订单
		v1.POST("order/list", Controllers.ListOrder) // 订单列表
		v1.POST("order/detail", Controllers.GetOneOrder) //获取订单详情
		v1.POST("order/change", Controllers.PutOneOrder) //更新订单
		v1.POST("order/delete", Controllers.DeleteOrder) // 删除订单
		v1.POST("order/upload", Controllers.Fileupload) // 上传文件
		v1.POST("order/ship", Controllers.GetOrderShip) // 订单物流
		v1.POST("order/order_do", Controllers.OrderDo) // 订单签收 订单支付 订单取消
		//地址相关
		v1.POST("address/index", Controllers.ListAddress) // 获取地址列表
		v1.POST("address/add", Controllers.AddNewAddress)// 地址添加
		v1.POST("address/detail", Controllers.GetOneAddress)// 获取地址详情
		v1.POST("address/change", Controllers.PutOneAddress)// 修改地址
		v1.POST("address/delete", Controllers.DeleteAddress)// 删除地址

		v1.POST("payment/qrcode", Controllers.GetOnePaymentCode) // 获取付款二维码
	}

	return r
}
