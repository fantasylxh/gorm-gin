package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
)

func AddNewFeedback(c *gin.Context) {
	var feed Models.FeedBack
	c.ShouldBind(&feed)
	err := Models.AddNewFeedback(&feed)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, feed, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, feed, "success")
	}
}
