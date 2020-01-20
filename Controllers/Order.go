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
	"strconv"
	"strings"
	"time"

	//	"time"
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
	uInt, err := strconv.Atoi(uid)
	order_method := strings.TrimSpace(c.PostForm("order_method"))
	// 检查参数
	method_slice := map[string]bool{"order_pay_confirm": true, "order_cancel": true, "order_sign": true}
	if _, ok := method_slice[order_method]; !ok {
		ApiHelpers.RespondJSON(c, 0, "", "order_method参数非法")
		return
	}
	// 获取用户信息
	var user Models.BackendUser
	Models.GetOneUser(&user, uid)

	err = Models.GetOneOrder(&order, id)

	if err != nil {
		ApiHelpers.RespondJSON(c, 0, "", err.Error())
		return
	}
	if uid != order.CreatorId {
		ApiHelpers.RespondJSON(c, 0, "", "订单不存在")
		return
	}

	var orderActionNote string

	switch order_method {
	case "order_pay_confirm": // 确认支付order_status =1 更新付款时间
		Models.UpdateOrderPayTime(&order, id)
		order.OrderStatus = 1
		orderActionNote = user.UserName + ":" + user.RealName + "修改了订单支付状态为:已付款";

	case "order_cancel": // 取消订单 order_status =2
		order.OrderStatus = 2
		orderActionNote = user.UserName + ":" + user.RealName + "修改了订单状态为:取消订单";
	case "order_sign": // 签收订单 ship_status=4 记录物流 未支付 未发货的不能签收
		if order.OrderStatus < 1 || order.ShipStatus < 3 {
			ApiHelpers.RespondJSON(c, 0, "", "未支付/未发货的快递无法操作")
			return
		}
		// 避免重复签收
		if order.ShipStatus == 4 {
			ApiHelpers.RespondJSON(c, 0, "", "该快递已签收,无需重复操作")
			return
		}

		order.ShipStatus = 4 // 更新订单状态为签收
		orderActionNote = user.UserName + ":" + user.RealName + "修改了物流状态为:已签收";
		// 记录物流信息
		var saction Models.ShipAction
		c.ShouldBind(&saction)
		saction.ActionNote = "收件人已签收,欢迎再次使用。"
		saction.CreatorId = uInt
		saction.ShipStatus = 4 // 标记签收状态
		saction.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		Models.AddNewShipAction(&saction)
	}
	// 记录order_action
	var oaction Models.OrderAction
	c.ShouldBind(&oaction)
	oaction.ActionNote = orderActionNote
	oaction.CreatorId = uInt
	Models.AddNewOrderAction(&oaction)

	c.Bind(&order)
	err = Models.PutOneOrder(&order, id)

	if err != nil {
		ApiHelpers.RespondJSON(c, 0, order, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success")
	}

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

	code := c.PostForm("code")
	code_slice := map[string]bool{"wechat": true, "alipay": true, "cash ": true}
	if _, ok := code_slice[code]; !ok {
		ApiHelpers.RespondJSON(c, 0, "", "code参数非法")
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

	order.PayQrcode = filename // 更新二维码
	order.OrderCode = code     // 更新支付类型
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
