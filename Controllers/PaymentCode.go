package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
)

func GetOnePaymentCode(c *gin.Context) {
	code := c.PostForm("code")
	code_slice := map[string]bool{"wechat": true,"alipay": true}
	if _, ok := code_slice[code]; !ok {
		ApiHelpers.RespondJSON(c, 0, "", "code参数非法")
		return
	}

	var payment_code Models.PaymentCode
	err := Models.GetOnePaymentCode(&payment_code,code)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0,payment_code , err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, payment_code, "success")
	}
}
