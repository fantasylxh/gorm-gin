package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Models"
	"io"
	"log"
	"net/http"
	"os"
	"path"
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
	uid := strings.TrimSpace(c.PostForm("uid"))
	order_sn := strings.TrimSpace(c.PostForm("order_sn"))
	if order_sn == "" {
		ApiHelpers.RespondJSON(c, 0, "", "order_sn不能为空")
		return
	}
	order.CreatorId = uid
	c.ShouldBind(&order)
	err := Models.AddNewOrder(&order)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
		return
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func GetOneOrder(c *gin.Context) {
	id := strings.TrimSpace(c.PostForm("order_id"))
	var order Models.Order
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func GetOrderShip(c *gin.Context) {
	id := strings.TrimSpace(c.PostForm("order_id"))
	var order []Models.ShipAction
	err := Models.GetOrderShip(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func OrderDo(c *gin.Context) {

	var order Models.Order
	id := strings.TrimSpace(c.PostForm("order_id"))
	uid := strings.TrimSpace(c.PostForm("uid"))
	order_method := strings.TrimSpace(c.PostForm("order_method"))
	//value := strings.TrimSpace(c.PostForm("value"))
	err := Models.GetOneOrder(&order, id)
	var grade string = "B"
	switch order_method {
	case "order_pay_confirm": // 确认支付 更新付款时间
		grade = "A"
	case "order_cancel":// 取消订单
		grade = "B"
	case "order_sign":// 签收订单 记录物流
		grade = "C"
	default:
		fmt.Printf("你的等级是 %s\n", grade );
	}
	// 记录order_action

	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	if uid != order.CreatorId {
		ApiHelpers.RespondJSON(c, 0, "", "订单不存在")
		return
	}
	ApiHelpers.RespondJSON(c, 200, "", "success")
}

func PutOneOrder(c *gin.Context) {
	var order Models.Order
	id := strings.TrimSpace(c.PostForm("order_id"))
	uid := strings.TrimSpace(c.PostForm("uid"))
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	if uid != order.CreatorId {
		ApiHelpers.RespondJSON(c, 0, "", "订单不存在")
		return
	}
	if order.ShipStatus > 1 {
		ApiHelpers.RespondJSON(c, 0, "", "已发货无法修改")
		return
	}
	c.ShouldBind(&order)

	err = Models.PutOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := strings.TrimSpace(c.PostForm("order_id"))
	uid := strings.TrimSpace(c.PostForm("uid"))
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	if uid != order.CreatorId {
		ApiHelpers.RespondJSON(c, 0, "", "订单不存在")
		return
	}
	err = Models.DeleteOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}

/**上传方法**/
func Fileupload(c *gin.Context) {
	id := strings.TrimSpace(c.PostForm("order_id"))
	uid := strings.TrimSpace(c.PostForm("uid"))
	var order Models.Order
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	if uid != order.CreatorId {
		ApiHelpers.RespondJSON(c, 0, "", "订单不存在")
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename
	filename = ApiHelpers.GetRandomString(10) + path.Ext(filename)
	fmt.Println(file, err, filename)
	//创建文件
	out, err := os.Create("static/uploadfile/" + filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	// 更新二维码
	order.PayQrcode = filename
	c.ShouldBind(&order)

	err = Models.PutOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	} else {
		order.PayQrcode = ApiHelpers.Geturl(c.Request) + "/static/uploadfile/" + filename
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}
}
