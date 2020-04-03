package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func GetOnePaymentCode(b *PaymentCode, code string) (err error) {
	if err = Config.DB.Where("status =1 and order_code = ?", code).Order("id desc").First(b).Error; err != nil {
		return err
	}
	return nil
}
