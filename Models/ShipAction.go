package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func AddNewShipAction(b *ShipAction) (err error) {

	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneShip(b *ShipAction, order_id string, ship_status int) (err error) {
	if err := Config.DB.Where("order_id = ? and ship_status = ?", order_id, ship_status).First(b).Error; err != nil {
		return err
	}
	return nil
}
