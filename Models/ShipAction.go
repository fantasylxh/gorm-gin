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
