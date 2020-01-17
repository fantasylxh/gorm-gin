package Models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-gin/Config"
)

func GetAllCategory(b *[]Category) (err error) {
	if err = Config.DB.Order("seq asc, id desc").Find(b).Error; err != nil {
		return err
	}
	return nil
}
