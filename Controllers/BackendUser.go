package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gozpt/express/utils"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"strings"
)

func UserLogin(c *gin.Context) {
	user_name := strings.TrimSpace(c.PostForm("user_name"))
	password := strings.TrimSpace(c.PostForm("password"))
	if len(user_name) == 0 || len(password) == 0 {
		ApiHelpers.RespondJSON(c, 0, "", "用户名和密码不正确")
		return
	}

	password = utils.String2md5(password)
	var user Models.BackendUser
	err := Models.GetUser(&user, user_name, password)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, Models.BackendUser{Id: user.Id}, "success")
	}

}

func UserChangePwd(c *gin.Context) {
	uid := strings.TrimSpace(c.PostForm("uid"))
	password := strings.TrimSpace(c.PostForm("password"))
	if len(uid) == 0 || len(password) == 0 {
		ApiHelpers.RespondJSON(c, 0, "", "用户名和密码不能为空")
		return
	}

	password = utils.String2md5(password)
	var user Models.BackendUser

	err := Models.GetOneUser(&user, uid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	user.UserPwd = password;
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, uid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, "", "success")
	}
}
