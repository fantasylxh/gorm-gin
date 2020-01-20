package Models

import (
	"gorm-gin/Config"
	_ "github.com/go-sql-driver/mysql"
)


func AddNewFeedback(b *FeedBack) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}



