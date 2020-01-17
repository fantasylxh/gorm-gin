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
		v1.POST("order/order_add", Controllers.AddNewOrder)
		v1.POST("order/list", Controllers.ListOrder)
		v1.POST("order/detail", Controllers.GetOneOrder)
	}

	return r
}

