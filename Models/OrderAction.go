package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func AddNewOrderAction(b *OrderAction) (err error) {

	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}
