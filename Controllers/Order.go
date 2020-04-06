package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-gin/ApiHelpers"
	"gorm-gin/Config"
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
	uid := strings.TrimSpace(c.PostForm("uid"))
	// 获取当前用户角色
	var role Models.RoleBackenduserRel
	err := Models.GetOneRole(&role, uid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 0, role, "当前用户没有分配角色")
		return
	}
	var keyMap map[string]string = map[string]string{
		"q_type":       c.DefaultPostForm("q_type", "1"), //查询类型 1：签收 2：发走 默认1为签收
		"keyword":      strings.TrimSpace(c.PostForm("keyword")),
		"page":         c.DefaultPostForm("page", "2"),
		"order_status": strings.TrimSpace(c.PostForm("order_status")),
		"ship_status":  strings.TrimSpace(c.PostForm("ship_status")),
		"uid":          uid,
		"role_id":      role.RoleId,
	}

	err, count := Models.GetAllOrder(&order, keyMap)

	if err != nil {
		ApiHelpers.RespondJSON(c, 0, count, err.Error())
	} else {
		ApiHelpers.RespondJSON(c, 200, order, "success", )
	}
}

func AddNewOrder(c *gin.Context) {
	var order Models.Order
	uid := strings.TrimSpace(c.PostForm("uid"))
	address_id := strings.TrimSpace(c.PostForm("address_id"))
	rec_address_id := strings.TrimSpace(c.PostForm("rec_address_id"))
	order_sn := strings.TrimSpace(c.PostForm("order_sn"))
	if order_sn == "" {
		ApiHelpers.RespondJSON(c, 0, "", "order_sn不能为空")
		return
	}
	order.CreatorId = uid
	c.ShouldBind(&order)
	// 同步处理address_id rec_address_id 到order表
	if address_id != "" {
		var address Models.Address
		err := Models.GetOneAddress(&address, address_id)
		if err != nil {
			ApiHelpers.RespondJSON(c, 0, "", "address_id地址记录不存在")
			return
		}
		order.Name = address.Name
		order.Country = address.Country
		order.Province = address.Province
		order.City = address.City
		order.Address = address.Address
		order.Mobile = address.Mobile
	}
	if rec_address_id != "" {
		var address Models.Address
		err := Models.GetOneAddress(&address, rec_address_id)
		if err != nil {
			ApiHelpers.RespondJSON(c, 0, "", "rec_address_id地址记录不存在")
			return
		}
		order.RecName = address.Name
		order.RecCountry = address.Country
		order.RecProvince = address.Province
		order.RecCity = address.City
		order.RecAddress = address.Address
		order.RecMobile = address.Mobile
	}

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
	method_slice := map[string]bool{"order_pay_confirm": true, "order_cancel": true, "order_collect": true, "order_deliver": true, "order_arrive": true, "order_sign": true}
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
	/* 创建map */
	shipMap := map[string]int{"order_collect": 1, "order_deliver": 2, "order_arrive": 3, "order_sign": 4}
	// 开始事务
	tx := Config.DB.Begin()
	switch order_method {
	case "order_pay_confirm": // 确认支付order_status =1 更新付款时间
		Models.UpdateOrderPayTime(&order, id)
		order.OrderStatus = 1
		order.PayStatus = 1
		orderActionNote = user.UserName + ":" + user.RealName + "修改了订单支付状态为:已付款";

	case "order_cancel": // 取消订单 order_status =2
		order.OrderStatus = 2
		orderActionNote = user.UserName + ":" + user.RealName + "修改了订单状态为:取消订单";
	case "order_collect", "order_deliver", "order_arrive", "order_sign": // 揽件1/发走2/已达到3/签收订单4 ship_status=4 记录物流 未支付 未发货的不能签收
		if order.PayStatus < 1 {
			ApiHelpers.RespondJSON(c, 0, "", "未支付的快递无法操作")
			return
		}
		// 避免重复物流操作
		if shipMap[order_method] == order.ShipStatus {
			ApiHelpers.RespondJSON(c, 0, "", order_method+"该快递操作已存在,无需重复操作")
			return
		}
		// 验证签收物流程序 已揽件->已发货(运输中)->已到达->已签收
		var saction Models.ShipAction
		if shipMap[order_method] == 3 || shipMap[order_method] == 2 {
			// 查询是否走了上一个流程
			shipStatus := shipMap[order_method] - 1
			err := Models.GetOneShip(&saction, id, shipStatus)
			if err != nil {
				ApiHelpers.RespondJSON(c, 0, saction, "请确认流程是否正确：已揽件->已发货->已到达->已签收")
				return
			}
		}
		order.ShipStatus = shipMap[order_method] // 更新订单状态为签收
		orderActionNote = user.UserName + ":" + user.RealName + "修改了物流状态为:已签收";
		// 记录物流信息
		c.ShouldBind(&saction)
		if shipMap[order_method] == 1 {
			saction.ActionNote = "快递员：" + user.RealName + "已揽件";
		} else if shipMap[order_method] == 2 {
			saction.ActionNote = "快递员：" + user.RealName + "已发货(运输中)";
		} else if shipMap[order_method] == 3 {
			saction.ActionNote = "订单已到达取件点,请保持收件人电话畅通)";
		} else {
			saction.ActionNote = "收件人已签收,欢迎再次使用。"
		}
		saction.CreatorId = uInt
		saction.ShipStatus = shipMap[order_method] // 标记物流状态
		saction.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
		err := Models.AddNewShipAction(&saction)
		if err != nil {
			tx.Rollback()
		}
	}
	// 记录order_action
	var oaction Models.OrderAction
	c.ShouldBind(&oaction)
	oaction.ActionNote = orderActionNote
	oaction.CreatorId = uInt
	err = Models.AddNewOrderAction(&oaction)
	if err != nil {
		tx.Rollback()
	}
	c.Bind(&order)
	err = Models.PutOneOrder(&order, id)
	if err != nil {
		tx.Rollback()
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

	order_code := c.PostForm("order_code")
	code_slice := map[string]bool{"wechat": true, "alipay": true, "cash ": true}
	if _, ok := code_slice[order_code]; !ok {
		ApiHelpers.RespondJSON(c, 0, "", "order_code参数非法")
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
