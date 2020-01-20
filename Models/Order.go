package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
	"time"
)

func GetAllOrder(b *[]Order) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
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
