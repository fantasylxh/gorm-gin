package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"strings"
)

func GetOneCategoryPrice(c *gin.Context) {
	id := c.PostForm("id")
	weight := strings.TrimSpace(c.PostForm("weight"))

	if weight == "" || id == "" {
		ApiHelpers.RespondJSON(c, 0, "", "weight 重量或分类id不能为空")
		return
	}
	var category_price Models.CategoryPrice
	err := Models.GetOneCategoryPrice(&category_price, id, weight)
	if err != nil {
		category_price.Price = "10" // 初始值
		ApiHelpers.RespondJSON(c, 200, category_price, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, category_price, "success")
	}
}
