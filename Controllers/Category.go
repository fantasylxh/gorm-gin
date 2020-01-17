package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
)

func ListCategory(c *gin.Context) {
	var category []Models.Category
	err := Models.GetAllCategory(&category)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, category,err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, category,"success")
	}
}
