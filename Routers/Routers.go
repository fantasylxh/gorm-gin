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
		v1.POST("order/order_add", Controllers.AddNewOrder) //添加订单
		v1.POST("order/list", Controllers.ListOrder) // 订单列表
		v1.POST("order/detail", Controllers.GetOneOrder) //获取订单详情
		v1.POST("order/change", Controllers.PutOneOrder) //更新订单
		v1.POST("order/delete", Controllers.DeleteOrder) // 删除订单
		v1.POST("order/upload", Controllers.Fileupload) // 上传文件
		v1.POST("order/ship", Controllers.GetOrderShip) // 订单物流
	}

	return r
}
