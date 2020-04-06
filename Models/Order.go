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
		return db.Where("order_sn like ?", "%"+keyword+"%").Or("name like ?", "%"+keyword+"%").Or("rec_name like ?", "%"+keyword+"%")
	}
}

//订单状态 待支付 已支付
func OrderStatus(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(GetByKeyword("")).Where("order_status in (?)", status)
	}
}

//物流状态
func ShipStatus(status []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Scopes(GetByKeyword("")).Where("ship_status in (?)", status)
	}
}

//订单来源 0:中国到孟加拉 1:孟加拉到中国 3:孟加拉到孟加拉国内
func OrderFrom(order_from []string) func(db *gorm.DB) *gorm.DB {
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
	var pageSize int = 2

	// 默认为签收,获取取指page，指定pagesize的记录,
	if conditions["q_type"] == "1" {
		// 验证当前用户角色
		if conditions["role_id"] == "24" { // 快递管理员-孟加拉, 则签收订单为 所有来着中国滴订单 order_from =0
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]string{"0"})).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
			// 获取总条数
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]string{"0"})).Count(&count)
		} else if conditions["role_id"] == "25" { // 快递管理员-中国 则签收订单为 所有来着孟加拉达卡滴订单 order_from =1
			Config.DB.Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]string{"1"})).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
			// 获取总条数
			Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderFrom([]string{"0"})).Count(&count)
		}
	}
	// 查询发走的订单
	if conditions["q_type"] == "2" {
		Config.DB.Scopes(GetByKeyword(conditions["keyword"]), OrderWithUid(conditions["uid"])).Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("created_at desc").Find(&b)
		// 获取总条数
		Config.DB.Model(&Order{}).Scopes(GetByKeyword(conditions["keyword"]), OrderWithUid(conditions["uid"])).Count(&count)
	}

	/*	if conditions["order_status"] != "" {
		if err = Config.DB.Scopes(OrderStatus([]string{conditions["order_status"]})).Order("id desc").Find(b).Error; err != nil {
			return err, 0
		}

		fmt.Println(b, "总数：", count)
	}*/

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
