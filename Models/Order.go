package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gorm-gin/Config"
	"strconv"
	"time"
)

//关键字查询
func GetByKeyword(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("order_sn like ? or  name like ? or rec_name like ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
}

//订单状态 待支付 已支付
func OrderStatus(status []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("order_status in (?)", status)
	}
}

//订单支付状态 待支付 已支付
func PayStatus(status []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("pay_status in (?)", status)
	}
}

//物流状态0，未发货；1，已揽件；2，已发货(运输中)；3，已到达；4，已签收',
func ShipStatus(status []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("ship_status in (?)", status)
	}
}

//订单来源 0:中国到孟加拉 1:孟加拉到中国 3:孟加拉到孟加拉国内
func OrderFrom(order_from []int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("order_from in (?)", order_from)
	}
}

// 根据UID 获取创建的订单
func OrderWithUid(uid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("creator_id = ?", uid)
	}
}

//查询类型 1：签收 2：发走 默认1为签收

func GetAllOrder(b *[]Order, conditions map[string]string) (err error, count int) {
	var page string = conditions["page"]
	pageInt, err := strconv.Atoi(page)
	var pageSize int = 10 //每页显示数量
	//全局基础查询条件
	var intOrderArr = []int{0, 1, 2} // 查询初始化订单0：待支付,1；已支付, 2：已取消',
	order_status, _ := strconv.Atoi(conditions["order_status"])
	if conditions["order_status"] != "" {
		intOrderArr = []int{order_status}
	}
	var intPayArr = []int{0, 1} // 查询初始化订单支付状态 0：未支付  1：已支付
	pay_status, _ := strconv.Atoi(conditions["pay_status"])
	if conditions["pay_status"] != "" {
		intPayArr = []int{pay_status}
	}
	var intShipArr = []int{0, 1, 2, 3, 4} // 查询初始化物流0，未发货；1，已揽件；2，已发货(运输中)；3，已到达；4，已签收',
	ship_status, _ := strconv.Atoi(conditions["ship_status"])
	if conditions["ship_status"] != "" {
		intShipArr = []int{ship_status}
	}
	// 默认为签收,获取取指page，指定pagesize的记录,
	if conditions["q_type"] == "1" {
		// 验证当前用户角色
		if conditions["role_id"] == "24" { // 快递管理员-孟加拉, 则签收订单为 所有来着中国滴订单 order_from =0
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]int{0}), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
			// 获取总条数
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]int{0}), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Count(&count)
		} else if conditions["role_id"] == "25" { // 快递管理员-中国 则签收订单为 所有来着孟加拉达卡滴订单 order_from =1
			Config.DB.Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]int{1}), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
			// 获取总条数
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]int{1}), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Count(&count)
		}
	}
	// 查询发走的订单
	if conditions["q_type"] == "2" {
		Config.DB.Scopes(GetByKeyword(conditions["keyword"]), OrderWithUid(conditions["uid"]), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
		// 获取总条数
		Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderWithUid(conditions["uid"]), PayStatus(intPayArr), OrderStatus(intOrderArr), ShipStatus(intShipArr)).Count(&count)
	}

	return nil, count
}

func AddNewOrder(b *Order) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOrder(b *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderShip(b *[]ShipAction, id string) (err error) {
	if err := Config.DB.Order("id desc").Where("order_id = ?", id).Find(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneOrder(b *Order, id string) (err error) {
	fmt.Println(b)
	Config.DB.Omit("order_sn").Save(b)
	return nil
}

func DeleteOrder(b *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}

// 更新订单付款时间
func UpdateOrderPayTime(b *Order, id string) (err error) {
	Config.DB.Model(b).Where("id = ?", id).Omit("order_sn").Update("pay_time", time.Now())
	return nil
}
