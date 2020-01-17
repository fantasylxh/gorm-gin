package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func GetOneCategoryPrice(b *CategoryPrice, id string, weight string) (err error) {
	if err := Config.DB.Where("category_id = ? and min_kg <= ? AND max_kg >= ?", id, weight, weight).First(b).Error; err != nil {
		return err
	}
	return nil
}
