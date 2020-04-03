package Controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
)

func GetOnePaymentCode(c *gin.Context) {
	order_code := c.PostForm("order_code")
	code_slice := map[string]bool{"wechat": true, "alipay": true}
	if _, ok := code_slice[order_code]; !ok {
		ApiHelpers.RespondJSON(c, 0, "", "order_code参数非法")
		return
	}

	var payment_code Models.PaymentCode
	err := Models.GetOnePaymentCode(&payment_code, order_code)
	if err != nil {
		payment_code.img = "http://express.zhang6.net" + payment_code.img
		ApiHelpers.RespondJSON(c, 0, payment_code, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, payment_code, "success")
	}
}
