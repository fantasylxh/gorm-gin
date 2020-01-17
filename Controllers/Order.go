package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"strings"
)

func ListOrder(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrder(&order)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {

		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func AddNewOrder(c *gin.Context) {
	var order Models.Order
	c.ShouldBind(&order)
	err := Models.AddNewOrder(&order)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func GetOneOrder(c *gin.Context) {
	id := strings.TrimSpace(c.PostForm("order_id"))
	var order Models.Order
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func PutOneOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	}
	c.BindJSON(&order)
	err = Models.PutOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order, "api接口")
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "api接口")
	}
}
