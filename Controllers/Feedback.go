package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"net/http"
	"strings"
)

func AddNewFeedback(c *gin.Context) {
	uid := strings.TrimSpace(c.PostForm("uid"))

	var feed Models.FeedBack
	feed.CreatorId = uid
	if err := c.ShouldBind(&feed); nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"Code": 0,
			"Msg":  err.Error(),
		})
		return
	}
	err := Models.AddNewFeedback(&feed)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, "", "success")
	}
}
